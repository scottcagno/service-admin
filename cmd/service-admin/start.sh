#!/usr/bin/env bash

tmls=$(pgrep tmux)
if [[ "$tmls" == "" ]]; then
  tmux -f conf/tmux.conf attach -t aware
  clear
  sleep 15
else
  echo "cannot start tmux, it is already running."
fi

p_tmux=$(./info.sh | sed -n '1p')
p_java=$(./info.sh | sed -n '2p')

if [[ $p_tmux == *"not"* && $p_tmux == *"not"* ]]; then
  echo "tmux and java are not running"
else
  echo "success"
fi
