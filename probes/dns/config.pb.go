// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/probes/dns/config.proto

/*
Package dns is a generated protocol buffer package.

It is generated from these files:
	github.com/google/cloudprober/probes/dns/config.proto

It has these top-level messages:
	ProbeConf
*/
package dns

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProbeConf struct {
	// Domain to use when making DNS queries
	ResolvedDomain *string `protobuf:"bytes,1,opt,name=resolved_domain,json=resolvedDomain,def=www.google.com." json:"resolved_domain,omitempty"`
	// Export stats after these many milliseconds
	StatsExportIntervalMsec *int32 `protobuf:"varint,2,opt,name=stats_export_interval_msec,json=statsExportIntervalMsec,def=10000" json:"stats_export_interval_msec,omitempty"`
	XXX_unrecognized        []byte `json:"-"`
}

func (m *ProbeConf) Reset()                    { *m = ProbeConf{} }
func (m *ProbeConf) String() string            { return proto.CompactTextString(m) }
func (*ProbeConf) ProtoMessage()               {}
func (*ProbeConf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_ProbeConf_ResolvedDomain string = "www.google.com."
const Default_ProbeConf_StatsExportIntervalMsec int32 = 10000

func (m *ProbeConf) GetResolvedDomain() string {
	if m != nil && m.ResolvedDomain != nil {
		return *m.ResolvedDomain
	}
	return Default_ProbeConf_ResolvedDomain
}

func (m *ProbeConf) GetStatsExportIntervalMsec() int32 {
	if m != nil && m.StatsExportIntervalMsec != nil {
		return *m.StatsExportIntervalMsec
	}
	return Default_ProbeConf_StatsExportIntervalMsec
}

func init() {
	proto.RegisterType((*ProbeConf)(nil), "cloudprober.probes.dns.ProbeConf")
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/probes/dns/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0xce,
	0xc9, 0x2f, 0x4d, 0x29, 0x28, 0xca, 0x4f, 0x4a, 0x2d, 0xd2, 0x07, 0x53, 0xc5, 0xfa, 0x29, 0x79,
	0xc5, 0xfa, 0xc9, 0xf9, 0x79, 0x69, 0x99, 0xe9, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x62,
	0x48, 0x8a, 0xf4, 0x20, 0x8a, 0xf4, 0x52, 0xf2, 0x8a, 0x95, 0x3a, 0x19, 0xb9, 0x38, 0x03, 0x40,
	0x5c, 0xe7, 0xfc, 0xbc, 0x34, 0x21, 0x0b, 0x2e, 0xfe, 0xa2, 0xd4, 0xe2, 0xfc, 0x9c, 0xb2, 0xd4,
	0x94, 0xf8, 0x94, 0xfc, 0xdc, 0xc4, 0xcc, 0x3c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x4e, 0x2b, 0xfe,
	0xf2, 0xf2, 0x72, 0x3d, 0x88, 0x75, 0x20, 0x9b, 0xf5, 0x82, 0xf8, 0x60, 0xea, 0x5c, 0xc0, 0xca,
	0x84, 0x9c, 0xb8, 0xa4, 0x8a, 0x4b, 0x12, 0x4b, 0x8a, 0xe3, 0x53, 0x2b, 0x0a, 0xf2, 0x8b, 0x4a,
	0xe2, 0x33, 0xf3, 0x4a, 0x52, 0x8b, 0xca, 0x12, 0x73, 0xe2, 0x73, 0x8b, 0x53, 0x93, 0x25, 0x98,
	0x14, 0x18, 0x35, 0x58, 0xad, 0x58, 0x0d, 0x0d, 0x0c, 0x0c, 0x0c, 0x82, 0xc4, 0xc1, 0x0a, 0x5d,
	0xc1, 0xea, 0x3c, 0xa1, 0xca, 0x7c, 0x8b, 0x53, 0x93, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2c,
	0x3e, 0x2f, 0xb9, 0xdb, 0x00, 0x00, 0x00,
}