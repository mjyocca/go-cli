fmt:
	gofmt -s -l -w .

build:
	go build -ldflags="-X 'version.Version=dev' -X 'main.Application=app'"
