.PHONY: assets

GENERATE_TLS_CERT = $(GOPATH)/bin/generate-tls-cert
GO_BINDATA := $(GOPATH)/bin/go-bindata
JUSTRUN := $(GOPATH)/bin/justrun

# Add files that change frequently to this list.
GO_FILES = $(shell find . -name '*.go')
GO_NOASSET_FILES := $(filter-out ./assets/bindata.go,$(GO_FILES))
WATCH_TARGETS = $(shell find ./static ./templates -type f)

$(GO_BINDATA):
	go get -u github.com/kevinburke/go-bindata/...

assets/bindata.go: $(WATCH_TARGETS) | $(GO_BINDATA)
	$(GO_BINDATA) -o=assets/bindata.go --nocompress --nometadata --pkg=assets templates/... static/...

assets: assets/bindata.go

$(GENERATE_TLS_CERT):
	go get -u github.com/Shyp/generate-tls-cert

certs/leaf.pem: | $(GENERATE_TLS_CERT)
	mkdir -p certs
	cd certs && $(GENERATE_TLS_CERT) --host=localhost,127.0.0.1

# Generate TLS certificates for local development.
generate_cert: certs/leaf.pem | $(GENERATE_TLS_CERT)

$(GOPATH)/bin/clipper-server: $(GO_FILES)
	go install ./cmd/clipper-server

$(GOPATH)/bin/clipper-server:

serve: assets $(GOPATH)/bin/clipper-server
	$(GOPATH)/bin/clipper-server

$(JUSTRUN):
	go get -u github.com/jmhodges/justrun

watch: | $(JUSTRUN)
	$(JUSTRUN) --delay=100ms -c 'make assets serve' $(GO_NOASSET_FILES) $(WATCH_TARGETS)

test:
	go test ./...

release:
	bump_version minor cmd/clipper-server/main.go
