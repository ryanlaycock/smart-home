import os
import glob
import logging
import grpc

os.system('modprobe w1-gpio')
os.system('modprobe w1-therm')

BASE_DIR = '/sys/bus/w1/devices/'
DEVICE_FOLDER_NAME = '/w1_slave'

device_folders = glob.glob(BASE_DIR + '28*')

def read_temp_raw(device_file):
    f = open(device_file, 'r')
    lines = f.readlines()
    f.close()

    return lines


def read_temp(context):
    if len(device_folders) == 0:
        context.set_details('Cannot find sensor in devices directory')
        context.set_code(grpc.StatusCode.INTERNAL)
        return 0

    device_file = device_folders[0] + DEVICE_FOLDER_NAME
    lines = read_temp_raw(device_file)
    logging.debug('Read device lines: %v', lines)

    while lines[0].strip()[-3:] != 'YES':
        context.set_details('Sensor has no values')
        context.set_code(grpc.StatusCode.INTERNAL)
        logging.debug('Couldn\'t find temperature value')
        return 0

    if len(lines) > 1:
        equals_pos = lines[1].find('t=')
        if equals_pos != -1:
            temp_string = lines[1][equals_pos + 2:]
            temp_c = float(temp_string) / 1000.0
            logging.info('Parsed temperature value to celsius: %v', temp_c)

            return temp_c

    context.set_details('Sensor has no values')
    context.set_code(grpc.StatusCode.INTERNAL)
    logging.debug('Error parsing temperature value')

    return 0
