// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/ad_group_ad_label.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A relationship between an ad group ad and a label.
type AdGroupAdLabel struct {
	// Immutable. The resource name of the ad group ad label.
	// Ad group ad label resource names have the form:
	// `customers/{customer_id}/adGroupAdLabels/{ad_group_id}~{ad_id}~{label_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The ad group ad to which the label is attached.
	AdGroupAd *wrappers.StringValue `protobuf:"bytes,2,opt,name=ad_group_ad,json=adGroupAd,proto3" json:"ad_group_ad,omitempty"`
	// Immutable. The label assigned to the ad group ad.
	Label                *wrappers.StringValue `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AdGroupAdLabel) Reset()         { *m = AdGroupAdLabel{} }
func (m *AdGroupAdLabel) String() string { return proto.CompactTextString(m) }
func (*AdGroupAdLabel) ProtoMessage()    {}
func (*AdGroupAdLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_2df918faaa4182f5, []int{0}
}

func (m *AdGroupAdLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAdLabel.Unmarshal(m, b)
}
func (m *AdGroupAdLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAdLabel.Marshal(b, m, deterministic)
}
func (m *AdGroupAdLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAdLabel.Merge(m, src)
}
func (m *AdGroupAdLabel) XXX_Size() int {
	return xxx_messageInfo_AdGroupAdLabel.Size(m)
}
func (m *AdGroupAdLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAdLabel.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAdLabel proto.InternalMessageInfo

func (m *AdGroupAdLabel) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupAdLabel) GetAdGroupAd() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupAd
	}
	return nil
}

func (m *AdGroupAdLabel) GetLabel() *wrappers.StringValue {
	if m != nil {
		return m.Label
	}
	return nil
}

func init() {
	proto.RegisterType((*AdGroupAdLabel)(nil), "google.ads.googleads.v1.resources.AdGroupAdLabel")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/ad_group_ad_label.proto", fileDescriptor_2df918faaa4182f5)
}

var fileDescriptor_2df918faaa4182f5 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcf, 0x8a, 0xd4, 0x30,
	0x18, 0xa7, 0x1d, 0x56, 0xd8, 0xac, 0x7a, 0xa8, 0x97, 0x71, 0x59, 0xd6, 0xd9, 0xc5, 0xd5, 0xd5,
	0x43, 0x42, 0xf5, 0xa2, 0xf1, 0x94, 0x5e, 0x06, 0x44, 0x64, 0x19, 0xa5, 0x87, 0xa5, 0x50, 0xd2,
	0x26, 0x93, 0x2d, 0xb4, 0x4d, 0x49, 0xda, 0xf1, 0x30, 0x0c, 0xf8, 0x10, 0x3e, 0x81, 0x47, 0x1f,
	0xc5, 0xa7, 0xf0, 0x3c, 0x8f, 0xe0, 0x49, 0xda, 0x34, 0x9d, 0x19, 0x44, 0x67, 0x6e, 0xbf, 0xf0,
	0xfd, 0xfe, 0xe5, 0xe3, 0x03, 0x6f, 0x85, 0x94, 0x22, 0xe7, 0x88, 0x32, 0x8d, 0x0c, 0x6c, 0xd1,
	0xc2, 0x47, 0x8a, 0x6b, 0xd9, 0xa8, 0x94, 0x6b, 0x44, 0x59, 0x2c, 0x94, 0x6c, 0xaa, 0x98, 0xb2,
	0x38, 0xa7, 0x09, 0xcf, 0x61, 0xa5, 0x64, 0x2d, 0xbd, 0x0b, 0xc3, 0x87, 0x94, 0x69, 0x38, 0x48,
	0xe1, 0xc2, 0x87, 0x83, 0xf4, 0xf4, 0x89, 0x75, 0xaf, 0x32, 0x34, 0xcf, 0x78, 0xce, 0xe2, 0x84,
	0xdf, 0xd1, 0x45, 0x26, 0x95, 0xf1, 0x38, 0x7d, 0xbc, 0x45, 0xb0, 0xb2, 0x7e, 0x74, 0xde, 0x8f,
	0xba, 0x57, 0xd2, 0xcc, 0xd1, 0x17, 0x45, 0xab, 0x8a, 0x2b, 0xdd, 0xcf, 0xcf, 0xb6, 0xa4, 0xb4,
	0x2c, 0x65, 0x4d, 0xeb, 0x4c, 0x96, 0xfd, 0xf4, 0xf2, 0xdb, 0x08, 0x3c, 0x24, 0x6c, 0xda, 0xf6,
	0x26, 0xec, 0x43, 0xdb, 0xda, 0xfb, 0x0c, 0x1e, 0xd8, 0x88, 0xb8, 0xa4, 0x05, 0x1f, 0x3b, 0x13,
	0xe7, 0xfa, 0x38, 0x40, 0xbf, 0xc8, 0xd1, 0x6f, 0xf2, 0x02, 0x3c, 0xdf, 0xfc, 0xa1, 0x47, 0x55,
	0xa6, 0x61, 0x2a, 0x0b, 0xb4, 0xeb, 0x33, 0xbb, 0x6f, 0x5d, 0x3e, 0xd2, 0x82, 0x7b, 0x77, 0xe0,
	0x64, 0x6b, 0x41, 0x63, 0x77, 0xe2, 0x5c, 0x9f, 0xbc, 0x3a, 0xeb, 0x2d, 0xa0, 0x2d, 0x0f, 0x3f,
	0xd5, 0x2a, 0x2b, 0x45, 0x48, 0xf3, 0x86, 0x07, 0x2f, 0xbb, 0xc4, 0xa7, 0xe0, 0x72, 0x7f, 0xe2,
	0xec, 0x98, 0x5a, 0xe8, 0xdd, 0x82, 0xa3, 0x6e, 0xfd, 0xe3, 0xd1, 0x01, 0x19, 0xcf, 0xba, 0x8c,
	0x09, 0x38, 0xff, 0x67, 0x86, 0xf9, 0x8c, 0xb1, 0xc4, 0xf3, 0x35, 0x49, 0x0f, 0xde, 0x80, 0xf7,
	0x26, 0x6d, 0x74, 0x2d, 0x0b, 0xae, 0x34, 0x5a, 0x5a, 0xb8, 0x42, 0x74, 0x87, 0xa4, 0xd1, 0xf2,
	0xaf, 0xc3, 0x59, 0x05, 0x5f, 0x5d, 0x70, 0x95, 0xca, 0x02, 0xee, 0x3d, 0x9d, 0xe0, 0xd1, 0x6e,
	0xe6, 0x4d, 0xfb, 0xc9, 0x1b, 0xe7, 0xf6, 0x7d, 0xaf, 0x14, 0x32, 0xa7, 0xa5, 0x80, 0x52, 0x09,
	0x24, 0x78, 0xd9, 0xad, 0x00, 0x6d, 0x4a, 0xff, 0xe7, 0x9c, 0xdf, 0x0d, 0xe8, 0xbb, 0x3b, 0x9a,
	0x12, 0xf2, 0xc3, 0xbd, 0x98, 0x1a, 0x4b, 0xc2, 0x34, 0x34, 0xb0, 0x45, 0xa1, 0x0f, 0x67, 0x96,
	0xf9, 0xd3, 0x72, 0x22, 0xc2, 0x74, 0x34, 0x70, 0xa2, 0xd0, 0x8f, 0x06, 0xce, 0xda, 0xbd, 0x32,
	0x03, 0x8c, 0x09, 0xd3, 0x18, 0x0f, 0x2c, 0x8c, 0x43, 0x1f, 0xe3, 0x81, 0x97, 0xdc, 0xeb, 0xca,
	0xbe, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x75, 0x2f, 0xf4, 0xbf, 0x7a, 0x03, 0x00, 0x00,
}
