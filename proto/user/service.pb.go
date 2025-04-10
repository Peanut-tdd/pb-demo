// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: user/service.proto

package user

import (
	article "github.com/pb-demo/proto/article"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_user_service_proto protoreflect.FileDescriptor

const file_user_service_proto_rawDesc = "" +
	"\n" +
	"\x12user/service.proto\x12\x04user\x1a\x12user/message.proto\x1a\x15article/article.proto2i\n" +
	"\x04User\x12+\n" +
	"\vGetUserInfo\x12\f.user.UserID\x1a\x0e.user.UserInfo\x124\n" +
	"\x11GetUserFavArticle\x12\f.user.UserID\x1a\x11.article.ArticlesB\x1fZ\x1dgithub.com/pb-demo/proto/userb\x06proto3"

var file_user_service_proto_goTypes = []any{
	(*UserID)(nil),           // 0: user.UserID
	(*UserInfo)(nil),         // 1: user.UserInfo
	(*article.Articles)(nil), // 2: article.Articles
}
var file_user_service_proto_depIdxs = []int32{
	0, // 0: user.User.GetUserInfo:input_type -> user.UserID
	0, // 1: user.User.GetUserFavArticle:input_type -> user.UserID
	1, // 2: user.User.GetUserInfo:output_type -> user.UserInfo
	2, // 3: user.User.GetUserFavArticle:output_type -> article.Articles
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_service_proto_init() }
func file_user_service_proto_init() {
	if File_user_service_proto != nil {
		return
	}
	file_user_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_service_proto_rawDesc), len(file_user_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_service_proto_goTypes,
		DependencyIndexes: file_user_service_proto_depIdxs,
	}.Build()
	File_user_service_proto = out.File
	file_user_service_proto_goTypes = nil
	file_user_service_proto_depIdxs = nil
}
