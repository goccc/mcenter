// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: apps/resource/pb/resource.proto

package resource

import (
	request "github.com/infraboard/mcube/http/request"
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

// QueryResourceRequest todo
type QueryResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// @gotags: json:"service_ids"
	ServiceIds []string `protobuf:"bytes,2,rep,name=service_ids,json=serviceIds,proto3" json:"service_ids"`
	// @gotags: json:"resources"
	Resources []string `protobuf:"bytes,3,rep,name=resources,proto3" json:"resources"`
	// @gotags: json:"permission_enable"
	PermissionEnable *bool `protobuf:"varint,4,opt,name=permission_enable,json=permissionEnable,proto3,oneof" json:"permission_enable"`
}

func (x *QueryResourceRequest) Reset() {
	*x = QueryResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_resource_pb_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResourceRequest) ProtoMessage() {}

func (x *QueryResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_resource_pb_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResourceRequest.ProtoReflect.Descriptor instead.
func (*QueryResourceRequest) Descriptor() ([]byte, []int) {
	return file_apps_resource_pb_resource_proto_rawDescGZIP(), []int{0}
}

func (x *QueryResourceRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryResourceRequest) GetServiceIds() []string {
	if x != nil {
		return x.ServiceIds
	}
	return nil
}

func (x *QueryResourceRequest) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *QueryResourceRequest) GetPermissionEnable() bool {
	if x != nil && x.PermissionEnable != nil {
		return *x.PermissionEnable
	}
	return false
}

// Resource todo
type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源所属服务名称
	// @gotags: json:"service_id"
	ServiceId string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id"`
	// 资源名称
	// @gotags: json:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// 资源支持操作的path
	// @gotags: json:"paths"
	Paths []string `protobuf:"bytes,3,rep,name=paths,proto3" json:"paths"`
	// 支持操作的方法
	// @gotags: json:"methods"
	Methods []string `protobuf:"bytes,4,rep,name=methods,proto3" json:"methods"`
	// 支持操作的函数
	// @gotags: json:"functions"
	Functions []string `protobuf:"bytes,5,rep,name=functions,proto3" json:"functions"`
	// 支持操作的动作
	// @gotags: json:"actions"
	Actions []string `protobuf:"bytes,6,rep,name=actions,proto3" json:"actions"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_resource_pb_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_apps_resource_pb_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_apps_resource_pb_resource_proto_rawDescGZIP(), []int{1}
}

func (x *Resource) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *Resource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Resource) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *Resource) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

func (x *Resource) GetFunctions() []string {
	if x != nil {
		return x.Functions
	}
	return nil
}

func (x *Resource) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

type ResourceSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*Resource `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *ResourceSet) Reset() {
	*x = ResourceSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_resource_pb_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceSet) ProtoMessage() {}

func (x *ResourceSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_resource_pb_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceSet.ProtoReflect.Descriptor instead.
func (*ResourceSet) Descriptor() ([]byte, []int) {
	return file_apps_resource_pb_resource_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ResourceSet) GetItems() []*Resource {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_apps_resource_pb_resource_proto protoreflect.FileDescriptor

var file_apps_resource_pb_resource_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70,
	0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5,
	0x01, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x30,
	0x0a, 0x11, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x10, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x14, 0x0a, 0x12, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xa5, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x60,
	0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x3b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x32, 0x74, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x6d, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x31, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x53, 0x65, 0x74, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_resource_pb_resource_proto_rawDescOnce sync.Once
	file_apps_resource_pb_resource_proto_rawDescData = file_apps_resource_pb_resource_proto_rawDesc
)

func file_apps_resource_pb_resource_proto_rawDescGZIP() []byte {
	file_apps_resource_pb_resource_proto_rawDescOnce.Do(func() {
		file_apps_resource_pb_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_resource_pb_resource_proto_rawDescData)
	})
	return file_apps_resource_pb_resource_proto_rawDescData
}

var file_apps_resource_pb_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apps_resource_pb_resource_proto_goTypes = []interface{}{
	(*QueryResourceRequest)(nil), // 0: infraboard.mcenter.resource.QueryResourceRequest
	(*Resource)(nil),             // 1: infraboard.mcenter.resource.Resource
	(*ResourceSet)(nil),          // 2: infraboard.mcenter.resource.ResourceSet
	(*request.PageRequest)(nil),  // 3: infraboard.mcube.page.PageRequest
}
var file_apps_resource_pb_resource_proto_depIdxs = []int32{
	3, // 0: infraboard.mcenter.resource.QueryResourceRequest.page:type_name -> infraboard.mcube.page.PageRequest
	1, // 1: infraboard.mcenter.resource.ResourceSet.items:type_name -> infraboard.mcenter.resource.Resource
	0, // 2: infraboard.mcenter.resource.RPC.QueryResources:input_type -> infraboard.mcenter.resource.QueryResourceRequest
	2, // 3: infraboard.mcenter.resource.RPC.QueryResources:output_type -> infraboard.mcenter.resource.ResourceSet
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_apps_resource_pb_resource_proto_init() }
func file_apps_resource_pb_resource_proto_init() {
	if File_apps_resource_pb_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_resource_pb_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResourceRequest); i {
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
		file_apps_resource_pb_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_apps_resource_pb_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceSet); i {
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
	file_apps_resource_pb_resource_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_resource_pb_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_resource_pb_resource_proto_goTypes,
		DependencyIndexes: file_apps_resource_pb_resource_proto_depIdxs,
		MessageInfos:      file_apps_resource_pb_resource_proto_msgTypes,
	}.Build()
	File_apps_resource_pb_resource_proto = out.File
	file_apps_resource_pb_resource_proto_rawDesc = nil
	file_apps_resource_pb_resource_proto_goTypes = nil
	file_apps_resource_pb_resource_proto_depIdxs = nil
}