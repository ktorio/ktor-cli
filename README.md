# Ktor-cli
The tool allows generating [Ktor](https://ktor.io/) applications through the command line interface.

## Building
To build an executable issue the following command in the root directory of the repository:
```shell
go build github.com/ktorio/ktor-cli/cmd/ktor
```

If the build is successful the `ktor` executable should appear in the current directory.
Also, the `go` command can be issued through Docker:
```shell
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.21 go build -buildvcs=false -v github.com/ktorio/ktor-cli/cmd/ktor
```

## Running
To run the tool without an intermediate build, execute the following command:
```shell
go run github.com/ktorio/ktor-cli/cmd/ktor # followed by CLI args
```


## Create a project

To create a new Ktor project, pass a project name to the `ktor new` command:

```
ktor new ktor-sample
```

The `-v` option can be used to enable verbose output:
```shell
ktor -v new ktor-sample
```
