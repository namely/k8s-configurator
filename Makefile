.PHONY: deps
deps:
	@go get -u github.com/kardianos/govendor
	@govendor fetch +external