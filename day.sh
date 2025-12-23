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
        *)
            echo "Unknown option: $1"
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
day_dir="$SCRIPT_DIR/$YEAR/day$DAY_PADDED"

mkdir -p "$day_dir"

case "$LANG" in
    go)
        TEMPLATE="$SCRIPT_DIR/templates/template.go"
        TARGET="$day_dir/main.go"
        ;;
    python)
        TEMPLATE="$SCRIPT_DIR/templates/template.py"
        TARGET="$day_dir/main.py"
        ;;
    *)
        echo "Unsupported language: $LANG"
        exit 1
        ;;
esac

if [ ! -f "$TARGET" ]; then
    cp "$TEMPLATE" "$TARGET"
fi
if [ ! -f "$day_dir/input.txt" ]; then
    touch "$day_dir/input.txt"
fi
if [ ! -f "$day_dir/sample.txt" ]; then
    echo "Filled in with sample input for each problem" > "$day_dir/sample.txt"
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
    > "$day_dir/input.txt"

