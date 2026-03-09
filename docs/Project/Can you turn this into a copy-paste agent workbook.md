<img src="https://r2cdn.perplexity.ai/pplx-full-logo-primary-dark%402x.png" style="height:64px;margin-right:32px"/>

# Can you turn this into a copy-paste agent workbook  for Google's Antigravity?

Perfect — here's a **copy-paste Antigravity workbook** for building Tabbre Core Chain MVP. Antigravity uses a distinct workflow: it prefers Planning mode with task lists, implementation plans, and artifact generation, plus it can orchestrate multiple agents simultaneously through the Agent Manager.[^1][^2]

***

# Tabbre Core Chain MVP — Antigravity Agent Workbook

## How to use this workbook

1. **Install Antigravity** from https://antigravity.google/download[^3]
2. **Create a workspace folder** for Tabbre Core Chain (e.g., `~/tabbre-core-chain-mvp`)
3. **Open Agent Manager** in Antigravity
4. **Set to Planning mode** (not Fast mode) for complex architectural work[^1]
5. **Use Gemini 3 Pro** as your model (the default)[^4]
6. **Copy prompts sequentially** — each prompt builds on the previous
7. **Review artifacts** (implementation plans, task lists, walkthroughs) before approving[^2][^1]
8. **Leave Google Docs-style comments** on artifacts to steer the agent[^1]
9. **Use multi-agent orchestration** when appropriate (e.g., parallel workstreams for docs + infra + core)[^1]

***

## Global Rules Setup (one-time configuration)

Before starting, add these **global rules** to ensure the agent respects Tabbre constraints across all conversations.

**How to add:**

1. In Antigravity, go to `...` menu → `Customizations` → `Rules` → `+ Global`[^1]
2. Create a rule file named `tabbre-tokenomics-constraints.md`
3. Paste the following:
```markdown
# Tabbre Core Chain Non-Negotiable Constraints

When working on Tabbre Core Chain, you MUST respect these product requirements:

## Chain Identity
- Chain name: Tabbre Core Chain (or TabbreChain in code)
- Based on: Cosmos SDK + Cosmos EVM
- Chain ID conventions: `tabbre-local-1` (devnet), `tabbre-testnet-1` (testnet), `tabbre-mainnet-1` (mainnet)

## Token Model
- TABB: fixed supply of 1,000,000,000 tokens created once at genesis
- TABB is NOT inflationary — do not add ongoing TABB emissions
- TABB is used for: staking, governance, collateral
- BRE: native transactional token for fees and usage charges
- Transaction fees are paid in BRE, not TABB
- Validator rewards are paid in BRE from fee income, not from TABB inflation

## Validator Model
- Launch mode: permissioned consortium validator set
- Validators stake TABB
- Validators earn BRE-based fee rewards
- Slashed TABB should be burned or sent to burn-compatible mechanism
- Future evolution: controlled transition toward broader PoS participation

## Governance
- Governance token: TABB
- Principle: one TABB, one vote (implement via standard Cosmos SDK governance where practical)
- Governance controls: protocol parameters, upgrades, validator admission (future)

## Technical Requirements
- EVM smart contract support via Cosmos EVM
- JSON-RPC compatibility for Ethereum tooling
- Standard Cosmos SDK modules where possible
- Minimize custom modules — prefer configuration and wrappers

## What NOT to do
- Do not redesign Tabbre tokenomics
- Do not invent TABB inflation schedules
- Do not make TABB the gas token
- Do not create complex custom modules if simpler solutions exist
- Do not silently change allocation percentages or supply caps

## When uncertain
- Choose the simplest MVP-compatible implementation
- Document all assumptions explicitly
- Propose alternatives with clear trade-offs
- Ask for clarification rather than guessing business rules
```


***

## Workspace Setup

**Step 1: Create workspace**

In Antigravity Agent Manager:

1. Click `Start Conversation`
2. Select or create workspace: `tabbre-core-chain-mvp`
3. Ensure Planning mode is selected
4. Model: Gemini 3 Pro

***

