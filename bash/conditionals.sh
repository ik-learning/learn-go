#!/bin/bash

read X

if [[ $X =~ ^[yY]{1}$ ]]; then echo "YES"; fi
if [[ $X =~ ^[nN]{1}$ ]]; then echo "NO"; fi