// Code generated by protoc-gen-go.
// source: google/type/date.proto
// DO NOT EDIT!

package google_type

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Represents a whole calendar date, e.g. date of birth. The time of day and
// time zone are either specified elsewhere or are not significant. The date
// is relative to the Proleptic Gregorian Calendar. The day may be 0 to
// represent a year and month where the day is not significant, e.g. credit card
// expiration date. The year may be 0 to represent a month and day independent
// of year, e.g. anniversary date. Related types are [google.type.TimeOfDay][google.type.TimeOfDay]
// and [google.protobuf.Timestamp][google.protobuf.Timestamp].
type Date struct {
	// Year of date. Must be from 1 to 9,999, or 0 if specifying a date without
	// a year.
	Year int32 `protobuf:"varint,1,opt,name=year" json:"year,omitempty"`
	// Month of year of date. Must be from 1 to 12.
	Month int32 `protobuf:"varint,2,opt,name=month" json:"month,omitempty"`
	// Day of month. Must be from 1 to 31 and valid for the year and month, or 0
	// if specifying a year/month where the day is not sigificant.
	Day int32 `protobuf:"varint,3,opt,name=day" json:"day,omitempty"`
}

func (m *Date) Reset()                    { *m = Date{} }
func (m *Date) String() string            { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()               {}
func (*Date) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterType((*Date)(nil), "google.type.Date")
}

func init() { proto.RegisterFile("google/type/date.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x4f, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x86, 0x88, 0xeb, 0x81, 0xc4, 0x95, 0x9c, 0xb8, 0x58, 0x5c, 0x80, 0x52,
	0x42, 0x42, 0x5c, 0x2c, 0x95, 0xa9, 0x89, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x60,
	0xb6, 0x90, 0x08, 0x17, 0x6b, 0x6e, 0x7e, 0x5e, 0x49, 0x86, 0x04, 0x13, 0x58, 0x10, 0xc2, 0x11,
	0x12, 0xe0, 0x62, 0x4e, 0x49, 0xac, 0x94, 0x60, 0x06, 0x8b, 0x81, 0x98, 0x4e, 0x8a, 0x5c, 0xfc,
	0xc9, 0xf9, 0xb9, 0x7a, 0x48, 0xc6, 0x3a, 0x71, 0x82, 0x0c, 0x0d, 0x00, 0x59, 0x17, 0xc0, 0xb8,
	0x80, 0x91, 0x31, 0x89, 0x0d, 0x6c, 0xb5, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x55, 0xc3, 0xd3,
	0x52, 0x94, 0x00, 0x00, 0x00,
}