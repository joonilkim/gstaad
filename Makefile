svcs = $(shell find svc/*)
thishash = $(shell git log --pretty=format:'%H')
newhash = $(shell git log --pretty=format:'%H' -n 1)

.PHONY: all ${svcs}
all: ${svcs}

${svcs}:
	support/githash.sh $@ && cd $@ && $(MAKE) build && $(MAKE) deploy

githash:
	$(thishash) = $(shell git log --pretty=format:'%H' -n 1 -- $@)
