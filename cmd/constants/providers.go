package constants

import (
	"github.com/warden-protocol/connect/oracle/config"
	"github.com/warden-protocol/connect/oracle/constants"
	"github.com/warden-protocol/connect/oracle/types"
	binanceapi "github.com/warden-protocol/connect/providers/apis/binance"
	bitstampapi "github.com/warden-protocol/connect/providers/apis/bitstamp"
	coinbaseapi "github.com/warden-protocol/connect/providers/apis/coinbase"
	"github.com/warden-protocol/connect/providers/apis/coingecko"
	"github.com/warden-protocol/connect/providers/apis/coinmarketcap"
	"github.com/warden-protocol/connect/providers/apis/defi/osmosis"
	"github.com/warden-protocol/connect/providers/apis/defi/raydium"
	"github.com/warden-protocol/connect/providers/apis/defi/uniswapv3"
	"github.com/warden-protocol/connect/providers/apis/dydx"
	krakenapi "github.com/warden-protocol/connect/providers/apis/kraken"
	"github.com/warden-protocol/connect/providers/apis/marketmap"
	"github.com/warden-protocol/connect/providers/apis/polymarket"
	"github.com/warden-protocol/connect/providers/volatile"
	binancews "github.com/warden-protocol/connect/providers/websockets/binance"
	"github.com/warden-protocol/connect/providers/websockets/bitfinex"
	"github.com/warden-protocol/connect/providers/websockets/bitstamp"
	"github.com/warden-protocol/connect/providers/websockets/bybit"
	"github.com/warden-protocol/connect/providers/websockets/coinbase"
	"github.com/warden-protocol/connect/providers/websockets/cryptodotcom"
	"github.com/warden-protocol/connect/providers/websockets/gate"
	"github.com/warden-protocol/connect/providers/websockets/huobi"
	"github.com/warden-protocol/connect/providers/websockets/kraken"
	"github.com/warden-protocol/connect/providers/websockets/kucoin"
	"github.com/warden-protocol/connect/providers/websockets/mexc"
	"github.com/warden-protocol/connect/providers/websockets/okx"
	mmtypes "github.com/warden-protocol/connect/service/clients/marketmap/types"
)

var (
	Providers = []config.ProviderConfig{
		// DEFI providers
		{
			Name: raydium.Name,
			API:  raydium.DefaultAPIConfig,
			Type: types.ConfigType,
		},
		{
			Name: uniswapv3.ProviderNames[constants.ETHEREUM],
			API:  uniswapv3.DefaultETHAPIConfig,
			Type: types.ConfigType,
		},
		{
			Name: uniswapv3.ProviderNames[constants.BASE],
			API:  uniswapv3.DefaultBaseAPIConfig,
			Type: types.ConfigType,
		},
		{
			Name: osmosis.Name,
			API:  osmosis.DefaultAPIConfig,
			Type: types.ConfigType,
		},

		// Exchange API providers
		{
			Name: binanceapi.Name,
			API:  binanceapi.DefaultNonUSAPIConfig,
			Type: types.ConfigType,
		},
		{
			Name: bitstampapi.Name,
			API:  bitstampapi.DefaultAPIConfig,
			Type: types.ConfigType,
		},
		{
			Name: coinbaseapi.Name,
			API:  coinbaseapi.DefaultAPIConfig,
			Type: types.ConfigType,
		},
		{
			Name: coingecko.Name,
			API:  coingecko.DefaultAPIConfig,
			Type: types.ConfigType,
		},
		{
			Name: coinmarketcap.Name,
			API:  coinmarketcap.DefaultAPIConfig,
			Type: types.ConfigType,
		},
		{
			Name: krakenapi.Name,
			API:  krakenapi.DefaultAPIConfig,
			Type: types.ConfigType,
		},
		{
			Name: volatile.Name,
			API:  volatile.DefaultAPIConfig,
			Type: types.ConfigType,
		},
		// Exchange WebSocket providers
		{
			Name:      binancews.Name,
			WebSocket: binancews.DefaultWebSocketConfig,
			Type:      types.ConfigType,
		},
		{
			Name:      bitfinex.Name,
			WebSocket: bitfinex.DefaultWebSocketConfig,
			Type:      types.ConfigType,
		},
		{
			Name:      bitstamp.Name,
			WebSocket: bitstamp.DefaultWebSocketConfig,
			Type:      types.ConfigType,
		},
		{
			Name:      bybit.Name,
			WebSocket: bybit.DefaultWebSocketConfig,
			Type:      types.ConfigType,
		},
		{
			Name:      coinbase.Name,
			WebSocket: coinbase.DefaultWebSocketConfig,
			Type:      types.ConfigType,
		},
		{
			Name:      cryptodotcom.Name,
			WebSocket: cryptodotcom.DefaultWebSocketConfig,
			Type:      types.ConfigType,
		},
		{
			Name:      gate.Name,
			WebSocket: gate.DefaultWebSocketConfig,
			Type:      types.ConfigType,
		},
		{
			Name:      huobi.Name,
			WebSocket: huobi.DefaultWebSocketConfig,
			Type:      types.ConfigType,
		},
		{
			Name:      kraken.Name,
			WebSocket: kraken.DefaultWebSocketConfig,
			Type:      types.ConfigType,
		},
		{
			Name:      kucoin.Name,
			WebSocket: kucoin.DefaultWebSocketConfig,
			API:       kucoin.DefaultAPIConfig,
			Type:      types.ConfigType,
		},
		{
			Name:      mexc.Name,
			WebSocket: mexc.DefaultWebSocketConfig,
			Type:      types.ConfigType,
		},
		{
			Name:      okx.Name,
			WebSocket: okx.DefaultWebSocketConfig,
			Type:      types.ConfigType,
		},

		// Polymarket provider
		{
			Name: polymarket.Name,
			API:  polymarket.DefaultAPIConfig,
			Type: types.ConfigType,
		},

		// MarketMap provider
		{
			Name: marketmap.Name,
			API:  marketmap.DefaultAPIConfig,
			Type: mmtypes.ConfigType,
		},
	}

	AlternativeMarketMapProviders = []config.ProviderConfig{
		{
			Name: dydx.Name,
			API:  dydx.DefaultAPIConfig,
			Type: mmtypes.ConfigType,
		},
		{
			Name: dydx.SwitchOverAPIHandlerName,
			API:  dydx.DefaultSwitchOverAPIConfig,
			Type: mmtypes.ConfigType,
		},
		{
			Name: dydx.ResearchAPIHandlerName,
			API:  dydx.DefaultResearchAPIConfig,
			Type: mmtypes.ConfigType,
		},
		{
			Name: dydx.ResearchCMCAPIHandlerName,
			API:  dydx.DefaultResearchCMCAPIConfig,
			Type: mmtypes.ConfigType,
		},
	}

	MarketMapProviderNames = map[string]struct{}{
		dydx.Name:                      {},
		dydx.SwitchOverAPIHandlerName:  {},
		dydx.ResearchAPIHandlerName:    {},
		dydx.ResearchCMCAPIHandlerName: {},
		marketmap.Name:                 {},
	}
)
