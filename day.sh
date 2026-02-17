#!/usr/bin/env bash
set -e

LANG="go"
YEAR=$1
DAY=$2
shift 2

while [[ $# -gt 0 ]]; do
    case "$1" in
        --python)
            LANG="python"
            ;;
        --go)
            LANG="go"
            ;;
        --zig)
            LANG="zig"
            ;;
        *)
            echo "No template for language: $1"
            exit 1
            ;;
    esac
    shift
done

if [ -z "$YEAR" ] || [ -z "$DAY" ]; then
    echo "Usage: ./day.sh <year> <day> [--python|--go]"
    exit 1
fi
if [ "$DAY" -lt 1 ] || [ "$DAY" -gt 25 ]; then
    echo "Day should be between 1 and 25"
    exit 1
fi
if [ "$YEAR" -lt 2015 ]; then
    echo "Year should be greater than or equal to 2015"
    exit 1
fi

SCRIPT_DIR="$(dirname "$(realpath "$0")")"
DAY_PADDED=$(printf "%02d" "$DAY")
DAY_DIR="$SCRIPT_DIR/$YEAR/day$DAY_PADDED"

mkdir -p "$DAY_DIR"

TEMPLATE="$SCRIPT_DIR/templates/template.$LANG"
TARGET="$DAY_DIR/main.$LANG"

if [ ! -f "$TARGET" ]; then
    cp "$TEMPLATE" "$TARGET"
fi
if [ ! -f "$DAY_DIR/input.txt" ]; then
    touch "$DAY_DIR/input.txt"
fi
if [ ! -f "$DAY_DIR/sample.txt" ]; then
    echo "Filled in with sample input for each problem" > "$DAY_DIR/sample.txt"
fi

if [ ! -f "$SCRIPT_DIR/.env" ]; then
    echo "No .env file found"
    exit 1
fi
source "$SCRIPT_DIR/.env"

if [ -z "$AOC_SESSION_COOKIE" ]; then
    echo "AOC_SESSION_COOKIE is not set in env file, unable to pull input"
    exit 1
fi

curl -s -b "session=$AOC_SESSION_COOKIE" \
    "https://adventofcode.com/$YEAR/day/$DAY/input" \
    > "$DAY_DIR/input.txt"

