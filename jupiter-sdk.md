Exploring Jupiter SDK APIs

API for On-Chain jupiter program. - https://github.com/jup-ag/instruction-parser/blob/main/src/idl/jupiter.ts

Pattern for on-chain jupiter swap.

There is some instruction called `SetTokenLedger` which seems to have same "instruction data" and "token ledger address" but different "token account address". 

API in TS is `export declare function createSetTokenLedgerInstruction(tokenAccountAddress: PublicKey): TransactionInstruction;` which means instruction only has tokenAccountAddress as var ans rest is constant for all swaps through Jupiter.

TokenAccount address given is middle currency of the arb trade. For example - if trade is USDC -> SOL -> USDC, then tokenAccountAddress will be SOL address.

Example 
- https://solscan.io/tx/3t3x5q74j4H1cTNWh6eb7pySeoTtGZjxYzMkCJK8mjA84cxggtXRUGDrdJfpMwfNoG16orAQPXtL5Pj5wizzkiV
In this stSOL -> LDO -> stSOL, then tokenAccountAddress is LDO address (https://solscan.io/account/Cbumh9Pve9EDsoifnv9j2mKcC8qBw9azT3LdWhcKJCzC)
- https://solscan.io/tx/3fg8LnUsRD8KK6wBVsms9ZapzjBG9hhvP1bveCxtzs8DfH1CMiqeGgrz9G5gX6dtQTcD2DsZrdXvPmy9rozAgVHh
In this RAY -> ATLAS -> RAY, then tokenAccountAddress is ATLAS address (https://solscan.io/account/B3hMKxtAXzdUR7NvRCMpHKWXJbAi1kJXP9yZtzH8coBz)


Instruction order for ARB using Jupiter

- SetTokenLedger instruction
- Swap1 (A->B)
- Swap2 (B->A)

Find out how to cancel this arb transaction if not profitable? Does adding minOutAmount in Swap2 works? 


Exported functions for instruction creation
```
export declare const JUPITER_PROGRAM_ID_STAGING: PublicKey;
export declare const JUPITER_PROGRAM_ID_PRODUCTION: PublicKey;
export declare function createMercurialExchangeInstruction(swapLayout: MercurialSwapLayoutState, userSourceTokenAccountAddress: PublicKey, userDestinationTokenAccountAddress: PublicKey, user: PublicKey, amount: number | null, minimumOutAmount: number, platformFee: PlatformFee | undefined): TransactionInstruction;
export declare function createSerumSwapInstruction(market: Market, inputMint: PublicKey, openOrdersAddress: PublicKey, userSourceTokenAccountAddress: PublicKey, userDestinationTokenAccountAddress: PublicKey, user: PublicKey, amount: number | null, minimumOutAmount: number, platformFee: PlatformFee | undefined, referrer: PublicKey | undefined): TransactionInstruction;
export declare function createTokenSwapInstruction(tokenSwapState: TokenSwapState, inputMint: PublicKey, userSourceTokenAccountAddress: PublicKey, userDestinationTokenAccountAddress: PublicKey, user: PublicKey, amount: number | null, minimumOutAmount: number, platformFee: PlatformFee | undefined, isStep: boolean): TransactionInstruction;
export declare function createSenchaSwapInstruction(poolState: SenchaPoolState, sourceMint: PublicKey, userSourceTokenAccountAddress: PublicKey, userDestinationTokenAccountAddress: PublicKey, user: PublicKey, amount: number | null, minimumOutAmount: number, platformFee: PlatformFee | undefined): TransactionInstruction;
export declare function createCropperSwapInstruction(poolState: CropperPoolState, sourceMint: PublicKey, userSourceTokenAccountAddress: PublicKey, userDestinationTokenAccountAddress: PublicKey, user: PublicKey, feeAccount: PublicKey, amount: number | null, minimumOutAmount: number, platformFee: PlatformFee | undefined): TransactionInstruction;
export declare function createRaydiumSwapInstruction(raydiumAmm: RaydiumAmm, userSourceTokenAccountAddress: PublicKey, userDestinationTokenAccountAddress: PublicKey, user: PublicKey, amount: number | null, minimumOutAmount: number, platformFee: PlatformFee | undefined): TransactionInstruction;
export declare function createAldrinSwapInstruction(poolState: AldrinPoolState, sourceMint: PublicKey, userSourceTokenAccountAddress: PublicKey, userDestinationTokenAccountAddress: PublicKey, user: PublicKey, amount: number | null, minimumOutAmount: number, platformFee: PlatformFee | undefined): TransactionInstruction;
export declare function createAldrinV2SwapInstruction(poolState: AldrinPoolState, sourceMint: PublicKey, userSourceTokenAccountAddress: PublicKey, userDestinationTokenAccountAddress: PublicKey, curve: PublicKey, user: PublicKey, amount: number | null, minimumOutAmount: number, platformFee: PlatformFee | undefined): TransactionInstruction;
export declare function createCremaSwapInstruction(poolState: CremaPoolState, sourceMint: PublicKey, sourceTokenAccount: PublicKey, destinationTokenAccount: PublicKey, userTransferAuthority: PublicKey, amount: number | null, minimumOutAmount: number, platformFee: PlatformFee | undefined): TransactionInstruction;
export declare function createRiskCheckAndFeeInstruction(userDestinationTokenAccount: PublicKey, userTransferAuthority: PublicKey, minimumOutAmount: number, platformFee: PlatformFee | undefined): TransactionInstruction;
export declare function createSetTokenLedgerInstruction(tokenAccountAddress: PublicKey): TransactionInstruction;
export declare function createCreateTokenLedgerInstruction(user: PublicKey): TransactionInstruction;
export declare function createOpenOrdersInstruction(market: Market, user: PublicKey): [PublicKey, TransactionInstruction];
export declare function createSaberExchangeInstruction(saberPool: StableSwap, inputMint: PublicKey, userSourceTokenAccountAddress: PublicKey, userDestinationTokenAccountAddress: PublicKey, user: PublicKey, amount: number | null, minimumOutAmount: number, platformFee: PlatformFee | undefined): TransactionInstruction;
export declare function createSaberAddDecimalsDepositInstruction(addDecimals: AddDecimals, sourceTokenAccountAddress: PublicKey, destinationTokenAccountAddress: PublicKey, userTransferAuthority: PublicKey, amount: number | null, minimumOutAmount: number, platformFee: PlatformFee | undefined): TransactionInstruction;
export declare function createSaberAddDecimalsWithdrawInstruction(addDecimals: AddDecimals, sourceTokenAccountAddress: PublicKey, destinationTokenAccountAddress: PublicKey, userTransferAuthority: PublicKey, amount: number | null, minimumOutAmount: number, platformFee: PlatformFee | undefined): TransactionInstruction;
export declare function createLifinitySwapInstruction(swapState: LifinitySwapLayoutState, sourceMint: PublicKey, userSourceTokenAccountAddress: PublicKey, userDestinationTokenAccountAddress: PublicKey, userTransferAuthority: PublicKey, amount: number | null, minimumOutAmount: number, platformFee: PlatformFee | undefined): TransactionInstruction;
```
