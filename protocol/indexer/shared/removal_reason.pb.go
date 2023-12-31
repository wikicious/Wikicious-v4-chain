// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dydxprotocol/indexer/shared/removal_reason.proto

package shared

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// OrderRemovalReason is an enum of all the reasons an order was removed.
type OrderRemovalReason int32

const (
	// Default value, this is invalid and unused.
	OrderRemovalReason_ORDER_REMOVAL_REASON_UNSPECIFIED OrderRemovalReason = 0
	// The order was removed due to being expired.
	OrderRemovalReason_ORDER_REMOVAL_REASON_EXPIRED OrderRemovalReason = 1
	// The order was removed due to being canceled by a user.
	OrderRemovalReason_ORDER_REMOVAL_REASON_USER_CANCELED OrderRemovalReason = 2
	// The order was removed due to being undercollateralized.
	OrderRemovalReason_ORDER_REMOVAL_REASON_UNDERCOLLATERALIZED OrderRemovalReason = 3
	// The order caused an internal error during order placement and was
	// removed.
	OrderRemovalReason_ORDER_REMOVAL_REASON_INTERNAL_ERROR OrderRemovalReason = 4
	// The order would have matched against another order placed by the same
	// subaccount and was removed.
	OrderRemovalReason_ORDER_REMOVAL_REASON_SELF_TRADE_ERROR OrderRemovalReason = 5
	// The order would have matched against maker orders on the orderbook
	// despite being a post-only order and was removed.
	OrderRemovalReason_ORDER_REMOVAL_REASON_POST_ONLY_WOULD_CROSS_MAKER_ORDER OrderRemovalReason = 6
	// The order was an ICO order and would have been placed on the orderbook as
	// resting liquidity and was removed.
	OrderRemovalReason_ORDER_REMOVAL_REASON_IMMEDIATE_OR_CANCEL_WOULD_REST_ON_BOOK OrderRemovalReason = 7
	// The order was a fill-or-kill order that could not be fully filled and was
	// removed.
	OrderRemovalReason_ORDER_REMOVAL_REASON_FOK_ORDER_COULD_NOT_BE_FULLY_FULLED OrderRemovalReason = 8
	// The order was a reduce-only order that was removed due to either:
	// - being a taker order and fully-filling the order would flip the side of
	//    the subaccount's position, in this case the remaining size of the
	//    order is removed
	// - being a maker order resting on the book and being removed when either
	//    the subaccount's position is closed or flipped sides
	OrderRemovalReason_ORDER_REMOVAL_REASON_REDUCE_ONLY_RESIZE OrderRemovalReason = 9
	// The order should be expired, according to the Indexer's cached data, but
	// the Indexer has yet to receive a message to remove the order. In order to
	// keep the data cached by the Indexer up-to-date and accurate, clear out
	// the data if it's expired by sending an order removal with this reason.
	// Protocol should never send this reason to Indexer.
	OrderRemovalReason_ORDER_REMOVAL_REASON_INDEXER_EXPIRED OrderRemovalReason = 10
	// The order has been replaced.
	OrderRemovalReason_ORDER_REMOVAL_REASON_REPLACED OrderRemovalReason = 11
	// The order has been fully-filled. Only sent by the Indexer for stateful
	// orders.
	OrderRemovalReason_ORDER_REMOVAL_REASON_FULLY_FILLED OrderRemovalReason = 12
	// The order has been removed since the subaccount does not satisfy the
	// equity tier requirements.
	OrderRemovalReason_ORDER_REMOVAL_REASON_EQUITY_TIER OrderRemovalReason = 13
)

var OrderRemovalReason_name = map[int32]string{
	0:  "ORDER_REMOVAL_REASON_UNSPECIFIED",
	1:  "ORDER_REMOVAL_REASON_EXPIRED",
	2:  "ORDER_REMOVAL_REASON_USER_CANCELED",
	3:  "ORDER_REMOVAL_REASON_UNDERCOLLATERALIZED",
	4:  "ORDER_REMOVAL_REASON_INTERNAL_ERROR",
	5:  "ORDER_REMOVAL_REASON_SELF_TRADE_ERROR",
	6:  "ORDER_REMOVAL_REASON_POST_ONLY_WOULD_CROSS_MAKER_ORDER",
	7:  "ORDER_REMOVAL_REASON_IMMEDIATE_OR_CANCEL_WOULD_REST_ON_BOOK",
	8:  "ORDER_REMOVAL_REASON_FOK_ORDER_COULD_NOT_BE_FULLY_FULLED",
	9:  "ORDER_REMOVAL_REASON_REDUCE_ONLY_RESIZE",
	10: "ORDER_REMOVAL_REASON_INDEXER_EXPIRED",
	11: "ORDER_REMOVAL_REASON_REPLACED",
	12: "ORDER_REMOVAL_REASON_FULLY_FILLED",
	13: "ORDER_REMOVAL_REASON_EQUITY_TIER",
}