## Phase 1: Architecture and Gap Analysis

### Prompt 1: Solution Architecture

**What this does:** Creates technical architecture, module map, repo structure, and implementation plan. The agent will generate artifacts including architecture.md, module-map.md, and implementation-plan.md.[^2][^1]

**Copy and paste this into Agent Manager:**

```
I need you to design the technical architecture for Tabbre Core Chain MVP, a production-oriented blockchain based on Cosmos SDK + Cosmos EVM.

Context and constraints:
- TABB is a fixed-supply token (1,000,000,000 total) minted once at genesis
- TABB is used for staking and governance
- BRE is the fee token (transaction fees paid in BRE, not TABB)
- Validators stake TABB but earn BRE-based fee rewards
- Launch as permissioned consortium validator set, evolve later to broader PoS
- Must support EVM smart contracts via Cosmos EVM
- Must support JSON-RPC for Ethereum tooling compatibility
- Slashed TABB should be burned or routed to burn mechanism
- Keep implementation minimal and reuse Cosmos modules where possible

Your tasks:
1. Propose the target architecture and complete repo structure
2. Identify which Cosmos SDK and Cosmos EVM modules can be used unchanged
3. Identify which modules need configuration changes only
4. Identify where custom code is unavoidable:
   - BRE as fee token instead of staking token
   - TABB as staking token (not the fee token)
   - Fixed supply genesis mint for TABB
   - Permissioned validator admission
   - Slashing burn behavior
5. Recommend the simplest implementation path for MVP
6. Create these artifacts:
   - architecture.md (comprehensive technical architecture)
   - module-map.md (which Cosmos modules, what config needed)
   - implementation-plan.md (step-by-step build approach)
   - repo-structure.txt (proposed directory tree)

Be specific about:
- app wiring and initialization
- token denomination strategy for TABB and BRE
- fee collection and distribution flow
- validator lifecycle and admission logic
- governance and upgrade flow
- EVM integration points

At the end of your implementation plan, list:
- All assumptions you made
- Unresolved technical questions
- Recommended next step
```

**What to expect:**

- Agent creates implementation plan artifact (review and comment if needed)
- Agent creates task list artifact
- Agent generates architecture.md, module-map.md, implementation-plan.md files
- Agent creates walkthrough showing what was produced

**Your action:** Review the implementation plan artifact. If the agent proposes TABB inflation or makes TABB the gas token, leave a comment: "This conflicts with Tabbre tokenomics. TABB supply is fixed at genesis. BRE is the fee token. Please revise."

***

### Prompt 2: Gap Analysis

**What this does:** Identifies exactly what differs between standard Cosmos EVM and Tabbre requirements, with risk assessment and recommendations.

**Copy and paste:**

```
Perform a detailed gap analysis between standard Cosmos EVM starter architecture and Tabbre Core Chain MVP requirements.

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
```

**Your action:** Review gap-analysis.md carefully. Pay special attention to how the agent proposes handling the TABB (staking) vs BRE (fee) split, since this is the hardest part. If the solution looks too invasive or risky, leave a comment asking for alternatives.

***

## Phase 2: Foundation and Configuration

### Prompt 3: Naming, Denoms, and Genesis Strategy

**What this does:** Locks down all naming conventions, token denominations, genesis approach, and local dev workflow.

**Copy and paste:**

```
Design the naming conventions, token denominations, and genesis strategy for Tabbre Core Chain MVP.

Objectives:
1. Define chain binary name (e.g., `tabbred`)
2. Define chain-id conventions:
   - Local devnet: `tabbre-local-1`
   - Public testnet: `tabbre-testnet-1`
   - Mainnet: `tabbre-mainnet-1`
3. Define denom strategy:
   - TABB denom (suggest `utabb` for micro-TABB, 1 TABB = 1,000,000 utabb)
   - BRE denom (suggest `ubre` for micro-BRE)
   - Display names and precision
4. Define genesis allocation approach:
   - Total TABB supply: 1,000,000,000
   - Allocations: 20% private sale, 10% public sale, 10% founders, 10% community, 50% reserve
   - How to represent these allocations in genesis accounts
   - How to handle reserve vesting (0.5% per month for 120 months)
5. Define module account strategy (fee collector, staking pool, governance, etc.)
6. Define validator setup approach for genesis validators
7. Define local dev workflow (init, keygen, genesis setup, node start)

Create these artifacts:
- naming-and-denoms.md
- genesis-strategy.md
- environment-strategy.md
- local-dev-quickstart.md (with shell commands to get started)

Assume:
- TABB and BRE both exist as native on-chain assets
- TABB total supply minted at genesis
- BRE initial balances seeded for testing purposes
- Consortium validators known at genesis
```

