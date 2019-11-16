#!/bin/bash

cd $GITHUB_WORKSPACE/src/grammar
ANTLR_EXEC="java -jar $1" make -e
