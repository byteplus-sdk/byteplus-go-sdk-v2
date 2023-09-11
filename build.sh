#!/bin/bash

ROOT=$(pwd)

function build(){
  flag=0
  for file in `ls $ROOT"/"$1`; do
       if [ ! -d $ROOT"/"$1"/"$file ]; then
         flag=1
         break
       fi
  done
  if [ $flag == 1 ];then
      echo 'go build '$ROOT'/'$1'/*.go'
      go build "$ROOT"/"$1"/*.go
  fi
  for file in `ls $ROOT"/"$1`; do
     if [  -d $ROOT"/"$1"/"$file ]; then
       build "$1"/"$file"
     fi
  done
}

build 'internal'
build 'private'
build 'byteplus'
build 'service'
#build 'example'