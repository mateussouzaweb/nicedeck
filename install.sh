#!/bin/bash

set -e
set -u

VERSION="v0.0.1"
REPOSITORY="https://github.com/mateussouzaweb/nicedeck"
BINARY="$REPOSITORY/releases/download/$VERSION/nicedeck"
BINARIES="$HOME/.local/bin"

# Make sure binaries folder exists
if [ -d "$BINARIES" ]; then
  mkdir -p $BINARIES
fi

# Make sure binaries path works
export PATH="$PATH:$BINARIES"

# Install nicedeck
echo "[INFO] Downloading nicedeck..."
wget -q $BINARY -O $BINARIES/nicedeck
chmod +x $BINARIES/nicedeck

echo "[INFO] nicedeck ${VERSION} installed at $BINARIES/nicedeck"