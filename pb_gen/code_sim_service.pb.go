// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.5.0
// source: code_sim/remote_code_service.proto

package pb_gen

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

type CodeSimUploadFileType int32

const (
	CodeSimUploadFileType_code_sim_upload_file_type_unknown CodeSimUploadFileType = 0
	CodeSimUploadFileType_code_sim_upload_file_type_zip     CodeSimUploadFileType = 1
)

// Enum value maps for CodeSimUploadFileType.
var (
	CodeSimUploadFileType_name = map[int32]string{
		0: "code_sim_upload_file_type_unknown",
		1: "code_sim_upload_file_type_zip",
	}
	CodeSimUploadFileType_value = map[string]int32{
		"code_sim_upload_file_type_unknown": 0,
		"code_sim_upload_file_type_zip":     1,
	}
)

func (x CodeSimUploadFileType) Enum() *CodeSimUploadFileType {
	p := new(CodeSimUploadFileType)
	*p = x
	return p
}

func (x CodeSimUploadFileType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CodeSimUploadFileType) Descriptor() protoreflect.EnumDescriptor {
	return file_code_sim_code_sim_service_proto_enumTypes[0].Descriptor()
}

func (CodeSimUploadFileType) Type() protoreflect.EnumType {
	return &file_code_sim_code_sim_service_proto_enumTypes[0]
}

