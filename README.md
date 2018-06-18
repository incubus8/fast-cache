<p align="center"> 
    <img src="assets/logo.png"/>
</p>

# Fast Cache

Ultra Lightweight, Simple, and High Performance In-Memory Caching Server via REST.

## Features

* Hoshicorp LRU
* Hoshicorp ARC-LRU
* [CCache LRU](github.com/karlseguin/ccache)

## Requirements

1. Golang 1.10.2 <
2. Install `dep` to manage Golang dependency

## How to run?

You must get dependencies via Dep. You will see vendor folder if success.

```shell
$ dep ensure -update 
$ go run main.go
```

## How to use?

1. Adding LRU cache of `xyz` with JSON data `{'a':"2", 'c':1}`

```shell
$ curl -X POST -d '{'a':"2", 'c':1}' localhost:8080/api/v1/lru/xyz
```

2. Get LRU cache of `xyz`

```shell
$ curl -X GET localhost:8080/api/v1/lru/xyz
```

3. Remove LRU cache of `xyz`

```shell
$ curl -X DELETE localhost:8080/api/v1/lru/xyz
```


## Credits

Bosan Tech (c) 2018

