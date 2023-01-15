all:
	go build -o ./build/main.exe main.go
	mkdir -p build/keys
	cp .env ./build/.env

clean:
	rm -r build