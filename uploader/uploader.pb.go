// Code generated by protoc-gen-go. DO NOT EDIT.
// source: code.xxxxxx.com/micro/proto/uploader/uploader.proto

/*
Package uploader is a generated protocol buffer package.

It is generated from these files:
	code.xxxxxx.com/micro/proto/uploader/uploader.proto

It has these top-level messages:
	Request
	Response
*/
package uploader

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

type Request struct {
	Bucket      string `protobuf:"bytes,1,opt,name=bucket" json:"bucket,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	ContentType string `protobuf:"bytes,3,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	Expires     int32  `protobuf:"varint,4,opt,name=expires" json:"expires,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Request) GetExpires() int32 {
	if m != nil {
		return m.Expires
	}
	return 0
}

type Response struct {
	Url       string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	UploadUrl string `protobuf:"bytes,2,opt,name=upload_url,json=uploadUrl" json:"upload_url,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Response) GetUploadUrl() string {
	if m != nil {
		return m.UploadUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "uploader.Request")
	proto.RegisterType((*Response)(nil), "uploader.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Uploader service

type UploaderClient interface {
	Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type uploaderClient struct {
	c           client.Client
	serviceName string
}

func NewUploaderClient(serviceName string, c client.Client) UploaderClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "uploader"
	}
	return &uploaderClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *uploaderClient) Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Uploader.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Uploader service

type UploaderHandler interface {
	Create(context.Context, *Request, *Response) error
}

func RegisterUploaderHandler(s server.Server, hdlr UploaderHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Uploader{hdlr}, opts...))
}

type Uploader struct {
	UploaderHandler
}

func (h *Uploader) Create(ctx context.Context, in *Request, out *Response) error {
	return h.UploaderHandler.Create(ctx, in, out)
}

func init() {
	proto.RegisterFile("code.xxxxxx.com/micro/proto/uploader/uploader.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4f, 0xc3, 0x40,
	0x0c, 0x85, 0x09, 0x2d, 0x69, 0x6a, 0x18, 0xc0, 0x03, 0x3a, 0x21, 0x21, 0x95, 0x4c, 0x9d, 0x12,
	0x41, 0x47, 0xc4, 0xc4, 0x3f, 0x88, 0xe8, 0x5c, 0x5d, 0xaf, 0x1e, 0x2a, 0x92, 0xf8, 0x70, 0x1c,
	0x89, 0xfe, 0x7b, 0xd4, 0xbb, 0x6b, 0xc4, 0xf6, 0xde, 0xe7, 0xe1, 0xbd, 0x67, 0xd8, 0x38, 0x3e,
	0x50, 0x65, 0x9d, 0xaa, 0xe5, 0xca, 0x71, 0x57, 0x77, 0x47, 0x27, 0x5c, 0x7b, 0x61, 0xe5, 0x7a,
	0xf4, 0x2d, 0xdb, 0x03, 0xc9, 0x24, 0xaa, 0xc0, 0xb1, 0xb8, 0xf8, 0x52, 0x60, 0xd1, 0xd0, 0xcf,
	0x48, 0x83, 0xe2, 0x23, 0xe4, 0xfb, 0xd1, 0x7d, 0x93, 0x9a, 0x6c, 0x95, 0xad, 0x97, 0x4d, 0x72,
	0x88, 0x30, 0xef, 0x6d, 0x47, 0xe6, 0x3a, 0xd0, 0xa0, 0xf1, 0x05, 0xee, 0x1c, 0xf7, 0x4a, 0xbd,
	0xee, 0xf4, 0xe4, 0xc9, 0xcc, 0xc2, 0xed, 0x36, 0xb1, 0xaf, 0x93, 0x27, 0x34, 0xb0, 0xa0, 0x5f,
	0x7f, 0x14, 0x1a, 0xcc, 0x7c, 0x95, 0xad, 0x6f, 0x9a, 0x8b, 0x2d, 0xdf, 0xa1, 0x68, 0x68, 0xf0,
	0xdc, 0x0f, 0x84, 0xf7, 0x30, 0x1b, 0xa5, 0x4d, 0x89, 0x67, 0x89, 0xcf, 0x00, 0xb1, 0xdd, 0xee,
	0x7c, 0x88, 0xa1, 0xcb, 0x48, 0xb6, 0xd2, 0xbe, 0x7d, 0x40, 0xb1, 0x4d, 0xe5, 0xf1, 0x15, 0xf2,
	0x4f, 0x21, 0xab, 0x84, 0x0f, 0xd5, 0xb4, 0x30, 0xcd, 0x79, 0xc2, 0xff, 0x28, 0xa6, 0x95, 0x57,
	0xfb, 0x3c, 0x3c, 0x60, 0xf3, 0x17, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x0c, 0x92, 0xf0, 0x37, 0x01,
	0x00, 0x00,
}
