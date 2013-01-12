all: githook 6g tool 

PREFIX=code.google.com/p/mx3

PKGS=\
	$(PREFIX)/gpu\
	$(PREFIX)/gpu/ptx\
	$(PREFIX)/cpu\
	$(PREFIX)/uni\
	$(PREFIX)/mag\
	$(PREFIX)/dump\
	$(PREFIX)/nimble\
	$(PREFIX)/core\


6g: ptx
	go install -v $(PKGS)
	go install -v 
	go install $(PREFIX)/render

tool:
	make -C tools/dump
	make -C tools/table

GCCGO=gccgo -gccgoflags '-static-libgcc -O4 -Ofast -march=native'

gccgo: ptx
	go install -v -compiler $(GCCGO) $(PKGS)
	go install -v -compiler $(GCCGO)

ptx:
	make -C gpu/ptx -j8

githook:
	ln -sf $(CURDIR)/pre-commit .git/hooks/pre-commit
	ln -sf $(CURDIR)/post-commit .git/hooks/post-commit

racetest:
	go test -race $(PKGS)
	make -C test

test: 6gtest  unittest #gccgotest #re-enable gccgotest when gcc up to date with go 1.1

unittest:
	make -C test

6gtest: 6g
	go test -i $(PKGS) 
	go test $(PKGS) 

gccgotest: gccgo
	go test -i -compiler=$(gccgo) $(PKGS)
	go test -compiler=$(gccgo) $(PKGS)

.PHONY: bench
bench: 6g
	go test -test.bench $(PKGS)

.PHONY: gccgobench
gccgobench: gccgo
	go test -compiler=$(gccgo) -test.bench $(PKGS)

.PHONY: clean
clean:
	go clean -i -x $(PKGS)
	rm -rf $(GOPATH)/pkg/gccgo/$(PREFIX)/
	make clean -C gpu/ptx

.PHONY: count
count:
	wc -l *.go */*.go gpu/ptx/*.cu
