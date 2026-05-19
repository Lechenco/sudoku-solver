clean:
	rm -rf build

build:
	go build -o build/app cmd/app/main.go 

start:
	./build/app

run:
	go run cmd/app/main.go