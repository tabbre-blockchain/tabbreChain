<img src="https://r2cdn.perplexity.ai/pplx-full-logo-primary-dark%402x.png" style="height:64px;margin-right:32px"/>

# Core Chain MVP GENESIS.md

## 1. Status

This document defines the genesis design for the Tabbre Core Chain MVP. It accompanies `SPEC.md` and `ARCHITECTURE.md` and describes the required initial state, allocation structure, authority accounts, validator bootstrap, and parameter defaults for the first live single-chain Tabbre deployment.

Where this document uses MUST, MUST NOT, REQUIRED, SHOULD, SHOULD NOT, and MAY, those words are normative.

## 2. Purpose

The genesis file initializes the first Tabbre Core Chain as a single-chain proof-of-stake ledger with:

- TABB as the staking asset.
- BRE as the transaction fee and transfer asset.
- An initial BRE allocation to Tabbre-controlled accounts.
- A permissioned validator set.
- Validator-restricted smart-contract deployment.
- Reserved authority domains for governance, treasury, and future BRE monetary control.

The genesis design SHOULD favor operational simplicity and clean forward migration to the later multichain architecture.

## 3. Design principles

The genesis state SHOULD follow these principles:

### 3.1 Fixed identity

The chain identity, EVM chain ID, base denoms, and address conventions MUST be stable from genesis onward.

### 3.2 Role separation

Genesis accounts and module authorities SHOULD separate:

- operational funds,
- reserve funds,
- validator stake,
- governance authority,
- treasury authority,
- monetary authority.


### 3.3 Minimal mutable assumptions

Anything that is expected to change over time SHOULD be parameterized or governed rather than hard-coded into application logic.

### 3.4 Future compatibility

Genesis naming, metadata, and authority layout SHOULD leave room for:

- future treasury activation,
- future BRE monetary policy activation,
- future multichain registry activation.


## 4. Genesis components

The Core Chain genesis MUST define the following categories of state:

- Chain identity.
- Consensus and validator bootstrap state.
- Native token metadata and supply.
- Account allocations.
- Governance authority.
- Validator policy state.
- EVM runtime configuration.
- Fee and staking parameters.
- Reserved module authorities for treasury and BRE monetary control.
- Core registry metadata.


## 5. Chain identity

The genesis file MUST define a unique chain identity for the Tabbre Core Chain.

### 5.1 Required fields

The genesis MUST include:

- `chain_id`
- `genesis_time`
- `initial_height`
- `evm_chain_id`
- `network_name`
- `environment`


### 5.2 Recommended values

Recommended naming shape:

```text
chain_id: tabbre-core-1
network_name: Tabbre Core Chain
environment: mainnet
```

For non-production deployments, test identifiers SHOULD be distinct, for example:

- `tabbre-core-dev-1`
- `tabbre-core-testnet-1`

The EVM chain ID MUST be unique and MUST NOT be reused for another incompatible network.

## 6. Denominations

The genesis MUST define two primary native denoms:

- `utabb` as the base unit of TABB.
- `ubre` as the base unit of BRE.

Optional display metadata MAY define:

- `TABB` = $10^6$ `utabb`
- `BRE` = $10^6$ `ubre`

If a different decimal precision is chosen, it MUST be applied consistently across:

- bank balances,
- staking,
- governance display,
- wallet APIs,
- explorer indexing.


## 7. Token supply model

### 7.1 TABB supply

TABB MUST have a fixed total supply defined at genesis. The genesis state MUST mint the full canonical TABB supply at chain creation.

Recommended canonical total supply:

```text
1,000,000,000 TABB
```

If using 6 decimal places, the base-unit supply is:

```text
1,000,000,000,000,000 utabb
```

The chain MUST preserve the fixed-supply model for TABB.

### 7.2 BRE supply

BRE MUST be initialized with a defined starting supply on the Core Chain and allocated to accounts under Tabbre Project control.

The exact initial BRE amount is a launch policy decision. It SHOULD be sufficient to support:

- validator and operator fee needs,
- initial operational transfers,
- ecosystem testing,
- public node and wallet usage.

The genesis MUST NOT leave BRE supply undefined.

## 8. TABB allocation model

The genesis SHOULD represent the tokenomics allocation categories as separate accounts or vesting buckets.

Recommended allocation model:


| Category | Percent | Quantity |
| :-- | --: | --: |
| Private Sale | 20% | 200,000,000 TABB |
| Public Sale | 10% | 100,000,000 TABB |
| Founders | 10% | 100,000,000 TABB |
| Community Development | 10% | 100,000,000 TABB |
| Reserve | 50% | 500,000,000 TABB |

