<img src="https://r2cdn.perplexity.ai/pplx-full-logo-primary-dark%402x.png" style="height:64px;margin-right:32px"/>

# My plan is to use an AI Agent to create the Tabbre Chain MVP using Cosmos EVM as the starting point. Please create the prompts that I need for the Agent.

Yes — the best way to do this is to give the agent a **sequenced prompt pack** rather than one giant prompt. Since Tabbre has specific tokenomics and validator rules, each prompt should lock the agent to those constraints so it does not “helpfully” invent a generic Cosmos chain. The prompts below are written so you can use them with a coding agent, one phase at a time.

## How to use them

Run these prompts in order, and require the agent to produce files, diffs, and decision logs at each step. Keep one rule constant across all prompts: the agent must not change Tabbre tokenomics or validator policy unless explicitly told to do so.

- Start with architecture and repo scaffolding.
- Then move to chain configuration and module customization.
- After that, generate infra, tests, docs, and launch artifacts.
- Require every step to end with: assumptions, open questions, risks, and next actions.


## Master system prompt

Use this as the persistent instruction for the agent across the whole project:

```text
You are the lead protocol engineer for the Tabbre Core Chain MVP.

Your job is to create a production-oriented MVP blockchain based on Cosmos SDK + Cosmos EVM, using Cosmos EVM as the starting point. You must act like a senior blockchain engineer, DevOps engineer, and technical writer working together.

Non-negotiable product constraints:
1. The chain is called Tabbre Core Chain (or TabbreChain in code/docs where appropriate).
2. The MVP is based on Cosmos SDK and Cosmos EVM.
3. TABB is the fixed-supply governance, staking, fundraising, and collateral token.
4. Total TABB supply is 1,000,000,000 created in a single token creation event.
5. TABB is not inflationary. Do not add ongoing TABB emissions unless explicitly instructed.
6. BRE is the native transactional token for fees and usage charges.
7. Transaction fees and usage fees are payable in BRE.
8. Validators stake TABB.
9. Validator rewards are paid in BRE-based fee income, not TABB inflation.
10. The chain launches first as a permissioned consortium chain, then can evolve later toward broader PoS participation.
11. Governance is token-based: one TABB, one vote, subject to implementation design.
12. The chain must support smart contracts through Cosmos EVM and standard Ethereum-style JSON-RPC compatibility where practical for the MVP.
13. Slashed TABB should be burned or sent to a burn mechanism compatible with Tabbre tokenomics.
14. Do not redesign the Tabbre business model, tokenomics, or mission.
15. Where product requirements are unspecified, choose the simplest MVP-compatible implementation and clearly document assumptions.

Working rules:
- Prefer minimal, working, testable implementations over ambitious architecture.
- Reuse Cosmos EVM patterns and existing module structures where possible.
- Avoid introducing unnecessary custom modules if configuration, hooks, wrappers, or minimal additions are enough.
- Produce code, config, docs, scripts, tests, and runbooks.
- At every stage, explain what you changed, why, what remains unresolved, and how to verify it.
- If a requirement conflicts with Cosmos EVM defaults, explicitly identify the conflict and propose the least risky implementation path.
- Never silently invent economics that contradict the Tabbre model.
- Favor deterministic local development, Docker-based workflows, and reproducible builds.

Output requirements for every task:
- Summary
- Files created/changed
- Key design decisions
- Step-by-step implementation notes
- How to run/test
- Risks/issues
- Next recommended prompt
```


## Prompt 1: solution architecture

```text
Using the project constraints below, create the technical architecture for the Tabbre Core Chain MVP built on Cosmos EVM.

Project constraints:
- Cosmos SDK + Cosmos EVM base
- TABB fixed supply: 1,000,000,000 minted once at genesis
- BRE is the fee token
- TABB is the staking token
- Permissioned consortium validator set at launch
- EVM smart contract support required
- JSON-RPC support required
- Minimal viable governance and upgrade path required
- Slashed TABB should be burned or routed to a burn-compatible sink
- Avoid unnecessary custom modules

Tasks:
1. Propose the target architecture and repo structure.
2. Identify which Cosmos SDK and Cosmos EVM modules should be used unchanged.
3. Identify which modules need configuration changes.
4. Identify where custom code is unavoidable for:
   - BRE fee payments
   - TABB staking
   - Genesis mint/allocation
   - Permissioned validator admission
   - Slashing burn behavior
5. Recommend the simplest implementation path for MVP.
6. Produce:
   - architecture.md
   - module-map.md
   - implementation-plan.md
   - repo tree proposal

Be specific about app wiring, token handling, fee flow, validator flow, governance flow, and upgrade flow.

End with:
- assumptions
- unresolved issues
- recommended next step
```


