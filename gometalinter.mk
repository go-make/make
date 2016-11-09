export GOMETALINTER:=$(GOPATH)/bin/gometalinter

$(GOMETALINTER): | $(GOPATH)
	$(GO) get github.com/alecthomas/gometalinter
	$(GOMETALINTER) --install

ifndef GOMETALINTER_LINELENGTH
GOMETALINTER_LINELENGTH:=100
endif

ifndef GOMETALINTER_DEADLINE
GOMETALINTER_DEADLINE:=30s
endif

ifndef GOMETALINTER_CYCLO
GOMETALINTER_CYCLO:=15
endif

ifndef GOMETALINTER_OPT
GOMETALINTER_OPT:=\
	--exclude='should have comment or be unexported' \
	--exclude='should have comment \(or a comment on this block\) or be unexported' \
	--exclude='error return value not checked \(defer ' \
	--deadline=$(GOMETALINTER_DEADLINE) \
	--line-length=$(GOMETALINTER_LINELENGTH) \
	--cyclo-over=$(GOMETALINTER_CYCLO)
endif
