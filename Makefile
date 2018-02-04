.PHONY: deps
deps:
	go get -u github.com/kardianos/govendor
	govendor sync
	go get github.com/mattn/goveralls
	go get github.com/go-playground/overalls

.PHONY: coveralls
coveralls:
	overalls -project=github.com/namely/k8s-configurator -covermode=count
	goveralls -coverprofile=overalls.coverprofile -service=travis-ci