WORK_DIR:=${CURDIR}
PRO_NAME:=$(notdir ${WORK_DIR})
pro_bin:=bin/${PRO_NAME}
target:=${pro_bin}

GitCommitLog:=$(shell git rev-parse --verify HEAD)
BuildTime:=$(shell  date "+%Y-%m-%d %H:%M:%S")
BuildGoVersion:=$(shell go version)

LDFLAG:="-X 'main.GitCommitLog=${GitCommitLog}' -X 'main.BuildTime=${BuildTime}' -X 'main.BuildGoVersion=${BuildGoVersion}'"

.PHONY: all build clean run test fmt help

all: fmt build vet test
	
build:
	@echo "go build -ldflags ${LDFLAG} -o ${target} cmd/${PRO_NAME}/main.go"
	@go build -ldflags ${LDFLAG} -o ${target} cmd/${PRO_NAME}/main.go
	@./${target} -version

fmt:
	@echo "go  fmt ./..."
	@go fmt ./...
	@echo "fmt over"

vet:
	@echo "go vet ./..."
	@go vet ./...
	@echo "vet successful"

test:
	@go test -v --count=1 ./...

clean:
	rm -f ${target}

help:
	@echo "make ${PRO_NAME}"
	@echo "GitCommitLog ${GitCommitLog}"
	@echo "BuildGoVersion ${BuildGoVersion}"
