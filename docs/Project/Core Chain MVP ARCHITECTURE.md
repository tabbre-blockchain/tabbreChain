<img src="https://r2cdn.perplexity.ai/pplx-full-logo-primary-dark%402x.png" style="height:64px;margin-right:32px"/>

# Core Chain MVP ARCHITECTURE.md

## 1. Status

This document describes the target architecture for the Tabbre Core Chain MVP. It is intended to accompany `SPEC.md` and explains how the MVP SHOULD be structured at the system, runtime, module, and service levels so that the initial single-chain deployment can evolve cleanly into the later Tabbre multichain design.

This document is descriptive and architectural. Where implementation behavior is normative, `SPEC.md` takes precedence.

## 2. Purpose

The Core Chain MVP architecture exists to deliver:

- An initial live Tabbre Core Chain.
- A permissioned proof-of-stake validator network.
- BRE-denominated fees and TABB-denominated staking.
- EVM-compatible smart-contract execution.
- Public node access for wallet and explorer use.
- Clean extension points for treasury, monetary authority, and future multichain control.

The architecture SHOULD prioritize:

- Simplicity in the first release.
- Strong operational safety.
- Explicit authority boundaries.
- Future compatibility with sub-chain governance and lifecycle management.


## 3. Architectural principles

The architecture SHOULD follow these principles:

### 3.1 Core first

The MVP SHOULD implement the single-chain system first and MUST NOT depend on live sub-chains. The Core Chain is the root authority of the future Tabbre ledger and SHOULD already be designed as that authority even before multichain activation.

### 3.2 Clear asset separation

TABB and BRE have different roles and the architecture SHOULD preserve that separation everywhere:

- TABB for staking, validator collateral, and governance weight.
- BRE for transaction execution, gas, and payment flows.


### 3.3 Permission before openness

The first production version SHOULD favor explicit allowlists, authority checks, and operational controls over open participation. This is especially important for validator onboarding, contract deployment, token issuance, and treasury movement.

### 3.4 Stable interfaces

Public APIs, internal module interfaces, and event schemas SHOULD be stable enough to survive the future transition from single-chain to multichain.

### 3.5 Replaceable services

Off-chain services such as indexers, wallet APIs, and public gateways SHOULD be replaceable without altering consensus-critical behavior.

## 4. System overview

At a high level, the Core Chain MVP consists of:

- The on-chain runtime.
- Validator node infrastructure.
- Public read and transaction gateways.
- Indexing and wallet support services.
- Governance and authority domains.

```text
                    +----------------------+
                    |   Wallet / Explorer  |
                    +----------+-----------+
                               |
                    +----------v-----------+
                    | Public API Gateway   |
                    | REST / gRPC / JSON   |
                    +----------+-----------+
                               |
              +----------------+----------------+
              |                                 |
   +----------v-----------+         +-----------v----------+
   | Public Full Node     |         | Indexer / Query DB   |
   | No validator key     |         | Explorer / Wallet    |
   +----------+-----------+         +-----------+----------+
              |                                 ^
              |                                 |
   +----------v---------------------------------+----------+
   |               Core Chain Network                      |
   |  CometBFT + Cosmos SDK App + EVM Runtime             |
   +----------+----------------------+---------------------+
              |                      |
   +----------v-----------+  +-------v------------+
   | Validator Node A     |  | Validator Node B   |
   | Signer + Full State  |  | Signer + Full State|
   +----------------------+  +--------------------+
```


## 5. Layers

The architecture SHOULD be understood in five layers.

### 5.1 Consensus layer

This layer is responsible for:

- Validator coordination.
- Proposal, prevote, precommit, and commit flow.
- Final block ordering.
- Liveness and safety under Byzantine assumptions.

This layer SHOULD be implemented with CometBFT or an equivalent Tendermint-family engine.

### 5.2 Application layer

This layer is the state machine. It is responsible for:

- Accounts and balances.
- Staking and slashing.
- Governance.
- Fees and distribution.
- EVM execution.
- Tabbre-specific policy modules.

This layer SHOULD be implemented as a Cosmos SDK application.

### 5.3 Execution layer

This layer is where transactions are interpreted and state transitions are applied. It includes:

- Native Cosmos messages.
- EVM message execution.
- Ante handling.
- Fee checks.
- Validator policy enforcement.


### 5.4 Data access layer

This layer provides:

- Query services.
- Event indexing.
- Explorer-friendly normalized data.
- Wallet-friendly metadata and fee quoting.

This layer SHOULD NOT modify consensus state directly.

### 5.5 Edge layer

This layer is the external interface and includes:

- JSON-RPC for EVM wallets.
- gRPC and REST for application queries.
- Load balancers.
- API auth and rate limiting where required.


