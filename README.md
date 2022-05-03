# Go Fibonacci

<img src="https://img.shields.io/badge/status-development-yellowgreen"/></a>

Golang Fibonacci is a single endpoint service that listens for `GET` requests on `/fib?n=<n>` and calculates n-th Fibonacci sequence number of the given number.

## Dependencies

This project was designed against Go 1.16.
The service uses Kubernetes and Docker to startup.

- Docker
- Kubernetes
  - Minikube

## Decisions taken on this implementation

The service as it is, solves up to element 93 of the Fibonacci sequence with an iterative approach, which has been chosen for its efficiency compared to the alternatives. The application is limited by the length of the int64 type, so the maximum value that can be stored in that data type is 18,446,744,073,709,551,615.

However, I have implemented a solution using the [`math/big` library](./internal/models/fibonacci.go#L81) to demonstrate that even if the computational cost skyrockets, the Fibonacci sequence can still be calculated and stored on data types beyond 64 bits using Golang.

Both benchmarking and testing can be proved by running the functions contained in [`fibonacci_test.go`](./internal/models/fibonacci_test.go)

    make bench
    make test

To explore what else you can do with the Makefile:

    make

## Instructions to run the project

By running the following command, the project will be (docker) built, deployed on minikube and started.

    make k8-minikube-start

The next step is to open a [minikube tunnel](https://minikube.sigs.k8s.io/docs/commands/tunnel/#minikube-tunnel), to access the service through localhost. I have opted for this option to avoid having to make too many changes to `fib.yaml` in case of sending the service to production.

    minikube tunnel

After successful execution, the service should be running on port 8080. Hit the *health* endpoint to check that all is running as expected:

- [:8080](http://localhost:8080/)

You should receive a `{"status": "ok"}` as the response.

The next step is to hit the `fib` endpoint and start querying the API

    curl http://localhost:8080/fib?n=1
    > 1

    curl http://localhost:8080/fib?n=72
    > 498454011879264

## Packaging

    .
    ├── cmd                     # Entrypoint, Main API
    └── internal
        ├── handlers            # HTTP layer
        ├── models              # Business logic
        └── web                 # Framework for common HTTP related tasks

## Author

Noel Ruault - [@noelruault](https://github.com/noelruault)

## Further development

[High-performance Fibonacci numbers generator in Go](https://blog.abelotech.com/posts/fibonacci-numbers-golang/)

### Matrix

[The Go Playground](https://play.golang.org/p/MpFWnMlEiSR)

[The optimal algorithm of Fibonacci sequence (golang code) - Programmer Sought](https://www.programmersought.com/article/66706671612/)

[Go Fibonacci sequence generator](https://stackoverflow.com/a/17604464)

[Fibonacci matrix-exponentiation](https://rosettacode.org/wiki/Fibonacci_matrix-exponentiation#Go)

### Phi

[The Golden Section - the Number](http://www.maths.surrey.ac.uk/hosted-sites/R.Knott/Fibonacci/phi.html#section5)

### Complexity

[Determining complexity for recursive functions (Big O notation)](https://stackoverflow.com/questions/13467674/determining-complexity-for-recursive-functions-big-o-notation)
