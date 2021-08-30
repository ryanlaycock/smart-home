import grpc
import logging
from concurrent.futures import ThreadPoolExecutor

# python -m grpc_tools.protoc -I../ --python_out=. --grpc_python_out=. ../protos/sensors.proto
from protos import sensors_pb2_grpc, sensors_pb2
import temperature_sensor


class SensorServer(sensors_pb2_grpc.SensorServicer):
    def Read(self, request, context):
        logging.info('Read sensor request received')
        temp = temperature_sensor.read_temp(context)
        reading = sensors_pb2.Reading(temp)

        return reading


if __name__ == '__main__':
    logging.basicConfig(
        level=logging.INFO,
        format='%(asctime)s - %(levelname)s - %(message)s',
    )
    server = grpc.server(ThreadPoolExecutor())
    sensors_pb2_grpc.add_SensorServicer_to_server(SensorServer(), server)
    port = 9999
    server.add_insecure_port(f'[::]:{port}')

    server.start()
    logging.info('Temperature Sensor Server ready on port %r', port)
    server.wait_for_termination()
    logging.info('Temperature Sensor Shutdown')
