BINARY_NAME := colnum

clean:
	rm -rf ./bin

build:
	@make clean
	env GOOS=linux GOARCH=amd64 go build -o ./bin/linux/amd64/${BINARY_NAME}
	env GOOS=linux GOARCH=386 go build -o ./bin/linux/386/${BINARY_NAME}
	env GOOS=darwin GOARCH=amd64 go build -o ./bin/ios/${BINARY_NAME}
	env GOOS=linux GOARCH=arm GOARM=5 go build -o ./bin/raspberry-pi/${BINARY_NAME}
	mkdir -p ./bin/releases
	tar -czvf ./bin/releases/${BINARY_NAME}_linux_amd64.tar.gz ./bin/linux/amd64/${BINARY_NAME}
	tar -czvf ./bin/releases/${BINARY_NAME}_linux_386.tar.gz ./bin/linux/386/${BINARY_NAME}
	tar -czvf ./bin/releases/${BINARY_NAME}_darwin_amd64.tar.gz ./bin/ios/${BINARY_NAME}
	tar -czvf ./bin/releases/${BINARY_NAME}_raspberry_pi.tar.gz ./bin/raspberry-pi/${BINARY_NAME}
