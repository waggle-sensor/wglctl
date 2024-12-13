#!/bin/bash

# build the cli app
go build .

#move to bin
mv wglctl /usr/local/bin/

echo "Installed, try using 'wglctl' in your terminal"