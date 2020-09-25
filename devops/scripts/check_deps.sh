#!/bin/bash

sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install -y golang-go

sudo apt install -y zip