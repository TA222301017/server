PLATFORMS=darwin linux windows
ARCHITECTURES=386 amd64 arm
BINARY=server

release:
	mkdir -p ./release
	cp .env ./release/.env;\
	for GOOS in $(PLATFORMS); do for GOARCH in $(ARCHITECTURES); do \
	export GOOS=$$GOOS;\
	export GOARCH=$$GOARCH;\
	if [[ "$$GOOS" == "windows" ]]; then\
		go build -o ./release/$(BINARY)-$$GOOS-$$GOARCH.exe main.go;\
	else\
		go build -o ./release/$(BINARY)-$$GOOS-$$GOARCH main.go;\
	fi;\
	done done\

build:
	if [[ "$(shell go env GOOS)" == "windows" ]]; then\
		go build -o ./build/$(BINARY).exe main.go; \
	else \
		go build -o ./build/$(BINARY) main.go; \
	fi; \
	cp .env ./build/.env

dump-db:
	pg_dump -U root -f dump/ta_database.sql ta_database

clean:
	rm -rf build
	rm -rf release