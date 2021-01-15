# HellDB

HellDB is a distributed key-value store that lives in memory. It's written in Go to be performant and fault tolerant with wrappers existing in multiple languages.

## Installation

Download a binary from the releases page or build it yourself - 

```sh
$ go get github.com/heldatabase/helldb
$ helldb
2021/01/15 15:59:04 serving hell on http://localhost:8080
```

## Usage

HellDB is extremely simple to use. After running it on port 8080, you can send queries over http like so - 

```sh
$ curl -s -X POST http://localhost:8080/query -d "query=GET hello;" | json_pp
```

Or using some programming language - 

```python
import requests
url = 'http://localhost:8080/query'
res = requests.post(url, data={'query': 'GET name;'}).json()
```

Or you can use a client libraries - 

1. [HellPy](https://github.com/helldatabase/hellpy/)
2. [HellJS](https://github.com/helldatabase/helljs/)
3. [HellGo](https://github.com/helldatabase/hellgo/)
4. [HellRb](https://github.com/helldatabase/hellrb/)
