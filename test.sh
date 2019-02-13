#!/usr/bin/env bash

>&1 echo "hello from stdout, argvs are $@"
>&2 echo "hello from stderr, argvs are $@"
exit 2
