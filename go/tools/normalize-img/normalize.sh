#!/bin/bash

(
    cd "$(cd $(dirname $0); pwd)"
    go run . -q 90 cat1
    go run . -q 90 cat2
    go run . -q 90 cat3
    go run . -q 90 catp
)