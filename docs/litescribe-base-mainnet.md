# Litescribe Base Mainnet Deployment

This repo is pinned to the EntryPoint v0.6 line for Coinbase Smart Wallet compatibility on Base.

## Network

- Chain: Base Mainnet
- Chain ID: `8453`
- EntryPoint: `0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789`

## Required Environment

Set the following values before deployment:

- `PRIVATE_KEY`
  The EOA that will broadcast the deployment transaction.
- `VERIFYING_SIGNER`
  The off-chain signer address that will sign paymaster approvals.
- `OWNER`
  The admin/owner address for the deployed paymaster.
- `RPC_URL`
  Base Mainnet RPC URL for `forge script`.

Copy the root template first:

```powershell
Copy-Item .env.example .env
```

## Deploy

With Foundry installed:

```powershell
$env:PRIVATE_KEY='<deployer-private-key>'
$env:ENTRYPOINT='0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789'
$env:VERIFYING_SIGNER='<verifying-signer-address>'
$env:OWNER='<owner-address>'
$env:RPC_URL='<base-rpc-url>'
forge script script/Deploy.s.sol:DeployScript --rpc-url $env:RPC_URL --broadcast
```

## After Deployment

1. Record the deployed paymaster address.
2. Fund the paymaster's EntryPoint deposit with a small amount of Base ETH.
3. Configure the Litescribe backend with:
   - `BASE_ERC20_PAYMASTER_ADDRESS`
   - `BASE_PAYMASTER_SIGNER_ADDRESS`
   - `BASE_PAYMASTER_SIGNER_PRIVATE_KEY`
   - `BASE_PAYMASTER_TOKEN_RECEIVER_ADDRESS`
4. Point the extension paymaster URL at the self-hosted paymaster RPC service.

## Current Litescribe Deployment

- Paymaster contract: `0xdcAce6220e4Fb37CF91684c87f5164CDA675C366`
- Deployment transaction: `0x5ad99e880be3c27e4b971f8ef9ce7768507764c47359dc96bb0954b2e4dde050`

Deployment broadcast artifacts remain local-only and are not intended to be pushed upstream.

## Litescribe Policy Expectations

The off-chain paymaster service should stay narrow:

- Base Mainnet only
- Base USDC only:
  - `0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913`
- 0x AllowanceHolder only:
  - `0x0000000000001fF3684f28c67538d4D072C22734`
- `USDC -> cbLTC` swap path only

The extension already injects the approval order expected by the contract:

1. approve USDC to the paymaster if needed
2. approve USDC to the swap spender if needed
3. execute the swap
