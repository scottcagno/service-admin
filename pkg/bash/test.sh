#!/usr/bin/env bash

function exists() {
  [[ -f "$1" ]]
}

if exists "cmd/.keepz"; then
  echo "file exists"
else
  echo "file does not exist"
fi
