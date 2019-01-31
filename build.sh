#!/bin/bash
if [ ! $1 ]
  then
  echo "input build or restart"
  exit
fi

if [ $1 == "mac" ]
then
{
  CGO_ENABLED=0 GOOS=darwin GOARCH=amd64  go build --tags "libsqlite3 linux" -o dist/main-mac
  echo "------>   build ok"
}||{
  echo "build error"
}
fi

if [ $1 == "linux64" ]
then
{
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build --tags "libsqlite3 linux" -o dist/main-linux
  echo "------>   build ok"
}||{
  echo "build error"
}
fi

if [ $1 == "linux386" ]
then
{
  CGO_ENABLED=0 GOOS=linux GOARCH=386 go build --tags "libsqlite3 linux" -o dist/main-linux
  echo "------>   build ok"
}||{
  echo "build error"
}
fi