var OrderRemovalReason_value = map[string]int32{
	"ORDER_REMOVAL_REASON_UNSPECIFIED":                            0,
	"ORDER_REMOVAL_REASON_EXPIRED":                                1,
	"ORDER_REMOVAL_REASON_USER_CANCELED":                          2,
	"ORDER_REMOVAL_REASON_UNDERCOLLATERALIZED":                    3,
	"ORDER_REMOVAL_REASON_INTERNAL_ERROR":                         4,
	"ORDER_REMOVAL_REASON_SELF_TRADE_ERROR":                       5,
	"ORDER_REMOVAL_REASON_POST_ONLY_WOULD_CROSS_MAKER_ORDER":      6,
	"ORDER_REMOVAL_REASON_IMMEDIATE_OR_CANCEL_WOULD_REST_ON_BOOK": 7,
	"ORDER_REMOVAL_REASON_FOK_ORDER_COULD_NOT_BE_FULLY_FULLED":    8,
	"ORDER_REMOVAL_REASON_REDUCE_ONLY_RESIZE":                     9,
	"ORDER_REMOVAL_REASON_INDEXER_EXPIRED":                        10,
	"ORDER_REMOVAL_REASON_REPLACED":                               11,
	"ORDER_REMOVAL_REASON_FULLY_FILLED":                           12,
	"ORDER_REMOVAL_REASON_EQUITY_TIER":                            13,
}

func (x OrderRemovalReason) String() string {
	return proto.EnumName(OrderRemovalReason_name, int32(x))
}

func (OrderRemovalReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d5eea5cab8c58ba, []int{0}
}

func init() {
	proto.RegisterEnum("dydxprotocol.indexer.shared.OrderRemovalReason", OrderRemovalReason_name, OrderRemovalReason_value)
}

func init() {
	proto.RegisterFile("dydxprotocol/indexer/shared/removal_reason.proto", fileDescriptor_0d5eea5cab8c58ba)
}

