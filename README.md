# Twitter Wall
The Twitter Wall was a project of my university and got me the highest grade. 

## Installation
The server needs a local MongoDB and the Go language. You need to add the twitter API key to .env file. Following command ín the terminal starts the server:

```bash
$ go run ./main.go
```

The server can also be compiled:

```bash
$ go build && ./twitter-wall-server.exe // Windows
$ go build && ./twitter-wall-server     // Unix
```

The website is then on `https://127.0.0.1` and `https://localhost`. TLS / https is a self-signed certificat and only for testing.

The frontend is based on Svelte. If you want to change something to it you need to use following commands:

```bash
$ npm install
$ npm run build
```