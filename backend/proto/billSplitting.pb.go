// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/billSplitting.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_billSplitting_proto protoreflect.FileDescriptor

var file_proto_billSplitting_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x53, 0x70, 0x6c, 0x69,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe9, 0x02, 0x0a,
	0x0d, 0x42, 0x69, 0x6c, 0x6c, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x6e,
	0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x27, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73,
	0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x41, 0x64, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x64, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x2d, 0x73, 0x70, 0x6c, 0x69,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_proto_billSplitting_proto_goTypes = []interface{}{
	(*CreateExpenseSummaryChartRequest)(nil),  // 0: proto.CreateExpenseSummaryChartRequest
	(*GetAuthTokenRequest)(nil),               // 1: proto.GetAuthTokenRequest
	(*CreateLineGroupRequest)(nil),            // 2: proto.CreateLineGroupRequest
	(*AddGroupMemberRequest)(nil),             // 3: proto.AddGroupMemberRequest
	(*CreateExpenseSummaryChartResponse)(nil), // 4: proto.CreateExpenseSummaryChartResponse
	(*GetAuthTokenResponse)(nil),              // 5: proto.GetAuthTokenResponse
	(*CreateLineGroupResponse)(nil),           // 6: proto.CreateLineGroupResponse
	(*AddGroupMemberResponse)(nil),            // 7: proto.AddGroupMemberResponse
}
var file_proto_billSplitting_proto_depIdxs = []int32{
	0, // 0: proto.BillSplitting.CreateExpenseSummaryChart:input_type -> proto.CreateExpenseSummaryChartRequest
	1, // 1: proto.BillSplitting.GetAuthToken:input_type -> proto.GetAuthTokenRequest
	2, // 2: proto.BillSplitting.CreateLineGroup:input_type -> proto.CreateLineGroupRequest
	3, // 3: proto.BillSplitting.AddGroupMember:input_type -> proto.AddGroupMemberRequest
	4, // 4: proto.BillSplitting.CreateExpenseSummaryChart:output_type -> proto.CreateExpenseSummaryChartResponse
	5, // 5: proto.BillSplitting.GetAuthToken:output_type -> proto.GetAuthTokenResponse
	6, // 6: proto.BillSplitting.CreateLineGroup:output_type -> proto.CreateLineGroupResponse
	7, // 7: proto.BillSplitting.AddGroupMember:output_type -> proto.AddGroupMemberResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_billSplitting_proto_init() }
func file_proto_billSplitting_proto_init() {
	if File_proto_billSplitting_proto != nil {
		return
	}
	file_proto_expense_proto_init()
	file_proto_auth_proto_init()
	file_proto_group_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_billSplitting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_billSplitting_proto_goTypes,
		DependencyIndexes: file_proto_billSplitting_proto_depIdxs,
	}.Build()
	File_proto_billSplitting_proto = out.File
	file_proto_billSplitting_proto_rawDesc = nil
	file_proto_billSplitting_proto_goTypes = nil
	file_proto_billSplitting_proto_depIdxs = nil
}
