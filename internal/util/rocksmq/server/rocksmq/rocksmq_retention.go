// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License.

package rocksmq

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	rocksdbkv "github.com/milvus-io/milvus/internal/kv/rocksdb"
	"github.com/milvus-io/milvus/internal/log"
	"github.com/tecbot/gorocksdb"
	"go.uber.org/zap"
)

var RocksmqRetentionTimeInMinutes int64 = 1
var RocksmqRetentionSizeInMB int64 = 1024
var TickerTimeInMinutes int64 = 1

type topicPageInfo struct {
	pageEndID   []UniqueID
	pageMsgSize map[int64]UniqueID
}

type topicAckedInfo struct {
	topicBeginID    UniqueID
	consumerBeginID map[string]UniqueID
	// TODO(yukun): may need to delete ackedTs
	ackedTs   map[UniqueID]UniqueID
	ackedSize int64
}

type retentionInfo struct {
	topics    []string
	consumers []*Consumer
	pageInfo  map[string]*topicPageInfo
	ackedInfo map[string]*topicAckedInfo
	// Key is last_retention_time/${topic}
	lastRetentionTime map[string]int64

	kv *rocksdbkv.RocksdbKV
	db *gorocksdb.DB
}

func (ri *retentionInfo) loadRetentionInfo(kv *rocksdbkv.RocksdbKV, db *gorocksdb.DB) error {
	// Get topic from topic begin id
	beginIDKeys, _, err := kv.LoadWithPrefix(TopicBeginIDTitle)
	if err != nil {
		return err
	}
	for _, key := range beginIDKeys {
		topic := key[len(TopicBeginIDTitle):]
		ri.topics = append(ri.topics, topic)
		topicMu.Store(topic, new(sync.Mutex))
	}

	for _, topic := range ri.topics {
		// Load all page infos
		pageEndID := make([]UniqueID, 0)
		pageMsgSize := make(map[int64]UniqueID)

		fixedPageSizeKey, err := constructKey(PageMsgSizeTitle, topic)
		if err != nil {
			return err
		}
		pageMsgSizePrefix := fixedPageSizeKey + "/"
		pageMsgSizeKeys, pageMsgSizeVals, err := kv.LoadWithPrefix(pageMsgSizePrefix)
		if err != nil {
			return err
		}
		for i, key := range pageMsgSizeKeys {
			endID, err := strconv.ParseInt(key[FixedChannelNameLen+1:], 10, 64)
			if err != nil {
				return err
			}
			pageEndID = append(pageEndID, endID)

			msgSize, err := strconv.ParseInt(pageMsgSizeVals[i], 10, 64)
			if err != nil {
				return err
			}
			pageMsgSize[endID] = msgSize
		}
		topicPageInfo := &topicPageInfo{
			pageEndID:   pageEndID,
			pageMsgSize: pageMsgSize,
		}

		// Load all acked infos
		ackedTs := make(map[UniqueID]UniqueID)

		topicBeginIDKey := TopicBeginIDTitle + topic
		topicBeginIDVal, err := kv.Load(topicBeginIDKey)
		if err != nil {
			return err
		}
		topicBeginID, err := strconv.ParseInt(topicBeginIDVal, 10, 64)
		if err != nil {
			return err
		}

		consumerBeginIDs := make(map[string]UniqueID)
		for _, consumer := range ri.consumers {
			beginIDKey := BeginIDTitle + consumer.Topic + "/" + consumer.GroupName
			beginIDVal, err := kv.Load(beginIDKey)
			if err != nil {
				return err
			}
			beginID, err := strconv.Atoi(beginIDVal)
			if err != nil {
				return err
			}
			consumerBeginIDs[consumer.GroupName] = UniqueID(beginID)
		}

		ackedTsPrefix, err := constructKey(AckedTsTitle, topic)
		if err != nil {
			return err
		}
		keys, vals, err := kv.LoadWithPrefix(ackedTsPrefix)
		if err != nil {
			return err
		}
		if len(keys) != len(vals) {
			return errors.New("LoadWithPrefix return unequal value length of keys and values")
		}

		for i, key := range keys {
			offset := FixedChannelNameLen + 1
			ackedID, err := strconv.ParseInt((key)[offset:], 10, 64)
			if err != nil {
				log.Debug("RocksMQ: parse int " + key[offset:] + " failed")
				return err
			}
			ts, err := strconv.ParseInt(vals[i], 10, 64)
			if err != nil {
				return err
			}
			ackedTs[ackedID] = ts
		}

		ackedSizeKey := AckedSizeTitle + topic
		ackedSizeVal, err := kv.Load(ackedSizeKey)
		if err != nil {
			return err
		}
		ackedSize, err := strconv.ParseInt(ackedSizeVal, 10, 64)
		if err != nil {
			return err
		}

		ackedInfo := &topicAckedInfo{
			topicBeginID:    topicBeginID,
			consumerBeginID: consumerBeginIDs,
			ackedTs:         ackedTs,
			ackedSize:       ackedSize,
		}

		//Load last retention timestamp
		lastRetentionTsKey := LastRetTsTitle + topic
		lastRetentionTsVal, err := kv.Load(lastRetentionTsKey)
		if err != nil {
			return err
		}
		lastRetentionTs, err := strconv.ParseInt(lastRetentionTsVal, 10, 64)
		if err != nil {
			return err
		}

		ri.ackedInfo[topic] = ackedInfo
		ri.pageInfo[topic] = topicPageInfo
		ri.lastRetentionTime[topic] = lastRetentionTs
	}
	ri.kv = kv
	ri.db = db

	return nil
}

