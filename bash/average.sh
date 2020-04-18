#!/bin/bash

COUNTER=0
SUM=0

read first
while read -r line;
do
  if (( $COUNTER <= $first ));
   then
    echo $line;
     let SUM=$SUM+$line;
     let COUNTER=COUNTER+1;
  fi
done

div="$(($SUM/$COUNTER))"
printf "%.3f" `(echo $div | bc -l)`