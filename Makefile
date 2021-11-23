
build:
		go build cmd/media/main.go -o bin/media

web:
		npm --prefix=svelte run build

run: web
		go run cmd/media/main.go --mediaDir /home/anuar/media/