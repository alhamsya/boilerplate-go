#!/bin/bash

if [[ ! -f .env ]]
then
  cp .env.example .env
  printf "\033[0;30m\033[42m Copy file '.env' SUCCEEDED \033[0m\n"
fi

if [[ ! -f files/etc/service/main.development.ini ]]
then
  cp files/etc/service/main.development.example.ini files/etc/service/main.development.ini
  printf "\033[0;30m\033[42m Copy file 'main.development.ini' SUCCEEDED \033[0m\n"
fi

if [[ ! -f files/etc/service/toggle.development.ini ]]
then
  cp files/etc/service/toggle.development.example.ini files/etc/service/toggle.development.ini
  printf "\033[0;30m\033[42m Copy file 'toggle.development.ini' SUCCEEDED \033[0m\n"
fi

