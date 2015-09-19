#!/bin/bash

export MAJOR_VERSION=0
export MINOR_VERSION=1
export PATCH_VERSION=0
export BUILD_VERSION=$TRAVIS_BUILD_ID

export PACKAGE_NAME=swagger_$MAJOR_VERSION.$MINOR_VERSION.$PATCH_VERSION

# Create directories for the deb package
mkdir $PACKAGE_NAME
mkdir $PACKAGE_NAME/DEBIAN
mkdir $PACKAGE_NAME/etc
mkdir -p $PACKAGE_NAME/usr/local/bin

# Copy configuration file to /etc
cp $TRAVIS_BUILD_DIR/swell.conf $PACKAGE_NAME/etc
# Copy new binary to /usr/local/bin
cp $HOME/gopath/bin/swell $PACKAGE_NAME/usr/local/bin

cat << EOF > $PACKAGE_NAME/DEBIAN/control
Package: swagger
Version: $MAJOR_VERSION.$MINOR_VERSION.$PATCH_VERSION-$BUILD_VERSION
Section: utils
Priority: optional
Architecture: amd64
Maintainer: Luke Segars <luke@lukesegars.com>
Description: Simple alternative to popular automated
 configuration management tools. 
EOF

dpkg-deb --build $PACKAGE_NAME