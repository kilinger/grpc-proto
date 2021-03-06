// Code generated by protoc-gen-go. DO NOT EDIT.
// source: code.xxxxxx.com/micro/proto/qrcode/qrcode.proto

/*
Package qrcode is a generated protocol buffer package.

It is generated from these files:
	code.xxxxxx.com/micro/proto/qrcode/qrcode.proto

It has these top-level messages:
	Request
	Response
*/
package qrcode

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
	Content string `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
	Size    int32  `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Request) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
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

func init() {
	proto.RegisterType((*Request)(nil), "qrcode.Request")
	proto.RegisterType((*Response)(nil), "qrcode.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for QRcode service

type QRcodeClient interface {
	Generate(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type qRcodeClient struct {
	c           client.Client
	serviceName string
}

func NewQRcodeClient(serviceName string, c client.Client) QRcodeClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "qrcode"
	}
	return &qRcodeClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *qRcodeClient) Generate(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "QRcode.Generate", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for QRcode service

type QRcodeHandler interface {
	Generate(context.Context, *Request, *Response) error
}

func RegisterQRcodeHandler(s server.Server, hdlr QRcodeHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&QRcode{hdlr}, opts...))
}

type QRcode struct {
	QRcodeHandler
}

func (h *QRcode) Generate(ctx context.Context, in *Request, out *Response) error {
	return h.QRcodeHandler.Generate(ctx, in, out)
}

func init() { proto.RegisterFile("code.xxxxxx.com/micro/proto/qrcode/qrcode.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0xb1, 0x8e, 0xc2, 0x30,
	0x10, 0x44, 0x2f, 0x97, 0xbb, 0x24, 0xb7, 0xcd, 0x45, 0xae, 0x2c, 0x44, 0x11, 0xa5, 0x4a, 0x65,
	0x4b, 0x50, 0xf1, 0x05, 0x14, 0x54, 0xf8, 0x0f, 0x8c, 0xd9, 0x22, 0x12, 0xf1, 0x26, 0xb6, 0xd3,
	0xf0, 0xf5, 0xc8, 0x8e, 0x11, 0xd5, 0xce, 0xcc, 0xae, 0x66, 0x1f, 0x48, 0x43, 0x77, 0x14, 0xda,
	0x84, 0xa0, 0x49, 0x18, 0x9a, 0xe4, 0x34, 0x1a, 0x47, 0x72, 0x76, 0x14, 0x48, 0x2e, 0x2e, 0x6e,
	0xf3, 0x10, 0x29, 0x63, 0xd5, 0xe6, 0xfa, 0x0b, 0xd4, 0x0a, 0x97, 0x15, 0x7d, 0x60, 0x1c, 0x6a,
	0x43, 0x36, 0xa0, 0x0d, 0xbc, 0xe8, 0x8a, 0xe1, 0x4f, 0xbd, 0x2d, 0x63, 0xf0, 0xe3, 0xc7, 0x27,
	0xf2, 0xef, 0xae, 0x18, 0x7e, 0x55, 0xd2, 0x31, 0xb3, 0x7a, 0x42, 0x5e, 0xa6, 0xd3, 0xa4, 0xfb,
	0x3d, 0x34, 0x0a, 0xfd, 0x4c, 0xd6, 0x23, 0x6b, 0xa1, 0x5c, 0xdd, 0x23, 0x37, 0x45, 0x79, 0x38,
	0x41, 0x75, 0x55, 0xf1, 0x29, 0x93, 0xd0, 0x9c, 0xd1, 0xa2, 0xd3, 0x01, 0xd9, 0xbf, 0xc8, 0x5c,
	0x19, 0x63, 0xd7, 0x7e, 0x82, 0xad, 0xaa, 0xff, 0xba, 0x55, 0x09, 0xfa, 0xf8, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0xc0, 0xab, 0x53, 0x88, 0xe7, 0x00, 0x00, 0x00,
}
