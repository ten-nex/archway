// DONTCOVER
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: archway/tracking/v1/tracking.proto

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

// ContractOperation denotes which operation consumed gas.
type ContractOperation int32

const (
	ContractOperation_CONTRACT_OPERATION_UNSPECIFIED   ContractOperation = 0
	ContractOperation_CONTRACT_OPERATION_INSTANTIATION ContractOperation = 1
	ContractOperation_CONTRACT_OPERATION_EXECUTION     ContractOperation = 2
	ContractOperation_CONTRACT_OPERATION_QUERY         ContractOperation = 3
	ContractOperation_CONTRACT_OPERATION_MIGRATE       ContractOperation = 4
	ContractOperation_CONTRACT_OPERATION_IBC           ContractOperation = 5
	ContractOperation_CONTRACT_OPERATION_SUDO          ContractOperation = 6
	ContractOperation_CONTRACT_OPERATION_REPLY         ContractOperation = 7
)

var ContractOperation_name = map[int32]string{
	0: "CONTRACT_OPERATION_UNSPECIFIED",
	1: "CONTRACT_OPERATION_INSTANTIATION",
	2: "CONTRACT_OPERATION_EXECUTION",
	3: "CONTRACT_OPERATION_QUERY",
	4: "CONTRACT_OPERATION_MIGRATE",
	5: "CONTRACT_OPERATION_IBC",
	6: "CONTRACT_OPERATION_SUDO",
	7: "CONTRACT_OPERATION_REPLY",
}

var ContractOperation_value = map[string]int32{
	"CONTRACT_OPERATION_UNSPECIFIED":   0,
	"CONTRACT_OPERATION_INSTANTIATION": 1,
	"CONTRACT_OPERATION_EXECUTION":     2,
	"CONTRACT_OPERATION_QUERY":         3,
	"CONTRACT_OPERATION_MIGRATE":       4,
	"CONTRACT_OPERATION_IBC":           5,
	"CONTRACT_OPERATION_SUDO":          6,
	"CONTRACT_OPERATION_REPLY":         7,
}

func (x ContractOperation) String() string {
	return proto.EnumName(ContractOperation_name, int32(x))
}

func (ContractOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_43be6db97c4ee6cb, []int{0}
}

// TxInfo keeps a transaction gas tracking data.
// Object is being created at the module EndBlocker.
type TxInfo struct {
	// id defines the unique transaction ID.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// height defines the block height of the transaction.
	Height int64 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	// total_gas defines total gas consumption by the transaction.
	// It is the sum of gas consumed by all contract operations (VM + SDK gas).
	TotalGas uint64 `protobuf:"varint,3,opt,name=total_gas,json=totalGas,proto3" json:"total_gas,omitempty"`
}

func (m *TxInfo) Reset()         { *m = TxInfo{} }
func (m *TxInfo) String() string { return proto.CompactTextString(m) }
func (*TxInfo) ProtoMessage()    {}
func (*TxInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_43be6db97c4ee6cb, []int{0}
}
func (m *TxInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxInfo.Merge(m, src)
}
func (m *TxInfo) XXX_Size() int {
	return m.Size()
}
func (m *TxInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TxInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TxInfo proto.InternalMessageInfo

func (m *TxInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TxInfo) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *TxInfo) GetTotalGas() uint64 {
	if m != nil {
		return m.TotalGas
	}
	return 0
}

// ContractOperationInfo keeps a single contract operation gas consumption data.
// Object is being created by the IngestGasRecord call from the wasmd.
type ContractOperationInfo struct {
	// id defines the unique operation ID.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// tx_id defines a transaction ID operation relates to (TxInfo.id).
	TxId uint64 `protobuf:"varint,2,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	// contract_address defines the contract address operation relates to.
	ContractAddress string `protobuf:"bytes,3,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// operation_type defines the gas consumption type.
	OperationType ContractOperation `protobuf:"varint,4,opt,name=operation_type,json=operationType,proto3,enum=archway.tracking.v1.ContractOperation" json:"operation_type,omitempty"`
	// vm_gas is the gas consumption reported by the WASM VM.
	// Value is adjusted by this module (CalculateUpdatedGas func).
	VmGas uint64 `protobuf:"varint,5,opt,name=vm_gas,json=vmGas,proto3" json:"vm_gas,omitempty"`
	// sdk_gas is the gas consumption reported by the SDK gas meter and the WASM
	// GasRegister (cost of Execute/Query/etc). Value is adjusted by this module
	// (CalculateUpdatedGas func).
	SdkGas uint64 `protobuf:"varint,6,opt,name=sdk_gas,json=sdkGas,proto3" json:"sdk_gas,omitempty"`
}

