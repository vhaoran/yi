#!/bin/bash

fn(){
   
   DIR="$1"
   echo "-------->-""$DIR""----"
   cd $DIR
   git pull origin master

};





fn   "$GOPATH/src/yi/";
fn   "$GOPATH/src/yiintf/";
fn   "$GOPATH/src/yiuser/";
fn   "$GOPATH/src/yicms/";
fn   "$GOPATH/src/yitrade/";


echo "------complete------------------"
