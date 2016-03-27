# tt-back

```bash
$ go get github.com/l3mncakes/tt-back
$ cd $GOPATH/github.com/l3mncakes/tt-back
$ cp config.go.example config.go
```

Open *config.go* and set your RIAK_ADDRESSES to the correct ip and
ports.

```bash
$ make build
$ make test
$ make start
```

Navigate to http://$DOCKER_HOST:8080/
