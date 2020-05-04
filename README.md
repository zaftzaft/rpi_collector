rpi-collector
=============

## Install
```
go build
install -Dm755 rpi_collector /usr/bin/rpi_collector

sh install-systemd.sh

systemctl enable rpi-collector.timer
systemctl start rpi-collector.timer


cat /tmp/node_exporter/rpi.prom
```

## Mertics
```
# HELP rpi_clock current frequency of the specified clock
# TYPE rpi_clock gauge
rpi_clock{clock="H264"} 0
rpi_clock{clock="arm"} 1.2e+09
rpi_clock{clock="core"} 4e+08
rpi_clock{clock="dpi"} 0
rpi_clock{clock="emmc"} 2e+08
rpi_clock{clock="hdmi"} 0
rpi_clock{clock="isp"} 2.5e+08
rpi_clock{clock="pixel"} 0
rpi_clock{clock="pwm"} 0
rpi_clock{clock="uart"} 4.8e+07
rpi_clock{clock="v3d"} 2.99999e+08
rpi_clock{clock="vec"} 0
# HELP rpi_temp temperature of the SoC as measured by the on-board temperature sensor
# TYPE rpi_temp gauge
rpi_temp 46.7
# HELP rpi_throttled  throttled state of the system
# TYPE rpi_throttled gauge
rpi_throttled{meaning="arm_frequency_capped"} 0
rpi_throttled{meaning="arm_frequency_capping_has_occurred"} 0
rpi_throttled{meaning="currently_throttled"} 0
rpi_throttled{meaning="soft_temperature_limit_active"} 0
rpi_throttled{meaning="soft_temperature_limit_has_occurred"} 0
rpi_throttled{meaning="throttling_has_occurred"} 1
rpi_throttled{meaning="under_voltage_detected"} 0
rpi_throttled{meaning="under_voltage_has_occurred"} 1
# HELP rpi_volts current voltages used by the specific block.
# TYPE rpi_volts gauge
rpi_volts{block="core"} 1.325
rpi_volts{block="sdram_c"} 1.2
rpi_volts{block="sdram_i"} 1.2
rpi_volts{block="sdram_p"} 1.225

```
