
all: build_go

build_go:
	go get k8s.io/apimachinery/pkg/api/resource@v0.0.0-20190413052414-40a3f73b0fa2
	go build resource_converter.go

clean:
	rm -f resource_converter
