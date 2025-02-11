// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: staffer-protos/vacancy.proto

package genprotos

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

type VacancyCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description  string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	DepartmentId string `protobuf:"bytes,3,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
	PositionId   string `protobuf:"bytes,4,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`
	Status       string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *VacancyCreate) Reset() {
	*x = VacancyCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staffer_protos_vacancy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VacancyCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VacancyCreate) ProtoMessage() {}

func (x *VacancyCreate) ProtoReflect() protoreflect.Message {
	mi := &file_staffer_protos_vacancy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VacancyCreate.ProtoReflect.Descriptor instead.
func (*VacancyCreate) Descriptor() ([]byte, []int) {
	return file_staffer_protos_vacancy_proto_rawDescGZIP(), []int{0}
}

func (x *VacancyCreate) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *VacancyCreate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *VacancyCreate) GetDepartmentId() string {
	if x != nil {
		return x.DepartmentId
	}
	return ""
}

func (x *VacancyCreate) GetPositionId() string {
	if x != nil {
		return x.PositionId
	}
	return ""
}

func (x *VacancyCreate) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type VacancyGetResUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title        string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DepartmentId string `protobuf:"bytes,4,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
	PositionId   string `protobuf:"bytes,5,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`
	Status       string `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *VacancyGetResUpdateReq) Reset() {
	*x = VacancyGetResUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staffer_protos_vacancy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VacancyGetResUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VacancyGetResUpdateReq) ProtoMessage() {}

func (x *VacancyGetResUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_staffer_protos_vacancy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VacancyGetResUpdateReq.ProtoReflect.Descriptor instead.
func (*VacancyGetResUpdateReq) Descriptor() ([]byte, []int) {
	return file_staffer_protos_vacancy_proto_rawDescGZIP(), []int{1}
}

func (x *VacancyGetResUpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VacancyGetResUpdateReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *VacancyGetResUpdateReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *VacancyGetResUpdateReq) GetDepartmentId() string {
	if x != nil {
		return x.DepartmentId
	}
	return ""
}

func (x *VacancyGetResUpdateReq) GetPositionId() string {
	if x != nil {
		return x.PositionId
	}
	return ""
}

func (x *VacancyGetResUpdateReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type VacancyGetAllReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string      `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description  string      `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	DepartmentId string      `protobuf:"bytes,3,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
	PositionId   string      `protobuf:"bytes,4,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`
	Status       string      `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Pagination   *Pagination `protobuf:"bytes,6,opt,name=Pagination,proto3" json:"Pagination,omitempty"`
}

func (x *VacancyGetAllReq) Reset() {
	*x = VacancyGetAllReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staffer_protos_vacancy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VacancyGetAllReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VacancyGetAllReq) ProtoMessage() {}

func (x *VacancyGetAllReq) ProtoReflect() protoreflect.Message {
	mi := &file_staffer_protos_vacancy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VacancyGetAllReq.ProtoReflect.Descriptor instead.
func (*VacancyGetAllReq) Descriptor() ([]byte, []int) {
	return file_staffer_protos_vacancy_proto_rawDescGZIP(), []int{2}
}

func (x *VacancyGetAllReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *VacancyGetAllReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *VacancyGetAllReq) GetDepartmentId() string {
	if x != nil {
		return x.DepartmentId
	}
	return ""
}

func (x *VacancyGetAllReq) GetPositionId() string {
	if x != nil {
		return x.PositionId
	}
	return ""
}

func (x *VacancyGetAllReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *VacancyGetAllReq) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type VacancyGetAllRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vacancies *VacancyGetResUpdateReq `protobuf:"bytes,1,opt,name=Vacancies,proto3" json:"Vacancies,omitempty"`
	Success   bool                    `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *VacancyGetAllRes) Reset() {
	*x = VacancyGetAllRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staffer_protos_vacancy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VacancyGetAllRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VacancyGetAllRes) ProtoMessage() {}

