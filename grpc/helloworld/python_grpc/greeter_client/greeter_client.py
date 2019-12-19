from __future__ import print_function
import logging

import grpc

import sys
sys.path.append("/Users/luoruofeng/Desktop/go/grpc/helloworld/python_grpc")
#  if u don't wanna use above line,u can export PYTHONPATH
#  export PYTHONPATH=/Users/luoruofeng/Desktop/go/grpc/helloworld/python_grpc

from helloworld.protos import helloworld_pb2
from helloworld.protos import helloworld_pb2_grpc
#Add __init__.py file to python dir,this will be become to python model,if u want to run module,please use `python -m xxxx.py`

def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = helloworld_pb2_grpc.GreeterStub(channel)
        response = stub.SayHello(helloworld_pb2.HelloRequest(name='you'))
        response2 = stub.SayHelloAgain(helloworld_pb2.HelloRequest(name='you'))
    print("Greeter client received: " + response.message)
    print("Greeter client received: " + response2.message)


if __name__ == '__main__':
    logging.basicConfig()
    run()