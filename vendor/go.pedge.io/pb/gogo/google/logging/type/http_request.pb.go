// Code generated by protoc-gen-gogo.
// source: google/logging/type/http_request.proto
// DO NOT EDIT!

package google_logging_type

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// A common proto for logging HTTP requests.
//
type HttpRequest struct {
	// The request method. Examples: `"GET"`, `"HEAD"`, `"PUT"`, `"POST"`.
	RequestMethod string `protobuf:"bytes,1,opt,name=request_method,json=requestMethod,proto3" json:"request_method,omitempty"`
	// The scheme (http, https), the host name, the path and the query
	// portion of the URL that was requested.
	// Example: `"http://example.com/some/info?color=red"`.
	RequestUrl string `protobuf:"bytes,2,opt,name=request_url,json=requestUrl,proto3" json:"request_url,omitempty"`
	// The size of the HTTP request message in bytes, including the request
	// headers and the request body.
	RequestSize int64 `protobuf:"varint,3,opt,name=request_size,json=requestSize,proto3" json:"request_size,omitempty"`
	// The response code indicating the status of response.
	// Examples: 200, 404.
	Status int32 `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	// The size of the HTTP response message sent back to the client, in bytes,
	// including the response headers and the response body.
	ResponseSize int64 `protobuf:"varint,5,opt,name=response_size,json=responseSize,proto3" json:"response_size,omitempty"`
	// The user agent sent by the client. Example:
	// `"Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; Q312461; .NET CLR 1.0.3705)"`.
	UserAgent string `protobuf:"bytes,6,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	// The IP address (IPv4 or IPv6) of the client that issued the HTTP
	// request. Examples: `"192.168.1.1"`, `"FE80::0202:B3FF:FE1E:8329"`.
	RemoteIp string `protobuf:"bytes,7,opt,name=remote_ip,json=remoteIp,proto3" json:"remote_ip,omitempty"`
	// The referer URL of the request, as defined in
	// [HTTP/1.1 Header Field Definitions](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html).
	Referer string `protobuf:"bytes,8,opt,name=referer,proto3" json:"referer,omitempty"`
	// Whether or not an entity was served from cache
	// (with or without validation).
	CacheHit bool `protobuf:"varint,9,opt,name=cache_hit,json=cacheHit,proto3" json:"cache_hit,omitempty"`
	// Whether or not the response was validated with the origin server before
	// being served from cache. This field is only meaningful if `cache_hit` is
	// True.
	ValidatedWithOriginServer bool `protobuf:"varint,10,opt,name=validated_with_origin_server,json=validatedWithOriginServer,proto3" json:"validated_with_origin_server,omitempty"`
}

func (m *HttpRequest) Reset()                    { *m = HttpRequest{} }
func (m *HttpRequest) String() string            { return proto.CompactTextString(m) }
func (*HttpRequest) ProtoMessage()               {}
func (*HttpRequest) Descriptor() ([]byte, []int) { return fileDescriptorHttpRequest, []int{0} }

func init() {
	proto.RegisterType((*HttpRequest)(nil), "google.logging.type.HttpRequest")
}

var fileDescriptorHttpRequest = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x4e, 0xeb, 0x30,
	0x10, 0x46, 0xd5, 0xf6, 0xb6, 0x4d, 0xa6, 0xbd, 0x08, 0x19, 0x09, 0x0c, 0x14, 0xf1, 0x27, 0x10,
	0x1b, 0x92, 0x05, 0x0f, 0x80, 0xe8, 0xaa, 0x2c, 0x10, 0x55, 0x2a, 0xc4, 0xd2, 0x0a, 0xed, 0x90,
	0x58, 0x4a, 0xe3, 0x60, 0x4f, 0x8a, 0xe0, 0x8d, 0x79, 0x0b, 0x1c, 0x27, 0x41, 0x2c, 0x58, 0xfa,
	0xcc, 0xf9, 0x66, 0x46, 0x1e, 0xb8, 0x4c, 0x94, 0x4a, 0x32, 0x0c, 0x33, 0x95, 0x24, 0x32, 0x4f,
	0x42, 0xfa, 0x28, 0x30, 0x4c, 0x89, 0x0a, 0xa1, 0xf1, 0xad, 0x44, 0x43, 0x41, 0xa1, 0x15, 0x29,
	0xb6, 0x53, 0x7b, 0x41, 0xe3, 0x05, 0x95, 0x77, 0x30, 0x69, 0xc2, 0x71, 0x21, 0xc3, 0x38, 0xcf,
	0x15, 0xc5, 0x24, 0x55, 0x6e, 0xea, 0xc8, 0xd9, 0x57, 0x17, 0x46, 0x33, 0xdb, 0x29, 0xaa, 0x1b,
	0xb1, 0x0b, 0xd8, 0x6a, 0x7a, 0x8a, 0x35, 0x52, 0xaa, 0x56, 0xbc, 0x73, 0xd2, 0xb9, 0xf2, 0xa3,
	0xff, 0x0d, 0x7d, 0x70, 0x90, 0x1d, 0xc3, 0xa8, 0xd5, 0x4a, 0x9d, 0xf1, 0xae, 0x73, 0xa0, 0x41,
	0x4f, 0x3a, 0x63, 0xa7, 0x30, 0x6e, 0x05, 0x23, 0x3f, 0x91, 0xf7, 0xac, 0xd1, 0x8b, 0xda, 0xd0,
	0xc2, 0x22, 0xb6, 0x0b, 0x03, 0x63, 0x97, 0x29, 0x0d, 0xff, 0x67, 0x8b, 0xfd, 0xa8, 0x79, 0xb1,
	0x73, 0xb0, 0xc3, 0x4c, 0x61, 0x77, 0xc4, 0x3a, 0xdb, 0x77, 0xd9, 0x71, 0x0b, 0x5d, 0xf8, 0x08,
	0xa0, 0x34, 0xa8, 0x45, 0x9c, 0x60, 0x4e, 0x7c, 0xe0, 0xe6, 0xfb, 0x15, 0xb9, 0xab, 0x00, 0x3b,
	0x04, 0x5f, 0xe3, 0x5a, 0x11, 0x0a, 0x59, 0xf0, 0xa1, 0xab, 0x7a, 0x35, 0xb8, 0x2f, 0x18, 0x87,
	0xa1, 0xc6, 0x57, 0xd4, 0xa8, 0xb9, 0xe7, 0x4a, 0xed, 0xb3, 0x8a, 0x2d, 0xe3, 0x65, 0x8a, 0x22,
	0x95, 0xc4, 0x7d, 0x5b, 0xf3, 0x22, 0xcf, 0x81, 0x99, 0x24, 0x76, 0x0b, 0x93, 0x4d, 0x9c, 0xc9,
	0x55, 0x4c, 0xb8, 0x12, 0xef, 0x92, 0x52, 0xa1, 0xb4, 0xb4, 0xff, 0x2c, 0xec, 0xd4, 0x8d, 0xed,
	0x05, 0xce, 0xdf, 0xff, 0x71, 0x9e, 0xad, 0xf2, 0xe8, 0x8c, 0x85, 0x13, 0xa6, 0xd7, 0xb0, 0xb7,
	0x54, 0xeb, 0xe0, 0x8f, 0x23, 0x4d, 0xb7, 0x7f, 0xdd, 0x60, 0x5e, 0x1d, 0x66, 0xde, 0x79, 0x19,
	0xb8, 0x0b, 0xdd, 0x7c, 0x07, 0x00, 0x00, 0xff, 0xff, 0x05, 0xb2, 0x0b, 0xf0, 0xfe, 0x01, 0x00,
	0x00,
}
