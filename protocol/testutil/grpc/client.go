package grpc

import (
	bridgetypes "github.com/dydxprotocol/v4-chain/protocol/daemons/bridge/api"
	liquidationtypes "github.com/dydxprotocol/v4-chain/protocol/daemons/liquidation/api"
	pricefeedtypes "github.com/dydxprotocol/v4-chain/protocol/daemons/pricefeed/api"
	clobtypes "github.com/dydxprotocol/v4-chain/protocol/x/clob/types"
	pricetypes "github.com/dydxprotocol/v4-chain/protocol/x/prices/types"
	satypes "github.com/dydxprotocol/v4-chain/protocol/x/subaccounts/types"
)

// QueryClient combines all the query clients used in testing into a single mock interface for testing convenience.
type QueryClient interface {
	satypes.QueryClient
	clobtypes.QueryClient
	bridgetypes.BridgeServiceClient
	liquidationtypes.LiquidationServiceClient
	pricefeedtypes.PriceFeedServiceClient
	pricetypes.QueryClient
}
