PROJECT_NAME := "github-actions-demo-go"
PKG := "github.com/brpaz/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

test: ## Run unittests
	@go test -short ${PKG_LIST}