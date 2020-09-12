#!/bin/bash

fn(){
   
   DIR="$1"
   echo "-------->-""$DIR""----"
   cd $DIR
   git pull origin master

};





fn   "$WORK/yi/";
fn   "$WORK/yiintf/";

fn   "$WORK/yigw/";

fn   "$WORK/yiuser/";
fn   "$WORK/yicms/";
fn   "$WORK/vchat/";
#fn   "$GOPATH/src/yitrade/";
#fn   "$GOPATH/src/yimsg/";
#fn   "$GOPATH/src/yiutil/";



echo "------complete------------------"
