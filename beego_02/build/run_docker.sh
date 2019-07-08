#!/bin/bash
docker run --rm -v "$PWD/..":/go/src/a_beego/ -p8080:8080 a_beego