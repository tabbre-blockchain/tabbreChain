<img src="https://r2cdn.perplexity.ai/pplx-full-logo-primary-dark%402x.png" style="height:64px;margin-right:32px"/>

# Core Chain MVP PROTO.md

## 1. Status

This document defines the protobuf and API contract design for the Tabbre Core Chain MVP. It accompanies `SPEC.md`, `ARCHITECTURE.md`, and `GENESIS.md` and describes the message schemas, gRPC services, query surfaces, naming conventions, and versioning strategy for the custom Core Chain modules.

This document is normative for protobuf package layout, service boundaries, field intent, and compatibility expectations. Implementation details may vary, but externally exposed APIs SHOULD conform to this document.

## 2. Purpose

The Core Chain MVP requires a stable API contract for:

- validator policy enforcement,
- wallet-facing queries,
- core registry metadata,
- BRE monetary authority scaffolding,
- treasury scaffolding.

The protobuf design SHOULD:

- be small and explicit,
- be forward-compatible with future multichain expansion,
- separate write services from read services,
- avoid field naming that will break when sub-chains are introduced later.


## 3. Design principles

### 3.1 Versioned packages

All protobuf packages MUST be versioned from the start.

### 3.2 Stable field numbering

Once released, field numbers MUST NOT be reused for different meanings.

### 3.3 Append-only evolution

New fields SHOULD be added with new field numbers. Existing fields SHOULD NOT change semantic meaning.

### 3.4 Read/write separation

Write operations SHOULD be exposed through `Msg` services. Read operations SHOULD be exposed through `Query` services.

### 3.5 Future-compatible naming

Single-chain MVP APIs SHOULD use names that remain valid once the Core Chain governs many sub-chains.

## 4. Directory layout

The repository SHOULD use the following protobuf structure:

```text
proto/
└─ tabbre/
   ├─ validator_policy/v1/
   │  ├─ tx.proto
   │  ├─ query.proto
   │  ├─ types.proto
   │  └─ genesis.proto
   ├─ wallet_query/v1/
   │  ├─ query.proto
   │  ├─ types.proto
   │  └─ genesis.proto
   ├─ core_registry/v1/
   │  ├─ tx.proto
   │  ├─ query.proto
   │  ├─ types.proto
   │  └─ genesis.proto
   ├─ bre_monetary/v1/
   │  ├─ tx.proto
   │  ├─ query.proto
   │  ├─ types.proto
   │  └─ genesis.proto
   └─ treasury/v1/
      ├─ tx.proto
      ├─ query.proto
      ├─ types.proto
      └─ genesis.proto
```


## 5. Package naming

Each module MUST use a package name in this shape:

```proto
package tabbre.<module>.v1;
```

Examples:

```proto
package tabbre.validator_policy.v1;
package tabbre.wallet_query.v1;
package tabbre.core_registry.v1;
package tabbre.bre_monetary.v1;
package tabbre.treasury.v1;
```

The Go package option SHOULD follow this shape:

```proto
option go_package = "github.com/tabbre/tabbre/x/<module>/types";
```


## 6. Common conventions

### 6.1 Address fields

All account-like addresses SHOULD be encoded as strings in public protobuf messages.

Field naming SHOULD distinguish:

- `address` for generic account address,
- `authority` for module admin authority,
- `operator_address` for validator operator identity,
- `validator_address` only where consensus/staking context requires that distinction.


### 6.2 Amount fields

Amounts SHOULD be encoded as strings to avoid integer overflow and preserve SDK compatibility.

Examples:

- `amount = "1000000"`
- `balance = "250000000"`


### 6.3 Denom fields

All coin-bearing structures SHOULD include an explicit `denom`.

### 6.4 Timestamps

Timestamps SHOULD use `google.protobuf.Timestamp`.

### 6.5 Pagination

List queries SHOULD use Cosmos-style pagination request and response types where available.

### 6.6 Enums

Enums SHOULD be used for compact protocol state where the state machine is closed and well-defined.

## 7. Shared types

The following shared message patterns SHOULD be used across modules.

### 7.1 Coin

```proto
message Coin {
  string denom = 1;
  string amount = 2;
}
```


### 7.2 FeatureFlag

```proto
message FeatureFlag {
  string name = 1;
  bool enabled = 2;
}
```


### 7.3 ChainRef

```proto
message ChainRef {
  string chain_id = 1;
  uint64 latest_height = 2;
  string network_name = 3;
  bool is_core = 4;
}
```


### 7.4 AccountOverview

