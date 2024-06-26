// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: api/shared/v1/common.proto

package sharedpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ContentType int32

const (
	ContentType_JSON_TEXT ContentType = 0 // JSON 文本
	ContentType_PROTOBUF  ContentType = 1 // PROTOBUF
)

// Enum value maps for ContentType.
var (
	ContentType_name = map[int32]string{
		0: "JSON_TEXT",
		1: "PROTOBUF",
	}
	ContentType_value = map[string]int32{
		"JSON_TEXT": 0,
		"PROTOBUF":  1,
	}
)

func (x ContentType) Enum() *ContentType {
	p := new(ContentType)
	*p = x
	return p
}

func (x ContentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContentType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_shared_v1_common_proto_enumTypes[0].Descriptor()
}

func (ContentType) Type() protoreflect.EnumType {
	return &file_api_shared_v1_common_proto_enumTypes[0]
}

func (x ContentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContentType.Descriptor instead.
func (ContentType) EnumDescriptor() ([]byte, []int) {
	return file_api_shared_v1_common_proto_rawDescGZIP(), []int{0}
}

type NullValue int32

const (
	NullValue_NULL_VALUE NullValue = 0
)

// Enum value maps for NullValue.
var (
	NullValue_name = map[int32]string{
		0: "NULL_VALUE",
	}
	NullValue_value = map[string]int32{
		"NULL_VALUE": 0,
	}
)

func (x NullValue) Enum() *NullValue {
	p := new(NullValue)
	*p = x
	return p
}

func (x NullValue) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NullValue) Descriptor() protoreflect.EnumDescriptor {
	return file_api_shared_v1_common_proto_enumTypes[1].Descriptor()
}

func (NullValue) Type() protoreflect.EnumType {
	return &file_api_shared_v1_common_proto_enumTypes[1]
}

func (x NullValue) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NullValue.Descriptor instead.
func (NullValue) EnumDescriptor() ([]byte, []int) {
	return file_api_shared_v1_common_proto_rawDescGZIP(), []int{1}
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//	*Value_Int64Value
	//	*Value_Int32Value
	//	*Value_FloatValue
	//	*Value_DoubleValue
	//	*Value_StringValue
	//	*Value_BoolValue
	//	*Value_NullValue
	//	*Value_ListValue
	//	*Value_MapValue
	Kind isValue_Kind `protobuf_oneof:"kind"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shared_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_api_shared_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_api_shared_v1_common_proto_rawDescGZIP(), []int{0}
}

func (m *Value) GetKind() isValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Value) GetInt64Value() int64 {
	if x, ok := x.GetKind().(*Value_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (x *Value) GetInt32Value() int32 {
	if x, ok := x.GetKind().(*Value_Int32Value); ok {
		return x.Int32Value
	}
	return 0
}

func (x *Value) GetFloatValue() float32 {
	if x, ok := x.GetKind().(*Value_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (x *Value) GetDoubleValue() float64 {
	if x, ok := x.GetKind().(*Value_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (x *Value) GetStringValue() string {
	if x, ok := x.GetKind().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *Value) GetBoolValue() bool {
	if x, ok := x.GetKind().(*Value_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (x *Value) GetNullValue() NullValue {
	if x, ok := x.GetKind().(*Value_NullValue); ok {
		return x.NullValue
	}
	return NullValue_NULL_VALUE
}

func (x *Value) GetListValue() *ListValue {
	if x, ok := x.GetKind().(*Value_ListValue); ok {
		return x.ListValue
	}
	return nil
}

func (x *Value) GetMapValue() *MapValue {
	if x, ok := x.GetKind().(*Value_MapValue); ok {
		return x.MapValue
	}
	return nil
}

type isValue_Kind interface {
	isValue_Kind()
}

type Value_Int64Value struct {
	Int64Value int64 `protobuf:"varint,1,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type Value_Int32Value struct {
	Int32Value int32 `protobuf:"varint,2,opt,name=int32_value,json=int32Value,proto3,oneof"`
}

