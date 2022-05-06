#!/bin/bash

if [[ ! -f .env ]]
then
  cp .env.example .env
  printf "\033[0;30m\033[42m Copy file '.env' SUCCEEDED \033[0m\n"
fi

if [[ ! -f files/etc/configuration/main.development.ini ]]
then
  cp files/etc/configuration/main.development.example.ini files/etc/configuration/main.development.ini
  printf "\033[0;30m\033[42m Copy file 'main.development.ini' SUCCEEDED \033[0m\n"
fi

if [[ ! -f files/etc/configuration/toggle.development.ini ]]
then
  cp files/etc/configuration/toggle.development.example.ini files/etc/configuration/toggle.development.ini
  printf "\033[0;30m\033[42m Copy file 'toggle.development.ini' SUCCEEDED \033[0m\n"
fi

