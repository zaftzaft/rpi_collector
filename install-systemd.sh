service=rpi-collector

systemctl is-enabled $service
if [ $? -eq 0 ]; then
  systemctl stop $service
fi


if [ -d /usr/lib/systemd/system/ ]; then
  unit_dir=/usr/lib/systemd/system
else
  unit_dir=/etc/systemd/system
fi

install -Dm644 systemd/$service.service $unit_dir/$service.service
install -Dm644 systemd/$service.timer $unit_dir/$service.timer

mkdir -p /etc/conf.d
install -Dm644 systemd/$service /etc/conf.d/$service


systemctl daemon-reload

