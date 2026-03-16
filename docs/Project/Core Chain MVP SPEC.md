<img src="https://r2cdn.perplexity.ai/pplx-full-logo-primary-dark%402x.png" style="height:64px;margin-right:32px"/>

# Core Chain MVP SPEC.md

## 1. Status

This document defines the normative engineering specification for the Tabbre Core Chain MVP. The MVP is the initial single-chain implementation of TabbreChain and MUST provide the foundation for later expansion into the Tabbre multichain ledger, validator-governed chain creation, BRE monetary operations, and DAO-controlled treasury flows.[^1][^2][^3]

This specification uses the key words MUST, MUST NOT, REQUIRED, SHOULD, SHOULD NOT, and MAY as described in RFC 2119.[^1]

## 2. Purpose

The Core Chain MVP MUST implement the single-chain proof-of-stake phase of the Tabbre protocol before any multichain rollout occurs.  The chain MUST be based on the Cosmos EVM SDK, MUST use TABB as the staking asset, MUST use BRE as the fee token, and MUST support smart contracts on the core chain for governance and protocol evolution.[^3][^1]

The MVP SHOULD optimize for correctness, deterministic finality, operational simplicity, and a clean migration path to the future Tabbre Global Ledger.  The MVP MUST NOT attempt to implement full multichain balance aggregation or live chain spawning/splitting in its first production release.[^2][^1]

## 3. Scope

### 3.1 In Scope

The MVP MUST include the following:

- A production-capable Core Chain binary.[^1]
- Proof-of-stake consensus using TABB.[^3][^1]
- BRE-denominated gas and transaction fees.[^3][^1]
- EVM-compatible smart-contract execution.[^1][^3]
- Validator-restricted contract deployment policy.[^1]
- Governance and upgrade support.[^3][^1]
- Public access endpoints suitable for wallet and explorer integration.[^2]
- Wallet-facing balance and metadata query support.[^2][^1]
- Reserved protocol scaffolding for future BRE Monetary Authority, treasury, and multichain registry features.[^2][^3][^1]


### 3.2 Out of Scope

The MVP MUST NOT require the following:

- Live sub-chains.[^1]
- Aggregate balances across multiple chains.[^1]
- Production chain spawning execution.[^1]
- Production chain splitting execution.[^1]
- Full BRE credit creation and CCA lending flows.[^3]
- Global UBI distribution.[^2][^3]


## 4. Terminology

For the purposes of this document:

- **Core Chain** means the initial single-chain Tabbre ledger.[^1]
- **TABB** means the staking, governance, and collateral token of the Tabbre system.[^2][^3]
- **BRE** means the native transaction and utility currency of the Tabbre system.[^3][^2][^1]
- **Validator Account** means the Tabbre account associated one-to-one with a consensus node or validator operator.[^1]
- **Authorized Deployer** means a validator-linked account permitted to deploy smart contracts.[^1]
- **Monetary Authority** means the future BRE authority responsible for issuance and supply policy.[^3]
- **Treasury** means the future DAO-directed allocation and buy-burn control domain.[^2][^3]


## 5. Goals

The Core Chain MVP MUST satisfy the following goals:

- Provide an initial live TabbreChain.[^2]
- Enable a MetaMask-like wallet experience through EVM-compatible APIs.[^2]
- Support an explorer and fully hosted public access node.[^2]
- Establish the staking, governance, and smart-contract foundations for later BRE operational framework work.[^3][^2]
- Preserve compatibility with the later multichain architecture in which the Core Chain governs future sub-chains.[^1]

The implementation SHOULD minimize protocol debt that would complicate migration to the later multichain design.[^1]

## 6. Protocol Requirements

### 6.1 Base Platform

The Core Chain MUST be implemented on a Cosmos SDK-based application stack with Cosmos EVM execution support.  The chain MUST use a Tendermint/CometBFT-style Byzantine fault tolerant proof-of-stake consensus engine with deterministic finality.[^1]

### 6.2 Consensus

The chain MUST implement proof of stake using TABB as the staking asset.  The network MUST achieve finality through a validator set using BFT consensus with block commitment after supermajority validator approval.[^3][^1]

The initial validator set SHOULD be permissioned, since the early-stage design assumes validator operation by the Tabbre Project or consortium participants before broader decentralization.  The protocol MUST support later expansion toward broader validator participation without changing the core staking asset or fee asset.[^3][^1]

