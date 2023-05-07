// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: policy.proto

package pb

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

// ListUsersRequest defines ListUsers request struct.
type ListPoliciesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset *int64 `protobuf:"varint,1,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
	Limit  *int64 `protobuf:"varint,2,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
}

func (x *ListPoliciesRequest) Reset() {
	*x = ListPoliciesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPoliciesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPoliciesRequest) ProtoMessage() {}

func (x *ListPoliciesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPoliciesRequest.ProtoReflect.Descriptor instead.
func (*ListPoliciesRequest) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{0}
}

func (x *ListPoliciesRequest) GetOffset() int64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *ListPoliciesRequest) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

// UserInfo contains user details.
type PolicyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username     string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	AuthzPolicy  string `protobuf:"bytes,2,opt,name=AuthzPolicy,proto3" json:"AuthzPolicy,omitempty"`
	PolicyShadow string `protobuf:"bytes,3,opt,name=PolicyShadow,proto3" json:"PolicyShadow,omitempty"`
}

func (x *PolicyInfo) Reset() {
	*x = PolicyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyInfo) ProtoMessage() {}

func (x *PolicyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyInfo.ProtoReflect.Descriptor instead.
func (*PolicyInfo) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{1}
}

func (x *PolicyInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PolicyInfo) GetAuthzPolicy() string {
	if x != nil {
		return x.AuthzPolicy
	}
	return ""
}

func (x *PolicyInfo) GetPolicyShadow() string {
	if x != nil {
		return x.PolicyShadow
	}
	return ""
}

// ListUsersResponse defines ListUsers response struct.
type ListPoliciesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64         `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Items      []*PolicyInfo `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListPoliciesResponse) Reset() {
	*x = ListPoliciesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPoliciesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPoliciesResponse) ProtoMessage() {}

func (x *ListPoliciesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPoliciesResponse.ProtoReflect.Descriptor instead.
func (*ListPoliciesResponse) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{2}
}

func (x *ListPoliciesResponse) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListPoliciesResponse) GetItems() []*PolicyInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_policy_proto protoreflect.FileDescriptor

var file_policy_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x62, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x6e, 0x0a, 0x0a, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x7a,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x53, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x53, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x22, 0x61, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x55, 0x0a,
	0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x4b, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x5c, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x62, 0x69, 0x72, 0x74, 0x68, 0x6d, 0x6f, 0x6e, 0x6b, 0x65, 0x79,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x63, 0x61, 0x66, 0x66, 0x6f, 0x6c, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_policy_proto_rawDescOnce sync.Once
	file_policy_proto_rawDescData = file_policy_proto_rawDesc
)

func file_policy_proto_rawDescGZIP() []byte {
	file_policy_proto_rawDescOnce.Do(func() {
		file_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_policy_proto_rawDescData)
	})
	return file_policy_proto_rawDescData
}

var file_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_policy_proto_goTypes = []interface{}{
	(*ListPoliciesRequest)(nil),  // 0: policy.ListPoliciesRequest
	(*PolicyInfo)(nil),           // 1: policy.PolicyInfo
	(*ListPoliciesResponse)(nil), // 2: policy.ListPoliciesResponse
}
var file_policy_proto_depIdxs = []int32{
	1, // 0: policy.ListPoliciesResponse.items:type_name -> policy.PolicyInfo
	0, // 1: policy.Policy.ListPolicies:input_type -> policy.ListPoliciesRequest
	2, // 2: policy.Policy.ListPolicies:output_type -> policy.ListPoliciesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_policy_proto_init() }
func file_policy_proto_init() {
	if File_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPoliciesRequest); i {
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
		file_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyInfo); i {
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
		file_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPoliciesResponse); i {
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
	file_policy_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_policy_proto_goTypes,
		DependencyIndexes: file_policy_proto_depIdxs,
		MessageInfos:      file_policy_proto_msgTypes,
	}.Build()
	File_policy_proto = out.File
	file_policy_proto_rawDesc = nil
	file_policy_proto_goTypes = nil
	file_policy_proto_depIdxs = nil
}
