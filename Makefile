ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "template-single"

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

get-lint-tools:
	@echo "Installing lint tools..."
	go install github.com/mgechev/revive@latest

get-tools: get-lint-tools
	@echo "Install tools done."


lint: get-lint-tools ## Lint Golang files
	@echo;
	@echo "Linting go codes with revive...";
	@REVIVE_FORCE_COLOR=1 revive -config ./.revive.toml -formatter stylish ${PKG_LIST}

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

precommit: vet lint test
	go fmt ./...
	go mod tidy
	git add .
