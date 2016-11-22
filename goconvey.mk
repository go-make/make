#  See https://github.com/boyvinall/gomake/goconvey.md

export GOCONVEY:=$(GOPATH)/bin/goconvey

.PHONY: goconvey
goconvey: $(GOCONVEY)
	nohup $(GOCONVEY) > goconvey.out &

$(GOCONVEY):
	$(GO) get -u github.com/smartystreets/goconvey