```proto
message AccountOverview {
  string address = 1;
  repeated Coin balances = 2;
  uint64 sequence = 3;
  uint64 account_number = 4;
  uint64 evm_nonce = 5;
}
```


## 8. Validator Policy module

## 8.1 Scope

This module controls validator-linked authorization for contract deployment and other privileged actions.

### 8.2 `types.proto`

```proto
syntax = "proto3";

package tabbre.validator_policy.v1;
option go_package = "github.com/tabbre/tabbre/x/validator_policy/types";

message Params {
  bool enabled = 1;
  bool validator_only_contract_deploy = 2;
}

message ValidatorAccountLink {
  string operator_address = 1;
  string account_address = 2;
  bool active = 3;
}

message AuthorizedDeployer {
  string address = 1;
  string linked_operator_address = 2;
  bool active = 3;
}

message PolicyStatus {
  Params params = 1;
  uint64 validator_link_count = 2;
  uint64 authorized_deployer_count = 3;
}
```


### 8.3 `tx.proto`

```proto
syntax = "proto3";

package tabbre.validator_policy.v1;
option go_package = "github.com/tabbre/tabbre/x/validator_policy/types";

service Msg {
  rpc RegisterValidatorAccount(MsgRegisterValidatorAccount)
      returns (MsgRegisterValidatorAccountResponse);

  rpc UpdateParams(MsgUpdateParams)
      returns (MsgUpdateParamsResponse);

  rpc AuthorizeDeployer(MsgAuthorizeDeployer)
      returns (MsgAuthorizeDeployerResponse);

  rpc RevokeDeployer(MsgRevokeDeployer)
      returns (MsgRevokeDeployerResponse);
}

message MsgRegisterValidatorAccount {
  string authority = 1;
  string operator_address = 2;
  string account_address = 3;
}

message MsgRegisterValidatorAccountResponse {}

message MsgUpdateParams {
  string authority = 1;
  Params params = 2;
}

message MsgUpdateParamsResponse {}

message MsgAuthorizeDeployer {
  string authority = 1;
  string address = 2;
  string linked_operator_address = 3;
}

message MsgAuthorizeDeployerResponse {}

message MsgRevokeDeployer {
  string authority = 1;
  string address = 2;
}

message MsgRevokeDeployerResponse {}
```


### 8.4 `query.proto`

```proto
syntax = "proto3";

package tabbre.validator_policy.v1;
option go_package = "github.com/tabbre/tabbre/x/validator_policy/types";

import "cosmos/base/query/v1beta1/pagination.proto";

service Query {
  rpc Params(QueryParamsRequest) returns (QueryParamsResponse);
  rpc PolicyStatus(QueryPolicyStatusRequest) returns (QueryPolicyStatusResponse);
  rpc ValidatorAccount(QueryValidatorAccountRequest) returns (QueryValidatorAccountResponse);
  rpc AuthorizedDeployer(QueryAuthorizedDeployerRequest) returns (QueryAuthorizedDeployerResponse);
  rpc AuthorizedDeployers(QueryAuthorizedDeployersRequest) returns (QueryAuthorizedDeployersResponse);
}

message QueryParamsRequest {}
message QueryParamsResponse {
  Params params = 1;
}

message QueryPolicyStatusRequest {}
message QueryPolicyStatusResponse {
  PolicyStatus status = 1;
}

message QueryValidatorAccountRequest {
  string operator_address = 1;
}

message QueryValidatorAccountResponse {
  ValidatorAccountLink link = 1;
}

message QueryAuthorizedDeployerRequest {
  string address = 1;
}

message QueryAuthorizedDeployerResponse {
  AuthorizedDeployer deployer = 1;
}

message QueryAuthorizedDeployersRequest {
  cosmos.base.query.v1beta1.PageRequest pagination = 1;
}

message QueryAuthorizedDeployersResponse {
  repeated AuthorizedDeployer deployers = 1;
  cosmos.base.query.v1beta1.PageResponse pagination = 2;
}
```


### 8.5 `genesis.proto`

```proto
syntax = "proto3";

package tabbre.validator_policy.v1;
option go_package = "github.com/tabbre/tabbre/x/validator_policy/types";

message GenesisState {
  Params params = 1;
  repeated ValidatorAccountLink validator_links = 2;
  repeated AuthorizedDeployer authorized_deployers = 3;
}
```


## 9. Wallet Query module

## 9.1 Scope

This module exposes wallet-friendly read APIs for account, chain, validator, and fee data. It is read-only in MVP.

