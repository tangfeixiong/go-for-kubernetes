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
	Name     string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Group    string                     `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Version  string                     `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Scope    CrdRecipient_ResourceScope `protobuf:"varint,4,opt,name=scope,proto3,enum=pb.CrdRecipient_ResourceScope" json:"scope,omitempty"`
	Plural   string                     `protobuf:"bytes,5,opt,name=plural,proto3" json:"plural,omitempty"`
	Singular string                     `protobuf:"bytes,6,opt,name=singular,proto3" json:"singular,omitempty"`
	Kind     string                     `protobuf:"bytes,7,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (m *CrdRecipient) Reset()                    { *m = CrdRecipient{} }
func (m *CrdRecipient) String() string            { return proto.CompactTextString(m) }
func (*CrdRecipient) ProtoMessage()               {}
func (*CrdRecipient) Descriptor() ([]byte, []int) { return fileDescriptorDatatype, []int{0} }

func (m *CrdRecipient) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

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

func (m *CrdRecipient) GetScope() CrdRecipient_ResourceScope {
	if m != nil {
		return m.Scope
	}
	return CrdRecipient_Cluster
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

func init() {
	proto.RegisterType((*CrdRecipient)(nil), "pb.CrdRecipient")
	proto.RegisterEnum("pb.CrdRecipient_ResourceScope", CrdRecipient_ResourceScope_name, CrdRecipient_ResourceScope_value)
}

func init() { proto.RegisterFile("pb/datatype.proto", fileDescriptorDatatype) }

var fileDescriptorDatatype = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xcd, 0xba, 0x6d, 0x75, 0xd4, 0x65, 0x1d, 0x44, 0x82, 0x07, 0x59, 0xf6, 0xb4, 0x07,
	0xa9, 0xa0, 0xfe, 0x83, 0xbd, 0x7b, 0xa8, 0xbf, 0x20, 0x6d, 0x87, 0x25, 0xd8, 0x26, 0xc3, 0x24,
	0x11, 0xfc, 0xf3, 0x22, 0x4d, 0x55, 0xdc, 0xdb, 0xfb, 0x5e, 0x1e, 0x2f, 0x8f, 0x81, 0x6b, 0x6e,
	0x1f, 0x7b, 0x13, 0x4d, 0xfc, 0x64, 0xaa, 0x59, 0x7c, 0xf4, 0xb8, 0xe0, 0x76, 0xfb, 0xa5, 0xe0,
	0x72, 0x2f, 0x7d, 0x43, 0x9d, 0x65, 0x4b, 0x2e, 0x22, 0xc2, 0xd2, 0x99, 0x91, 0xb4, 0xda, 0xa8,
	0xdd, 0x79, 0x93, 0x35, 0xde, 0x40, 0x71, 0x10, 0x9f, 0x58, 0x2f, 0xb2, 0x39, 0x03, 0x6a, 0xa8,
	0x3e, 0x48, 0x82, 0xf5, 0x4e, 0x9f, 0x66, 0xff, 0x17, 0xf1, 0x05, 0x8a, 0xd0, 0x79, 0x26, 0xbd,
	0xdc, 0xa8, 0xdd, 0xea, 0xe9, 0xbe, 0xe6, 0xb6, 0xfe, 0xff, 0x49, 0xdd, 0x50, 0xf0, 0x49, 0x3a,
	0x7a, 0x9b, 0x52, 0xcd, 0x1c, 0xc6, 0x5b, 0x28, 0x79, 0x48, 0x62, 0x06, 0x5d, 0xe4, 0xba, 0x1f,
	0xc2, 0x3b, 0x38, 0x0b, 0xd6, 0x1d, 0xd2, 0x60, 0x44, 0x97, 0xf9, 0xe5, 0x8f, 0xa7, 0xb5, 0xef,
	0xd6, 0xf5, 0xba, 0x9a, 0xd7, 0x4e, 0x7a, 0xfb, 0x00, 0x57, 0x47, 0xfd, 0x78, 0x01, 0xd5, 0x7e,
	0x48, 0x21, 0x92, 0xac, 0x4f, 0x70, 0x05, 0xf0, 0x6a, 0x46, 0x0a, 0x6c, 0x3a, 0xea, 0xd7, 0xaa,
	0x2d, 0xf3, 0x2d, 0x9e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x52, 0x9e, 0x0d, 0xa6, 0x20, 0x01,
	0x00, 0x00,
}