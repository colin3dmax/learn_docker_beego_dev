#!/bin/bash
docker run --rm -v "$PWD/..":/go/src/beego_03/ -p8080:8080 beego_03