### 9.2 `types.proto`

```proto
syntax = "proto3";

package tabbre.wallet_query.v1;
option go_package = "github.com/tabbre/tabbre/x/wallet_query/types";

import "google/protobuf/timestamp.proto";

message Coin {
  string denom = 1;
  string amount = 2;
}

message ChainInfo {
  string chain_id = 1;
  uint64 evm_chain_id = 2;
  string network_name = 3;
  string environment = 4;
  uint64 latest_height = 5;
  string staking_denom = 6;
  string fee_denom = 7;
  bool multichain_enabled = 8;
  google.protobuf.Timestamp latest_block_time = 9;
}

message AccountOverview {
  string address = 1;
  repeated Coin balances = 2;
  uint64 sequence = 3;
  uint64 account_number = 4;
  uint64 evm_nonce = 5;
}

message ValidatorMeta {
  string operator_address = 1;
  string account_address = 2;
  string moniker = 3;
  string status = 4;
  string tokens = 5;
  string commission_rate = 6;
  bool jailed = 7;
}

message FeeQuote {
  string denom = 1;
  string base_fee = 2;
  string suggested_gas_price = 3;
  string suggested_priority_fee = 4;
}

message FeatureFlag {
  string name = 1;
  bool enabled = 2;
}
```


### 9.3 `query.proto`

```proto
syntax = "proto3";

package tabbre.wallet_query.v1;
option go_package = "github.com/tabbre/tabbre/x/wallet_query/types";

import "cosmos/base/query/v1beta1/pagination.proto";

service Query {
  rpc ChainInfo(QueryChainInfoRequest) returns (QueryChainInfoResponse);
  rpc AccountOverview(QueryAccountOverviewRequest) returns (QueryAccountOverviewResponse);
  rpc Balance(QueryBalanceRequest) returns (QueryBalanceResponse);
  rpc Balances(QueryBalancesRequest) returns (QueryBalancesResponse);
  rpc Validators(QueryValidatorsRequest) returns (QueryValidatorsResponse);
  rpc FeeQuote(QueryFeeQuoteRequest) returns (QueryFeeQuoteResponse);
  rpc FeatureFlags(QueryFeatureFlagsRequest) returns (QueryFeatureFlagsResponse);
}

message QueryChainInfoRequest {}

message QueryChainInfoResponse {
  ChainInfo chain = 1;
}

message QueryAccountOverviewRequest {
  string address = 1;
}

message QueryAccountOverviewResponse {
  AccountOverview account = 1;
}

message QueryBalanceRequest {
  string address = 1;
  string denom = 2;
}

message QueryBalanceResponse {
  Coin balance = 1;
}

message QueryBalancesRequest {
  string address = 1;
}

message QueryBalancesResponse {
  repeated Coin balances = 1;
}

message QueryValidatorsRequest {
  cosmos.base.query.v1beta1.PageRequest pagination = 1;
}

message QueryValidatorsResponse {
  repeated ValidatorMeta validators = 1;
  cosmos.base.query.v1beta1.PageResponse pagination = 2;
}

message QueryFeeQuoteRequest {
  string tx_type = 1;
}

message QueryFeeQuoteResponse {
  FeeQuote quote = 1;
}

message QueryFeatureFlagsRequest {}

message QueryFeatureFlagsResponse {
  repeated FeatureFlag flags = 1;
}
```


### 9.4 `genesis.proto`

```proto
syntax = "proto3";

package tabbre.wallet_query.v1;
option go_package = "github.com/tabbre/tabbre/x/wallet_query/types";

message GenesisState {
  string network_name = 1;
  string environment = 2;
  uint64 evm_chain_id = 3;
  string staking_denom = 4;
  string fee_denom = 5;
}
```


## 10. Core Registry module

## 10.1 Scope

This module stores Core Chain metadata and future multichain registry scaffolding.

### 10.2 `types.proto`

```proto
syntax = "proto3";

package tabbre.core_registry.v1;
option go_package = "github.com/tabbre/tabbre/x/core_registry/types";

message Params {
  bool multichain_enabled = 1;
  bool spawn_enabled = 2;
  bool split_enabled = 3;
}

message CoreMetadata {
  string chain_id = 1;
  uint64 evm_chain_id = 2;
  string network_name = 3;
  bool is_core_chain = 4;
}

message ReservedChainId {
  string chain_id = 1;
  string description = 2;
  bool active = 3;
}

message RegisteredChain {
  string chain_id = 1;
  string chain_type = 2;
  string status = 3;
  uint64 activation_height = 4;
}
```