func (ri retentionInfo) retention() error {
	log.Debug("In retention")
	ticker := time.NewTicker(time.Duration(TickerTimeInMinutes * int64(time.Minute) / 10))
	done := make(chan bool)

	for {
		select {
		case <-done:
			break
		case t := <-ticker.C:
			timeNow := t.Unix()
			checkTime := RocksmqRetentionTimeInMinutes * 60 / 10
			log.Debug("In ticker: ", zap.Any("ticker", timeNow))
			for k, v := range ri.lastRetentionTime {
				if v+checkTime > timeNow {
					err := ri.expiredCleanUp(k)
					if err != nil {
						panic(err)
					}
				}
			}
		}
	}
}

func (ri retentionInfo) expiredCleanUp(topic string) error {
	log.Debug("In expiredCleanUp")
	ackedInfo := ri.ackedInfo[topic]

	ll, ok := topicMu.Load(topic)
	if !ok {
		return fmt.Errorf("topic name = %s not exist", topic)
	}
	lock, ok := ll.(*sync.Mutex)
	if !ok {
		return fmt.Errorf("get mutex failed, topic name = %s", topic)
	}
	lock.Lock()
	defer lock.Unlock()

	readOpts := gorocksdb.NewDefaultReadOptions()
	defer readOpts.Destroy()
	readOpts.SetPrefixSameAsStart(true)
	iter := ri.kv.DB.NewIterator(readOpts)
	defer iter.Close()
	ackedTsPrefix, err := constructKey(AckedTsTitle, topic)
	if err != nil {
		return err
	}
	iter.Seek([]byte(ackedTsPrefix))
	if !iter.Valid() {
		return nil
	}
	var startID, endID UniqueID
	startID, err = strconv.ParseInt(string(iter.Key().Data())[FixedChannelNameLen+1:], 10, 64)
	if err != nil {
		return err
	}
	pageRetentionOffset := 0
	for i, pageEndID := range ri.pageInfo[topic].pageEndID {
		// Clean by RocksmqRetentionTimeInMinutes
		if msgExpiredCheck(ackedInfo.ackedTs[pageEndID]) {
			// All of the page expired, set the pageEndID to current endID
			endID = pageEndID
			fixedAckedTsKey, err := constructKey(AckedTsTitle, topic)
			if err != nil {
				return err
			}
			newKey := fixedAckedTsKey + "/" + strconv.Itoa(int(pageEndID))
			iter.Seek([]byte(newKey))
			pageRetentionOffset = i + 1
		}
	}

	pageEndID := endID
	// The end msg of the page is not expired, find the last expired msg in this page
	for ; iter.Valid(); iter.Next() {
		ackedTs, err := strconv.ParseInt(string(iter.Value().Data()), 10, 64)
		if err != nil {
			return err
		}
		if msgExpiredCheck(ackedTs) {
			endID, err = strconv.ParseInt(string(iter.Key().Data())[FixedChannelNameLen+1:], 10, 64)
			if err != nil {
				return err
			}
		} else {
			break
		}
	}

	// Delete page message size in rocksdb_kv
	pageStartID := ri.pageInfo[topic].pageEndID[0]
	fixedPageSizeKey, err := constructKey(PageMsgSizeTitle, topic)
	if err != nil {
		return err
	}
	pageStartKey := fixedPageSizeKey + "/" + strconv.Itoa(int(pageStartID))
	pageEndKey := fixedPageSizeKey + "/" + strconv.Itoa(int(pageEndID))
	pageWriteBatch := gorocksdb.NewWriteBatch()
	defer pageWriteBatch.Clear()
	pageWriteBatch.DeleteRange([]byte(pageStartKey), []byte(pageEndKey))
	ri.kv.DB.Write(gorocksdb.NewDefaultWriteOptions(), pageWriteBatch)

	ri.pageInfo[topic].pageEndID = ri.pageInfo[topic].pageEndID[pageRetentionOffset:]

	// Delete acked_ts in rocksdb_kv
	fixedAckedTsTitle, err := constructKey(AckedTsTitle, topic)
	if err != nil {
		return err
	}
	ackedStartIDKey := fixedAckedTsTitle + "/" + strconv.Itoa(int(startID))
	ackedEndIDKey := fixedAckedTsTitle + "/" + strconv.Itoa(int(endID))
	ackedTsWriteBatch := gorocksdb.NewWriteBatch()
	defer ackedTsWriteBatch.Clear()
	ackedTsWriteBatch.DeleteRange([]byte(ackedStartIDKey), []byte(ackedEndIDKey))
	ri.kv.DB.Write(gorocksdb.NewDefaultWriteOptions(), ackedTsWriteBatch)

	// Update acked_size in rocksdb_kv

	// Update last retention ts
	lastRetentionTsKey := LastRetTsTitle + topic
	err = ri.kv.Save(lastRetentionTsKey, strconv.FormatInt(time.Now().Unix(), 10))
	if err != nil {
		return err
	}

	log.Debug("Delete finish!")

	return DeleteMessages(ri.db, topic, startID, endID)
}

func DeleteMessages(db *gorocksdb.DB, topic string, startID, endID UniqueID) error {
	// Delete msg by range of startID and endID
	startKey, err := combKey(topic, startID)
	if err != nil {
		log.Debug("RocksMQ: combKey(" + topic + "," + strconv.FormatInt(startID, 10) + ")")
		return err
	}
	endKey, err := combKey(topic, endID)
	if err != nil {
		log.Debug("RocksMQ: combKey(" + topic + "," + strconv.FormatInt(endID, 10) + ")")
		return err
	}

	writeBatch := gorocksdb.NewWriteBatch()
	defer writeBatch.Clear()
	writeBatch.DeleteRange([]byte(startKey), []byte(endKey))
	err = db.Write(gorocksdb.NewDefaultWriteOptions(), writeBatch)
	if err != nil {
		return err
	}

	return nil
}

func msgExpiredCheck(ackedTs int64) bool {
	return ackedTs+RocksmqRetentionTimeInMinutes*60 < time.Now().Unix()
}
