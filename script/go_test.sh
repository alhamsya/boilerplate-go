#!/bin/bash

PASS=true

go clean -testcache -race ; files=`go list ./... | grep -v mocks` ;

printf "\nRUN GO VET\n" ;
go vet $files ;
if [[ $? != 0 ]]; then
    PASS=false
fi

printf "\nRUN GO TEST START\n" ;
go test $files -cover
if [[ $? != 0 ]]; then
  PASS=false
fi
printf "\nRUN GO TEST END\n\n" ;

if ! $PASS; then
  printf "\033[0;30m\033[41mFAILED\033[0m\n"
  exit 1
else
  printf "\033[0;30m\033[42mSUCCEEDED\033[0m\n"
fi

exit 0