func (m *ContractOperationInfo) Reset()         { *m = ContractOperationInfo{} }
func (m *ContractOperationInfo) String() string { return proto.CompactTextString(m) }
func (*ContractOperationInfo) ProtoMessage()    {}
func (*ContractOperationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_43be6db97c4ee6cb, []int{1}
}
func (m *ContractOperationInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractOperationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractOperationInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractOperationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractOperationInfo.Merge(m, src)
}
func (m *ContractOperationInfo) XXX_Size() int {
	return m.Size()
}
func (m *ContractOperationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractOperationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ContractOperationInfo proto.InternalMessageInfo

func (m *ContractOperationInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ContractOperationInfo) GetTxId() uint64 {
	if m != nil {
		return m.TxId
	}
	return 0
}

func (m *ContractOperationInfo) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *ContractOperationInfo) GetOperationType() ContractOperation {
	if m != nil {
		return m.OperationType
	}
	return ContractOperation_CONTRACT_OPERATION_UNSPECIFIED
}

func (m *ContractOperationInfo) GetVmGas() uint64 {
	if m != nil {
		return m.VmGas
	}
	return 0
}

func (m *ContractOperationInfo) GetSdkGas() uint64 {
	if m != nil {
		return m.SdkGas
	}
	return 0
}

// BlockTracking is the tracking information for a block.
type BlockTracking struct {
	// txs defines the list of transactions tracked in the block.
	Txs []TxTracking `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs"`
}

func (m *BlockTracking) Reset()         { *m = BlockTracking{} }
func (m *BlockTracking) String() string { return proto.CompactTextString(m) }
func (*BlockTracking) ProtoMessage()    {}
func (*BlockTracking) Descriptor() ([]byte, []int) {
	return fileDescriptor_43be6db97c4ee6cb, []int{2}
}
func (m *BlockTracking) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockTracking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockTracking.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockTracking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockTracking.Merge(m, src)
}
func (m *BlockTracking) XXX_Size() int {
	return m.Size()
}
func (m *BlockTracking) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockTracking.DiscardUnknown(m)
}

var xxx_messageInfo_BlockTracking proto.InternalMessageInfo

func (m *BlockTracking) GetTxs() []TxTracking {
	if m != nil {
		return m.Txs
	}
	return nil
}

// TxTracking is the tracking information for a single transaction.
type TxTracking struct {
	// info defines the transaction details.
	Info TxInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info"`
	// contract_operations defines the list of contract operations consumed by the
	// transaction.
	ContractOperations []ContractOperationInfo `protobuf:"bytes,2,rep,name=contract_operations,json=contractOperations,proto3" json:"contract_operations"`
}

func (m *TxTracking) Reset()         { *m = TxTracking{} }
func (m *TxTracking) String() string { return proto.CompactTextString(m) }
func (*TxTracking) ProtoMessage()    {}
func (*TxTracking) Descriptor() ([]byte, []int) {
	return fileDescriptor_43be6db97c4ee6cb, []int{3}
}
func (m *TxTracking) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxTracking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxTracking.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxTracking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxTracking.Merge(m, src)
}
func (m *TxTracking) XXX_Size() int {
	return m.Size()
}
func (m *TxTracking) XXX_DiscardUnknown() {
	xxx_messageInfo_TxTracking.DiscardUnknown(m)
}

var xxx_messageInfo_TxTracking proto.InternalMessageInfo

func (m *TxTracking) GetInfo() TxInfo {
	if m != nil {
		return m.Info
	}
	return TxInfo{}
}

func (m *TxTracking) GetContractOperations() []ContractOperationInfo {
	if m != nil {
		return m.ContractOperations
	}
	return nil
}

