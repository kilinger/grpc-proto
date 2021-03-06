// Code generated by protoc-gen-go. DO NOT EDIT.
// source: code.xxxxxx.com/micro/proto/verification/verification.proto

/*
Package verification is a generated protocol buffer package.

It is generated from these files:
	code.xxxxxx.com/micro/proto/verification/verification.proto

It has these top-level messages:
	Pair
	Record
	CreateRequest
	CreateResponse
	UpdateRequest
	UpdateResponse
	ReadRequest
	ReadResponse
	ListRequest
	ListResponse
*/
package verification

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

type Pair struct {
	Key    string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
}

func (m *Pair) Reset()                    { *m = Pair{} }
func (m *Pair) String() string            { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()               {}
func (*Pair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Pair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Pair) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type Record struct {
	Id   int64            `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Data map[string]*Pair `protobuf:"bytes,2,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Record) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Record) GetData() map[string]*Pair {
	if m != nil {
		return m.Data
	}
	return nil
}

type CreateRequest struct {
	Data map[string]*Pair `protobuf:"bytes,1,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateRequest) GetData() map[string]*Pair {
	if m != nil {
		return m.Data
	}
	return nil
}

type CreateResponse struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UpdateRequest struct {
	Id   int64            `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Data map[string]*Pair `protobuf:"bytes,2,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateRequest) GetData() map[string]*Pair {
	if m != nil {
		return m.Data
	}
	return nil
}

type UpdateResponse struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateResponse) GetId() int64 {
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
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadResponse struct {
	Id   int64            `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Data map[string]*Pair `protobuf:"bytes,2,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReadResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReadResponse) GetData() map[string]*Pair {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListRequest struct {
	Data map[string]*Pair `protobuf:"bytes,1,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListRequest) GetData() map[string]*Pair {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListResponse struct {
	Count   int64     `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	Results []*Record `protobuf:"bytes,2,rep,name=results" json:"results,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListResponse) GetCount() int64 {
	if m != nil {
		return m.Count
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
	proto.RegisterType((*Pair)(nil), "verification.Pair")
	proto.RegisterType((*Record)(nil), "verification.Record")
	proto.RegisterType((*CreateRequest)(nil), "verification.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "verification.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "verification.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "verification.UpdateResponse")
	proto.RegisterType((*ReadRequest)(nil), "verification.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "verification.ReadResponse")
	proto.RegisterType((*ListRequest)(nil), "verification.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "verification.ListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Verification service

type VerificationClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
}

type verificationClient struct {
	c           client.Client
	serviceName string
}

func NewVerificationClient(serviceName string, c client.Client) VerificationClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "verification"
	}
	return &verificationClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *verificationClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Verification.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verificationClient) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Verification.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verificationClient) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Verification.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verificationClient) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Verification.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Verification service

type VerificationHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
}

func RegisterVerificationHandler(s server.Server, hdlr VerificationHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Verification{hdlr}, opts...))
}

type Verification struct {
	VerificationHandler
}

func (h *Verification) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.VerificationHandler.Create(ctx, in, out)
}

func (h *Verification) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.VerificationHandler.Update(ctx, in, out)
}

func (h *Verification) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.VerificationHandler.Read(ctx, in, out)
}

func (h *Verification) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.VerificationHandler.List(ctx, in, out)
}

