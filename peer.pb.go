//
//Copyright 2018-2022 Mailgun Technologies Inc
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: peer.proto

package gubernator

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

type ForwardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Must specify at least one RateLimit. The peer that receives this request MUST be authoritative for
	// each rate_limit[x].unique_key provided, as the peer will not forward the request to any other peers
	Requests []*RateLimitRequest `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
}

func (x *ForwardRequest) Reset() {
	*x = ForwardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForwardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardRequest) ProtoMessage() {}

func (x *ForwardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardRequest.ProtoReflect.Descriptor instead.
func (*ForwardRequest) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{0}
}

func (x *ForwardRequest) GetRequests() []*RateLimitRequest {
	if x != nil {
		return x.Requests
	}
	return nil
}

type ForwardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Responses are in the same order as they appeared in the PeerRateLimitRequestuests
	RateLimits []*RateLimitResponse `protobuf:"bytes,1,rep,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
}

func (x *ForwardResponse) Reset() {
	*x = ForwardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForwardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardResponse) ProtoMessage() {}

func (x *ForwardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardResponse.ProtoReflect.Descriptor instead.
func (*ForwardResponse) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{1}
}

func (x *ForwardResponse) GetRateLimits() []*RateLimitResponse {
	if x != nil {
		return x.RateLimits
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Must specify at least one RateLimit
	Globals []*UpdateRateLimit `protobuf:"bytes,1,rep,name=globals,proto3" json:"globals,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRequest) GetGlobals() []*UpdateRateLimit {
	if x != nil {
		return x.Globals
	}
	return nil
}

type UpdateRateLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Uniquely identifies this rate limit IE: 'ip:10.2.10.7' or 'account:123445'
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The rate limit state to update
	State *RateLimitResponse `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	// The algorithm used to calculate the rate limit. The algorithm may change on
	// subsequent requests, when this occurs any previous rate limit hit counts are reset.
	Algorithm Algorithm `protobuf:"varint,3,opt,name=algorithm,proto3,enum=gubernator.v3.Algorithm" json:"algorithm,omitempty"`
	// The duration of the rate limit in milliseconds
	Duration int64 `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	// The exact time the original request was created in Epoch milliseconds.
	// Due to time drift between systems, it may be advantageous for a client to
	// set the exact time the request was created. It possible the system clock
	// for the client has drifted from the system clock where gubernator daemon
	// is running.
	//
	// The created time is used by gubernator to calculate the reset time for
	// both token and leaky algorithms. If it is not set by the client,
	// gubernator will set the created time when it receives the rate limit
	// request.
	CreatedAt int64 `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *UpdateRateLimit) Reset() {
	*x = UpdateRateLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRateLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRateLimit) ProtoMessage() {}

func (x *UpdateRateLimit) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRateLimit.ProtoReflect.Descriptor instead.
func (*UpdateRateLimit) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRateLimit) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UpdateRateLimit) GetState() *RateLimitResponse {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *UpdateRateLimit) GetAlgorithm() Algorithm {
	if x != nil {
		return x.Algorithm
	}
	return Algorithm_TOKEN_BUCKET
}

func (x *UpdateRateLimit) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *UpdateRateLimit) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

var File_peer_proto protoreflect.FileDescriptor

var file_peer_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x67, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x10, 0x67, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a,
	0x0e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3b, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76,
	0x33, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x54, 0x0a, 0x0f,
	0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x41, 0x0a, 0x0b, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x73, 0x22, 0x49, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x07, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x52, 0x07, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x73, 0x22, 0xce, 0x01,
	0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x36, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x76, 0x33, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x61,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x67, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x41,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x25,
	0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peer_proto_rawDescOnce sync.Once
	file_peer_proto_rawDescData = file_peer_proto_rawDesc
)

func file_peer_proto_rawDescGZIP() []byte {
	file_peer_proto_rawDescOnce.Do(func() {
		file_peer_proto_rawDescData = protoimpl.X.CompressGZIP(file_peer_proto_rawDescData)
	})
	return file_peer_proto_rawDescData
}

var file_peer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_peer_proto_goTypes = []interface{}{
	(*ForwardRequest)(nil),    // 0: gubernator.v3.ForwardRequest
	(*ForwardResponse)(nil),   // 1: gubernator.v3.ForwardResponse
	(*UpdateRequest)(nil),     // 2: gubernator.v3.UpdateRequest
	(*UpdateRateLimit)(nil),   // 3: gubernator.v3.UpdateRateLimit
	(*RateLimitRequest)(nil),  // 4: gubernator.v3.RateLimitRequest
	(*RateLimitResponse)(nil), // 5: gubernator.v3.RateLimitResponse
	(Algorithm)(0),            // 6: gubernator.v3.Algorithm
}
var file_peer_proto_depIdxs = []int32{
	4, // 0: gubernator.v3.ForwardRequest.requests:type_name -> gubernator.v3.RateLimitRequest
	5, // 1: gubernator.v3.ForwardResponse.rate_limits:type_name -> gubernator.v3.RateLimitResponse
	3, // 2: gubernator.v3.UpdateRequest.globals:type_name -> gubernator.v3.UpdateRateLimit
	5, // 3: gubernator.v3.UpdateRateLimit.state:type_name -> gubernator.v3.RateLimitResponse
	6, // 4: gubernator.v3.UpdateRateLimit.algorithm:type_name -> gubernator.v3.Algorithm
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_peer_proto_init() }
func file_peer_proto_init() {
	if File_peer_proto != nil {
		return
	}
	file_gubernator_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_peer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForwardRequest); i {
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
		file_peer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForwardResponse); i {
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
		file_peer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_peer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRateLimit); i {
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
			RawDescriptor: file_peer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_peer_proto_goTypes,
		DependencyIndexes: file_peer_proto_depIdxs,
		MessageInfos:      file_peer_proto_msgTypes,
	}.Build()
	File_peer_proto = out.File
	file_peer_proto_rawDesc = nil
	file_peer_proto_goTypes = nil
	file_peer_proto_depIdxs = nil
}
