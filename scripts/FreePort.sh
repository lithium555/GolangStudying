#!/bin/bash



function free_port(){


# Free port - kill pid that use port
  echo "Gonna to free port ${PORT}..."
  PID=$(lsof -i :"${PORT}" | awk '{if ($2 != "PID")print $2}' | uniq)
  if [[  "${PID}" != "" ]];
  then
    echo "To free the port enter sudo password"
    sudo kill -9 "${PID}"
  fi
  echo Port "${PORT}" is free
}


if [ $# -eq 0 ]
then
  echo "Free_Port.sh is used as imported"
else
  PORT=$1
  echo port is "${PORT}"
  free_port
fi