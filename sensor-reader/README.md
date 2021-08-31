# Sensor Reader (for RPi)
This package is designed to provide a gRPC interface to multiple sensors on a raspberry pi. Configuration of the service,
and each sensor on the RPi is defined below.
 
## Temperature Sensors
This service is designed to be run on a raspberry Pi, that has a number of DS18b20 sensors connected to the GPIO pins
to measure temperature. 

### Raspberry Pi Configuration
The DS18b20 uses the one wire communication protocol, which first needs to be enabled on the RPi

Add the following line to the `/boot/config.txt` file, and `reboot`:

``dtoverlay=w1-gpio``

To view the connected sensors to configure and name them later, run the following commands (these commands are run at 
runtime in the service normally):

``modprobe w1-gpio``

``modprobe w1-therm``

You can then view the sensors by navigating to `/sys/bus/w1/devices`. Each device will appear as a sub folder here
beginning with `28-`. To view the temperature of each sensor and to check they're working `cd` into the device's folder
and examine the `w1-slave` file.

## External Links and References
- DS18b20 sensors used
  - https://www.amazon.co.uk/gp/product/B01MDQO9C2/ref=ppx_yo_dt_b_search_asin_title?ie=UTF8&psc=1
- Thermometer sensor diagrams
  - https://www.circuitbasics.com/raspberry-pi-ds18b20-temperature-sensor-tutorial/