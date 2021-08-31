import grpc
import logging
from concurrent.futures import ThreadPoolExecutor

from protos import sensors_pb2_grpc, sensors_pb2
import temperature_sensor


class SensorServer(sensors_pb2_grpc.SensorServicer):
    def ReadTemp(self, request, context):
        sensorId = request.sensorId
        logging.info('Read temp sensor request received for %s', sensorId)
        temp = temperature_sensor.read_temp(context, sensorId)
        reading = sensors_pb2.Reading(value=temp)

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
    logging.info('Sensor Server ready on port %r', port)

    for sensor_id in temperature_sensor.device_folders:
        logging.info('Found temperature sensor: %s', sensor_id)

    server.wait_for_termination()
    logging.info('Sensor Server shutdown')
