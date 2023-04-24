// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: apps/notify/pb/notify.proto

package notify

import (
	resource "github.com/infraboard/mcube/pb/resource"
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

type PROVIDER int32

const (
	// 腾讯短信服务
	PROVIDER_TENCENT PROVIDER = 0
	// 阿里短信服务
	PROVIDER_ALI PROVIDER = 1
)

// Enum value maps for PROVIDER.
var (
	PROVIDER_name = map[int32]string{
		0: "TENCENT",
		1: "ALI",
	}
	PROVIDER_value = map[string]int32{
		"TENCENT": 0,
		"ALI":     1,
	}
)

func (x PROVIDER) Enum() *PROVIDER {
	p := new(PROVIDER)
	*p = x
	return p
}

func (x PROVIDER) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PROVIDER) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_notify_pb_notify_proto_enumTypes[0].Descriptor()
}

func (PROVIDER) Type() protoreflect.EnumType {
	return &file_apps_notify_pb_notify_proto_enumTypes[0]
}

func (x PROVIDER) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PROVIDER.Descriptor instead.
func (PROVIDER) EnumDescriptor() ([]byte, []int) {
	return file_apps_notify_pb_notify_proto_rawDescGZIP(), []int{0}
}

type NOTIFY_TYPE int32

const (
	// 邮件通知
	NOTIFY_TYPE_MAIL NOTIFY_TYPE = 0
	// 短信通知
	NOTIFY_TYPE_SMS NOTIFY_TYPE = 1
	// 语音通知
	NOTIFY_TYPE_VOICE NOTIFY_TYPE = 2
	// IM个人消息
	NOTIFY_TYPE_IM NOTIFY_TYPE = 3
)

// Enum value maps for NOTIFY_TYPE.
var (
	NOTIFY_TYPE_name = map[int32]string{
		0: "MAIL",
		1: "SMS",
		2: "VOICE",
		3: "IM",
	}
	NOTIFY_TYPE_value = map[string]int32{
		"MAIL":  0,
		"SMS":   1,
		"VOICE": 2,
		"IM":    3,
	}
)

func (x NOTIFY_TYPE) Enum() *NOTIFY_TYPE {
	p := new(NOTIFY_TYPE)
	*p = x
	return p
}

func (x NOTIFY_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NOTIFY_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_notify_pb_notify_proto_enumTypes[1].Descriptor()
}

func (NOTIFY_TYPE) Type() protoreflect.EnumType {
	return &file_apps_notify_pb_notify_proto_enumTypes[1]
}

func (x NOTIFY_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NOTIFY_TYPE.Descriptor instead.
func (NOTIFY_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_apps_notify_pb_notify_proto_rawDescGZIP(), []int{1}
}

type SendNotifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 域
	// @gotags: json:"domain" bson:"domain" validate:"required"
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain" bson:"domain" validate:"required"`
	// 空间
	// @gotags: json:"namespace" bson:"namespace" validate:"required"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace" bson:"namespace" validate:"required"`
	// 通知类型
	// @gotags: json:"notify_tye" bson:"notify_tye"
	NotifyTye NOTIFY_TYPE `protobuf:"varint,3,opt,name=notify_tye,json=notifyTye,proto3,enum=infraboard.mcenter.notify.NOTIFY_TYPE" json:"notify_tye" bson:"notify_tye"`
	// 用户
	// @gotags: json:"users" bson:"users" validate:"required"
	Users []string `protobuf:"bytes,4,rep,name=users,proto3" json:"users" bson:"users" validate:"required"`
	// 消息标题
	// @gotags: json:"title" bson:"title" validate:"required"
	Title string `protobuf:"bytes,5,opt,name=title,proto3" json:"title" bson:"title" validate:"required"`
	// 消息内容格式, 格式如果不填由 由具体渠道自行适配
	// @gotags: json:"content_type" bson:"content_type"
	ContentType string `protobuf:"bytes,10,opt,name=content_type,json=contentType,proto3" json:"content_type" bson:"content_type"`
	// 消息内容
	// @gotags: json:"content" bson:"content"
	Content string `protobuf:"bytes,6,opt,name=content,proto3" json:"content" bson:"content"`
	// 用户的 session 内容，腾讯 server 回包中会原样返回
	// @gotags: bson:"session_context" json:"session_context"
	SessionContext string `protobuf:"bytes,9,opt,name=session_context,json=sessionContext,proto3" json:"session_context" bson:"session_context"`
	// 短信通知请求参数
	// @gotags: json:"sms_request" bson:"sms_request"
	SmsRequest *SMSRequest `protobuf:"bytes,7,opt,name=sms_request,json=smsRequest,proto3" json:"sms_request" bson:"sms_request"`
	// 语音通知请求参数
	// @gotags: json:"voice_request" bson:"voice_request"
	VoiceRequest *VoiceRequest `protobuf:"bytes,8,opt,name=voice_request,json=voiceRequest,proto3" json:"voice_request" bson:"voice_request"`
	// 标签, 方便查找
	// @gotags: json:"labels" bson:"labels"
	Labels map[string]string `protobuf:"bytes,14,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"labels"`
	// 扩展参数
	// @gotags: json:"extra" bson:"extra"
	Extra map[string]string `protobuf:"bytes,15,rep,name=extra,proto3" json:"extra" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"extra"`
}

