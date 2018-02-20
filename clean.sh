#!/bin/sh

_commands=$(ls cmd)
(
  cd $GOPATH/bin
  rm -f $_commands
)
