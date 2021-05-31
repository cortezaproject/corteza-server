#!/bin/sh

BRANCH_NAME=$(git symbolic-ref -q HEAD | sed 's/refs\/heads\///g')

# TODO: Actual branch name rules
if  [[ ! $BRANCH_NAME =~ ^2021.6.x-.*|^2021.9.x-.* ]];
then
    echo "Wrong branch name: $BRANCH_NAME"
    exit 1
fi