## 6. On-chain application architecture

The Core Chain application SHOULD be built around a small set of standard runtime modules plus a thin set of Tabbre-specific modules.

### 6.1 Standard runtime modules

The application SHOULD include:

- Auth.
- Bank.
- Staking.
- Slashing.
- Distribution.
- Governance.
- Upgrade.
- Params or equivalent configuration support.
- Fee market support.
- EVM runtime.

These provide the core ledger, staking, governance, and execution behavior needed by the MVP.

### 6.2 Tabbre-specific modules

The MVP SHOULD add the following custom modules:

- `validator_policy`
- `wallet_query`
- `core_registry`
- `bre_monetary`
- `treasury`

These modules SHOULD remain small and narrowly scoped in the MVP.

## 7. Module responsibilities

### 7.1 `validator_policy`

Purpose:

- Enforce that only validator-authorized accounts can deploy contracts.
- Maintain validator-account and deployer authorization mappings.
- Provide a central policy boundary for privileged operations.

Responsibilities:

- Track validator-linked accounts.
- Track authorized deployers.
- Expose queryable authorization state.
- Integrate with the ante pipeline or EVM create path.

This module SHOULD be the only custom authority layer that directly gates contract deployment.

### 7.2 `wallet_query`

Purpose:

- Present a stable wallet-facing read interface.
- Normalize chain and account information for clients.
- Reduce client dependence on low-level protocol details.

Responsibilities:

- Account overview query.
- Balance query.
- Validator set query.
- Fee quote query.
- Chain metadata query.
- Feature-flag query.

Although the MVP is single-chain, the schema SHOULD be future-compatible with multichain discovery.

### 7.3 `core_registry`

Purpose:

- Hold metadata about the Core Chain and future multichain activation.
- Reserve namespaces and feature flags for later chain creation logic.

Responsibilities:

- Chain metadata.
- Reserved chain IDs.
- Future feature activation flags.
- Protocol identity metadata.

This module SHOULD NOT yet attempt to manage live sub-chain state in the MVP.

### 7.4 `bre_monetary`

Purpose:

- Reserve the architectural boundary for BRE issuance and policy control.
- Enable limited authority-controlled BRE mint/burn behavior if needed.

Responsibilities:

- Monetary authority account reference.
- BRE policy parameters.
- Optional restricted mint and burn messages.
- Future oracle and rate-policy storage.

This module SHOULD be minimal in MVP and SHOULD NOT implement the full mature lending architecture yet.

### 7.5 `treasury`

Purpose:

- Reserve the architectural boundary for future treasury-directed flows.
- Provide minimal authority-controlled treasury state if needed.

Responsibilities:

- Treasury authority.
- Allocation policy records.
- Optional controlled treasury transfer execution.
- Future buy-burn and allocation hooks.


## 8. Runtime flow

### 8.1 Transaction lifecycle

A transaction SHOULD flow through the system in this order:

```text
Client
  -> Public Gateway
  -> Full Node RPC
  -> Mempool Admission
  -> AnteHandler
  -> Consensus Proposal
  -> Block Execution
  -> Event Emission
  -> Commit
  -> Query / Index Propagation
```

Key enforcement points:

- Signature and sequence checks in ante.
- BRE fee validation in ante.
- Contract deployment authorization before EVM create execution.
- Governance authority checks during message handling.


### 8.2 Query lifecycle

Queries SHOULD follow this pattern:

```text
Client
  -> Public API or Wallet API
  -> Full Node Query Interface or Index DB
  -> Response
```

Consensus-critical data SHOULD come directly from node queries where correctness is essential. Derived or aggregated UX data MAY come from indexers.

## 9. Node roles

The MVP SHOULD define distinct node roles.

### 9.1 Validator node

A validator node:

- Participates in consensus.
- Holds full state.
- Signs votes and proposals.
- May expose restricted operator interfaces.

A validator node SHOULD NOT serve as the main public internet-facing RPC endpoint unless necessary.

### 9.2 Public full node

A public full node:

- Tracks full chain state.
- Exposes query and transaction endpoints.
- Does not hold validator signing keys.

This node role SHOULD be horizontally scalable.

### 9.3 Seed or peer-discovery node

The network MAY include dedicated seed infrastructure for peer bootstrapping. These nodes SHOULD NOT require validator authority.

### 9.4 Indexer node

An indexer node:

- Subscribes to chain events and blocks.
- Writes normalized data into query databases.
- Powers explorer and wallet views.

This role SHOULD be operationally isolated from consensus-critical infrastructure.

## 10. Validator infrastructure

Validator infrastructure SHOULD separate:

- Consensus process.
- Application process.
- Signing keys.
- Monitoring and alerting.

