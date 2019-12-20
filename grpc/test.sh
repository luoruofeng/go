#/bin/bash

set +e

export TEMPDIR=$(mktemp)
trap "rm -rf ${TEMPDIR}" EXIT

clean(){
    jobss -p | xargs kill -9
    wait
}

fail () {
    echo "$(tput setaf 1) $1 $(tput sgr 0)"
    clean
    exit 1
}

important () {
    echo "$(tput setaf 4) $1 $(tput sgr 0)"
}

pass () {
    echo "$(tput setaf 2) $1 $(tput sgr 0)"
}

EXAMPLES=(
    "helloworld/helloworld"
    "helloworld/python_grpc/helloworld/protos"
)

PROTOS=(
    "helloworld/helloworld/"
    "helloworld/python_grpc/helloworld/protos/"
)

reset_data(){
    echo "reset:"
    for example in $*;do
        for file in `ls ./${example}/`;do
            if test -f ./$example/$file || test -d ./$example/$file
            then
                f="./$example/$file"
                if test ${f##*.} != "proto";then
                    rm -rf $f
                    important "remove $f"
                fi
            fi
        done
    done
}

protoc(){
    PWD=`pwd`

    cd "$PWD/${PROTOS[1]}"
    protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld

    # cd "$PWD/${PROTOS[2]}"
    # python -m grpc_tools.protoc -I./protos --python_out=./protos --grpc_python_out=./protos ./protos/helloworld.proto
}

reset_data ${EXAMPLES[@]}
protoc
