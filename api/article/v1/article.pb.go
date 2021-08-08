// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.5.1
// source: api/article/v1/article.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The request message containing the user's name.
type ArticlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Num  int64 `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *ArticlesRequest) Reset() {
	*x = ArticlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_v1_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticlesRequest) ProtoMessage() {}

func (x *ArticlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_v1_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticlesRequest.ProtoReflect.Descriptor instead.
func (*ArticlesRequest) Descriptor() ([]byte, []int) {
	return file_api_article_v1_article_proto_rawDescGZIP(), []int{0}
}

func (x *ArticlesRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ArticlesRequest) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

// The response message containing the greetings
type ArticlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*ArticlesResponse_Article `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ArticlesResponse) Reset() {
	*x = ArticlesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_v1_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticlesResponse) ProtoMessage() {}

func (x *ArticlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_v1_article_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticlesResponse.ProtoReflect.Descriptor instead.
func (*ArticlesResponse) Descriptor() ([]byte, []int) {
	return file_api_article_v1_article_proto_rawDescGZIP(), []int{1}
}

func (x *ArticlesResponse) GetResults() []*ArticlesResponse_Article {
	if x != nil {
		return x.Results
	}
	return nil
}

type ArticlesResponse_Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ArticlesResponse_Article) Reset() {
	*x = ArticlesResponse_Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_article_v1_article_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticlesResponse_Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticlesResponse_Article) ProtoMessage() {}

func (x *ArticlesResponse_Article) ProtoReflect() protoreflect.Message {
	mi := &file_api_article_v1_article_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticlesResponse_Article.ProtoReflect.Descriptor instead.
func (*ArticlesResponse_Article) Descriptor() ([]byte, []int) {
	return file_api_article_v1_article_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ArticlesResponse_Article) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ArticlesResponse_Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ArticlesResponse_Article) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_api_article_v1_article_proto protoreflect.FileDescriptor

var file_api_article_v1_article_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x0f,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0xa1, 0x01, 0x0a, 0x10, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x49,
	0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0xdc, 0x01, 0x0a, 0x07, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x67, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e,
	0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x68,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2f, 0x3a, 0x69, 0x64, 0x42, 0x10, 0x5a, 0x0e, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_article_v1_article_proto_rawDescOnce sync.Once
	file_api_article_v1_article_proto_rawDescData = file_api_article_v1_article_proto_rawDesc
)

func file_api_article_v1_article_proto_rawDescGZIP() []byte {
	file_api_article_v1_article_proto_rawDescOnce.Do(func() {
		file_api_article_v1_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_article_v1_article_proto_rawDescData)
	})
	return file_api_article_v1_article_proto_rawDescData
}

var file_api_article_v1_article_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_article_v1_article_proto_goTypes = []interface{}{
	(*ArticlesRequest)(nil),          // 0: api.article.v1.ArticlesRequest
	(*ArticlesResponse)(nil),         // 1: api.article.v1.ArticlesResponse
	(*ArticlesResponse_Article)(nil), // 2: api.article.v1.ArticlesResponse.Article
}
var file_api_article_v1_article_proto_depIdxs = []int32{
	2, // 0: api.article.v1.ArticlesResponse.results:type_name -> api.article.v1.ArticlesResponse.Article
	0, // 1: api.article.v1.Article.ListArticles:input_type -> api.article.v1.ArticlesRequest
	0, // 2: api.article.v1.Article.GetArticle:input_type -> api.article.v1.ArticlesRequest
	1, // 3: api.article.v1.Article.ListArticles:output_type -> api.article.v1.ArticlesResponse
	1, // 4: api.article.v1.Article.GetArticle:output_type -> api.article.v1.ArticlesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_article_v1_article_proto_init() }
func file_api_article_v1_article_proto_init() {
	if File_api_article_v1_article_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_article_v1_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticlesRequest); i {
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
		file_api_article_v1_article_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticlesResponse); i {
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
		file_api_article_v1_article_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticlesResponse_Article); i {
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
			RawDescriptor: file_api_article_v1_article_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_article_v1_article_proto_goTypes,
		DependencyIndexes: file_api_article_v1_article_proto_depIdxs,
		MessageInfos:      file_api_article_v1_article_proto_msgTypes,
	}.Build()
	File_api_article_v1_article_proto = out.File
	file_api_article_v1_article_proto_rawDesc = nil
	file_api_article_v1_article_proto_goTypes = nil
	file_api_article_v1_article_proto_depIdxs = nil
}