### 6.3 Currency Roles

The protocol MUST define TABB and BRE as distinct native assets with distinct roles:

- TABB MUST be used for staking and validator collateral.[^3][^1]
- BRE MUST be used for transaction gas and fees.[^3][^1]

The protocol MUST NOT accept TABB as the default fee token for standard transaction execution.[^3][^1]

## 7. Token Specification

### 7.1 TABB

TABB MUST be represented on the Core Chain as a native asset with fixed total issuance defined at genesis.  The MVP MUST preserve the fixed-supply model described in the tokenomics paper.[^3]

The canonical token allocation model SHOULD match the following distribution:

- 20% private sale.[^3]
- 10% public sale.[^3]
- 10% founders.[^3]
- 10% community development.[^3]
- 50% reserve.[^3]

If vesting contracts are not included in the MVP, the genesis state MUST still preserve account separation for these categories.[^3]

### 7.2 BRE

BRE MUST be represented on the Core Chain as the native gas, fee, and transfer currency.  The Core Chain MUST begin with an initial BRE supply allocated to accounts under Tabbre Project control.[^1][^3]

The MVP MAY implement BRE minting and burning under a restricted authority account.  The MVP MUST NOT expose unrestricted public BRE minting.[^3]

### 7.3 Base Units

The implementation MUST define base units:

- `utabb` for TABB.
- `ubre` for BRE.

All state transitions MUST use integer base units internally.

## 8. Account Model

The Core Chain MUST support an account model compatible with both Cosmos-native operations and Ethereum-style wallet interaction.  The chain SHOULD use secp256k1-compatible keys for EVM wallet interoperability.[^2][^1]

A single canonical address derivation scheme MUST be chosen and MUST remain stable across protocol upgrades, because future multichain operation depends on the same account address existing across chains.  Validator operator identity MUST be mapped one-to-one to a Tabbre validator account.[^1]

## 9. Validator Model

### 9.1 Validator Eligibility

In the MVP, validator participation SHOULD be permissioned.  Only approved validator operators MUST be permitted to join the active set.[^1][^3]

### 9.2 Bonding Asset

All validator bonding, self-bonding, and delegated stake MUST be denominated in TABB.[^1][^3]

### 9.3 Rewards

Transaction fees collected in BRE MUST be distributable to validators and delegators according to protocol parameters.  The implementation SHOULD support configurable distribution percentages to allow later governance tuning.[^3]

### 9.4 Slashing

The protocol MUST support slashing for double-signing and validator liveness failures.  Slashing penalties MUST be denominated in TABB.  The final destination of slashed TABB MUST be explicitly defined by parameters or governance; it MAY be burned in accordance with the tokenomics model.[^3]

## 10. Smart Contracts

### 10.1 EVM Support

The Core Chain MUST support smart-contract execution on the core chain.  The EVM interface MUST be sufficient to support governance-related and protocol-related contracts.[^2][^1][^3]

### 10.2 Deployment Policy

Only validator accounts or accounts explicitly authorized under validator policy MUST be permitted to deploy smart contracts.  Non-authorized accounts MUST be rejected when attempting EVM contract creation.[^1]

### 10.3 Contract Calls

The protocol MAY allow non-validator accounts to call deployed contracts unless a contract’s own logic restricts access.[^1]

### 10.4 System Contracts

The chain SHOULD reserve addresses or predeploy contracts for future governance, treasury, and BRE monetary authority integration.[^2][^3]

## 11. Governance

The Core Chain MUST support on-chain governance for protocol evolution, treasury control, policy changes, and upgrades.  Governance voting power SHOULD be aligned with TABB ownership or delegated TABB stake.[^2][^3]

The governance system MUST support at minimum:

- Parameter change proposals.
- Software upgrade proposals.
- Validator policy proposals.
- Treasury authority proposals.
- Monetary authority proposals.[^2][^1][^3]

The protocol MUST support hard-fork style upgrades because the Core Chain paper anticipates multiple hard forks during development.[^1]

## 12. Fees and Gas

All transaction charges MUST be payable in BRE.  The protocol SHOULD support congestion-sensitive pricing or an equivalent fee market mechanism because the multichain design anticipates variable gas pricing under congestion.[^1][^3]

