// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/icaoracle/icaoracle.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MetricStatus indicates whether the Metric update ICA has been sent
type MetricStatus int32

const (
	MetricStatus_UNSPECIFIED MetricStatus = 0
	MetricStatus_QUEUED      MetricStatus = 1
	MetricStatus_IN_PROGRESS MetricStatus = 2
)

var MetricStatus_name = map[int32]string{
	0: "METRIC_STATUS_UNSPECIFIED",
	1: "METRIC_STATUS_QUEUED",
	2: "METRIC_STATUS_IN_PROGRESS",
}

var MetricStatus_value = map[string]int32{
	"METRIC_STATUS_UNSPECIFIED": 0,
	"METRIC_STATUS_QUEUED":      1,
	"METRIC_STATUS_IN_PROGRESS": 2,
}

func (x MetricStatus) String() string {
	return proto.EnumName(MetricStatus_name, int32(x))
}

func (MetricStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_842e38c1f0da9e66, []int{0}
}

// Oracle structure stores context about the CW oracle sitting a different chain
type Oracle struct {
	ChainId         string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	ConnectionId    string `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	ChannelId       string `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	PortId          string `protobuf:"bytes,4,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	IcaAddress      string `protobuf:"bytes,5,opt,name=ica_address,json=icaAddress,proto3" json:"ica_address,omitempty"`
	ContractAddress string `protobuf:"bytes,6,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	Active          bool   `protobuf:"varint,7,opt,name=active,proto3" json:"active,omitempty"`
}

func (m *Oracle) Reset()         { *m = Oracle{} }
func (m *Oracle) String() string { return proto.CompactTextString(m) }
func (*Oracle) ProtoMessage()    {}
func (*Oracle) Descriptor() ([]byte, []int) {
	return fileDescriptor_842e38c1f0da9e66, []int{0}
}
func (m *Oracle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Oracle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Oracle.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Oracle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Oracle.Merge(m, src)
}
func (m *Oracle) XXX_Size() int {
	return m.Size()
}
func (m *Oracle) XXX_DiscardUnknown() {
	xxx_messageInfo_Oracle.DiscardUnknown(m)
}

var xxx_messageInfo_Oracle proto.InternalMessageInfo

func (m *Oracle) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *Oracle) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *Oracle) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *Oracle) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *Oracle) GetIcaAddress() string {
	if m != nil {
		return m.IcaAddress
	}
	return ""
}

func (m *Oracle) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *Oracle) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

