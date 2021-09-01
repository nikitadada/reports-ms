// Code generated by protoc-gen-go. DO NOT EDIT.
// source: citilink/cmsfiles/file/v1/file_api.proto

package filev1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Запрос создания файла.
type CreateRequest struct {
	// Название файла
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Тип файла
	Type Type `protobuf:"varint,2,opt,name=type,proto3,enum=citilink.cmsfiles.file.v1.Type" json:"type,omitempty"`
	// Расширение файла
	Extension            string   `protobuf:"bytes,3,opt,name=extension,proto3" json:"extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf23e0ed0cdee215, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_TYPE_INVALID
}

func (m *CreateRequest) GetExtension() string {
	if m != nil {
		return m.Extension
	}
	return ""
}

// Ответ на запрос создания файла.
type CreateResponse struct {
	// Идентификатор файла
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf23e0ed0cdee215, []int{1}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Чанк запроса загрузки файла.
type UploadRequest struct {
	// Идентификатор файла
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Чанк с данными
	Chunk                []byte   `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadRequest) Reset()         { *m = UploadRequest{} }
func (m *UploadRequest) String() string { return proto.CompactTextString(m) }
func (*UploadRequest) ProtoMessage()    {}
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf23e0ed0cdee215, []int{2}
}

func (m *UploadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRequest.Unmarshal(m, b)
}
func (m *UploadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRequest.Marshal(b, m, deterministic)
}
func (m *UploadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRequest.Merge(m, src)
}
func (m *UploadRequest) XXX_Size() int {
	return xxx_messageInfo_UploadRequest.Size(m)
}
func (m *UploadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRequest proto.InternalMessageInfo

func (m *UploadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UploadRequest) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

// Ответ на запрос загрузки файла.
type UploadResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadResponse) Reset()         { *m = UploadResponse{} }
func (m *UploadResponse) String() string { return proto.CompactTextString(m) }
func (*UploadResponse) ProtoMessage()    {}
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf23e0ed0cdee215, []int{3}
}

