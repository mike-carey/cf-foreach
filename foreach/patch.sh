#!/usr/bin/env bash

###
# This file modifies the generated file from ifacemaker to import cfclient and modify the classes that needed to be referenced  from that package.
##

function patch() {
  keep_backup=false
  args=()
  while [[ -n "$1" ]]; do
    case $1 in
      --keep-backup )
        keep_backup=true
        ;;
      -- )
        args+=($@)
        break
        ;;
      * )
        args+=("$1")
        ;;
    esac
    shift
  done

  for file in *.go; do
    if [[ $file == definition.go ]]; then
      continue
    fi

    sed -i.bak 's/^func ForEachCfclient\(.*\)ToCfclient\(.*\)/func \1To\2/g' $file
    if [[ $keep_backup != true ]]; then
      rm -f $file.bak
    fi
  done
}

if [[ ${BASH_SOURCE[0]} != $0 ]]; then
  export -f patch
else
  set -euo pipefail

  patch "${@:-}"
  exit $?
fi
