// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: protos/translate/translate.proto

package translate

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

type TranslateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Word            []string `protobuf:"bytes,1,rep,name=word,proto3" json:"word,omitempty"`
	SourceLanguage  string   `protobuf:"bytes,2,opt,name=source_language,json=sourceLanguage,proto3" json:"source_language,omitempty"`
	TargetLanguages []string `protobuf:"bytes,3,rep,name=target_languages,json=targetLanguages,proto3" json:"target_languages,omitempty"`
}

func (x *TranslateRequest) Reset() {
	*x = TranslateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_translate_translate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslateRequest) ProtoMessage() {}

func (x *TranslateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_translate_translate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslateRequest.ProtoReflect.Descriptor instead.
func (*TranslateRequest) Descriptor() ([]byte, []int) {
	return file_protos_translate_translate_proto_rawDescGZIP(), []int{0}
}

func (x *TranslateRequest) GetWord() []string {
	if x != nil {
		return x.Word
	}
	return nil
}

func (x *TranslateRequest) GetSourceLanguage() string {
	if x != nil {
		return x.SourceLanguage
	}
	return ""
}

func (x *TranslateRequest) GetTargetLanguages() []string {
	if x != nil {
		return x.TargetLanguages
	}
	return nil
}

type TranslateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TranslatedWord map[string]string `protobuf:"bytes,1,rep,name=translated_word,json=translatedWord,proto3" json:"translated_word,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TranslateResponse) Reset() {
	*x = TranslateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_translate_translate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslateResponse) ProtoMessage() {}

func (x *TranslateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_translate_translate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslateResponse.ProtoReflect.Descriptor instead.
func (*TranslateResponse) Descriptor() ([]byte, []int) {
	return file_protos_translate_translate_proto_rawDescGZIP(), []int{1}
}

func (x *TranslateResponse) GetTranslatedWord() map[string]string {
	if x != nil {
		return x.TranslatedWord
	}
	return nil
}

var File_protos_translate_translate_proto protoreflect.FileDescriptor

var file_protos_translate_translate_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x7a, 0x0a,
	0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x22, 0xb1, 0x01, 0x0a, 0x11, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x59, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x57, 0x6f, 0x72, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x64, 0x1a, 0x41, 0x0a, 0x13, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x5d, 0x0a,
	0x0a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x72, 0x12, 0x4f, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x64, 0x12,
	0x1b, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_translate_translate_proto_rawDescOnce sync.Once
	file_protos_translate_translate_proto_rawDescData = file_protos_translate_translate_proto_rawDesc
)

func file_protos_translate_translate_proto_rawDescGZIP() []byte {
	file_protos_translate_translate_proto_rawDescOnce.Do(func() {
		file_protos_translate_translate_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_translate_translate_proto_rawDescData)
	})
	return file_protos_translate_translate_proto_rawDescData
}

var file_protos_translate_translate_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_translate_translate_proto_goTypes = []interface{}{
	(*TranslateRequest)(nil),  // 0: translate.TranslateRequest
	(*TranslateResponse)(nil), // 1: translate.TranslateResponse
	nil,                       // 2: translate.TranslateResponse.TranslatedWordEntry
}
var file_protos_translate_translate_proto_depIdxs = []int32{
	2, // 0: translate.TranslateResponse.translated_word:type_name -> translate.TranslateResponse.TranslatedWordEntry
	0, // 1: translate.Translater.GetTranslateWord:input_type -> translate.TranslateRequest
	1, // 2: translate.Translater.GetTranslateWord:output_type -> translate.TranslateResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_translate_translate_proto_init() }
func file_protos_translate_translate_proto_init() {
	if File_protos_translate_translate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_translate_translate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslateRequest); i {
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
		file_protos_translate_translate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslateResponse); i {
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
			RawDescriptor: file_protos_translate_translate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_translate_translate_proto_goTypes,
		DependencyIndexes: file_protos_translate_translate_proto_depIdxs,
		MessageInfos:      file_protos_translate_translate_proto_msgTypes,
	}.Build()
	File_protos_translate_translate_proto = out.File
	file_protos_translate_translate_proto_rawDesc = nil
	file_protos_translate_translate_proto_goTypes = nil
	file_protos_translate_translate_proto_depIdxs = nil
}
