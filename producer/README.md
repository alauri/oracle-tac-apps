# The UDP producer

To make this test as realistic as possible, we have simulated a real streaming
of data.

In this folder you can find an UDP server written in Golang which sends all the
events (or just some of them) to any UDP client connected to it.

> [!NOTE]
> This project does not support Go modules, so the env **GO111MODULE** must be
> specified.


## Setup

    $ GO111MODULE=off go get github.com/pterm/pterm


## Build

    $ GO111MODULE=off go build main.go


## Execution

The server is a small CLI with a few information needed, just have a look at the
*usage* message:

    $ GO111MODULE=off go run main.go -h