**Your action:** Review naming-and-denoms.md to ensure TABB and BRE are clearly distinguished. Confirm genesis-strategy.md allocates exactly 1B TABB with the correct percentages.

***

### Prompt 4: Token Model Implementation

**What this does:** Designs exact implementation for TABB and BRE tokenomics within Cosmos SDK + EVM.

**Copy and paste:**

```
Design and implement the token model for Tabbre Core Chain MVP.

Token requirements:
- TABB: 1,000,000,000 total supply minted at genesis (no inflation)
- TABB: used for staking (bonded to validators and delegators)
- TABB: used for governance voting weight
- BRE: used for transaction fees (gas payments)
- BRE: used for validator reward distribution from fee income
- Slashed TABB: burned or sent to permanent burn address

Tasks:
1. Define token metadata:
   - TABB: base denom, display denom, precision (suggest 6 decimals)
   - BRE: base denom, display denom, precision (suggest 6 decimals)
2. Show exact genesis setup:
   - How to mint 1B TABB at genesis
   - How to allocate to genesis accounts (private sale, founders, reserve, etc.)
   - How to seed initial BRE for testing
3. Show how to configure TABB as the bonded staking token in x/staking module
4. Show how to configure BRE as the fee token for transactions
5. Show how to route validator rewards in BRE from fee collector to validators
6. Show slashing burn path for TABB
7. Provide exact code patches, config changes, or pseudocode for:
   - Genesis setup
   - Staking token config
   - Fee token config
   - Reward distribution routing
   - Slashing burn implementation

Create tokenomics-implementation.md explaining:
- How each Tabbre requirement maps to Cosmos SDK implementation
- What standard modules are used
- What custom logic (if any) is needed
- Trade-offs for MVP vs future phases

If there are multiple approaches to make BRE the fee token on Cosmos EVM, compare them and recommend one for MVP with justification.
```

**Your action:** This is critical. Review tokenomics-implementation.md carefully. Confirm:

- No TABB inflation mechanism was added
- BRE is clearly the fee token
- Validator reward flow is BRE-based, not TABB-based
- If the agent proposes complex custom modules, ask: "Is there a simpler config-based approach?"

***

## Phase 3: Fee Handling and Validator Policy

### Prompt 5: Fee and Ante Handler Design

**What this does:** Solves the technical challenge of making BRE the fee token while keeping TABB as staking token.

**Copy and paste:**

```
Design the transaction fee and ante-handler model for Tabbre Core Chain MVP.

Tabbre requirements:
- All transaction fees paid in BRE (not TABB)
- Applies to both EVM transactions and native Cosmos transactions
- Validators are rewarded from BRE fee income (collected fees distributed to validators)
- Solution must preserve Cosmos EVM compatibility as much as possible
- Minimize invasive custom changes

Tasks:
1. Explain how fees are currently handled in standard Cosmos EVM chain
2. Design the minimal modifications to accept BRE as the fee token
3. Address these flows:
   - Mempool admission (checking min gas price in BRE)
   - EVM transaction fee accounting (gas price * gas used in BRE)
   - Non-EVM transaction fee accounting (standard SDK tx fees in BRE)
   - Fee collector module account setup
   - Fee distribution to validators (from fee collector to validator rewards)
4. Address compatibility:
   - Wallet integration (MetaMask, Keplr)
   - JSON-RPC fee estimation
   - Impact on existing Cosmos EVM tooling
5. Generate artifacts:
   - fee-model.md (comprehensive fee design)
   - ante-handler-plan.md (what ante handlers need to change)
   - compatibility-notes.md (wallet and tooling impacts)
   - code-changes-summary.md (what files need editing)

Highlight any place where MVP requires a short-term compromise that should be revisited post-launch.
```

