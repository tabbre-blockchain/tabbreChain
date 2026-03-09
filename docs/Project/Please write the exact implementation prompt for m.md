<img src="https://r2cdn.perplexity.ai/pplx-full-logo-primary-dark%402x.png" style="height:64px;margin-right:32px"/>

# Please write the exact implementation prompt for my agent: “Modify Cosmos EVM starter so BRE is the gas token and TABB is the staking token."

Use this as a copy-paste implementation prompt for your coding agent.

## Implementation prompt

```text
Modify the Cosmos EVM starter so that BRE is the gas/fee token and TABB is the staking token for the Tabbre Core Chain MVP.

You are working in an existing Cosmos SDK + Cosmos EVM starter repository. Your job is to make the minimum safe set of changes needed to support the Tabbre token model.

Non-negotiable requirements:
1. TABB is the staking token.
2. BRE is the gas/fee token for all transactions.
3. Validators and delegators are rewarded from BRE fee income, not TABB inflation.
4. TABB has a fixed total supply minted at genesis only.
5. No post-genesis TABB inflation or minting path may remain active.
6. The chain should remain EVM-compatible and expose Ethereum JSON-RPC.
7. Keep changes as small, reviewable, and testable as possible.
8. Prefer configuration and app wiring changes over large custom modules.
9. Do not silently introduce assumptions that contradict Tabbre tokenomics.

Business rules to preserve:
- TABB = staking, governance, collateral
- BRE = transaction charges and usage fees
- Launch model = permissioned consortium validators
- Slashed TABB should be burned or routed to a burn-compatible sink
- MVP should support localnet, testnet, and mainnet configuration paths

Your objectives:
A. Identify all code and config locations that assume a single token is both gas token and staking token.
B. Modify the app so staking uses TABB and fees use BRE.
C. Ensure EVM transactions sent through JSON-RPC pay gas in BRE.
D. Ensure native Cosmos SDK transactions also pay fees in BRE.
E. Ensure fee collection and distribution work with BRE.
F. Ensure TABB fixed-supply semantics remain intact.
G. Produce a working localnet proving the model end-to-end.

Implementation scope:
1. Token denomination setup
2. Genesis supply and balances
3. Staking module configuration
4. Fee market configuration
5. Ante handler / fee deduction path
6. Fee collector and distribution path
7. Slashing handling for TABB
8. Localnet scripts
9. Integration tests
10. Documentation

Detailed requirements:

1. Token model
- Define TABB as the staking bond denom.
- Define BRE as the canonical fee denom.
- Recommend sensible base denoms and metadata.
- Prefer an EVM-friendly decimal strategy for BRE if needed for wallet/tooling compatibility.
- Ensure token metadata is registered correctly.

2. Genesis
- Mint all TABB supply at genesis only.
- Do not leave any active path that can mint more TABB later.
- Seed BRE balances for localnet and tests.
- Set up genesis validator accounts with TABB stake.
- Create deterministic localnet accounts for testing both TABB and BRE flows.

3. Staking
- Configure x/staking bond denom to TABB.
- Ensure validator creation, delegation, undelegation, and redelegation use TABB.
- Ensure slashing affects TABB stake, not BRE balances.
- If slashed TABB is not automatically burned by the starter, implement the smallest safe burn-compatible approach and document it.

4. Fees and gas
- Make BRE the required fee token for native Cosmos transactions.
- Make BRE the gas token for EVM transactions through Cosmos EVM.
- Ensure EIP-1559 fee market values are interpreted in BRE.
- Ensure mempool admission / min gas price checks use BRE.
- Ensure fee deduction removes BRE from sender accounts.
- Ensure fee collector module receives BRE.

5. Rewards and distribution
- Ensure validators receive BRE fee income through the normal fee collection / distribution flow.
- Ensure delegator rewards, if enabled in the starter, are paid in BRE.
- Ensure no TABB inflation-based rewards remain active.
- Document any limitation if distribution in BRE requires careful handling.

6. Governance and upgrades
- Keep standard governance and upgrade modules working.
- Governance deposit denom may remain configurable, but clearly state the current implementation choice.
- Do not introduce custom governance logic unless required for compilation or correctness.

7. EVM / JSON-RPC compatibility
- Keep Ethereum JSON-RPC enabled.
- Test at least:
  - eth_chainId
  - eth_blockNumber
  - eth_gasPrice or equivalent fee endpoint
  - contract deployment
  - value transfer
- Validate that gas is actually paid in BRE.
- Document any wallet display quirks caused by denom/decimal choices.

8. Localnet
- Produce a reproducible 3-validator localnet.
- Include scripts or Make targets for:
  - init
  - genesis
  - start
  - stop
  - reset
  - test
- Seed accounts with TABB and BRE.
- Include one account primarily for staking tests and one for EVM fee tests.

9. Tests
Create or update tests to prove:
- TABB bond denom is active.
- BRE fee denom is active.
- TABB total supply does not increase after chain start.
- Native Cosmos tx fees are charged in BRE.
- EVM tx fees are charged in BRE.
- Validator rewards accrue in BRE.
- Staking actions use TABB.
- Slashing reduces TABB as expected.
- JSON-RPC remains operational.

10. Documentation
Create or update:
- README section explaining TABB vs BRE
- docs/token-model.md
- docs/fees.md
- docs/staking.md
- docs/localnet.md
- docs/known-limitations.md

Required working style:
- Work in small, reviewable steps.
- Before changing code, first inspect and list the exact files and code paths likely involved.
- Then propose a step-by-step implementation plan.
- Then implement only step 1 and stop for review.
- After each step, provide:
  - summary of changes
  - files changed
  - rationale
  - how to test
  - any consensus/state/rpc impact
  - open issues

Important constraints:
- Do not refactor unrelated code.
- Do not add major new modules unless absolutely necessary.
- If BRE-as-fee-token requires more invasive changes than expected, stop and provide:
  1. the blocker,
  2. why it occurs,
  3. the least-risk workaround,
  4. the trade-offs.
- If the starter’s mint/inflation module conflicts with fixed-supply TABB, disable, remove, or neutralize it in the smallest safe way and explain exactly how.
- If an assumption is needed, record it in ASSUMPTIONS_FOR_AGENT.md.
- Maintain:
  - CHANGELOG_FOR_AGENT.md
  - OPEN_ISSUES_FOR_AGENT.md
  - ASSUMPTIONS_FOR_AGENT.md

Deliverables expected by the end of the task:
1. Modified source code
2. Localnet scripts/config
3. Passing tests for the token split model
4. Clear docs
5. A short architecture note explaining how BRE fee flow and TABB staking flow work

Start now with:
Step 0: inspect the repository and identify all relevant files for staking denom, fee denom, fee market, ante handlers, mint/inflation, genesis, distribution, and localnet setup.
Do not implement anything yet.
Return:
- repository impact map
- implementation plan
- key risks
- first code step proposal
```


