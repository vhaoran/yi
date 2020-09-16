#!/bin/bash

up(){
   
   DIR="$1"
   echo "-------->-""$DIR""----"
   cd $DIR
   git add .
   git commit -m "bbb"
   git push
};




up   "$WORK/yi/";
up   "$WORK/yiintf/";

up   "$WORK/yigw/";

up   "$WORK/yiuser/";
up   "$WORK/yicms/";
# up   "$WORK/yitrade/";
# up   "$WORK/yimsg/";
# up   "$WORK/yiutil/";






echo "$VER"