func init() {
	proto.RegisterEnum("archway.tracking.v1.ContractOperation", ContractOperation_name, ContractOperation_value)
	proto.RegisterType((*TxInfo)(nil), "archway.tracking.v1.TxInfo")
	proto.RegisterType((*ContractOperationInfo)(nil), "archway.tracking.v1.ContractOperationInfo")
	proto.RegisterType((*BlockTracking)(nil), "archway.tracking.v1.BlockTracking")
	proto.RegisterType((*TxTracking)(nil), "archway.tracking.v1.TxTracking")
}

func init() {
	proto.RegisterFile("archway/tracking/v1/tracking.proto", fileDescriptor_43be6db97c4ee6cb)
}

var fileDescriptor_43be6db97c4ee6cb = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0x9b, 0x40,
	0x10, 0x35, 0x18, 0x93, 0x66, 0xa2, 0xa4, 0x74, 0xdd, 0x24, 0xc8, 0x8e, 0x08, 0xb2, 0xaa, 0xca,
	0x8d, 0x54, 0x5b, 0x71, 0x54, 0xf5, 0x6c, 0x13, 0xea, 0x22, 0xd5, 0x1f, 0xc5, 0x58, 0x6a, 0x7a,
	0x41, 0x04, 0x08, 0x46, 0x8e, 0x59, 0x0b, 0xb6, 0x0e, 0xfe, 0x17, 0x3d, 0xf4, 0x37, 0xf4, 0xb7,
	0xe4, 0x98, 0x63, 0x4f, 0x55, 0x65, 0x4b, 0xfd, 0x1d, 0x15, 0x6b, 0xec, 0x54, 0x0d, 0x91, 0x7a,
	0xdb, 0x7d, 0xf3, 0xe6, 0xbd, 0x37, 0x03, 0x0b, 0x15, 0x2b, 0xb4, 0x47, 0x37, 0xd6, 0xbc, 0x4e,
	0x42, 0xcb, 0x1e, 0xfb, 0x81, 0x57, 0x9f, 0x9d, 0x6e, 0xce, 0xb5, 0x69, 0x88, 0x09, 0x46, 0xc5,
	0x94, 0x53, 0xdb, 0xe0, 0xb3, 0xd3, 0xd2, 0x73, 0x0f, 0x7b, 0x98, 0xd6, 0xeb, 0xc9, 0x69, 0x45,
	0xad, 0x74, 0x80, 0x37, 0x62, 0x2d, 0xb8, 0xc2, 0x68, 0x0f, 0x58, 0xdf, 0x11, 0x19, 0x99, 0xa9,
	0x72, 0x3a, 0xeb, 0x3b, 0xe8, 0x00, 0xf8, 0x91, 0xeb, 0x7b, 0x23, 0x22, 0xb2, 0x32, 0x53, 0xcd,
	0xeb, 0xe9, 0x0d, 0x95, 0x61, 0x9b, 0x60, 0x62, 0x5d, 0x9b, 0x9e, 0x15, 0x89, 0x79, 0x4a, 0x7f,
	0x42, 0x81, 0xb6, 0x15, 0x55, 0x7e, 0x33, 0xb0, 0xaf, 0xe0, 0x20, 0xf1, 0x25, 0xbd, 0xa9, 0x1b,
	0x5a, 0xc4, 0xc7, 0x41, 0xa6, 0x7c, 0x11, 0x0a, 0x24, 0x36, 0x7d, 0x87, 0xaa, 0x73, 0x3a, 0x47,
	0x62, 0xcd, 0x41, 0xaf, 0x40, 0xb0, 0xd3, 0x6e, 0xd3, 0x72, 0x9c, 0xd0, 0x8d, 0x56, 0x16, 0xdb,
	0xfa, 0xd3, 0x35, 0xde, 0x5c, 0xc1, 0xa8, 0x03, 0x7b, 0x78, 0x6d, 0x60, 0x92, 0xf9, 0xd4, 0x15,
	0x39, 0x99, 0xa9, 0xee, 0x35, 0x5e, 0xd6, 0x32, 0x86, 0xaf, 0x3d, 0xc8, 0xa4, 0xef, 0x6e, 0xba,
	0x8d, 0xf9, 0xd4, 0x45, 0xfb, 0xc0, 0xcf, 0x26, 0x74, 0xa4, 0x02, 0xcd, 0x53, 0x98, 0x4d, 0xda,
	0x56, 0x84, 0x0e, 0x61, 0x2b, 0x72, 0xc6, 0x14, 0xe7, 0x29, 0xce, 0x47, 0xce, 0x38, 0x19, 0xf4,
	0x3d, 0xec, 0xb6, 0xae, 0xb1, 0x3d, 0x36, 0x52, 0x13, 0xf4, 0x16, 0xf2, 0x24, 0x8e, 0x44, 0x46,
	0xce, 0x57, 0x77, 0x1a, 0xc7, 0x99, 0x21, 0x8c, 0x78, 0xcd, 0x6e, 0x71, 0xb7, 0x3f, 0x8f, 0x73,
	0x7a, 0xd2, 0x51, 0xf9, 0xce, 0x00, 0xdc, 0x57, 0xd0, 0x1b, 0xe0, 0xfc, 0xe0, 0x0a, 0xd3, 0x4d,
	0xed, 0x34, 0xca, 0x8f, 0x08, 0x25, 0x2b, 0x4d, 0x45, 0x28, 0x1d, 0x59, 0x50, 0xdc, 0x6c, 0x6e,
	0x33, 0x59, 0x24, 0xb2, 0x34, 0xce, 0xc9, 0xff, 0xed, 0xe4, 0x2f, 0x51, 0x64, 0xff, 0x5b, 0x8c,
	0x4e, 0xbe, 0xb1, 0xf0, 0xec, 0x41, 0x0f, 0xaa, 0x80, 0xa4, 0xf4, 0xba, 0x86, 0xde, 0x54, 0x0c,
	0xb3, 0xd7, 0x57, 0xf5, 0xa6, 0xa1, 0xf5, 0xba, 0xe6, 0xb0, 0x3b, 0xe8, 0xab, 0x8a, 0xf6, 0x4e,
	0x53, 0xcf, 0x85, 0x1c, 0x7a, 0x01, 0x72, 0x06, 0x47, 0xeb, 0x0e, 0x8c, 0x66, 0xd7, 0xd0, 0xe8,
	0x4d, 0x60, 0x90, 0x0c, 0x47, 0x19, 0x2c, 0xf5, 0x93, 0xaa, 0x0c, 0x29, 0x83, 0x45, 0x47, 0x20,
	0x66, 0x30, 0x3e, 0x0e, 0x55, 0xfd, 0x42, 0xc8, 0x23, 0x09, 0x4a, 0x19, 0xd5, 0x8e, 0xd6, 0xd6,
	0x9b, 0x86, 0x2a, 0x70, 0xa8, 0x04, 0x07, 0x59, 0x29, 0x5a, 0x8a, 0x50, 0x40, 0x65, 0x38, 0xcc,
	0xa8, 0x0d, 0x86, 0xe7, 0x3d, 0x81, 0x7f, 0xc4, 0x56, 0x57, 0xfb, 0x1f, 0x2e, 0x84, 0xad, 0x56,
	0xe7, 0x76, 0x21, 0x31, 0x77, 0x0b, 0x89, 0xf9, 0xb5, 0x90, 0x98, 0xaf, 0x4b, 0x29, 0x77, 0xb7,
	0x94, 0x72, 0x3f, 0x96, 0x52, 0xee, 0xf3, 0x99, 0xe7, 0x93, 0xd1, 0x97, 0xcb, 0x9a, 0x8d, 0x27,
	0xf5, 0xf4, 0x03, 0xbc, 0x0e, 0x5c, 0x72, 0x83, 0xc3, 0xf1, 0xfa, 0x5e, 0x8f, 0xef, 0xdf, 0x71,
	0xf2, 0x13, 0x47, 0x97, 0x3c, 0x7d, 0x97, 0x67, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xe4,
	0x67, 0xe4, 0xe8, 0x03, 0x00, 0x00,
}

