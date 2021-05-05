#!/usr/bin/env bash

function stop_tmux() {
  echo "attempting to stop tmux..."
  tmux kill-session -t aware
}

function stop_process() {
  echo "attempting to stop process..."
  #output=`~/AwareIM/bin/shutdown.sh`
  output=$(../../internal/bash/java-shutdown.sh)
  if [[ "$output" == *"is now shutdown"* ]]; then
    echo "successfully shutdown"
  else
    echo "unable to shutdown!!"
  fi
}

# check to make sure tmux *IS* running
tmls=$(pgrep tmux)
if [[ "$tmls" != "" ]]; then
  stop_tmux
else
  echo "cannot stop tmux, it is not running."
fi

# check to make sure the process *IS* running
pid=$(pgrep java)
if [[ "$pid" != "" ]]; then
  stop_process
else
  echo "cannot stop process, it is not running."
fi

# check for and remove started file
if [[ -f "started.txt" ]]; then
  echo "found started.txt file, removing"
  rm started.txt
fi
