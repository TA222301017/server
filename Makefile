all:
	go build -o ./build/main.exe main.go
	mkdir -p build/keys
	cp .env ./build/.env

dump-db:
	pg_dump -U root -f dump/ta_database.sql ta_database

clean:
	rm -r build