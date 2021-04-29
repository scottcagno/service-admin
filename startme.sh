#!/usr/bin/env bash
# startme.sh is a test script meant to mimic "starting" a process
FILE=fakeproc
[[ ! -f "$FILE" ]] && { # if the file does not exist, create it
  touch "$FILE"
  echo "12345" >"$FILE"
}
