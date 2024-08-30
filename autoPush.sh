#!/bin/bash

bg_blue='\033[44m'
bg_green='\033[42m'
reset_color='\033[0m'

echo "+--------------------+"

echo -e "${bg_blue}==>Start adding...${reset_color}"

git add .

echo -e "${bg_green}==>Add complete${reset_color}"

echo "+--------------------+"

echo -e "${bg_blue}==>Start committing...${reset_color}"

git commit -m "update via git"

echo -e "${bg_green}==>Commit complete${reset_color}"

echo "+--------------------+"

echo -e "${bg_blue}==>Start pushing...${reset_color}"

git push

echo -e "${bg_green}==>Push complete${reset_color}"

echo "+--------------------+"