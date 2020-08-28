#!/bin/bash

up(){
   
   DIR="$1"
   echo "-------->-""$DIR""----"
   cd $DIR
   git add .
   git commit -m "auto commit"
   git push
};




up   "$GOPATH/src/yi/";
up   "$GOPATH/src/yiintf/";
up   "$GOPATH/src/yigw/";
up   "$GOPATH/src/yiuser/";
up   "$GOPATH/src/yicms/";
up   "$GOPATH/src/yitrade/";
up   "$GOPATH/src/yimsg/";





echo "$VER"
