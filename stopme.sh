#!/usr/bin/env bash
# stopme.sh is a test script meant to mimic "stopping" a process
FILE=fakeproc
[[ -f "$FILE" ]] && rm "$FILE" # if the file exists, remove it
