// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pulsar.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pulsar.proto

It has these top-level messages:
	Status
	SegmentRecord
	VectorRowRecord
	AttrRecord
	VectorRecord
	VectorParam
	FieldValue
	PulsarMessage
	PulsarMessages
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ErrorCode int32

const (
	ErrorCode_SUCCESS                 ErrorCode = 0
	ErrorCode_UNEXPECTED_ERROR        ErrorCode = 1
	ErrorCode_CONNECT_FAILED          ErrorCode = 2
	ErrorCode_PERMISSION_DENIED       ErrorCode = 3
	ErrorCode_COLLECTION_NOT_EXISTS   ErrorCode = 4
	ErrorCode_ILLEGAL_ARGUMENT        ErrorCode = 5
	ErrorCode_ILLEGAL_DIMENSION       ErrorCode = 7
	ErrorCode_ILLEGAL_INDEX_TYPE      ErrorCode = 8
	ErrorCode_ILLEGAL_COLLECTION_NAME ErrorCode = 9
	ErrorCode_ILLEGAL_TOPK            ErrorCode = 10
	ErrorCode_ILLEGAL_ROWRECORD       ErrorCode = 11
	ErrorCode_ILLEGAL_VECTOR_ID       ErrorCode = 12
	ErrorCode_ILLEGAL_SEARCH_RESULT   ErrorCode = 13
	ErrorCode_FILE_NOT_FOUND          ErrorCode = 14
	ErrorCode_META_FAILED             ErrorCode = 15
	ErrorCode_CACHE_FAILED            ErrorCode = 16
	ErrorCode_CANNOT_CREATE_FOLDER    ErrorCode = 17
	ErrorCode_CANNOT_CREATE_FILE      ErrorCode = 18
	ErrorCode_CANNOT_DELETE_FOLDER    ErrorCode = 19
	ErrorCode_CANNOT_DELETE_FILE      ErrorCode = 20
	ErrorCode_BUILD_INDEX_ERROR       ErrorCode = 21
	ErrorCode_ILLEGAL_NLIST           ErrorCode = 22
	ErrorCode_ILLEGAL_METRIC_TYPE     ErrorCode = 23
	ErrorCode_OUT_OF_MEMORY           ErrorCode = 24
)