func (m *UploadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadResponse.Unmarshal(m, b)
}
func (m *UploadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadResponse.Marshal(b, m, deterministic)
}
func (m *UploadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadResponse.Merge(m, src)
}
func (m *UploadResponse) XXX_Size() int {
	return xxx_messageInfo_UploadResponse.Size(m)
}
func (m *UploadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadResponse proto.InternalMessageInfo

// Запрос получения файла по идентификатору.
type GetRequest struct {
	// Идентификатор файла
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf23e0ed0cdee215, []int{4}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Ответ на запрос получения файла по идентификатору.
type GetResponse struct {
	// Информация о файле
	File                 *File    `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf23e0ed0cdee215, []int{5}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetFile() *File {
	if m != nil {
		return m.File
	}
	return nil
}

// Запрос поиска файлов по критериям.
type FilterRequest struct {
	// Идентификаторы файлов для поиска
	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	// Тип файла
	Type Type `protobuf:"varint,2,opt,name=type,proto3,enum=citilink.cmsfiles.file.v1.Type" json:"type,omitempty"`
	// Статус загрузки файла
	Status Status `protobuf:"varint,3,opt,name=status,proto3,enum=citilink.cmsfiles.file.v1.Status" json:"status,omitempty"`
	// Смещение при выборке данных
	Offset int32 `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	// Максимальное количество возвращаемых файлов
	Limit int32 `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	// Признак нужно ли считать общее количество файлов
	ComputeTotal         bool     `protobuf:"varint,6,opt,name=compute_total,json=computeTotal,proto3" json:"compute_total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterRequest) Reset()         { *m = FilterRequest{} }
func (m *FilterRequest) String() string { return proto.CompactTextString(m) }
func (*FilterRequest) ProtoMessage()    {}
func (*FilterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf23e0ed0cdee215, []int{6}
}

func (m *FilterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterRequest.Unmarshal(m, b)
}
func (m *FilterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterRequest.Marshal(b, m, deterministic)
}
func (m *FilterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterRequest.Merge(m, src)
}
func (m *FilterRequest) XXX_Size() int {
	return xxx_messageInfo_FilterRequest.Size(m)
}
func (m *FilterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FilterRequest proto.InternalMessageInfo

func (m *FilterRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *FilterRequest) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_TYPE_INVALID
}

func (m *FilterRequest) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_STATUS_INVALID
}

func (m *FilterRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *FilterRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *FilterRequest) GetComputeTotal() bool {
	if m != nil {
		return m.ComputeTotal
	}
	return false
}

// Ответ на запрос поиска файлов по критериям.
type FilterResponse struct {
	// Массив удовлетворяющих фильтрации файлов
	Files []*File `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	// Смещение при выборке данных
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Максимальное количество возвращаемых файлов
	Limit int32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// Общее количество файлов
	Total                int32    `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterResponse) Reset()         { *m = FilterResponse{} }
func (m *FilterResponse) String() string { return proto.CompactTextString(m) }
func (*FilterResponse) ProtoMessage()    {}
func (*FilterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf23e0ed0cdee215, []int{7}
}

func (m *FilterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterResponse.Unmarshal(m, b)
}
func (m *FilterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterResponse.Marshal(b, m, deterministic)
}
func (m *FilterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterResponse.Merge(m, src)
}
func (m *FilterResponse) XXX_Size() int {
	return xxx_messageInfo_FilterResponse.Size(m)
}
func (m *FilterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FilterResponse proto.InternalMessageInfo

func (m *FilterResponse) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *FilterResponse) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *FilterResponse) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *FilterResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

// Запрос обновления файла.
type UpdateRequest struct {
	// Идентификатор обновляемого файла
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Название файла
	Name *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Тип файла
	Type *UpdateRequest_TypeValue `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// Статус загрузки файла
	Status               *UpdateRequest_StatusValue `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf23e0ed0cdee215, []int{8}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UpdateRequest) GetType() *UpdateRequest_TypeValue {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *UpdateRequest) GetStatus() *UpdateRequest_StatusValue {
	if m != nil {
		return m.Status
	}
	return nil
}

// Обертка над enum'ом Type для возможности передачи nil.
type UpdateRequest_TypeValue struct {
	Type                 Type     `protobuf:"varint,1,opt,name=type,proto3,enum=citilink.cmsfiles.file.v1.Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest_TypeValue) Reset()         { *m = UpdateRequest_TypeValue{} }
func (m *UpdateRequest_TypeValue) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest_TypeValue) ProtoMessage()    {}
func (*UpdateRequest_TypeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf23e0ed0cdee215, []int{8, 0}
}

func (m *UpdateRequest_TypeValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest_TypeValue.Unmarshal(m, b)
}
func (m *UpdateRequest_TypeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest_TypeValue.Marshal(b, m, deterministic)
}
func (m *UpdateRequest_TypeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest_TypeValue.Merge(m, src)
}
func (m *UpdateRequest_TypeValue) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest_TypeValue.Size(m)
}
func (m *UpdateRequest_TypeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest_TypeValue.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest_TypeValue proto.InternalMessageInfo

func (m *UpdateRequest_TypeValue) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_TYPE_INVALID
}

// Обертка над enum'ом Status для возможности передачи nil.
type UpdateRequest_StatusValue struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=citilink.cmsfiles.file.v1.Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest_StatusValue) Reset()         { *m = UpdateRequest_StatusValue{} }
func (m *UpdateRequest_StatusValue) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest_StatusValue) ProtoMessage()    {}
func (*UpdateRequest_StatusValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf23e0ed0cdee215, []int{8, 1}
}

func (m *UpdateRequest_StatusValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest_StatusValue.Unmarshal(m, b)
}
func (m *UpdateRequest_StatusValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest_StatusValue.Marshal(b, m, deterministic)
}
func (m *UpdateRequest_StatusValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest_StatusValue.Merge(m, src)
}
func (m *UpdateRequest_StatusValue) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest_StatusValue.Size(m)
}
func (m *UpdateRequest_StatusValue) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest_StatusValue.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest_StatusValue proto.InternalMessageInfo

func (m *UpdateRequest_StatusValue) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_STATUS_INVALID
}

// Ответ на запрос обновления файла.
type UpdateResponse struct {
	// Информация о файле
	File                 *File    `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf23e0ed0cdee215, []int{9}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetFile() *File {
	if m != nil {
		return m.File
	}
	return nil
}

// Запрос удаления файлов.
type DeleteRequest struct {
	// Идентификаторы файлов для удаления
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf23e0ed0cdee215, []int{10}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

// Ответ на запрос удаления файлов.
type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf23e0ed0cdee215, []int{11}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateRequest)(nil), "citilink.cmsfiles.file.v1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "citilink.cmsfiles.file.v1.CreateResponse")
	proto.RegisterType((*UploadRequest)(nil), "citilink.cmsfiles.file.v1.UploadRequest")
	proto.RegisterType((*UploadResponse)(nil), "citilink.cmsfiles.file.v1.UploadResponse")
	proto.RegisterType((*GetRequest)(nil), "citilink.cmsfiles.file.v1.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "citilink.cmsfiles.file.v1.GetResponse")
	proto.RegisterType((*FilterRequest)(nil), "citilink.cmsfiles.file.v1.FilterRequest")
	proto.RegisterType((*FilterResponse)(nil), "citilink.cmsfiles.file.v1.FilterResponse")
	proto.RegisterType((*UpdateRequest)(nil), "citilink.cmsfiles.file.v1.UpdateRequest")
	proto.RegisterType((*UpdateRequest_TypeValue)(nil), "citilink.cmsfiles.file.v1.UpdateRequest.TypeValue")
	proto.RegisterType((*UpdateRequest_StatusValue)(nil), "citilink.cmsfiles.file.v1.UpdateRequest.StatusValue")
	proto.RegisterType((*UpdateResponse)(nil), "citilink.cmsfiles.file.v1.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "citilink.cmsfiles.file.v1.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "citilink.cmsfiles.file.v1.DeleteResponse")
}

func init() {
	proto.RegisterFile("citilink/cmsfiles/file/v1/file_api.proto", fileDescriptor_bf23e0ed0cdee215)
}

var fileDescriptor_bf23e0ed0cdee215 = []byte{
	// 700 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0x96, 0xed, 0xc4, 0xff, 0x9f, 0x69, 0x13, 0x55, 0xab, 0x0a, 0xb9, 0x51, 0x81, 0xd4, 0x1c,
	0x64, 0x6e, 0x1c, 0x92, 0xd2, 0x0b, 0xee, 0xa0, 0x81, 0x14, 0x24, 0x2e, 0x2a, 0xb7, 0xf4, 0x02,
	0x35, 0xaa, 0xdc, 0x78, 0x53, 0x56, 0x75, 0x6c, 0x13, 0xaf, 0x03, 0x7d, 0x06, 0xde, 0x82, 0x4b,
	0x9e, 0x81, 0x27, 0xe0, 0x39, 0x78, 0x03, 0x5e, 0x00, 0xed, 0xa9, 0xb1, 0x0b, 0xdd, 0x1e, 0x6e,
	0x12, 0xef, 0xee, 0xcc, 0x37, 0xdf, 0xcc, 0x7c, 0x33, 0xe0, 0x8d, 0x09, 0x25, 0x31, 0x49, 0x4e,
	0xbb, 0xe3, 0x69, 0x3e, 0x21, 0x31, 0xce, 0xbb, 0xec, 0xb7, 0x3b, 0xef, 0xf1, 0xff, 0xa3, 0x30,
	0x23, 0x7e, 0x36, 0x4b, 0x69, 0x8a, 0xd6, 0x94, 0xa5, 0xaf, 0x2c, 0x7d, 0xf6, 0xeb, 0xcf, 0x7b,
	0xed, 0x87, 0x7a, 0x10, 0x01, 0xd0, 0xbe, 0x77, 0x92, 0xa6, 0x27, 0x31, 0xee, 0xf2, 0xd3, 0x71,
	0x31, 0xe9, 0x7e, 0x9e, 0x85, 0x59, 0x86, 0x67, 0xb9, 0x78, 0x77, 0xe7, 0xd0, 0x1c, 0xcc, 0x70,
	0x48, 0x71, 0x80, 0x3f, 0x15, 0x38, 0xa7, 0x08, 0x41, 0x2d, 0x09, 0xa7, 0xd8, 0x31, 0x3a, 0x86,
	0xd7, 0x08, 0xf8, 0x37, 0xda, 0x84, 0x1a, 0x3d, 0xcb, 0xb0, 0x63, 0x76, 0x0c, 0xaf, 0xd5, 0xbf,
	0xef, 0x5f, 0x4a, 0xca, 0xdf, 0x3f, 0xcb, 0x70, 0xc0, 0x8d, 0xd1, 0x3a, 0x34, 0xf0, 0x17, 0x8a,
	0x93, 0x9c, 0xa4, 0x89, 0x63, 0x71, 0xb4, 0xc5, 0x85, 0xdb, 0x81, 0x96, 0x8a, 0x9b, 0x67, 0x69,
	0x92, 0x63, 0xd4, 0x02, 0x93, 0x44, 0x32, 0xac, 0x49, 0x22, 0x77, 0x0b, 0x9a, 0xef, 0xb3, 0x38,
	0x0d, 0x23, 0xc5, 0xec, 0x82, 0x01, 0x5a, 0x85, 0xfa, 0xf8, 0x63, 0x91, 0x9c, 0x72, 0x5a, 0xcb,
	0x81, 0x38, 0xb8, 0x2b, 0xd0, 0x52, 0x6e, 0x02, 0xd8, 0x5d, 0x07, 0xd8, 0xc1, 0xf4, 0x12, 0x14,
	0x77, 0x1b, 0x96, 0xf8, 0xab, 0x64, 0xb1, 0x09, 0x35, 0x96, 0x0b, 0x37, 0x58, 0xd2, 0xa6, 0x3a,
	0x24, 0x31, 0x0e, 0xb8, 0xb1, 0xfb, 0xcb, 0x80, 0xe6, 0x90, 0xc4, 0x14, 0xcf, 0x54, 0x94, 0x15,
	0xb0, 0x48, 0x94, 0x3b, 0x46, 0xc7, 0xf2, 0x1a, 0x01, 0xfb, 0xbc, 0x5d, 0x0d, 0x9f, 0x83, 0x9d,
	0xd3, 0x90, 0x16, 0x39, 0x2f, 0x60, 0xab, 0xbf, 0xa1, 0x71, 0xdb, 0xe3, 0x86, 0x81, 0x74, 0x40,
	0x77, 0xc0, 0x4e, 0x27, 0x93, 0x1c, 0x53, 0xa7, 0xd6, 0x31, 0xbc, 0x7a, 0x20, 0x4f, 0xac, 0x6a,
	0x31, 0x99, 0x12, 0xea, 0xd4, 0xf9, 0xb5, 0x38, 0xa0, 0x07, 0xd0, 0x1c, 0xa7, 0xd3, 0xac, 0xa0,
	0xf8, 0x88, 0xa6, 0x34, 0x8c, 0x1d, 0xbb, 0x63, 0x78, 0xff, 0x07, 0xcb, 0xf2, 0x72, 0x9f, 0xdd,
	0xb9, 0x5f, 0x0d, 0x68, 0xa9, 0x34, 0x65, 0xb9, 0xb6, 0xa0, 0xce, 0x59, 0xf0, 0x4c, 0xaf, 0x51,
	0x2f, 0x61, 0x5d, 0x22, 0x67, 0xfe, 0x9b, 0x9c, 0x55, 0x26, 0xb7, 0x0a, 0x75, 0x41, 0x4a, 0x64,
	0x22, 0x0e, 0xee, 0x6f, 0x93, 0x09, 0x24, 0x2a, 0x49, 0xf7, 0xa2, 0x40, 0x9e, 0x4a, 0x29, 0x9b,
	0xbc, 0x97, 0xeb, 0xbe, 0x18, 0x05, 0x5f, 0x8d, 0x82, 0xbf, 0x47, 0x67, 0x24, 0x39, 0x39, 0x08,
	0xe3, 0x02, 0x4b, 0xa1, 0x0f, 0x65, 0x93, 0x2c, 0xee, 0xd1, 0xd7, 0x64, 0x53, 0x89, 0xcc, 0x5b,
	0x26, 0x71, 0x78, 0xdf, 0xde, 0x9d, 0xf7, 0xad, 0xc6, 0x91, 0x9e, 0x5d, 0x1b, 0x49, 0x74, 0x51,
	0x60, 0x49, 0x8c, 0xf6, 0x0b, 0x68, 0x9c, 0x07, 0x38, 0xd7, 0x91, 0x71, 0x03, 0x1d, 0xb5, 0xdf,
	0xc0, 0x52, 0x09, 0xb8, 0x24, 0x2b, 0xe3, 0x86, 0xb2, 0x72, 0x5f, 0xb3, 0xf1, 0x8a, 0xca, 0x73,
	0x7b, 0xab, 0x89, 0xd9, 0x80, 0xe6, 0x2b, 0x1c, 0xe3, 0x45, 0xef, 0xfe, 0x1a, 0x18, 0x36, 0xc8,
	0xca, 0x44, 0x44, 0xea, 0xff, 0xa8, 0xc1, 0x7f, 0x0c, 0xe3, 0xe5, 0xee, 0x5b, 0x34, 0x02, 0x5b,
	0xec, 0x0f, 0xe4, 0x69, 0x22, 0x56, 0x56, 0x5b, 0xfb, 0xc9, 0x35, 0x2c, 0x65, 0x52, 0x47, 0x60,
	0x8b, 0x2d, 0xa2, 0x85, 0xaf, 0xec, 0x27, 0x2d, 0x7c, 0x75, 0x25, 0x79, 0x06, 0xda, 0x07, 0x6b,
	0x07, 0x53, 0xf4, 0x48, 0xe3, 0xb3, 0x58, 0x5a, 0xed, 0xc7, 0x57, 0x99, 0x49, 0xda, 0x23, 0xb0,
	0xc5, 0x80, 0x6a, 0x69, 0x57, 0x56, 0x95, 0x96, 0xf6, 0x85, 0x69, 0x1f, 0xb1, 0xaa, 0x44, 0x57,
	0x15, 0xbd, 0x22, 0xe8, 0x2b, 0xaa, 0x52, 0x51, 0xd2, 0x08, 0x6c, 0xd1, 0x71, 0x2d, 0x7c, 0x45,
	0x37, 0x5a, 0xf8, 0xaa, 0x7c, 0xb6, 0x0b, 0xb8, 0x3b, 0x4e, 0xa7, 0x97, 0xdb, 0x6f, 0x2f, 0x73,
	0x71, 0x65, 0x64, 0x97, 0x2d, 0x88, 0x5d, 0xe3, 0x83, 0xcd, 0x1e, 0xe6, 0xbd, 0x6f, 0xa6, 0x35,
	0x18, 0x0c, 0xbf, 0x9b, 0x6b, 0x03, 0xe5, 0x39, 0x50, 0x9e, 0xcc, 0xc1, 0x3f, 0xe8, 0xfd, 0x5c,
	0xbc, 0x1d, 0xaa, 0xb7, 0x43, 0xf6, 0x76, 0x78, 0xd0, 0x3b, 0xb6, 0xf9, 0xbe, 0xd9, 0xfc, 0x13,
	0x00, 0x00, 0xff, 0xff, 0x37, 0x71, 0xcc, 0xee, 0xf5, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileAPIClient is the client API for FileAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileAPIClient interface {
	// Создает новый файл.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Загружает файл.
	Upload(ctx context.Context, opts ...grpc.CallOption) (FileAPI_UploadClient, error)
	// Получает информацию о файле.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Находит файлы по критериям.
	Filter(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (*FilterResponse, error)
	// Обновляет информацию о файле.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Удаляет информацию о файле.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type fileAPIClient struct {
	cc *grpc.ClientConn
}

func NewFileAPIClient(cc *grpc.ClientConn) FileAPIClient {
	return &fileAPIClient{cc}
}

func (c *fileAPIClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/citilink.cmsfiles.file.v1.FileAPI/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileAPIClient) Upload(ctx context.Context, opts ...grpc.CallOption) (FileAPI_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileAPI_serviceDesc.Streams[0], "/citilink.cmsfiles.file.v1.FileAPI/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileAPIUploadClient{stream}
	return x, nil
}

type FileAPI_UploadClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type fileAPIUploadClient struct {
	grpc.ClientStream
}

func (x *fileAPIUploadClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileAPIUploadClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileAPIClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/citilink.cmsfiles.file.v1.FileAPI/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileAPIClient) Filter(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (*FilterResponse, error) {
	out := new(FilterResponse)
	err := c.cc.Invoke(ctx, "/citilink.cmsfiles.file.v1.FileAPI/Filter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileAPIClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/citilink.cmsfiles.file.v1.FileAPI/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileAPIClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/citilink.cmsfiles.file.v1.FileAPI/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileAPIServer is the server API for FileAPI service.
type FileAPIServer interface {
	// Создает новый файл.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Загружает файл.
	Upload(FileAPI_UploadServer) error
	// Получает информацию о файле.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Находит файлы по критериям.
	Filter(context.Context, *FilterRequest) (*FilterResponse, error)
	// Обновляет информацию о файле.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Удаляет информацию о файле.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterFileAPIServer(s *grpc.Server, srv FileAPIServer) {
	s.RegisterService(&_FileAPI_serviceDesc, srv)
}

func _FileAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citilink.cmsfiles.file.v1.FileAPI/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileAPIServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileAPI_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileAPIServer).Upload(&fileAPIUploadServer{stream})
}

type FileAPI_UploadServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type fileAPIUploadServer struct {
	grpc.ServerStream
}

func (x *fileAPIUploadServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileAPIUploadServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citilink.cmsfiles.file.v1.FileAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileAPIServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileAPI_Filter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileAPIServer).Filter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citilink.cmsfiles.file.v1.FileAPI/Filter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileAPIServer).Filter(ctx, req.(*FilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileAPI_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileAPIServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citilink.cmsfiles.file.v1.FileAPI/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileAPIServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileAPI_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileAPIServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citilink.cmsfiles.file.v1.FileAPI/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileAPIServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "citilink.cmsfiles.file.v1.FileAPI",
	HandlerType: (*FileAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _FileAPI_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _FileAPI_Get_Handler,
		},
		{
			MethodName: "Filter",
			Handler:    _FileAPI_Filter_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FileAPI_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FileAPI_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _FileAPI_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "citilink/cmsfiles/file/v1/file_api.proto",
}