export GLIDE:=$(GOPATH)/bin/glide

#
#  Glide v0.12 no longer creates git repo under the vendor directory, which
#  makes things slower IMHO. I also sometimes find it useful to edit files
#  in vendored packages during debug, and it's less easy to see what you edited
#  if you can't just "git diff" in there.
#
GLIDE_VERSION:=v0.11.1

GLIDE_ARCH:=linux-amd64
GLIDE_OPT_INSTALL:=--cache --cache-gopath

$(GLIDE): | $(GOPATH)
	curl -sL https://github.com/Masterminds/glide/releases/download/$(GLIDE_VERSION)/glide-$(GLIDE_VERSION)-$(GLIDE_ARCH).tar.gz | tar -xz --to-stdout -f- $(GLIDE_ARCH)/glide > $@


