
build:
	go build -o bin/media cmd/media/main.go

web:
	npm --prefix=front run build

run: web
	go run cmd/media/main.go --mediaDir /home/anuar/pics
