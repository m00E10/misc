#!/bin/bash
# Positional args: break time, work time, extended break

chill() {
break_time=$(( break_time * 60 ))
date0=$((`date +%s` + $break_time))
ffplay ~/bin/apb-intro.mp3 -nodisp -nostats -hide_banner &
while [ "$date0" -ge `date +%s` ]; do
    sleep 1; clear; 
	echo "Time to Chill" | figlet -f slant
	echo -ne "$(date -u --date @$(($date0 - `date +%s` )) +%H:%M:%S)\r" | figlet -f shadow
done
}

work() {
work_time=$(( work_time * 60 ))
date1=$((`date +%s` + $work_time)) 
ffplay ~/bin/apb-intro.mp3 -nodisp -nostats -hide_banner &
while [ "$date1" -ge `date +%s` ]; do 
    sleep 1; clear;
	echo "Time to Grind" | figlet
    echo -ne "$(date -u --date @$(($date1 - `date +%s` )) +%H:%M:%S)\r" | figlet
done
}

# One chill session before work to setup study environment. This setup period
# benefits from being timed, it prepares the mind for work and ensures I won't
# spend too long trying to find a particular study mix or w/e
break_time=${1:-5}
chill

while true; do
    counter=0
    while [ $counter -lt 4 ]; do
        break_time=${1:-5}
        work_time=${2:-25}
	    work
    	chill
        let "counter++"
    done
    break_time=${3:-30}
    chill
done
