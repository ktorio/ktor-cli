# Ktor CLI
The `ktor` tool allows generating [Ktor](https://ktor.io/) applications through the command line interface.

## Prerequisites
To build the tool, the `go` compiler needs to be installed first. You can find the [installation guide](https://go.dev/doc/install) on the official website.


## Building
To build an executable, issue the following command in the root directory of the repository:
```shell
go build github.com/ktorio/ktor-cli/cmd/ktor
```

If the build is successful, the `ktor` executable should appear in the current directory.
Also, the `go` command can be issued through Docker using an [official Go image](https://hub.docker.com/_/golang):
```shell
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.21 go build -buildvcs=false -v github.com/ktorio/ktor-cli/cmd/ktor
```

## Running
To run the tool without making an intermediate build, execute the following command:
```shell
go run github.com/ktorio/ktor-cli/cmd/ktor # followed by CLI args
```

Effectively, the `go run github.com/ktorio/ktor-cli/cmd/ktor` line can replace the `ktor` executable in the below commands.


## Create a project

To create a new Ktor project, pass a project name to the `ktor new` command:

```
ktor new ktor-sample
```

The `-v` option can be used to enable verbose output:
```shell
ktor -v new ktor-project
```

## Get the version

To get the version of the tool, use the `--version` flag or the `version` command:
```shell
ktor --version
ktor version
```

## Get the usage info

To get the help page about the tool usage, use the `--help` flag or the `help` command:
```shell
ktor --help
ktor help
```

## HTTP proxy

To use a proxy server while making requests to the generation server, set the `HTTPS_PROXY` environment variable. Here is an example:
```shell
HTTPS_PROXY=http://localhost:3128 ktor new ktor-project
```