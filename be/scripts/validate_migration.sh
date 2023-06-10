#!/bin/bash
if [[ -z $(which migrate) ]]; then
  echo "Error: gomigrate is not installed";
  exit 1;
fi