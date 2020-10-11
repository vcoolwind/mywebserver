#!/bin/bash
go build -o mywebserver && sudo docker build . -t yanfengking/mywebserver:1.0
