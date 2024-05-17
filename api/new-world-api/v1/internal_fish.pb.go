// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: api/new-world-api/v1/internal_fish.proto

package apipb

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

type UserFishStatusUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    int64         `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	FishId int64         `protobuf:"varint,2,opt,name=fish_id,json=fishId,proto3" json:"fish_id,omitempty"`
	From   v1.FishStatus `protobuf:"varint,3,opt,name=from,proto3,enum=shared.v1.FishStatus" json:"from,omitempty"`
	To     v1.FishStatus `protobuf:"varint,4,opt,name=to,proto3,enum=shared.v1.FishStatus" json:"to,omitempty"`
}

func (x *UserFishStatusUpdateRequest) Reset() {
	*x = UserFishStatusUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_api_v1_internal_fish_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFishStatusUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFishStatusUpdateRequest) ProtoMessage() {}

func (x *UserFishStatusUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_api_v1_internal_fish_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFishStatusUpdateRequest.ProtoReflect.Descriptor instead.
func (*UserFishStatusUpdateRequest) Descriptor() ([]byte, []int) {
	return file_api_new_world_api_v1_internal_fish_proto_rawDescGZIP(), []int{0}
}

func (x *UserFishStatusUpdateRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *UserFishStatusUpdateRequest) GetFishId() int64 {
	if x != nil {
		return x.FishId
	}
	return 0
}

func (x *UserFishStatusUpdateRequest) GetFrom() v1.FishStatus {
	if x != nil {
		return x.From
	}
	return v1.FishStatus(0)
}

func (x *UserFishStatusUpdateRequest) GetTo() v1.FishStatus {
	if x != nil {
		return x.To
	}
	return v1.FishStatus(0)
}

type UserFishStatusUpdateResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Effect int32 `protobuf:"varint,1,opt,name=effect,proto3" json:"effect,omitempty"`
}

func (x *UserFishStatusUpdateResult) Reset() {
	*x = UserFishStatusUpdateResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_new_world_api_v1_internal_fish_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFishStatusUpdateResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFishStatusUpdateResult) ProtoMessage() {}

func (x *UserFishStatusUpdateResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_new_world_api_v1_internal_fish_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFishStatusUpdateResult.ProtoReflect.Descriptor instead.
func (*UserFishStatusUpdateResult) Descriptor() ([]byte, []int) {
	return file_api_new_world_api_v1_internal_fish_proto_rawDescGZIP(), []int{1}
}

func (x *UserFishStatusUpdateResult) GetEffect() int32 {
	if x != nil {
		return x.Effect
	}
	return 0
}

var File_api_new_world_api_v1_internal_fish_proto protoreflect.FileDescriptor

var file_api_new_world_api_v1_internal_fish_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x77, 0x2d, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x66, 0x69, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x65, 0x77, 0x5f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x09, 0x62, 0x69, 0x7a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x1b, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x73, 0x68,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x73, 0x68, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69, 0x73, 0x68, 0x49, 0x64, 0x12, 0x29,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x73, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x25, 0x0a, 0x02, 0x74, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x69, 0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x02, 0x74, 0x6f,
	0x22, 0x34, 0x0a, 0x1a, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x73, 0x68, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x32, 0x7b, 0x0a, 0x0c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x46, 0x69, 0x73, 0x68, 0x12, 0x6b, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x29,
	0x2e, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x46, 0x69, 0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6e, 0x65, 0x77, 0x5f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x73,
	0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x42, 0x2c, 0x5a, 0x2a, 0x6e, 0x65, 0x77, 0x2d, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2d, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x77, 0x2d, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_new_world_api_v1_internal_fish_proto_rawDescOnce sync.Once
	file_api_new_world_api_v1_internal_fish_proto_rawDescData = file_api_new_world_api_v1_internal_fish_proto_rawDesc
)

func file_api_new_world_api_v1_internal_fish_proto_rawDescGZIP() []byte {
	file_api_new_world_api_v1_internal_fish_proto_rawDescOnce.Do(func() {
		file_api_new_world_api_v1_internal_fish_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_new_world_api_v1_internal_fish_proto_rawDescData)
	})
	return file_api_new_world_api_v1_internal_fish_proto_rawDescData
}

var file_api_new_world_api_v1_internal_fish_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_new_world_api_v1_internal_fish_proto_goTypes = []interface{}{
	(*UserFishStatusUpdateRequest)(nil), // 0: new_world.v1.UserFishStatusUpdateRequest
	(*UserFishStatusUpdateResult)(nil),  // 1: new_world.v1.UserFishStatusUpdateResult
	(v1.FishStatus)(0),                  // 2: shared.v1.FishStatus
}
var file_api_new_world_api_v1_internal_fish_proto_depIdxs = []int32{
	2, // 0: new_world.v1.UserFishStatusUpdateRequest.from:type_name -> shared.v1.FishStatus
	2, // 1: new_world.v1.UserFishStatusUpdateRequest.to:type_name -> shared.v1.FishStatus
	0, // 2: new_world.v1.InternalFish.UpdateUserFishStatus:input_type -> new_world.v1.UserFishStatusUpdateRequest
	1, // 3: new_world.v1.InternalFish.UpdateUserFishStatus:output_type -> new_world.v1.UserFishStatusUpdateResult
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_new_world_api_v1_internal_fish_proto_init() }
func file_api_new_world_api_v1_internal_fish_proto_init() {
	if File_api_new_world_api_v1_internal_fish_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_new_world_api_v1_internal_fish_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFishStatusUpdateRequest); i {
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
		file_api_new_world_api_v1_internal_fish_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFishStatusUpdateResult); i {
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
			RawDescriptor: file_api_new_world_api_v1_internal_fish_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_new_world_api_v1_internal_fish_proto_goTypes,
		DependencyIndexes: file_api_new_world_api_v1_internal_fish_proto_depIdxs,
		MessageInfos:      file_api_new_world_api_v1_internal_fish_proto_msgTypes,
	}.Build()
	File_api_new_world_api_v1_internal_fish_proto = out.File
	file_api_new_world_api_v1_internal_fish_proto_rawDesc = nil
	file_api_new_world_api_v1_internal_fish_proto_goTypes = nil
	file_api_new_world_api_v1_internal_fish_proto_depIdxs = nil
}