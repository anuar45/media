
build:
	go build cmd/media/main.go -o bin/media

web:
	npm --prefix=front run build

run: web
	go run cmd/media/main.go --mediaDir data/pics 
