# URL Shortener App

## Overview

An URL shortener app written in golang gin web framework. 

## How to run

```
$ git clone git@github.com:balajisa09/url-shortener.git

$ cd url-shortener

$ go run main.go 
```

## Run as a docker container

```
From the repository 

$ docker build -t url-shortener

$ docker run -p 8080:8080 url-shortener 

or 

From remote docker registry 

docker pull balajis30/url-shortener

docker run -p 8080:8080 balajis30/url-shortener 

```


The app will be running on localhost:8080. visit the browser and enter the url to be shortened.