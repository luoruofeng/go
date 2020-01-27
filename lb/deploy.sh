#!/bin/bash

nohup ./bin/proxy </dev/null &>/dev/null &
nohup ./bin/rs -port=8080 </dev/null &>/dev/null &
nohup ./bin/rs2 -port=8081 </dev/null &>/dev/null &
nohup ./bin/rs3 -port=8082 </dev/null &>/dev/null &

# ./bin/proxy &
# ./bin/rs -port=8080 &
# ./bin/rs2 -port=8081 &
# ./bin/rs3 -port=8082 &