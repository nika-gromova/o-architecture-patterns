// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: api.proto

package rules_v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RuleV1 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BaseLink      string                 `protobuf:"bytes,2,opt,name=base_link,json=baseLink,proto3" json:"base_link,omitempty"`
	DefaultLink   *string                `protobuf:"bytes,3,opt,name=default_link,json=defaultLink,proto3,oneof" json:"default_link,omitempty"`
	Redirects     []*RedirectV1          `protobuf:"bytes,4,rep,name=redirects,proto3" json:"redirects,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RuleV1) Reset() {
	*x = RuleV1{}
	mi := &file_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RuleV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleV1) ProtoMessage() {}

func (x *RuleV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleV1.ProtoReflect.Descriptor instead.
func (*RuleV1) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *RuleV1) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RuleV1) GetBaseLink() string {
	if x != nil {
		return x.BaseLink
	}
	return ""
}

func (x *RuleV1) GetDefaultLink() string {
	if x != nil && x.DefaultLink != nil {
		return *x.DefaultLink
	}
	return ""
}

func (x *RuleV1) GetRedirects() []*RedirectV1 {
	if x != nil {
		return x.Redirects
	}
	return nil
}

type RedirectV1 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Formula       *FormulaV1             `protobuf:"bytes,1,opt,name=formula,proto3" json:"formula,omitempty"`
	TargetLink    *TargetLinkV1          `protobuf:"bytes,2,opt,name=target_link,json=targetLink,proto3" json:"target_link,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RedirectV1) Reset() {
	*x = RedirectV1{}
	mi := &file_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RedirectV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedirectV1) ProtoMessage() {}

func (x *RedirectV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedirectV1.ProtoReflect.Descriptor instead.
func (*RedirectV1) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *RedirectV1) GetFormula() *FormulaV1 {
	if x != nil {
		return x.Formula
	}
	return nil
}

func (x *RedirectV1) GetTargetLink() *TargetLinkV1 {
	if x != nil {
		return x.TargetLink
	}
	return nil
}

type FormulaV1 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Expression    string                 `protobuf:"bytes,1,opt,name=expression,proto3" json:"expression,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FormulaV1) Reset() {
	*x = FormulaV1{}
	mi := &file_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FormulaV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FormulaV1) ProtoMessage() {}

func (x *FormulaV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FormulaV1.ProtoReflect.Descriptor instead.
func (*FormulaV1) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *FormulaV1) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

type TargetLinkV1 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Link          string                 `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TargetLinkV1) Reset() {
	*x = TargetLinkV1{}
	mi := &file_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TargetLinkV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetLinkV1) ProtoMessage() {}

func (x *TargetLinkV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetLinkV1.ProtoReflect.Descriptor instead.
func (*TargetLinkV1) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *TargetLinkV1) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type ListRulesV1Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRulesV1Request) Reset() {
	*x = ListRulesV1Request{}
	mi := &file_api_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRulesV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRulesV1Request) ProtoMessage() {}

func (x *ListRulesV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRulesV1Request.ProtoReflect.Descriptor instead.
func (*ListRulesV1Request) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

type ListRulesV1Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Rules         []*RuleV1              `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRulesV1Response) Reset() {
	*x = ListRulesV1Response{}
	mi := &file_api_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRulesV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRulesV1Response) ProtoMessage() {}

func (x *ListRulesV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRulesV1Response.ProtoReflect.Descriptor instead.
func (*ListRulesV1Response) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListRulesV1Response) GetRules() []*RuleV1 {
	if x != nil {
		return x.Rules
	}
	return nil
}

type CreateRuleV1Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Rule          *RuleV1                `protobuf:"bytes,1,opt,name=rule,proto3" json:"rule,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRuleV1Request) Reset() {
	*x = CreateRuleV1Request{}
	mi := &file_api_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRuleV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRuleV1Request) ProtoMessage() {}

