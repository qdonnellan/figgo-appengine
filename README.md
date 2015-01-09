# Google App Engine (Go) on Docker/Fig
This is a simple demonstration of getting `go_appengine` up and running inside a local Fig (Docker) development environment. The purpose is to prove that Fig can be a suitable protoytping dev environment for Go apps on any machine without having to perform the `GOPATH` dance with appengine!

## Before you start
You will need [Fig](http://www.fig.sh/install.html) and [Docker](https://docs.docker.com/installation/#installation) installed on your machine.

## Running the Dev Environment
1. Download this repo
2. `fig build`
3. `fig up`

That's it! During the `fig build` step you should see Go, and Google Appengine download and install inside your docker container. Notice that the `Dockerfile` takes care of setting your `GOPATH` for you!