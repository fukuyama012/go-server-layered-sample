# go-server-layered-sample
Go API Sever, layered architecture sample

## Install

```
$ git clone https://github.com/fukuyama012/go-server-layered-sample.git
```

- Setup [goenv](https://github.com/syndbg/goenv)

- Install Specified Go Version
```
$ goenv install `cat .go-version`
```

- Check Go version (same .go-version)
```
$ go version
go version go1.15.3 darwin/amd64 (for example)
```


## How to run

- Run server
```
$ make run
```

- Access with browsers
```
http://localhost:1323/v1/example/1
```

## Reflect changes

- Setup [wire](https://github.com/google/wire)
 
- If you change dependency injection info (handler/wire.go), you need to call this command.
```
$ make wire
```