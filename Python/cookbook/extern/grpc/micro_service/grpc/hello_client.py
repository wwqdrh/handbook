import grpc
from micro_service.grpc.proto import hello_pb2, hello_pb2_grpc


def run():
    # 连接 rpc 服务器
    channel = grpc.insecure_channel("localhost:50051")
    # 调用 rpc 服务
    stub = hello_pb2_grpc.GreeterStub(channel)
    response = stub.SayHello(hello_pb2.HelloRequest(name="czl"))
    print("Greeter client received: " + response.message)
    response = stub.SayHelloAgain(hello_pb2.HelloRequest(name="daydaygo"))
    print("Greeter client received: " + response.message)


if __name__ == "__main__":
    run()