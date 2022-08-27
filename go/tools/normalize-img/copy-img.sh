#!/bin/bash

file=normalized.jpg

(
    TOOLDIR=$(cd $(dirname $0); pwd)
    cd $TOOLDIR
    SRCDIR=$TOOLDIR/img
    DSTDIR=$(realpath ../../assets/img)
    rm $DSTDIR/* -rf
    find $SRCDIR -type f -name "$file" -exec dirname {} \; \
        | while read t ; do
        cp -p $t/$file $DSTDIR/$(basename $t).jpg
    done
    ls -l $DSTDIR
)