default: gotest/html

.PHONY: gotest gotest/html

gotest:
	go test -covermode=count -v ./...

gotest/html:
	go test -v ./... -coverprofile=coverage.out && go tool cover -html=coverage.out