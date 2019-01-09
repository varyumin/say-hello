# say-hello
say-hello is a simple tool for checking connections tcp, upd (layer 4) and http, https (layer 7). It requires a simple configuration.
## Installation
### Build from source (requires Go 1.10+)
```
$ git clone https://github.com/varyumin/say-hello.git
$ cd say-hello
$ dep ensure
$ go build -ldflags "-X main.VERSION=v0.0.1" -o "webserver" .
```
### Docker
```
$ docker run -d -p 8080:8080 varyumin/say-hello
```
### Kubernetes (Helm Chart)
```
$ helm upgrade say-hello helm/say-hello -i --debug --wait
```
## Configuration
### Environments
| NAME  | DEFAULT  | DESCRIPTION |
|---|---|---|
| PORT  | 8080  |Bind port web server|
| TIMEOUT-WEB  | 30  |Timeout from web server|
| TIMEOUT-CHECK  |  5 |Timeout from check resurce|
### Flags
`--port` - Bind port web server

`--timeout-web` - Timeout from web server

`--timeout-check` - Timeout from check resurce

## Example

![](art/screencast.gif)