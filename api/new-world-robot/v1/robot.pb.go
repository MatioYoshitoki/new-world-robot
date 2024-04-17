// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: api/new-world-robot/v1/robot.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1 "new-world-robot/api/shared/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RobotMemory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth                    *RobotMemory_Auth     `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	Assets                  *RobotMemory_Assets   `protobuf:"bytes,2,opt,name=assets,proto3" json:"assets,omitempty"`
	Fishes                  *RobotMemory_Fishes   `protobuf:"bytes,3,opt,name=fishes,proto3" json:"fishes,omitempty"`
	ParkingMap              map[string]string     `protobuf:"bytes,4,rep,name=parkingMap,proto3" json:"parkingMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MarketMap               []*RobotMemory_Market `protobuf:"bytes,5,rep,name=marketMap,proto3" json:"marketMap,omitempty"`
	MineFishProductRelation map[int64]int64       `protobuf:"bytes,6,rep,name=mine_fish_product_relation,json=mineFishProductRelation,proto3" json:"mine_fish_product_relation,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *RobotMemory) Reset() {
	*x = RobotMemory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_robot_v1_robot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RobotMemory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RobotMemory) ProtoMessage() {}

func (x *RobotMemory) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_robot_v1_robot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RobotMemory.ProtoReflect.Descriptor instead.
func (*RobotMemory) Descriptor() ([]byte, []int) {
	return file_api_new_world_robot_v1_robot_proto_rawDescGZIP(), []int{0}
}

func (x *RobotMemory) GetAuth() *RobotMemory_Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *RobotMemory) GetAssets() *RobotMemory_Assets {
	if x != nil {
		return x.Assets
	}
	return nil
}

func (x *RobotMemory) GetFishes() *RobotMemory_Fishes {
	if x != nil {
		return x.Fishes
	}
	return nil
}

func (x *RobotMemory) GetParkingMap() map[string]string {
	if x != nil {
		return x.ParkingMap
	}
	return nil
}

func (x *RobotMemory) GetMarketMap() []*RobotMemory_Market {
	if x != nil {
		return x.MarketMap
	}
	return nil
}

func (x *RobotMemory) GetMineFishProductRelation() map[int64]int64 {
	if x != nil {
		return x.MineFishProductRelation
	}
	return nil
}

type RobotMemory_Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Face        string `protobuf:"bytes,1,opt,name=face,proto3" json:"face,omitempty"`
	FKey        string `protobuf:"bytes,2,opt,name=f_key,json=fKey,proto3" json:"f_key,omitempty"`
	AccessToken string `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Uid         string `protobuf:"bytes,4,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *RobotMemory_Auth) Reset() {
	*x = RobotMemory_Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_robot_v1_robot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RobotMemory_Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RobotMemory_Auth) ProtoMessage() {}

func (x *RobotMemory_Auth) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_robot_v1_robot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RobotMemory_Auth.ProtoReflect.Descriptor instead.
func (*RobotMemory_Auth) Descriptor() ([]byte, []int) {
	return file_api_new_world_robot_v1_robot_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RobotMemory_Auth) GetFace() string {
	if x != nil {
		return x.Face
	}
	return ""
}

func (x *RobotMemory_Auth) GetFKey() string {
	if x != nil {
		return x.FKey
	}
	return ""
}

func (x *RobotMemory_Auth) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *RobotMemory_Auth) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type RobotMemory_Assets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exp   int64 `protobuf:"varint,1,opt,name=exp,proto3" json:"exp,omitempty"`
	Gold  int64 `protobuf:"varint,2,opt,name=gold,proto3" json:"gold,omitempty"`
	Level int32 `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *RobotMemory_Assets) Reset() {
	*x = RobotMemory_Assets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_robot_v1_robot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RobotMemory_Assets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RobotMemory_Assets) ProtoMessage() {}

func (x *RobotMemory_Assets) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_robot_v1_robot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RobotMemory_Assets.ProtoReflect.Descriptor instead.
func (*RobotMemory_Assets) Descriptor() ([]byte, []int) {
	return file_api_new_world_robot_v1_robot_proto_rawDescGZIP(), []int{0, 1}
}

func (x *RobotMemory_Assets) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *RobotMemory_Assets) GetGold() int64 {
	if x != nil {
		return x.Gold
	}
	return 0
}

func (x *RobotMemory_Assets) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type RobotMemory_Fishes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FishList []*RobotMemory_Fishes_Fish `protobuf:"bytes,1,rep,name=fish_list,json=fishList,proto3" json:"fish_list,omitempty"`
}

func (x *RobotMemory_Fishes) Reset() {
	*x = RobotMemory_Fishes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_robot_v1_robot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RobotMemory_Fishes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RobotMemory_Fishes) ProtoMessage() {}

func (x *RobotMemory_Fishes) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_robot_v1_robot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RobotMemory_Fishes.ProtoReflect.Descriptor instead.
func (*RobotMemory_Fishes) Descriptor() ([]byte, []int) {
	return file_api_new_world_robot_v1_robot_proto_rawDescGZIP(), []int{0, 2}
}

func (x *RobotMemory_Fishes) GetFishList() []*RobotMemory_Fishes_Fish {
	if x != nil {
		return x.FishList
	}
	return nil
}

type RobotMemory_Market struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int64 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Price     int64 `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *RobotMemory_Market) Reset() {
	*x = RobotMemory_Market{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_robot_v1_robot_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RobotMemory_Market) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RobotMemory_Market) ProtoMessage() {}

