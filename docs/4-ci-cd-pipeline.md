# CI/CD Pipeline

This RFC contains initial comprehensive analysis of how the CI/CD pipeline should be introduced. It explains about the code quality, build strategy, release strategy, and deploy strategy.

## Code Quality

For ensuring the quality of the code, linter and comprehensive test should be introduced. This RFC breakdowns about which linters could be used using golangci-lint and guide of standardization the test structure.

**Linters**

These following linters will be used as guardrails for ensuring every piece of code.
```
- bodyclose
- contextcheck
- decorder
- dupl
- durationcheck
- errname
- errorlint
- exhaustive  
- exportloopref
- gci
- gocritic -> think to add this later?
- gocognit
- goconst
- gofmt
- goimports
- gomnd
- importas
- interfacebloat
- lll
- maintidx -> interesting but it contains cyclomatic complexity too
- misspell
- nilerr
- nilnil
- prealloc
- tenv
- unconvert
- unparam
- usestdlibvars
- wrapcheck -> rootcause err
```

## Build And Release Strategy

Build and release strategy will be used goreleaser. The server binary will be dockerized for the linux and macos platform.

## Deploy Strategy

For now, deployment strategy is scoped only on GCP VM. The CI/CD will be deploy the build through terraform.
