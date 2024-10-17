package main

import "fmt"

// build時にldfflag経由でバージョンを埋め込む
// go build -o example -ldflags "-X main.version=v1.2.3" go/example/main.go
var version string

func main() {
	fmt.Printf("Example %s\n", version)
}
