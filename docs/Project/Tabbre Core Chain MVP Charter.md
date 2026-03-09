<img src="https://r2cdn.perplexity.ai/pplx-full-logo-primary-dark%402x.png" style="height:64px;margin-right:32px"/>

## Tabbre Core Chain MVP Charter

**Version 1.0**
**Date: March 7, 2026**
**Sponsor: Tabbre Foundation / DAO**
**Product Lead: [TBD]**
**Protocol Lead: [TBD]**

### 1. Vision and purpose

Tabbre Core Chain is the foundational proof‑of‑stake blockchain powering the Tabbre ecosystem, enabling decentralized governance via TABB, collateralization for energy project financing, and the operational framework for BRE issuance.  This MVP delivers Phase 2 of the Tabbre Roadmap: an initial live TabbreChain with hosted public access node, wallet support, and explorer, built on Cosmos SDK with Cosmos EVM for Ethereum compatibility.  Success means external users can stake TABB, govern via on‑chain proposals, and interact via familiar Web3 tools, providing the secure base for Tabbre’s sustainable energy mission.[^1][^2][^3]

### 2. Objectives

- Launch a production‑ready PoS chain secured by TABB staking within 16 weeks.[^2][^1]
- Enable basic DAO governance for protocol parameters and upgrades, aligning with TABB’s role as governance token.[^3][^1]
- Provide developer‑friendly access via JSON‑RPC, MetaMask, and explorer to accelerate ecosystem growth.[^2]
- Achieve 99% uptime on public testnet and mainnet, with 10+ active validators at launch.[^2]


### 3. MVP scope (in)

**Core chain features:**

- Cosmos SDK app‑chain with Cosmos EVM module for EVM execution and Ethereum JSON‑RPC.[^4][^5]
- TABB as native staking, governance, and gas token, with fixed supply mechanics and halving schedule prepared for activation.[^1][^3]
- Standard SDK modules: x/staking, x/bank, x/distribution, x/slashing, x/gov, x/upgrade.[^6][^7]
- Hosted RPC endpoints (HTTP/WebSocket), chain explorer, MetaMask wallet integration, testnet faucet.[^8][^2]

**Environments:** Local devnet, public testnet, mainnet.[^2]

### 4. MVP scope (out)

- Full BRE issuance, monetary policy, or stablecoin peg mechanisms (Phase 3).[^3][^2]
- IBC cross‑chain messaging or Tabbre Global Ledger (Phase 3+).[^9][^2]
- Advanced features: sharding, L2 scaling, buyback/burn automation, UBI distribution.[^1][^3]
- Mobile wallet, advanced explorer analytics, permissioned validators beyond genesis.[^2]


### 5. Key assumptions

- TABB exists on an exchange or bridgeable source for genesis allocations and testnet funding.[^1][^2]
- Initial validator set (8‑15 nodes) commits during testnet phase.[^2]
- Cosmos SDK v0.50+ and Cosmos EVM latest stable are production‑ready.[^5][^4]
- External audit completes within hardening phase (weeks 13‑14).[^10]


### 6. Success metrics

| Category | Metric | Target |
| :-- | :-- | :-- |
| **Technical** | Public testnet uptime | 95% over 14 days [^2] |
| **Technical** | Mainnet block time | <7 seconds, 99% uptime first week |
| **Adoption** | Active validators at mainnet | 10+ with >50% stake delegated [^1] |
| **Adoption** | Daily active addresses | 50+ in first week post‑launch |
| **Usability** | MetaMask integration | Users can stake TABB and propose via wallet |
| **Security** | Audit findings | All high/critical resolved pre‑mainnet [^10] |

### 7. Timeline and milestones

- **Weeks 1‑3:** Spec freeze, design (gate: signed charter).
- **Weeks 4‑10:** Build devnet/testnet (gate: public testnet live).
- **Weeks 11‑14:** Hardening, audit (gate: testnet stable).
- **Weeks 15‑16:** Mainnet launch (gate: genesis signed).[^11][^2]


### 8. Team and resources

| Role | Count | Key skills |
| :-- | :-- | :-- |
| Product Lead | 1 | Roadmap alignment, stakeholder mgmt |
| Protocol Engineer | 1‑2 | Cosmos SDK, EVM integration [^5] |
| DevOps/SRE | 1 | Kubernetes, monitoring, node ops |
| Security Engineer | 1 | Audits, threat modeling [^10] |
| EVM/Wallet Dev | 1 | JSON‑RPC, Solidity testing [^8] |

**Budget bands:** \$300K‑\$500K (team, infra, audit), excluding TABB allocations.[^12]

### 9. Risks and mitigations

| Risk | Likelihood/Impact | Mitigation |
| :-- | :-- | :-- |
| EVM module incompatibility | Medium/High | Use Cosmos EVM example_chain as base; integration tests early [^4][^13] |
| Validator recruitment slow | High/Medium | Recruit from Cosmos ecosystem; incentives via testnet rewards [^2] |
| Audit delays | Medium/High | Scope to core modules; parallel internal review [^10] |
| Token supply/bridge issues | Medium/Medium | Pre‑allocate testnet TABB; mainnet multisig [^1] |

**Approval:**
This charter defines the MVP boundary. Changes require Product Lead + Sponsor sign‑off.[^1][^2]
<span style="display:none">[^14][^15][^16][^17][^18][^19][^20][^21][^22][^23][^24]</span>

<div align="center">⁂</div>

[^1]: https://www.tabbre.co.uk/faq

[^2]: https://tabbre.com/roadmap

[^3]: https://tabbre.com/blog

[^4]: https://github.com/cosmos/evm

[^5]: https://cosmos-docs.mintlify.app/evm/v0.5.0/documentation/concepts/overview

[^6]: https://evm.cosmos.network/protocol/modules

[^7]: https://cosmos-docs.mintlify.app/evm/v0.5.0/documentation/cosmos-sdk/overview

[^8]: https://docs.cosmos.network/evm/next/api-reference/ethereum-json-rpc

[^9]: https://docs.cosmos.network/sdk/v0.53/learn/intro/overview

[^10]: https://www.certik.com/resources/blog/evm-cosmos-convergence-research-from-security-base-part-2

[^11]: https://www.cosmoslabs.io/blog/the-cosmos-stack-roadmap-2026

[^12]: https://www.youtube.com/watch?v=w_lwSeMcJt4

[^13]: https://libraries.io/go/github.com%2Fcosmos%2Fevm%2Fexample_chain

[^14]: https://www.tabbre.com/documents/html/TokenomicsOfTheTabbreSystem20250718

[^15]: https://tabbre.com/documents/000_TokenomicsOfTheTabbreSystem20250617.pdf

[^16]: https://tabbre.co.uk/documents/documents/manifesto/info.html

[^17]: https://github.com/tabbre-blockchain/Documents/

[^18]: https://tabbre.com/documents/ebook/ThepathToTheAgeOfSuperAbundance:HowFloatingSolarFarmsWillSaveTheWorld.md.pdf

[^19]: https://uk.linkedin.com/company/tabbre

[^20]: https://www.gate.com/learn/articles/tabi-chain-the-1st-modular-cosmos-gaming-chain-with-high-scalability-and-omni-execution-layer/3414

[^21]: https://www.linkedin.com/posts/cicunningham_tokenomics-tabbreproject-crypto-activity-7349140919285145602-qyho

[^22]: https://www.binance.com/en/square/post/17962887761306

[^23]: https://www.tabbre.co.uk/documents/investors/TabbreLightPaper20240825_1.pdf

[^24]: https://tabbre.io/documents/investors/TabbreFloatingSolar20251229.pdf