var ErrorCode_name = map[int32]string{
	0:  "SUCCESS",
	1:  "UNEXPECTED_ERROR",
	2:  "CONNECT_FAILED",
	3:  "PERMISSION_DENIED",
	4:  "COLLECTION_NOT_EXISTS",
	5:  "ILLEGAL_ARGUMENT",
	7:  "ILLEGAL_DIMENSION",
	8:  "ILLEGAL_INDEX_TYPE",
	9:  "ILLEGAL_COLLECTION_NAME",
	10: "ILLEGAL_TOPK",
	11: "ILLEGAL_ROWRECORD",
	12: "ILLEGAL_VECTOR_ID",
	13: "ILLEGAL_SEARCH_RESULT",
	14: "FILE_NOT_FOUND",
	15: "META_FAILED",
	16: "CACHE_FAILED",
	17: "CANNOT_CREATE_FOLDER",
	18: "CANNOT_CREATE_FILE",
	19: "CANNOT_DELETE_FOLDER",
	20: "CANNOT_DELETE_FILE",
	21: "BUILD_INDEX_ERROR",
	22: "ILLEGAL_NLIST",
	23: "ILLEGAL_METRIC_TYPE",
	24: "OUT_OF_MEMORY",
}
var ErrorCode_value = map[string]int32{
	"SUCCESS":                 0,
	"UNEXPECTED_ERROR":        1,
	"CONNECT_FAILED":          2,
	"PERMISSION_DENIED":       3,
	"COLLECTION_NOT_EXISTS":   4,
	"ILLEGAL_ARGUMENT":        5,
	"ILLEGAL_DIMENSION":       7,
	"ILLEGAL_INDEX_TYPE":      8,
	"ILLEGAL_COLLECTION_NAME": 9,
	"ILLEGAL_TOPK":            10,
	"ILLEGAL_ROWRECORD":       11,
	"ILLEGAL_VECTOR_ID":       12,
	"ILLEGAL_SEARCH_RESULT":   13,
	"FILE_NOT_FOUND":          14,
	"META_FAILED":             15,
	"CACHE_FAILED":            16,
	"CANNOT_CREATE_FOLDER":    17,
	"CANNOT_CREATE_FILE":      18,
	"CANNOT_DELETE_FOLDER":    19,
	"CANNOT_DELETE_FILE":      20,
	"BUILD_INDEX_ERROR":       21,
	"ILLEGAL_NLIST":           22,
	"ILLEGAL_METRIC_TYPE":     23,
	"OUT_OF_MEMORY":           24,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DataType int32

const (
	DataType_NONE          DataType = 0
	DataType_BOOL          DataType = 1
	DataType_INT8          DataType = 2
	DataType_INT16         DataType = 3
	DataType_INT32         DataType = 4
	DataType_INT64         DataType = 5
	DataType_FLOAT         DataType = 10
	DataType_DOUBLE        DataType = 11
	DataType_STRING        DataType = 20
	DataType_VECTOR_BINARY DataType = 100
	DataType_VECTOR_FLOAT  DataType = 101
)

var DataType_name = map[int32]string{
	0:   "NONE",
	1:   "BOOL",
	2:   "INT8",
	3:   "INT16",
	4:   "INT32",
	5:   "INT64",
	10:  "FLOAT",
	11:  "DOUBLE",
	20:  "STRING",
	100: "VECTOR_BINARY",
	101: "VECTOR_FLOAT",
}
var DataType_value = map[string]int32{
	"NONE":          0,
	"BOOL":          1,
	"INT8":          2,
	"INT16":         3,
	"INT32":         4,
	"INT64":         5,
	"FLOAT":         10,
	"DOUBLE":        11,
	"STRING":        20,
	"VECTOR_BINARY": 100,
	"VECTOR_FLOAT":  101,
}

func (x DataType) String() string {
	return proto.EnumName(DataType_name, int32(x))
}
func (DataType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type OpType int32

const (
	OpType_Insert     OpType = 0
	OpType_Delete     OpType = 1
	OpType_Search     OpType = 2
	OpType_TimeSync   OpType = 3
	OpType_Key2Seg    OpType = 4
	OpType_Statistics OpType = 5
)

var OpType_name = map[int32]string{
	0: "Insert",
	1: "Delete",
	2: "Search",
	3: "TimeSync",
	4: "Key2Seg",
	5: "Statistics",
}
var OpType_value = map[string]int32{
	"Insert":     0,
	"Delete":     1,
	"Search":     2,
	"TimeSync":   3,
	"Key2Seg":    4,
	"Statistics": 5,
}

func (x OpType) String() string {
	return proto.EnumName(OpType_name, int32(x))
}
func (OpType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Status struct {
	ErrorCode ErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,enum=pb.ErrorCode" json:"error_code,omitempty"`
	Reason    string    `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Status) GetErrorCode() ErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return ErrorCode_SUCCESS
}

func (m *Status) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type SegmentRecord struct {
	SegInfo []string `protobuf:"bytes,1,rep,name=seg_info,json=segInfo" json:"seg_info,omitempty"`
}

func (m *SegmentRecord) Reset()                    { *m = SegmentRecord{} }
func (m *SegmentRecord) String() string            { return proto.CompactTextString(m) }
func (*SegmentRecord) ProtoMessage()               {}
func (*SegmentRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SegmentRecord) GetSegInfo() []string {
	if m != nil {
		return m.SegInfo
	}
	return nil
}

type VectorRowRecord struct {
	FloatData  []float32 `protobuf:"fixed32,1,rep,packed,name=float_data,json=floatData" json:"float_data,omitempty"`
	BinaryData []byte    `protobuf:"bytes,2,opt,name=binary_data,json=binaryData,proto3" json:"binary_data,omitempty"`
}

func (m *VectorRowRecord) Reset()                    { *m = VectorRowRecord{} }
func (m *VectorRowRecord) String() string            { return proto.CompactTextString(m) }
func (*VectorRowRecord) ProtoMessage()               {}
func (*VectorRowRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *VectorRowRecord) GetFloatData() []float32 {
	if m != nil {
		return m.FloatData
	}
	return nil
}

func (m *VectorRowRecord) GetBinaryData() []byte {
	if m != nil {
		return m.BinaryData
	}
	return nil
}

type AttrRecord struct {
	Int32Value  []int32   `protobuf:"varint,1,rep,packed,name=int32_value,json=int32Value" json:"int32_value,omitempty"`
	Int64Value  []int64   `protobuf:"varint,2,rep,packed,name=int64_value,json=int64Value" json:"int64_value,omitempty"`
	FloatValue  []float32 `protobuf:"fixed32,3,rep,packed,name=float_value,json=floatValue" json:"float_value,omitempty"`
	DoubleValue []float64 `protobuf:"fixed64,4,rep,packed,name=double_value,json=doubleValue" json:"double_value,omitempty"`
}

func (m *AttrRecord) Reset()                    { *m = AttrRecord{} }
func (m *AttrRecord) String() string            { return proto.CompactTextString(m) }
func (*AttrRecord) ProtoMessage()               {}
func (*AttrRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AttrRecord) GetInt32Value() []int32 {
	if m != nil {
		return m.Int32Value
	}
	return nil
}

func (m *AttrRecord) GetInt64Value() []int64 {
	if m != nil {
		return m.Int64Value
	}
	return nil
}

func (m *AttrRecord) GetFloatValue() []float32 {
	if m != nil {
		return m.FloatValue
	}
	return nil
}

func (m *AttrRecord) GetDoubleValue() []float64 {
	if m != nil {
		return m.DoubleValue
	}
	return nil
}

type VectorRecord struct {
	Records []*VectorRowRecord `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
}

func (m *VectorRecord) Reset()                    { *m = VectorRecord{} }
func (m *VectorRecord) String() string            { return proto.CompactTextString(m) }
func (*VectorRecord) ProtoMessage()               {}
func (*VectorRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *VectorRecord) GetRecords() []*VectorRowRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

type VectorParam struct {
	Json      string        `protobuf:"bytes,1,opt,name=json" json:"json,omitempty"`
	RowRecord *VectorRecord `protobuf:"bytes,2,opt,name=row_record,json=rowRecord" json:"row_record,omitempty"`
}

func (m *VectorParam) Reset()                    { *m = VectorParam{} }
func (m *VectorParam) String() string            { return proto.CompactTextString(m) }
func (*VectorParam) ProtoMessage()               {}
func (*VectorParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *VectorParam) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

func (m *VectorParam) GetRowRecord() *VectorRecord {
	if m != nil {
		return m.RowRecord
	}
	return nil
}

type FieldValue struct {
	FieldName    string        `protobuf:"bytes,1,opt,name=field_name,json=fieldName" json:"field_name,omitempty"`
	Type         DataType      `protobuf:"varint,2,opt,name=type,enum=pb.DataType" json:"type,omitempty"`
	AttrRecord   *AttrRecord   `protobuf:"bytes,3,opt,name=attr_record,json=attrRecord" json:"attr_record,omitempty"`
	VectorRecord *VectorRecord `protobuf:"bytes,4,opt,name=vector_record,json=vectorRecord" json:"vector_record,omitempty"`
}

func (m *FieldValue) Reset()                    { *m = FieldValue{} }
func (m *FieldValue) String() string            { return proto.CompactTextString(m) }
func (*FieldValue) ProtoMessage()               {}
func (*FieldValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FieldValue) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

func (m *FieldValue) GetType() DataType {
	if m != nil {
		return m.Type
	}
	return DataType_NONE
}

func (m *FieldValue) GetAttrRecord() *AttrRecord {
	if m != nil {
		return m.AttrRecord
	}
	return nil
}

func (m *FieldValue) GetVectorRecord() *VectorRecord {
	if m != nil {
		return m.VectorRecord
	}
	return nil
}

type PulsarMessage struct {
	CollectionName string         `protobuf:"bytes,1,opt,name=collection_name,json=collectionName" json:"collection_name,omitempty"`
	Fields         []*FieldValue  `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty"`
	EntityId       int64          `protobuf:"varint,3,opt,name=entity_id,json=entityId" json:"entity_id,omitempty"`
	PartitionTag   string         `protobuf:"bytes,4,opt,name=partition_tag,json=partitionTag" json:"partition_tag,omitempty"`
	VectorParam    *VectorParam   `protobuf:"bytes,5,opt,name=vector_param,json=vectorParam" json:"vector_param,omitempty"`
	Segments       *SegmentRecord `protobuf:"bytes,6,opt,name=segments" json:"segments,omitempty"`
	Timestamp      int64          `protobuf:"varint,7,opt,name=timestamp" json:"timestamp,omitempty"`
	ClientId       int64          `protobuf:"varint,8,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	MsgType        OpType         `protobuf:"varint,9,opt,name=msg_type,json=msgType,enum=pb.OpType" json:"msg_type,omitempty"`
}

func (m *PulsarMessage) Reset()                    { *m = PulsarMessage{} }
func (m *PulsarMessage) String() string            { return proto.CompactTextString(m) }
func (*PulsarMessage) ProtoMessage()               {}
func (*PulsarMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PulsarMessage) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *PulsarMessage) GetFields() []*FieldValue {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *PulsarMessage) GetEntityId() int64 {
	if m != nil {
		return m.EntityId
	}
	return 0
}

func (m *PulsarMessage) GetPartitionTag() string {
	if m != nil {
		return m.PartitionTag
	}
	return ""
}

func (m *PulsarMessage) GetVectorParam() *VectorParam {
	if m != nil {
		return m.VectorParam
	}
	return nil
}

func (m *PulsarMessage) GetSegments() *SegmentRecord {
	if m != nil {
		return m.Segments
	}
	return nil
}

func (m *PulsarMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *PulsarMessage) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *PulsarMessage) GetMsgType() OpType {
	if m != nil {
		return m.MsgType
	}
	return OpType_Insert
}

type PulsarMessages struct {
	CollectionName string           `protobuf:"bytes,1,opt,name=collection_name,json=collectionName" json:"collection_name,omitempty"`
	Fields         []*FieldValue    `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty"`
	EntityId       []int64          `protobuf:"varint,3,rep,packed,name=entity_id,json=entityId" json:"entity_id,omitempty"`
	PartitionTag   string           `protobuf:"bytes,4,opt,name=partition_tag,json=partitionTag" json:"partition_tag,omitempty"`
	VectorParam    []*VectorParam   `protobuf:"bytes,5,rep,name=vector_param,json=vectorParam" json:"vector_param,omitempty"`
	Segments       []*SegmentRecord `protobuf:"bytes,6,rep,name=segments" json:"segments,omitempty"`
	Timestamp      []int64          `protobuf:"varint,7,rep,packed,name=timestamp" json:"timestamp,omitempty"`
	ClientId       []int64          `protobuf:"varint,8,rep,packed,name=client_id,json=clientId" json:"client_id,omitempty"`
	MsgType        OpType           `protobuf:"varint,9,opt,name=msg_type,json=msgType,enum=pb.OpType" json:"msg_type,omitempty"`
}

func (m *PulsarMessages) Reset()                    { *m = PulsarMessages{} }
func (m *PulsarMessages) String() string            { return proto.CompactTextString(m) }
func (*PulsarMessages) ProtoMessage()               {}
func (*PulsarMessages) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PulsarMessages) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *PulsarMessages) GetFields() []*FieldValue {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *PulsarMessages) GetEntityId() []int64 {
	if m != nil {
		return m.EntityId
	}
	return nil
}

func (m *PulsarMessages) GetPartitionTag() string {
	if m != nil {
		return m.PartitionTag
	}
	return ""
}

func (m *PulsarMessages) GetVectorParam() []*VectorParam {
	if m != nil {
		return m.VectorParam
	}
	return nil
}

func (m *PulsarMessages) GetSegments() []*SegmentRecord {
	if m != nil {
		return m.Segments
	}
	return nil
}

func (m *PulsarMessages) GetTimestamp() []int64 {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *PulsarMessages) GetClientId() []int64 {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *PulsarMessages) GetMsgType() OpType {
	if m != nil {
		return m.MsgType
	}
	return OpType_Insert
}

func init() {
	proto.RegisterType((*Status)(nil), "pb.Status")
	proto.RegisterType((*SegmentRecord)(nil), "pb.SegmentRecord")
	proto.RegisterType((*VectorRowRecord)(nil), "pb.VectorRowRecord")
	proto.RegisterType((*AttrRecord)(nil), "pb.AttrRecord")
	proto.RegisterType((*VectorRecord)(nil), "pb.VectorRecord")
	proto.RegisterType((*VectorParam)(nil), "pb.VectorParam")
	proto.RegisterType((*FieldValue)(nil), "pb.FieldValue")
	proto.RegisterType((*PulsarMessage)(nil), "pb.PulsarMessage")
	proto.RegisterType((*PulsarMessages)(nil), "pb.PulsarMessages")
	proto.RegisterEnum("pb.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("pb.DataType", DataType_name, DataType_value)
	proto.RegisterEnum("pb.OpType", OpType_name, OpType_value)
}

func init() { proto.RegisterFile("pulsar.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1101 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x5d, 0x6e, 0xdb, 0x46,
	0x17, 0x0d, 0x45, 0xfd, 0xf1, 0x52, 0x92, 0xc7, 0x13, 0x27, 0x51, 0x90, 0xef, 0x43, 0x54, 0x15,
	0x6d, 0x0d, 0xa3, 0x49, 0x50, 0x25, 0x35, 0xfa, 0xd2, 0x07, 0x9a, 0x1c, 0x25, 0x44, 0x28, 0x52,
	0x1d, 0x52, 0x8e, 0xfd, 0x44, 0xd0, 0xd2, 0x58, 0x65, 0x21, 0x91, 0x02, 0x39, 0x76, 0xa0, 0x65,
	0xb4, 0x4b, 0xe8, 0x1e, 0xba, 0xa6, 0x76, 0x19, 0xc5, 0x0c, 0x49, 0x4b, 0x09, 0xd0, 0x02, 0x29,
	0xda, 0xb7, 0xcb, 0x33, 0xe7, 0xde, 0x39, 0xe7, 0x5c, 0x52, 0x10, 0x74, 0x36, 0x37, 0xab, 0x3c,
	0xca, 0x9e, 0x6f, 0xb2, 0x94, 0xa7, 0xb8, 0xb6, 0xb9, 0x1a, 0xba, 0xd0, 0xf4, 0x79, 0xc4, 0x6f,
	0x72, 0xfc, 0x35, 0x00, 0xcb, 0xb2, 0x34, 0x0b, 0xe7, 0xe9, 0x82, 0xf5, 0x95, 0x81, 0x72, 0xdc,
	0x1b, 0x75, 0x9f, 0x6f, 0xae, 0x9e, 0x13, 0x81, 0x9a, 0xe9, 0x82, 0x51, 0x8d, 0x55, 0x25, 0x7e,
	0x08, 0xcd, 0x8c, 0x45, 0x79, 0x9a, 0xf4, 0x6b, 0x03, 0xe5, 0x58, 0xa3, 0xe5, 0xd3, 0xf0, 0x04,
	0xba, 0x3e, 0x5b, 0xae, 0x59, 0xc2, 0x29, 0x9b, 0xa7, 0xd9, 0x02, 0x3f, 0x86, 0x76, 0xce, 0x96,
	0x61, 0x9c, 0x5c, 0xa7, 0x7d, 0x65, 0xa0, 0x1e, 0x6b, 0xb4, 0x95, 0xb3, 0xa5, 0x9d, 0x5c, 0xa7,
	0xc3, 0x1f, 0xe0, 0xe0, 0x9c, 0xcd, 0x79, 0x9a, 0xd1, 0xf4, 0x7d, 0xc9, 0xfe, 0x3f, 0xc0, 0xf5,
	0x2a, 0x8d, 0x78, 0xb8, 0x88, 0x78, 0x24, 0xf9, 0x35, 0xaa, 0x49, 0xc4, 0x8a, 0x78, 0x84, 0x9f,
	0x82, 0x7e, 0x15, 0x27, 0x51, 0xb6, 0x2d, 0xce, 0xc5, 0xd5, 0x1d, 0x0a, 0x05, 0x24, 0x08, 0xc3,
	0x5f, 0x14, 0x00, 0x83, 0xf3, 0xac, 0x1c, 0xf7, 0x14, 0xf4, 0x38, 0xe1, 0x2f, 0x47, 0xe1, 0x6d,
	0xb4, 0xba, 0x61, 0x72, 0x5e, 0x83, 0x82, 0x84, 0xce, 0x05, 0x52, 0x12, 0x4e, 0x5f, 0x95, 0x84,
	0xda, 0x40, 0x3d, 0x56, 0x25, 0xe1, 0xf4, 0xd5, 0x1d, 0xa1, 0x10, 0x54, 0x10, 0x54, 0xa9, 0xa8,
	0xd0, 0x58, 0x10, 0x3e, 0x83, 0xce, 0x22, 0xbd, 0xb9, 0x5a, 0xb1, 0x92, 0x51, 0x1f, 0xa8, 0xc7,
	0x0a, 0xd5, 0x0b, 0x4c, 0x52, 0x86, 0xdf, 0x43, 0xa7, 0xf4, 0x59, 0xa8, 0x7a, 0x06, 0xad, 0x4c,
	0x56, 0xb9, 0x54, 0xa4, 0x8f, 0xee, 0x8b, 0x98, 0x3f, 0x8a, 0x82, 0x56, 0x9c, 0x21, 0x05, 0xbd,
	0x38, 0x9b, 0x46, 0x59, 0xb4, 0xc6, 0x18, 0xea, 0x3f, 0x89, 0xdc, 0x15, 0x99, 0xbb, 0xac, 0xf1,
	0x0b, 0x80, 0x2c, 0x7d, 0x1f, 0x16, 0x1d, 0x32, 0x16, 0x7d, 0x84, 0xf6, 0x86, 0x16, 0x13, 0xb5,
	0xac, 0x1a, 0x3e, 0xfc, 0x4d, 0x01, 0x18, 0xc7, 0x6c, 0xb5, 0x28, 0x4c, 0x88, 0xd8, 0xc5, 0x53,
	0x98, 0x44, 0x6b, 0x56, 0x4e, 0xd6, 0x24, 0xe2, 0x46, 0x6b, 0x86, 0x07, 0x50, 0xe7, 0xdb, 0x0d,
	0x93, 0x83, 0x7b, 0xa3, 0x8e, 0x18, 0x2c, 0xd2, 0x0e, 0xb6, 0x1b, 0x46, 0xe5, 0x09, 0x7e, 0x01,
	0x7a, 0xc4, 0x79, 0x56, 0x29, 0x50, 0xa5, 0x82, 0x9e, 0x20, 0xee, 0xb6, 0x41, 0x21, 0xda, 0x6d,
	0xe6, 0x5b, 0xe8, 0xde, 0x4a, 0x6d, 0x55, 0x4b, 0xfd, 0x2f, 0x44, 0x77, 0x6e, 0xf7, 0x9e, 0x86,
	0xbf, 0xd7, 0xa0, 0x3b, 0x95, 0xef, 0xf0, 0x84, 0xe5, 0x79, 0xb4, 0x64, 0xf8, 0x2b, 0x38, 0x98,
	0xa7, 0xab, 0x15, 0x9b, 0xf3, 0x38, 0x4d, 0xf6, 0xf5, 0xf7, 0x76, 0xb0, 0x34, 0xf1, 0x25, 0x34,
	0xa5, 0xa3, 0x5c, 0x6e, 0xb9, 0x54, 0xb7, 0xcb, 0x80, 0x96, 0xa7, 0xf8, 0x09, 0x68, 0x2c, 0xe1,
	0x31, 0xdf, 0x86, 0x71, 0x61, 0x44, 0xa5, 0xed, 0x02, 0xb0, 0x17, 0xf8, 0x73, 0xe8, 0x6e, 0xa2,
	0x8c, 0xc7, 0xf2, 0x32, 0x1e, 0x2d, 0xa5, 0x6c, 0x8d, 0x76, 0xee, 0xc0, 0x20, 0x5a, 0xe2, 0x11,
	0x94, 0xa2, 0xc3, 0x8d, 0xd8, 0x58, 0xbf, 0x21, 0xad, 0x1d, 0xec, 0xac, 0xc9, 0x45, 0x52, 0xfd,
	0x76, 0x6f, 0xab, 0xcf, 0xe4, 0x67, 0x22, 0xbe, 0x9b, 0xbc, 0xdf, 0x94, 0xfc, 0x43, 0xc1, 0xff,
	0xe0, 0x5b, 0xa2, 0x77, 0x14, 0xfc, 0x3f, 0xd0, 0x78, 0xbc, 0x66, 0x39, 0x8f, 0xd6, 0x9b, 0x7e,
	0x4b, 0x8a, 0xdc, 0x01, 0xc2, 0xc2, 0x7c, 0x15, 0xb3, 0x84, 0x0b, 0x0b, 0xed, 0xc2, 0x42, 0x01,
	0xd8, 0x0b, 0xfc, 0x05, 0xb4, 0xd7, 0xf9, 0x32, 0x94, 0x0b, 0xd5, 0xe4, 0x42, 0x41, 0xdc, 0xe4,
	0x6d, 0xe4, 0x3a, 0x5b, 0xeb, 0x7c, 0x29, 0x8a, 0xe1, 0x1f, 0x35, 0xe8, 0x7d, 0x90, 0x74, 0xfe,
	0x9f, 0x47, 0xad, 0xfe, 0x1b, 0x51, 0xab, 0x9f, 0x18, 0xb5, 0xfa, 0x89, 0x51, 0xab, 0x7f, 0x1b,
	0xb5, 0xfa, 0x0f, 0xa2, 0x3e, 0xf9, 0xb5, 0x0e, 0xda, 0xdd, 0x8f, 0x2c, 0xd6, 0xa1, 0xe5, 0xcf,
	0x4c, 0x93, 0xf8, 0x3e, 0xba, 0x87, 0x8f, 0x00, 0xcd, 0x5c, 0x72, 0x31, 0x25, 0x66, 0x40, 0xac,
	0x90, 0x50, 0xea, 0x51, 0xa4, 0x60, 0x0c, 0x3d, 0xd3, 0x73, 0x5d, 0x62, 0x06, 0xe1, 0xd8, 0xb0,
	0x1d, 0x62, 0xa1, 0x1a, 0x7e, 0x00, 0x87, 0x53, 0x42, 0x27, 0xb6, 0xef, 0xdb, 0x9e, 0x1b, 0x5a,
	0xc4, 0xb5, 0x89, 0x85, 0x54, 0xfc, 0x18, 0x1e, 0x98, 0x9e, 0xe3, 0x10, 0x33, 0x10, 0xb0, 0xeb,
	0x05, 0x21, 0xb9, 0xb0, 0xfd, 0xc0, 0x47, 0x75, 0x31, 0xdb, 0x76, 0x1c, 0xf2, 0xda, 0x70, 0x42,
	0x83, 0xbe, 0x9e, 0x4d, 0x88, 0x1b, 0xa0, 0x86, 0x98, 0x53, 0xa1, 0x96, 0x3d, 0x21, 0xae, 0x18,
	0x87, 0x5a, 0xf8, 0x21, 0xe0, 0x0a, 0xb6, 0x5d, 0x8b, 0x5c, 0x84, 0xc1, 0xe5, 0x94, 0xa0, 0x36,
	0x7e, 0x02, 0x8f, 0x2a, 0x7c, 0xff, 0x1e, 0x63, 0x42, 0x90, 0x86, 0x11, 0x74, 0xaa, 0xc3, 0xc0,
	0x9b, 0xbe, 0x45, 0xb0, 0x3f, 0x9d, 0x7a, 0xef, 0x28, 0x31, 0x3d, 0x6a, 0x21, 0x7d, 0x1f, 0x3e,
	0x27, 0x66, 0xe0, 0xd1, 0xd0, 0xb6, 0x50, 0x47, 0x88, 0xaf, 0x60, 0x9f, 0x18, 0xd4, 0x7c, 0x13,
	0x52, 0xe2, 0xcf, 0x9c, 0x00, 0x75, 0x45, 0x04, 0x63, 0xdb, 0x21, 0xd2, 0xd1, 0xd8, 0x9b, 0xb9,
	0x16, 0xea, 0xe1, 0x03, 0xd0, 0x27, 0x24, 0x30, 0xaa, 0x4c, 0x0e, 0xc4, 0xfd, 0xa6, 0x61, 0xbe,
	0x21, 0x15, 0x82, 0x70, 0x1f, 0x8e, 0x4c, 0xc3, 0x15, 0x4d, 0x26, 0x25, 0x46, 0x40, 0xc2, 0xb1,
	0xe7, 0x58, 0x84, 0xa2, 0x43, 0x61, 0xf0, 0xa3, 0x13, 0xdb, 0x21, 0x08, 0xef, 0x75, 0x58, 0xc4,
	0x21, 0xbb, 0x8e, 0xfb, 0x7b, 0x1d, 0xd5, 0x89, 0xe8, 0x38, 0x12, 0x66, 0xce, 0x66, 0xb6, 0x63,
	0x95, 0x41, 0x15, 0x4b, 0x7b, 0x80, 0x0f, 0xa1, 0x5b, 0x99, 0x71, 0x1d, 0xdb, 0x0f, 0xd0, 0x43,
	0xfc, 0x08, 0xee, 0x57, 0xd0, 0x84, 0x04, 0xd4, 0x36, 0x8b, 0x54, 0x1f, 0x09, 0xae, 0x37, 0x0b,
	0x42, 0x6f, 0x1c, 0x4e, 0xc8, 0xc4, 0xa3, 0x97, 0xa8, 0x7f, 0xf2, 0xb3, 0x02, 0xed, 0xea, 0x47,
	0x17, 0xb7, 0xa1, 0xee, 0x7a, 0x2e, 0x41, 0xf7, 0x44, 0x75, 0xe6, 0x79, 0x0e, 0x52, 0x44, 0x65,
	0xbb, 0xc1, 0x77, 0xa8, 0x86, 0x35, 0x68, 0xd8, 0x6e, 0xf0, 0xcd, 0x29, 0x52, 0xcb, 0xf2, 0xe5,
	0x08, 0xd5, 0xcb, 0xf2, 0xf4, 0x15, 0x6a, 0x88, 0x72, 0xec, 0x78, 0x46, 0x80, 0x00, 0x03, 0x34,
	0x2d, 0x6f, 0x76, 0xe6, 0x10, 0xa4, 0x8b, 0xda, 0x0f, 0xa8, 0xed, 0xbe, 0x46, 0x47, 0x42, 0x41,
	0xb9, 0x89, 0x33, 0xdb, 0x35, 0xe8, 0x25, 0x5a, 0x88, 0x34, 0x4b, 0xa8, 0x68, 0x66, 0x27, 0xef,
	0xa0, 0x59, 0xbc, 0xcb, 0xa2, 0xd5, 0x4e, 0x72, 0x96, 0x71, 0x74, 0x4f, 0x8e, 0x64, 0x2b, 0xc6,
	0x19, 0x52, 0xe4, 0x48, 0x16, 0x65, 0xf3, 0x1f, 0x51, 0x0d, 0x77, 0xa0, 0x1d, 0xc4, 0x6b, 0xe6,
	0x6f, 0x93, 0x39, 0x52, 0xc5, 0x6b, 0xfe, 0x96, 0x6d, 0x47, 0x3e, 0x5b, 0xa2, 0x3a, 0xee, 0x01,
	0x88, 0x7f, 0x21, 0x71, 0xce, 0xe3, 0x79, 0x8e, 0x1a, 0x57, 0x4d, 0xf9, 0x07, 0xe5, 0xe5, 0x9f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x27, 0x2a, 0xd3, 0x11, 0xb0, 0x08, 0x00, 0x00,
}