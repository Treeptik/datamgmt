# datamgmt
Cloudunit logging management agent

# Contributing

## Setting up your dev environment

Current Go version used for development is Golang 1.7.4.

install version 0.11.0-2

The location where you clone is important. Please clone under the source
directory of your `GOPATH`. If you don't have `GOPATH` already set, you can
simply set it to your home directory (`export GOPATH=$HOME`).

    $ mkdir -p ${GOPATH}/src/github.com/treeptik
    $ cd ${GOPATH}/src/github.com/treeptik
    $ git clone https://github.com/treeptik/datamgmt.git
    $ glide up
