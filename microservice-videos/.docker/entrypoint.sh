#!/bin/bash

#compile app
go build

#start app after wait the database
sh -c "/wait && ./backend"