Makefile snippets that I use for building golang projects.
Everyone likes their own, I like these :)  Just wanted a public
place to put them so I don't end up with scattered copies everywhere.

Use in a makefile as follows:

````
-include $(GOPATH)/src/github.com/boyvinall/gomake/gotools.mk
$(GOPATH)/src/github.com/boyvinall/gomake/gotools.mk:
	go get github.com/boyvinall/gomake/gotools.mk
````