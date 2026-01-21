#!/bin/bash 

OS_ARCH="$(uname -m)" 
APP_DIR="$(dirname "$0")" 
    
if [ "$OS_ARCH" == "x86_64" ]; then 
    exec "$APP_DIR/nicedeck-amd64" "$@" 
elif [ "$OS_ARCH" == "arm64" ]; then 
    exec "$APP_DIR/nicedeck-arm64" "$@" 
fi