**Your action:** Review fee-model.md and ante-handler-plan.md. If the proposed solution requires deep custom ante handler logic or breaks EVM compatibility, comment: "This seems invasive. Is there a way to configure fee denomination without custom ante logic? What are the trade-offs?"

***

### Prompt 6: Validator Policy Implementation

**What this does:** Implements permissioned consortium launch with path to future decentralization.

**Copy and paste:**

```
Design and implement the validator policy for permissioned consortium launch of Tabbre Core Chain MVP.

Requirements:
- Only approved consortium members can validate at launch
- Validators must stake TABB (minimum amount TBD)
- Validators earn BRE-based rewards from transaction fees
- Slashing applies for downtime and double-sign misbehavior
- Design should allow controlled future transition to open validator participation
- Policy must be simple, auditable, and operationally clear

Tasks:
1. Propose the MVP mechanism for permissioned validator admission. Compare:
   - Hardcoded genesis validator set only (no new validators post-genesis)
   - Allowlist-based admission (new validators require allowlist entry)
   - Governance-controlled validator onboarding (x/gov proposals to add validators)
2. Recommend one approach for MVP launch and one for phase 2 evolution
3. Create artifacts:
   - validator-policy.md (policy overview and rationale)
   - validator-lifecycle.md (how validators join, operate, leave)
   - genesis-validator-onboarding.md (step-by-step for consortium members)
   - governance-transition-plan.md (how to move toward open PoS later)
4. Define operational parameters (propose MVP defaults, mark as governance-adjustable):
   - Minimum self-stake in TABB
   - Maximum validator count at launch
   - Commission rate limits
   - Uptime requirements
   - Slashing conditions and penalties
   - Key rotation approach
   - Validator removal/suspension process

Important: Do not invent business-specific rules (like exact TABB amounts) without labeling them as "proposed MVP defaults subject to governance approval."
```

**Your action:** Review validator-policy.md. Ensure it clearly states that this is a permissioned consortium at launch (not open PoS). Check that slashing penalties result in TABB being burned, not just redistributed.

***

## Phase 4: Governance and Local Development

### Prompt 7: Governance and Upgrade Framework

**Copy and paste:**

```
Design the governance and upgrade framework for Tabbre Core Chain MVP.

Requirements:
- TABB is the governance asset (voting weight based on TABB holdings)
- Target principle: one TABB = one vote
- Minimal working governance sufficient for MVP
- Must support safe chain upgrades
- Permissioned consortium launch should not block future decentralization

Tasks:
1. Propose MVP governance design using Cosmos SDK x/gov module where possible
2. Explain how close this gets to true one-TABB-one-vote (note any deviations, e.g., staking-weighted vs token-weighted)
3. Define proposal types enabled at MVP launch:
   - Parameter change proposals
   - Software upgrade proposals
   - Text proposals
4. Define what should be disabled or deferred until later phases
5. Define governance parameters for MVP:
   - Minimum deposit (in TABB)
   - Voting period
   - Quorum threshold
   - Pass threshold
   - Veto threshold
6. Create artifacts:
   - governance-design.md
   - upgrade-policy.md
   - gov-params-proposal.md (proposed genesis governance params)
   - open-questions.md (unresolved governance questions)

Recommendation: Should we use standard staking-weighted governance (votes weighted by staked TABB) initially, or introduce custom tokenocracy logic (votes weighted by all TABB holdings) later? Analyze trade-offs.
```

**Your action:** Review governance-design.md. Confirm TABB is the governance token. If the agent deviates from one-TABB-one-vote due to technical constraints, ensure it's clearly documented with a plan to address later.

***

### Prompt 8: Local Devnet Setup

