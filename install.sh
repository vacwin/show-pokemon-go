#!/bin/sh

pkgname=pokemon-colorscripts-go

sudo mkdir -p "/usr/local/share/$pkgname"

sudo cp -rf colorscripts "/usr/local/share/$pkgname"
sudo cp pokemon.json "/usr/local/share/$pkgname"

go install -ldflags "-X main.PROGRAM_DIR=/usr/local/share/$pkgname" . 
