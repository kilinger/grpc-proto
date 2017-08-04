// Code generated by protoc-gen-go. DO NOT EDIT.
// source: code.xxxxxx.com/micro/proto/sms/sms.proto

/*
Package sms is a generated protocol buffer package.

It is generated from these files:
	code.xxxxxx.com/micro/proto/sms/sms.proto

It has these top-level messages:
	Param
	Request
	Response
*/
package sms

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

type Param struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Param) Reset()                    { *m = Param{} }
func (m *Param) String() string            { return proto.CompactTextString(m) }
func (*Param) ProtoMessage()               {}
func (*Param) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Param) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Param) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Request struct {
	Channel     string   `protobuf:"bytes,1,opt,name=channel" json:"channel,omitempty"`
	RecNum      string   `protobuf:"bytes,2,opt,name=rec_num,json=recNum" json:"rec_num,omitempty"`
	ParamString string   `protobuf:"bytes,3,opt,name=param_string,json=paramString" json:"param_string,omitempty"`
	Params      []*Param `protobuf:"bytes,4,rep,name=params" json:"params,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *Request) GetRecNum() string {
	if m != nil {
		return m.RecNum
	}
	return ""
}

func (m *Request) GetParamString() string {
	if m != nil {
		return m.ParamString
	}
	return ""
}

func (m *Request) GetParams() []*Param {
	if m != nil {
		return m.Params
	}
	return nil
}

type Response struct {
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*Param)(nil), "sms.Param")
	proto.RegisterType((*Request)(nil), "sms.Request")
	proto.RegisterType((*Response)(nil), "sms.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Sms service

type SmsClient interface {
	Send(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type smsClient struct {
	c           client.Client
	serviceName string
}

func NewSmsClient(serviceName string, c client.Client) SmsClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "sms"
	}
	return &smsClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *smsClient) Send(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Sms.Send", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sms service

type SmsHandler interface {
	Send(context.Context, *Request, *Response) error
}

func RegisterSmsHandler(s server.Server, hdlr SmsHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Sms{hdlr}, opts...))
}

type Sms struct {
	SmsHandler
}

func (h *Sms) Send(ctx context.Context, in *Request, out *Response) error {
	return h.SmsHandler.Send(ctx, in, out)
}

func init() { proto.RegisterFile("code.xxxxxx.com/micro/proto/sms/sms.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0xd1, 0x4a, 0x03, 0x31,
	0x10, 0x45, 0x5d, 0xb7, 0xdd, 0xd5, 0x69, 0x05, 0x09, 0x82, 0x41, 0x7c, 0xa8, 0x01, 0xa1, 0x82,
	0xec, 0x42, 0xfd, 0x0f, 0x91, 0xdd, 0x0f, 0x28, 0x31, 0x1d, 0xb4, 0xd8, 0x24, 0x6b, 0x26, 0x11,
	0x7c, 0xf6, 0xc7, 0x25, 0xd3, 0xf4, 0x21, 0x30, 0xe7, 0x0e, 0x37, 0x39, 0x04, 0x9e, 0x8c, 0xdf,
	0x61, 0xa7, 0x4d, 0x8c, 0xda, 0x77, 0xc6, 0xdb, 0xde, 0xee, 0x4d, 0xf0, 0xfd, 0x14, 0x7c, 0xf4,
	0x3d, 0x59, 0xca, 0xa7, 0x63, 0x12, 0x35, 0x59, 0x52, 0x3d, 0xcc, 0xdf, 0x74, 0xd0, 0x56, 0x5c,
	0x43, 0xfd, 0x85, 0xbf, 0xb2, 0x5a, 0x55, 0xeb, 0xcb, 0x21, 0x8f, 0xe2, 0x06, 0xe6, 0x3f, 0xfa,
	0x90, 0x50, 0x9e, 0x73, 0x76, 0x04, 0xf5, 0x57, 0x41, 0x3b, 0xe0, 0x77, 0x42, 0x8a, 0x42, 0x42,
	0x6b, 0x3e, 0xb5, 0x73, 0x78, 0x28, 0xbd, 0x13, 0x8a, 0x5b, 0x68, 0x03, 0x9a, 0xad, 0x4b, 0xb6,
	0xb4, 0x9b, 0x80, 0xe6, 0x35, 0x59, 0xf1, 0x00, 0xcb, 0x29, 0xbf, 0xb7, 0xa5, 0x18, 0xf6, 0xee,
	0x43, 0xd6, 0xbc, 0x5d, 0x70, 0x36, 0x72, 0x24, 0x14, 0x34, 0x8c, 0x24, 0x67, 0xab, 0x7a, 0xbd,
	0xd8, 0x40, 0x97, 0x9d, 0xd9, 0x72, 0x28, 0x1b, 0x75, 0x0f, 0x17, 0x03, 0xd2, 0xe4, 0x1d, 0x61,
	0x36, 0x4f, 0xe1, 0x64, 0x90, 0xc7, 0xcd, 0x33, 0xd4, 0xa3, 0x25, 0xf1, 0x08, 0xb3, 0x11, 0xdd,
	0x4e, 0x2c, 0xf9, 0x82, 0x22, 0x7d, 0x77, 0x55, 0xe8, 0xd8, 0x56, 0x67, 0xef, 0x0d, 0x7f, 0xc7,
	0xcb, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x0b, 0x91, 0x02, 0x3b, 0x01, 0x00, 0x00,
}