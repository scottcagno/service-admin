#!/usr/bin/env bash

# check to make sure tmux *IS* running
tmls=$(pgrep tmux)
if [[ "$tmls" != "" ]]; then
  tmux kill-session -t aware
else
  echo "cannot stop tmux, it is not running."
fi

# check to make sure the process *IS* running
pid=$(pgrep java)
if [[ "$pid" != "" ]]; then
  output=$(../../internal/bash/java-shutdown.sh)
  if [[ "$output" == *"is now shutdown"* ]]; then
    echo "successfully shutdown"
  else
    echo "unable to shutdown!!"
  fi
else
  echo "cannot stop process, it is not running."
fi

# check for and remove started file
#if [[ -f "started.txt" ]]; then
#  echo "found started.txt file, removing"
#  rm started.txt
#fi
