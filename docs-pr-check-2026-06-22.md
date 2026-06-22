# Documentation PR check: merged Tempo PRs, 2026-06-15 through 2026-06-22

Generated on 2026-06-22 for merged PRs in `grafana/tempo`.

Source query:

```bash
gh pr list --repo grafana/tempo --state merged --search "merged:2026-06-15..2026-06-22" --json number,title,mergedAt,author,labels,url --limit 200
```

## Table

| PR | Title | Classification | Notes |
|----|-------|----------------|-------|
| [#7502](https://github.com/grafana/tempo/pull/7502) | chore(deps): update grafana/tempo-vulture docker tag to v3.0.2 (main) | No docs required | Renovate dependency update in `example/docker-compose/migrate-to-3/docker-compose.yaml`; no user-facing behavior change. |
| [#7493](https://github.com/grafana/tempo/pull/7493) | ci: sign and attest tempo-vulture container images | No docs required | CI/CD signing and attestation workflow for `tempo-vulture`. PR added internal workflow documentation in `.github/workflows/sign-and-attest.md`, but no public docs update appears required under the docs-pr-check criteria. |
| [#7492](https://github.com/grafana/tempo/pull/7492) | perf(collector): make scoped distinct string collector lock-free | No docs required | Performance optimization in internal collectors; PR states limits remain soft and behavior is result-equivalent for users. |
| [#7490](https://github.com/grafana/tempo/pull/7490) | chore(deps): update grafana/shared-workflows/login-to-gar action to v1.0.3 (main) | No docs required | Renovate patch update for GitHub Actions workflows only. |
| [#7481](https://github.com/grafana/tempo/pull/7481) | [release-v2.10] chore: use large runner for release workflow | No docs required | Backport of a CI-only release workflow runner change. |
| [#7468](https://github.com/grafana/tempo/pull/7468) | feat: add experimental trace diff command to Tempo CLI | Docs needed | Adds user-facing experimental `tempo-cli experimental trace-diff` command. PR checklist has "Documentation added" unchecked, no `docs/` files changed, and existing `docs/sources/tempo/operations/tempo_cli.md` does not mention `trace-diff` or `trace-patch-v0`. |
| [#7465](https://github.com/grafana/tempo/pull/7465) | spanfilter: fast paths for intrinsic policy matching | No docs required | Internal performance fast path for span filter intrinsic policy matching; PR says custom patterns fall back to regex with identical results. |
| [#7459](https://github.com/grafana/tempo/pull/7459) | chore(deps): update actions/checkout action to v6.0.3 (main) | No docs required | Renovate patch update for GitHub Actions workflows only. |
| [#7378](https://github.com/grafana/tempo/pull/7378) | chore(deps): update actions/checkout digest to df4cb1c (main) | No docs required | Renovate digest update for GitHub Actions workflows only. |

## Gap summary

1. **PR [#7468](https://github.com/grafana/tempo/pull/7468): add public Tempo CLI documentation for experimental trace diff.**
   - Suggested existing file to update: `docs/sources/tempo/operations/tempo_cli.md`.
   - The new command is registered as `tempo-cli experimental trace-diff`.
   - Required arguments from the command source:
     - `trace-a`: baseline local trace JSON file.
     - `trace-b`: comparison local trace JSON file.
   - Options from the command source:
     - `--format trace-patch-v0`: output format; currently the only supported format.
     - `-o, --out <path>`: write output to a file instead of stdout.
     - `--pretty`: pretty-print JSON output.
   - The docs should explain that the command compares two local trace JSON files and emits experimental `trace-patch-v0` output, including added, removed, and modified spans plus per-span field or attribute changes.
   - The docs should mention the experimental status and known current limitations from the PR body: no DRAIN normalization for high-cardinality span names, and no differences for scope, event, or link attributes.

2. **No other public documentation gaps found for the reviewed PRs.**
   - PR [#7493](https://github.com/grafana/tempo/pull/7493) may become a public docs topic if Tempo decides to document image signature and provenance verification for users, but under the current docs-pr-check rules it is classified as CI/CD with no required docs update.
