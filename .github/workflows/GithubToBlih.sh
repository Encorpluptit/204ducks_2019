#!/bin/bash -e
# GitHubToBlih Copyright (C) 2020 Maxime Houis
# This program comes with ABSOLUTELY NO WARRANTY.
# This is free software, and you are welcome to redistribute it
# under certain conditions; see LICENSE for details.

if [[ $# == 4 ]]; then
    USER_GHUB="$1"
    REPO_GHUB="$2"
    USER_BLIH="$3"
    REPO_BLIH="$4"
elif [[ $# == 1 ]]; then
    USER_GHUB="$USER_GHUB"
    USER_BLIH="$USER_BLIH"
    REPO_BLIH="$1"
    REPO_GHUB=$REPO_GHUB
else
    echo "Usage: $0 REPO_NAME"
    exit 1
fi


set -xe
git push --mirror git@git.epitech.eu:$USER_BLIH/$REPO_BLIH


#TMP_DIR="./temporary_repository"
#set -xe
#OLDPWD=$(pwd)

#rm -rf $TMP_DIR
#mkdir $TMP_DIR

#git clone --bare git@github.com:$USER_GHUB/$REPO_GHUB $TMP_DIR
#cd $TMP_DIR
#git push --mirror git@git.epitech.eu:$USER_BLIH/$REPO_BLIH
#
#cd $OLDPWD
#rm -rf $TMP_DIR
