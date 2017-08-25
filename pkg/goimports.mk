GOIMPORTS:=$(GOPATH)/bin/goimports

$(GOIMPORTS):
	go get golang.org/x/tools/cmd/goimports
