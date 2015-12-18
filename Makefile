mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(dir $(mkfile_path))
volume_source := $(current_dir):/go/src/github.com/tadasv/gorocksdb
workdir := /go/src/github.com/tadasv/gorocksdb

shell:
	docker run --rm -ti -v $(volume_source) -w $(workdir) golang:1.5.1 bash
