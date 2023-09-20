run:
	go run main.go
build:
	go build main.go
windows:

	CGO_CPPFLAGS="-Wno-error" CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build main.go 
#https://github.com/gen2brain/raylib-go/issues/184