**Copy and paste:**

```
Build a reproducible local multi-node development network for Tabbre Core Chain MVP.

Requirements:
- At least 3 validator nodes
- Permissioned consortium-style setup
- TABB as staking token, BRE as fee token
- EVM and JSON-RPC enabled
- Seeded test accounts with TABB and BRE balances
- Faucet script or pre-funded accounts for testing
- Fully reproducible developer workflow

Deliverables:
1. Local network architecture (docker-compose or shell-based multi-node setup)
2. Scripts for:
   - init: Initialize chain config and home directories
   - genesis: Generate genesis file with correct TABB/BRE allocations and validators
   - accounts: Create and fund test accounts
   - validators: Set up validator keys and gentx
   - start: Launch all nodes
   - stop: Stop all nodes
   - reset: Clean state and restart from genesis
3. Artifacts:
   - localnet.md (comprehensive local devnet guide)
   - docker-compose.yml or equivalent multi-node setup
   - scripts/ directory with all automation
   - test-accounts.json (list of test accounts with keys and balances)
   - Sample RPC endpoints and explorer integration instructions
4. Example commands for:
   - Querying TABB and BRE balances
   - Sending transactions
   - Staking TABB
   - Submitting governance proposals
   - Deploying EVM contracts via JSON-RPC

Design this for deterministic local testing by engineers and CI systems.
```

**Your action:** Once localnet.md is created, actually try running the local devnet following the instructions. Verify that:

- TABB and BRE denoms are correct
- You can stake TABB
- You can pay fees in BRE
- JSON-RPC works with MetaMask or Hardhat

***

## Phase 5: Testing, Security, and Documentation

### Prompt 9: Test Strategy and Suite

**Copy and paste:**

```
Create comprehensive test strategy and initial test suite for Tabbre Core Chain MVP.

Core flows to test:
1. Genesis initialization: verify TABB total supply = 1B, allocations correct
2. TABB staking: delegate, undelegate, redelegate flows work
3. BRE fee payment: transactions deduct BRE as gas fees
4. Validator rewards: validators receive BRE from fee collector
5. Slashing: misbehavior slashes TABB, slashed amount is burned
6. Permissioned validators: only consortium members can validate
7. EVM compatibility: deploy Solidity contracts, call functions, verify events
8. JSON-RPC: standard Ethereum JSON-RPC methods work (eth_blockNumber, eth_sendTransaction, etc.)
9. Governance: submit proposals, vote with TABB, execute param changes
10. Upgrades: perform chain software upgrade without state corruption

Deliverables:
1. test-plan.md (comprehensive test strategy)
2. Test matrix with categories:
   - Unit tests (per module)
   - Integration tests (multi-module flows)
   - End-to-end tests (full user workflows)
   - Smoke tests (quick sanity checks for localnet)
3. Example test implementations:
   - CLI-based tests (using chain binary)
   - EVM tooling tests (using Hardhat/Foundry)
   - JSON-RPC tests (using web3.js or ethers.js)
4. CI strategy: how to run these tests in continuous integration
5. Prioritized test list: which tests prove Tabbre-specific tokenomics and validator rules are correct

Generate test-plan.md and sample test scripts in tests/ directory.
```

**Your action:** Review test-plan.md and prioritize implementing the TABB/BRE tokenomics tests first (supply, fees, staking, rewards). These are the highest risk areas.

***

### Prompt 10: Security Review Preparation

**Copy and paste:**

```
Prepare security review package for Tabbre Core Chain MVP implementation.

Security focus areas:
1. Fee token customization (BRE as fee token instead of staking token)
2. Staking token / fee token split (TABB staked, BRE paid as fees)
3. Permissioned validator logic (admission control, removal)
4. Slashing and burn handling (slashed TABB burned correctly)
5. Genesis allocation correctness (1B TABB allocated correctly, no minting exploits)
6. Module account permissions (who can mint, burn, transfer from module accounts)
7. Governance and upgrade safety (proposal execution, upgrade coordination)
8. RPC exposure and operational security (rate limits, abuse vectors)

Deliverables:
1. threat-model.md (identify attack vectors specific to Tabbre design)
2. audit-readiness-checklist.md (what must be reviewed before mainnet)
3. privileged-accounts-and-permissions.md (list all privileged accounts and capabilities)
4. invariants.md (critical invariants that must always hold, e.g., "total TABB supply never exceeds 1B")
5. known-risks.md (acknowledged risks and mitigations)
6. recommended-fixes.md (pre-audit hardening recommendations)

Be explicit about which customizations are highest risk compared to standard Cosmos EVM chain.
```

