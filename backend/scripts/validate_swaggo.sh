#!/bin/bash

if ! ls "$(go env GOPATH)"/bin/swag 1>/dev/null 2>/dev/null; then
  echo "There's no swagger binary";
  echo "Processing to download the swagger";

  if ! go install github.com/swaggo/swag/cmd/swag@v1.8.12; then 
    echo "The download exit with error code $?"
  fi
fi