type Value_FloatValue struct {
	FloatValue float32 `protobuf:"fixed32,3,opt,name=float_value,json=floatValue,proto3,oneof"`
}

type Value_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Value_StringValue struct {
	StringValue string `protobuf:"bytes,5,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_BoolValue struct {
	BoolValue bool `protobuf:"varint,6,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type Value_NullValue struct {
	NullValue NullValue `protobuf:"varint,7,opt,name=null_value,json=nullValue,proto3,enum=shared.v1.NullValue,oneof"`
}

type Value_ListValue struct {
	ListValue *ListValue `protobuf:"bytes,8,opt,name=list_value,json=listValue,proto3,oneof"`
}

type Value_MapValue struct {
	MapValue *MapValue `protobuf:"bytes,9,opt,name=map_value,json=mapValue,proto3,oneof"`
}

func (*Value_Int64Value) isValue_Kind() {}

func (*Value_Int32Value) isValue_Kind() {}

func (*Value_FloatValue) isValue_Kind() {}

func (*Value_DoubleValue) isValue_Kind() {}

func (*Value_StringValue) isValue_Kind() {}

func (*Value_BoolValue) isValue_Kind() {}

func (*Value_NullValue) isValue_Kind() {}

func (*Value_ListValue) isValue_Kind() {}

func (*Value_MapValue) isValue_Kind() {}

type ListValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *ListValue) Reset() {
	*x = ListValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shared_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListValue) ProtoMessage() {}

func (x *ListValue) ProtoReflect() protoreflect.Message {
	mi := &file_api_shared_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListValue.ProtoReflect.Descriptor instead.
func (*ListValue) Descriptor() ([]byte, []int) {
	return file_api_shared_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *ListValue) GetValues() []*Value {
	if x != nil {
		return x.Values
	}
	return nil
}

type MapValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields map[string]*Value `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MapValue) Reset() {
	*x = MapValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shared_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapValue) ProtoMessage() {}

func (x *MapValue) ProtoReflect() protoreflect.Message {
	mi := &file_api_shared_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapValue.ProtoReflect.Descriptor instead.
func (*MapValue) Descriptor() ([]byte, []int) {
	return file_api_shared_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *MapValue) GetFields() map[string]*Value {
	if x != nil {
		return x.Fields
	}
	return nil
}

type ClientSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid            int64             `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	BrokerId       string            `protobuf:"bytes,3,opt,name=broker_id,json=brokerId,proto3" json:"broker_id,omitempty"`
	BrokerEndpoint string            `protobuf:"bytes,4,opt,name=broker_endpoint,json=brokerEndpoint,proto3" json:"broker_endpoint,omitempty"`
	Metadata       map[string]*Value `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ClientSession) Reset() {
	*x = ClientSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shared_v1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientSession) ProtoMessage() {}

func (x *ClientSession) ProtoReflect() protoreflect.Message {
	mi := &file_api_shared_v1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientSession.ProtoReflect.Descriptor instead.
func (*ClientSession) Descriptor() ([]byte, []int) {
	return file_api_shared_v1_common_proto_rawDescGZIP(), []int{3}
}

func (x *ClientSession) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClientSession) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ClientSession) GetBrokerId() string {
	if x != nil {
		return x.BrokerId
	}
	return ""
}

func (x *ClientSession) GetBrokerEndpoint() string {
	if x != nil {
		return x.BrokerEndpoint
	}
	return ""
}

func (x *ClientSession) GetMetadata() map[string]*Value {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type ServiceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	ServiceId   string `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
}

func (x *ServiceInfo) Reset() {
	*x = ServiceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shared_v1_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceInfo) ProtoMessage() {}

func (x *ServiceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_shared_v1_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceInfo.ProtoReflect.Descriptor instead.
func (*ServiceInfo) Descriptor() ([]byte, []int) {
	return file_api_shared_v1_common_proto_rawDescGZIP(), []int{4}
}

func (x *ServiceInfo) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *ServiceInfo) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

var File_api_shared_v1_common_proto protoreflect.FileDescriptor

