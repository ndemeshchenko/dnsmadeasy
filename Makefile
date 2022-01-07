OS ?= $(shell go env GOOS)
ARCH ?= $(shell go env GOARCH)

ifeq (Darwin, $(shell uname))
	GREP_PREGEX_FLAG := E
else
	GREP_PREGEX_FLAG := P
endif

GO_VERSION ?= $(shell go mod edit -json | grep -${GREP_PREGEX_FLAG}o '"Go":\s+"([0-9.]+)"' | sed -E 's/.+"([0-9.]+)"/\1/')

test:
	go test -APIAccessKey ${API_ACCESS_KEY} -APISecretKey ${API_SECRET_KEY} -v .

lint:
	golangci-lint run ./...

vendor:
	go mod vendor