### 10.3 `tx.proto`

```proto
syntax = "proto3";

package tabbre.core_registry.v1;
option go_package = "github.com/tabbre/tabbre/x/core_registry/types";

service Msg {
  rpc UpdateParams(MsgUpdateParams)
      returns (MsgUpdateParamsResponse);

  rpc RegisterReservedChainId(MsgRegisterReservedChainId)
      returns (MsgRegisterReservedChainIdResponse);

  rpc RemoveReservedChainId(MsgRemoveReservedChainId)
      returns (MsgRemoveReservedChainIdResponse);
}

message MsgUpdateParams {
  string authority = 1;
  Params params = 2;
}

message MsgUpdateParamsResponse {}

message MsgRegisterReservedChainId {
  string authority = 1;
  string chain_id = 2;
  string description = 3;
}

message MsgRegisterReservedChainIdResponse {}

message MsgRemoveReservedChainId {
  string authority = 1;
  string chain_id = 2;
}

message MsgRemoveReservedChainIdResponse {}
```


### 10.4 `query.proto`

```proto
syntax = "proto3";

package tabbre.core_registry.v1;
option go_package = "github.com/tabbre/tabbre/x/core_registry/types";

import "cosmos/base/query/v1beta1/pagination.proto";

service Query {
  rpc Params(QueryParamsRequest) returns (QueryParamsResponse);
  rpc CoreMetadata(QueryCoreMetadataRequest) returns (QueryCoreMetadataResponse);
  rpc ReservedChainIds(QueryReservedChainIdsRequest) returns (QueryReservedChainIdsResponse);
  rpc RegisteredChains(QueryRegisteredChainsRequest) returns (QueryRegisteredChainsResponse);
}

message QueryParamsRequest {}
message QueryParamsResponse {
  Params params = 1;
}

message QueryCoreMetadataRequest {}
message QueryCoreMetadataResponse {
  CoreMetadata metadata = 1;
}

message QueryReservedChainIdsRequest {
  cosmos.base.query.v1beta1.PageRequest pagination = 1;
}

message QueryReservedChainIdsResponse {
  repeated ReservedChainId chain_ids = 1;
  cosmos.base.query.v1beta1.PageResponse pagination = 2;
}

message QueryRegisteredChainsRequest {
  cosmos.base.query.v1beta1.PageRequest pagination = 1;
}

message QueryRegisteredChainsResponse {
  repeated RegisteredChain chains = 1;
  cosmos.base.query.v1beta1.PageResponse pagination = 2;
}
```


### 10.5 `genesis.proto`

```proto
syntax = "proto3";

package tabbre.core_registry.v1;
option go_package = "github.com/tabbre/tabbre/x/core_registry/types";

message GenesisState {
  Params params = 1;
  CoreMetadata metadata = 2;
  repeated ReservedChainId reserved_chain_ids = 3;
  repeated RegisteredChain registered_chains = 4;
}
```


## 11. BRE Monetary module

## 11.1 Scope

This module defines the authority boundary and future state for BRE issuance and policy.

### 11.2 `types.proto`

```proto
syntax = "proto3";

package tabbre.bre_monetary.v1;
option go_package = "github.com/tabbre/tabbre/x/bre_monetary/types";

message Params {
  bool enabled = 1;
  bool public_mint_enabled = 2;
  bool public_burn_enabled = 3;
  string policy_mode = 4;
}

message PolicyState {
  string authority = 1;
  string base_rate = 2;
  string target_mode = 3;
  string oracle_config_ref = 4;
}

message SupplyInfo {
  string denom = 1;
  string total_supply = 2;
}
```


### 11.3 `tx.proto`

```proto
syntax = "proto3";

package tabbre.bre_monetary.v1;
option go_package = "github.com/tabbre/tabbre/x/bre_monetary/types";

service Msg {
  rpc UpdateParams(MsgUpdateParams)
      returns (MsgUpdateParamsResponse);

  rpc SetAuthority(MsgSetAuthority)
      returns (MsgSetAuthorityResponse);

  rpc MintBRE(MsgMintBRE)
      returns (MsgMintBREResponse);

  rpc BurnBRE(MsgBurnBRE)
      returns (MsgBurnBREResponse);

  rpc UpdatePolicyState(MsgUpdatePolicyState)
      returns (MsgUpdatePolicyStateResponse);
}

message MsgUpdateParams {
  string authority = 1;
  Params params = 2;
}

message MsgUpdateParamsResponse {}

message MsgSetAuthority {
  string authority = 1;
  string new_authority = 2;
}

message MsgSetAuthorityResponse {}

message MsgMintBRE {
  string authority = 1;
  string recipient = 2;
  string amount = 3;
}

message MsgMintBREResponse {}

message MsgBurnBRE {
  string authority = 1;
  string from_address = 2;
  string amount = 3;
}

message MsgBurnBREResponse {}

message MsgUpdatePolicyState {
  string authority = 1;
  PolicyState policy_state = 2;
}

message MsgUpdatePolicyStateResponse {}
```


