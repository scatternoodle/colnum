BINARY_NAME := colnum

build:
	env GOOS=linux GOARCH=amd64 go build -o ./bin/linux/amd64/${BINARY_NAME}
	env GOOS=linux GOARCH=386 go build -o ./bin/linux/386/${BINARY_NAME}
	env GOOS=darwin GOARCH=amd64 go build -o ./bin/ios/${BINARY_NAME}
	env GOOS=linux GOARCH=arm GOARM=5 go build -o ./bin/raspberry-pi/${BINARY_NAME}

clean:
	rm -rf ./bin