func init() {
	proto.RegisterFile("code.xxxxxx.com/micro/proto/verification/verification.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x3d, 0xaf, 0xd3, 0x30,
	0x14, 0xad, 0xd3, 0xbe, 0xa0, 0xde, 0xe4, 0x3d, 0x21, 0xeb, 0x09, 0x95, 0xf2, 0xa1, 0xc8, 0x80,
	0x94, 0x29, 0x45, 0x61, 0xa0, 0x85, 0x81, 0x01, 0x3a, 0xc1, 0x80, 0x2c, 0x60, 0x37, 0x89, 0x91,
	0x2c, 0xda, 0xb8, 0x38, 0x4e, 0xa5, 0xfe, 0x0e, 0x06, 0xc4, 0xc2, 0x86, 0xf8, 0x9b, 0x28, 0x71,
	0x02, 0xb1, 0x95, 0xbe, 0x2d, 0x5b, 0x6c, 0xdf, 0x73, 0xee, 0xb9, 0xc7, 0xc7, 0x81, 0x97, 0x99,
	0xcc, 0x79, 0xc2, 0x32, 0xad, 0x99, 0x4c, 0x32, 0xb9, 0x5f, 0xed, 0x45, 0xa6, 0xe4, 0xea, 0xa0,
	0xa4, 0x96, 0xab, 0x23, 0x57, 0xe2, 0x8b, 0xc8, 0x98, 0x16, 0xb2, 0xb0, 0x16, 0x49, 0x73, 0x8e,
	0xc3, 0xfe, 0x1e, 0x79, 0x0a, 0xb3, 0xf7, 0x4c, 0x28, 0x7c, 0x1b, 0xa6, 0x5f, 0xf9, 0x69, 0x81,
	0x22, 0x14, 0xcf, 0x69, 0xfd, 0x89, 0xef, 0x80, 0x7f, 0x64, 0xbb, 0x8a, 0x97, 0x0b, 0x2f, 0x9a,
	0xc6, 0x73, 0xda, 0xae, 0xc8, 0x4f, 0x04, 0x3e, 0xe5, 0x99, 0x54, 0x39, 0xbe, 0x02, 0x4f, 0xe4,
	0x0d, 0x66, 0x4a, 0x3d, 0x91, 0xe3, 0x14, 0x66, 0x39, 0xd3, 0xac, 0x01, 0x04, 0xe9, 0xc3, 0xc4,
	0xea, 0x6e, 0x30, 0xc9, 0x1b, 0xa6, 0xd9, 0xb6, 0xd0, 0xea, 0x44, 0x9b, 0xda, 0xe5, 0x5b, 0x98,
	0xff, 0xdb, 0x1a, 0x50, 0x11, 0xc3, 0x45, 0xd3, 0x77, 0xe1, 0x45, 0x28, 0x0e, 0x52, 0x6c, 0x73,
	0xd6, 0xd2, 0xa9, 0x29, 0x78, 0xe1, 0xad, 0x11, 0xf9, 0x81, 0xe0, 0xf2, 0xb5, 0xe2, 0x4c, 0x73,
	0xca, 0xbf, 0x55, 0xbc, 0xd4, 0x78, 0xd3, 0x4a, 0x42, 0x8d, 0xa4, 0x27, 0x36, 0xdc, 0x2a, 0x1d,
	0x57, 0x59, 0x04, 0x57, 0x5d, 0xb7, 0xf2, 0x20, 0x8b, 0x92, 0xbb, 0xe6, 0x91, 0x3f, 0x08, 0x2e,
	0x3f, 0x1e, 0xf2, 0x9e, 0x76, 0xd7, 0xde, 0x8d, 0x65, 0xaf, 0x33, 0x8b, 0x05, 0x1d, 0x7d, 0x96,
	0xae, 0xdb, 0x99, 0x59, 0x1e, 0x40, 0x40, 0x39, 0xcb, 0xcf, 0x0c, 0x42, 0x7e, 0x23, 0x08, 0xcd,
	0xf9, 0x30, 0x1e, 0xaf, 0xad, 0x49, 0x1f, 0xbb, 0x41, 0xfa, 0x8f, 0x1c, 0x77, 0xd0, 0xef, 0x08,
	0x82, 0x77, 0xa2, 0xd4, 0xdd, 0x1c, 0xcf, 0xad, 0x30, 0x3d, 0xb2, 0xc1, 0xbd, 0xc2, 0x71, 0x55,
	0x7d, 0x80, 0xd0, 0xf4, 0x6a, 0xcd, 0xbb, 0x86, 0x8b, 0x4c, 0x56, 0x85, 0x6e, 0xfd, 0x33, 0x0b,
	0x9c, 0xc0, 0x2d, 0xc5, 0xcb, 0x6a, 0xa7, 0xcb, 0xd6, 0xc5, 0xeb, 0xa1, 0xe7, 0x48, 0xbb, 0xa2,
	0xf4, 0x97, 0x07, 0xe1, 0xa7, 0x5e, 0x01, 0xde, 0x82, 0x6f, 0x12, 0x8b, 0xef, 0xdd, 0xf0, 0x6a,
	0x96, 0xf7, 0x87, 0x0f, 0x8d, 0x36, 0x32, 0xa9, 0x69, 0x4c, 0x58, 0x5c, 0x1a, 0x2b, 0xb0, 0x2e,
	0x8d, 0x9d, 0x2f, 0x32, 0xc1, 0xaf, 0x60, 0x56, 0xdf, 0x3b, 0xbe, 0x3b, 0x94, 0x05, 0x43, 0xb1,
	0x3c, 0x1f, 0x13, 0x43, 0x50, 0xbb, 0xe6, 0x12, 0xf4, 0x6e, 0xcd, 0x25, 0xe8, 0x9b, 0x4c, 0x26,
	0x9f, 0xfd, 0xe6, 0xf7, 0xf9, 0xec, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xe9, 0x44, 0x2d,
	0x7d, 0x05, 0x00, 0x00,
}
