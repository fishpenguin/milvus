package writenode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParamTable_WriteNode(t *testing.T) {

	Params.Init()

	t.Run("Test PulsarAddress", func(t *testing.T) {
		address := Params.PulsarAddress
		split := strings.Split(address, ":")
		assert.Equal(t, split[0], "pulsar")
		assert.Equal(t, split[len(split)-1], "6650")
	})

	t.Run("Test WriteNodeID", func(t *testing.T) {
		id := Params.WriteNodeID
		assert.Equal(t, id, UniqueID(3))
	})

	t.Run("Test insertChannelRange", func(t *testing.T) {
		channelRange := Params.InsertChannelRange
		assert.Equal(t, len(channelRange), 2)
		assert.Equal(t, channelRange[0], 0)
		assert.Equal(t, channelRange[1], 2)
	})

	t.Run("Test insertMsgStreamReceiveBufSize", func(t *testing.T) {
		bufSize := Params.InsertReceiveBufSize
		assert.Equal(t, bufSize, int64(1024))
	})

	t.Run("Test insertPulsarBufSize", func(t *testing.T) {
		bufSize := Params.InsertPulsarBufSize
		assert.Equal(t, bufSize, int64(1024))
	})

	t.Run("Test flowGraphMaxQueueLength", func(t *testing.T) {
		length := Params.FlowGraphMaxQueueLength
		assert.Equal(t, length, int32(1024))
	})

	t.Run("Test flowGraphMaxParallelism", func(t *testing.T) {
		maxParallelism := Params.FlowGraphMaxParallelism
		assert.Equal(t, maxParallelism, int32(1024))
	})

	t.Run("Test insertChannelNames", func(t *testing.T) {
		names := Params.InsertChannelNames
		assert.Equal(t, len(names), 2)
		assert.Equal(t, names[0], "insert-0")
		assert.Equal(t, names[1], "insert-1")
	})

	t.Run("Test msgChannelSubName", func(t *testing.T) {
		name := Params.MsgChannelSubName
		assert.Equal(t, name, "writeNode-3")
	})

	t.Run("Test timeTickChannelName", func(t *testing.T) {
		name := Params.WriteNodeTimeTickChannelName
		assert.Equal(t, name, "writeNodeTimeTick")
	})
}