## Shorter version

If you want a tighter version for an agent that already has your project context loaded, use this:

```text
Modify the Cosmos EVM starter for Tabbre Core Chain MVP so that:
- TABB is the staking token
- BRE is the gas/fee token
- validator rewards come from BRE fee income
- TABB is fixed supply at genesis only
- no TABB inflation remains active

Tasks:
1. Inspect the repo and identify all files affecting:
   - staking bond denom
   - fee denom
   - ante handlers
   - fee market / EIP-1559
   - distribution
   - mint/inflation
   - genesis
   - localnet
2. Produce an implementation plan.
3. Implement in small steps:
   - set TABB as bond denom
   - set BRE as fee token for Cosmos txs
   - set BRE as gas token for EVM txs
   - route collected BRE through fee collector/distribution
   - disable or remove post-genesis TABB minting
   - seed genesis balances and validators
   - update localnet scripts
   - add integration tests
4. After each step, report:
   - files changed
   - rationale
   - how to test
   - state/consensus/rpc impact
5. Stop and ask for review after each step.

Start with repo inspection only. No code changes yet.
```


## Recommended companion guardrail

Paste this immediately after the main prompt if you want to keep the agent from drifting:

```text
Guardrails:
- Never make TABB the fee token.
- Never add TABB inflation.
- Never assume the staking token and fee token must be the same.
- Prefer the smallest possible change that preserves Cosmos EVM compatibility.
- If BRE fee support is blocked by a starter assumption, explain the exact blocker before proceeding.
- Keep all changes auditable and easy to review.
```


## Best way to use it

Send the main prompt first, then when the agent returns the repo impact map, reply with:

```text
Proceed with step 1 only.
Implement the smallest safe change first.
Do not touch unrelated files.
```

After that, keep it on a short leash:

- review every step,
- require tests after each meaningful change,
- do the fee-path spike before broader cleanup.

If you want, I can also write:

- a **Cursor/Claude Code version**,
- an **Antigravity version**,
- or a **two-agent version** where one agent handles protocol code and another handles localnet/tests.

