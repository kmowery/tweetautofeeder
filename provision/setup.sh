#!/bin/bash

apt-get update
apt-get -y install git

# preconfigure and install golang
echo "golang-go golang-go/dashboard boolean false" | debconf-set-selections
apt-get -y install golang

go get github.com/gorilla/mux

cp /vagrant/home/bashrc /home/vagrant/.bashrc

