// Code generated by protoc-gen-go. DO NOT EDIT.
// source: code.xxxxxx.com/micro/proto/category/category.proto

/*
Package category is a generated protocol buffer package.

It is generated from these files:
	code.xxxxxx.com/micro/proto/category/category.proto

It has these top-level messages:
	Record
	CreateRequest
	CreateResponse
	ReadRequest
	ReadResponse
	DeleteRequest
	DeleteResponse
	ListRequest
	ListResponse
	StatsRequest
	StatsResponse
*/
package category

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Record struct {
	Id          int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ParentId    int64  `protobuf:"varint,2,opt,name=parent_id,json=parentId" json:"parent_id,omitempty"`
	UserId      int64  `protobuf:"varint,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Name        string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	CreatedAt   string `protobuf:"bytes,6,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Position    int32  `protobuf:"varint,7,opt,name=position" json:"position,omitempty"`
	Tag         string `protobuf:"bytes,8,opt,name=tag" json:"tag,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Record) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Record) GetParentId() int64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *Record) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Record) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Record) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Record) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Record) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *Record) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type CreateRequest struct {
	Record *Record `protobuf:"bytes,1,opt,name=record" json:"record,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateRequest) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type CreateResponse struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadResponse struct {
	Record *Record `protobuf:"bytes,1,opt,name=record" json:"record,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReadResponse) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type DeleteRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteResponse struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListRequest struct {
	Limit    int64  `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	Offset   int64  `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	UserId   int64  `protobuf:"varint,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	ParentId int64  `protobuf:"varint,4,opt,name=parent_id,json=parentId" json:"parent_id,omitempty"`
	Tag      string `protobuf:"bytes,5,opt,name=tag" json:"tag,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ListRequest) GetParentId() int64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *ListRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type ListResponse struct {
	Total   int32     `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
	Results []*Record `protobuf:"bytes,2,rep,name=results" json:"results,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ListResponse) GetResults() []*Record {
	if m != nil {
		return m.Results
	}
	return nil
}

type StatsRequest struct {
	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Tag    string `protobuf:"bytes,2,opt,name=tag" json:"tag,omitempty"`
}

func (m *StatsRequest) Reset()                    { *m = StatsRequest{} }
func (m *StatsRequest) String() string            { return proto.CompactTextString(m) }
func (*StatsRequest) ProtoMessage()               {}
func (*StatsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *StatsRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *StatsRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type StatsResponse struct {
	Total int32 `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
}

func (m *StatsResponse) Reset()                    { *m = StatsResponse{} }
func (m *StatsResponse) String() string            { return proto.CompactTextString(m) }
func (*StatsResponse) ProtoMessage()               {}
func (*StatsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *StatsResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*Record)(nil), "category.Record")
	proto.RegisterType((*CreateRequest)(nil), "category.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "category.CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "category.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "category.ReadResponse")
	proto.RegisterType((*DeleteRequest)(nil), "category.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "category.DeleteResponse")
	proto.RegisterType((*ListRequest)(nil), "category.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "category.ListResponse")
	proto.RegisterType((*StatsRequest)(nil), "category.StatsRequest")
	proto.RegisterType((*StatsResponse)(nil), "category.StatsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Category service

type CategoryClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Stats(ctx context.Context, in *StatsRequest, opts ...client.CallOption) (*StatsResponse, error)
}

type categoryClient struct {
	c           client.Client
	serviceName string
}

func NewCategoryClient(serviceName string, c client.Client) CategoryClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "category"
	}
	return &categoryClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *categoryClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Category.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryClient) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Category.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryClient) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Category.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryClient) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Category.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryClient) Stats(ctx context.Context, in *StatsRequest, opts ...client.CallOption) (*StatsResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Category.Stats", in)
	out := new(StatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Category service

type CategoryHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
	Stats(context.Context, *StatsRequest, *StatsResponse) error
}

func RegisterCategoryHandler(s server.Server, hdlr CategoryHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Category{hdlr}, opts...))
}

type Category struct {
	CategoryHandler
}

func (h *Category) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.CategoryHandler.Create(ctx, in, out)
}

func (h *Category) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.CategoryHandler.Read(ctx, in, out)
}

func (h *Category) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.CategoryHandler.Delete(ctx, in, out)
}

func (h *Category) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.CategoryHandler.List(ctx, in, out)
}

func (h *Category) Stats(ctx context.Context, in *StatsRequest, out *StatsResponse) error {
	return h.CategoryHandler.Stats(ctx, in, out)
}