func (x *RobotMemory_Market) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_robot_v1_robot_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RobotMemory_Market.ProtoReflect.Descriptor instead.
func (*RobotMemory_Market) Descriptor() ([]byte, []int) {
	return file_api_new_world_robot_v1_robot_proto_rawDescGZIP(), []int{0, 3}
}

func (x *RobotMemory_Market) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *RobotMemory_Market) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type RobotMemory_Fishes_Fish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FishId int64         `protobuf:"varint,1,opt,name=fish_id,json=fishId,proto3" json:"fish_id,omitempty"`
	Statue v1.FishStatus `protobuf:"varint,13,opt,name=statue,proto3,enum=shared.v1.FishStatus" json:"statue,omitempty"`
}

func (x *RobotMemory_Fishes_Fish) Reset() {
	*x = RobotMemory_Fishes_Fish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_robot_v1_robot_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RobotMemory_Fishes_Fish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RobotMemory_Fishes_Fish) ProtoMessage() {}

func (x *RobotMemory_Fishes_Fish) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_robot_v1_robot_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RobotMemory_Fishes_Fish.ProtoReflect.Descriptor instead.
func (*RobotMemory_Fishes_Fish) Descriptor() ([]byte, []int) {
	return file_api_new_world_robot_v1_robot_proto_rawDescGZIP(), []int{0, 2, 0}
}

func (x *RobotMemory_Fishes_Fish) GetFishId() int64 {
	if x != nil {
		return x.FishId
	}
	return 0
}

func (x *RobotMemory_Fishes_Fish) GetStatue() v1.FishStatus {
	if x != nil {
		return x.Statue
	}
	return v1.FishStatus(0)
}

var File_api_new_world_robot_v1_robot_proto protoreflect.FileDescriptor

var file_api_new_world_robot_v1_robot_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x77, 0x2d, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2d,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x09, 0x62, 0x69, 0x7a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x07, 0x0a, 0x0b, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x12, 0x38, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x3e, 0x0a,
	0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x6e, 0x65, 0x77, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x3e, 0x0a,
	0x06, 0x66, 0x69, 0x73, 0x68, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x6e, 0x65, 0x77, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x46,
	0x69, 0x73, 0x68, 0x65, 0x73, 0x52, 0x06, 0x66, 0x69, 0x73, 0x68, 0x65, 0x73, 0x12, 0x4f, 0x0a,
	0x0a, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x2e, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x12, 0x44,
	0x0a, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x4d, 0x61, 0x70, 0x12, 0x79, 0x0a, 0x1a, 0x6d, 0x69, 0x6e, 0x65, 0x5f, 0x66, 0x69, 0x73,
	0x68, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f,
	0x62, 0x6f, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x4d, 0x69, 0x6e, 0x65, 0x46, 0x69,
	0x73, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x17, 0x6d, 0x69, 0x6e, 0x65, 0x46, 0x69, 0x73, 0x68,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x64, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x61, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x61, 0x63, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x66,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x4b, 0x65, 0x79,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x1a, 0x44, 0x0a, 0x06, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x78,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x67, 0x6f, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x1a, 0xa2, 0x01, 0x0a, 0x06,
	0x46, 0x69, 0x73, 0x68, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x09, 0x66, 0x69, 0x73, 0x68, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x65, 0x77, 0x5f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x6f, 0x62, 0x6f, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x46, 0x69, 0x73, 0x68, 0x65,
	0x73, 0x2e, 0x46, 0x69, 0x73, 0x68, 0x52, 0x08, 0x66, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74,
	0x1a, 0x4e, 0x0a, 0x04, 0x46, 0x69, 0x73, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x73, 0x68,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69, 0x73, 0x68, 0x49,
	0x64, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69,
	0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x65,
	0x1a, 0x3d, 0x0a, 0x06, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x1a,
	0x3d, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4a,
	0x0a, 0x1c, 0x4d, 0x69, 0x6e, 0x65, 0x46, 0x69, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x6e, 0x65,
	0x77, 0x2d, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2d, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6e, 0x65, 0x77, 0x2d, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2d, 0x72, 0x6f, 0x62, 0x6f,
	0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_new_world_robot_v1_robot_proto_rawDescOnce sync.Once
	file_api_new_world_robot_v1_robot_proto_rawDescData = file_api_new_world_robot_v1_robot_proto_rawDesc
)

