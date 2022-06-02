# solana-pool-monitor

This repo is POC to test account monitoring of Solana blockchain. I tried to listen to pair of addresses of Orca SOL-USDC AMM pool and find updated prices after each trade is done against the pool.

For simple AMM, there is subscription based monitoring.

For AMMs whose liquidity is distributed between pool and orderbook like Raydium and Serum, price calculation is a bit tricky. Code in this repo finds price from Raydium SOL-USDC pool by requesting a few things (It's not subscription based)

## Try locally

Build

`go build`

Run

`go run solana-pool-monitor`

Price and other details should be logged in terminal

![Screenshot 2022-06-02 at 11 29 02 PM](https://user-images.githubusercontent.com/12713808/171701744-9134d310-4ddf-47b5-9efc-6b36cabb4940.png)
