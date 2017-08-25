export GOCONVEY:=$(GOPATH)/bin/goconvey

.PHONY: goconvey
goconvey: $(GOCONVEY)
	nohup $(GOCONVEY) > goconvey.out &

$(GOCONVEY):
	$(call PROMPT,Installing $@)
	$(GO) get -u github.com/smartystreets/goconvey
