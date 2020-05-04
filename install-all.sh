go build

install -Dm755 rpi_collector /usr/bin/rpi_collector

sh install-systemd.sh

systemctl enable rpi-collector.timer
systemctl start rpi-collector.timer