If vesting support is available at launch:

- Founder allocations SHOULD be vested.
- Reserve allocations SHOULD follow a long-duration release schedule.
- Community allocations MAY be controlled by a designated community account or vesting authority.

If vesting support is not available at launch:

- each category MUST still be assigned to a distinct genesis account or custody account,
- and a governance or operational migration plan SHOULD exist for later conversion to vesting controls.


## 9. Required genesis account classes

The genesis SHOULD define the following account classes.

### 9.1 Foundation operational account

Purpose:

- pay for operational expenses,
- fund initial chain operations,
- seed ecosystem activities.

Typical balances:

- BRE operational balance,
- optional limited TABB balance.


### 9.2 BRE supply custody account

Purpose:

- hold initial BRE supply before later policy activation,
- provide liquidity or distribution capacity under governance or operational control.


### 9.3 Reserve TABB account

Purpose:

- hold reserve-category TABB,
- support future collateral, treasury, or market operations.


### 9.4 Private sale custody account

Purpose:

- hold or distribute private sale allocations.


### 9.5 Public sale custody account

Purpose:

- hold or distribute public sale allocations.


### 9.6 Founder allocation accounts

Purpose:

- receive founder allocations,
- ideally under vesting or restricted custody.


### 9.7 Community development account

Purpose:

- ecosystem growth,
- developer grants,
- partner incentives.


### 9.8 Treasury authority account

Purpose:

- future treasury governance and controlled treasury actions.

This account MAY begin dormant but MUST be initialized if the treasury module is present.

### 9.9 Monetary authority account

Purpose:

- future BRE issuance and policy control.

This account MAY begin dormant but MUST be initialized if the monetary module is present.

### 9.10 Governance authority account

Purpose:

- module authority where modules require explicit admin ownership.
- initial governance transition support if needed.


### 9.11 Validator self-bond accounts

Purpose:

- hold TABB used as self-bond for initial validators,
- pay BRE fees as needed.


## 10. Validator bootstrap

### 10.1 Initial validator set

The genesis MUST include an initial permissioned validator set.

Each validator entry MUST define:

- operator address,
- consensus public key,
- self-bond amount in TABB,
- status,
- commission configuration,
- metadata such as moniker.


### 10.2 Validator count

A small initial validator set is RECOMMENDED for MVP stability.

Suggested ranges:

- localnet: 1–4 validators
- testnet: 4–7 validators
- early mainnet: 4–11 validators

These are operational recommendations, not protocol invariants.

### 10.3 Self-bonding

Each genesis validator MUST have:

- sufficient TABB balance to cover self-bond,
- sufficient BRE balance to cover operational transaction fees.


### 10.4 Validator account mapping

Genesis MUST initialize the one-to-one mapping between validator operators and validator accounts required by validator policy.

## 11. Validator policy genesis

The validator policy module MUST be initialized so that contract deployment is restricted to validator-authorized accounts from block 1.

Required initial state:

- validator-policy enabled,
- deployer restriction enabled,
- all initial validator accounts marked as authorized deployers or linked to deployer eligibility,
- non-validator accounts unauthorized by default.

Recommended shape:

```yaml
validator_policy:
  enabled: true
  validator_only_contract_deploy: true
  authorized_deployers:
    - val_account_1
    - val_account_2
    - val_account_3
```


## 12. Bank balances

The genesis MUST initialize bank balances for all required accounts.

### 12.1 Minimum balance requirements

The following account classes SHOULD hold BRE at genesis:

- validator accounts,
- foundation operational account,
- public infrastructure operator accounts,
- treasury authority if treasury actions may occur,
- monetary authority if BRE mint/burn actions may occur.

The following account classes SHOULD hold TABB at genesis where applicable:

- validator self-bond accounts,
- reserve account,
- allocation custody accounts,
- founder and community accounts.


### 12.2 Dust prevention

Genesis allocations SHOULD avoid creating large numbers of trivial dust balances.

### 12.3 Operational sufficiency

Validator operators SHOULD have enough BRE at genesis to avoid immediate operational deadlock from fee starvation.

## 13. Governance genesis

The governance module MUST be initialized at genesis.

Recommended governance settings SHOULD include:

- proposal submission enabled,
- voting enabled,
- parameter change proposal support,
- software upgrade proposal support,
- expedited or emergency proposal support only if governance design is mature.


### 13.1 Voting power basis

