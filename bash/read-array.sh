#!/bin/bash

COUNTER=0
while read -r line;
do
 if [ $COUNTER == 3 ]
  then
   echo $line;
   exit 0
 fi
let COUNTER=COUNTER+1
done