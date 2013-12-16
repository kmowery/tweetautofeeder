#!/bin/bash

apt-get update
apt-get -y install git

# preconfigure and install golang
echo "golang-go golang-go/dashboard boolean false" | debconf-set-selections
apt-get -y install golang

cp /vagrant/home/bashrc /home/vagrant/.bashrc
export GOPATH="/vagrant/golang"
mkdir -p ${GOPATH}

go get github.com/gorilla/mux
go get github.com/hoisie/mustache