The governance system SHOULD derive voting power from TABB staking or TABB-linked stake weight.

### 13.2 Initial authority

Module authority SHOULD default to governance where feasible rather than to a permanently centralized admin account.

If direct governance control is not feasible at launch, temporary admin authority:

- MUST be explicit,
- MUST be documented,
- SHOULD be migratable to governance later.


## 14. Staking genesis

The staking module MUST be initialized with TABB as the bond denom.

Recommended initial staking parameters:

- bond denom: `utabb`
- unbonding time: 14 days
- max validators: small initial value suitable for MVP
- max entries: standard SDK-compatible value
- historical entries: enabled for explorer and slashing support

These values MAY be adjusted by governance later.

## 15. Slashing genesis

The slashing module MUST be initialized and enabled.

Recommended initial slashing configuration SHOULD include:

- downtime jail duration,
- slash fraction for downtime,
- slash fraction for double-signing,
- signed blocks window,
- minimum signed per window.

The slashing denom effect MUST apply to TABB stake.

## 16. Distribution genesis

The distribution module MUST be initialized to distribute BRE-denominated fees collected by the chain.

The distribution design SHOULD support:

- validator rewards,
- delegator rewards if delegation is enabled,
- governance-tunable community or burn share if later desired.

If a fee-burning component is included, it MUST be explicitly parameterized.

## 17. Fee market genesis

The genesis MUST initialize BRE as the required fee denom for normal transactions.

Recommended fee-market parameters:

- base fee enabled,
- minimum gas price or equivalent enabled,
- priority fee optional,
- fee quote compatibility for wallet APIs.

Genesis SHOULD set conservative defaults to prevent spam while keeping early testing usable.

## 18. EVM genesis

The EVM runtime MUST be initialized at genesis.

Required EVM fields:

- chain ID,
- gas schedule or EVM config,
- base fee configuration if applicable,
- predeploy reservations if system contracts are planned.


### 18.1 Predeploys

The genesis MAY reserve or initialize system contracts for:

- governance hooks,
- treasury registry,
- monetary authority registry,
- protocol metadata.

If no predeploys are included in MVP, reserved addresses SHOULD still be documented to avoid future collisions.

## 19. Wallet query genesis

The wallet query module SHOULD be initialized with:

- chain metadata,
- supported denoms,
- display metadata,
- feature flags,
- environment identifiers.

Although the MVP is single-chain, the query schema SHOULD present a future-compatible structure.

## 20. Core registry genesis

The core registry module SHOULD initialize:

- core chain metadata,
- network identity,
- reserved future chain namespace,
- multichain feature flags disabled by default.

Recommended flags:

- `multichain_enabled = false`
- `spawn_enabled = false`
- `split_enabled = false`

This keeps the Core Chain ready for future activation without exposing unfinished functionality.

## 21. BRE monetary module genesis

If the `bre_monetary` module is included, genesis MUST define:

- authority account,
- enabled/disabled state,
- mint permissions,
- burn permissions,
- optional placeholder policy parameters.

Recommended initial state:

- module present,
- public issuance disabled,
- authority-only mint/burn enabled or fully disabled,
- policy mode set to placeholder.

Example:

```yaml
bre_monetary:
  enabled: true
  authority: tabbre1...
  public_mint: false
  public_burn: false
  policy_mode: bootstrap
```


## 22. Treasury module genesis

If the `treasury` module is included, genesis MUST define:

- treasury authority account,
- enabled/disabled state,
- optional treasury balances,
- placeholder allocation policy.

Recommended initial state:

- treasury present,
- discretionary transfer disabled unless needed,
- governance migration path documented.

Example:

```yaml
treasury:
  enabled: true
  authority: tabbre1...
  transfers_enabled: false
  allocation_policy: bootstrap
```


## 23. Module authorities

Each module requiring elevated authority MUST have a clearly defined genesis authority.

Recommended mapping:

- governance-controlled modules → governance authority
- treasury module → treasury authority or governance
- BRE monetary module → monetary authority or governance
- validator policy module → governance authority
- registry module → governance authority

Authority mappings MUST be explicit and MUST NOT rely on undocumented defaults.

## 24. Metadata and display config

The genesis SHOULD include human-readable metadata for:

- network name,
- token names and symbols,
- decimals,
- explorer references,
- wallet display names.

Recommended asset metadata:

- TABB: symbol `TABB`, display `TABB`, base `utabb`
- BRE: symbol `BRE`, display `BRE`, base `ubre`


## 25. Example logical genesis layout

