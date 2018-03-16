#!/bin/bash

NAME=John \
SURNAME=Connor \
./templater -t - -o /tmp/hello.txt <<EOF
Hello, {{env "NAME"}} {{env "SURNAME"}}!
EOF
