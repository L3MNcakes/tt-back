# tt-back

`go get github.com/l3mncakes/tt-back`

`cd $GOPATH/github.com/l3mncakes/tt-back`

`cp config.go.example config.go`

Open *config.go* and set your RIAK_ADDRESSES to the correct ip and
ports.

`make build`

`make start`

Navigate to http://$DOCKER_HOST:8080/
