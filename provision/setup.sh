#!/bin/bash

apt-get update
apt-get -y install git
apt-get -y install golang

go get github.com/gorilla/mux

