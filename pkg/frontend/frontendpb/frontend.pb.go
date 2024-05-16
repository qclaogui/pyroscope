// SPDX-License-Identifier: AGPL-3.0-only
// Provenance-includes-location: https://github.com/cortexproject/cortex/blob/master/pkg/frontend/v2/frontendv2pb/frontend.proto
// Provenance-includes-license: Apache-2.0
// Provenance-includes-copyright: The Cortex Authors.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: frontend/frontendpb/frontend.proto

package frontendpb

import (
	stats "github.com/grafana/pyroscope/pkg/querier/stats"
	httpgrpc "github.com/grafana/pyroscope/pkg/util/httpgrpc"
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

type QueryResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryID      uint64                 `protobuf:"varint,1,opt,name=queryID,proto3" json:"queryID,omitempty"`
	HttpResponse *httpgrpc.HTTPResponse `protobuf:"bytes,2,opt,name=httpResponse,proto3" json:"httpResponse,omitempty"`
	Stats        *stats.Stats           `protobuf:"bytes,3,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *QueryResultRequest) Reset() {
	*x = QueryResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frontend_frontendpb_frontend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResultRequest) ProtoMessage() {}

func (x *QueryResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_frontend_frontendpb_frontend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResultRequest.ProtoReflect.Descriptor instead.
func (*QueryResultRequest) Descriptor() ([]byte, []int) {
	return file_frontend_frontendpb_frontend_proto_rawDescGZIP(), []int{0}
}

func (x *QueryResultRequest) GetQueryID() uint64 {
	if x != nil {
		return x.QueryID
	}
	return 0
}

func (x *QueryResultRequest) GetHttpResponse() *httpgrpc.HTTPResponse {
	if x != nil {
		return x.HttpResponse
	}
	return nil
}

func (x *QueryResultRequest) GetStats() *stats.Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

type QueryResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryResultResponse) Reset() {
	*x = QueryResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_frontend_frontendpb_frontend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResultResponse) ProtoMessage() {}

func (x *QueryResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_frontend_frontendpb_frontend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResultResponse.ProtoReflect.Descriptor instead.
func (*QueryResultResponse) Descriptor() ([]byte, []int) {
	return file_frontend_frontendpb_frontend_proto_rawDescGZIP(), []int{1}
}

var File_frontend_frontendpb_frontend_proto protoreflect.FileDescriptor

var file_frontend_frontendpb_frontend_proto_rawDesc = []byte{
	0x0a, 0x22, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x70, 0x62, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x70, 0x62,
	0x1a, 0x19, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x75, 0x74, 0x69,
	0x6c, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x12, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x44, 0x12, 0x3a, 0x0a, 0x0c, 0x68, 0x74,
	0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x54, 0x54, 0x50,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x66, 0x0a, 0x12, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x46, 0x6f, 0x72,
	0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x12, 0x50, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x70, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x70, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x9d, 0x01, 0x0a, 0x0e, 0x63, 0x6f,
	0x6d, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x42, 0x0d, 0x46, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x61, 0x6e,
	0x61, 0x2f, 0x70, 0x79, 0x72, 0x6f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x46, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x46, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x70, 0x62, 0xca, 0x02, 0x0a, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x70, 0x62, 0xe2, 0x02, 0x16, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x70, 0x62,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x46,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_frontend_frontendpb_frontend_proto_rawDescOnce sync.Once
	file_frontend_frontendpb_frontend_proto_rawDescData = file_frontend_frontendpb_frontend_proto_rawDesc
)

func file_frontend_frontendpb_frontend_proto_rawDescGZIP() []byte {
	file_frontend_frontendpb_frontend_proto_rawDescOnce.Do(func() {
		file_frontend_frontendpb_frontend_proto_rawDescData = protoimpl.X.CompressGZIP(file_frontend_frontendpb_frontend_proto_rawDescData)
	})
	return file_frontend_frontendpb_frontend_proto_rawDescData
}

var file_frontend_frontendpb_frontend_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_frontend_frontendpb_frontend_proto_goTypes = []interface{}{
	(*QueryResultRequest)(nil),    // 0: frontendpb.QueryResultRequest
	(*QueryResultResponse)(nil),   // 1: frontendpb.QueryResultResponse
	(*httpgrpc.HTTPResponse)(nil), // 2: httpgrpc.HTTPResponse
	(*stats.Stats)(nil),           // 3: stats.Stats
}
var file_frontend_frontendpb_frontend_proto_depIdxs = []int32{
	2, // 0: frontendpb.QueryResultRequest.httpResponse:type_name -> httpgrpc.HTTPResponse
	3, // 1: frontendpb.QueryResultRequest.stats:type_name -> stats.Stats
	0, // 2: frontendpb.FrontendForQuerier.QueryResult:input_type -> frontendpb.QueryResultRequest
	1, // 3: frontendpb.FrontendForQuerier.QueryResult:output_type -> frontendpb.QueryResultResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_frontend_frontendpb_frontend_proto_init() }
func file_frontend_frontendpb_frontend_proto_init() {
	if File_frontend_frontendpb_frontend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_frontend_frontendpb_frontend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResultRequest); i {
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
		file_frontend_frontendpb_frontend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResultResponse); i {
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
			RawDescriptor: file_frontend_frontendpb_frontend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_frontend_frontendpb_frontend_proto_goTypes,
		DependencyIndexes: file_frontend_frontendpb_frontend_proto_depIdxs,
		MessageInfos:      file_frontend_frontendpb_frontend_proto_msgTypes,
	}.Build()
	File_frontend_frontendpb_frontend_proto = out.File
	file_frontend_frontendpb_frontend_proto_rawDesc = nil
	file_frontend_frontendpb_frontend_proto_goTypes = nil
	file_frontend_frontendpb_frontend_proto_depIdxs = nil
}