func file_api_new_world_robot_v1_robot_proto_rawDescGZIP() []byte {
	file_api_new_world_robot_v1_robot_proto_rawDescOnce.Do(func() {
		file_api_new_world_robot_v1_robot_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_new_world_robot_v1_robot_proto_rawDescData)
	})
	return file_api_new_world_robot_v1_robot_proto_rawDescData
}

var file_api_new_world_robot_v1_robot_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_new_world_robot_v1_robot_proto_goTypes = []interface{}{
	(*RobotMemory)(nil),             // 0: new_world_robot.v1.RobotMemory
	(*RobotMemory_Auth)(nil),        // 1: new_world_robot.v1.RobotMemory.Auth
	(*RobotMemory_Assets)(nil),      // 2: new_world_robot.v1.RobotMemory.Assets
	(*RobotMemory_Fishes)(nil),      // 3: new_world_robot.v1.RobotMemory.Fishes
	(*RobotMemory_Market)(nil),      // 4: new_world_robot.v1.RobotMemory.Market
	nil,                             // 5: new_world_robot.v1.RobotMemory.ParkingMapEntry
	nil,                             // 6: new_world_robot.v1.RobotMemory.MineFishProductRelationEntry
	(*RobotMemory_Fishes_Fish)(nil), // 7: new_world_robot.v1.RobotMemory.Fishes.Fish
	(v1.FishStatus)(0),              // 8: shared.v1.FishStatus
}
var file_api_new_world_robot_v1_robot_proto_depIdxs = []int32{
	1, // 0: new_world_robot.v1.RobotMemory.auth:type_name -> new_world_robot.v1.RobotMemory.Auth
	2, // 1: new_world_robot.v1.RobotMemory.assets:type_name -> new_world_robot.v1.RobotMemory.Assets
	3, // 2: new_world_robot.v1.RobotMemory.fishes:type_name -> new_world_robot.v1.RobotMemory.Fishes
	5, // 3: new_world_robot.v1.RobotMemory.parkingMap:type_name -> new_world_robot.v1.RobotMemory.ParkingMapEntry
	4, // 4: new_world_robot.v1.RobotMemory.marketMap:type_name -> new_world_robot.v1.RobotMemory.Market
	6, // 5: new_world_robot.v1.RobotMemory.mine_fish_product_relation:type_name -> new_world_robot.v1.RobotMemory.MineFishProductRelationEntry
	7, // 6: new_world_robot.v1.RobotMemory.Fishes.fish_list:type_name -> new_world_robot.v1.RobotMemory.Fishes.Fish
	8, // 7: new_world_robot.v1.RobotMemory.Fishes.Fish.statue:type_name -> shared.v1.FishStatus
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_api_new_world_robot_v1_robot_proto_init() }
func file_api_new_world_robot_v1_robot_proto_init() {
	if File_api_new_world_robot_v1_robot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_new_world_robot_v1_robot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RobotMemory); i {
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
		file_api_new_world_robot_v1_robot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RobotMemory_Auth); i {
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
		file_api_new_world_robot_v1_robot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RobotMemory_Assets); i {
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
		file_api_new_world_robot_v1_robot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RobotMemory_Fishes); i {
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
		file_api_new_world_robot_v1_robot_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RobotMemory_Market); i {
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
		file_api_new_world_robot_v1_robot_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RobotMemory_Fishes_Fish); i {
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
			RawDescriptor: file_api_new_world_robot_v1_robot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_new_world_robot_v1_robot_proto_goTypes,
		DependencyIndexes: file_api_new_world_robot_v1_robot_proto_depIdxs,
		MessageInfos:      file_api_new_world_robot_v1_robot_proto_msgTypes,
	}.Build()
	File_api_new_world_robot_v1_robot_proto = out.File
	file_api_new_world_robot_v1_robot_proto_rawDesc = nil
	file_api_new_world_robot_v1_robot_proto_goTypes = nil
	file_api_new_world_robot_v1_robot_proto_depIdxs = nil
}