**Your action:** Review threat-model.md and invariants.md. Share these with your security team or auditor early. Prioritize fixes for any High risk items in known-risks.md before testnet launch.

***

### Prompt 11: Developer and Validator Documentation

**Copy and paste:**

```
Create complete documentation set for Tabbre Core Chain MVP.

Audience:
- Protocol engineers (internal team)
- Validator operators (consortium members)
- Wallet and dev-tool integrators (MetaMask, Hardhat users)
- Internal stakeholders (Tabbre Foundation, governance)

Create these docs:
1. README.md (repo root, quick overview and links)
2. docs/architecture.md (technical architecture overview)
3. docs/tokenomics.md (TABB and BRE roles, allocations, economics)
4. docs/run-localnet.md (how to run local devnet)
5. docs/validator-operations.md (how to run a validator node)
6. docs/json-rpc.md (JSON-RPC endpoints and Ethereum compatibility)
7. docs/governance.md (how governance works, how to submit proposals)
8. docs/evm-contracts.md (how to deploy Solidity contracts)
9. docs/known-limitations.md (what is MVP-only, what is deferred)
10. docs/troubleshooting.md (common issues and fixes)

Requirements:
- Clearly explain TABB (staking/governance) vs BRE (fees) distinction throughout
- Explain permissioned consortium launch model and future evolution
- Provide concrete examples: CLI commands, JSON-RPC calls, Solidity deployment
- Call out MVP limitations and future roadmap items
- Keep language clear and accessible (avoid unnecessary jargon)

Generate all docs in docs/ directory and update README.md with navigation links.
```

**Your action:** Review README.md and docs/tokenomics.md carefully. These will be the first things external users see. Ensure tokenomics are explained clearly and accurately.

***

## Phase 6: Testnet and Launch Prep

### Prompt 12: Testnet Launch Preparation

**Copy and paste:**

```
Prepare Tabbre Core Chain MVP for public testnet release.

Requirements:
- Stable chain configuration
- Consortium validator cohort identified and onboarded
- Public RPC endpoints available
- Block explorer integration plan
- Testnet faucet for BRE (test tokens)
- Release artifacts (binaries, genesis file, docs)
- Rollback and incident response plan

Deliverables:
1. testnet-launch-checklist.md (step-by-step launch checklist)
2. release-runbook.md (how to build, test, and release testnet version)
3. validator-onboarding-pack.md (send this to consortium validators with setup instructions)
4. incident-response-runbook.md (what to do if testnet halts, security issue, etc.)
5. config-audit.md (audit all chain params, ensure nothing is misconfigured)
6. mainnet-readiness-gaps.md (what still needs to be done before mainnet)

Include:
- Pre-launch testing checklist (run all tests, genesis dry run, upgrade rehearsal)
- Monitoring setup (block production, validator uptime, RPC health)
- Communication plan (how to announce testnet, where to publish docs and RPC)
- Validator SLA expectations (uptime, response time)
- Faucet setup (how users get test BRE tokens)

Assume testnet is still consortium-led permissioned chain (not open to public validators yet).
```

**Your action:** Use testnet-launch-checklist.md as your actual launch checklist. Don't skip steps. Run config-audit.md to verify all genesis parameters are correct before testnet genesis.

***

## Parallel Workstream: Multi-Agent Orchestration

**Antigravity's killer feature:** You can run multiple agents in parallel via Agent Manager.[^2][^1]

Once you've completed Phase 1-2 (architecture and foundation), you can split work across agents:

