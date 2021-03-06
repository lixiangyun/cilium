// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/core/protocol.proto

package core

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/lyft/protoc-gen-validate/validate"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// [#not-implemented-hide:]
type TcpProtocolOptions struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpProtocolOptions) Reset()         { *m = TcpProtocolOptions{} }
func (m *TcpProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*TcpProtocolOptions) ProtoMessage()    {}
func (*TcpProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_d86476e078060b60, []int{0}
}

func (m *TcpProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProtocolOptions.Unmarshal(m, b)
}
func (m *TcpProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProtocolOptions.Marshal(b, m, deterministic)
}
func (m *TcpProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProtocolOptions.Merge(m, src)
}
func (m *TcpProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_TcpProtocolOptions.Size(m)
}
func (m *TcpProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProtocolOptions proto.InternalMessageInfo

type HttpProtocolOptions struct {
	// The idle timeout for upstream connection pool connections. The idle timeout is defined as the
	// period in which there are no active requests. If not set, there is no idle timeout. When the
	// idle timeout is reached the connection will be closed. Note that request based timeouts mean
	// that HTTP/2 PINGs will not keep the connection alive.
	IdleTimeout          *duration.Duration `protobuf:"bytes,1,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *HttpProtocolOptions) Reset()         { *m = HttpProtocolOptions{} }
func (m *HttpProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*HttpProtocolOptions) ProtoMessage()    {}
func (*HttpProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_d86476e078060b60, []int{1}
}

func (m *HttpProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpProtocolOptions.Unmarshal(m, b)
}
func (m *HttpProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpProtocolOptions.Marshal(b, m, deterministic)
}
func (m *HttpProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpProtocolOptions.Merge(m, src)
}
func (m *HttpProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_HttpProtocolOptions.Size(m)
}
func (m *HttpProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_HttpProtocolOptions proto.InternalMessageInfo

func (m *HttpProtocolOptions) GetIdleTimeout() *duration.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

type Http1ProtocolOptions struct {
	// Handle HTTP requests with absolute URLs in the requests. These requests
	// are generally sent by clients to forward/explicit proxies. This allows clients to configure
	// envoy as their HTTP proxy. In Unix, for example, this is typically done by setting the
	// *http_proxy* environment variable.
	AllowAbsoluteUrl *wrappers.BoolValue `protobuf:"bytes,1,opt,name=allow_absolute_url,json=allowAbsoluteUrl,proto3" json:"allow_absolute_url,omitempty"`
	// Handle incoming HTTP/1.0 and HTTP 0.9 requests.
	// This is off by default, and not fully standards compliant. There is support for pre-HTTP/1.1
	// style connect logic, dechunking, and handling lack of client host iff
	// *default_host_for_http_10* is configured.
	AcceptHttp_10 bool `protobuf:"varint,2,opt,name=accept_http_10,json=acceptHttp10,proto3" json:"accept_http_10,omitempty"`
	// A default host for HTTP/1.0 requests. This is highly suggested if *accept_http_10* is true as
	// Envoy does not otherwise support HTTP/1.0 without a Host header.
	// This is a no-op if *accept_http_10* is not true.
	DefaultHostForHttp_10 string   `protobuf:"bytes,3,opt,name=default_host_for_http_10,json=defaultHostForHttp10,proto3" json:"default_host_for_http_10,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Http1ProtocolOptions) Reset()         { *m = Http1ProtocolOptions{} }
func (m *Http1ProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*Http1ProtocolOptions) ProtoMessage()    {}
func (*Http1ProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_d86476e078060b60, []int{2}
}

func (m *Http1ProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Http1ProtocolOptions.Unmarshal(m, b)
}
func (m *Http1ProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Http1ProtocolOptions.Marshal(b, m, deterministic)
}
func (m *Http1ProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Http1ProtocolOptions.Merge(m, src)
}
func (m *Http1ProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_Http1ProtocolOptions.Size(m)
}
func (m *Http1ProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_Http1ProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_Http1ProtocolOptions proto.InternalMessageInfo

func (m *Http1ProtocolOptions) GetAllowAbsoluteUrl() *wrappers.BoolValue {
	if m != nil {
		return m.AllowAbsoluteUrl
	}
	return nil
}

func (m *Http1ProtocolOptions) GetAcceptHttp_10() bool {
	if m != nil {
		return m.AcceptHttp_10
	}
	return false
}

func (m *Http1ProtocolOptions) GetDefaultHostForHttp_10() string {
	if m != nil {
		return m.DefaultHostForHttp_10
	}
	return ""
}

type Http2ProtocolOptions struct {
	// `Maximum table size <http://httpwg.org/specs/rfc7541.html#rfc.section.4.2>`_
	// (in octets) that the encoder is permitted to use for the dynamic HPACK table. Valid values
	// range from 0 to 4294967295 (2^32 - 1) and defaults to 4096. 0 effectively disables header
	// compression.
	HpackTableSize *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=hpack_table_size,json=hpackTableSize,proto3" json:"hpack_table_size,omitempty"`
	// `Maximum concurrent streams <http://httpwg.org/specs/rfc7540.html#rfc.section.5.1.2>`_
	// allowed for peer on one HTTP/2 connection. Valid values range from 1 to 2147483647 (2^31 - 1)
	// and defaults to 2147483647.
	MaxConcurrentStreams *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=max_concurrent_streams,json=maxConcurrentStreams,proto3" json:"max_concurrent_streams,omitempty"`
	// This field also acts as a soft limit on the number of bytes Envoy will buffer per-stream in the
	// HTTP/2 codec buffers. Once the buffer reaches this pointer, watermark callbacks will fire to
	// stop the flow of data to the codec buffers.
	InitialStreamWindowSize *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=initial_stream_window_size,json=initialStreamWindowSize,proto3" json:"initial_stream_window_size,omitempty"`
	// Similar to *initial_stream_window_size*, but for connection-level flow-control
	// window. Currently, this has the same minimum/maximum/default as *initial_stream_window_size*.
	InitialConnectionWindowSize *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=initial_connection_window_size,json=initialConnectionWindowSize,proto3" json:"initial_connection_window_size,omitempty"`
	// [#not-implemented-hide:] Hiding until nghttp2 has native support.
	//
	// Allows proxying Websocket and other upgrades over H2 connect.
	//
	// THIS IS NOT SAFE TO USE IN PRODUCTION
	//
	// This currently works via disabling all HTTP sanity checks for H2 traffic
	// which is a much larger hammer than we'd like to use. Eventually when
	// https://github.com/nghttp2/nghttp2/issues/1181 is resolved, this will work
	// with simply enabling CONNECT for H2. This may require some tweaks to the
	// headers making pre-CONNECT-support proxying not backwards compatible with
	// post-CONNECT-support proxying.
	AllowConnect         bool     `protobuf:"varint,5,opt,name=allow_connect,json=allowConnect,proto3" json:"allow_connect,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Http2ProtocolOptions) Reset()         { *m = Http2ProtocolOptions{} }
func (m *Http2ProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*Http2ProtocolOptions) ProtoMessage()    {}
func (*Http2ProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_d86476e078060b60, []int{3}
}

func (m *Http2ProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Http2ProtocolOptions.Unmarshal(m, b)
}
func (m *Http2ProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Http2ProtocolOptions.Marshal(b, m, deterministic)
}
func (m *Http2ProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Http2ProtocolOptions.Merge(m, src)
}
func (m *Http2ProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_Http2ProtocolOptions.Size(m)
}
func (m *Http2ProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_Http2ProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_Http2ProtocolOptions proto.InternalMessageInfo

func (m *Http2ProtocolOptions) GetHpackTableSize() *wrappers.UInt32Value {
	if m != nil {
		return m.HpackTableSize
	}
	return nil
}

func (m *Http2ProtocolOptions) GetMaxConcurrentStreams() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConcurrentStreams
	}
	return nil
}

func (m *Http2ProtocolOptions) GetInitialStreamWindowSize() *wrappers.UInt32Value {
	if m != nil {
		return m.InitialStreamWindowSize
	}
	return nil
}

func (m *Http2ProtocolOptions) GetInitialConnectionWindowSize() *wrappers.UInt32Value {
	if m != nil {
		return m.InitialConnectionWindowSize
	}
	return nil
}

func (m *Http2ProtocolOptions) GetAllowConnect() bool {
	if m != nil {
		return m.AllowConnect
	}
	return false
}

// [#not-implemented-hide:]
type GrpcProtocolOptions struct {
	Http2ProtocolOptions *Http2ProtocolOptions `protobuf:"bytes,1,opt,name=http2_protocol_options,json=http2ProtocolOptions,proto3" json:"http2_protocol_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GrpcProtocolOptions) Reset()         { *m = GrpcProtocolOptions{} }
func (m *GrpcProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*GrpcProtocolOptions) ProtoMessage()    {}
func (*GrpcProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_d86476e078060b60, []int{4}
}

func (m *GrpcProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcProtocolOptions.Unmarshal(m, b)
}
func (m *GrpcProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcProtocolOptions.Marshal(b, m, deterministic)
}
func (m *GrpcProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcProtocolOptions.Merge(m, src)
}
func (m *GrpcProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_GrpcProtocolOptions.Size(m)
}
func (m *GrpcProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcProtocolOptions proto.InternalMessageInfo

func (m *GrpcProtocolOptions) GetHttp2ProtocolOptions() *Http2ProtocolOptions {
	if m != nil {
		return m.Http2ProtocolOptions
	}
	return nil
}

func init() {
	proto.RegisterType((*TcpProtocolOptions)(nil), "envoy.api.v2.core.TcpProtocolOptions")
	proto.RegisterType((*HttpProtocolOptions)(nil), "envoy.api.v2.core.HttpProtocolOptions")
	proto.RegisterType((*Http1ProtocolOptions)(nil), "envoy.api.v2.core.Http1ProtocolOptions")
	proto.RegisterType((*Http2ProtocolOptions)(nil), "envoy.api.v2.core.Http2ProtocolOptions")
	proto.RegisterType((*GrpcProtocolOptions)(nil), "envoy.api.v2.core.GrpcProtocolOptions")
}

func init() { proto.RegisterFile("envoy/api/v2/core/protocol.proto", fileDescriptor_d86476e078060b60) }

var fileDescriptor_d86476e078060b60 = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x65, 0x12, 0xa0, 0x6c, 0x43, 0x69, 0x5d, 0xab, 0x0d, 0x01, 0x85, 0x28, 0x20, 0x11,
	0xf5, 0x60, 0xb7, 0xae, 0xc4, 0x9d, 0x14, 0x95, 0x70, 0x02, 0xb9, 0x29, 0x88, 0x03, 0x5a, 0x6d,
	0x36, 0x9b, 0x64, 0xc5, 0xc6, 0x63, 0xad, 0xc7, 0x49, 0xe9, 0x93, 0xf0, 0x08, 0x3c, 0x03, 0x07,
	0xc4, 0x03, 0xf0, 0x0e, 0x48, 0xdc, 0x78, 0x0a, 0x23, 0x7b, 0x37, 0x11, 0x24, 0x95, 0x40, 0xdc,
	0x46, 0x33, 0xf3, 0xff, 0xdf, 0x8c, 0x3d, 0x4b, 0x5a, 0x22, 0x9e, 0xc1, 0x87, 0x80, 0x25, 0x32,
	0x98, 0x85, 0x01, 0x07, 0x2d, 0x82, 0x44, 0x03, 0x02, 0x07, 0xe5, 0x97, 0x81, 0xbb, 0x53, 0x76,
	0xf8, 0x2c, 0x91, 0xfe, 0x2c, 0xf4, 0x8b, 0x8e, 0x46, 0x73, 0x0c, 0x30, 0x56, 0xb6, 0x73, 0x90,
	0x8d, 0x82, 0x61, 0xa6, 0x19, 0x4a, 0x88, 0x8d, 0x64, 0xbd, 0x3e, 0xd7, 0x2c, 0x49, 0x84, 0x4e,
	0x6d, 0x7d, 0x7f, 0xc6, 0x94, 0x1c, 0x32, 0x14, 0xc1, 0x22, 0xb0, 0x05, 0x6f, 0x0c, 0x63, 0x28,
	0xc3, 0xa0, 0x88, 0x4c, 0xb6, 0xed, 0x11, 0xb7, 0xcf, 0x93, 0x57, 0x76, 0xac, 0x97, 0x49, 0x41,
	0x4a, 0xdb, 0x6f, 0xc9, 0x6e, 0x0f, 0x71, 0x35, 0xed, 0x76, 0x49, 0x4d, 0x0e, 0x95, 0xa0, 0x28,
	0xa7, 0x02, 0x32, 0xac, 0x3b, 0x2d, 0xa7, 0xb3, 0x19, 0xde, 0xf5, 0xcd, 0x48, 0xfe, 0x62, 0x24,
	0xff, 0x99, 0x1d, 0xb9, 0x5b, 0xfd, 0xf8, 0xfd, 0x81, 0x13, 0x6d, 0x16, 0xa2, 0xbe, 0xd1, 0xb4,
	0xbf, 0x38, 0xc4, 0x2b, 0xbc, 0x8f, 0x56, 0xcd, 0x7b, 0xc4, 0x65, 0x4a, 0xc1, 0x9c, 0xb2, 0x41,
	0x0a, 0x2a, 0x43, 0x41, 0x33, 0xad, 0x2c, 0xa2, 0xb1, 0x86, 0xe8, 0x02, 0xa8, 0xd7, 0x4c, 0x65,
	0x22, 0xda, 0x2e, 0x55, 0x4f, 0xad, 0xe8, 0x5c, 0x2b, 0xf7, 0x11, 0xd9, 0x62, 0x9c, 0x8b, 0x04,
	0xe9, 0x04, 0x31, 0xa1, 0x47, 0x87, 0xf5, 0x6b, 0x2d, 0xa7, 0xb3, 0x11, 0xd5, 0x4c, 0xb6, 0xa4,
	0x1f, 0xba, 0x4f, 0x48, 0x7d, 0x28, 0x46, 0x2c, 0x53, 0x48, 0x27, 0x90, 0x22, 0x1d, 0x81, 0x5e,
	0xf6, 0x57, 0x5a, 0x4e, 0xe7, 0x56, 0xe4, 0xd9, 0x7a, 0x0f, 0x52, 0x3c, 0x05, 0x6d, 0x74, 0xed,
	0x6f, 0x15, 0xb3, 0x40, 0xb8, 0xba, 0xc0, 0x29, 0xd9, 0x9e, 0x24, 0x8c, 0xbf, 0xa7, 0xc8, 0x06,
	0x4a, 0xd0, 0x54, 0x5e, 0x0a, 0x3b, 0xfe, 0xfd, 0xb5, 0xf1, 0xcf, 0x5f, 0xc4, 0x78, 0x1c, 0x9a,
	0x05, 0xb6, 0x4a, 0x55, 0xbf, 0x10, 0x9d, 0xc9, 0x4b, 0xe1, 0x72, 0xb2, 0x37, 0x65, 0x17, 0x94,
	0x43, 0xcc, 0x33, 0xad, 0x45, 0x8c, 0x34, 0x45, 0x2d, 0xd8, 0x34, 0x2d, 0xd7, 0xf8, 0x8b, 0x5b,
	0xf7, 0xce, 0xe7, 0x9f, 0x5f, 0x2b, 0xe4, 0x60, 0xa3, 0x9e, 0xe7, 0x79, 0x7e, 0xb3, 0xe3, 0x44,
	0xde, 0x94, 0x5d, 0x9c, 0x2c, 0xbd, 0xce, 0x8c, 0x95, 0xab, 0x48, 0x43, 0xc6, 0x12, 0x25, 0x53,
	0xd6, 0x9d, 0xce, 0x65, 0x3c, 0x84, 0xb9, 0x19, 0xbb, 0xf2, 0x0f, 0xa0, 0x9d, 0x02, 0x54, 0x3b,
	0x20, 0x16, 0x94, 0xe7, 0x95, 0x68, 0xdf, 0x5a, 0x1a, 0xc8, 0x9b, 0xd2, 0xb0, 0x5c, 0x09, 0x49,
	0x73, 0x41, 0xe3, 0x10, 0xc7, 0x82, 0x17, 0x5f, 0xec, 0x0f, 0x62, 0xf5, 0xff, 0x88, 0xf7, 0xac,
	0xed, 0xc9, 0xd2, 0xf5, 0x37, 0xea, 0x43, 0x72, 0xdb, 0x5c, 0x94, 0x65, 0xd6, 0xaf, 0xdb, 0x33,
	0x28, 0x92, 0x56, 0xd1, 0x46, 0xb2, 0xfb, 0x5c, 0x27, 0x7c, 0xf5, 0x67, 0xbe, 0x23, 0x7b, 0xc5,
	0x31, 0x84, 0x74, 0xf1, 0x62, 0x29, 0x98, 0x8a, 0xfd, 0xa5, 0x8f, 0xfd, 0xb5, 0xa7, 0xeb, 0x5f,
	0x75, 0x15, 0x91, 0x37, 0xb9, 0x22, 0xdb, 0xad, 0x7e, 0xfa, 0xd1, 0x74, 0x06, 0x37, 0x4a, 0xf3,
	0xe3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x48, 0x26, 0xf8, 0x27, 0x29, 0x04, 0x00, 0x00,
}
