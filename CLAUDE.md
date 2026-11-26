# Project: eramba-go-client

## Commands

- `golangci-lint fmt`: golang code format
- `golangci-lint run`: golang code linting
- `go test ./...`: unit tests
- `go run ./examples/<name>`: run an example against a live Eramba instance (needs `ERAMBA_URL`, `ERAMBA_USER`, `ERAMBA_PASSWORD`, loaded from `.env`)


## Commits

- Always sign off commits: `git commit -s` (adds a `Signed-off-by:` trailer).


## Architecture

- `/model`: Eramba type models
- `/examples`: Example runs with marshal and unmarshal
- `types.go`: endpoint constructors on `Client`; v2 endpoints use a `v2/` path and the `ReadWriteClientV2*` clients


## Eramba v2 API

The swagger is not reliable for v2. Verify new v2 endpoints by running an example against a live instance.

- Relation lists (`risks`, `third_party_risks`, ...) accept string ids, so the v1 list types (`model.Risks` etc.) work.
- Sending `[]` / `null` on `Put` removes relations and dates.
- Owners and contacts come back as `{"users":[...],"groups":[...]}` and are sent as `["User-1","Group-2"]` (`model.UsersAndGroups`). v1 uses `[]UserOrGroup`.
- Comments are created with POST `{path}/{id}/comments` (v1 appends `/add`).
