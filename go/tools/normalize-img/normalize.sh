#!/bin/bash

(
    cd "$(cd $(dirname $0); pwd)"
    go run . -q 100 cat1
    go run . -q 100 cat2
    go run . -q 100 cat3
    go run . -q 100 catp
)