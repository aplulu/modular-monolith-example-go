// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: example/article/v1/common.proto

package article

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

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string       `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string       `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	User    *ArticleUser `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_article_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_example_article_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_example_article_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *Article) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Article) GetUser() *ArticleUser {
	if x != nil {
		return x.User
	}
	return nil
}

type ArticleUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ArticleUser) Reset() {
	*x = ArticleUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_article_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleUser) ProtoMessage() {}

func (x *ArticleUser) ProtoReflect() protoreflect.Message {
	mi := &file_example_article_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleUser.ProtoReflect.Descriptor instead.
func (*ArticleUser) Descriptor() ([]byte, []int) {
	return file_example_article_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *ArticleUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ArticleUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_example_article_v1_common_proto protoreflect.FileDescriptor

var file_example_article_v1_common_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x7e, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x33, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x31, 0x0a, 0x0b, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x58, 0x5a, 0x56, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x6c, 0x75, 0x6c, 0x75, 0x2f, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x61, 0x72, 0x2d, 0x6d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68, 0x2d, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_article_v1_common_proto_rawDescOnce sync.Once
	file_example_article_v1_common_proto_rawDescData = file_example_article_v1_common_proto_rawDesc
)

func file_example_article_v1_common_proto_rawDescGZIP() []byte {
	file_example_article_v1_common_proto_rawDescOnce.Do(func() {
		file_example_article_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_article_v1_common_proto_rawDescData)
	})
	return file_example_article_v1_common_proto_rawDescData
}

var file_example_article_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_example_article_v1_common_proto_goTypes = []interface{}{
	(*Article)(nil),     // 0: example.article.v1.Article
	(*ArticleUser)(nil), // 1: example.article.v1.ArticleUser
}
var file_example_article_v1_common_proto_depIdxs = []int32{
	1, // 0: example.article.v1.Article.user:type_name -> example.article.v1.ArticleUser
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_example_article_v1_common_proto_init() }
func file_example_article_v1_common_proto_init() {
	if File_example_article_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_article_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
		file_example_article_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleUser); i {
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
			RawDescriptor: file_example_article_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_article_v1_common_proto_goTypes,
		DependencyIndexes: file_example_article_v1_common_proto_depIdxs,
		MessageInfos:      file_example_article_v1_common_proto_msgTypes,
	}.Build()
	File_example_article_v1_common_proto = out.File
	file_example_article_v1_common_proto_rawDesc = nil
	file_example_article_v1_common_proto_goTypes = nil
	file_example_article_v1_common_proto_depIdxs = nil
}
