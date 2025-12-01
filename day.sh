#!/usr/bin/env bash

# Get input
if [ -z "$1" ]; then
    echo "Usage: ./day.sh <day_number> <optional year>"
    exit 1
fi

DAY=$1
YEAR=${2:-$(date +'%Y')}

# Validate input
if [ $DAY -lt 1 ] || [ $DAY -gt 25 ]; then
    echo "Day should be between 1 and 25"
    exit 1
fi
if [ $YEAR -lt 2015 ]; then
    echo "Year should be greater than or equal to 2015"
    exit 1
fi

# Zero-pad the day for the directory
DAY_PADDED=$(printf "%02d" $DAY)

# Make dir & files
day_dir=$(dirname $(realpath $0))/$YEAR/day$DAY_PADDED
if [ ! -d $day_dir ]; then
    mkdir -p $day_dir
fi

# Create files if they don't exist
if [ ! -f $day_dir/main.go ]; then
    cp $(dirname $(realpath $0))/template.go $day_dir/main.go
fi
if [ ! -f $day_dir/input.txt ]; then
    echo "Filled in automatically with actual problem input" > $day_dir/sample.txt
fi
if [ ! -f $day_dir/input.txt ]; then
    echo "Filled in with sample input for each problem" > $day_dir/sample.txt
fi

# Validate & curl from adventofcode website
if [ ! -f $(dirname $(realpath $0))/.env ]; then
    echo "No .env file found"
    exit 1
fi
source $(dirname $(realpath $0))/.env
if [ -z "$AOC_SESSION_COOKIE" ]; then
    echo "AOC_SESSION_COOKIE is not set in env file, unable to pull input"
    exit 1
fi

curl -s -b "session=$AOC_SESSION_COOKIE" https://adventofcode.com/$YEAR/day/$DAY/input > $day_dir/input.txt

