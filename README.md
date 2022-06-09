# oauth-sample-api-golang

To facilitate integration with the Dock's Caradhras API, a series of code examples were developed considering best practices and the best way to manage the data returned by the endpoints. So feel free to use this code example as you see fit.

## Description

This code example demonstrates how to implement oauth 2.0 authorization considering the use of caching to not make unnecessary extra calls.

To make this possible, we use the field returned with the expiration time information. Every time we use the token, we check if the expiration time has passed and perform a new authentication only if necessary.

## Getting Started

### Prerequisites

* Some programming experience. The code here is pretty simple, but it helps to know something about functions.
* A tool to edit your code. Any text editor you have will work fine. Most text editors have good support for Go. The most popular are VSCode (free), GoLand (paid), and Vim (free).
* A command terminal. Go works well using any terminal on Linux and Mac, and on PowerShell or cmd in Windows.

### Dependencies

* Go - <a href="https://go.dev/doc/install">Download and install</a> steps.

### Installing

```
$ git clone https://github.com/merci-app/oauth-sample-api-golang
```

### Configuring

Set this global variables
```go
var (
    username = "username"
    password = "password"
)
```


### Executing program

```
cd cmd
$ go run main.go
```

## Authors

- Experiences team

## Version History

* 0.1
    * Initial Release
