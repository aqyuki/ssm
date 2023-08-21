build:
	@go build

test:
	@go test ./... -v

cover:
	@go test ./... -cover

clean:
	@echo Remove binalys
	@rm ssm.exe