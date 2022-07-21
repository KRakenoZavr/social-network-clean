#!/bin/bash

export BACKEND_PORT=3003
export MIGRATION=true
export DB_NAME=test.db
export RUN_ENV=test

# (cd backend && go run cmd/main.go &
# cd test && npm i && npm run test && kill -9 `lsof -t -i:3003` && rm ../backend/test.db)

PIDFILE="tmpfile-$LOGNAME.txt"
STARTDIR=$(pwd)

runTest() {
  cd test && npm i && npm run test
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

  if runTest; then
    stop && killProcess
  else
    stop && killProcess
    exit 1
  fi

  # Print a message indicating the script has been stopped
  echo -e "everything killed...\n\n\n"
}

stop() {
  cd $STARTDIR

  read PID <$PIDFILE

  # Remove the PIDFILE
  rm -f $PIDFILE

  # Send a 'terminate' signal to process
  kill $PID

  # Print a message indicating the script has been stopped
  echo -e "everything stopped...\n\n\n"
}

killProcess() {
  kill -9 $(lsof -t -i:$BACKEND_PORT) && rm /backend/$DB_NAME
}

start
