.PHONY: assets

GENERATE_TLS_CERT = $(GOPATH)/bin/generate-tls-cert
GO_BINDATA := $(GOPATH)/bin/go-bindata

# Add files that change frequently to this list.
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


serve:
	go run cmd/clipper-server/main.go
