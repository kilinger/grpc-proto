// Code generated by protoc-gen-go. DO NOT EDIT.
// source: code.xxxxxx.com/micro/proto/brand/brand.proto

/*
Package brand is a generated protocol buffer package.

It is generated from these files:
	code.xxxxxx.com/micro/proto/brand/brand.proto

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
*/
package brand

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
	UserId      int64  `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Intro       string `protobuf:"bytes,4,opt,name=intro" json:"intro,omitempty"`
	LogoUrl     string `protobuf:"bytes,5,opt,name=logo_url,json=logoUrl" json:"logo_url,omitempty"`
	CoverUrl    string `protobuf:"bytes,6,opt,name=cover_url,json=coverUrl" json:"cover_url,omitempty"`
	UrlFragment string `protobuf:"bytes,7,opt,name=url_fragment,json=urlFragment" json:"url_fragment,omitempty"`
	CreatedAt   string `protobuf:"bytes,8,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
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

func (m *Record) GetIntro() string {
	if m != nil {
		return m.Intro
	}
	return ""
}

func (m *Record) GetLogoUrl() string {
	if m != nil {
		return m.LogoUrl
	}
	return ""
}

func (m *Record) GetCoverUrl() string {
	if m != nil {
		return m.CoverUrl
	}
	return ""
}

func (m *Record) GetUrlFragment() string {
	if m != nil {
		return m.UrlFragment
	}
	return ""
}

