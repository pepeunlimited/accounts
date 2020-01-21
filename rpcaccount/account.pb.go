// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package rpcaccount

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateDepositParams struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	AccountType          string   `protobuf:"bytes,3,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDepositParams) Reset()         { *m = CreateDepositParams{} }
func (m *CreateDepositParams) String() string { return proto.CompactTextString(m) }
func (*CreateDepositParams) ProtoMessage()    {}
func (*CreateDepositParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

func (m *CreateDepositParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDepositParams.Unmarshal(m, b)
}
func (m *CreateDepositParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDepositParams.Marshal(b, m, deterministic)
}
func (m *CreateDepositParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDepositParams.Merge(m, src)
}
func (m *CreateDepositParams) XXX_Size() int {
	return xxx_messageInfo_CreateDepositParams.Size(m)
}
func (m *CreateDepositParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDepositParams.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDepositParams proto.InternalMessageInfo

func (m *CreateDepositParams) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CreateDepositParams) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *CreateDepositParams) GetAccountType() string {
	if m != nil {
		return m.AccountType
	}
	return ""
}

type CreateWithdrawParams struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateWithdrawParams) Reset()         { *m = CreateWithdrawParams{} }
func (m *CreateWithdrawParams) String() string { return proto.CompactTextString(m) }
func (*CreateWithdrawParams) ProtoMessage()    {}
func (*CreateWithdrawParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}

func (m *CreateWithdrawParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWithdrawParams.Unmarshal(m, b)
}
func (m *CreateWithdrawParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWithdrawParams.Marshal(b, m, deterministic)
}
func (m *CreateWithdrawParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWithdrawParams.Merge(m, src)
}
func (m *CreateWithdrawParams) XXX_Size() int {
	return xxx_messageInfo_CreateWithdrawParams.Size(m)
}
func (m *CreateWithdrawParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWithdrawParams.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWithdrawParams proto.InternalMessageInfo

func (m *CreateWithdrawParams) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CreateWithdrawParams) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type CreateTransferParams struct {
	FromUserId           int64    `protobuf:"varint,1,opt,name=from_user_id,json=fromUserId,proto3" json:"from_user_id,omitempty"`
	FromAmount           int64    `protobuf:"varint,2,opt,name=from_amount,json=fromAmount,proto3" json:"from_amount,omitempty"`
	ToUserId             int64    `protobuf:"varint,3,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"`
	ToAmount             int64    `protobuf:"varint,4,opt,name=to_amount,json=toAmount,proto3" json:"to_amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTransferParams) Reset()         { *m = CreateTransferParams{} }
func (m *CreateTransferParams) String() string { return proto.CompactTextString(m) }
func (*CreateTransferParams) ProtoMessage()    {}
func (*CreateTransferParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{2}
}