func (x *SendNotifyRequest) Reset() {
	*x = SendNotifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_notify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendNotifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendNotifyRequest) ProtoMessage() {}

func (x *SendNotifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_notify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendNotifyRequest.ProtoReflect.Descriptor instead.
func (*SendNotifyRequest) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_notify_proto_rawDescGZIP(), []int{0}
}

func (x *SendNotifyRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *SendNotifyRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *SendNotifyRequest) GetNotifyTye() NOTIFY_TYPE {
	if x != nil {
		return x.NotifyTye
	}
	return NOTIFY_TYPE_MAIL
}

func (x *SendNotifyRequest) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *SendNotifyRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SendNotifyRequest) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *SendNotifyRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SendNotifyRequest) GetSessionContext() string {
	if x != nil {
		return x.SessionContext
	}
	return ""
}

func (x *SendNotifyRequest) GetSmsRequest() *SMSRequest {
	if x != nil {
		return x.SmsRequest
	}
	return nil
}

func (x *SendNotifyRequest) GetVoiceRequest() *VoiceRequest {
	if x != nil {
		return x.VoiceRequest
	}
	return nil
}

func (x *SendNotifyRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *SendNotifyRequest) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type SMSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 短信模版的Id
	// @gotags: bson:"template_id" json:"template_id"
	TemplateId string `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id" bson:"template_id"`
	// 模版参数
	// @gotags: bson:"template_params" json:"template_params"
	TemplateParams []string `protobuf:"bytes,2,rep,name=template_params,json=templateParams,proto3" json:"template_params" bson:"template_params"`
}

func (x *SMSRequest) Reset() {
	*x = SMSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_notify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SMSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SMSRequest) ProtoMessage() {}

func (x *SMSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_notify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SMSRequest.ProtoReflect.Descriptor instead.
func (*SMSRequest) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_notify_proto_rawDescGZIP(), []int{1}
}

func (x *SMSRequest) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *SMSRequest) GetTemplateParams() []string {
	if x != nil {
		return x.TemplateParams
	}
	return nil
}

type VoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 短信模版的Id
	// @gotags: bson:"template_id" json:"template_id"
	TemplateId string `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id" bson:"template_id"`
	// 模版参数
	// @gotags: bson:"template_params" json:"template_params"
	TemplateParams []string `protobuf:"bytes,2,rep,name=template_params,json=templateParams,proto3" json:"template_params" bson:"template_params"`
	// 播放次数，可选，最多3次，默认2次
	// @gotags: bson:"play_times" json:"play_times"
	PlayTimes uint64 `protobuf:"varint,4,opt,name=play_times,json=playTimes,proto3" json:"play_times" bson:"play_times"`
}

func (x *VoiceRequest) Reset() {
	*x = VoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_notify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoiceRequest) ProtoMessage() {}

func (x *VoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_notify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoiceRequest.ProtoReflect.Descriptor instead.
func (*VoiceRequest) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_notify_proto_rawDescGZIP(), []int{2}
}

func (x *VoiceRequest) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *VoiceRequest) GetTemplateParams() []string {
	if x != nil {
		return x.TemplateParams
	}
	return nil
}

func (x *VoiceRequest) GetPlayTimes() uint64 {
	if x != nil {
		return x.PlayTimes
	}
	return 0
}

type RecordSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数
	// @gotags: json:"total" bson:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// 发生记录
	// @gotags: json:"items" bson:"items"
	Items []*Record `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *RecordSet) Reset() {
	*x = RecordSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_notify_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordSet) ProtoMessage() {}

func (x *RecordSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_notify_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordSet.ProtoReflect.Descriptor instead.
func (*RecordSet) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_notify_proto_rawDescGZIP(), []int{3}
}

func (x *RecordSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *RecordSet) GetItems() []*Record {
	if x != nil {
		return x.Items
	}
	return nil
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 元信息
	// @gotags: bson:",inline" json:"meta"
	Meta *resource.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta" bson:",inline"`
	// 请求
	// @gotags: json:"request" bson:",inline"
	Request *SendNotifyRequest `protobuf:"bytes,2,opt,name=request,proto3" json:"request" bson:",inline"`
	// 发送响应
	// @gotags: json:"response" bson:"response"
	Response []*SendResponse `protobuf:"bytes,15,rep,name=response,proto3" json:"response" bson:"response"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_notify_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_notify_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_notify_proto_rawDescGZIP(), []int{4}
}

func (x *Record) GetMeta() *resource.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Record) GetRequest() *SendNotifyRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Record) GetResponse() []*SendResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type SendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户
	// @gotags: json:"user" bson:"user"
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user" bson:"user"`
	// 通知目标, 如果是邮件就是邮箱, 电话就是电话号码, 一般不用填写, 通过查询用户信息获取
	// @gotags: json:"target" bson:"target"
	Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target" bson:"target"`
	// 是否发生成功
	// @gotags: json:"success" bson:"success"
	Success bool `protobuf:"varint,3,opt,name=success,proto3" json:"success" bson:"success"`
	// 失败时的错误
	// @gotags: json:"message" bson:"message"
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message" bson:"message"`
	// 语音通知响应
	// @gotags: json:"voice_response" bson:"voice_response,omitempty"
	VoiceResponse *VoiceResponse `protobuf:"bytes,5,opt,name=voice_response,json=voiceResponse,proto3" json:"voice_response" bson:"voice_response,omitempty"`
	// 其他信息
	// @gotags: json:"extra" bson:"extra"
	Extra map[string]string `protobuf:"bytes,15,rep,name=extra,proto3" json:"extra" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"extra"`
}

func (x *SendResponse) Reset() {
	*x = SendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_notify_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResponse) ProtoMessage() {}

func (x *SendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_notify_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResponse.ProtoReflect.Descriptor instead.
func (*SendResponse) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_notify_proto_rawDescGZIP(), []int{5}
}

func (x *SendResponse) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *SendResponse) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *SendResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SendResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SendResponse) GetVoiceResponse() *VoiceResponse {
	if x != nil {
		return x.VoiceResponse
	}
	return nil
}

func (x *SendResponse) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type VoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 呼叫Id
	// @gotags: bson:"call_id" json:"call_id"
	CallId string `protobuf:"bytes,1,opt,name=call_id,json=callId,proto3" json:"call_id" bson:"call_id"`
	// 用户的 session 内容，腾讯 server 回包中会原样返回
	// @gotags: bson:"session_context" json:"session_context"
	SessionContext string `protobuf:"bytes,2,opt,name=session_context,json=sessionContext,proto3" json:"session_context" bson:"session_context"`
}

func (x *VoiceResponse) Reset() {
	*x = VoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_notify_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoiceResponse) ProtoMessage() {}

func (x *VoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_notify_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoiceResponse.ProtoReflect.Descriptor instead.
func (*VoiceResponse) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_notify_proto_rawDescGZIP(), []int{6}
}

func (x *VoiceResponse) GetCallId() string {
	if x != nil {
		return x.CallId
	}
	return ""
}

func (x *VoiceResponse) GetSessionContext() string {
	if x != nil {
		return x.SessionContext
	}
	return ""
}

var File_apps_notify_pb_notify_proto protoreflect.FileDescriptor

var file_apps_notify_pb_notify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x70, 0x62,
	0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x1a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f,
	0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x05, 0x0a,
	0x11, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x5f, 0x74, 0x79, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x79, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x46, 0x0a, 0x0b, 0x73, 0x6d, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x53, 0x4d, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x73,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x0d, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x4d, 0x0a, 0x05, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x56, 0x0a,
	0x0a, 0x53, 0x4d, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x77, 0x0a, 0x0c, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x22, 0x5a,
	0x0a, 0x09, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x37, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xca, 0x01, 0x0a, 0x06, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x33, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x46, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x43, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x0f,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xc3, 0x02, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0d, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x05, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x51, 0x0a,
	0x0d, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x2a, 0x20, 0x0a, 0x08, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x12, 0x0b, 0x0a, 0x07,
	0x54, 0x45, 0x4e, 0x43, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x49,
	0x10, 0x01, 0x2a, 0x33, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x53,
	0x4d, 0x53, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x10, 0x02, 0x12,
	0x06, 0x0a, 0x02, 0x49, 0x4d, 0x10, 0x03, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_notify_pb_notify_proto_rawDescOnce sync.Once
	file_apps_notify_pb_notify_proto_rawDescData = file_apps_notify_pb_notify_proto_rawDesc
)