func (m *Record) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
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
	Limit     int64  `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	Offset    int64  `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	UserId    int64  `protobuf:"varint,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	NameIndex string `protobuf:"bytes,5,opt,name=name_index,json=nameIndex" json:"name_index,omitempty"`
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

func (m *ListRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListRequest) GetNameIndex() string {
	if m != nil {
		return m.NameIndex
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

func init() {
	proto.RegisterType((*Record)(nil), "brand.Record")
	proto.RegisterType((*CreateRequest)(nil), "brand.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "brand.CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "brand.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "brand.ReadResponse")
	proto.RegisterType((*DeleteRequest)(nil), "brand.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "brand.DeleteResponse")
	proto.RegisterType((*ListRequest)(nil), "brand.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "brand.ListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Brand service

type BrandClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
}

type brandClient struct {
	c           client.Client
	serviceName string
}

func NewBrandClient(serviceName string, c client.Client) BrandClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "brand"
	}
	return &brandClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *brandClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Brand.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandClient) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Brand.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandClient) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Brand.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandClient) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Brand.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Brand service

type BrandHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
}

func RegisterBrandHandler(s server.Server, hdlr BrandHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Brand{hdlr}, opts...))
}

type Brand struct {
	BrandHandler
}

func (h *Brand) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.BrandHandler.Create(ctx, in, out)
}

func (h *Brand) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.BrandHandler.Read(ctx, in, out)
}

func (h *Brand) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.BrandHandler.Delete(ctx, in, out)
}

func (h *Brand) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.BrandHandler.List(ctx, in, out)
}

func init() { proto.RegisterFile("code.xxxxxx.com/micro/proto/brand/brand.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0xdd, 0xb4, 0x4d, 0xda, 0x4e, 0x3f, 0x0e, 0x43, 0x81, 0xb0, 0xa8, 0x22, 0x44, 0x42, 0xec,
	0x85, 0x56, 0x2c, 0x82, 0x3d, 0xf3, 0x21, 0xa4, 0x95, 0xe0, 0x12, 0x89, 0x73, 0xe4, 0x8d, 0xa7,
	0x2b, 0x4b, 0x4e, 0x5c, 0x1c, 0x07, 0xf1, 0x0f, 0xf8, 0x85, 0x9c, 0xf9, 0x2b, 0xc8, 0x76, 0x12,
	0xda, 0x02, 0xd2, 0x5e, 0x92, 0xbc, 0xf7, 0xc6, 0x7e, 0x33, 0xf3, 0x14, 0x78, 0x51, 0x28, 0x4e,
	0x1b, 0x56, 0x18, 0xc3, 0xd4, 0xa6, 0x50, 0xe5, 0xb6, 0x14, 0x85, 0x56, 0xdb, 0xbd, 0x56, 0x46,
	0x6d, 0x6f, 0x34, 0xab, 0xb8, 0x7f, 0x6e, 0x1c, 0x83, 0xa1, 0x03, 0xe9, 0xcf, 0x00, 0xa2, 0x8c,
	0x0a, 0xa5, 0x39, 0x2e, 0x61, 0x20, 0x78, 0x1c, 0x24, 0xc1, 0xc5, 0x30, 0x1b, 0x08, 0x8e, 0x0f,
	0x61, 0xdc, 0xd4, 0xa4, 0x73, 0xc1, 0xe3, 0x81, 0x23, 0x23, 0x0b, 0xaf, 0x39, 0x22, 0x8c, 0x2a,
	0x56, 0x52, 0x3c, 0x4c, 0x82, 0x8b, 0x69, 0xe6, 0xbe, 0x71, 0x05, 0xa1, 0xa8, 0x8c, 0x56, 0xf1,
	0xc8, 0x91, 0x1e, 0xe0, 0x23, 0x98, 0x48, 0x75, 0xab, 0xf2, 0x46, 0xcb, 0x38, 0x74, 0xc2, 0xd8,
	0xe2, 0x2f, 0x5a, 0xe2, 0x63, 0x98, 0x16, 0xea, 0x1b, 0x69, 0xa7, 0x45, 0x4e, 0x9b, 0x38, 0xc2,
	0x8a, 0x4f, 0x61, 0xde, 0x68, 0x99, 0xef, 0x34, 0xbb, 0x2d, 0xa9, 0x32, 0xf1, 0xd8, 0xe9, 0xb3,
	0x46, 0xcb, 0x8f, 0x2d, 0x85, 0x6b, 0x80, 0x42, 0x13, 0x33, 0xc4, 0x73, 0x66, 0xe2, 0x89, 0x2b,
	0x98, 0xb6, 0xcc, 0x5b, 0x93, 0xbe, 0x81, 0xc5, 0x7b, 0x07, 0x32, 0xfa, 0xda, 0x50, 0x6d, 0xf0,
	0x19, 0x44, 0xda, 0xcd, 0xe9, 0x26, 0x9c, 0x5d, 0x2e, 0x36, 0x7e, 0x1b, 0x7e, 0xf8, 0xac, 0x15,
	0xd3, 0x04, 0x96, 0xdd, 0xb9, 0x7a, 0xaf, 0xaa, 0x9a, 0x4e, 0xd7, 0x92, 0xae, 0x61, 0x96, 0x11,
	0xe3, 0xdd, 0xbd, 0xa7, 0xf2, 0x6b, 0x98, 0x7b, 0xb9, 0x3d, 0x7e, 0x47, 0xdf, 0x27, 0xb0, 0xf8,
	0x40, 0x92, 0xfe, 0xf4, 0x7b, 0x7a, 0x6f, 0x02, 0xcb, 0xae, 0xe0, 0x3f, 0x8d, 0xfd, 0x08, 0x60,
	0xf6, 0x49, 0xd4, 0xa6, 0xbb, 0x61, 0x05, 0xa1, 0x14, 0xa5, 0x30, 0x6d, 0x89, 0x07, 0xf8, 0x00,
	0x22, 0xb5, 0xdb, 0xd5, 0x64, 0xba, 0x50, 0x3d, 0x3a, 0x4c, 0x7b, 0xf8, 0xcf, 0xb4, 0x47, 0x07,
	0x69, 0xaf, 0x01, 0xec, 0x3b, 0x17, 0x15, 0xa7, 0xef, 0x6d, 0xb2, 0x53, 0xcb, 0x5c, 0x5b, 0x22,
	0xfd, 0x0c, 0x73, 0xdf, 0x48, 0xdb, 0xe9, 0x0a, 0x42, 0xa3, 0x0c, 0x93, 0xae, 0x93, 0x30, 0xf3,
	0x00, 0x9f, 0xc3, 0x58, 0x53, 0xdd, 0x48, 0x53, 0xc7, 0x83, 0x64, 0xf8, 0xf7, 0x6a, 0x3a, 0xf5,
	0xf2, 0x57, 0x00, 0xe1, 0x3b, 0xab, 0xe0, 0x15, 0x44, 0x3e, 0x1d, 0x5c, 0xb5, 0xb5, 0x47, 0x21,
	0x9f, 0xdf, 0x3f, 0x61, 0xbd, 0x7f, 0x7a, 0x86, 0x2f, 0x61, 0x64, 0x53, 0x41, 0xec, 0x2d, 0xfa,
	0x04, 0xcf, 0xef, 0x1d, 0x71, 0xfd, 0x91, 0x2b, 0x88, 0xfc, 0xc2, 0x7b, 0xaf, 0xa3, 0x80, 0x7a,
	0xaf, 0xe3, 0x54, 0xbc, 0x97, 0x9d, 0xbe, 0xf7, 0x3a, 0xc8, 0xa4, 0xf7, 0x3a, 0x5c, 0x4f, 0x7a,
	0x76, 0x13, 0xb9, 0x7f, 0xf2, 0xd5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0x43, 0x95, 0xce,
	0xc4, 0x03, 0x00, 0x00,
}
