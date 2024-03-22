// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: api/new-world-api/v1/user.proto

package apipb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type RankType int32

const (
	RankType_level RankType = 0
)

// Enum value maps for RankType.
var (
	RankType_name = map[int32]string{
		0: "level",
	}
	RankType_value = map[string]int32{
		"level": 0,
	}
)

func (x RankType) Enum() *RankType {
	p := new(RankType)
	*p = x
	return p
}

func (x RankType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RankType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_new_world_api_v1_user_proto_enumTypes[0].Descriptor()
}

func (RankType) Type() protoreflect.EnumType {
	return &file_api_new_world_api_v1_user_proto_enumTypes[0]
}

func (x RankType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RankType.Descriptor instead.
func (RankType) EnumDescriptor() ([]byte, []int) {
	return file_api_new_world_api_v1_user_proto_rawDescGZIP(), []int{0}
}

type PropStatus int32

const (
	PropStatus_prop_unused  PropStatus = 0
	PropStatus_prop_used    PropStatus = 1
	PropStatus_prop_expired PropStatus = 2
)

// Enum value maps for PropStatus.
var (
	PropStatus_name = map[int32]string{
		0: "prop_unused",
		1: "prop_used",
		2: "prop_expired",
	}
	PropStatus_value = map[string]int32{
		"prop_unused":  0,
		"prop_used":    1,
		"prop_expired": 2,
	}
)

func (x PropStatus) Enum() *PropStatus {
	p := new(PropStatus)
	*p = x
	return p
}

func (x PropStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PropStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_new_world_api_v1_user_proto_enumTypes[1].Descriptor()
}

func (PropStatus) Type() protoreflect.EnumType {
	return &file_api_new_world_api_v1_user_proto_enumTypes[1]
}

func (x PropStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PropStatus.Descriptor instead.
func (PropStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_new_world_api_v1_user_proto_rawDescGZIP(), []int{1}
}

type SignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignRequest) Reset() {
	*x = SignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_api_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRequest) ProtoMessage() {}

func (x *SignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_api_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRequest.ProtoReflect.Descriptor instead.
func (*SignRequest) Descriptor() ([]byte, []int) {
	return file_api_new_world_api_v1_user_proto_rawDescGZIP(), []int{0}
}

type SignResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignResult) Reset() {
	*x = SignResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_api_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignResult) ProtoMessage() {}

func (x *SignResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_api_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignResult.ProtoReflect.Descriptor instead.
func (*SignResult) Descriptor() ([]byte, []int) {
	return file_api_new_world_api_v1_user_proto_rawDescGZIP(), []int{1}
}

type AssetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AssetRequest) Reset() {
	*x = AssetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_api_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetRequest) ProtoMessage() {}

func (x *AssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_api_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetRequest.ProtoReflect.Descriptor instead.
func (*AssetRequest) Descriptor() ([]byte, []int) {
	return file_api_new_world_api_v1_user_proto_rawDescGZIP(), []int{2}
}

type AssetResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level int32 `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	Exp   int64 `protobuf:"varint,2,opt,name=exp,proto3" json:"exp,omitempty"`
	Gold  int64 `protobuf:"varint,3,opt,name=gold,proto3" json:"gold,omitempty"`
}

func (x *AssetResult) Reset() {
	*x = AssetResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_api_v1_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetResult) ProtoMessage() {}

func (x *AssetResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_api_v1_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetResult.ProtoReflect.Descriptor instead.
func (*AssetResult) Descriptor() ([]byte, []int) {
	return file_api_new_world_api_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *AssetResult) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *AssetResult) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *AssetResult) GetGold() int64 {
	if x != nil {
		return x.Gold
	}
	return 0
}

type RankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankType RankType `protobuf:"varint,1,opt,name=rank_type,json=rankType,proto3,enum=new_world.v1.RankType" json:"rank_type,omitempty"`
}

func (x *RankRequest) Reset() {
	*x = RankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_api_v1_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankRequest) ProtoMessage() {}

func (x *RankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_api_v1_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankRequest.ProtoReflect.Descriptor instead.
func (*RankRequest) Descriptor() ([]byte, []int) {
	return file_api_new_world_api_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *RankRequest) GetRankType() RankType {
	if x != nil {
		return x.RankType
	}
	return RankType_level
}

type RankResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RankResult) Reset() {
	*x = RankResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_api_v1_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankResult) ProtoMessage() {}

func (x *RankResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_api_v1_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankResult.ProtoReflect.Descriptor instead.
func (*RankResult) Descriptor() ([]byte, []int) {
	return file_api_new_world_api_v1_user_proto_rawDescGZIP(), []int{5}
}

type ExpandParkingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExpandParkingRequest) Reset() {
	*x = ExpandParkingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_api_v1_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandParkingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandParkingRequest) ProtoMessage() {}

func (x *ExpandParkingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_api_v1_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandParkingRequest.ProtoReflect.Descriptor instead.
func (*ExpandParkingRequest) Descriptor() ([]byte, []int) {
	return file_api_new_world_api_v1_user_proto_rawDescGZIP(), []int{6}
}

type ExpandParkingResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExpandParkingResult) Reset() {
	*x = ExpandParkingResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_api_v1_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandParkingResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandParkingResult) ProtoMessage() {}

func (x *ExpandParkingResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_api_v1_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandParkingResult.ProtoReflect.Descriptor instead.
func (*ExpandParkingResult) Descriptor() ([]byte, []int) {
	return file_api_new_world_api_v1_user_proto_rawDescGZIP(), []int{7}
}

type EatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropId int64 `protobuf:"varint,1,opt,name=prop_id,json=propId,proto3" json:"prop_id,omitempty"`
}

func (x *EatRequest) Reset() {
	*x = EatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_api_v1_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EatRequest) ProtoMessage() {}

func (x *EatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_api_v1_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EatRequest.ProtoReflect.Descriptor instead.
func (*EatRequest) Descriptor() ([]byte, []int) {
	return file_api_new_world_api_v1_user_proto_rawDescGZIP(), []int{8}
}

func (x *EatRequest) GetPropId() int64 {
	if x != nil {
		return x.PropId
	}
	return 0
}

type EatResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EatResult) Reset() {
	*x = EatResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_api_v1_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EatResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EatResult) ProtoMessage() {}

func (x *EatResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_api_v1_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EatResult.ProtoReflect.Descriptor instead.
func (*EatResult) Descriptor() ([]byte, []int) {
	return file_api_new_world_api_v1_user_proto_rawDescGZIP(), []int{9}
}

type SkillUpgradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkillId int64 `protobuf:"varint,1,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
}

func (x *SkillUpgradeRequest) Reset() {
	*x = SkillUpgradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_api_v1_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkillUpgradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkillUpgradeRequest) ProtoMessage() {}

func (x *SkillUpgradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_api_v1_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkillUpgradeRequest.ProtoReflect.Descriptor instead.
func (*SkillUpgradeRequest) Descriptor() ([]byte, []int) {
	return file_api_new_world_api_v1_user_proto_rawDescGZIP(), []int{10}
}

func (x *SkillUpgradeRequest) GetSkillId() int64 {
	if x != nil {
		return x.SkillId
	}
	return 0
}

type SkillUpgradeResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SkillUpgradeResult) Reset() {
	*x = SkillUpgradeResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_api_v1_user_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkillUpgradeResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkillUpgradeResult) ProtoMessage() {}

func (x *SkillUpgradeResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_api_v1_user_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkillUpgradeResult.ProtoReflect.Descriptor instead.
func (*SkillUpgradeResult) Descriptor() ([]byte, []int) {
	return file_api_new_world_api_v1_user_proto_rawDescGZIP(), []int{11}
}

type BuildingUpgradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildingId int64 `protobuf:"varint,1,opt,name=building_id,json=buildingId,proto3" json:"building_id,omitempty"`
}

func (x *BuildingUpgradeRequest) Reset() {
	*x = BuildingUpgradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_api_v1_user_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildingUpgradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildingUpgradeRequest) ProtoMessage() {}

func (x *BuildingUpgradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_api_v1_user_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildingUpgradeRequest.ProtoReflect.Descriptor instead.
func (*BuildingUpgradeRequest) Descriptor() ([]byte, []int) {
	return file_api_new_world_api_v1_user_proto_rawDescGZIP(), []int{12}
}

func (x *BuildingUpgradeRequest) GetBuildingId() int64 {
	if x != nil {
		return x.BuildingId
	}
	return 0
}

type BuildingUpgradeResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BuildingUpgradeResult) Reset() {
	*x = BuildingUpgradeResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_api_v1_user_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildingUpgradeResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildingUpgradeResult) ProtoMessage() {}

func (x *BuildingUpgradeResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_api_v1_user_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildingUpgradeResult.ProtoReflect.Descriptor instead.
func (*BuildingUpgradeResult) Descriptor() ([]byte, []int) {
	return file_api_new_world_api_v1_user_proto_rawDescGZIP(), []int{13}
}

var File_api_new_world_api_v1_user_proto protoreflect.FileDescriptor

var file_api_new_world_api_v1_user_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x77, 0x2d, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a,
	0x0b, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0c, 0x0a, 0x0a,
	0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x49, 0x0a, 0x0b, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x78,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x67, 0x6f, 0x6c, 0x64, 0x22, 0x42, 0x0a, 0x0b, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x09, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x72, 0x61, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x52, 0x61, 0x6e,
	0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x45, 0x78, 0x70, 0x61, 0x6e,
	0x64, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x15, 0x0a, 0x13, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x25, 0x0a, 0x0a, 0x45, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x70, 0x49, 0x64, 0x22, 0x0b, 0x0a,
	0x09, 0x45, 0x61, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x30, 0x0a, 0x13, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x39, 0x0a, 0x16, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x70,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x17, 0x0a,
	0x15, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2a, 0x15, 0x0a, 0x08, 0x52, 0x61, 0x6e, 0x6b, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x10, 0x00, 0x2a, 0x3e, 0x0a,
	0x0a, 0x50, 0x72, 0x6f, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x70, 0x5f, 0x75, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x70, 0x72, 0x6f, 0x70, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x70,
	0x72, 0x6f, 0x70, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x10, 0x02, 0x32, 0xf7, 0x05,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x04, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x19,
	0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6e, 0x65, 0x77, 0x5f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x3a, 0x01,
	0x2a, 0x12, 0x5d, 0x0a, 0x05, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x6e, 0x65, 0x77,
	0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x3a, 0x01, 0x2a,
	0x12, 0x59, 0x0a, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x19, 0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x1c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x3a, 0x01, 0x2a, 0x12, 0x7e, 0x0a, 0x0d, 0x45,
	0x78, 0x70, 0x61, 0x6e, 0x64, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x2e, 0x6e,
	0x65, 0x77, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x61,
	0x6e, 0x64, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64,
	0x2f, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x3a, 0x01, 0x2a, 0x12, 0x55, 0x0a, 0x03, 0x45,
	0x61, 0x74, 0x12, 0x18, 0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6e,
	0x65, 0x77, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x65, 0x61, 0x74, 0x3a,
	0x01, 0x2a, 0x12, 0x7a, 0x0a, 0x0c, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x55, 0x70, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x12, 0x21, 0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22,
	0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x2f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x86,
	0x01, 0x0a, 0x0f, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x70, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x12, 0x24, 0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67,
	0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x28, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x75, 0x70, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x50, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56,
	0x31, 0x50, 0x01, 0x5a, 0x24, 0x6e, 0x65, 0x77, 0x2d, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6e, 0x65, 0x77, 0x2d, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_new_world_api_v1_user_proto_rawDescOnce sync.Once
	file_api_new_world_api_v1_user_proto_rawDescData = file_api_new_world_api_v1_user_proto_rawDesc
)

func file_api_new_world_api_v1_user_proto_rawDescGZIP() []byte {
	file_api_new_world_api_v1_user_proto_rawDescOnce.Do(func() {
		file_api_new_world_api_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_new_world_api_v1_user_proto_rawDescData)
	})
	return file_api_new_world_api_v1_user_proto_rawDescData
}