func (x *VacancyGetAllRes) ProtoReflect() protoreflect.Message {
	mi := &file_staffer_protos_vacancy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VacancyGetAllRes.ProtoReflect.Descriptor instead.
func (*VacancyGetAllRes) Descriptor() ([]byte, []int) {
	return file_staffer_protos_vacancy_proto_rawDescGZIP(), []int{3}
}

func (x *VacancyGetAllRes) GetVacancies() *VacancyGetResUpdateReq {
	if x != nil {
		return x.Vacancies
	}
	return nil
}

func (x *VacancyGetAllRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_staffer_protos_vacancy_proto protoreflect.FileDescriptor

var file_staffer_protos_vacancy_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x74, 0x61, 0x66, 0x66, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x1a, 0x19, 0x73, 0x74, 0x61, 0x66, 0x66, 0x65, 0x72, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa5, 0x01, 0x0a, 0x0d, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xbe, 0x01, 0x0a, 0x16, 0x56, 0x61, 0x63,
	0x61, 0x6e, 0x63, 0x79, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xdb, 0x01, 0x0a, 0x10, 0x56, 0x61,
	0x63, 0x61, 0x6e, 0x63, 0x79, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x69, 0x0a, 0x10, 0x56, 0x61, 0x63, 0x61, 0x6e,
	0x63, 0x79, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x09, 0x56,
	0x61, 0x63, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x52, 0x09, 0x56,
	0x61, 0x63, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x32, 0x8a, 0x02, 0x0a, 0x0e, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x14, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x0b, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x35, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0b, 0x2e,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x42, 0x79, 0x69, 0x64, 0x1a, 0x1d, 0x2e, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x2e, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x34, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x56, 0x61, 0x63, 0x61,
	0x6e, 0x63, 0x79, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12,
	0x22, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x2e, 0x42, 0x79, 0x69, 0x64, 0x1a, 0x0b, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x17, 0x2e,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x56,
	0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x42,
	0x0c, 0x5a, 0x0a, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_staffer_protos_vacancy_proto_rawDescOnce sync.Once
	file_staffer_protos_vacancy_proto_rawDescData = file_staffer_protos_vacancy_proto_rawDesc
)

func file_staffer_protos_vacancy_proto_rawDescGZIP() []byte {
	file_staffer_protos_vacancy_proto_rawDescOnce.Do(func() {
		file_staffer_protos_vacancy_proto_rawDescData = protoimpl.X.CompressGZIP(file_staffer_protos_vacancy_proto_rawDescData)
	})
	return file_staffer_protos_vacancy_proto_rawDescData
}

var file_staffer_protos_vacancy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_staffer_protos_vacancy_proto_goTypes = []interface{}{
	(*VacancyCreate)(nil),          // 0: staff.VacancyCreate
	(*VacancyGetResUpdateReq)(nil), // 1: staff.VacancyGetResUpdateReq
	(*VacancyGetAllReq)(nil),       // 2: staff.VacancyGetAllReq
	(*VacancyGetAllRes)(nil),       // 3: staff.VacancyGetAllRes
	(*Pagination)(nil),             // 4: staff.Pagination
	(*Byid)(nil),                   // 5: staff.Byid
	(*Void)(nil),                   // 6: staff.Void
}
var file_staffer_protos_vacancy_proto_depIdxs = []int32{
	4, // 0: staff.VacancyGetAllReq.Pagination:type_name -> staff.Pagination
	1, // 1: staff.VacancyGetAllRes.Vacancies:type_name -> staff.VacancyGetResUpdateReq
	0, // 2: staff.VacancyService.Create:input_type -> staff.VacancyCreate
	5, // 3: staff.VacancyService.GetByID:input_type -> staff.Byid
	1, // 4: staff.VacancyService.Update:input_type -> staff.VacancyGetResUpdateReq
	5, // 5: staff.VacancyService.Delete:input_type -> staff.Byid
	2, // 6: staff.VacancyService.GetAll:input_type -> staff.VacancyGetAllReq
	6, // 7: staff.VacancyService.Create:output_type -> staff.Void
	1, // 8: staff.VacancyService.GetByID:output_type -> staff.VacancyGetResUpdateReq
	6, // 9: staff.VacancyService.Update:output_type -> staff.Void
	6, // 10: staff.VacancyService.Delete:output_type -> staff.Void
	3, // 11: staff.VacancyService.GetAll:output_type -> staff.VacancyGetAllRes
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_staffer_protos_vacancy_proto_init() }
func file_staffer_protos_vacancy_proto_init() {
	if File_staffer_protos_vacancy_proto != nil {
		return
	}
	file_staffer_protos_void_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_staffer_protos_vacancy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VacancyCreate); i {
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
		file_staffer_protos_vacancy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VacancyGetResUpdateReq); i {
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
		file_staffer_protos_vacancy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VacancyGetAllReq); i {
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
		file_staffer_protos_vacancy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VacancyGetAllRes); i {
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
			RawDescriptor: file_staffer_protos_vacancy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_staffer_protos_vacancy_proto_goTypes,
		DependencyIndexes: file_staffer_protos_vacancy_proto_depIdxs,
		MessageInfos:      file_staffer_protos_vacancy_proto_msgTypes,
	}.Build()
	File_staffer_protos_vacancy_proto = out.File
	file_staffer_protos_vacancy_proto_rawDesc = nil
	file_staffer_protos_vacancy_proto_goTypes = nil
	file_staffer_protos_vacancy_proto_depIdxs = nil
}
