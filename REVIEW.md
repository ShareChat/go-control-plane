# Review guide for go-control-plane (ShareChat fork)

Code owners: `@shivamnarula` `@krhitesh7` `@jensoncs` `@prakharjoshi` `@abhiroop93` `@prateek010` `@Abhipsa20` (the nexus owners; this repo has no `CODEOWNERS`, and `OWNERS.md` lists the upstream maintainers). Last reviewed: 2026-11-07.

This is the `nexus` branch of a fork of `envoyproxy/go-control-plane`, based on upstream `main` at `v0.12.0-106-gbdba4bba1`. ShareChat's changes are in `pkg/cache/` and `pkg/server/stream/v3/stream.go`; the rest is upstream code. The only consumer is xlr8 in ShareChat/nexus (`xlr8/go.mod:148`). It links `pkg/cache/types`, `pkg/cache/v3`, `pkg/log`, `pkg/resource/v3`, `pkg/server/stream/v3`, `pkg/wellknown` and the generated `envoy/` types.

## What failure looks like here
- `pkg/cache/v3/` (`simple.go`, `snapshot.go`, `delta.go`, `cache.go`), `pkg/cache/types/`, `pkg/server/stream/v3/`: it misroutes. Envoy keeps a deleted resource, misses an update, or a stream stops getting pushes, all with no error.
- Upstream code xlr8 never runs: `pkg/server/` other than `stream/v3`, `pkg/client/`, `pkg/integration/`, `pkg/test/`, `LinearCache`, `MuxCache`. A change there does not reach production, so reach alone does not make it Important.

## Protected behaviours (Important whenever a change reaches them)
These apply when a change alters the behaviour. A log-only or metric-only edit nearby is not Important.
- Watch registration: `CreateDeltaWatch` and `CreateWatch` decide "client is up to date" and register the watch under one `info.mu` hold (`simple.go:1033`, `simple.go:881-892`). Splitting the two loses a push with no error.
- Locks and sends: writers take `info.mu`, then `Snapshot.Mu` (`notifyWatches`, `simple.go:613-627`). `respondDelta` never blocks on the channel while `info.mu` is held (`simple.go:1163-1172`). A blocking send, or the reverse lock order, deadlocks the node.
- Dead streams: a send on a closed channel returns `ErrResponseChannelClosed` and the watch is dropped (`simple.go:966-972`, `:1145-1152`, `:734`, `:769`). Folding it back into `(nil, nil)` leaks watches that look healthy.
- Deletes: `DeleteResources` bumps the type version and notifies (`simple.go:568-602`). A type whose last resource was deleted (empty map) is answered with removals (`simple.go:1048`). A type never populated (nil map) is parked (`simple.go:1116`). Changing nil-versus-empty handling makes Envoy keep deleted resources or drop live ones.
- In-place edits: `UpsertResources`, `BatchUpsertResources` and `DeleteResources` edit the stored `*Snapshot` under `Snapshot.Mu`. `notifyWatches` skips a snapshot that was replaced meanwhile (`simple.go:620`). Writing the edit back would resurrect a cleared node or roll back a newer snapshot.
- Delta versions come from the caller (`types.ResourceWithTTL.Version`), not a hash. `ConstructVersionMap` errors on an empty one (`snapshot.go:182`), and `GetDeltaDiscoveryResponse` drops a resource with no name or version (`cache.go:242-249`). A change to how versions are set or compared means no push, or a push loop.
- Node keys: snapshots and status are keyed by `hash.ID(node)` in 256 shards (`simple.go:200-253`), and `GetStatusKeys` returns those keys (`simple.go:1245`). Keying by hash slot again serves one node's snapshot to another.

