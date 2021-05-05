#!/usr/bin/env bash

function disable_ufw() {
  echo "attempting to disable ufw..."
  #echo $pas | sudo -S ufw disable
}

function start_process() {
  echo "attempting to start process..."
  # try to start it up
  #path="~/AwareIM/bin/"
  #prog="startAwareIMNoGUI.sh"
  path="../../internal/bash"
  prog="java-startup.sh"
  cd $path
  bash $prog
}

function enable_ufw() {
  echo "attempting to enable ufw..."
  #echo $pas | sudo -S ufw enable
}

# make sure to shutdown ufw while the process starts up
disable_ufw

# make sure the process is *NOT* running
pid=$(pgrep java)
if [[ "$pid" == "" ]]; then
  start_process &
else
  echo "cannot start process, it is already running."
fi

# make sure to enable ufw again after the process starts up
enable_ufw

# write status
#now=$(date +"%m-%d-%Y %H:%M")
#echo "$now" >started.txt

# detach from tmux session
tmux detach -s aware
