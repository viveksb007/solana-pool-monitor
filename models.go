package main

import (
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
)

type LiquidityStateLayoutV4 struct {
	Status   bin.Uint64
	Nonce    bin.Uint64
	MaxOrder bin.Uint64
	Depth    bin.Uint64

	BaseDecimal  bin.Uint64
	QuoteDecimal bin.Uint64
	State        bin.Uint64
	ResetFlag    bin.Uint64

	MinSize                bin.Uint64
	VolMaxCutRatio         bin.Uint64
	AmountWaveRatio        bin.Uint64
	BaseLotSize            bin.Uint64
	QuoteLotSize           bin.Uint64
	MinPriceMultiplier     bin.Uint64
	MaxPriceMultiplier     bin.Uint64
	SystemDecimalValue     bin.Uint64
	MinSeparateNumerator   bin.Uint64
	MinSeparateDenominator bin.Uint64
	TradeFeeNumerator      bin.Uint64
	TradeFeeDenominator    bin.Uint64
	PnlNumerator           bin.Uint64
	PnlDenominator         bin.Uint64
	SwapFeeNumerator       bin.Uint64
	SwapFeeDenominator     bin.Uint64
	BaseNeedTakePnl        bin.Uint64
	QuoteNeedTakePnl       bin.Uint64

	QuoteTotalPnl       bin.Uint64
	BaseTotalPnl        bin.Uint64
	QuoteTotalDeposited bin.Uint128
	BaseTotalDeposited  bin.Uint128
	SwapBaseInAmount    bin.Uint128
	SwapQuoteOutAmount  bin.Uint128

	SwapBase2QuoteFee bin.Uint64
	SwapQuoteInAmount bin.Uint128
	SwapBaseOutAmount bin.Uint128

	SwapQuote2BaseFee bin.Uint64

	BaseVault       solana.PublicKey
	QuoteVault      solana.PublicKey
	BaseMint        solana.PublicKey
	QuoteMint       solana.PublicKey
	LpMint          solana.PublicKey
	OpenOrders      solana.PublicKey
	MarketId        solana.PublicKey
	MarketProgramId solana.PublicKey
	TargetOrders    solana.PublicKey
	WithdrawQueue   solana.PublicKey
	LpVault         solana.PublicKey
	Owner           solana.PublicKey
	PnlOwner        solana.PublicKey
}
