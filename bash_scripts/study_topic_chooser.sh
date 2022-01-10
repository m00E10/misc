#!/bin/bash

# From a defined file, output a random string that is surrounded in double
# quotes. Primarily used to have the computer pick random projects for myself to
# work on when I have time to kill, and don't know what to do with it.

file=~/file.txt


total=$(grep -o '".*"' $file | tr -d '"' | wc -l)
random=$(shuf -i 1-$total -n 1)
result=$(grep -o '".*"' $file | tr -d '"' | head -$random | tail -1)

echo "Why not work on"
echo "$result"
