# docker-bug

## Build error

```
$ go build main.go
# github.com/docker/cli/cli/command/container
../go/pkg/mod/github.com/docker/cli@v29.3.0+incompatible/cli/command/container/tty.go:63:23: \
cannot range over 10 (untyped int constant): requires go1.22 or later (-lang was set to \
go1.16; check go.mod)
```