## Prompt 2: gap analysis against Cosmos EVM

```text
Perform a gap analysis between standard Cosmos EVM starter architecture and the requirements of Tabbre Core Chain MVP.

Requirements to test against:
- TABB fixed one-time supply at genesis
- BRE as gas/fee token
- TABB as staking token
- BRE fee rewards to validators
- Permissioned consortium validator set at launch
- EVM compatibility
- On-chain governance for upgrades/parameters
- Burn treatment for slashed TABB
- Public RPC for users and developers

Deliver:
1. A table with columns:
   - requirement
   - supported out of the box?
   - partial support?
   - custom work needed?
   - risk level
   - recommended implementation
2. A ranked list of the hardest engineering issues.
3. A recommendation for what should be implemented now vs deferred.

Do not stay abstract. Reference likely Cosmos SDK / Cosmos EVM integration points, app config, ante handlers, staking params, fee logic, and module account design.

End with:
- top 5 technical risks
- top 5 simplifications for MVP
- exact next build task
```


## Prompt 3: scaffold the chain

```text
Create the initial implementation plan for scaffolding Tabbre Core Chain from a Cosmos EVM starter.

Objectives:
- define chain binary name
- define chain-id conventions for local, testnet, and mainnet
- define denom strategy for TABB and BRE
- define module account strategy
- define genesis allocation approach
- define validator key and account setup approach
- define local dev workflow

Please produce:
1. naming-and-denoms.md
2. genesis-strategy.md
3. environment-strategy.md
4. a step-by-step scaffold plan from starter repo to first running localnet
5. shell commands and/or Makefile targets to initialize local development

Assume:
- TABB and BRE both exist natively on-chain for MVP
- TABB supply is minted once at genesis
- BRE initial balances may be seeded for testing
- consortium validators are known at genesis

The output should be practical enough that an engineer can start coding immediately.
```


## Prompt 4: implement token model

```text
Design and implement the token model for Tabbre Core Chain MVP.

Rules:
- TABB total supply = 1,000,000,000 at genesis
- No inflationary TABB emissions
- TABB is staking token
- BRE is fee token
- Validator rewards come from BRE fee income
- TABB slashing should reduce effective supply via burn-compatible handling
- Keep implementation as simple as possible

Tasks:
1. Define denom metadata, precision, and display units for TABB and BRE.
2. Show how genesis minting and allocation should work.
3. Show how TABB is made the bonded staking token.
4. Show how BRE becomes the fee token for transactions.
5. Propose the exact minimal code/config changes required.
6. Produce code patches or pseudocode for:
   - genesis setup
   - staking token config
   - fee handling
   - validator reward routing
   - slashing burn path
7. Write tokenomics-implementation.md explaining how implementation matches Tabbre requirements.

If there are multiple ways to implement BRE-as-fee-token on Cosmos EVM, compare them and recommend one for MVP.
```


## Prompt 5: fee handling and ante logic

```text
Design the transaction fee and ante-handler model for Tabbre Core Chain MVP.

Tabbre requirements:
- fees payable in BRE
- EVM and non-EVM transactions should both be considered
- TABB is not used for normal transaction gas
- validators are rewarded from BRE fee income
- solution should minimize invasive changes and preserve Cosmos EVM compatibility as much as possible

Tasks:
1. Explain how fees are currently handled in a standard Cosmos EVM chain.
2. Design the minimal modifications required so BRE is accepted as the fee token.
3. Address:
   - mempool admission
   - min gas price handling
   - EVM tx fee accounting
   - non-EVM tx fee accounting
   - fee collector distribution
   - compatibility concerns with wallets and tooling
4. Generate:
   - fee-model.md
   - ante-handler-plan.md
   - compatibility-notes.md
   - recommended code changes

Highlight any place where a short-term compromise is needed for MVP.
```


## Prompt 6: validator policy implementation

