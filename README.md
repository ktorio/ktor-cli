# Ktor CLI
The `ktor` tool allows generating [Ktor](https://ktor.io/) applications through the command line interface.

For a web interface, visit https://start.ktor.io.

## Install

### Linux and macOS

The tool can be installed via Homebrew:
```shell
brew install ktor
```

### Windows
The tool can be installed via WinGet:
```shell
winget install JetBrains.KtorCLI
```

## Prerequisites
To build the tool, the `go` compiler needs to be installed first. You can find the [installation guide](https://go.dev/doc/install) on the official website.


## Build
To build an executable, issue the following command in the root directory of the repository:
```shell
go build github.com/ktorio/ktor-cli/cmd/ktor
```

If the build is successful, the `ktor` executable should appear in the current directory.
Also, the `go` command can be issued through Docker using an [official Go image](https://hub.docker.com/_/golang):
```shell
docker run --rm -v "$PWD":/usr/src/build -w /usr/src/build golang:1.21 git config --global --add safe.directory . && go build -v github.com/ktorio/ktor-cli/cmd/ktor
```

## Run
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

## Create a project in an interactive mode

To create a new project in the interactive mode, simply use the `new` command without a project name:

```shell
ktor new
```

## Generate a project from an OpenAPI specification

To generate a project in the current directory from a given [OpenAPI specification](https://swagger.io/specification/), use the `openapi` command:
```shell
ktor openapi petstore.yaml
```

You can specify a different output directory with the `-o` or `--output` flag:
```shell
ktor openapi -o path/to/project petstore.yaml
```

## Add a Ktor dependency to an existing project

To add a Ktor dependency to a Gradle project in the current working directory, use the `add` command:
```shell
ktor add server-core
```
Use the `-p` or `--project` option to specify a path to the project directory:
```shell
ktor add -p /path/to/project server-core
```

You can add multiple modules with a single command:
```shell
ktor add -p /path/to/project server-core client-core json
```

Currently, Ktor dependencies can only be added to **non-multiplatform** Gradle projects using Kotlin DSL.

## Run a Ktor project in development mode

The `dev` command executes the `run` Gradle task while continuously rebuilding source files upon changes.
This allows uninterrupted development of a Ktor server application without requiring a restart to observe recent changes.
For the `dev` command to function correctly, ensure the Ktor Gradle plugin is applied to your project.

To run the application in development mode from a project in the current working directory, use the following command:
```shell
ktor dev
```

To specify a path to the project directory, use the `-p` or `--project` option:
```shell
ktor dev --project /path/to/project
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