## Contracts with other repos
An exported-API change in `pkg/cache/` or `pkg/server/stream/v3/` changes together with ShareChat/nexus `xlr8/`. The finding names the xlr8 caller:
- `SnapshotCache` methods (`simple.go:75-107`): `xlr8/pkg/providers/xds_state/state.go:204-271`, `xlr8/pkg/xds/reconcile.go:291`.
- `CreateDeltaWatch` returns `(cancel, delayed)` (`cache.go:68`): `xlr8/pkg/xds/ads.go:430` and `xlr8/pkg/xds/ack_timeout.go:86` start the ACK clock when `delayed` is false.
- `Snapshot.Mu` (`snapshot.go:38`): read-locked by `xlr8/pkg/providers/xds_state/snapshot_read.go:31`.
- `GetResourceName` for `*LbEndpoint` reads `xdstp.endpoint` `uri` (`resource.go:93-100`), which `xlr8/pkg/model/leds.go:100` writes.
- `StreamState.WatchesResources` and its `/*` prefixes (`stream.go:94`): `xlr8/pkg/xds/delta_xds_watches.go:94`.
- `NodeHash` (`status.go:27-31`): implemented by `xlr8/pkg/providers/xds_state/state.go:71-90`. A new method breaks xlr8's build.
- `RawDeltaResponse.GetNextVersionMap` and `GetContext` (`cache.go:295-299`): `xlr8/pkg/xds/ads.go:318,331`.
- `types.Resource` needs `MarshalVTStrict` (`types.go:12`), which exists only with `-tags vtprotobuf` (`envoy/**/*_vtproto.pb.go:1`). xlr8 builds with that tag (`xlr8/Dockerfile:19`). `go.mod:1` must stay `github.com/envoyproxy/go-control-plane`, because xlr8 swaps it in with `replace`.
- A merge reaches Envoy only after a `v0.12.0-v3.2.3-beta.N` tag is cut from `nexus` and `xlr8/go.mod:148` is bumped. Ask which tag, and which xlr8 PR.

## Checks this repo needs
- CI runs no tests. `ci.yaml` fails at `build/do_ci.sh:9` (`goimports@latest` needs a newer Go than `Dockerfile.ci:1`), before `make test`. `golangci-lint.yaml` fails typecheck because it builds without `-tags vtprotobuf`. Both were red on every recent PR, including the merged #22. A cache change with no local test run is a finding.
- Run `go test -race -tags vtprotobuf -count=1 ./pkg/cache/... ./pkg/server/...` (about 20 s). `make test` lacks the tag (`Makefile:27`). At `beta.25`, 14 upstream tests in `pkg/cache/v3` fail and the `pkg/server/v3` tests don't compile. The tests in `delete_resources_test.go`, `failures_test.go`, `simple_race_test.go`, `sotw_watches_test.go`, `status_keys_test.go` and `respond_delta_internal_test.go` pass with no race. A new failure, or a cache fix with no test in those files, is a finding.
- A change to an xlr8-facing symbol: run xlr8's `./pkg/providers/xds_state/` and `./pkg/xds/` tests against the branch (local `replace`, `TEST_RUN=true`, `-tags vtprotobuf`).

## Always needs a person
- `.github/workflows/`, `build/`, `Dockerfile.ci`, `Makefile`.
- A merge from upstream `envoyproxy/go-control-plane`, or an `envoy/` proto sync (`ci/sync_envoy.sh`; `envoy-sync.yaml` checks out `main`, not `nexus`). Its conflicts land on the protected code above.
- A bump of `planetscale/vtprotobuf`, `google.golang.org/protobuf` or `grpc` in `go.mod`: it changes how every resource is encoded.

## Skip
- Generated, mirrored from `envoyproxy/envoy` and unchanged since the fork base: `envoy/`, `contrib/envoy/`, `ratelimit/`.
- Upstream examples, fixtures and docs, unchanged: `examples/`, `internal/`, `sample/`, `support/`, `xdsmatcher/`, `docs/`.

## Ask the author, don't flag
- A changed park-or-respond decision (full channel, nil map, closed stream): what triggers the next push, and which test pins it?
- An upstream test edited or deleted to go green: is that upstream behaviour meant to be dropped?
- A new type URL on an edit path: does xlr8 send it, and does the path guard `types.UnknownType` as `DeleteResources` does (`simple.go:570`)?
- A new `Warnf` or `Errorf` per watch or per push: xlr8 forwards only those two levels (`xlr8/pkg/providers/xds_state/state.go:121-126`). What is the volume across the fleet?

## Never delete (cleanup agents)
- `DrainResources` and `UpdateVirtualHosts` (`simple.go:495-606`): no-ops, but part of `SnapshotCache` and called by xlr8 (`xlr8/pkg/providers/xds_state/state.go:260,271`).
- The `recover()` in `respond` and `respondDelta` (`simple.go:966`, `simple.go:1145`): it is the only signal that a stream is gone.

## How findings should look
- At most 3 nits. No style points; `goimports` and `golangci-lint` own style.
- Evidence is a file:line, or a command with its output. Never a guess from a name.
- Start the summary with a count, and say plainly when there is nothing to fix.
