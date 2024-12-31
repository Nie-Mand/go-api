#!/bin/sh

# Since we're using Taskfile
TOOL="task utils:rename-project --"

DEFAULT_PROJECT_NAME=github.com/Nie-Mand/go-api
PROJECT_NAME=$1

function escape() {
    echo $1 | sed 's/\./\\./g' | sed 's/\//\\\//g'
}

if [ -z "$PROJECT_NAME" ]; then
    echo "Usage: $TOOL <project-name>"
    echo "Example: $TOOL github.com/Nie-Mand/go-api"
    exit 0
fi

ESCAPED_DEFAULT_PROJECT_NAME=$(escape $DEFAULT_PROJECT_NAME)
ESCAPED_PROJECT_NAME=$(escape $PROJECT_NAME)

find . -type f -not -path "./.git/*" -exec sed -i '' "s/$ESCAPED_DEFAULT_PROJECT_NAME/$ESCAPED_PROJECT_NAME/g" {} +
echo "Done"
