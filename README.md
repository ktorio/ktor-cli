# Ktor CLI
Ktor CLI is a command-line tool that brings the capability to create, run, and interact with your [Ktor](https://ktor.io/) application from the command line.

## Install CLI tools

### macOS

You can install Ktor CLI on macOS using [Homebrew](https://brew.sh/) as follows:
1. Add a Ktor repository using the `brew tap` command:
   ```
   brew tap ktorio/ktor\
   ```
2. Install Ktor CLI using `brew install`:
   ```
   brew install --build-from-source ktor\
   ```

### Linux

On Linux, you can install Ktor CLI using [snaps](https://snapcraft.io/):

```
snap install --beta --classic ktor
```

### Available commands
You can get a list of Ktor commands available to you by typing `ktor --help`.


## Create a project

To create a new Ktor project, pass a project name to the `ktor generate` command:

```
ktor generate ktor-sample
```

This command generates a simple Ktor project that uses the Gradle build system with Kotlin DSL.


## Run a project

To run the existing Ktor application, use the `ktor start` command.
This command accepts the name of the directory where the project is placed:

```
ktor start ktor-sample
```

With the default configuration, the terminal should show the following message:

```
[main] INFO  ktor.application - Responding at http://0.0.0.0:8080
```

This means that the server is ready to accept requests at the http://0.0.0.0:8080 address. 
