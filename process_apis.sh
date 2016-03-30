#!/usr/bin/env bash

if [ $# -lt 1 ]; then
  echo "Usage $0 <path to source docs>"
  exit 1
fi

mkdir api-docs

for f in ${1}/*.json; do
  name=$(basename $f)
  jqscript=jq-scripts/${name}
  if [ -f "${jqscript}" ]; then
    jq -f "${jqscript}" $f > "api-docs/${name}"
  else
    cp $f api-docs
  fi
done