var fileDescriptor_0d5eea5cab8c58ba = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x13, 0xa0, 0x05, 0x16, 0x90, 0x56, 0x7b, 0x05, 0xac, 0x16, 0x5a, 0x5a, 0xfe, 0xc5,
	0x48, 0x20, 0x54, 0x01, 0x12, 0xda, 0x78, 0x27, 0xd2, 0x2a, 0x1b, 0x6f, 0x18, 0xdb, 0xd0, 0xe4,
	0x32, 0x72, 0x63, 0x8b, 0x44, 0x6a, 0x63, 0xe4, 0x96, 0xaa, 0xbc, 0x01, 0x47, 0x1e, 0x8b, 0x63,
	0x8e, 0x1c, 0x51, 0xf2, 0x22, 0x48, 0x76, 0x40, 0x54, 0x5a, 0x5f, 0x7c, 0xf1, 0xef, 0xf7, 0xcd,
	0x67, 0xef, 0x2c, 0x7b, 0x91, 0x7d, 0xcb, 0x2e, 0xbe, 0x94, 0xc5, 0x59, 0x31, 0x29, 0x8e, 0xfd,
	0xd9, 0x3c, 0xcb, 0x2f, 0xf2, 0xd2, 0x3f, 0x9d, 0xa6, 0x65, 0x9e, 0xf9, 0x65, 0x7e, 0x52, 0x9c,
	0xa7, 0xc7, 0x54, 0xe6, 0xe9, 0x69, 0x31, 0xef, 0x54, 0x98, 0xb8, 0xfb, 0xbf, 0xd1, 0x59, 0x1b,
	0x9d, 0xda, 0x78, 0xf2, 0x7d, 0x83, 0x09, 0x5b, 0x66, 0x79, 0x89, 0xb5, 0x8a, 0x95, 0x29, 0x76,
	0xd8, 0x96, 0x45, 0x05, 0x48, 0x08, 0x03, 0xfb, 0x51, 0x1a, 0x42, 0x90, 0x91, 0x0d, 0x29, 0x09,
	0xa3, 0x21, 0x04, 0xba, 0xa7, 0x41, 0xf1, 0x96, 0xd8, 0x62, 0xf7, 0x9c, 0x14, 0x1c, 0x0e, 0x35,
	0x82, 0xe2, 0x6d, 0xf1, 0x88, 0x3d, 0x70, 0xe7, 0x44, 0x80, 0x14, 0xc8, 0x30, 0x00, 0x03, 0x8a,
	0x5f, 0x11, 0xcf, 0xd8, 0x7e, 0xc3, 0x3c, 0x05, 0x18, 0x58, 0x63, 0x64, 0x0c, 0x28, 0x8d, 0x1e,
	0x83, 0xe2, 0x57, 0xc5, 0x1e, 0x7b, 0xe8, 0xa4, 0x75, 0x18, 0x03, 0x86, 0xd2, 0x10, 0x20, 0x5a,
	0xe4, 0xd7, 0xc4, 0x63, 0xb6, 0xeb, 0x04, 0x23, 0x30, 0x3d, 0x8a, 0x51, 0x2a, 0x58, 0xa3, 0x1b,
	0xe2, 0x0d, 0x7b, 0xed, 0x44, 0x87, 0x36, 0x8a, 0xc9, 0x86, 0x66, 0x44, 0x9f, 0x6c, 0x62, 0x14,
	0x05, 0x68, 0xa3, 0x88, 0x06, 0xb2, 0x0f, 0x48, 0x95, 0xc0, 0x37, 0xc5, 0x7b, 0xf6, 0xd6, 0xdd,
	0x67, 0x30, 0x00, 0xa5, 0x65, 0x0c, 0x64, 0xff, 0x7e, 0xed, 0x3a, 0x05, 0xa1, 0x4a, 0xa5, 0xae,
	0xb5, 0x7d, 0x7e, 0x5d, 0xbc, 0x63, 0x07, 0xce, 0x80, 0x9e, 0xed, 0xd7, 0x43, 0x28, 0xa8, 0xb4,
	0xd0, 0xc6, 0xd4, 0x05, 0xea, 0x25, 0xc6, 0x8c, 0xaa, 0x27, 0x28, 0x7e, 0x43, 0x3c, 0x65, 0x7b,
	0x4e, 0x1b, 0x41, 0x25, 0x01, 0xd4, 0xe5, 0x11, 0x22, 0x3d, 0x06, 0x7e, 0x53, 0xec, 0xb3, 0x9d,
	0x86, 0x7f, 0xa7, 0xe0, 0x10, 0xf0, 0xdf, 0xd9, 0x31, 0xb1, 0xcd, 0xee, 0x37, 0xc4, 0x0e, 0x8d,
	0x0c, 0x40, 0xf1, 0x5b, 0x62, 0x97, 0x6d, 0xbb, 0x7b, 0xd7, 0x05, 0x75, 0x55, 0xf0, 0x76, 0xe3,
	0x36, 0xc1, 0x87, 0x44, 0xc7, 0x23, 0x8a, 0x35, 0x20, 0xbf, 0xd3, 0xc5, 0x9f, 0x4b, 0xaf, 0xbd,
	0x58, 0x7a, 0xed, 0xdf, 0x4b, 0xaf, 0xfd, 0x63, 0xe5, 0xb5, 0x16, 0x2b, 0xaf, 0xf5, 0x6b, 0xe5,
	0xb5, 0xc6, 0x07, 0x9f, 0x67, 0x67, 0xd3, 0xaf, 0x47, 0x9d, 0x49, 0x71, 0xe2, 0x5f, 0x5a, 0xff,
	0xf3, 0x57, 0xcf, 0x27, 0xd3, 0x74, 0x36, 0xf7, 0x1b, 0x2e, 0xc4, 0xd1, 0x66, 0xf5, 0xe2, 0xe5,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x84, 0xfa, 0x05, 0x36, 0x03, 0x00, 0x00,
}
