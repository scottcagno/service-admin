#!/usr/bin/env bash

function start_tmux() {
  echo "attempting to start tmux..."
  tmux new-session -d -s aware
  tmux send-keys './start-everything.sh' 'C-m'
  #tmux attach -t aware
  #tmux detach -s aware
}

# make sure tmux is *NOT* running
tmls=$(pgrep tmux)
if [[ "$tmls" == "" ]]; then
  start_tmux
  sleep 10
else
  echo "cannot start tmux, it is already running."
fi

# get timestamp and check to make sure everything was good
now=$(date +"%m-%d-%Y %H:%M")
out=$(cat started.txt)
if [[ "$now" == "$out" ]]; then
  echo "SUCCESS!!"
else
  echo "SOMETHING WENT WRONG!"
fi