```text
Implement the validator policy for a permissioned consortium launch of Tabbre Core Chain MVP.

Requirements:
- only approved consortium members can validate at launch
- validators must stake TABB
- slashing applies for misconduct
- the design should later allow controlled transition toward broader participation
- validator policy should be simple and auditable

Tasks:
1. Propose the best MVP mechanism for permissioned validator admission.
2. Compare:
   - hardcoded genesis validator set only
   - allowlist-based admission
   - governance-controlled validator onboarding
3. Recommend one design for launch and one for phase 2.
4. Produce:
   - validator-policy.md
   - validator-lifecycle.md
   - genesis-validator-onboarding.md
   - governance-transition-plan.md
5. Include operational rules:
   - minimum self-stake parameter
   - key rotation approach
   - downtime expectations
   - slashing events
   - removal/suspension flow

Do not invent business rules without labeling them as proposed MVP defaults.
```


## Prompt 7: governance and upgrades

```text
Design the governance and upgrade framework for Tabbre Core Chain MVP.

Requirements:
- TABB is the governance asset
- one TABB one vote is the target principle
- minimal working governance is enough for MVP
- chain upgrades must be manageable safely
- permissioned consortium launch should not block future decentralization

Tasks:
1. Propose the MVP governance design using existing Cosmos SDK capabilities where possible.
2. Explain how close this gets to one-TABB-one-vote and note any deviations.
3. Define what proposal types should be enabled at MVP launch.
4. Define what should stay disabled until later phases.
5. Produce:
   - governance-design.md
   - upgrade-policy.md
   - gov-params-proposal.md
   - open-questions.md

Include a recommendation on whether to use standard staking-weighted governance initially or introduce custom tokenocracy logic later.
```


## Prompt 8: local devnet build

```text
Build a local multi-node development network plan for Tabbre Core Chain MVP.

Requirements:
- at least 3 validators
- permissioned consortium-style setup
- TABB staking token
- BRE fee token
- EVM RPC enabled
- faucet or seeded accounts for testing
- reproducible developer workflow

Deliver:
1. docker-compose or similar localnet design
2. scripts for:
   - init
   - genesis generation
   - account seeding
   - validator setup
   - node startup
   - chain reset
3. localnet.md with step-by-step instructions
4. test accounts and sample balances
5. sample RPC endpoints and chain config

Design this for deterministic local testing by engineers and the AI agent itself.
```


## Prompt 9: end-to-end testing

```text
Create the test strategy and initial test suite for Tabbre Core Chain MVP.

Core flows to test:
- genesis starts correctly
- TABB total supply is correct
- BRE fee payment works
- TABB staking works
- validator set behaves correctly
- permissioned validator rules are enforced
- slashing path works
- EVM contract deployment works
- JSON-RPC works with standard Ethereum tools
- governance proposal flow works
- upgrade rehearsal works

Deliver:
1. test-plan.md
2. a categorized test matrix
3. integration test list
4. smoke tests for localnet
5. examples using CLI and EVM tooling
6. CI strategy for running tests

Prioritize tests that prove Tabbre-specific economics and validator rules are implemented correctly.
```


## Prompt 10: security review preparation

```text
Prepare a security review package for the Tabbre Core Chain MVP implementation.

Focus areas:
- fee token customization
- staking token / fee token split
- permissioned validator logic
- slashing burn handling
- genesis allocation correctness
- module account permissions
- governance and upgrade safety
- RPC exposure and operational risk

Deliver:
1. threat-model.md
2. audit-readiness-checklist.md
3. privileged-accounts-and-module-permissions.md
4. key invariants.md
5. known-risks.md
6. recommended pre-mainnet fixes

Be explicit about which customizations are highest risk relative to a standard Cosmos EVM chain.
```


## Prompt 11: docs for developers and validators

```text
Create the first full documentation set for Tabbre Core Chain MVP.

Audience:
- protocol engineers
- validator operators
- wallet/dev-tool integrators
- internal Tabbre stakeholders

Create:
1. README.md
2. docs/architecture.md
3. docs/run-localnet.md
4. docs/token-model.md
5. docs/validator-operations.md
6. docs/json-rpc.md
7. docs/governance.md
8. docs/known-limitations.md

Requirements:
- document TABB vs BRE responsibilities clearly
- explain the permissioned validator launch model
- explain how fees work
- explain how developers deploy Solidity contracts
- explain what is MVP-only and what is deferred
```


