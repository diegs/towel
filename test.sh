#!/usr/bin/env bash

>&2 echo "12345 {this is a valid log output, argvs are $@}"
>&2 echo "this is invalid logging"
>&2 echo "on many lines"
>&2 echo "12345 {back to valid logging}"
>&2 echo "12345 {for a few lines}"
>&2 echo "and then a panic!"
>&2 echo "on a few lines"
exit 2
