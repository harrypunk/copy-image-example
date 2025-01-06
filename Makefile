.PHONY: build zip

build:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o build/bootstrap -tags lambda.norpc main.go

zip:
	zip -j build/copy.zip build/bootstrap
