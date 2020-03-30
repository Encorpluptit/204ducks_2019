#!/bin/bash -e
# Build Copyright (C) 2020 Damien Bernard
# This program comes with ABSOLUTELY NO WARRANTY.
# This is free software, and you are welcome to redistribute it
# under certain conditions; see LICENSE for details.

# for arg in $@; do
#     $arg
# done


echo "Hello $1"
time=$(date)
echo "::set-output name=time::$time"
