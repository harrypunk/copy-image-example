.PHONY: clean

build/bootstrap: example/aws-lambda/main.go
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $@ -tags lambda.norpc $<

build/lambda-copy.zip: build/bootstrap
	zip -j $@ $<

clean:
	rm build/*
