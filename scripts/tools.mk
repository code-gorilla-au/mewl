
get-tools: ## Get tools used 
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
	go install github.com/mitranim/gow@latest