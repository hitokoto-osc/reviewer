ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "template-single"

include ./hack/hack.mk

init-env:
	@echo "Initializing environment..."
	@go mod tidy;
	@npm install;

get-tools:
	@echo "Installing tools..."
	go install github.com/mgechev/revive@latest


lint: get-tools ## Lint Golang files
	@echo;
	@echo "Linting go codes with revive...";
	@revive -config ./.revive.toml -formatter stylish ${PKG_LIST}

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