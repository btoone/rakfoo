BINARY_NAME=rakfoo

build:
	go build -o bin/${BINARY_NAME}

run:
	go run .

clean:
	go clean
	rm bin/${BINARY_NAME}

test:
	go test --short -v