Recommended shape:

```text
+-----------------------------------+
| Validator Host                    |
|  +-----------------------------+  |
|  | CometBFT / Node Process     |  |
|  +-----------------------------+  |
|  | Application Process         |  |
|  | cored                       |  |
|  +-----------------------------+  |
|  | Remote Signer Client        |  |
|  +-----------------------------+  |
+----------------+------------------+
                 |
        +--------v--------+
        | Remote Signer   |
        | HSM / isolated  |
        +-----------------+
```

Validator signing keys SHOULD be isolated from the main node runtime. Remote signers or HSM-backed signing SHOULD be preferred.

## 11. Service architecture

The MVP SHOULD include the following off-chain services.

### 11.1 Public API gateway

Purpose:

- Present stable public endpoints.
- Route traffic to public nodes.
- Provide rate limiting, request shaping, and observability.

This gateway MAY front REST, gRPC-web, and JSON-RPC.

### 11.2 Wallet API service

Purpose:

- Expose wallet-friendly, normalized views.
- Hide low-level chain-specific complexity.
- Provide fee quotes and chain metadata.

The wallet API SHOULD remain read-focused in MVP. Transaction signing SHOULD remain client-side.

### 11.3 Explorer indexer

Purpose:

- Build a searchable representation of accounts, blocks, validators, and events.
- Support explorer UI and operator troubleshooting.


### 11.4 Supply and metrics service

Purpose:

- Track BRE and TABB supply.
- Track validator and fee metrics.
- Support governance and operational dashboards.


## 12. API architecture

The public interface SHOULD have three main surfaces.

### 12.1 JSON-RPC

Used primarily for EVM wallet compatibility.

Expected consumers:

- Browser wallets.
- Web3 libraries.
- Contract tooling.


### 12.2 gRPC / REST

Used primarily for:

- Native chain queries.
- Wallet API backends.
- Operational tools.
- Explorer services.


### 12.3 Wallet API

Used for:

- Chain info.
- Account overview.
- Balance views.
- Fee quote views.
- Validator metadata.

This layer SHOULD present stable application-level semantics even if lower-level chain APIs evolve.

## 13. Data architecture

The architecture SHOULD separate canonical data from derived data.

### 13.1 Canonical data

Canonical state includes:

- Token balances.
- Validator state.
- Governance proposals and votes.
- EVM contract state.
- BRE authority state.
- Treasury authority state.
- Registry metadata.

Canonical data MUST live on-chain.

### 13.2 Derived data

Derived data includes:

- Explorer summaries.
- Historical materialized views.
- Wallet dashboards.
- Performance metrics.
- Search indexes.

Derived data SHOULD live off-chain and SHOULD be rebuildable from chain history.

## 14. Authority domains

The architecture SHOULD define explicit authority domains.

### 14.1 Consensus authority

Controls:

- Block production.
- Validator participation.
- Liveness and safety.


### 14.2 Governance authority

Controls:

- Upgrades.
- Parameters.
- Policy shifts.
- Treasury and monetary authority appointment.


### 14.3 Validator policy authority

Controls:

- Authorized contract deployers.
- Validator-linked privileged permissions.


### 14.4 Monetary authority

Controls:

- BRE issuance policy.
- BRE mint/burn permissions.
- Future monetary parameters.


### 14.5 Treasury authority

Controls:

- Treasury transfers.
- Allocation rules.
- Future buy-burn and distribution operations.

These domains SHOULD be separated even if some are initially governed by the same operational team.

## 15. Security architecture

### 15.1 Trust boundaries

The architecture SHOULD treat the following as separate trust zones:

- Validator signer zone.
- Consensus node zone.
- Public API zone.
- Indexer and analytics zone.
- Governance key or authority zone.


### 15.2 Critical protections

The system SHOULD protect against:

- Unauthorized contract deployment.
- Unauthorized BRE minting.
- Unauthorized treasury transfers.
- Validator key compromise.
- Governance operator misuse.
- Public RPC abuse.


### 15.3 Defensive patterns

The deployment SHOULD use:

- Isolated validator signing.
- Minimal public exposure for validator nodes.
- Network segmentation.
- Principle of least privilege.
- Comprehensive audit logs.
- Alerting on authority actions.


## 16. Event architecture

The Core Chain SHOULD emit stable events with enough structure for explorer and wallet indexing.

Important event classes:

- BRE transfer.
- TABB transfer.
- Staking events.
- Slashing and jailing events.
- Contract deployment events.
- Contract execution events.
- Governance proposal and vote events.
- BRE monetary events.
- Treasury events.
- Feature-flag events.

Events SHOULD include:

