<img src="https://r2cdn.perplexity.ai/pplx-full-logo-primary-dark%402x.png" style="height:64px;margin-right:32px"/>

# Perform a detailed gap analysis between standard Cosmos EVM starter architecture and Tabbre Core Chain MVP requirements.

Test these requirements against standard Cosmos EVM capabilities:

1. TABB: fixed 1B supply minted once at genesis (no ongoing inflation)
2. BRE as the gas/fee token for all transactions
3. TABB as staking token (not fee token)
4. Validator rewards paid in BRE from fee income (not new token emissions)
5. Permissioned consortium validator set at launch
6. EVM compatibility and JSON-RPC support
7. On-chain governance for parameters and upgrades
8. Slashed TABB burned or sent to burn mechanism
9. Public RPC endpoints for developers

Create these deliverables:

1. A detailed table with columns:
    - Requirement
    - Supported out-of-box by Cosmos EVM?
    - Partial support / needs config?
    - Custom work required?
    - Risk level (Low/Medium/High)
    - Recommended implementation approach
2. Ranked list of the 5 hardest engineering challenges
3. Recommendations: what to implement in MVP vs defer to later phases
4. Specific integration points to address (ante handlers, fee collection, staking params, module accounts)

Create gap-analysis.md with this information.

End with:

- Top 5 technical risks
- Top 5 possible simplifications for MVP
- Exact next build task recommendation

Below is a practical gap analysis for Tabbre Core Chain MVP against a standard Cosmos EVM starter. The short version is: Cosmos EVM already gives you the EVM, JSON-RPC, governance, upgrades, and standard PoS building blocks, but Tabbre’s split-token model—TABB for staking/governance and BRE for fees/rewards—is the main area where standard behavior does not cleanly match requirements and will likely need custom fee-path work.[^1][^2][^3]

## gap-analysis.md

### Detailed table

| Requirement | Supported out-of-box by Cosmos EVM? | Partial support / needs config? | Custom work required? | Risk level | Recommended implementation approach |
| :-- | :-- | :-- | :-- | :-- | :-- |
| 1. TABB fixed 1B supply minted once at genesis, no ongoing inflation | Partially. Cosmos SDK chains can set arbitrary genesis balances and staking token denomination, but standard staking/distribution designs often assume a reward flow that is frequently paired with inflationary issuance. [^2][^4] | Yes. Genesis can mint and allocate a fixed supply, and inflation can be disabled or omitted depending on module choices and app wiring. [^4][^2] | Low custom work. Mostly genesis config, token metadata, and making sure no mint/inflation path remains enabled for TABB. [^4][^2] | Low | Mint all 1,000,000,000 TABB at genesis; set TABB as bond denom in staking params; exclude or disable any inflationary mint logic for TABB; add invariant tests asserting total TABB supply never exceeds genesis supply except downward via burns. [^2][^4] |
| 2. BRE as gas/fee token for all transactions | Not cleanly out-of-box. Cosmos EVM provides fee market and EIP-1559 style logic, but standard setups generally assume a primary fee token aligned with the chain’s core denomination. [^3][^5][^6] | Partial. JSON-RPC, EVM fee market, and fee collection exist, but using a fee token different from the staking token requires explicit handling across Cosmos and EVM transaction paths. [^1][^5] | Yes. Likely ante-handler and fee validation customization, plus fee collector/accounting changes and careful denom handling for both EVM and Cosmos txs. [^5][^3] | High | For MVP, implement a single canonical BRE base denom for all fees, customize fee admission and fee deduction so both native SDK txs and EVM txs pay in BRE, and keep only one fee denom active to reduce complexity. [^5][^1] |
| 3. TABB as staking token, not fee token | Yes for staking token selection. Cosmos staking supports a configured bond denom. [^2] | Yes. Straightforward if TABB is set as bond denom, but operational interactions with a separate fee token still need design. [^2][^5] | Minimal direct custom work for staking denom itself; indirect custom work exists because fee and staking tokens differ. [^2][^5] | Medium | Set TABB as staking bond denom in x/staking params and keep staking, delegation, redelegation, and slashing entirely in TABB. Avoid hybrid staking logic in MVP. [^2] |
| 4. Validator rewards paid in BRE from fee income, not new token emissions | Partially. Cosmos distribution and fee collection already support distributing collected fees to validators/delegators, and Cosmos EVM has fee distribution mechanics. [^3][^5] | Yes. Works if fees are collected in BRE, but reward routing, commissions, and delegator accounting must all remain BRE-denominated without relying on inflationary minting. [^3][^5] | Moderate. Need to verify x/distribution and fee collector behavior with BRE-only fee flow and ensure no assumptions about reward denom conflict with staking denom. [^2][^5] | Medium-High | Use fee_collector + distribution pipeline with BRE as fee denom; disable TABB inflation rewards entirely; test validator commission, community pool behavior, and delegator withdrawals in BRE. [^5][^2] |
| 5. Permissioned consortium validator set at launch | Not directly as a first-class “consortium mode” feature in standard Cosmos EVM starter templates. [^3][^2] | Partial. Genesis validators can be pre-defined, and operationally you can launch with a fixed set. [^2] | Yes, if you need ongoing permissioned admission enforcement after genesis rather than simply starting with a closed initial set. [^2] | Medium | MVP approach: start with genesis-defined validators and no open validator onboarding path; if new validator admission is needed, gate it through governance or an allowlist check later rather than building a complex custom permissioning module now. [^2] |
| 6. EVM compatibility and JSON-RPC support | Yes. Cosmos EVM explicitly provides EVM execution and Ethereum JSON-RPC over HTTP and WebSocket. [^3][^1][^7] | Mostly config. Need to enable namespaces, endpoints, gas caps, and timeouts in app.toml or flags. [^1] | Little to none unless custom RPC behavior is desired. [^1] | Low | Use standard Cosmos EVM JSON-RPC server, enable only required namespaces (`eth,web3,net,txpool` initially), and test with MetaMask and Hardhat/Foundry early. [^1][^7] |
| 7. On-chain governance for parameters and upgrades | Yes. Cosmos SDK provides governance and upgrade modules suitable for MVP chain parameter changes and software upgrades. [^4][^8] | Mostly configuration and scope control. Proposal types and thresholds need to be chosen. [^4] | Low | Use standard x/gov and x/upgrade for MVP; keep proposal types narrow at launch (params, software upgrade, validator/gov operations only if needed). [^4][^8] |  |
| 8. Slashed TABB burned or sent to burn mechanism | Not guaranteed by default. Cosmos staking/slashing reduce validator/delegator stake, but exact sink behavior must be verified in the app’s module-account flow. [^2] | Partial. Slashing exists out-of-box, but burn semantics for slashed TABB should be made explicit in implementation. [^2] | Yes, likely moderate customization or explicit module-account routing to a burn-compatible sink with invariants. [^2] | Medium | Keep slashing in x/slashing/x/staking, then explicitly route slashed TABB to a burn module account or burn address path, and add tests proving total TABB supply decreases after slash events. [^2] |
| 9. Public RPC endpoints for developers | Yes. Cosmos EVM supports Ethereum JSON-RPC, and standard Cosmos nodes expose RPC services configurable in app settings. [^1][^7] | Mostly infra and config. Public exposure requires rate limits, namespace choices, capacity planning, and monitoring. [^1] | Low protocol custom work; medium operational work. [^1] | Low-Medium | Expose public JSON-RPC on dedicated RPC nodes, not validator nodes; enable standard namespaces, rate-limit aggressively, and separate public RPC from consensus infrastructure. [^1] |