The following is a non-binding logical layout example:

```yaml
chain:
  chain_id: tabbre-core-1
  evm_chain_id: 782001
  network_name: Tabbre Core Chain
  environment: mainnet

assets:
  - base: utabb
    display: TABB
    exponent: 6
  - base: ubre
    display: BRE
    exponent: 6

supply:
  utabb_total: 1000000000000000
  ubre_total: <launch_defined>

accounts:
  - foundation_ops
  - bre_custody
  - tabb_reserve
  - private_sale_custody
  - public_sale_custody
  - founders_1
  - founders_2
  - community_dev
  - treasury_authority
  - monetary_authority
  - governance_authority
  - validator_1
  - validator_2
  - validator_3
  - validator_4

staking:
  bond_denom: utabb
  validators:
    - validator_1
    - validator_2
    - validator_3
    - validator_4

fees:
  fee_denom: ubre

validator_policy:
  validator_only_contract_deploy: true

registry:
  multichain_enabled: false
  spawn_enabled: false
  split_enabled: false
```


## 26. Invariants

The genesis builder and validation logic MUST enforce the following invariants:

1. The sum of all TABB account balances plus bonded TABB MUST equal total TABB supply.
2. The sum of all BRE account balances MUST equal total BRE supply.
3. Every genesis validator MUST have a valid consensus key.
4. Every genesis validator MUST have sufficient TABB self-bond.
5. Every genesis validator MUST have a mapped validator account.
6. Contract deployment restriction MUST be enabled by default.
7. No non-authorized account may appear in the authorized deployer set unless explicitly intended.
8. Bond denom MUST be `utabb`.
9. Fee denom MUST be `ubre`.
10. Chain ID and EVM chain ID MUST both be set.

## 27. Validation checklist

Before accepting a genesis file for deployment, the release process SHOULD verify:

- Chain ID is correct for the target environment.
- EVM chain ID is unique and correct.
- TABB supply matches intended canonical issuance.
- TABB allocations match intended category totals.
- BRE starting supply is defined and allocated.
- Validator set matches approved operator list.
- Validator self-bond amounts are correct.
- Validator-policy restrictions are enabled.
- Governance authority mappings are correct.
- Treasury and monetary authority accounts are correct.
- Multichain features are disabled.
- Public test accounts, if any, are removed from mainnet genesis.
- Metadata displays correctly in wallet and explorer test environments.


## 28. Environment profiles

### 28.1 Localnet

Localnet genesis SHOULD:

- use small supplies,
- use 1–4 validators,
- enable rapid blocks,
- allow easier testing and account funding.


### 28.2 Testnet

Testnet genesis SHOULD:

- use realistic parameter values,
- use multiple validators,
- include public testing accounts where necessary,
- mirror mainnet layout where possible.


### 28.3 Mainnet

Mainnet genesis MUST:

- use final chain identifiers,
- use final authority accounts,
- use audited validator keys,
- use finalized token allocations,
- exclude temporary testing permissions.


## 29. Migration considerations

Genesis design SHOULD minimize pain for future upgrades.

To support later migration:

- address formats MUST remain unchanged,
- authority accounts SHOULD remain stable or cleanly upgradable,
- registry metadata SHOULD already reflect the Core Chain’s root role,
- token metadata SHOULD remain immutable where possible.

Future migrations MAY add:

- vesting conversions,
- treasury activation,
- BRE policy activation,
- multichain registry population,
- chain factory logic.


## 30. Acceptance criteria

A genesis file SHALL be considered compliant with this document when:

- it initializes a valid single-chain Core Chain,
- TABB is the staking denom,
- BRE is the fee denom,
- total TABB and BRE supplies are fully defined,
- initial BRE is allocated to Tabbre-controlled accounts,
- validator set is permissioned and bonded,
- contract deployment is restricted to validator-authorized accounts,
- governance is initialized,
- treasury and BRE monetary authority domains are reserved if those modules exist,
- multichain features remain disabled by default.


## 31. Recommended companion files

This genesis document SHOULD be used alongside:

- `SPEC.md`
- `ARCHITECTURE.md`
- `CHAIN_PARAMS.md`
- `ALLOCATIONS.csv`
- `VALIDATORS.json`
- `AUTHORITIES.json`


## 32. Non-normative next step

A practical next step is to convert this document into:

- a concrete `genesis.template.json`,
- an `ALLOCATIONS.csv`,
- and a `VALIDATORS.json` schema so the genesis build can be automated.

If you want, I can next create the matching `PROTO.md`.

