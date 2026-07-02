# terraform-provider-sendgrid

Terraform provider (Go) for SendGrid, published to the Terraform Registry
under `bid-ops-development/sendgrid`.

## Layout

- `main.go` — provider entry point
- `sendgrid/` — resource + data-source implementations
- `sdk/` — SendGrid API client wrappers
- `docs/` — generated provider docs (registry consumes these)
- `examples/` — usage examples
- `scripts/`, `tools/` — build helpers
- `terraform-registry-manifest.json` — registry publication manifest
- `GNUmakefile` — build targets
- `go.mod`, `go.sum`

## Commands

```bash
make build             # compile to $GOPATH/bin
make test              # unit tests
make testacc           # acceptance tests (hits SendGrid API — needs creds)
make docs              # regenerate docs/ (must be committed)
```

## Conventions & Gotchas

- **`docs/` is committed and consumed by the Terraform Registry.** Always
  regenerate with `make docs` after adding/changing a resource — stale docs
  ship to registry users.
- **Acceptance tests hit the real SendGrid API** — needs `SENDGRID_API_KEY`
  and can create real resources. Run them with intent, not by default.
- **Registry manifest changes** must be coordinated with a version bump.
- **This is publicly consumed** via the Terraform Registry — breaking
  changes need a major-version bump and release notes.
- Cross-cutting patterns: see the team
  [CLAUDE.md baseline](https://github.com/bid-ops-development/proposals-and-planning/tree/main/proposals/claude-md-baseline).
