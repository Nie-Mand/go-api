#!/bin/sh

# Since we're using Taskfile
TOOL="task utils:rename-project --"

DEFAULT_PROJECT_NAME=github.com/Nie-Mand/go-api
PROJECT_NAME=$1

if [ -z "$PROJECT_NAME" ]; then
    echo "Usage: $TOOL <project-name>"
    echo "Example: $TOOL github.com/Nie-Mand/go-api"
    exit 0
fi

# TODO: Replace all occurrences of the default project name with the new project name
# find . -path ./.git -prune -o -type f -exec sed -i "s,$DEFAULT_PROJECT_NAME,$PROJECT_NAME,g" {} +
# -exec sed -i "s,$DEFAULT_PROJECT_NAME,$PROJECT_NAME,g" {} +

echo "Done"
