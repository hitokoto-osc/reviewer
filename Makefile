ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "hitokoto-reviewer"
DOCKER_NAME = "hitokoto-reviewer"

include ./hack/hack.mk

get-gf:
	@echo "Installing gf..."
	wget -O gf "https://github.com/gogf/gf/releases/latest/download/gf_$(go env GOOS)_$(go env GOARCH)" && chmod +x gf && ./gf install -y && rm ./gf

init-env: dep
	@echo "Initializing environment..."
	@go mod tidy;
	@npm install;

dep: # get dependencies
	@echo "Installing Dependencies..."
	go mod download

lint: ## Lint Golang files
	@echo;
	@echo "Linting go codes with golangci-lint...";
	@golangci-lint run ./... --color always

vet:
	@echo "Checking go codes with go vet..."
	go vet ./...

test:
	@echo "Testing..."
	@go test -short ${PKG_LIST}

test-coverage:
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST}
	@cat cover.out >> coverage.txt

clean:
	@rm -f coverage.txt
	@rm -f cover.out

release:
	@echo "Releasing by GoReleaser..."
	@goreleaser release --rm-dist

precommit: lint test
	go fmt ./...
	go mod tidy
	git add .