## Prompt 12: testnet release prep

```text
Prepare Tabbre Core Chain MVP for public or semi-public testnet release.

Requirements:
- stable chain config
- known validator cohort
- RPC endpoints
- explorer integration plan
- faucet approach
- release artifacts
- rollback and recovery plan

Deliver:
1. testnet-launch-checklist.md
2. release-runbook.md
3. validator-invite-pack.md
4. incident-response-runbook.md
5. config-audit.md
6. mainnet-readiness-gaps.md

Assume the chain is still consortium-led at this stage.
```


## Prompt 13: coding execution prompt

Use this when you actually want the agent to start writing code:

```text
Start implementing Tabbre Core Chain MVP now.

Rules:
- Work in small, reviewable steps.
- Do not refactor unrelated code.
- Before each code change, state the objective and affected files.
- After each code change, provide:
  - diff summary
  - why it was needed
  - how to test it
  - whether it changes chain state, consensus behavior, or RPC behavior
- Maintain a running CHANGELOG_FOR_AGENT.md
- Maintain a running OPEN_ISSUES_FOR_AGENT.md
- Maintain a running ASSUMPTIONS_FOR_AGENT.md

Implementation priorities:
1. chain naming and app scaffold
2. genesis token setup for TABB and BRE
3. staking token = TABB
4. fee token = BRE
5. consortium validator admission
6. localnet scripts
7. smoke tests
8. docs

Important:
- If BRE fee-token implementation becomes too invasive, stop and propose the least-risk fallback with exact trade-offs instead of improvising.
- If one-TABB-one-vote is not cleanly achievable in MVP using standard governance, document the deviation and implement the safest minimal governance path.

Start with step 1 only and wait for approval before moving to step 2.
```


## Prompt 14: architecture decision record prompt

This is useful to keep the agent disciplined:

```text
For the current implementation stage, write Architecture Decision Records (ADRs) for every material decision.

Include ADRs for:
- choosing Cosmos EVM as base
- TABB as staking token
- BRE as fee token
- fixed-supply TABB genesis mint
- permissioned consortium validator launch
- governance model for MVP
- slashing burn treatment
- JSON-RPC exposure model

Each ADR must contain:
- title
- status
- context
- decision
- alternatives considered
- consequences
- follow-up actions
```


## Prompt 15: final audit-and-handover prompt

```text
Prepare the Tabbre Core Chain MVP for handover.

Deliver:
1. complete file tree
2. setup instructions from zero
3. localnet verification steps
4. test coverage summary
5. list of all deviations from ideal Tabbre end-state design
6. list of all temporary MVP compromises
7. list of all security-sensitive customizations
8. recommended next 30/60/90 day roadmap
9. operator handover notes
10. executive technical summary

Also provide:
- what is production-ready
- what is testnet-ready only
- what must be audited before mainnet
```


## Optional guardrail prompt

Use this whenever the agent starts drifting into generic chain design:

```text
Re-check your output against Tabbre constraints.

Specifically verify:
- no inflationary TABB design was introduced
- BRE remains the fee token
- TABB remains the staking and governance token
- validator rewards are not described as TABB emissions unless explicitly justified as a temporary workaround
- permissioned consortium launch remains intact
- token allocations do not contradict the Tabbre tokenomics
- no major custom module was introduced without necessity

List every place where your current output may conflict with these constraints.
Then revise the output accordingly.
```


## Recommended execution order

1. Master system prompt
2. Prompt 1
3. Prompt 2
4. Prompt 3
5. Prompt 4
6. Prompt 5
7. Prompt 6
8. Prompt 7
9. Prompt 8
10. Prompt 9
11. Prompt 10
12. Prompt 11
13. Prompt 12
14. Prompt 13
15. Prompt 14 and 15 near the end

One important note: the hardest part is likely the **split-token design** where TABB is the staking asset but BRE is the fee token, because many standard chain flows assume a tighter relationship between staking rewards, fee collection, and native gas denomination. So your prompts should force the agent to pause and document trade-offs instead of hacking around them.

I can also turn this into a **copy-paste agent workbook** with:

- one Claude Code version,
- one Cursor version,
- one OpenAI coding-agent version,
- and one “single prompt that manages the whole build” version.

