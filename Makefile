# When you first clone the repo, start here!
.PHONY: start
start: proto/generate 
	go mod tidy

.PHONY: serve
serve: # Runs the server for local development
	CONFIG_PATH=config/config_local.hcl go run cmd/%%wp_project%%-api/main.go

.PHONY: proto/lint
proto/lint: # Lint protobufs
	buf lint .

.PHONY: proto/generate
proto/generate: # Compile protobufs
	buf generate

.PHONY: test
test: # Run tests
	go test ./...
