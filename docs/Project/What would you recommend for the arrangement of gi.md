<img src="https://r2cdn.perplexity.ai/pplx-full-logo-primary-dark%402x.png" style="height:64px;margin-right:32px"/>

# What would you recommend for the arrangement of github repos for the Tabbre Blockchain?

I’d recommend a **hybrid model**: one primary code monorepo for the chain and tightly coupled developer assets, plus a small number of separate repos for infrastructure, governance artifacts, and public documentation. That gives Tabbre fast MVP iteration and atomic changes where the chain, tests, and tooling move together, while keeping security-sensitive ops and public-facing materials cleanly separated.[^1][^2][^3]

## Recommended layout

For a Cosmos SDK + Cosmos EVM chain, the protocol code, localnet tooling, integration tests, EVM examples, and developer docs are tightly interdependent, and monorepos work well when projects share components and need coordinated changes.  At the same time, infrastructure repos are often better separated by environment or operational purpose because access control and deployment concerns differ from application code.[^2][^3][^4][^1]

**Recommended GitHub org:** `tabbre-blockchain`

**Recommended repos**

- `tabbre-chain` — main monorepo for the blockchain itself.[^4][^2]
- `tabbre-infra` — infrastructure as code, deployment manifests, observability, secrets references, RPC and validator environment templates.[^3]
- `tabbre-docs` — public documentation site, tokenomics summaries, developer portal, validator guides, governance docs.[^2]
- `tabbre-governance` — governance proposals, ADRs, genesis decisions, parameter sets, audit reports, validator admissions.[^4]
- `tabbre-explorer` — only if you build or heavily customize your own explorer; otherwise keep explorer config inside `tabbre-chain`.[^2]
- `tabbre-contracts` — only if you expect a substantial independent Solidity codebase beyond chain smoke tests; otherwise keep EVM examples/tests inside `tabbre-chain`.[^5]


## Main monorepo

The main chain repo should be the engineering center of gravity because Cosmos SDK chains combine protocol logic, modules, config, and execution-layer integration in one codebase.[^6][^4]

Suggested structure for `tabbre-chain`:

```text
tabbre-chain/
  app/
  cmd/tabbred/
  proto/
  x/
  precompiles/              # only if you add custom EVM precompiles
  contracts/                # optional MVP examples and protocol-adjacent contracts
  tests/
    integration/
    e2e/
    evm/
  localnet/
  scripts/
  docker/
  docs/
  genesis/
    local/
    testnet/
    mainnet-candidate/
  config/
  ops/
    upgrade-handbooks/
    snapshots/
  Makefile
  README.md
```

Put anything that must evolve **in lockstep** with consensus logic in this repo: app wiring, token handling, ante/fee logic, staking config, localnet scripts, and end-to-end tests.  That is especially important for Tabbre because the TABB/BRE split is a protocol-level concern, so code and tests should change atomically in one PR.[^1][^6][^4][^2]

## Separate repos

`tabbre-infra` should be separate because validator infrastructure, RPC fleet management, Terraform, Helm charts, and environment promotion have different permissions and release cycles than protocol code.  I would structure it by platform and environment, with strict separation between dev/testnet/mainnet overlays.[^3]

Example:

```text
tabbre-infra/
  terraform/
    shared/
    testnet/
    mainnet/
  kubernetes/
    testnet/
    mainnet/
  ansible/
  monitoring/
    prometheus/
    grafana/
  runbooks/
  policies/
```

`tabbre-governance` should store the institutional memory of the chain: ADRs, genesis allocations, validator onboarding decisions, parameter change drafts, upgrade proposals, and signed release notes. This makes governance auditable and reduces the risk that critical decisions are buried in engineering PRs.[^4]

Example:

```text
tabbre-governance/
  adrs/
  genesis/
  proposals/
    testnet/
    mainnet/
  validator-admissions/
  economics/
  audits/
  release-approvals/
```


## What not to split yet

For the MVP, I would **not** start with lots of tiny repos. Multi-repo setups help when teams are independent, but they add coordination overhead and slow down small teams working on tightly coupled systems.  Since Tabbre Chain MVP likely has a small engineering team and many cross-cutting changes, too much repo fragmentation will create friction.[^7][^2]

So for MVP:

- Keep chain code, EVM integration tests, localnet scripts, and protocol-adjacent contracts together in `tabbre-chain`.[^1][^2]
- Split out `tabbre-infra` early because access control and deployment discipline matter.[^3]
- Split out `tabbre-docs` only if docs are meant to be publicly versioned and published independently; otherwise keep docs in the monorepo at first.[^2]


## My concrete recommendation

Start with **four repos**:

- `tabbre-chain`
- `tabbre-infra`
- `tabbre-governance`
- `tabbre-docs`

Then add `tabbre-explorer` or `tabbre-contracts` only when those become real standalone products or teams. That keeps the repo model simple, secure, and scalable for a Cosmos EVM chain MVP.[^3][^2]

A good rule is:

- If it affects consensus, keep it in `tabbre-chain`.[^4]
- If it affects deployment security, put it in `tabbre-infra`.[^3]
- If it affects formal decision-making, put it in `tabbre-governance`.
- If it targets external readers, put it in `tabbre-docs`.

Would you like me to turn this into a **GitHub organization blueprint** with repo names, branch strategy, CODEOWNERS, permissions, and CI/CD responsibilities?
<span style="display:none">[^10][^11][^12][^13][^14][^15][^8][^9]</span>

<div align="center">⁂</div>

[^1]: https://www.thoughtworks.com/en-es/insights/blog/agile-engineering-practices/monorepo-vs-multirepo

[^2]: https://www.gitkraken.com/blog/monorepo-vs-multi-repo-collaboration

[^3]: https://www.harness.io/blog/gitops-repo-structure

[^4]: https://docs.cosmos.network/sdk/v0.53/learn/intro/overview

[^5]: https://github.com/orgs/community/discussions/161996

[^6]: https://cosmos-docs.mintlify.app

[^7]: https://kodus.io/en/monorepo-vs-multi-repo-strategy/

[^8]: https://www.reddit.com/r/softwarearchitecture/comments/1pqlco8/monorepo_vs_multiple_repos_for_backend_mobile_web/

[^9]: https://www.sanity.io/answers/best-approach-to-structuring-a-repo-for-a-studio-and-using-separate-repos-for-front-end-and-studio-

[^10]: https://github.com/cosmos/cosmos-sdk

[^11]: https://www.reddit.com/r/git/comments/dp5w0m/best_practices_for_repository_structure/

[^12]: https://www.quicknode.com/builders-guide/tools/cosmos-sdk-by-interchain-foundation

[^13]: https://www.reddit.com/r/devops/comments/uyr6hx/thoughts_on_mono_repo_vs_multi_repo_how_do_you/

[^14]: https://www.reddit.com/r/devops/comments/rmadw8/monorepo_vs_multirepo/

[^15]: https://www.lfdecentralizedtrust.org/blog/2018/11/05/6-blockchain-best-practices-enterprises-need-to-know