func (m *TxInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalGas != 0 {
		i = encodeVarintTracking(dAtA, i, uint64(m.TotalGas))
		i--
		dAtA[i] = 0x18
	}
	if m.Height != 0 {
		i = encodeVarintTracking(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintTracking(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ContractOperationInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractOperationInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractOperationInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SdkGas != 0 {
		i = encodeVarintTracking(dAtA, i, uint64(m.SdkGas))
		i--
		dAtA[i] = 0x30
	}
	if m.VmGas != 0 {
		i = encodeVarintTracking(dAtA, i, uint64(m.VmGas))
		i--
		dAtA[i] = 0x28
	}
	if m.OperationType != 0 {
		i = encodeVarintTracking(dAtA, i, uint64(m.OperationType))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintTracking(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.TxId != 0 {
		i = encodeVarintTracking(dAtA, i, uint64(m.TxId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintTracking(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BlockTracking) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockTracking) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockTracking) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for iNdEx := len(m.Txs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Txs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTracking(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *TxTracking) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxTracking) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxTracking) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractOperations) > 0 {
		for iNdEx := len(m.ContractOperations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContractOperations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTracking(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTracking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTracking(dAtA []byte, offset int, v uint64) int {
	offset -= sovTracking(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TxInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTracking(uint64(m.Id))
	}
	if m.Height != 0 {
		n += 1 + sovTracking(uint64(m.Height))
	}
	if m.TotalGas != 0 {
		n += 1 + sovTracking(uint64(m.TotalGas))
	}
	return n
}

func (m *ContractOperationInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTracking(uint64(m.Id))
	}
	if m.TxId != 0 {
		n += 1 + sovTracking(uint64(m.TxId))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovTracking(uint64(l))
	}
	if m.OperationType != 0 {
		n += 1 + sovTracking(uint64(m.OperationType))
	}
	if m.VmGas != 0 {
		n += 1 + sovTracking(uint64(m.VmGas))
	}
	if m.SdkGas != 0 {
		n += 1 + sovTracking(uint64(m.SdkGas))
	}
	return n
}

func (m *BlockTracking) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for _, e := range m.Txs {
			l = e.Size()
			n += 1 + l + sovTracking(uint64(l))
		}
	}
	return n
}

func (m *TxTracking) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Info.Size()
	n += 1 + l + sovTracking(uint64(l))
	if len(m.ContractOperations) > 0 {
		for _, e := range m.ContractOperations {
			l = e.Size()
			n += 1 + l + sovTracking(uint64(l))
		}
	}
	return n
}

func sovTracking(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTracking(x uint64) (n int) {
	return sovTracking(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TxInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracking
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
			return fmt.Errorf("proto: TxInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalGas", wireType)
			}
			m.TotalGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTracking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTracking
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
func (m *ContractOperationInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracking
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
			return fmt.Errorf("proto: ContractOperationInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractOperationInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxId", wireType)
			}
			m.TxId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracking
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
				return ErrInvalidLengthTracking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTracking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperationType", wireType)
			}
			m.OperationType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OperationType |= ContractOperation(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VmGas", wireType)
			}
			m.VmGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VmGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SdkGas", wireType)
			}
			m.SdkGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SdkGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTracking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTracking
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
func (m *BlockTracking) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracking
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
			return fmt.Errorf("proto: BlockTracking: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockTracking: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracking
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
				return ErrInvalidLengthTracking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTracking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txs = append(m.Txs, TxTracking{})
			if err := m.Txs[len(m.Txs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTracking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTracking
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
func (m *TxTracking) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTracking
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
			return fmt.Errorf("proto: TxTracking: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxTracking: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracking
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
				return ErrInvalidLengthTracking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTracking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractOperations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTracking
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
				return ErrInvalidLengthTracking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTracking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractOperations = append(m.ContractOperations, ContractOperationInfo{})
			if err := m.ContractOperations[len(m.ContractOperations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTracking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTracking
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
func skipTracking(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTracking
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
					return 0, ErrIntOverflowTracking
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
					return 0, ErrIntOverflowTracking
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
				return 0, ErrInvalidLengthTracking
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTracking
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTracking
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTracking        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTracking          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTracking = fmt.Errorf("proto: unexpected end of group")
)
