// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/scm/pb/rpc.proto

package scm

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

type QueryProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 仓库类型
	// @gotags: json:"type" validate:"required"
	Type TYPE `protobuf:"varint,1,opt,name=type,proto3,enum=infraboard.mcenter.scm.TYPE" json:"type" validate:"required"`
	// 仓库地址
	// @gotags: json:"address" validate:"required"
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address" validate:"required"`
	// 仓库访问凭证
	// @gotags: json:"token" validate:"required"
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token" validate:"required"`
}

func (x *QueryProjectRequest) Reset() {
	*x = QueryProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_scm_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProjectRequest) ProtoMessage() {}

func (x *QueryProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_scm_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProjectRequest.ProtoReflect.Descriptor instead.
func (*QueryProjectRequest) Descriptor() ([]byte, []int) {
	return file_apps_scm_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryProjectRequest) GetType() TYPE {
	if x != nil {
		return x.Type
	}
	return TYPE_GITLAB
}

func (x *QueryProjectRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *QueryProjectRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_apps_scm_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_scm_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63, 0x6d, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x63, 0x6d, 0x1a,
	0x15, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63, 0x6d, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x63, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x73, 0x63, 0x6d, 0x2e, 0x54, 0x59, 0x50, 0x45, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32,
	0xc1, 0x01, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x5f, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x63, 0x6d,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x63, 0x6d, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x74, 0x12, 0x59, 0x0a, 0x0b, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x63, 0x6d,
	0x2e, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x24, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x73, 0x63, 0x6d, 0x2e, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_scm_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_scm_pb_rpc_proto_rawDescData = file_apps_scm_pb_rpc_proto_rawDesc
)

func file_apps_scm_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_scm_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_scm_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_scm_pb_rpc_proto_rawDescData)
	})
	return file_apps_scm_pb_rpc_proto_rawDescData
}

var file_apps_scm_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_apps_scm_pb_rpc_proto_goTypes = []interface{}{
	(*QueryProjectRequest)(nil), // 0: infraboard.mcenter.scm.QueryProjectRequest
	(TYPE)(0),                   // 1: infraboard.mcenter.scm.TYPE
	(*WebHookEvent)(nil),        // 2: infraboard.mcenter.scm.WebHookEvent
	(*ProjectSet)(nil),          // 3: infraboard.mcenter.scm.ProjectSet
}
var file_apps_scm_pb_rpc_proto_depIdxs = []int32{
	1, // 0: infraboard.mcenter.scm.QueryProjectRequest.type:type_name -> infraboard.mcenter.scm.TYPE
	0, // 1: infraboard.mcenter.scm.RPC.QueryProject:input_type -> infraboard.mcenter.scm.QueryProjectRequest
	2, // 2: infraboard.mcenter.scm.RPC.HandleEvent:input_type -> infraboard.mcenter.scm.WebHookEvent
	3, // 3: infraboard.mcenter.scm.RPC.QueryProject:output_type -> infraboard.mcenter.scm.ProjectSet
	2, // 4: infraboard.mcenter.scm.RPC.HandleEvent:output_type -> infraboard.mcenter.scm.WebHookEvent
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_scm_pb_rpc_proto_init() }
func file_apps_scm_pb_rpc_proto_init() {
	if File_apps_scm_pb_rpc_proto != nil {
		return
	}
	file_apps_scm_pb_scm_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_scm_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProjectRequest); i {
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
			RawDescriptor: file_apps_scm_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_scm_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_scm_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_scm_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_scm_pb_rpc_proto = out.File
	file_apps_scm_pb_rpc_proto_rawDesc = nil
	file_apps_scm_pb_rpc_proto_goTypes = nil
	file_apps_scm_pb_rpc_proto_depIdxs = nil
}