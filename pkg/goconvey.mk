export GOCONVEY:=$(GOPATH)/bin/goconvey

.PHONY: goconvey
goconvey: $(GOCONVEY)
	nohup $(GOCONVEY) > goconvey.out &

$(GOCONVEY):
	$(call PROMPT,Installing $@)
	$(GO) get github.com/smartystreets/goconvey

clean-tools::
	rm -f $(GOCONVEY)

update-tools::
	$(GO) get -u github.com/smartystreets/goconvey