func (x *CreateRuleV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRuleV1Request.ProtoReflect.Descriptor instead.
func (*CreateRuleV1Request) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *CreateRuleV1Request) GetRule() *RuleV1 {
	if x != nil {
		return x.Rule
	}
	return nil
}

type CreateRuleV1Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRuleV1Response) Reset() {
	*x = CreateRuleV1Response{}
	mi := &file_api_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRuleV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRuleV1Response) ProtoMessage() {}

func (x *CreateRuleV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRuleV1Response.ProtoReflect.Descriptor instead.
func (*CreateRuleV1Response) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

type UpdateRuleV1Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Redirects     []*RedirectV1          `protobuf:"bytes,2,rep,name=redirects,proto3" json:"redirects,omitempty"`
	DefaultLink   *TargetLinkV1          `protobuf:"bytes,3,opt,name=default_link,json=defaultLink,proto3,oneof" json:"default_link,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRuleV1Request) Reset() {
	*x = UpdateRuleV1Request{}
	mi := &file_api_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRuleV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRuleV1Request) ProtoMessage() {}

func (x *UpdateRuleV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRuleV1Request.ProtoReflect.Descriptor instead.
func (*UpdateRuleV1Request) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateRuleV1Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRuleV1Request) GetRedirects() []*RedirectV1 {
	if x != nil {
		return x.Redirects
	}
	return nil
}

func (x *UpdateRuleV1Request) GetDefaultLink() *TargetLinkV1 {
	if x != nil {
		return x.DefaultLink
	}
	return nil
}

type UpdateRuleV1Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Rule          *RuleV1                `protobuf:"bytes,1,opt,name=rule,proto3" json:"rule,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRuleV1Response) Reset() {
	*x = UpdateRuleV1Response{}
	mi := &file_api_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRuleV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRuleV1Response) ProtoMessage() {}

func (x *UpdateRuleV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRuleV1Response.ProtoReflect.Descriptor instead.
func (*UpdateRuleV1Response) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateRuleV1Response) GetRule() *RuleV1 {
	if x != nil {
		return x.Rule
	}
	return nil
}

type DeleteRuleV1Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRuleV1Request) Reset() {
	*x = DeleteRuleV1Request{}
	mi := &file_api_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRuleV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRuleV1Request) ProtoMessage() {}

func (x *DeleteRuleV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRuleV1Request.ProtoReflect.Descriptor instead.
func (*DeleteRuleV1Request) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteRuleV1Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteRuleV1Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRuleV1Response) Reset() {
	*x = DeleteRuleV1Response{}
	mi := &file_api_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRuleV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRuleV1Response) ProtoMessage() {}

func (x *DeleteRuleV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRuleV1Response.ProtoReflect.Descriptor instead.
func (*DeleteRuleV1Response) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{11}
}

type GetRuleV1Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRuleV1Request) Reset() {
	*x = GetRuleV1Request{}
	mi := &file_api_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRuleV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuleV1Request) ProtoMessage() {}

func (x *GetRuleV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRuleV1Request.ProtoReflect.Descriptor instead.
func (*GetRuleV1Request) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{12}
}

