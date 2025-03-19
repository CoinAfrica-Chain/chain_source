// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evmos/claims/v1/genesis.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"

	_ "github.com/cosmos/gogoproto/gogoproto"
	"github.com/evmos/evmos/v12/x/claims/types"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
	_ = time.Kitchen
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// V3GenesisState define the claims module's genesis state.
type V3GenesisState struct {
	// V3Params defines all the parameters of the module.
	V3Params V3Params `protobuf:"bytes,1,opt,name=V3Params,proto3" json:"V3Params"`
	// claims_records is a list of claim records with the corresponding airdrop recipient
	ClaimsRecords []types.ClaimsRecordAddress `protobuf:"bytes,2,rep,name=claims_records,json=claimsRecords,proto3" json:"claims_records"`
}

func (m *V3GenesisState) Reset()         { *m = V3GenesisState{} }
func (m *V3GenesisState) String() string { return proto.CompactTextString(m) }
func (*V3GenesisState) ProtoMessage()    {}
func (*V3GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2f8f1d6f18af278, []int{0}
}

func (m *V3GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *V3GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_V3GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *V3GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_V3GenesisState.Merge(m, src)
}

func (m *V3GenesisState) XXX_Size() int {
	return m.Size()
}

func (m *V3GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_V3GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_V3GenesisState proto.InternalMessageInfo

func (m *V3GenesisState) GetV3Params() V3Params {
	if m != nil {
		return m.V3Params
	}
	return V3Params{}
}

func (m *V3GenesisState) GetClaimsRecords() []types.ClaimsRecordAddress {
	if m != nil {
		return m.ClaimsRecords
	}
	return nil
}

// V3Params defines the claims module's parameters.
type V3Params struct {
	// enable_claims is the parameter to enable the claiming process
	EnableClaims bool `protobuf:"varint,1,opt,name=enable_claims,json=enableClaims,proto3" json:"enable_claims,omitempty"`
	// airdrop_start_time defines the timestamp of the airdrop start
	AirdropStartTime time.Time `protobuf:"bytes,2,opt,name=airdrop_start_time,json=airdropStartTime,proto3,stdtime" json:"airdrop_start_time"`
	// duration_until_decay of claimable tokens begin
	DurationUntilDecay time.Duration `protobuf:"bytes,3,opt,name=duration_until_decay,json=durationUntilDecay,proto3,stdduration" json:"duration_until_decay"`
	// duration_of_decay for token claim decay period
	DurationOfDecay time.Duration `protobuf:"bytes,4,opt,name=duration_of_decay,json=durationOfDecay,proto3,stdduration" json:"duration_of_decay"`
	// claims_denom is the denomination of the claimable coin
	ClaimsDenom string `protobuf:"bytes,5,opt,name=claims_denom,json=claimsDenom,proto3" json:"claims_denom,omitempty"`
	// authorized_channels is the list of authorized channel identifiers that can perform address
	// attestations via IBC.
	AuthorizedChannels []string `protobuf:"bytes,6,rep,name=authorized_channels,json=authorizedChannels,proto3" json:"authorized_channels,omitempty"`
	// evm_channels is the list of channel identifiers from EVM compatible chains
	EVMChannels []string `protobuf:"bytes,7,rep,name=evm_channels,json=evmChannels,proto3" json:"evm_channels,omitempty"`
}

func (m *V3Params) Reset()         { *m = V3Params{} }
func (m *V3Params) String() string { return proto.CompactTextString(m) }
func (*V3Params) ProtoMessage()    {}
func (*V3Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2f8f1d6f18af278, []int{1}
}

func (m *V3Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *V3Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_V3Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *V3Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_V3Params.Merge(m, src)
}

func (m *V3Params) XXX_Size() int {
	return m.Size()
}

func (m *V3Params) XXX_DiscardUnknown() {
	xxx_messageInfo_V3Params.DiscardUnknown(m)
}

var xxx_messageInfo_V3Params proto.InternalMessageInfo

func (m *V3Params) GetEnableClaims() bool {
	if m != nil {
		return m.EnableClaims
	}
	return false
}

func (m *V3Params) GetAirdropStartTime() time.Time {
	if m != nil {
		return m.AirdropStartTime
	}
	return time.Time{}
}

func (m *V3Params) GetDurationUntilDecay() time.Duration {
	if m != nil {
		return m.DurationUntilDecay
	}
	return 0
}

func (m *V3Params) GetDurationOfDecay() time.Duration {
	if m != nil {
		return m.DurationOfDecay
	}
	return 0
}

func (m *V3Params) GetClaimsDenom() string {
	if m != nil {
		return m.ClaimsDenom
	}
	return ""
}

func (m *V3Params) GetAuthorizedChannels() []string {
	if m != nil {
		return m.AuthorizedChannels
	}
	return nil
}

func (m *V3Params) GetEVMChannels() []string {
	if m != nil {
		return m.EVMChannels
	}
	return nil
}

func init() {
	proto.RegisterType((*V3GenesisState)(nil), "evmos.claims.v1.V3GenesisState")
	proto.RegisterType((*V3Params)(nil), "evmos.claims.v1.V3Params")
}

func init() { proto.RegisterFile("evmos/claims/v1/genesis.proto", fileDescriptor_f2f8f1d6f18af278) }

var fileDescriptor_f2f8f1d6f18af278 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6e, 0xd3, 0x40,
	0x14, 0xc6, 0xe3, 0xa6, 0x84, 0x76, 0x9c, 0x12, 0x18, 0x2a, 0x61, 0x22, 0x70, 0x42, 0x61, 0x11,
	0x36, 0x36, 0x09, 0xe2, 0x00, 0x24, 0x41, 0xac, 0x50, 0xc1, 0xa5, 0x2c, 0xd8, 0x58, 0x13, 0x7b,
	0xe2, 0x58, 0xf2, 0x78, 0xac, 0x99, 0xb1, 0x45, 0x39, 0x45, 0x97, 0xbd, 0x06, 0xb7, 0xe8, 0xb2,
	0x4b, 0x56, 0x05, 0x25, 0x17, 0x41, 0xf3, 0x8f, 0xa2, 0x74, 0xc3, 0xc6, 0x1a, 0xbf, 0xef, 0xf7,
	0x7d, 0xf3, 0xde, 0xd3, 0x80, 0xa7, 0xb8, 0x21, 0x94, 0x87, 0x49, 0x81, 0x72, 0xc2, 0xc3, 0x66,
	0x1c, 0x66, 0xb8, 0xc4, 0x3c, 0xe7, 0x41, 0xc5, 0xa8, 0xa0, 0xb0, 0xa7, 0xe4, 0x40, 0xcb, 0x41,
	0x33, 0xee, 0x3f, 0xd9, 0xe6, 0x8d, 0xa4, 0xf0, 0xfe, 0x61, 0x46, 0x33, 0xaa, 0x8e, 0xa1, 0x3c,
	0x99, 0xaa, 0x9f, 0x51, 0x9a, 0x15, 0x38, 0x54, 0x7f, 0x8b, 0x7a, 0x19, 0xa6, 0x35, 0x43, 0x22,
	0xa7, 0xa5, 0xd1, 0x07, 0xdb, 0xba, 0xc8, 0x09, 0xe6, 0x02, 0x91, 0x4a, 0x03, 0x47, 0x17, 0x0e,
	0xe8, 0xbe, 0xd7, 0x7d, 0x9d, 0x08, 0x24, 0x30, 0x7c, 0x03, 0x3a, 0x15, 0x62, 0x88, 0x70, 0xcf,
	0x19, 0x3a, 0x23, 0x77, 0xf2, 0x28, 0xd8, 0xea, 0x33, 0xf8, 0xa8, 0xe4, 0xe9, 0xee, 0xe5, 0xf5,
	0xa0, 0x15, 0x19, 0x18, 0x7e, 0x02, 0xf7, 0x34, 0x11, 0x33, 0x9c, 0x50, 0x96, 0x72, 0x6f, 0x67,
	0xd8, 0x1e, 0xb9, 0x93, 0x17, 0xb7, 0xec, 0x33, 0x75, 0x8a, 0x14, 0xf5, 0x36, 0x4d, 0x19, 0xe6,
	0x36, 0xeb, 0x20, 0xf9, 0x47, 0xe2, 0x47, 0x3f, 0xda, 0xa0, 0xa3, 0xef, 0x82, 0xcf, 0xc1, 0x01,
	0x2e, 0xd1, 0xa2, 0xc0, 0xb1, 0x46, 0x54, 0x6f, 0x7b, 0x51, 0x57, 0x17, 0x75, 0x22, 0x8c, 0x00,
	0x44, 0x39, 0x4b, 0x19, 0xad, 0x62, 0x2e, 0x10, 0x13, 0xb1, 0x9c, 0xd5, 0xdb, 0x51, 0x53, 0xf4,
	0x03, 0xbd, 0x88, 0xc0, 0x2e, 0x22, 0xf8, 0x6c, 0x17, 0x31, 0xdd, 0x93, 0x97, 0x9f, 0xff, 0x1a,
	0x38, 0xd1, 0x7d, 0xe3, 0x3f, 0x91, 0x76, 0x09, 0xc0, 0x53, 0x70, 0x68, 0x37, 0x1a, 0xd7, 0xa5,
	0xc8, 0x8b, 0x38, 0xc5, 0x09, 0x3a, 0xf3, 0xda, 0x2a, 0xf5, 0xf1, 0xad, 0xd4, 0xb9, 0x81, 0x75,
	0xe8, 0x85, 0x0c, 0x85, 0x36, 0xe0, 0x54, 0xfa, 0xe7, 0xd2, 0x0e, 0x8f, 0xc1, 0x83, 0xbf, 0xb1,
	0x74, 0x69, 0x32, 0x77, 0xff, 0x3f, 0xb3, 0x67, 0xdd, 0xc7, 0x4b, 0x1d, 0xf8, 0x0c, 0x74, 0xcd,
	0xfa, 0x53, 0x5c, 0x52, 0xe2, 0xdd, 0x19, 0x3a, 0xa3, 0xfd, 0xc8, 0xd5, 0xb5, 0xb9, 0x2c, 0xc1,
	0x10, 0x3c, 0x44, 0xb5, 0x58, 0x51, 0x96, 0x7f, 0xc7, 0x69, 0x9c, 0xac, 0x50, 0x59, 0xe2, 0x82,
	0x7b, 0x9d, 0x61, 0x7b, 0xb4, 0x1f, 0xc1, 0x1b, 0x69, 0x66, 0x14, 0x38, 0x01, 0x5d, 0xdc, 0x90,
	0x1b, 0xf2, 0xae, 0x24, 0xa7, 0xbd, 0xf5, 0xf5, 0xc0, 0x7d, 0xf7, 0xe5, 0x83, 0xc5, 0x22, 0x17,
	0x37, 0xc4, 0xfe, 0x4c, 0x67, 0x97, 0x6b, 0xdf, 0xb9, 0x5a, 0xfb, 0xce, 0xef, 0xb5, 0xef, 0x9c,
	0x6f, 0xfc, 0xd6, 0xd5, 0xc6, 0x6f, 0xfd, 0xdc, 0xf8, 0xad, 0xaf, 0x2f, 0xb3, 0x5c, 0xac, 0xea,
	0x45, 0x90, 0x50, 0x12, 0xea, 0x87, 0xae, 0xbf, 0xcd, 0xf8, 0x55, 0xf8, 0xcd, 0x3e, 0x7a, 0x71,
	0x56, 0x61, 0xbe, 0xe8, 0xa8, 0xd1, 0x5f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x17, 0x29, 0xc9,
	0xb8, 0x41, 0x03, 0x00, 0x00,
}

