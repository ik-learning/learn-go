#!/bin/bash

read X
read Y
read Z

if (( $X == $Y )) && (( $X == $Z )) && (( $Y == $Z )); then echo "EQUILATERAL"; exit 0; fi
if (( $X != $Y )) && (( $X != $Z )) && (( $Y != $Z )); then echo "SCALENE"; exit 0; fi
if (( $X == $Y )) || (( $Y == $Z )) || (( $X == $Z )); then echo "ISOSCELES"; exit 0; fi