func (x *GetRuleV1Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetRuleV1Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Rule          *RuleV1                `protobuf:"bytes,1,opt,name=rule,proto3" json:"rule,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRuleV1Response) Reset() {
	*x = GetRuleV1Response{}
	mi := &file_api_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRuleV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuleV1Response) ProtoMessage() {}

func (x *GetRuleV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRuleV1Response.ProtoReflect.Descriptor instead.
func (*GetRuleV1Response) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{13}
}

func (x *GetRuleV1Response) GetRule() *RuleV1 {
	if x != nil {
		return x.Rule
	}
	return nil
}

type GetRedirectV1Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRedirectV1Request) Reset() {
	*x = GetRedirectV1Request{}
	mi := &file_api_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRedirectV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRedirectV1Request) ProtoMessage() {}

func (x *GetRedirectV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRedirectV1Request.ProtoReflect.Descriptor instead.
func (*GetRedirectV1Request) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{14}
}

func (x *GetRedirectV1Request) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetRedirectV1Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRedirectV1Response) Reset() {
	*x = GetRedirectV1Response{}
	mi := &file_api_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRedirectV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRedirectV1Response) ProtoMessage() {}

func (x *GetRedirectV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRedirectV1Response.ProtoReflect.Descriptor instead.
func (*GetRedirectV1Response) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{15}
}

func (x *GetRedirectV1Response) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = string([]byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa3, 0x01, 0x0a, 0x06, 0x52, 0x75, 0x6c, 0x65, 0x56, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x26, 0x0a, 0x0c,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x09, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e,
	0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x56, 0x31, 0x52, 0x09, 0x72, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x73, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x6e, 0x0a, 0x0a, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x56, 0x31, 0x12, 0x2a, 0x0a, 0x07, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x46, 0x6f,
	0x72, 0x6d, 0x75, 0x6c, 0x61, 0x56, 0x31, 0x52, 0x07, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61,
	0x12, 0x34, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x56, 0x31, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x2b, 0x0a, 0x09, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c,
	0x61, 0x56, 0x31, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x22, 0x0a, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x56, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x75, 0x6c, 0x65, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3a, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65,
	0x56, 0x31, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x38, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x56, 0x31, 0x52, 0x04, 0x72,
	0x75, 0x6c, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c,
	0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa8, 0x01, 0x0a, 0x13,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x72, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x56, 0x31, 0x52, 0x09, 0x72,
	0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x0c, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x56, 0x31, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4c, 0x69,
	0x6e, 0x6b, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x39, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x75, 0x6c, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72,
	0x75, 0x6c, 0x65, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x56, 0x31, 0x52, 0x04, 0x72, 0x75, 0x6c,
	0x65, 0x22, 0x29, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x16, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x21, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x56, 0x31, 0x52, 0x04,
	0x72, 0x75, 0x6c, 0x65, 0x22, 0x28, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x29,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x56, 0x31, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0xc4, 0x04, 0x0a, 0x05, 0x52, 0x75,
	0x6c, 0x65, 0x73, 0x12, 0x5b, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73,
	0x56, 0x31, 0x12, 0x19, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x75, 0x6c, 0x65, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x56,
	0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x12, 0x61, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x56, 0x31,
	0x12, 0x1a, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x75, 0x6c, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x72,
	0x75, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x56,
	0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75,
	0x6c, 0x65, 0x73, 0x12, 0x68, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c,
	0x65, 0x56, 0x31, 0x12, 0x1a, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x75,
	0x6c, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x1a, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x65, 0x0a,
	0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x56, 0x31, 0x12, 0x1a, 0x2e,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x56, 0x31, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x2a, 0x14,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x7d, 0x12, 0x5c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x56,
	0x31, 0x12, 0x17, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c,
	0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x12, 0x4c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x56, 0x31, 0x12, 0x1b, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0xfa, 0x01, 0x92, 0x41, 0xb0, 0x01, 0x5a, 0x9b, 0x01, 0x0a, 0x98, 0x01, 0x0a, 0x0a, 0x42,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x89, 0x01, 0x08, 0x03, 0x20, 0x02,
	0x28, 0x02, 0x32, 0x3f, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x68, 0x6f, 0x73, 0x74, 0x3a, 0x37, 0x30, 0x38, 0x30, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x73,
	0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x3a, 0x40, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x37, 0x30, 0x38, 0x30, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x6d,
	0x73, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x6b, 0x61, 0x2d, 0x67, 0x72, 0x6f, 0x6d, 0x6f, 0x76, 0x61,
	0x2f, 0x6f, 0x2d, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2d,
	0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData []byte
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_proto_rawDesc), len(file_api_proto_rawDesc)))
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_api_proto_goTypes = []any{
	(*RuleV1)(nil),                // 0: rules.RuleV1
	(*RedirectV1)(nil),            // 1: rules.RedirectV1
	(*FormulaV1)(nil),             // 2: rules.FormulaV1
	(*TargetLinkV1)(nil),          // 3: rules.TargetLinkV1
	(*ListRulesV1Request)(nil),    // 4: rules.ListRulesV1Request
	(*ListRulesV1Response)(nil),   // 5: rules.ListRulesV1Response
	(*CreateRuleV1Request)(nil),   // 6: rules.CreateRuleV1Request
	(*CreateRuleV1Response)(nil),  // 7: rules.CreateRuleV1Response
	(*UpdateRuleV1Request)(nil),   // 8: rules.UpdateRuleV1Request
	(*UpdateRuleV1Response)(nil),  // 9: rules.UpdateRuleV1Response
	(*DeleteRuleV1Request)(nil),   // 10: rules.DeleteRuleV1Request
	(*DeleteRuleV1Response)(nil),  // 11: rules.DeleteRuleV1Response
	(*GetRuleV1Request)(nil),      // 12: rules.GetRuleV1Request
	(*GetRuleV1Response)(nil),     // 13: rules.GetRuleV1Response
	(*GetRedirectV1Request)(nil),  // 14: rules.GetRedirectV1Request
	(*GetRedirectV1Response)(nil), // 15: rules.GetRedirectV1Response
}
var file_api_proto_depIdxs = []int32{
	1,  // 0: rules.RuleV1.redirects:type_name -> rules.RedirectV1
	2,  // 1: rules.RedirectV1.formula:type_name -> rules.FormulaV1
	3,  // 2: rules.RedirectV1.target_link:type_name -> rules.TargetLinkV1
	0,  // 3: rules.ListRulesV1Response.rules:type_name -> rules.RuleV1
	0,  // 4: rules.CreateRuleV1Request.rule:type_name -> rules.RuleV1
	1,  // 5: rules.UpdateRuleV1Request.redirects:type_name -> rules.RedirectV1
	3,  // 6: rules.UpdateRuleV1Request.default_link:type_name -> rules.TargetLinkV1
	0,  // 7: rules.UpdateRuleV1Response.rule:type_name -> rules.RuleV1
	0,  // 8: rules.GetRuleV1Response.rule:type_name -> rules.RuleV1
	4,  // 9: rules.Rules.ListRulesV1:input_type -> rules.ListRulesV1Request
	6,  // 10: rules.Rules.CreateRuleV1:input_type -> rules.CreateRuleV1Request
	8,  // 11: rules.Rules.UpdateRuleV1:input_type -> rules.UpdateRuleV1Request
	10, // 12: rules.Rules.DeleteRuleV1:input_type -> rules.DeleteRuleV1Request
	12, // 13: rules.Rules.GetRuleV1:input_type -> rules.GetRuleV1Request
	14, // 14: rules.Rules.GetRedirectV1:input_type -> rules.GetRedirectV1Request
	5,  // 15: rules.Rules.ListRulesV1:output_type -> rules.ListRulesV1Response
	7,  // 16: rules.Rules.CreateRuleV1:output_type -> rules.CreateRuleV1Response
	9,  // 17: rules.Rules.UpdateRuleV1:output_type -> rules.UpdateRuleV1Response
	11, // 18: rules.Rules.DeleteRuleV1:output_type -> rules.DeleteRuleV1Response
	13, // 19: rules.Rules.GetRuleV1:output_type -> rules.GetRuleV1Response
	15, // 20: rules.Rules.GetRedirectV1:output_type -> rules.GetRedirectV1Response
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	file_api_proto_msgTypes[0].OneofWrappers = []any{}
	file_api_proto_msgTypes[8].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_proto_rawDesc), len(file_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