func (x CodeSimUploadFileType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CodeSimUploadFileType.Descriptor instead.
func (CodeSimUploadFileType) EnumDescriptor() ([]byte, []int) {
	return file_code_sim_code_sim_service_proto_rawDescGZIP(), []int{0}
}

type CodeSimUploadStatus int32

const (
	CodeSimUploadStatus_code_sim_upload_status_unknown CodeSimUploadStatus = 0
	CodeSimUploadStatus_code_sim_upload_status_OK      CodeSimUploadStatus = 1
	CodeSimUploadStatus_code_sim_upload_status_fail    CodeSimUploadStatus = 2
)

// Enum value maps for CodeSimUploadStatus.
var (
	CodeSimUploadStatus_name = map[int32]string{
		0: "code_sim_upload_status_unknown",
		1: "code_sim_upload_status_OK",
		2: "code_sim_upload_status_fail",
	}
	CodeSimUploadStatus_value = map[string]int32{
		"code_sim_upload_status_unknown": 0,
		"code_sim_upload_status_OK":      1,
		"code_sim_upload_status_fail":    2,
	}
)

func (x CodeSimUploadStatus) Enum() *CodeSimUploadStatus {
	p := new(CodeSimUploadStatus)
	*p = x
	return p
}

func (x CodeSimUploadStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CodeSimUploadStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_code_sim_code_sim_service_proto_enumTypes[1].Descriptor()
}

func (CodeSimUploadStatus) Type() protoreflect.EnumType {
	return &file_code_sim_code_sim_service_proto_enumTypes[1]
}

func (x CodeSimUploadStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CodeSimUploadStatus.Descriptor instead.
func (CodeSimUploadStatus) EnumDescriptor() ([]byte, []int) {
	return file_code_sim_code_sim_service_proto_rawDescGZIP(), []int{1}
}

type CodeSimSearchRequest_CodeType int32

const (
	CodeSimSearchRequest_unknown CodeSimSearchRequest_CodeType = 0
	CodeSimSearchRequest_python  CodeSimSearchRequest_CodeType = 1
	CodeSimSearchRequest_golang  CodeSimSearchRequest_CodeType = 2
)

// Enum value maps for CodeSimSearchRequest_CodeType.
var (
	CodeSimSearchRequest_CodeType_name = map[int32]string{
		0: "unknown",
		1: "python",
		2: "golang",
	}
	CodeSimSearchRequest_CodeType_value = map[string]int32{
		"unknown": 0,
		"python":  1,
		"golang":  2,
	}
)

func (x CodeSimSearchRequest_CodeType) Enum() *CodeSimSearchRequest_CodeType {
	p := new(CodeSimSearchRequest_CodeType)
	*p = x
	return p
}

func (x CodeSimSearchRequest_CodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CodeSimSearchRequest_CodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_code_sim_code_sim_service_proto_enumTypes[2].Descriptor()
}

func (CodeSimSearchRequest_CodeType) Type() protoreflect.EnumType {
	return &file_code_sim_code_sim_service_proto_enumTypes[2]
}

func (x CodeSimSearchRequest_CodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CodeSimSearchRequest_CodeType.Descriptor instead.
func (CodeSimSearchRequest_CodeType) EnumDescriptor() ([]byte, []int) {
	return file_code_sim_code_sim_service_proto_rawDescGZIP(), []int{4, 0}
}

type CodeSimHelloWorldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HelloText string `protobuf:"bytes,1,opt,name=hello_text,json=helloText,proto3" json:"hello_text,omitempty"`
}

func (x *CodeSimHelloWorldRequest) Reset() {
	*x = CodeSimHelloWorldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_sim_code_sim_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeSimHelloWorldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeSimHelloWorldRequest) ProtoMessage() {}

func (x *CodeSimHelloWorldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_code_sim_code_sim_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeSimHelloWorldRequest.ProtoReflect.Descriptor instead.
func (*CodeSimHelloWorldRequest) Descriptor() ([]byte, []int) {
	return file_code_sim_code_sim_service_proto_rawDescGZIP(), []int{0}
}

func (x *CodeSimHelloWorldRequest) GetHelloText() string {
	if x != nil {
		return x.HelloText
	}
	return ""
}

type CodeSimProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	Tag         string `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *CodeSimProject) Reset() {
	*x = CodeSimProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_sim_code_sim_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeSimProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeSimProject) ProtoMessage() {}

func (x *CodeSimProject) ProtoReflect() protoreflect.Message {
	mi := &file_code_sim_code_sim_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeSimProject.ProtoReflect.Descriptor instead.
func (*CodeSimProject) Descriptor() ([]byte, []int) {
	return file_code_sim_code_sim_service_proto_rawDescGZIP(), []int{1}
}

func (x *CodeSimProject) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *CodeSimProject) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type CodeSimProjectFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectInfo  *CodeSimProject `protobuf:"bytes,1,opt,name=project_info,json=projectInfo,proto3" json:"project_info,omitempty"`
	RelativePath string          `protobuf:"bytes,2,opt,name=relative_path,json=relativePath,proto3" json:"relative_path,omitempty"`
	Content      []byte          `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CodeSimProjectFile) Reset() {
	*x = CodeSimProjectFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_sim_code_sim_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeSimProjectFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeSimProjectFile) ProtoMessage() {}

func (x *CodeSimProjectFile) ProtoReflect() protoreflect.Message {
	mi := &file_code_sim_code_sim_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeSimProjectFile.ProtoReflect.Descriptor instead.
func (*CodeSimProjectFile) Descriptor() ([]byte, []int) {
	return file_code_sim_code_sim_service_proto_rawDescGZIP(), []int{2}
}

func (x *CodeSimProjectFile) GetProjectInfo() *CodeSimProject {
	if x != nil {
		return x.ProjectInfo
	}
	return nil
}

func (x *CodeSimProjectFile) GetRelativePath() string {
	if x != nil {
		return x.RelativePath
	}
	return ""
}

func (x *CodeSimProjectFile) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type CodeSimHelloWorldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThanksText string `protobuf:"bytes,1,opt,name=thanks_text,json=thanksText,proto3" json:"thanks_text,omitempty"`
}

func (x *CodeSimHelloWorldResponse) Reset() {
	*x = CodeSimHelloWorldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_sim_code_sim_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeSimHelloWorldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeSimHelloWorldResponse) ProtoMessage() {}

func (x *CodeSimHelloWorldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_code_sim_code_sim_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeSimHelloWorldResponse.ProtoReflect.Descriptor instead.
func (*CodeSimHelloWorldResponse) Descriptor() ([]byte, []int) {
	return file_code_sim_code_sim_service_proto_rawDescGZIP(), []int{3}
}

func (x *CodeSimHelloWorldResponse) GetThanksText() string {
	if x != nil {
		return x.ThanksText
	}
	return ""
}

type CodeSimSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchText  string                          `protobuf:"bytes,1,opt,name=match_text,json=matchText,proto3" json:"match_text,omitempty"`
	CodeTypes  []CodeSimSearchRequest_CodeType `protobuf:"varint,2,rep,packed,name=code_types,json=codeTypes,proto3,enum=pb.CodeSimSearchRequest_CodeType" json:"code_types,omitempty"` // code_type 为unknown时，按照plain-text查询
	Limit      int32                           `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset     int32                           `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	WithSource bool                            `protobuf:"varint,5,opt,name=with_source,json=withSource,proto3" json:"with_source,omitempty"`
}

func (x *CodeSimSearchRequest) Reset() {
	*x = CodeSimSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_sim_code_sim_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeSimSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeSimSearchRequest) ProtoMessage() {}

func (x *CodeSimSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_code_sim_code_sim_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeSimSearchRequest.ProtoReflect.Descriptor instead.
func (*CodeSimSearchRequest) Descriptor() ([]byte, []int) {
	return file_code_sim_code_sim_service_proto_rawDescGZIP(), []int{4}
}

func (x *CodeSimSearchRequest) GetMatchText() string {
	if x != nil {
		return x.MatchText
	}
	return ""
}

func (x *CodeSimSearchRequest) GetCodeTypes() []CodeSimSearchRequest_CodeType {
	if x != nil {
		return x.CodeTypes
	}
	return nil
}

func (x *CodeSimSearchRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *CodeSimSearchRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *CodeSimSearchRequest) GetWithSource() bool {
	if x != nil {
		return x.WithSource
	}
	return false
}

type CodeSimSearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*CodeSimProjectFile `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *CodeSimSearchResponse) Reset() {
	*x = CodeSimSearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_sim_code_sim_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeSimSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeSimSearchResponse) ProtoMessage() {}

func (x *CodeSimSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_code_sim_code_sim_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeSimSearchResponse.ProtoReflect.Descriptor instead.
func (*CodeSimSearchResponse) Descriptor() ([]byte, []int) {
	return file_code_sim_code_sim_service_proto_rawDescGZIP(), []int{5}
}

func (x *CodeSimSearchResponse) GetFiles() []*CodeSimProjectFile {
	if x != nil {
		return x.Files
	}
	return nil
}

type CodeSimUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project      *CodeSimProject       `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	FileType     CodeSimUploadFileType `protobuf:"varint,2,opt,name=file_type,json=fileType,proto3,enum=pb.CodeSimUploadFileType" json:"file_type,omitempty"`
	TotalSize    string                `protobuf:"bytes,3,opt,name=totalSize,proto3" json:"totalSize,omitempty"`
	Received     string                `protobuf:"bytes,4,opt,name=received,proto3" json:"received,omitempty"`
	ContentChunk []byte                `protobuf:"bytes,5,opt,name=content_chunk,json=contentChunk,proto3" json:"content_chunk,omitempty"`
}

func (x *CodeSimUploadRequest) Reset() {
	*x = CodeSimUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_sim_code_sim_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeSimUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeSimUploadRequest) ProtoMessage() {}

func (x *CodeSimUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_code_sim_code_sim_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeSimUploadRequest.ProtoReflect.Descriptor instead.
func (*CodeSimUploadRequest) Descriptor() ([]byte, []int) {
	return file_code_sim_code_sim_service_proto_rawDescGZIP(), []int{6}
}

func (x *CodeSimUploadRequest) GetProject() *CodeSimProject {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *CodeSimUploadRequest) GetFileType() CodeSimUploadFileType {
	if x != nil {
		return x.FileType
	}
	return CodeSimUploadFileType_code_sim_upload_file_type_unknown
}

func (x *CodeSimUploadRequest) GetTotalSize() string {
	if x != nil {
		return x.TotalSize
	}
	return ""
}

func (x *CodeSimUploadRequest) GetReceived() string {
	if x != nil {
		return x.Received
	}
	return ""
}

func (x *CodeSimUploadRequest) GetContentChunk() []byte {
	if x != nil {
		return x.ContentChunk
	}
	return nil
}

type CodeSimUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string              `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  CodeSimUploadStatus `protobuf:"varint,2,opt,name=status,proto3,enum=pb.CodeSimUploadStatus" json:"status,omitempty"`
}

func (x *CodeSimUploadResponse) Reset() {
	*x = CodeSimUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_sim_code_sim_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeSimUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeSimUploadResponse) ProtoMessage() {}

func (x *CodeSimUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_code_sim_code_sim_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeSimUploadResponse.ProtoReflect.Descriptor instead.
func (*CodeSimUploadResponse) Descriptor() ([]byte, []int) {
	return file_code_sim_code_sim_service_proto_rawDescGZIP(), []int{7}
}

func (x *CodeSimUploadResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CodeSimUploadResponse) GetStatus() CodeSimUploadStatus {
	if x != nil {
		return x.Status
	}
	return CodeSimUploadStatus_code_sim_upload_status_unknown
}

type CodeSimDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	Tag         string `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *CodeSimDeleteRequest) Reset() {
	*x = CodeSimDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_sim_code_sim_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeSimDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeSimDeleteRequest) ProtoMessage() {}

func (x *CodeSimDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_code_sim_code_sim_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeSimDeleteRequest.ProtoReflect.Descriptor instead.
func (*CodeSimDeleteRequest) Descriptor() ([]byte, []int) {
	return file_code_sim_code_sim_service_proto_rawDescGZIP(), []int{8}
}

func (x *CodeSimDeleteRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *CodeSimDeleteRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type CodeSimDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CodeSimDeleteResponse) Reset() {
	*x = CodeSimDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_sim_code_sim_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeSimDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeSimDeleteResponse) ProtoMessage() {}

func (x *CodeSimDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_code_sim_code_sim_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeSimDeleteResponse.ProtoReflect.Descriptor instead.
func (*CodeSimDeleteResponse) Descriptor() ([]byte, []int) {
	return file_code_sim_code_sim_service_proto_rawDescGZIP(), []int{9}
}

var File_code_sim_code_sim_service_proto protoreflect.FileDescriptor

var file_code_sim_code_sim_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x5f,
	0x73, 0x69, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x39, 0x0a, 0x18, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x54, 0x65, 0x78, 0x74,
	0x22, 0x45, 0x0a, 0x0e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x8a, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x64, 0x65,
	0x53, 0x69, 0x6d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x35,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69,
	0x6d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x22, 0x3c, 0x0a, 0x19, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x68, 0x61, 0x6e, 0x6b, 0x73, 0x5f, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x68, 0x61, 0x6e, 0x6b, 0x73, 0x54, 0x65,
	0x78, 0x74, 0x22, 0xf7, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x65, 0x78, 0x74, 0x12, 0x40, 0x0a, 0x0a, 0x63, 0x6f,
	0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x21,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x69,
	0x74, 0x68, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x77, 0x69, 0x74, 0x68, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x2f, 0x0a, 0x08, 0x43,
	0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x10, 0x02, 0x22, 0x45, 0x0a, 0x15,
	0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69,
	0x6d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x22, 0xdb, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x22, 0x62, 0x0a, 0x15, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69,
	0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4b, 0x0a, 0x14, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74,
	0x61, 0x67, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x61, 0x0a, 0x15, 0x43,
	0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x21, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x69, 0x6d,
	0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x63,
	0x6f, 0x64, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x7a, 0x69, 0x70, 0x10, 0x01, 0x2a, 0x79,
	0x0a, 0x13, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x1e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x69,
	0x6d, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x63, 0x6f, 0x64,
	0x65, 0x5f, 0x73, 0x69, 0x6d, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x63, 0x6f, 0x64, 0x65,
	0x5f, 0x73, 0x69, 0x6d, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x10, 0x02, 0x32, 0x9b, 0x02, 0x0a, 0x07, 0x43, 0x6f,
	0x64, 0x65, 0x53, 0x69, 0x6d, 0x12, 0x4b, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f,
	0x72, 0x6c, 0x64, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x53, 0x69, 0x6d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x53, 0x69, 0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x70, 0x62, 0x5f,
	0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_code_sim_code_sim_service_proto_rawDescOnce sync.Once
	file_code_sim_code_sim_service_proto_rawDescData = file_code_sim_code_sim_service_proto_rawDesc
)

func file_code_sim_code_sim_service_proto_rawDescGZIP() []byte {
	file_code_sim_code_sim_service_proto_rawDescOnce.Do(func() {
		file_code_sim_code_sim_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_code_sim_code_sim_service_proto_rawDescData)
	})
	return file_code_sim_code_sim_service_proto_rawDescData
}

var file_code_sim_code_sim_service_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_code_sim_code_sim_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_code_sim_code_sim_service_proto_goTypes = []interface{}{
	(CodeSimUploadFileType)(0),         // 0: pb.CodeSimUploadFileType
	(CodeSimUploadStatus)(0),           // 1: pb.CodeSimUploadStatus
	(CodeSimSearchRequest_CodeType)(0), // 2: pb.CodeSimSearchRequest.CodeType
	(*CodeSimHelloWorldRequest)(nil),   // 3: pb.CodeSimHelloWorldRequest
	(*CodeSimProject)(nil),             // 4: pb.CodeSimProject
	(*CodeSimProjectFile)(nil),         // 5: pb.CodeSimProjectFile
	(*CodeSimHelloWorldResponse)(nil),  // 6: pb.CodeSimHelloWorldResponse
	(*CodeSimSearchRequest)(nil),       // 7: pb.CodeSimSearchRequest
	(*CodeSimSearchResponse)(nil),      // 8: pb.CodeSimSearchResponse
	(*CodeSimUploadRequest)(nil),       // 9: pb.CodeSimUploadRequest
	(*CodeSimUploadResponse)(nil),      // 10: pb.CodeSimUploadResponse
	(*CodeSimDeleteRequest)(nil),       // 11: pb.CodeSimDeleteRequest
	(*CodeSimDeleteResponse)(nil),      // 12: pb.CodeSimDeleteResponse
}
var file_code_sim_code_sim_service_proto_depIdxs = []int32{
	4,  // 0: pb.CodeSimProjectFile.project_info:type_name -> pb.CodeSimProject
	2,  // 1: pb.CodeSimSearchRequest.code_types:type_name -> pb.CodeSimSearchRequest.CodeType
	5,  // 2: pb.CodeSimSearchResponse.files:type_name -> pb.CodeSimProjectFile
	4,  // 3: pb.CodeSimUploadRequest.project:type_name -> pb.CodeSimProject
	0,  // 4: pb.CodeSimUploadRequest.file_type:type_name -> pb.CodeSimUploadFileType
	1,  // 5: pb.CodeSimUploadResponse.status:type_name -> pb.CodeSimUploadStatus
	3,  // 6: pb.CodeSim.HelloWorld:input_type -> pb.CodeSimHelloWorldRequest
	7,  // 7: pb.CodeSim.Search:input_type -> pb.CodeSimSearchRequest
	9,  // 8: pb.CodeSim.Upload:input_type -> pb.CodeSimUploadRequest
	11, // 9: pb.CodeSim.Delete:input_type -> pb.CodeSimDeleteRequest
	6,  // 10: pb.CodeSim.HelloWorld:output_type -> pb.CodeSimHelloWorldResponse
	8,  // 11: pb.CodeSim.Search:output_type -> pb.CodeSimSearchResponse
	10, // 12: pb.CodeSim.Upload:output_type -> pb.CodeSimUploadResponse
	12, // 13: pb.CodeSim.Delete:output_type -> pb.CodeSimDeleteResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_code_sim_code_sim_service_proto_init() }
func file_code_sim_code_sim_service_proto_init() {
	if File_code_sim_code_sim_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_code_sim_code_sim_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeSimHelloWorldRequest); i {
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
		file_code_sim_code_sim_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeSimProject); i {
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
		file_code_sim_code_sim_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeSimProjectFile); i {
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
		file_code_sim_code_sim_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeSimHelloWorldResponse); i {
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
		file_code_sim_code_sim_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeSimSearchRequest); i {
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
		file_code_sim_code_sim_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeSimSearchResponse); i {
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
		file_code_sim_code_sim_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeSimUploadRequest); i {
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
		file_code_sim_code_sim_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeSimUploadResponse); i {
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
		file_code_sim_code_sim_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeSimDeleteRequest); i {
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
		file_code_sim_code_sim_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeSimDeleteResponse); i {
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
			RawDescriptor: file_code_sim_code_sim_service_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_code_sim_code_sim_service_proto_goTypes,
		DependencyIndexes: file_code_sim_code_sim_service_proto_depIdxs,
		EnumInfos:         file_code_sim_code_sim_service_proto_enumTypes,
		MessageInfos:      file_code_sim_code_sim_service_proto_msgTypes,
	}.Build()
	File_code_sim_code_sim_service_proto = out.File
	file_code_sim_code_sim_service_proto_rawDesc = nil
	file_code_sim_code_sim_service_proto_goTypes = nil
	file_code_sim_code_sim_service_proto_depIdxs = nil
}
