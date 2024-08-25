.PHONY: setup-new-year format-go
setup-new-year:
	scripts/setup-new-year.sh

format-go:
	gofmt -w .