### Agent 1: Core Chain Implementation

**Task:** Implement token model, fee handling, validator policy, governance

### Agent 2: Infrastructure and DevOps

**Task:** Build docker-compose localnet, CI/CD scripts, monitoring setup

### Agent 3: Documentation and Testing

**Task:** Write all docs, create test suite, prepare audit materials

**How to do this in Antigravity:**

1. Open Agent Manager
2. Click `Start Conversation` three times (creates three parallel agent sessions)
3. Assign one of the above task sets to each agent
4. Agents work asynchronously — you review artifacts from each as they complete

**Example parallel prompt for Agent 2 (Infrastructure):**

```
You are Agent 2 focused on infrastructure and DevOps for Tabbre Core Chain MVP.

Your job:
1. Create docker-compose setup for 3-node localnet
2. Create Makefile with targets: init, genesis, start, stop, reset, test
3. Create CI/CD pipeline (GitHub Actions) for:
   - Build chain binary
   - Run unit tests
   - Run integration tests
   - Package release artifacts
4. Create monitoring setup:
   - Prometheus metrics export
   - Grafana dashboard for block height, validator uptime, transaction throughput
5. Create deployment docs for testnet RPC nodes

Refer to architecture.md, tokenomics-implementation.md, and validator-policy.md from Agent 1 for context.

Generate all artifacts in infra/ directory.
```


***

## Workstream Skills (Optional Power-User Setup)

Create **workspace skills** for Tabbre-specific knowledge that agents should reference.

**How to add:**

1. In your workspace root, create `.agents/skills/tabbre-tokenomics/`
2. Create `.agents/skills/tabbre-tokenomics/SKILL.md`:
```markdown
---
name: tabbre-tokenomics
description: Tabbre Core Chain tokenomics reference. Use when implementing or documenting TABB and BRE token logic.
---

# Tabbre Tokenomics Skill

When working on Tabbre token logic, always refer to these constraints:

## TABB Token
- Fixed supply: 1,000,000,000 (1 billion)
- Minted once at genesis
- Never inflate or mint additional TABB post-genesis
- Used for: staking, governance, collateral
- Allocations:
  - 20% private sale (200M)
  - 10% public sale (100M)
  - 10% founders (100M)
  - 10% community (100M)
  - 50% reserve (500M, vesting 0.5%/month for 120 months)

## BRE Token
- Transactional token
- Used for: transaction fees, validator rewards, energy sales (future)
- Issued as credit by BRE Monetary Authority (future phase)
- MVP: seed initial supply for testing

## Validator Economics
- Validators stake TABB
- Validators earn BRE from transaction fee income
- No TABB block rewards or emissions
- Slashed TABB is burned

## When implementing
- Always check: am I preserving fixed TABB supply?
- Always check: are fees paid in BRE, not TABB?
- Always check: are validator rewards BRE-based, not TABB-based?
```

3. Agent will now auto-load this when working on token-related tasks[^1]

***

## Workflow: Common Tasks

Create **workspace workflows** for repeatable tasks.

**Example workflow: Generate unit tests**

Create `.agents/workflows/generate-tests.md`:

```markdown
---
name: generate-tests
description: Generate comprehensive unit tests for Tabbre Core Chain modules
---

# Test Generation Workflow

When generating tests for Tabbre Core Chain:

1. **Test coverage requirements:**
   - Test TABB supply is always 1B
   - Test BRE is deducted for transaction fees
   - Test TABB staking and delegation flows
   - Test slashed TABB is burned
   - Test validator reward distribution in BRE
   - Test governance proposals

2. **Test structure:**
   - Use Go testing framework
   - Create test files in module_test.go format
   - Include table-driven tests where appropriate
   - Mock external dependencies

3. **Test naming:**
   - TestTokenSupplyInvariant
   - TestFeePaymentInBRE
   - TestValidatorRewardsInBRE
   - TestSlashingBurn

4. **Always include:**
   - Setup and teardown
   - Positive and negative test cases
   - Edge cases (zero amounts, max values, etc.)
```