func (m *V3GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *V3GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *V3GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimsRecords) > 0 {
		for iNdEx := len(m.ClaimsRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimsRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.V3Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *V3Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *V3Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *V3Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EVMChannels) > 0 {
		for iNdEx := len(m.EVMChannels) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.EVMChannels[iNdEx])
			copy(dAtA[i:], m.EVMChannels[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.EVMChannels[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.AuthorizedChannels) > 0 {
		for iNdEx := len(m.AuthorizedChannels) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AuthorizedChannels[iNdEx])
			copy(dAtA[i:], m.AuthorizedChannels[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.AuthorizedChannels[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ClaimsDenom) > 0 {
		i -= len(m.ClaimsDenom)
		copy(dAtA[i:], m.ClaimsDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ClaimsDenom)))
		i--
		dAtA[i] = 0x2a
	}
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DurationOfDecay, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationOfDecay):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGenesis(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	n3, err3 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DurationUntilDecay, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationUntilDecay):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintGenesis(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x1a
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.AirdropStartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStartTime):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintGenesis(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x12
	if m.EnableClaims {
		i--
		if m.EnableClaims {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *V3GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.V3Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ClaimsRecords) > 0 {
		for _, e := range m.ClaimsRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *V3Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EnableClaims {
		n += 2
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStartTime)
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationUntilDecay)
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationOfDecay)
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.ClaimsDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.AuthorizedChannels) > 0 {
		for _, s := range m.AuthorizedChannels {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.EVMChannels) > 0 {
		for _, s := range m.EVMChannels {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *V3GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: V3GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: V3GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field V3Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.V3Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimsRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimsRecords = append(m.ClaimsRecords, types.ClaimsRecordAddress{})
			if err := m.ClaimsRecords[len(m.ClaimsRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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

func (m *V3Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: V3Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: V3Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableClaims", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			m.EnableClaims = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.AirdropStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationUntilDecay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DurationUntilDecay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationOfDecay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DurationOfDecay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimsDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimsDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizedChannels", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorizedChannels = append(m.AuthorizedChannels, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EVMChannels", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EVMChannels = append(m.EVMChannels, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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

func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
