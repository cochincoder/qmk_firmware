#!/bin/sh
# NOTE: This script uses tabs for indentation

errcho() {
  echo "$@" >&2
}

USAGE="Usage: $0 [keyboard[:keymap[:target]]]"

# Check preconditions
for arg; do
  if [ "$arg" = "--help" ]; then
    echo "$USAGE"
    exit 0
  fi
done
if [ $# -gt 1 ]; then
  errcho "$USAGE"
  exit 1
fi

# Allow $RUNTIME to be overriden by the user as an environment variable
# Else check if either docker or podman exit and set them as runtime
# if none are found error out
if [ -z "$RUNTIME" ]; then
  if command -v docker >/dev/null 2>&1; then
    RUNTIME="docker"
  elif command -v podman >/dev/null 2>&1; then
    RUNTIME="podman"
  else
    errcho "Error: no compatible container runtime found."
    errcho "Either podman or docker are required."
    errcho "See https://podman.io/getting-started/installation"
    errcho "or https://docs.docker.com/install/#supported-platforms"
    errcho "for installation instructions."
    exit 2
  fi
fi

usb_args="--privileged -v /dev:/dev"
dir=$(pwd -W 2>/dev/null) || dir=$PWD  # Use Windows path if on Windows

if [ "$RUNTIME" = "docker" ]; then
  uid_arg="--user $(id -u):$(id -g)"
fi

# Run container and build firmware
"$RUNTIME" run --rm -it $usb_args \
  $uid_arg \
  -w /qmk_firmware \
  -v "$dir":/qmk_firmware \
  -e ALT_GET_KEYBOARDS=true \
  -e SKIP_GIT="$SKIP_GIT" \
  -e MAKEFLAGS="$MAKEFLAGS" \
  qmkfm/qmk_cli \
  bash
