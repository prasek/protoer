# TODO: run golint, errcheck, staticcheck, unused, ineffassign
.PHONY: default
default: deps checkgofmt vet predeclared unused ineffassign predeclared errchack test

.PHONY: ci
ci: deps checkgofmt vet predeclared unused ineffassign predeclared errchack testcover

.PHONY: deps
deps:
	go get -d -v -t ./...

.PHONY: updatedeps
updatedeps:
	go get -d -v -t -u -f ./...

.PHONY: install
install:
	go install ./...

.PHONY: checkgofmt
checkgofmt:
	@if [ -n "$$(go version | awk '{ print $$3 }' | grep -v devel)" ]; then \
		echo "gofmt -s -l ." ; \
		if [ -n "$$(gofmt -s -l .)"  ]; then \
			for dir in $$(gofmt -s -l .); do \
				gofmt -d -s $$dir; \
			done; \
		  gofmt -s -l . ; \
			echo "Run gofmt on the above files!"; \
			exit 1; \
		fi; \
	fi

# workaround https://github.com/golang/protobuf/issues/214 until in master
.PHONY: vet
vet:
	@echo go vet ./...  --ignore internal/goprotos
	@for dir in $$(go list ./... | grep -v 'internal/golang/testprotos'); do \
		go vet $$dir ; \
	done

.PHONY: staticcheck
staticcheck:
	@go get honnef.co/go/tools/cmd/staticcheck
	staticcheck ./...

.PHONY: unused
unused:
	@go get honnef.co/go/tools/cmd/unused
	unused ./...

.PHONY: ineffassign
ineffassign:
	@go get github.com/gordonklaus/ineffassign
	ineffassign .

.PHONY: predeclared
predeclared:
	@go get github.com/nishanths/predeclared
	predeclared .

# Intentionally omitted from CI, but target here for ad-hoc reports.
.PHONY: golint
golint:
	@go get github.com/golang/lint/golint
	golint -min_confidence 0.9 -set_exit_status ./...

# Intentionally omitted from CI, but target here for ad-hoc reports.
.PHONY: errchack
errcheck:
	@go get github.com/kisielk/errcheck
	errcheck ./...

.PHONY: test
test:
	go test -race ./...

.PHONY: generate
generate:
	go generate github.com/prasek/protoer/internal/gogo/testprotos/
	go generate github.com/prasek/protoer/internal/golang/testprotos/

.PHONY: testcover
testcover:
	@echo go test -covermode=atomic ./... 
	@echo "mode: atomic" > coverage.out
	@for dir in $$(go list ./proto/... | grep -v 'internal/'); do \
		go test -race -coverprofile profile.out -covermode=atomic $$dir ; \
		if [ -f profile.out ]; then \
			tail -n +2 profile.out >> coverage.out && rm profile.out ; \
		fi; \
	done; \
	go test -race -coverprofile profile.out -covermode=atomic -coverpkg ./proto ./proto/gogo; \
	if [ -f profile.out ]; then \
		tail -n +2 profile.out >> coverage.out && rm profile.out ; \
	fi; \
	go test -race -coverprofile profile.out -covermode=atomic -coverpkg ./proto ./proto/golang; \
	if [ -f profile.out ]; then \
		tail -n +2 profile.out >> coverage.out && rm profile.out ; \
	fi; \