## Hardest challenges

The five hardest engineering issues are mostly concentrated around the split-token design and launch controls rather than the EVM layer itself.[^5][^3]

1. **BRE as universal fee token across both Cosmos and EVM transaction paths** because standard Cosmos EVM chains already have fee logic, but aligning both SDK txs and EVM txs on a non-staking fee token is the main customization surface.[^5][^1]
2. **BRE-denominated validator and delegator rewards with TABB-denominated staking** because you need distribution, commissions, and withdrawals to remain coherent when the reward asset differs from the bonded asset.[^2][^5]
3. **Permissioned validator admission after genesis** because starting with a fixed validator set is easy, but enforcing a consortium-only policy over time without overbuilding custom governance/permissioning adds design complexity.[^2]
4. **Burning slashed TABB cleanly and audibly** because default slashing semantics may not automatically guarantee the exact deflationary burn behavior Tabbre wants.[^2]
5. **Wallet/tooling compatibility under a nonstandard fee model** because EVM users and tools often assume the familiar network gas asset and standard fee semantics, so BRE handling must be tested carefully through JSON-RPC and UI flows.[^1][^5]

## MVP vs later

For MVP, implement only the minimum required for a functional chain with the Tabbre economic model. Cosmos EVM already covers EVM execution, JSON-RPC, governance, upgrades, and core staking, so the MVP should focus engineering effort on token denomination correctness, fee flow, validator launch policy, and invariant testing.[^3][^1][^2]

**Implement now**

- Fixed 1B TABB genesis mint and allocation.[^4][^2]
- TABB as staking bond denom.[^2]
- BRE as single accepted fee denom for SDK and EVM txs.[^5][^1]
- BRE fee collection and distribution to validators/delegators.[^5]
- Genesis-defined consortium validator set.[^2]
- Standard x/gov and x/upgrade.[^8][^4]
- Public JSON-RPC on dedicated RPC nodes.[^1]

**Defer**

- Sophisticated validator onboarding beyond genesis or simple governance gating.[^2]
- Advanced tokenocracy beyond standard governance weighting if it requires heavy customization.[^4]
- Multi-denom fees, fee abstraction, or complex gas sponsorship.[^5]
- EVM precompiles for staking/governance UX unless clearly necessary.[^9][^10]
- Cross-chain IBC and BRE monetary authority mechanics beyond seeded MVP behavior.[^4]


