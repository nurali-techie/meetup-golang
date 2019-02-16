from concurrent import futures

import time
import logging

import grpc

import model4_pb2
import model4_pb2_grpc

_ONE_DAY_IN_SECONDS = 60 * 60 * 24

class RoleService(model4_pb2_grpc.RoleServiceServicer):
    def GetRole(self, request, context):
        return model4_pb2.GetRoleReply(role="admin")

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    model4_pb2_grpc.add_RoleServiceServicer_to_server(RoleService(), server)
    server.add_insecure_port('[::]:8090')
    server.start()

    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    logging.basicConfig()
    serve()