### 11.4 `query.proto`

```proto
syntax = "proto3";

package tabbre.bre_monetary.v1;
option go_package = "github.com/tabbre/tabbre/x/bre_monetary/types";

service Query {
  rpc Params(QueryParamsRequest) returns (QueryParamsResponse);
  rpc PolicyState(QueryPolicyStateRequest) returns (QueryPolicyStateResponse);
  rpc Supply(QuerySupplyRequest) returns (QuerySupplyResponse);
}

message QueryParamsRequest {}
message QueryParamsResponse {
  Params params = 1;
}

message QueryPolicyStateRequest {}
message QueryPolicyStateResponse {
  PolicyState policy_state = 1;
}

message QuerySupplyRequest {}
message QuerySupplyResponse {
  SupplyInfo supply = 1;
}
```


### 11.5 `genesis.proto`

```proto
syntax = "proto3";

package tabbre.bre_monetary.v1;
option go_package = "github.com/tabbre/tabbre/x/bre_monetary/types";

message GenesisState {
  Params params = 1;
  PolicyState policy_state = 2;
}
```


## 12. Treasury module

## 12.1 Scope

This module provides the treasury authority boundary and future allocation scaffolding.

### 12.2 `types.proto`

```proto
syntax = "proto3";

package tabbre.treasury.v1;
option go_package = "github.com/tabbre/tabbre/x/treasury/types";

message Params {
  bool enabled = 1;
  bool transfers_enabled = 2;
}

message AllocationPolicy {
  string mode = 1;
  string research_and_development_bps = 2;
  string new_investment_bps = 3;
  string investor_returns_bps = 4;
  string global_ubi_bps = 5;
}

message TreasuryState {
  string authority = 1;
  string treasury_address = 2;
  AllocationPolicy allocation_policy = 3;
}
```


### 12.3 `tx.proto`

```proto
syntax = "proto3";

package tabbre.treasury.v1;
option go_package = "github.com/tabbre/tabbre/x/treasury/types";

service Msg {
  rpc UpdateParams(MsgUpdateParams)
      returns (MsgUpdateParamsResponse);

  rpc SetAuthority(MsgSetAuthority)
      returns (MsgSetAuthorityResponse);

  rpc SetAllocationPolicy(MsgSetAllocationPolicy)
      returns (MsgSetAllocationPolicyResponse);

  rpc TreasuryTransfer(MsgTreasuryTransfer)
      returns (MsgTreasuryTransferResponse);
}

message MsgUpdateParams {
  string authority = 1;
  Params params = 2;
}

message MsgUpdateParamsResponse {}

message MsgSetAuthority {
  string authority = 1;
  string new_authority = 2;
}

message MsgSetAuthorityResponse {}

message MsgSetAllocationPolicy {
  string authority = 1;
  AllocationPolicy allocation_policy = 2;
}

message MsgSetAllocationPolicyResponse {}

message MsgTreasuryTransfer {
  string authority = 1;
  string recipient = 2;
  string denom = 3;
  string amount = 4;
  string memo = 5;
}

message MsgTreasuryTransferResponse {}
```


### 12.4 `query.proto`

```proto
syntax = "proto3";

package tabbre.treasury.v1;
option go_package = "github.com/tabbre/tabbre/x/treasury/types";

service Query {
  rpc Params(QueryParamsRequest) returns (QueryParamsResponse);
  rpc TreasuryState(QueryTreasuryStateRequest) returns (QueryTreasuryStateResponse);
}

message QueryParamsRequest {}
message QueryParamsResponse {
  Params params = 1;
}

message QueryTreasuryStateRequest {}
message QueryTreasuryStateResponse {
  TreasuryState state = 1;
}
```


### 12.5 `genesis.proto`

```proto
syntax = "proto3";

package tabbre.treasury.v1;
option go_package = "github.com/tabbre/tabbre/x/treasury/types";

message GenesisState {
  Params params = 1;
  TreasuryState state = 2;
}
```


## 13. Error model

