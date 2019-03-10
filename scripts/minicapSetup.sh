#!/usr/bin/env bash

minicap=/Volumes/SAMSUNG/Library/minicap

# Launching minicap on device $1

if [[ $1 ]]; then
    ANDROID_SERIAL=$(echo $1)
    res=$(cd $minicap && ./run.sh autosize)
    echo $res
fi

# Forwarding to local port 1313

adb forward tcp:1313 localabstract:minicap
