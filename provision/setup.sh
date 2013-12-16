#!/bin/bash

apt-get update
apt-get -y install git

# preconfigure and install golang
echo "golang-go golang-go/dashboard boolean false" | debconf-set-selections
apt-get -y install golang

cp /vagrant/provision/home/bashrc /home/vagrant/.bashrc
export GOPATH="/home/vagrant/gopath"
mkdir -p ${GOPATH}
chown -R vagrant:vagrant ${GOPATH}

go get github.com/gorilla/mux
go get github.com/hoisie/mustache