func file_apps_notify_pb_notify_proto_rawDescGZIP() []byte {
	file_apps_notify_pb_notify_proto_rawDescOnce.Do(func() {
		file_apps_notify_pb_notify_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_notify_pb_notify_proto_rawDescData)
	})
	return file_apps_notify_pb_notify_proto_rawDescData
}

var file_apps_notify_pb_notify_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_apps_notify_pb_notify_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_apps_notify_pb_notify_proto_goTypes = []interface{}{
	(PROVIDER)(0),             // 0: infraboard.mcenter.notify.PROVIDER
	(NOTIFY_TYPE)(0),          // 1: infraboard.mcenter.notify.NOTIFY_TYPE
	(*SendNotifyRequest)(nil), // 2: infraboard.mcenter.notify.SendNotifyRequest
	(*SMSRequest)(nil),        // 3: infraboard.mcenter.notify.SMSRequest
	(*VoiceRequest)(nil),      // 4: infraboard.mcenter.notify.VoiceRequest
	(*RecordSet)(nil),         // 5: infraboard.mcenter.notify.RecordSet
	(*Record)(nil),            // 6: infraboard.mcenter.notify.Record
	(*SendResponse)(nil),      // 7: infraboard.mcenter.notify.SendResponse
	(*VoiceResponse)(nil),     // 8: infraboard.mcenter.notify.VoiceResponse
	nil,                       // 9: infraboard.mcenter.notify.SendNotifyRequest.LabelsEntry
	nil,                       // 10: infraboard.mcenter.notify.SendNotifyRequest.ExtraEntry
	nil,                       // 11: infraboard.mcenter.notify.SendResponse.ExtraEntry
	(*resource.Meta)(nil),     // 12: infraboard.mcube.resource.Meta
}
var file_apps_notify_pb_notify_proto_depIdxs = []int32{
	1,  // 0: infraboard.mcenter.notify.SendNotifyRequest.notify_tye:type_name -> infraboard.mcenter.notify.NOTIFY_TYPE
	3,  // 1: infraboard.mcenter.notify.SendNotifyRequest.sms_request:type_name -> infraboard.mcenter.notify.SMSRequest
	4,  // 2: infraboard.mcenter.notify.SendNotifyRequest.voice_request:type_name -> infraboard.mcenter.notify.VoiceRequest
	9,  // 3: infraboard.mcenter.notify.SendNotifyRequest.labels:type_name -> infraboard.mcenter.notify.SendNotifyRequest.LabelsEntry
	10, // 4: infraboard.mcenter.notify.SendNotifyRequest.extra:type_name -> infraboard.mcenter.notify.SendNotifyRequest.ExtraEntry
	6,  // 5: infraboard.mcenter.notify.RecordSet.items:type_name -> infraboard.mcenter.notify.Record
	12, // 6: infraboard.mcenter.notify.Record.meta:type_name -> infraboard.mcube.resource.Meta
	2,  // 7: infraboard.mcenter.notify.Record.request:type_name -> infraboard.mcenter.notify.SendNotifyRequest
	7,  // 8: infraboard.mcenter.notify.Record.response:type_name -> infraboard.mcenter.notify.SendResponse
	8,  // 9: infraboard.mcenter.notify.SendResponse.voice_response:type_name -> infraboard.mcenter.notify.VoiceResponse
	11, // 10: infraboard.mcenter.notify.SendResponse.extra:type_name -> infraboard.mcenter.notify.SendResponse.ExtraEntry
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_apps_notify_pb_notify_proto_init() }
func file_apps_notify_pb_notify_proto_init() {
	if File_apps_notify_pb_notify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_notify_pb_notify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendNotifyRequest); i {
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
		file_apps_notify_pb_notify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SMSRequest); i {
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
		file_apps_notify_pb_notify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoiceRequest); i {
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
		file_apps_notify_pb_notify_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordSet); i {
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
		file_apps_notify_pb_notify_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
		file_apps_notify_pb_notify_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResponse); i {
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
		file_apps_notify_pb_notify_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoiceResponse); i {
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
			RawDescriptor: file_apps_notify_pb_notify_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_notify_pb_notify_proto_goTypes,
		DependencyIndexes: file_apps_notify_pb_notify_proto_depIdxs,
		EnumInfos:         file_apps_notify_pb_notify_proto_enumTypes,
		MessageInfos:      file_apps_notify_pb_notify_proto_msgTypes,
	}.Build()
	File_apps_notify_pb_notify_proto = out.File
	file_apps_notify_pb_notify_proto_rawDesc = nil
	file_apps_notify_pb_notify_proto_goTypes = nil
	file_apps_notify_pb_notify_proto_depIdxs = nil
}
