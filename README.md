# dealls-go

### How To Run This Project

> Make Sure you have run the dealls.sql in your mysql

Since the project already use Go Module, I recommend to put the source code in any folder but GOPATH.

#### Run the Testing

```bash
$ make tests
```

#### Run the Applications

Here is the steps to run it with `docker-compose`

```bash
#move to directory
$ cd workspace

# Clone into your workspace
$ git clone https://github.com/dandyfirmansyah18/dealls-go.git

#move to project
$ cd go-clean-arch

# Run the application
$ make up

# The hot reload will running

# Execute the call in another terminal
$ curl localhost:9090/articles
```

### Tools Used:

In this project, I use some tools listed below. But you can use any simmilar library that have the same purposes. But, well, different library will have different implementation type. Just be creative and use anything that you really need.

- All libraries listed in [`go.mod`](https://github.com/dandyfirmansyah18/dealls-go/blob/master/go.mod)
- ["github.com/vektra/mockery".](https://github.com/vektra/mockery) To Generate Mocks for testing needs.