func (m *CreateTransferParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTransferParams.Unmarshal(m, b)
}
func (m *CreateTransferParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTransferParams.Marshal(b, m, deterministic)
}
func (m *CreateTransferParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTransferParams.Merge(m, src)
}
func (m *CreateTransferParams) XXX_Size() int {
	return xxx_messageInfo_CreateTransferParams.Size(m)
}
func (m *CreateTransferParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTransferParams.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTransferParams proto.InternalMessageInfo

func (m *CreateTransferParams) GetFromUserId() int64 {
	if m != nil {
		return m.FromUserId
	}
	return 0
}

func (m *CreateTransferParams) GetFromAmount() int64 {
	if m != nil {
		return m.FromAmount
	}
	return 0
}

func (m *CreateTransferParams) GetToUserId() int64 {
	if m != nil {
		return m.ToUserId
	}
	return 0
}

func (m *CreateTransferParams) GetToAmount() int64 {
	if m != nil {
		return m.ToAmount
	}
	return 0
}

type CreateTransferResponse struct {
	From                 *Account `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   *Account `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTransferResponse) Reset()         { *m = CreateTransferResponse{} }
func (m *CreateTransferResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTransferResponse) ProtoMessage()    {}
func (*CreateTransferResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{3}
}

func (m *CreateTransferResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTransferResponse.Unmarshal(m, b)
}
func (m *CreateTransferResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTransferResponse.Marshal(b, m, deterministic)
}
func (m *CreateTransferResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTransferResponse.Merge(m, src)
}
func (m *CreateTransferResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTransferResponse.Size(m)
}
func (m *CreateTransferResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTransferResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTransferResponse proto.InternalMessageInfo

func (m *CreateTransferResponse) GetFrom() *Account {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *CreateTransferResponse) GetTo() *Account {
	if m != nil {
		return m.To
	}
	return nil
}

type GetAccountsParams struct {
	AccountType          *wrappers.StringValue `protobuf:"bytes,1,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`
	UserId               int64                 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetAccountsParams) Reset()         { *m = GetAccountsParams{} }
func (m *GetAccountsParams) String() string { return proto.CompactTextString(m) }
func (*GetAccountsParams) ProtoMessage()    {}
func (*GetAccountsParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{4}
}

func (m *GetAccountsParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountsParams.Unmarshal(m, b)
}
func (m *GetAccountsParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountsParams.Marshal(b, m, deterministic)
}
func (m *GetAccountsParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountsParams.Merge(m, src)
}
func (m *GetAccountsParams) XXX_Size() int {
	return xxx_messageInfo_GetAccountsParams.Size(m)
}
func (m *GetAccountsParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountsParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountsParams proto.InternalMessageInfo

func (m *GetAccountsParams) GetAccountType() *wrappers.StringValue {
	if m != nil {
		return m.AccountType
	}
	return nil
}

func (m *GetAccountsParams) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CreateAccountParams struct {
	AccountType          string   `protobuf:"bytes,1,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountParams) Reset()         { *m = CreateAccountParams{} }
func (m *CreateAccountParams) String() string { return proto.CompactTextString(m) }
func (*CreateAccountParams) ProtoMessage()    {}
func (*CreateAccountParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{5}
}

func (m *CreateAccountParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountParams.Unmarshal(m, b)
}
func (m *CreateAccountParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountParams.Marshal(b, m, deterministic)
}
func (m *CreateAccountParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountParams.Merge(m, src)
}
func (m *CreateAccountParams) XXX_Size() int {
	return xxx_messageInfo_CreateAccountParams.Size(m)
}
func (m *CreateAccountParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountParams.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountParams proto.InternalMessageInfo

func (m *CreateAccountParams) GetAccountType() string {
	if m != nil {
		return m.AccountType
	}
	return ""
}

func (m *CreateAccountParams) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type GetAccountsResponse struct {
	Accounts             []*Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
	UserId               int64      `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetAccountsResponse) Reset()         { *m = GetAccountsResponse{} }
func (m *GetAccountsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAccountsResponse) ProtoMessage()    {}
func (*GetAccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{6}
}

func (m *GetAccountsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountsResponse.Unmarshal(m, b)
}
func (m *GetAccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountsResponse.Marshal(b, m, deterministic)
}
func (m *GetAccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountsResponse.Merge(m, src)
}
func (m *GetAccountsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAccountsResponse.Size(m)
}
func (m *GetAccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountsResponse proto.InternalMessageInfo

func (m *GetAccountsResponse) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *GetAccountsResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type GetAccountParams struct {
	AccountId            int64    `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountParams) Reset()         { *m = GetAccountParams{} }
func (m *GetAccountParams) String() string { return proto.CompactTextString(m) }
func (*GetAccountParams) ProtoMessage()    {}
func (*GetAccountParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{7}
}

func (m *GetAccountParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountParams.Unmarshal(m, b)
}
func (m *GetAccountParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountParams.Marshal(b, m, deterministic)
}
func (m *GetAccountParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountParams.Merge(m, src)
}
func (m *GetAccountParams) XXX_Size() int {
	return xxx_messageInfo_GetAccountParams.Size(m)
}
func (m *GetAccountParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountParams proto.InternalMessageInfo

func (m *GetAccountParams) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *GetAccountParams) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type Account struct {
	Balance              int64    `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	UserId               int64    `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Id                   int64    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{8}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Account) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Account) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Account) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateDepositParams)(nil), "pepeunlimited.accounts.CreateDepositParams")
	proto.RegisterType((*CreateWithdrawParams)(nil), "pepeunlimited.accounts.CreateWithdrawParams")
	proto.RegisterType((*CreateTransferParams)(nil), "pepeunlimited.accounts.CreateTransferParams")
	proto.RegisterType((*CreateTransferResponse)(nil), "pepeunlimited.accounts.CreateTransferResponse")
	proto.RegisterType((*GetAccountsParams)(nil), "pepeunlimited.accounts.GetAccountsParams")
	proto.RegisterType((*CreateAccountParams)(nil), "pepeunlimited.accounts.CreateAccountParams")
	proto.RegisterType((*GetAccountsResponse)(nil), "pepeunlimited.accounts.GetAccountsResponse")
	proto.RegisterType((*GetAccountParams)(nil), "pepeunlimited.accounts.GetAccountParams")
	proto.RegisterType((*Account)(nil), "pepeunlimited.accounts.Account")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5d, 0x6f, 0x12, 0x4d,
	0x14, 0x0e, 0x0b, 0x81, 0x72, 0xa0, 0xe4, 0x7d, 0xa7, 0x06, 0x09, 0xad, 0x96, 0xee, 0x15, 0xa6,
	0x66, 0x49, 0xda, 0x4b, 0x2f, 0x4c, 0xd5, 0xa4, 0xa9, 0x57, 0x4a, 0xab, 0x4d, 0x4c, 0x0c, 0x0e,
	0xec, 0x01, 0x37, 0xb2, 0x3b, 0x93, 0xd9, 0xc1, 0x86, 0x1b, 0x7f, 0x85, 0xff, 0xcf, 0xbf, 0x62,
	0x98, 0x8f, 0xed, 0x4e, 0x05, 0x77, 0xe3, 0xdd, 0xce, 0xf9, 0x78, 0x9e, 0xf3, 0xf1, 0x9c, 0x85,
	0x7d, 0x3a, 0x9b, 0xb1, 0x55, 0x22, 0x03, 0x2e, 0x98, 0x64, 0xa4, 0xcb, 0x91, 0xe3, 0x2a, 0x59,
	0x46, 0x71, 0x24, 0x31, 0x0c, 0x8c, 0x33, 0xed, 0x1f, 0x2e, 0x18, 0x5b, 0x2c, 0x71, 0xa4, 0xa2,
	0xa6, 0xab, 0xf9, 0x08, 0x63, 0x2e, 0xd7, 0x3a, 0xa9, 0xff, 0xf4, 0xa1, 0xf3, 0x4e, 0x50, 0xce,
	0x51, 0xa4, 0xda, 0xef, 0x47, 0x70, 0xf0, 0x5a, 0x20, 0x95, 0xf8, 0x06, 0x39, 0x4b, 0x23, 0xf9,
	0x8e, 0x0a, 0x1a, 0xa7, 0xe4, 0x31, 0x34, 0x56, 0x29, 0x8a, 0x49, 0x14, 0xf6, 0x2a, 0x83, 0xca,
	0xb0, 0x3a, 0xae, 0x6f, 0x9e, 0x57, 0x21, 0xe9, 0x42, 0x9d, 0xc6, 0x1b, 0xde, 0x9e, 0xa7, 0xed,
	0xfa, 0x45, 0x4e, 0xa0, 0x6d, 0x0a, 0x9a, 0xc8, 0x35, 0xc7, 0x5e, 0x75, 0x50, 0x19, 0x36, 0xc7,
	0x2d, 0x63, 0xbb, 0x59, 0x73, 0xf4, 0x2f, 0xe1, 0x91, 0xa6, 0xba, 0x8d, 0xe4, 0xd7, 0x50, 0xd0,
	0xbb, 0x7f, 0xe4, 0xf2, 0x7f, 0x56, 0x2c, 0xd2, 0x8d, 0xa0, 0x49, 0x3a, 0x47, 0x61, 0x90, 0x06,
	0xd0, 0x9e, 0x0b, 0x16, 0x4f, 0x5c, 0x38, 0xd8, 0xd8, 0x3e, 0x68, 0xc8, 0x63, 0x68, 0xa9, 0x08,
	0x07, 0x57, 0x05, 0x5c, 0xe8, 0x3e, 0x8e, 0x00, 0x24, 0xcb, 0x00, 0xaa, 0xca, 0xbf, 0x27, 0x99,
	0x49, 0x3f, 0x84, 0xa6, 0x64, 0x36, 0xb9, 0x66, 0x9d, 0x3a, 0xd5, 0xff, 0x01, 0x5d, 0xb7, 0xaa,
	0x31, 0xa6, 0x9c, 0x25, 0x29, 0x92, 0x73, 0xa8, 0x6d, 0x28, 0x54, 0x3d, 0xad, 0xb3, 0xe3, 0x60,
	0xfb, 0x22, 0x83, 0x0b, 0xfd, 0x31, 0x56, 0xc1, 0x64, 0x04, 0x9e, 0x64, 0xaa, 0xc2, 0x12, 0x29,
	0x9e, 0x64, 0x7e, 0x0c, 0xff, 0x5f, 0xa2, 0x34, 0x96, 0xd4, 0x8c, 0xe4, 0xe5, 0x83, 0xbd, 0xe8,
	0x12, 0x8e, 0x02, 0x2d, 0x8b, 0xc0, 0xca, 0x22, 0xb8, 0x96, 0x22, 0x4a, 0x16, 0x1f, 0xe9, 0x72,
	0x85, 0xce, 0xd6, 0xf2, 0xdb, 0xf1, 0xf2, 0xdb, 0xf1, 0xdf, 0x5b, 0xe5, 0x18, 0x46, 0x43, 0x78,
	0xb2, 0x85, 0xb0, 0x59, 0x12, 0xf2, 0x1b, 0x1c, 0xe4, 0x3a, 0xc8, 0xc6, 0xf7, 0x02, 0xf6, 0x6c,
	0xc3, 0xbd, 0xca, 0xa0, 0x5a, 0x66, 0x1e, 0x59, 0xc2, 0x6e, 0xb2, 0xb7, 0xf0, 0xdf, 0x3d, 0x99,
	0x29, 0xfe, 0x09, 0x80, 0x2d, 0x3e, 0x93, 0x4f, 0xd3, 0x58, 0xae, 0xc2, 0xdd, 0x58, 0x5f, 0xa0,
	0x61, 0x80, 0x48, 0x0f, 0x1a, 0x53, 0xba, 0xa4, 0xc9, 0x0c, 0x4d, 0xbe, 0x7d, 0x12, 0x02, 0x35,
	0x35, 0x11, 0x4f, 0x4d, 0x44, 0x7d, 0xe7, 0x11, 0xab, 0x8e, 0xf6, 0x3b, 0xe0, 0x45, 0xa1, 0x91,
	0x98, 0x17, 0x85, 0x67, 0xbf, 0x6a, 0xd0, 0x31, 0x14, 0xd7, 0x28, 0xbe, 0x47, 0x33, 0x24, 0x9f,
	0x61, 0xdf, 0x39, 0x5d, 0x72, 0xba, 0x6b, 0x2a, 0x5b, 0x2e, 0xbc, 0x5f, 0x34, 0x42, 0x32, 0x81,
	0x8e, 0x7b, 0xae, 0xe4, 0xf9, 0xdf, 0xf1, 0xdd, 0xb3, 0x2e, 0x26, 0x48, 0x2c, 0x81, 0xbd, 0x97,
	0x22, 0x02, 0xf7, 0xda, 0xfb, 0x41, 0xb9, 0xe8, 0x4c, 0x46, 0xd9, 0xbc, 0x6c, 0x01, 0x05, 0xf3,
	0x72, 0xa4, 0x51, 0xdc, 0x0e, 0x42, 0x2b, 0x27, 0x5e, 0xf2, 0x6c, 0x57, 0xfc, 0x1f, 0x37, 0xda,
	0x3f, 0x2d, 0x11, 0x9a, 0x75, 0x71, 0x0b, 0x70, 0x6f, 0x26, 0xc3, 0xe2, 0xd4, 0x92, 0xf5, 0xbf,
	0x6a, 0x7f, 0x02, 0xc1, 0x67, 0xc6, 0x3c, 0xad, 0xab, 0x3f, 0xc3, 0xf9, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xfd, 0x7c, 0xfc, 0x87, 0x84, 0x06, 0x00, 0x00,
}