Custom modules SHOULD return SDK-compatible errors with stable codespaces.

Recommended codespaces:

- `validator_policy`
- `wallet_query`
- `core_registry`
- `bre_monetary`
- `treasury`

Errors SHOULD distinguish:

- unauthorized,
- not found,
- invalid request,
- invalid state,
- feature disabled,
- duplicate entry.


## 14. Events

Each module SHOULD emit stable typed events for indexers and explorers.

Recommended event names:

### 14.1 Validator Policy

- `validator_account_registered`
- `validator_policy_updated`
- `deployer_authorized`
- `deployer_revoked`


### 14.2 Core Registry

- `core_registry_params_updated`
- `reserved_chain_id_registered`
- `reserved_chain_id_removed`


### 14.3 BRE Monetary

- `bre_monetary_params_updated`
- `bre_authority_updated`
- `bre_minted`
- `bre_burned`
- `bre_policy_updated`


### 14.4 Treasury

- `treasury_params_updated`
- `treasury_authority_updated`
- `treasury_allocation_policy_updated`
- `treasury_transfer`

Events SHOULD include enough indexed attributes to support:

- actor address,
- module authority,
- denom,
- amount,
- target address,
- chain metadata where relevant.


## 15. Compatibility rules

### 15.1 Backward compatibility

After a public release:

- field numbers MUST NOT be reused,
- field names SHOULD NOT be renamed unless compatibility wrappers exist,
- required semantic meaning MUST remain stable.


### 15.2 Deprecation

Deprecated fields SHOULD remain readable for at least one major module version cycle.

### 15.3 New versions

Breaking API changes MUST produce a new package version such as:

```proto
package tabbre.validator_policy.v2;
```


## 16. Query design guidance

Queries SHOULD follow these rules:

- singular query for one object by ID or address,
- plural query for list access with pagination,
- no mutation through query services,
- responses SHOULD prefer structured objects over flat ad hoc fields.


## 17. Authority design guidance

All mutating admin messages SHOULD include an `authority` field.

This makes it explicit that:

- the operation is privileged,
- governance or module admin may control it,
- future migration from operational admin to governance remains possible.


## 18. Single-chain to multichain compatibility

The MVP protobuf design SHOULD already anticipate future multichain behavior.

Examples:

- `ChainInfo` SHOULD include chain identity even though only one chain exists now.
- `FeatureFlags` SHOULD expose multichain activation state.
- `core_registry` SHOULD already model reserved and registered chains.
- wallet query schemas SHOULD avoid assumptions that there is only one execution chain forever.


## 19. Suggested code generation

The build pipeline SHOULD generate:

- Go types,
- gRPC server stubs,
- gRPC client stubs,
- gateway bindings where needed,
- OpenAPI output if public REST documentation is desired.

A typical generation command MAY use Buf or standard `protoc` workflows.

## 20. Validation checklist

Before freezing v1 protobufs, the project SHOULD verify:

- package names are versioned,
- all admin messages include authority,
- amount fields are strings,
- addresses are strings,
- field numbers are unique,
- list queries support pagination where needed,
- genesis messages fully represent module state,
- no field names assume impossible future changes.


## 21. Example Buf module layout

A matching Buf-oriented structure MAY look like:

```text
proto/
├─ buf.yaml
├─ buf.gen.yaml
└─ tabbre/
   ├─ validator_policy/v1/
   ├─ wallet_query/v1/
   ├─ core_registry/v1/
   ├─ bre_monetary/v1/
   └─ treasury/v1/
```

Example `buf.gen.yaml` conceptually SHOULD generate:

- Go protobuf types,
- Go gRPC types,
- grpc-gateway output.


## 22. Recommended next files

This document SHOULD be followed by:

- `STORE.md` for KV key layout,
- `EVENTS.md` for event attribute definitions,
- `STATE_MACHINE.md` for exact transition rules,
- `API_EXAMPLES.md` for JSON request/response samples.


## 23. Acceptance criteria

The protobuf layer SHALL be considered complete for MVP when:

- all five custom modules have versioned `types.proto`, `tx.proto`, `query.proto`, and `genesis.proto`,
- admin write messages include explicit authority fields,
- query services cover wallet, registry, policy, treasury, and monetary reads,
- schemas are compatible with the single-chain MVP and future multichain evolution,
- code generation can produce server and client bindings without schema ambiguity.

If you want, I can next create either:

- `STORE.md` with KV prefixes and key layouts, or
- `STATE_MACHINE.md` with per-message transition rules.

