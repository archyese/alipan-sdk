// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: protos/adrive_client/adrive_open_file_update_params.proto

package protos

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

type AdriveOpenFileUpdateParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriveId       string `protobuf:"bytes,1,opt,name=drive_id,json=driveId,proto3" json:"drive_id,omitempty"`
	FileId        string `protobuf:"bytes,2,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Name          string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CheckNameMode string `protobuf:"bytes,4,opt,name=check_name_mode,json=checkNameMode,proto3" json:"check_name_mode,omitempty"`
	Starred       bool   `protobuf:"varint,5,opt,name=starred,proto3" json:"starred,omitempty"`
}

func (x *AdriveOpenFileUpdateParams) Reset() {
	*x = AdriveOpenFileUpdateParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_adrive_client_adrive_open_file_update_params_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdriveOpenFileUpdateParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdriveOpenFileUpdateParams) ProtoMessage() {}

func (x *AdriveOpenFileUpdateParams) ProtoReflect() protoreflect.Message {
	mi := &file_protos_adrive_client_adrive_open_file_update_params_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdriveOpenFileUpdateParams.ProtoReflect.Descriptor instead.
func (*AdriveOpenFileUpdateParams) Descriptor() ([]byte, []int) {
	return file_protos_adrive_client_adrive_open_file_update_params_proto_rawDescGZIP(), []int{0}
}

func (x *AdriveOpenFileUpdateParams) GetDriveId() string {
	if x != nil {
		return x.DriveId
	}
	return ""
}

func (x *AdriveOpenFileUpdateParams) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *AdriveOpenFileUpdateParams) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdriveOpenFileUpdateParams) GetCheckNameMode() string {
	if x != nil {
		return x.CheckNameMode
	}
	return ""
}

func (x *AdriveOpenFileUpdateParams) GetStarred() bool {
	if x != nil {
		return x.Starred
	}
	return false
}

var File_protos_adrive_client_adrive_open_file_update_params_proto protoreflect.FileDescriptor

var file_protos_adrive_client_adrive_open_file_update_params_proto_rawDesc = []byte{
	0x0a, 0x39, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x64, 0x72, 0x69, 0x76, 0x65, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x64, 0x72, 0x69, 0x76, 0x65, 0x5f, 0x6f, 0x70,
	0x65, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x1a,
	0x41, 0x64, 0x72, 0x69, 0x76, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74,
	0x61, 0x72, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x74, 0x61,
	0x72, 0x72, 0x65, 0x64, 0x42, 0x17, 0x5a, 0x15, 0x61, 0x64, 0x72, 0x69, 0x76, 0x65, 0x5f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_adrive_client_adrive_open_file_update_params_proto_rawDescOnce sync.Once
	file_protos_adrive_client_adrive_open_file_update_params_proto_rawDescData = file_protos_adrive_client_adrive_open_file_update_params_proto_rawDesc
)

func file_protos_adrive_client_adrive_open_file_update_params_proto_rawDescGZIP() []byte {
	file_protos_adrive_client_adrive_open_file_update_params_proto_rawDescOnce.Do(func() {
		file_protos_adrive_client_adrive_open_file_update_params_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_adrive_client_adrive_open_file_update_params_proto_rawDescData)
	})
	return file_protos_adrive_client_adrive_open_file_update_params_proto_rawDescData
}

var file_protos_adrive_client_adrive_open_file_update_params_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_adrive_client_adrive_open_file_update_params_proto_goTypes = []interface{}{
	(*AdriveOpenFileUpdateParams)(nil), // 0: AdriveOpenFileUpdateParams
}
var file_protos_adrive_client_adrive_open_file_update_params_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_adrive_client_adrive_open_file_update_params_proto_init() }
func file_protos_adrive_client_adrive_open_file_update_params_proto_init() {
	if File_protos_adrive_client_adrive_open_file_update_params_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_adrive_client_adrive_open_file_update_params_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdriveOpenFileUpdateParams); i {
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
			RawDescriptor: file_protos_adrive_client_adrive_open_file_update_params_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_adrive_client_adrive_open_file_update_params_proto_goTypes,
		DependencyIndexes: file_protos_adrive_client_adrive_open_file_update_params_proto_depIdxs,
		MessageInfos:      file_protos_adrive_client_adrive_open_file_update_params_proto_msgTypes,
	}.Build()
	File_protos_adrive_client_adrive_open_file_update_params_proto = out.File
	file_protos_adrive_client_adrive_open_file_update_params_proto_rawDesc = nil
	file_protos_adrive_client_adrive_open_file_update_params_proto_goTypes = nil
	file_protos_adrive_client_adrive_open_file_update_params_proto_depIdxs = nil
}