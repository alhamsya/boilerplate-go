#!/bin/bash

if [[ ! -f .env ]]
then
  cp .env.example .env
  printf "\033[0;30m\033[42m Copy file '.env' SUCCEEDED \033[0m\n"
fi

if [[ ! -f files/etc/configure/main.development.ini ]]
then
  cp files/etc/configure/main.development.example.ini files/etc/configure/main.development.ini
  printf "\033[0;30m\033[42m Copy file 'main.development.ini' SUCCEEDED \033[0m\n"
fi

if [[ ! -f files/etc/configure/toggle.development.ini ]]
then
  cp files/etc/configure/toggle.development.example.ini files/etc/configure/toggle.development.ini
  printf "\033[0;30m\033[42m Copy file 'toggle.development.ini' SUCCEEDED \033[0m\n"
fi

