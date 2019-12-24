#/bin/bash

set +e

export TEMPDIR=$(mktemp)
trap "rm -rf ${TEMPDIR}" EXIT

clean(){
    jobs -p | xargs kill -9
    wait
}

trap clean EXIT

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
    "helloworld/"
    "helloworld/python_grpc/helloworld/"
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


CURRENT_FILEPATH=$(cd `dirname $0`; pwd)

create_protoc(){
    cd "$CURRENT_FILEPATH/${PROTOS[0]}"
    protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
    pass "execute protoc in $CURRENT_FILEPATH/${PROTOS[0]} SUCCESSED!"

    cd "$CURRENT_FILEPATH/${PROTOS[1]}"
    python -m grpc_tools.protoc -I./protos --python_out=./protos --grpc_python_out=./protos ./protos/helloworld.proto
    pass "execute protoc in $CURRENT_FILEPATH/${PROTOS[1]} SUCCESSED!"
}

reset_data ${EXAMPLES[@]}
create_protoc

# Build server
cd "$CURRENT_FILEPATH/${PROTOS[0]}"
if ! go build -o /dev/null ./*server/*.go; then
    fail "failed to build server"
else
    pass "successfully built server"
fi

# Start server
SERVER_LOG="$(mktemp)"
go run ./*server/*.go &> $SERVER_LOG  &


# Build client
if ! go build -o /dev/null ./*client/*.go; then
    fail "failed to build client"
else
    pass "successfully built client"
fi

#Start client
CLIENT_LOG="$(mktemp)"


#linux
# if ! timeout 20 go run ./*client/*.go &> $CLIENT_LOG; then
#Mac
#brew install coreutils before execute next line
if ! gtimeout 20 go run ./*client/*.go luoruofeng &> $CLIENT_LOG; then
    fail "client failed to communicate with server
    got server log:
    $(cat $SERVER_LOG)
    got client log:
    $(cat $CLIENT_LOG)
    "
else
    pass "client successfully communitcated with server"
    echo "
    got server log:
    $(cat $SERVER_LOG)
    got client log:
    $(cat $CLIENT_LOG)
    "
fi

#这对代码会出错
#protoc 生成的python文件会有错：
#import helloworld_pb2 as helloworld__pb2
#改为
#from . import helloworld_pb2 as helloworld__pb2
#再手动执行
#Start python client,communicate with go server in grpc
cd $CURRENT_FILEPATH/helloworld/python_grpc/*client
export PYTHONPATH=$CURRENT_FILEPATH/helloworld/python_grpc
if ! gtimeout 20 python ./*.py luoruofeng &> $CLIENT_LOG; then
    fail "python client failed to communicate with server
    got server log:
    $(cat $SERVER_LOG)
    got client log:
    $(cat $CLIENT_LOG)
    "
else
    pass "python client successfully communitcated with server"
    echo "
    got server log:
    $(cat $SERVER_LOG)
    got client log:
    $(cat $CLIENT_LOG)
    "
fi

