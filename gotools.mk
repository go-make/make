export GOMETALINTER:=$(GOPATH)/bin/gometalinter
export GLIDE:=$(GOPATH)/bin/glide
export RICE:=$(GOPATH)/bin/rice

$(GOMETALINTER):
	go get github.com/alecthomas/gometalinter
	$(GOMETALINTER) --install

GLIDE_ARCH:=linux-amd64
$(GLIDE):
	curl -sL https://github.com/Masterminds/glide/releases/download/v0.11.1/glide-v0.11.1-$(GLIDE_ARCH).tar.gz | tar -xz --to-stdout -f- $(GLIDE_ARCH)/glide > $@

$(RICE):
	go get github.com/GeertJohan/go.rice/rice

ifndef GOMETALINTER_LINELENGTH
GOMETALINTER_LINELENGTH:=150
endif

ifndef GOMETALINTER_DEADLINE
GOMETALINTER_DEADLINE:=1000s
endif

ifndef GOMETALINTER_CYCLO
GOMETALINTER_CYCLO:=15
endif

ifndef GOMETALINTER_OPT
GOMETALINTER_OPT:=\
	--exclude='should have comment or be unexported' \
	--exclude='should have comment \(or a comment on this block\) or be unexported' \
	--exclude='error return value not checked \(defer ' \
	--enable-all \
	--deadline=$(GOMETALINTER_DEADLINE) \
	--line-length=$(GOMETALINTER_LINELENGTH) \
	--cyclo-over=$(GOMETALINTER_CYCLO)
endif