**To trigger:** Type `/generate-tests` in agent chat[^1]

***

## Emergency: Reset Agent Context

If agent gets confused or starts violating Tabbre constraints:

**Prompt:**

```
Stop. Re-read the tabbre-tokenomics-constraints global rule.

Confirm you understand:
1. TABB supply is fixed at 1B
2. BRE is the fee token
3. Validators earn BRE rewards, not TABB
4. No TABB inflation

Now, review your last implementation. List any places where you violated these constraints. Then fix them.
```


***

## Final Launch Prompt

When you're ready to prepare handover:

**Prompt 13: Final Audit and Handover**

```
Prepare Tabbre Core Chain MVP for final handover and launch readiness review.

Deliver:
1. Complete file tree and repo state
2. End-to-end setup instructions from zero (clone → localnet running)
3. Localnet verification steps (prove TABB/BRE tokenomics work correctly)
4. Test coverage summary (what % of flows are tested)
5. List of all deviations from ideal Tabbre end-state design
6. List of all MVP compromises and technical debt
7. List of all security-sensitive customizations flagged for audit
8. Recommended 30/60/90 day post-launch roadmap
9. Operator handover notes (for validator consortium and DevOps team)
10. Executive technical summary (2-page overview for non-technical stakeholders)

Also provide clear assessment:
- What is production-ready now
- What is testnet-ready only
- What MUST be audited before mainnet
- What can be deferred to phase 2

Generate all artifacts in docs/handover/
```


***

## Tips for Antigravity-Specific Workflow

1. **Use Planning mode** for architecture, design, and complex implementation[^1]
2. **Use Fast mode** for quick fixes, renaming, simple edits[^1]
3. **Review Implementation Plans** before agent starts coding — this is where you catch misunderstandings[^1]
4. **Comment on artifacts** like Google Docs — highlight specific lines and leave feedback[^1]
5. **Use browser agent** for testing deployed contracts (agent can open browser, interact with dApp, capture screenshots)[^1]
6. **Enable terminal execution** carefully — use Review-driven mode so agent asks before running destructive commands[^1]
7. **Check artifacts panel** regularly — agent produces task lists, walkthroughs, screenshots, browser recordings[^2][^1]
8. **Undo is your friend** — if agent makes wrong changes, use "Undo changes up to this point" in chat[^1]

***

**This workbook gives you a production-grade prompt sequence for building Tabbre Core Chain MVP with Antigravity's agent-first architecture. The key is: review artifacts, leave comments, iterate. Let the agent do the heavy lifting, but you stay in control of business logic and architecture decisions.**
<span style="display:none">[^10][^11][^12][^13][^14][^5][^6][^7][^8][^9]</span>

<div align="center">⁂</div>

[^1]: https://codelabs.developers.google.com/getting-started-google-antigravity

[^2]: https://developers.googleblog.com/build-with-google-antigravity-our-new-agentic-development-platform/

[^3]: https://antigravity.google

[^4]: https://antigravity-ide.com

[^5]: https://www.youtube.com/watch?v=-0Irz8G0PEE

[^6]: https://venturebeat.com/ai/google-antigravity-introduces-agent-first-architecture-for-asynchronous

[^7]: https://www.theverge.com/news/822833/google-antigravity-ide-coding-agent-gemini-3-pro

[^8]: https://www.codecademy.com/article/how-to-set-up-and-use-google-antigravity

[^9]: https://www.reddit.com/r/ChatGPTCoding/comments/1p35bdl/i_tried_googles_new_antigravity_ide_so_you_dont/

[^10]: https://dev.to/fabianfrankwerner/an-honest-review-of-google-antigravity-4g6f

[^11]: https://en.wikipedia.org/wiki/Google_Antigravity

[^12]: https://antigravityide.org

[^13]: https://www.youtube.com/watch?v=Mz_epTVBIRE

[^14]: https://dev.to/this-is-learning/my-experience-with-google-antigravity-how-i-refactored-easy-kit-utils-with-ai-agents-2e54

