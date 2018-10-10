## Go Webassembly example

Learn more about Webassembly here: https://github.com/golang/go/wiki/WebAssembly

To try this example, compile the client-side Go program to WASM like this:

    GOOS=js GOARCH=wasm go build -o static/main.wasm wasm/main.go

Now start the web server like this:

    go run cmd/webserver.go

Go to http://localhost:8080 and check the console!
