#!/bin/bash

apt-get update
apt-get -y install git

apt-get -y install python-software-properties
add-apt-repository -y ppa:duh/golang
apt-get update

# preconfigure and install golang
echo "golang-go golang-go/dashboard boolean false" | debconf-set-selections
apt-get -y install golang

apt-get -y install sqlite3

sqlite3 /home/vagrant/storage.sqlite < /vagrant/provision/make_db.sql
chown vagrant:vagrant /home/vagrant/storage.sqlite

cp /vagrant/provision/home/bashrc /home/vagrant/.bashrc
export GOPATH="/home/vagrant/gopath"
mkdir -p ${GOPATH}
chown -R vagrant:vagrant ${GOPATH}

go get github.com/gorilla/mux
go get github.com/gorilla/sessions
go get github.com/hoisie/mustache
go get github.com/ChimeraCoder/anaconda
go get github.com/mrjones/oauth
go get github.com/mattn/go-sqlite3