// Metric structure stores a generic metric using a key value structure
// along with additional context
type Metric struct {
	Key               string       `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value             string       `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	MetricType        string       `protobuf:"bytes,3,opt,name=metric_type,json=metricType,proto3" json:"metric_type,omitempty"`
	UpdateTime        int64        `protobuf:"varint,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	BlockHeight       int64        `protobuf:"varint,5,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Attributes        string       `protobuf:"bytes,6,opt,name=attributes,proto3" json:"attributes,omitempty"`
	DestinationOracle string       `protobuf:"bytes,7,opt,name=destination_oracle,json=destinationOracle,proto3" json:"destination_oracle,omitempty"`
	Status            MetricStatus `protobuf:"varint,8,opt,name=status,proto3,enum=stride.icaoracle.MetricStatus" json:"status,omitempty"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_842e38c1f0da9e66, []int{1}
}
func (m *Metric) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return m.Size()
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Metric) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Metric) GetMetricType() string {
	if m != nil {
		return m.MetricType
	}
	return ""
}

func (m *Metric) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *Metric) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *Metric) GetAttributes() string {
	if m != nil {
		return m.Attributes
	}
	return ""
}

func (m *Metric) GetDestinationOracle() string {
	if m != nil {
		return m.DestinationOracle
	}
	return ""
}

func (m *Metric) GetStatus() MetricStatus {
	if m != nil {
		return m.Status
	}
	return MetricStatus_UNSPECIFIED
}

// Attributes associated with a RedemptionRate metric update
type RedemptionRateAttributes struct {
	SttokenDenom string `protobuf:"bytes,1,opt,name=sttoken_denom,json=sttokenDenom,proto3" json:"sttoken_denom,omitempty"`
}

func (m *RedemptionRateAttributes) Reset()         { *m = RedemptionRateAttributes{} }
func (m *RedemptionRateAttributes) String() string { return proto.CompactTextString(m) }
func (*RedemptionRateAttributes) ProtoMessage()    {}
func (*RedemptionRateAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_842e38c1f0da9e66, []int{2}
}
func (m *RedemptionRateAttributes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RedemptionRateAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RedemptionRateAttributes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RedemptionRateAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedemptionRateAttributes.Merge(m, src)
}
func (m *RedemptionRateAttributes) XXX_Size() int {
	return m.Size()
}
func (m *RedemptionRateAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_RedemptionRateAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_RedemptionRateAttributes proto.InternalMessageInfo

func (m *RedemptionRateAttributes) GetSttokenDenom() string {
	if m != nil {
		return m.SttokenDenom
	}
	return ""
}

func init() {
	proto.RegisterEnum("stride.icaoracle.MetricStatus", MetricStatus_name, MetricStatus_value)
	proto.RegisterType((*Oracle)(nil), "stride.icaoracle.Oracle")
	proto.RegisterType((*Metric)(nil), "stride.icaoracle.Metric")
	proto.RegisterType((*RedemptionRateAttributes)(nil), "stride.icaoracle.RedemptionRateAttributes")
}

func init() { proto.RegisterFile("stride/icaoracle/icaoracle.proto", fileDescriptor_842e38c1f0da9e66) }

var fileDescriptor_842e38c1f0da9e66 = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x93, 0xc1, 0x6e, 0xd3, 0x4c,
	0x14, 0x85, 0xe3, 0xe6, 0xaf, 0xdb, 0xde, 0xa6, 0x7f, 0xcd, 0x28, 0x82, 0x34, 0x12, 0xc6, 0xa4,
	0x2c, 0x02, 0x52, 0x1d, 0x41, 0xa5, 0x6e, 0x51, 0x68, 0x0c, 0x58, 0x22, 0x6d, 0xb1, 0x93, 0x0d,
	0x1b, 0x6b, 0x32, 0x33, 0x4a, 0x46, 0x89, 0x3d, 0x96, 0x3d, 0x89, 0xc8, 0x2b, 0x74, 0xc5, 0x82,
	0x25, 0x7d, 0x1f, 0x96, 0x5d, 0xb2, 0x44, 0x09, 0x0f, 0x82, 0x3c, 0x76, 0x5a, 0x03, 0xbb, 0xb9,
	0xdf, 0x39, 0xbe, 0xf2, 0xb9, 0x33, 0x17, 0xac, 0x54, 0x26, 0x9c, 0xb2, 0x0e, 0x27, 0x58, 0x24,
	0x98, 0xcc, 0x4a, 0x27, 0x3b, 0x4e, 0x84, 0x14, 0xc8, 0xc8, 0x1d, 0xf6, 0x1d, 0x6f, 0xd6, 0xc7,
	0x62, 0x2c, 0x94, 0xd8, 0xc9, 0x4e, 0xb9, 0xaf, 0xf5, 0x4b, 0x03, 0xfd, 0x52, 0x19, 0xd0, 0x11,
	0xec, 0x92, 0x09, 0xe6, 0x51, 0xc0, 0x69, 0x43, 0xb3, 0xb4, 0xf6, 0x9e, 0xb7, 0xa3, 0x6a, 0x97,
	0xa2, 0x63, 0x38, 0x20, 0x22, 0x8a, 0x18, 0x91, 0x5c, 0x28, 0x7d, 0x4b, 0xe9, 0xb5, 0x7b, 0xe8,
	0x52, 0xf4, 0x18, 0x80, 0x4c, 0x70, 0x14, 0xb1, 0x59, 0xe6, 0xa8, 0x2a, 0xc7, 0x5e, 0x41, 0x5c,
	0x8a, 0x1e, 0xc1, 0x4e, 0x2c, 0x12, 0x99, 0x69, 0xff, 0x29, 0x4d, 0xcf, 0x4a, 0x97, 0xa2, 0x27,
	0xb0, 0xcf, 0x09, 0x0e, 0x30, 0xa5, 0x09, 0x4b, 0xd3, 0xc6, 0xb6, 0x12, 0x81, 0x13, 0xdc, 0xcd,
	0x09, 0x7a, 0x0e, 0x06, 0x11, 0x91, 0x4c, 0x30, 0x91, 0x77, 0x2e, 0x5d, 0xb9, 0x0e, 0x37, 0x7c,
	0x63, 0x7d, 0x08, 0x3a, 0x26, 0x92, 0x2f, 0x58, 0x63, 0xc7, 0xd2, 0xda, 0xbb, 0x5e, 0x51, 0xb5,
	0xbe, 0x6d, 0x81, 0xde, 0x67, 0x32, 0xe1, 0x04, 0x19, 0x50, 0x9d, 0xb2, 0x65, 0x91, 0x30, 0x3b,
	0xa2, 0x3a, 0x6c, 0x2f, 0xf0, 0x6c, 0xce, 0x8a, 0x54, 0x79, 0x91, 0xfd, 0x56, 0xa8, 0xbe, 0x08,
	0xe4, 0x32, 0x66, 0x45, 0x1e, 0xc8, 0xd1, 0x60, 0x19, 0x2b, 0xc3, 0x3c, 0xa6, 0x58, 0xb2, 0x40,
	0xf2, 0x90, 0xa9, 0x50, 0x55, 0x0f, 0x72, 0x34, 0xe0, 0x21, 0x43, 0x4f, 0xa1, 0x36, 0x9a, 0x09,
	0x32, 0x0d, 0x26, 0x8c, 0x8f, 0x27, 0x52, 0x25, 0xab, 0x7a, 0xfb, 0x8a, 0xbd, 0x57, 0x08, 0x99,
	0x00, 0x58, 0xca, 0x84, 0x8f, 0xe6, 0x92, 0x6d, 0x42, 0x95, 0x08, 0x3a, 0x01, 0x44, 0x59, 0x2a,
	0x79, 0x84, 0xd5, 0xe4, 0xf3, 0xab, 0x54, 0xd9, 0xf6, 0xbc, 0x07, 0x25, 0xa5, 0xb8, 0xc2, 0x33,
	0xd0, 0x53, 0x89, 0xe5, 0x3c, 0x6d, 0xec, 0x5a, 0x5a, 0xfb, 0xff, 0x57, 0xa6, 0xfd, 0xf7, 0x33,
	0xb0, 0xf3, 0x29, 0xf8, 0xca, 0xe5, 0x15, 0xee, 0xd6, 0x6b, 0x68, 0x78, 0x8c, 0xb2, 0x30, 0xce,
	0x7a, 0x79, 0x58, 0xb2, 0xee, 0xfd, 0x2f, 0x1c, 0xc3, 0x41, 0x2a, 0xa5, 0x98, 0xb2, 0x28, 0xa0,
	0x2c, 0x12, 0x61, 0x31, 0xb9, 0x5a, 0x01, 0x7b, 0x19, 0x7b, 0xf1, 0x55, 0x83, 0x5a, 0xb9, 0x33,
	0xb2, 0xe1, 0xa8, 0xef, 0x0c, 0x3c, 0xf7, 0x3c, 0xf0, 0x07, 0xdd, 0xc1, 0xd0, 0x0f, 0x86, 0x17,
	0xfe, 0x95, 0x73, 0xee, 0xbe, 0x75, 0x9d, 0x9e, 0x51, 0x69, 0x1e, 0x5e, 0xdf, 0x58, 0xfb, 0x25,
	0x84, 0x9e, 0x41, 0xfd, 0x4f, 0xff, 0xc7, 0xa1, 0x33, 0x74, 0x7a, 0x86, 0xd6, 0x84, 0xeb, 0x1b,
	0x4b, 0xcf, 0xab, 0x7f, 0xbb, 0xba, 0x17, 0xc1, 0x95, 0x77, 0xf9, 0xce, 0x73, 0x7c, 0xdf, 0xd8,
	0xca, 0xbb, 0x96, 0xd0, 0x9b, 0xfe, 0xf7, 0x95, 0xa9, 0xdd, 0xae, 0x4c, 0xed, 0xe7, 0xca, 0xd4,
	0xbe, 0xac, 0xcd, 0xca, 0xed, 0xda, 0xac, 0xfc, 0x58, 0x9b, 0x95, 0x4f, 0xa7, 0x63, 0x2e, 0x27,
	0xf3, 0x91, 0x4d, 0x44, 0xd8, 0xf1, 0xd5, 0x8c, 0x4e, 0x3e, 0xe0, 0x51, 0xda, 0x29, 0x16, 0x6b,
	0xf1, 0xf2, 0xac, 0xf3, 0xb9, 0xb4, 0x5e, 0xd9, 0x13, 0x48, 0x47, 0xba, 0xda, 0x99, 0xd3, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x24, 0xa4, 0x53, 0xc9, 0x7f, 0x03, 0x00, 0x00,
}

func (m *Oracle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Oracle) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Oracle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.IcaAddress) > 0 {
		i -= len(m.IcaAddress)
		copy(dAtA[i:], m.IcaAddress)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.IcaAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Metric) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metric) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metric) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintIcaoracle(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x40
	}
	if len(m.DestinationOracle) > 0 {
		i -= len(m.DestinationOracle)
		copy(dAtA[i:], m.DestinationOracle)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.DestinationOracle)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Attributes) > 0 {
		i -= len(m.Attributes)
		copy(dAtA[i:], m.Attributes)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.Attributes)))
		i--
		dAtA[i] = 0x32
	}
	if m.BlockHeight != 0 {
		i = encodeVarintIcaoracle(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.UpdateTime != 0 {
		i = encodeVarintIcaoracle(dAtA, i, uint64(m.UpdateTime))
		i--
		dAtA[i] = 0x20
	}
	if len(m.MetricType) > 0 {
		i -= len(m.MetricType)
		copy(dAtA[i:], m.MetricType)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.MetricType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RedemptionRateAttributes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RedemptionRateAttributes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RedemptionRateAttributes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SttokenDenom) > 0 {
		i -= len(m.SttokenDenom)
		copy(dAtA[i:], m.SttokenDenom)
		i = encodeVarintIcaoracle(dAtA, i, uint64(len(m.SttokenDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIcaoracle(dAtA []byte, offset int, v uint64) int {
	offset -= sovIcaoracle(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Oracle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	l = len(m.IcaAddress)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	if m.Active {
		n += 2
	}
	return n
}

func (m *Metric) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	l = len(m.MetricType)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	if m.UpdateTime != 0 {
		n += 1 + sovIcaoracle(uint64(m.UpdateTime))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovIcaoracle(uint64(m.BlockHeight))
	}
	l = len(m.Attributes)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	l = len(m.DestinationOracle)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovIcaoracle(uint64(m.Status))
	}
	return n
}

func (m *RedemptionRateAttributes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SttokenDenom)
	if l > 0 {
		n += 1 + l + sovIcaoracle(uint64(l))
	}
	return n
}

func sovIcaoracle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIcaoracle(x uint64) (n int) {
	return sovIcaoracle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Oracle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIcaoracle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Oracle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Oracle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IcaAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IcaAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Active = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipIcaoracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Metric) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIcaoracle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Metric: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metric: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetricType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetricType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateTime", wireType)
			}
			m.UpdateTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdateTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationOracle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationOracle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= MetricStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIcaoracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RedemptionRateAttributes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIcaoracle
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RedemptionRateAttributes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RedemptionRateAttributes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SttokenDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIcaoracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SttokenDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIcaoracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIcaoracle
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipIcaoracle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIcaoracle
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIcaoracle
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthIcaoracle
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIcaoracle
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIcaoracle
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIcaoracle        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIcaoracle          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIcaoracle = fmt.Errorf("proto: unexpected end of group")
)
