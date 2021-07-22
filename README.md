# Fibber

(Yes, the name is deliberately cheeky.)

## Build and run

This app was developed on Debian Buster. We are assuming that postgresql server and golang are already 
installed. 

[Golang download and installation instructions][https://golang.org/dl/]

In order to run the app, execute the following 
from the root directory:

```
make install-db
make run
```

Unit test suite can be run with:

```
make test
```

## API

The paths below correspond to the 3 functions:

- fetch the Fibonacci number given an ordinal (e.g. Fib(11) == 89, Fib(12) == 144), 
- fetch the number of memoized results less than a given value (e.g. there are 12 intermediate results less than 120)
- clear the data store.

```
"/ordinal/{ordinal_n}"
"/cardinality_less/{cardinality_x}"
"/clear"
```

To test from the command line using curl:

```
curl http://127.0.0.1:8080/ordinal/20
curl http://127.0.0.1:8080/cardinality_less/20
curl http://127.0.0.1:8080/clear
```

Replies are returned as plain text. No provision was made to enforce GET/POST/etc.

## Database

If one wishes to query the database, use the command below with password 'secret123':

```
psql -h localhost -U fib_app -d fib_db
```

## Architecture

I am simply using golang as a cache, while storing the sequence values in the database.

It's hard to apply architecture including persistence to such a toy problem. For one thing, Binet's formula 
with an extended precision math library, makes any persistence redundant:

[Binet's formula][https://www.wikihow.com/Calculate-the-Fibonacci-Sequence]

However, I went ahead and implemented this anyhow as an exercise. This architecture can be extended by sharding
the app server by ordinal. The cache would be extended, such that a minimum ordinal and fibonaci sequence value 
would be held by each shard. Each shard would then be responsible for an interval in the domain.

I am left wondering if this exercise is meant to be an analogy to some kind of practical existing style of 
application. If the test taker is meant to guess this analogy, the test is unclear and somewhat arbitrary.

Also note that the golang's internal scheduler is being used as the queue for writing to the database.

## Algorithm

Instead of using the typical recursive functions, I've changed the recursion into a loop. This has the same 
complexity as the recursive implementation, but should be more efficient, since there is no usage of the stack.

## Performance

I've included a flag to execute a wall clock measurement of extending the cache to a chosen ordinal.

```
$ go run cmd/main.go -time=1000000
2021/07/21 18:27:37 Starting timing session for N=1000000
2021/07/21 18:27:38 SyncExtendToN took 1.055017569s
$ go run cmd/main.go -time=1000000000
2021/07/21 18:27:49 Starting timing session for N=1000000000
2021/07/21 18:27:49 SyncExtendToN took 488.187435ms
$
```

To scale this architecture, one would want to shard it according to the domain of ordinals as discussed above. 