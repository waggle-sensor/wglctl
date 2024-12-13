#!/bin/bash

# build the cli app
go build .

#move to bin
mv wglctl /usr/local/bin/