var file_api_shared_v1_common_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x85, 0x03, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x21, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x0a,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48,
	0x00, 0x52, 0x09, 0x6e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x35, 0x0a, 0x0a,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x6d, 0x61, 0x70, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x08, 0x6d,
	0x61, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22,
	0x35, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x28, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x4b, 0x0a, 0x0b,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8a, 0x02, 0x0a, 0x0d, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x4d, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4f, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x2a, 0x2a, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x4a, 0x53, 0x4f, 0x4e, 0x5f, 0x54,
	0x45, 0x58, 0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x42, 0x55,
	0x46, 0x10, 0x01, 0x2a, 0x1b, 0x0a, 0x09, 0x4e, 0x75, 0x6c, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x55, 0x4c, 0x4c, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x00,
	0x42, 0x48, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x26, 0x6e, 0x65, 0x77, 0x2d, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2d, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76,
	0x31, 0x3b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_shared_v1_common_proto_rawDescOnce sync.Once
	file_api_shared_v1_common_proto_rawDescData = file_api_shared_v1_common_proto_rawDesc
)

func file_api_shared_v1_common_proto_rawDescGZIP() []byte {
	file_api_shared_v1_common_proto_rawDescOnce.Do(func() {
		file_api_shared_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_shared_v1_common_proto_rawDescData)
	})
	return file_api_shared_v1_common_proto_rawDescData
}

var file_api_shared_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_shared_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_shared_v1_common_proto_goTypes = []interface{}{
	(ContentType)(0),      // 0: shared.v1.ContentType
	(NullValue)(0),        // 1: shared.v1.NullValue
	(*Value)(nil),         // 2: shared.v1.Value
	(*ListValue)(nil),     // 3: shared.v1.ListValue
	(*MapValue)(nil),      // 4: shared.v1.MapValue
	(*ClientSession)(nil), // 5: shared.v1.ClientSession
	(*ServiceInfo)(nil),   // 6: shared.v1.ServiceInfo
	nil,                   // 7: shared.v1.MapValue.FieldsEntry
	nil,                   // 8: shared.v1.ClientSession.MetadataEntry
}
var file_api_shared_v1_common_proto_depIdxs = []int32{
	1, // 0: shared.v1.Value.null_value:type_name -> shared.v1.NullValue
	3, // 1: shared.v1.Value.list_value:type_name -> shared.v1.ListValue
	4, // 2: shared.v1.Value.map_value:type_name -> shared.v1.MapValue
	2, // 3: shared.v1.ListValue.values:type_name -> shared.v1.Value
	7, // 4: shared.v1.MapValue.fields:type_name -> shared.v1.MapValue.FieldsEntry
	8, // 5: shared.v1.ClientSession.metadata:type_name -> shared.v1.ClientSession.MetadataEntry
	2, // 6: shared.v1.MapValue.FieldsEntry.value:type_name -> shared.v1.Value
	2, // 7: shared.v1.ClientSession.MetadataEntry.value:type_name -> shared.v1.Value
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_api_shared_v1_common_proto_init() }
func file_api_shared_v1_common_proto_init() {
	if File_api_shared_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_shared_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_shared_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_shared_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_shared_v1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientSession); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_shared_v1_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_api_shared_v1_common_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Value_Int64Value)(nil),
		(*Value_Int32Value)(nil),
		(*Value_FloatValue)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BoolValue)(nil),
		(*Value_NullValue)(nil),
		(*Value_ListValue)(nil),
		(*Value_MapValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_shared_v1_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_shared_v1_common_proto_goTypes,
		DependencyIndexes: file_api_shared_v1_common_proto_depIdxs,
		EnumInfos:         file_api_shared_v1_common_proto_enumTypes,
		MessageInfos:      file_api_shared_v1_common_proto_msgTypes,
	}.Build()
	File_api_shared_v1_common_proto = out.File
	file_api_shared_v1_common_proto_rawDesc = nil
	file_api_shared_v1_common_proto_goTypes = nil
	file_api_shared_v1_common_proto_depIdxs = nil
}
