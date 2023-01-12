.PHONY: list
.DEFAULT_GOAL := list

GO_FILES=./cmd/*

list:
	@echo "Available targets:"
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs -n1 echo "  "
build: 
	@go build -o ./bin $(GO_FILES)
clean:
	go clean
	rm -f ./bin/*
	rm .test-coverage
test:
	go test
check:
	go test
cover:
	go test -coverprofile .test-coverage $(GO_FILES)
	go tool cover -html=.test-coverage
# run:
# 	./"${BIN_FILE}"
lint:
	golangci-lint run --enable-all