var file_api_new_world_api_v1_user_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_new_world_api_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_api_new_world_api_v1_user_proto_goTypes = []interface{}{
	(RankType)(0),                  // 0: new_world.v1.RankType
	(PropStatus)(0),                // 1: new_world.v1.PropStatus
	(*SignRequest)(nil),            // 2: new_world.v1.SignRequest
	(*SignResult)(nil),             // 3: new_world.v1.SignResult
	(*AssetRequest)(nil),           // 4: new_world.v1.AssetRequest
	(*AssetResult)(nil),            // 5: new_world.v1.AssetResult
	(*RankRequest)(nil),            // 6: new_world.v1.RankRequest
	(*RankResult)(nil),             // 7: new_world.v1.RankResult
	(*ExpandParkingRequest)(nil),   // 8: new_world.v1.ExpandParkingRequest
	(*ExpandParkingResult)(nil),    // 9: new_world.v1.ExpandParkingResult
	(*EatRequest)(nil),             // 10: new_world.v1.EatRequest
	(*EatResult)(nil),              // 11: new_world.v1.EatResult
	(*SkillUpgradeRequest)(nil),    // 12: new_world.v1.SkillUpgradeRequest
	(*SkillUpgradeResult)(nil),     // 13: new_world.v1.SkillUpgradeResult
	(*BuildingUpgradeRequest)(nil), // 14: new_world.v1.BuildingUpgradeRequest
	(*BuildingUpgradeResult)(nil),  // 15: new_world.v1.BuildingUpgradeResult
}
var file_api_new_world_api_v1_user_proto_depIdxs = []int32{
	0,  // 0: new_world.v1.RankRequest.rank_type:type_name -> new_world.v1.RankType
	2,  // 1: new_world.v1.User.Sign:input_type -> new_world.v1.SignRequest
	4,  // 2: new_world.v1.User.Asset:input_type -> new_world.v1.AssetRequest
	6,  // 3: new_world.v1.User.Rank:input_type -> new_world.v1.RankRequest
	8,  // 4: new_world.v1.User.ExpandParking:input_type -> new_world.v1.ExpandParkingRequest
	10, // 5: new_world.v1.User.Eat:input_type -> new_world.v1.EatRequest
	12, // 6: new_world.v1.User.SkillUpgrade:input_type -> new_world.v1.SkillUpgradeRequest
	14, // 7: new_world.v1.User.BuildingUpgrade:input_type -> new_world.v1.BuildingUpgradeRequest
	3,  // 8: new_world.v1.User.Sign:output_type -> new_world.v1.SignResult
	5,  // 9: new_world.v1.User.Asset:output_type -> new_world.v1.AssetResult
	7,  // 10: new_world.v1.User.Rank:output_type -> new_world.v1.RankResult
	9,  // 11: new_world.v1.User.ExpandParking:output_type -> new_world.v1.ExpandParkingResult
	11, // 12: new_world.v1.User.Eat:output_type -> new_world.v1.EatResult
	13, // 13: new_world.v1.User.SkillUpgrade:output_type -> new_world.v1.SkillUpgradeResult
	15, // 14: new_world.v1.User.BuildingUpgrade:output_type -> new_world.v1.BuildingUpgradeResult
	8,  // [8:15] is the sub-list for method output_type
	1,  // [1:8] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_new_world_api_v1_user_proto_init() }
func file_api_new_world_api_v1_user_proto_init() {
	if File_api_new_world_api_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_new_world_api_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRequest); i {
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
		file_api_new_world_api_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignResult); i {
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
		file_api_new_world_api_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetRequest); i {
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
		file_api_new_world_api_v1_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetResult); i {
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
		file_api_new_world_api_v1_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankRequest); i {
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
		file_api_new_world_api_v1_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankResult); i {
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
		file_api_new_world_api_v1_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandParkingRequest); i {
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
		file_api_new_world_api_v1_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandParkingResult); i {
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
		file_api_new_world_api_v1_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EatRequest); i {
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
		file_api_new_world_api_v1_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EatResult); i {
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
		file_api_new_world_api_v1_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkillUpgradeRequest); i {
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
		file_api_new_world_api_v1_user_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkillUpgradeResult); i {
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
		file_api_new_world_api_v1_user_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildingUpgradeRequest); i {
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
		file_api_new_world_api_v1_user_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildingUpgradeResult); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_new_world_api_v1_user_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_new_world_api_v1_user_proto_goTypes,
		DependencyIndexes: file_api_new_world_api_v1_user_proto_depIdxs,
		EnumInfos:         file_api_new_world_api_v1_user_proto_enumTypes,
		MessageInfos:      file_api_new_world_api_v1_user_proto_msgTypes,
	}.Build()
	File_api_new_world_api_v1_user_proto = out.File
	file_api_new_world_api_v1_user_proto_rawDesc = nil
	file_api_new_world_api_v1_user_proto_goTypes = nil
	file_api_new_world_api_v1_user_proto_depIdxs = nil
}
