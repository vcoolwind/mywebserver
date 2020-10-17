#!/bin/bash
go build -o mywebserver && sudo docker build . -t yanfengking/mywebserver:1.1 && sudo docker push yanfengking/mywebserver:1.1 
