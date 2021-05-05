#!/usr/bin/env bash

p_tmux="tmux is not running"
p_java="java is not running"

id_tmux=$(pgrep tmux)
if [[ "$id_tmux" != "" ]]; then
  p_tmux="tmux is running"
fi

id_java=$(pgrep java)
if [[ "$id_java" != "" ]]; then
  p_java="java is running"
fi

echo "$p_tmux"
echo "$p_java"
