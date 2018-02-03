// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/datatype.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CrdRecipient_ResourceScope int32

const (
	CrdRecipient_Cluster    CrdRecipient_ResourceScope = 0
	CrdRecipient_Namespaced CrdRecipient_ResourceScope = 1
)

var CrdRecipient_ResourceScope_name = map[int32]string{
	0: "Cluster",
	1: "Namespaced",
}
var CrdRecipient_ResourceScope_value = map[string]int32{
	"Cluster":    0,
	"Namespaced": 1,
}

func (x CrdRecipient_ResourceScope) String() string {
	return proto.EnumName(CrdRecipient_ResourceScope_name, int32(x))
}
func (CrdRecipient_ResourceScope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorDatatype, []int{0, 0}
}

type CrdRecipient struct {
	Group         string                     `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Version       string                     `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Scope         string                     `protobuf:"bytes,3,opt,name=scope,proto3" json:"scope,omitempty"`
	Plural        string                     `protobuf:"bytes,4,opt,name=plural,proto3" json:"plural,omitempty"`
	Singular      string                     `protobuf:"bytes,5,opt,name=singular,proto3" json:"singular,omitempty"`
	Kind          string                     `protobuf:"bytes,6,opt,name=kind,proto3" json:"kind,omitempty"`
	ResourceScope CrdRecipient_ResourceScope `protobuf:"varint,7,opt,name=resource_scope,json=resourceScope,proto3,enum=pb.CrdRecipient_ResourceScope" json:"resource_scope,omitempty"`
}

func (m *CrdRecipient) Reset()                    { *m = CrdRecipient{} }
func (m *CrdRecipient) String() string            { return proto.CompactTextString(m) }
func (*CrdRecipient) ProtoMessage()               {}
func (*CrdRecipient) Descriptor() ([]byte, []int) { return fileDescriptorDatatype, []int{0} }

func (m *CrdRecipient) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *CrdRecipient) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *CrdRecipient) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *CrdRecipient) GetPlural() string {
	if m != nil {
		return m.Plural
	}
	return ""
}

func (m *CrdRecipient) GetSingular() string {
	if m != nil {
		return m.Singular
	}
	return ""
}

func (m *CrdRecipient) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *CrdRecipient) GetResourceScope() CrdRecipient_ResourceScope {
	if m != nil {
		return m.ResourceScope
	}
	return CrdRecipient_Cluster
}

func init() {
	proto.RegisterType((*CrdRecipient)(nil), "pb.CrdRecipient")
	proto.RegisterEnum("pb.CrdRecipient_ResourceScope", CrdRecipient_ResourceScope_name, CrdRecipient_ResourceScope_value)
}

func init() { proto.RegisterFile("pb/datatype.proto", fileDescriptorDatatype) }

var fileDescriptorDatatype = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0x6c, 0x13, 0x1d, 0x6d, 0xa8, 0x83, 0xc8, 0xe2, 0x41, 0x4a, 0x4f, 0x3d, 0x48,
	0x04, 0x7d, 0x84, 0xe2, 0xd5, 0x43, 0x7c, 0x00, 0xd9, 0x24, 0x43, 0x59, 0x5c, 0x77, 0x87, 0xd9,
	0x5d, 0xc1, 0xc7, 0xf0, 0x8d, 0xa5, 0x9b, 0x2a, 0xf5, 0x36, 0xdf, 0xff, 0x7f, 0x30, 0xc3, 0xc0,
	0x15, 0xf7, 0x0f, 0xa3, 0x8e, 0x3a, 0x7e, 0x31, 0xb5, 0x2c, 0x3e, 0x7a, 0x2c, 0xb9, 0x5f, 0x7f,
	0x97, 0x70, 0xb9, 0x95, 0xb1, 0xa3, 0xc1, 0xb0, 0x21, 0x17, 0xf1, 0x1a, 0xe6, 0x3b, 0xf1, 0x89,
	0x55, 0xb1, 0x2a, 0x36, 0xe7, 0xdd, 0x04, 0xa8, 0xa0, 0xfe, 0x24, 0x09, 0xc6, 0x3b, 0x55, 0xe6,
	0xfc, 0x17, 0xf7, 0x7e, 0x18, 0x3c, 0x93, 0x3a, 0x9d, 0xfc, 0x0c, 0x78, 0x03, 0x15, 0xdb, 0x24,
	0xda, 0xaa, 0x59, 0x8e, 0x0f, 0x84, 0xb7, 0x70, 0x16, 0x8c, 0xdb, 0x25, 0xab, 0x45, 0xcd, 0x73,
	0xf3, 0xc7, 0x88, 0x30, 0x7b, 0x37, 0x6e, 0x54, 0x55, 0xce, 0xf3, 0x8c, 0xcf, 0xd0, 0x08, 0x05,
	0x9f, 0x64, 0xa0, 0xb7, 0x69, 0x4d, 0xbd, 0x2a, 0x36, 0xcd, 0xe3, 0x5d, 0xcb, 0x7d, 0x7b, 0x7c,
	0x77, 0xdb, 0x1d, 0xb4, 0xd7, 0xbd, 0xd5, 0x2d, 0xe4, 0x18, 0xd7, 0xf7, 0xb0, 0xf8, 0xd7, 0xe3,
	0x05, 0xd4, 0x5b, 0x9b, 0x42, 0x24, 0x59, 0x9e, 0x60, 0x03, 0xf0, 0xa2, 0x3f, 0x28, 0xb0, 0x1e,
	0x68, 0x5c, 0x16, 0x7d, 0x95, 0xdf, 0xf3, 0xf4, 0x13, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xfe, 0x25,
	0x21, 0x33, 0x01, 0x00, 0x00,
}