## Specific integration points

These are the concrete places the engineering team should inspect first in a Cosmos EVM starter.

- **Ante handlers / fee admission**: enforce BRE as accepted fee denom for native Cosmos txs and ensure EVM txs satisfy base fee / min gas logic in BRE rather than an assumed default denom.[^1][^5]
- **Fee market / EIP-1559 path**: confirm base fee, priority tip, and mempool admission all operate in BRE and are reflected consistently in JSON-RPC responses.[^1][^5]
- **x/staking params**: set bond denom to TABB and keep all bonded/unbonding/slashing accounting in TABB.[^2]
- **x/distribution + fee_collector**: ensure collected BRE flows correctly to validators, commissions, delegators, and any community pool destination without TABB emissions.[^5][^2]
- **Module accounts**: define explicit module-account permissions for mint, burn, fee collection, staking pools, and slash handling, especially to prevent any post-genesis TABB mint path.[^4][^2]
- **Genesis builder**: create deterministic genesis setup for 1B TABB supply, correct allocations, seeded BRE balances, and consortium validator accounts.[^4]
- **JSON-RPC config**: enable namespaces and expose endpoints with conservative defaults for public use.[^7][^1]
- **Governance / upgrade app wiring**: include x/gov and x/upgrade in the standard app stack with narrow proposal surface at launch.[^8][^4]


## Top 5 technical risks

- BRE fee-token implementation may require more invasive Cosmos + EVM fee-path changes than expected.[^5]
- Reward distribution in BRE may expose edge cases where tooling or accounting implicitly assumes reward and staking denoms are aligned.[^5][^2]
- Wallet and RPC tooling may behave unexpectedly if gas estimation and display logic assume a different canonical asset.[^1]
- Slashed TABB may not burn exactly as intended unless module-account routing is explicitly customized and tested.[^2]
- Permissioned validator policy can become messy if you try to support post-genesis consortium changes without a simple governance-based rule.[^2]


## Top 5 simplifications for MVP

- Use exactly one BRE base denom for all fees, with no multi-denom support.[^5]
- Start with a fixed genesis validator set and no open onboarding path.[^2]
- Use standard x/gov and x/upgrade instead of custom tokenocracy logic in v1.[^4]
- Avoid custom precompiles unless required for contract-based staking/governance UX.[^10][^9]
- Keep BRE issuance for MVP simple, seeded via genesis or controlled setup, rather than implementing the full future BRE monetary authority immediately.[^4]


## Exact next build task recommendation

The next build task should be to produce a **design-and-code spike for fee flow**, specifically a small prototype branch that proves four things: TABB is the staking bond denom, BRE is the only accepted fee denom, EVM txs can pay fees in BRE through JSON-RPC, and collected BRE reaches the fee collector/distribution pipeline correctly.  If that spike works cleanly, the rest of the MVP architecture is relatively straightforward; if it does not, you will need to choose a fallback before deeper implementation begins.[^3][^1][^5][^2]
<span style="display:none">[^11][^12][^13][^14][^15][^16][^17]</span>

<div align="center">⁂</div>

[^1]: https://docs.cosmos.network/evm/next/api-reference/ethereum-json-rpc

[^2]: https://evm.cosmos.network/docs/sdk/v0.53/build/modules/staking/README

[^3]: https://cosmos-docs.mintlify.app/evm/v0.5.0/documentation/concepts/overview

[^4]: https://docs.cosmos.network/sdk/v0.53/learn/intro/overview

[^5]: https://docs.emoney.network/tendermint-and-evmos/gas-and-fees/how-are-gas-and-fees-handled-on-e-money-network

[^6]: https://cosmos-docs.mintlify.app/evm/next/documentation/concepts/tokens

[^7]: https://docs.cosmos.network/evm/v0.4.x/api-reference/ethereum-json-rpc

[^8]: https://evm.cosmos.network/protocol/modules

[^9]: https://forum.cosmos.network/t/evm-compatibility-and-the-price-of-atom/15721

[^10]: https://pkg.go.dev/github.com/zenanetwork/zena

[^11]: https://trustwallet.com/blog/staking/how-to-stake-cosmos-atom-and-earn-rewards-using-trust-wallet

[^12]: https://docs.cronos.org/cronos-chain-protocol/module_overview

[^13]: https://help.crypto.com/en/articles/4786335-staking-cosmos-atom-on-crypto-com-onchain-earn

[^14]: https://support.poloniex.com/hc/en-us/articles/360040016033-Cosmos-staking-What-is-it-And-other-Frequently-Asked-Questions

[^15]: https://docs.story.foundation/network/learn/node-software/cosmos-modules/evmstaking-module

[^16]: https://everstake.one/crypto-reports/cosmos-staking-insights-and-analysis-h1-2025

[^17]: https://castropatton.com/why-transaction-fees-and-staking-strategy-matter-more-than-you-think-in-cosmos/