func init() {
	proto.RegisterFile("code.xxxxxx.com/micro/proto/category/category.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0x8e, 0x64, 0x4b, 0x96, 0xc7, 0x3f, 0x84, 0x25, 0x8d, 0x85, 0x4b, 0xa8, 0x10, 0x14, 0x4c,
	0x0f, 0x36, 0x24, 0x87, 0x36, 0x85, 0x1c, 0x4a, 0x7a, 0x09, 0xf4, 0x50, 0xd4, 0x07, 0x08, 0x5b,
	0xed, 0x26, 0x2c, 0xd8, 0x5e, 0x75, 0x77, 0x7c, 0xe8, 0xb9, 0xef, 0xd4, 0xf7, 0xe8, 0x1b, 0x15,
	0xed, 0xae, 0xe4, 0x95, 0xd2, 0x94, 0xdc, 0x34, 0x3f, 0xdf, 0xce, 0x37, 0xdf, 0x7c, 0x08, 0xae,
	0x4a, 0xc9, 0xf8, 0x9a, 0x96, 0x88, 0x54, 0xae, 0x4b, 0xb9, 0xdb, 0xec, 0x44, 0xa9, 0xe4, 0xa6,
	0x52, 0x12, 0xe5, 0xa6, 0xa4, 0xc8, 0x1f, 0xa5, 0xfa, 0xd9, 0x7e, 0xac, 0x4d, 0x9e, 0x24, 0x4d,
	0x9c, 0xff, 0x09, 0x20, 0x2e, 0x78, 0x29, 0x15, 0x23, 0x73, 0x08, 0x05, 0x4b, 0x83, 0x2c, 0x58,
	0x0d, 0x8a, 0x50, 0x30, 0xf2, 0x1a, 0xc6, 0x15, 0x55, 0x7c, 0x8f, 0xf7, 0x82, 0xa5, 0xa1, 0x49,
	0x27, 0x36, 0x71, 0xc7, 0xc8, 0x02, 0x46, 0x07, 0xcd, 0x55, 0x5d, 0x1a, 0x98, 0x52, 0x5c, 0x87,
	0x77, 0x8c, 0x10, 0x18, 0xee, 0xe9, 0x8e, 0xa7, 0xc3, 0x2c, 0x58, 0x8d, 0x0b, 0xf3, 0x4d, 0x32,
	0x98, 0x30, 0xae, 0x4b, 0x25, 0x2a, 0x14, 0x72, 0x9f, 0x46, 0xa6, 0xe4, 0xa7, 0xc8, 0x05, 0x40,
	0xa9, 0x38, 0x45, 0xce, 0xee, 0x29, 0xa6, 0xb1, 0x69, 0x18, 0xbb, 0xcc, 0x27, 0x24, 0x4b, 0x48,
	0x2a, 0xa9, 0x85, 0x41, 0x8f, 0xb2, 0x60, 0x15, 0x15, 0x6d, 0x4c, 0x4e, 0x61, 0x80, 0xf4, 0x31,
	0x4d, 0x0c, 0xa6, 0xfe, 0xcc, 0xaf, 0x61, 0x76, 0x6b, 0xa0, 0x05, 0xff, 0x71, 0xe0, 0x1a, 0xc9,
	0x0a, 0x62, 0x65, 0x76, 0x34, 0xdb, 0x4d, 0x2e, 0x4f, 0xd7, 0xad, 0x1e, 0x76, 0xf7, 0xc2, 0xd5,
	0xf3, 0x0c, 0xe6, 0x0d, 0x54, 0x57, 0x72, 0xaf, 0x79, 0x5f, 0x95, 0xfc, 0x02, 0x26, 0x05, 0xa7,
	0xac, 0x79, 0xba, 0x5f, 0xfe, 0x00, 0x53, 0x5b, 0x76, 0xf0, 0x97, 0x8f, 0x7e, 0x03, 0xb3, 0xcf,
	0x7c, 0xcb, 0x8f, 0xac, 0xfb, 0x4f, 0x67, 0x30, 0x6f, 0x1a, 0x9e, 0xe1, 0xf6, 0x2b, 0x80, 0xc9,
	0x17, 0xa1, 0xb1, 0x79, 0xe1, 0x0c, 0xa2, 0xad, 0xd8, 0x09, 0x74, 0x2d, 0x36, 0x20, 0xe7, 0x10,
	0xcb, 0x87, 0x07, 0xcd, 0xd1, 0x1d, 0xd5, 0x45, 0xcf, 0x9f, 0xb4, 0x63, 0x84, 0x61, 0xcf, 0x08,
	0x4e, 0xfe, 0xe8, 0x28, 0xff, 0x57, 0x98, 0x5a, 0x12, 0x8e, 0xe5, 0x19, 0x44, 0x28, 0x91, 0x6e,
	0x0d, 0x8b, 0xa8, 0xb0, 0x01, 0x79, 0x07, 0x23, 0xc5, 0xf5, 0x61, 0x8b, 0x3a, 0x0d, 0xb3, 0xc1,
	0x3f, 0x95, 0x69, 0x1a, 0xf2, 0x6b, 0x98, 0x7e, 0x43, 0x8a, 0xba, 0xd9, 0xcb, 0x63, 0x1a, 0x74,
	0x98, 0x3a, 0x32, 0xe1, 0x91, 0xcc, 0x5b, 0x98, 0x39, 0xe8, 0xff, 0xd8, 0x5c, 0xfe, 0x0e, 0x21,
	0xb9, 0x75, 0xe3, 0xc9, 0x0d, 0xc4, 0xd6, 0x04, 0x64, 0x71, 0xe4, 0xd4, 0x71, 0xd4, 0x32, 0x7d,
	0x5a, 0xb0, 0xef, 0xe7, 0x27, 0xe4, 0x3d, 0x0c, 0x6b, 0x0b, 0x90, 0x57, 0xfe, 0x42, 0xad, 0x63,
	0x96, 0xe7, 0xfd, 0x74, 0x0b, 0xbc, 0x81, 0xd8, 0x1e, 0xd8, 0x9f, 0xdb, 0xf1, 0x84, 0x3f, 0xb7,
	0xeb, 0x05, 0x3b, 0xb7, 0xd6, 0xdd, 0x9f, 0xeb, 0x99, 0xc1, 0x9f, 0xeb, 0x9f, 0x27, 0x3f, 0x21,
	0x1f, 0x21, 0x32, 0x1a, 0x11, 0xaf, 0xc5, 0xd7, 0x7b, 0xb9, 0x78, 0x92, 0x6f, 0xb0, 0xdf, 0x63,
	0xf3, 0x43, 0xb9, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x69, 0xb5, 0x9b, 0x87, 0x04, 0x00,
	0x00,
}