The implementation MUST expose fee estimation suitable for wallet UX.  The fee mechanism SHOULD be parameterizable so governance can later tune base fee, priority fee, burn fraction, and validator distribution share.[^2]

## 13. Wallet and Query Support

The MVP MUST expose wallet-friendly query endpoints because the roadmap explicitly calls for a MetaMask-like wallet and public access node.  The read interface MUST support:[^2]

- Account overview.
- TABB balance.
- BRE balance.
- Validator set query.
- Fee quote query.
- Chain metadata query.[^2][^1]

Even in single-chain mode, the query surface SHOULD use future-compatible naming so it can evolve into multichain chain discovery and aggregate-balance queries later.[^1]

## 14. Public API Requirements

The Core Chain MUST expose:

- Cosmos-native gRPC and/or REST query interfaces.
- Ethereum-compatible JSON-RPC methods needed for wallet integration.[^2]

At minimum, Ethereum-style JSON-RPC SHOULD include:

- `eth_chainId`
- `eth_blockNumber`
- `eth_getBalance`
- `eth_estimateGas`
- `eth_sendRawTransaction`
- `eth_call`
- `eth_getTransactionReceipt`[^2]

Public read APIs MUST be available through non-validator public nodes.[^2]

## 15. Module Requirements

The MVP MUST include standard runtime modules required for accounts, balances, staking, slashing, governance, upgrades, fee handling, and EVM execution.[^3][^1]

The MVP MUST include the following Tabbre-specific modules or equivalent logic:

### 15.1 Validator Policy Module

This module MUST:

- Record validator-account mappings.[^1]
- Record authorized deployers.[^1]
- Enforce contract deployment restrictions.[^1]


### 15.2 Wallet Query Module

This module MUST:

- Expose wallet-friendly account and chain metadata queries.[^2][^1]
- Return balances and fee quotes in a stable schema.[^2][^1]


### 15.3 Core Registry Module

This module MUST reserve chain-level metadata and future multichain feature flags because the Core Chain later governs chain creation and validation.  The MVP MAY keep all multichain-related entries disabled.[^1]

### 15.4 BRE Monetary Module

The MVP MUST reserve a controlled authority model for BRE issuance and policy state.  The MVP SHOULD implement only minimal authority-controlled mint/burn support and MUST NOT implement unsecured public issuance.[^3]

### 15.5 Treasury Module

The MVP SHOULD reserve state and authority hooks for future treasury-controlled allocation, investor returns, and buy-burn actions.[^2][^3]

## 16. Genesis Requirements

The genesis state MUST define:

- A unique Core Chain ID.[^1]
- A fixed EVM chain ID.[^2]
- TABB genesis supply and category allocations.[^3]
- Initial BRE supply held by designated Tabbre-controlled accounts.[^1]
- Initial validator accounts and bonds in TABB.[^3][^1]
- Governance-capable authority accounts.[^2][^3]
- Validator-only contract deployment policy enabled by default.[^1]

If treasury and monetary authority modules are enabled, their authority accounts MUST be initialized at genesis.[^2][^3]

## 17. Message Types

The MVP MUST support transactions for:

- BRE transfers.[^3][^1]
- TABB transfers.[^3]
- Staking and unstaking of TABB.[^3][^1]
- Governance proposals and votes.[^2][^3]
- EVM contract calls and validator-authorized deployments.[^1]
- BRE mint/burn under authority control, if monetary module is enabled.[^3]
- Treasury-controlled transfers, if treasury module is enabled.[^2][^3]

The MVP SHOULD define protobuf services for its custom modules from the start, even if some methods are initially authority-only or operationally disabled.

## 18. State Transition Rules

The implementation MUST reject any transaction that violates one or more of the following:

- Invalid signature.
- Invalid account sequence.
- Fee not payable in BRE.[^3][^1]
- Contract deployment attempt by a non-authorized account.[^1]
- Unauthorized mint or treasury operation.[^2][^3]
- Invalid governance authority or malformed proposal.

The implementation SHOULD enforce these checks in the ante pipeline wherever possible.

## 19. Security Requirements

The MVP MUST be secure against unauthorized contract deployment, unauthorized token issuance, validator double-signing, and basic governance abuse.[^3][^1]

Operationally:

- Validator signing keys SHOULD be isolated from application nodes.
- Treasury and monetary authority accounts SHOULD be controlled by governance or robust multi-party authorization.[^2][^3]
- Public nodes MUST NOT hold validator signing keys.
- The chain MUST support jailing and slashing for validator faults.[^3]

The MVP SHOULD prefer conservative permissions over convenience in the initial release.[^1][^3]

## 20. Upgradeability

The chain MUST support coordinated software upgrades because protocol evolution is expected during development.  Store migrations for custom modules MUST preserve address stability, token balances, validator state, and governance records.[^1]

Any future activation of multichain functionality SHOULD occur through explicit upgrade and governance processes rather than ad hoc runtime toggles.[^1]

## 21. Observability

The Core Chain implementation SHOULD expose:

- Block production metrics.
- Validator liveness metrics.
- Fee market metrics.
- EVM execution metrics.
- Governance metrics.
- Token supply metrics for BRE and TABB.

The chain MUST emit stable events for:

- Asset transfers.
- Validator lifecycle actions.
- Contract deployment.
- Governance proposals and votes.
- BRE mint/burn.
- Treasury operations.

These events SHOULD support downstream explorer and wallet indexing.[^2]

## 22. Testing Requirements

The MVP MUST include:

- Unit tests for all custom modules.
- Integration tests for staking with TABB and fees in BRE.[^3][^1]
- Tests proving that unauthorized contract deployment fails.[^1]
- Tests proving that authorized deployment succeeds.[^1]
- Genesis validation tests for TABB and BRE allocations.[^3][^1]
- Governance and upgrade tests.[^3][^1]
- JSON-RPC compatibility tests for wallet interactions.[^2]

The MVP SHOULD include end-to-end localnet and testnet scenarios with multiple validators and public RPC nodes.

## 23. Interoperability and Future Compatibility

The Core Chain MVP MUST preserve forward compatibility with the later multichain design in the following ways:

- Address format MUST remain stable across future sub-chains.[^1]
- Core governance authority MUST remain the root of future chain creation and splitting.[^1]
- Wallet query surfaces SHOULD be extensible to chain discovery and aggregate-balance reporting.[^1]
- Registry and feature-flag state SHOULD be designed for later activation of sub-chain records.[^1]

The MVP MUST NOT hard-code assumptions that prevent future migration to a Core Chain plus many sub-chains model.[^1]

## 24. Recommended Initial Parameters

The following are recommended implementation defaults for the MVP and MAY be adjusted by governance:

- Block time target: approximately 2 seconds, consistent with the stated 1–3 second finality target.[^1]
- Validator count: small permissioned set for initial network stability.[^3][^1]
- Contract deployment policy: validator-only enabled.[^1]
- Fee denom: BRE only.[^3][^1]
- Staking denom: TABB only.[^3][^1]

These values are implementation defaults and are not protocol invariants unless codified by governance.

## 25. Acceptance Criteria

The Core Chain MVP SHALL be considered compliant with this specification when all of the following are true:

- A production-capable single Core Chain runs with BFT proof of stake.[^1]
- TABB is the staking asset.[^3][^1]
- BRE is the fee and gas asset.[^3][^1]
- Initial BRE supply is allocated to Tabbre-controlled accounts at genesis.[^1]
- Contract deployment is restricted to validator-authorized accounts.[^1]
- Public APIs support wallet and explorer integration.[^2]
- Governance and upgrades are functional.[^3][^1]
- Treasury, monetary authority, and registry extension points exist for later protocol phases.[^2][^3][^1]


## 26. Non-Normative Implementation Notes

A practical implementation MAY use a monorepo with a dedicated Core Chain binary, custom Cosmos SDK modules, protobuf services, EVM system contracts, a public RPC layer, and an indexer service. This structure is consistent with the roadmap requirement for an initial live chain, public access node, wallet, and explorer, while also preserving the Core Chain’s later role as the authority over multichain lifecycle operations.[^2][^1]

If you want, I can next produce:

- a matching `ARCHITECTURE.md`,
- a `PROTO.md` with concrete protobuf definitions,
- or a `GENESIS.md` covering accounts, allocations, and chain parameters.

<div align="center">⁂</div>

[^1]: TabbreMultichainLedger20260306.pdf

[^2]: TabbreWhitePaper-DARK-20251227.pdf

[^3]: TokenomicsOfTheTabbreSystem20250718.pdf