- Block height.
- Transaction hash.
- Module.
- Actor address.
- Amount and denom where relevant.
- Status code.


## 17. Configuration architecture

The system SHOULD distinguish between:

- Genesis-fixed values.
- Governance-tunable parameters.
- Node-local operator configuration.


### 17.1 Genesis-fixed examples

These SHOULD include:

- Chain identifiers.
- Base denoms.
- Initial token allocations.
- Initial validator set.
- Initial authority accounts.


### 17.2 Governance-tunable examples

These SHOULD include:

- Validator policy settings.
- Fee parameters.
- Slashing parameters.
- Treasury control parameters.
- BRE monetary policy parameters.


### 17.3 Node-local examples

These SHOULD include:

- RPC binding addresses.
- Peer lists.
- Logging verbosity.
- Metrics endpoints.
- API rate limits.


## 18. Deployment topology

A minimal production-like deployment SHOULD include:

- 4 to 7 validator nodes.
- 2 or more public full nodes.
- 1 or more indexer nodes.
- 1 API gateway tier.
- Shared observability stack.

Example:

```text
                Internet
                   |
         +---------v----------+
         | API Gateway / LB   |
         +----+----------+----+
              |          |
      +-------v--+   +---v--------+
      | Public 1 |   | Public 2   |
      +----+-----+   +------+-----+
           |                |
   +-------+----------------+-------+
   |        Core Chain P2P Network   |
   +--+---------+---------+--------+-+
      |         |         |        |
   +--v--+   +--v--+   +--v--+  +--v--+
   | Val1 |   | Val2 |  | Val3 | | Val4 |
   +----- +   +----- +  +----- + +----- +
```


## 19. Evolution path

The architecture SHOULD support this evolution sequence:

1. Single-chain Core Chain MVP.
2. Public node, explorer, and wallet stabilization.
3. Monetary and treasury authority hardening.
4. Registry and feature-flag activation for multichain readiness.
5. Validator-governed chain spawning.
6. Wallet-led multichain discovery and routing.
7. High-scale sub-chain operations.

Nothing in the MVP architecture SHOULD assume that the single chain remains the only execution domain forever.

## 20. Codebase shape

A matching repository structure SHOULD look broadly like this:

```text
tabbre/
├─ app/
│  ├─ app.go
│  ├─ ante.go
│  ├─ encoding.go
│  ├─ modules.go
│  └─ upgrades/
├─ cmd/
│  └─ cored/
├─ x/
│  ├─ validator_policy/
│  ├─ wallet_query/
│  ├─ core_registry/
│  ├─ bre_monetary/
│  └─ treasury/
├─ proto/
├─ contracts/
│  └─ system/
├─ services/
│  ├─ wallet-api/
│  ├─ explorer-indexer/
│  └─ api-gateway/
├─ deployments/
│  ├─ localnet/
│  ├─ testnet/
│  └─ mainnet/
└─ docs/
   ├─ SPEC.md
   ├─ ARCHITECTURE.md
   ├─ GENESIS.md
   └─ PROTO.md
```

The architecture SHOULD keep consensus-critical code in the chain app and keep UX or aggregation logic in services.

## 21. Operational architecture

The MVP SHOULD support:

- Structured logs.
- Prometheus-style metrics.
- Health endpoints.
- Block lag alerts.
- Validator liveness alerts.
- RPC saturation alerts.
- Supply and treasury action alerts.

Operational runbooks SHOULD exist for:

- Validator replacement.
- Emergency public node rotation.
- Planned software upgrade.
- Governance parameter rollback.
- Authority key rotation.


## 22. Testing architecture

The architecture SHOULD support four testing layers:

- Unit tests for modules.
- Integration tests for runtime interactions.
- End-to-end tests on localnet.
- Staging or testnet validation before mainnet promotion.

The test harness SHOULD include:

- Validator-policy tests.
- BRE fee enforcement tests.
- TABB staking tests.
- JSON-RPC wallet compatibility tests.
- Upgrade rehearsal tests.


## 23. Acceptance view

The architecture MAY be considered successfully implemented when:

- Validator nodes, public nodes, and services can be deployed independently.
- TABB staking and BRE fee logic are enforced end to end.
- Contract deployment is restricted by validator policy.
- Wallet and explorer traffic are served from public infrastructure without exposing validator keys.
- Governance and authority flows are operational.
- The system is structurally ready for later multichain expansion.


## 24. Non-goals

This architecture does not attempt to define:

- Full sub-chain runtime architecture.
- Mature BRE credit market mechanics.
- UBI distribution architecture.
- Energy market settlement architecture.
- Final production decentralization policy.

Those SHOULD be covered in later documents.

If you want, I can next create the matching `GENESIS.md` or `PROTO.md`.

