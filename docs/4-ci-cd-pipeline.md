# [RFC 4] CI/CD Pipeline
Status: Draft

This RFC contains initial comprehensive analysis of how the CI/CD pipeline should be introduced. It explains about the code quality, build strategy, release strategy, and deploy strategy.

## Flow

The flow of CI/CD pipelining will be as follow:
- lint -> linter to check the code quality -> job: quality check
- test -> test the source code (code coverage included) -> job: quality check
- build -> build the binary app (will be handled by goreleaser)
- release -> release to binary or dokerized component -> job: build and release (needs: quality check)
- deploy -> deploy to gcp vm -> job: deploy


jobs:
- quality check: [lint, test]
- build and release: [build, release] <- needs: quality check
- deploy: [deploy ]

branches jobs:
- deploy branches: [deploy ] -> needs to check this new thing
- release tag branches: [quality check, build and release]
- main/master and other branch: [quality check]

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

Build and release strategy will be used goreleaser. The server binary will be dockerized for the linux and macos platform. The app must have a version, and overrided it with ldflags.
In pipeline, use goreleaser for building the binary etc. Build command in makefile only being used for the development only (wrapping the goreleaser build).
Release strategy will use the default practices.

Release branches should have the major version syntax: eg, releases/v1, releases/v2, etc, while the tags is the complete version but under the same major version on the corresponding release branches. tag v1.0.2 will be on the branch releases/v1, etc.

- Release branch is only being used to create tag, any branches are not allowed to create a tag.
- It should provide the git hook to prevent the tag being created outside the release branches

## Deploy Strategy

For now, deployment strategy is scoped only on GCP VM. The CI/CD will be deploy the build through terraform. For now, deploy srategy is not provided yet.
