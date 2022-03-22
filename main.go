package main

import (
	"context"
	"fmt"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/serum"
	"github.com/gagliardetto/solana-go/rpc"
	"math"
	"sync"

	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func main() {

	rpcClient := rpc.New("https://ssc-dao.genesysgo.net")

	printRaydiumPoolReserve("58oQChx4yWmvKdwLLZzBi4ChoCc2fqCUWBkwMihLYQo2", rpcClient)

	//client, err := ws.Connect(context.Background(), "wss://ssc-dao.genesysgo.net")
	//if err != nil {
	//	panic(err)
	//}
	//
	//go monitorPool("ANP74VNsHwSrq9uUSjiSNyNWvf6ZPrKTmE4gHoNd13Lg", "75HgnSvXbWKZBpZHveX68ZzAhDqMzNDS29X6BGLtxMo1", client)
	//go monitorPool("DQyrAcCrDXQ7NeoqGgDCZwBvWDcYmFCjSb9JtteuvPpz", "HLmqeL62xR1QoZ1HKKbXRrdN1p3phKpxRMb2VVopvBBz", client)
	//select {}
}

// TODO - write monitor Raydium Pool method

func printRaydiumPoolReserve(poolId string, rpcClient *rpc.Client) {
	ctx := context.TODO()
	info, err := rpcClient.GetAccountInfo(ctx, solana.MustPublicKeyFromBase58(poolId))
	if err != nil {
		panic(err)
	}

	borshDec := bin.NewBorshDecoder(info.Value.Data.GetBinary())
	var liquidityPool LiquidityStateLayoutV4
	err = borshDec.Decode(&liquidityPool)
	if err != nil {
		panic(err)
	}
	fmt.Println(liquidityPool)

	openOrders := liquidityPool.OpenOrders
	baseVault := liquidityPool.BaseVault
	quoteVault := liquidityPool.QuoteVault
	basePnl := liquidityPool.BaseNeedTakePnl
	quotePnl := liquidityPool.QuoteNeedTakePnl
	fmt.Println(openOrders)
	fmt.Println(fmt.Sprintf("Base Pnl %d", basePnl))
	fmt.Println(fmt.Sprintf("Quote Pnl %d", quotePnl))

	baseVaultBalance, err := rpcClient.GetTokenAccountBalance(ctx, baseVault, rpc.CommitmentConfirmed)
	if err != nil {
		panic(err)
	}
	baseTokenAmount := *baseVaultBalance.Value.UiAmount
	fmt.Println(fmt.Sprintf("Base token amount %s : %f", liquidityPool.BaseVault, baseTokenAmount))

	quoteVaultBalance, err := rpcClient.GetTokenAccountBalance(ctx, quoteVault, rpc.CommitmentConfirmed)
	if err != nil {
		panic(err)
	}
	quoteTokenAmount := *quoteVaultBalance.Value.UiAmount
	fmt.Println(fmt.Sprintf("Quote token amount %s : %f", liquidityPool.QuoteVault, quoteTokenAmount))

	serumOpenOrders := getSerumOpenOrder(openOrders, rpcClient)
	fmt.Println(serumOpenOrders.NativeBaseTokenTotal)
	fmt.Println(serumOpenOrders.NativeQuoteTokenTotal)

	base := baseTokenAmount + (float64(serumOpenOrders.NativeBaseTokenTotal)-float64(basePnl))/SOL
	quote := quoteTokenAmount + (float64(serumOpenOrders.NativeQuoteTokenTotal)-float64(quotePnl))/USDC

	fmt.Println(fmt.Sprintf("BASE %f", base))
	fmt.Println(fmt.Sprintf("QUOTE %f", quote))
}

func getSerumOpenOrder(pubKey solana.PublicKey, rpcClient *rpc.Client) *serum.OpenOrders {
	info, err := rpcClient.GetAccountInfo(context.TODO(), pubKey)
	if err != nil {
		panic(err)
	}

	borshDec := bin.NewBorshDecoder(info.Value.Data.GetBinary())
	var meta serum.OpenOrders
	err = borshDec.Decode(&meta)
	if err != nil {
		panic(err)
	}
	return &meta
}

type PoolUpdateEvent struct {
	mu      sync.Mutex
	amounts [2]uint64
	slots   [2]uint64
}

func monitorPool(pair1 string, pair2 string, client *ws.Client) {
	pair1PK := solana.MustPublicKeyFromBase58(pair1)
	pair2PK := solana.MustPublicKeyFromBase58(pair2)

	poolUpdateEvent := PoolUpdateEvent{}

	go listenToAddress(pair1PK, client, &poolUpdateEvent, 0)
	go listenToAddress(pair2PK, client, &poolUpdateEvent, 1)

	select {}

}

var (
	SOL  = math.Pow10(9)
	USDC = math.Pow10(6)
	RAY  = math.Pow10(6)
)

func listenToAddress(publicKey solana.PublicKey, client *ws.Client, p *PoolUpdateEvent, id int) {
	sub, err := client.AccountSubscribe(
		publicKey,
		rpc.CommitmentConfirmed,
	)
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()
	for {
		got, err := sub.Recv()
		if err != nil {
			panic(err)
		}
		borshDec := bin.NewBorshDecoder(got.Value.Data.GetBinary())
		var meta token.Account
		err = borshDec.Decode(&meta)
		if err != nil {
			panic(err)
		}
		p.mu.Lock()
		if meta.Amount != p.amounts[id] {
			p.amounts[id] = meta.Amount
			p.slots[id] = got.Context.Slot
			fmt.Println(fmt.Sprintf("%s   %d  SLOT:%d", meta.Mint, p.amounts[id], p.slots[id]))
			if p.slots[0] == p.slots[1] {
				fmt.Println(fmt.Sprintf("Price %f", (float64(p.amounts[1])/USDC)/(float64(p.amounts[0])/SOL)))
			}
		}
		p.mu.Unlock()
	}
}
