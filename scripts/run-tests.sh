#!/bin/bash

export BACKEND_PORT=3003
export MIGRATION=true
export DB_NAME=test.db
export RUN_ENV=test

PIDFILE="tmpfile-$LOGNAME.txt"
STARTDIR=$(pwd)

runTest() {
  cd test

  npm run test
}

runBackend() {
  cd backend && go run cmd/main.go
}

start() {
  runBackend &

  # Save the backgound job process number into a file.
  jobs -p >$PIDFILE

  # Disconnect the job from this shell.
  # (Note that 'disown' command is only in the 'bash' shell.)
  disown %1

  # Print a message indicating the script has been started
  echo "Script has been started..."

  sleep 1

  # stop and kill created processes
  # if tests failed exit with code 1
  if runTest; then
    stop

    killProcess
  else
    stop

    killProcess
    exit 1
  fi

  # Print a message indicating the script
  # has been done successfully
  echo -e "Script has been done...\n\n\n"
}

stop() {
  # Go to basedir
  cd $STARTDIR

  read PID <$PIDFILE

  # Remove the PIDFILE
  rm -f $PIDFILE

  # Send a 'terminate' signal to process
  kill $PID

  # Print a message indicating the script has been stopped
  echo -e "Everything stopped..."
}

killProcess() {
  # Kill bg process of backend
  kill -9 $(lsof -t -i:$BACKEND_PORT)

  # Remove test db
  # rm ./backend/$DB_NAME

  # Print a message indicating the script has been killed
  echo -e "Everything killed..."
}

start
