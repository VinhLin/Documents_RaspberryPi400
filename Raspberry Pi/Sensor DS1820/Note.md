[ThamKhao1](https://www.waveshare.com/wiki/Raspberry_Pi_Tutorial_Series:_1-Wire_DS18B20_Sensor)
[ThamKhao2](https://www.raspberrypi-spy.co.uk/2013/03/raspberry-pi-1-wire-digital-thermometer-sensor/)


### Cấu hình cho Raspberry:

> sudo nano /boot/config.txt

- Thêm vào nội dung:

> dtoverlay=w1-gpio-pullup,gpiopin=4

- Sau đó reboot lại Pi.

### Test and Run:

> cd /sys/bus/w1/devices

> ls

- Mỗi con cảm biến DS1820 đều có một địa chỉ riêng.

- Nếu thấy địa chỉ bắt đầu bằng 00 thì tức là ta đã hàn sai gì đó rồi.

- Nếu ổn thì khi vào thư mục địa chỉ ID thì sẽ có w1_slave. Ví dụ id là: 28-01192eec59b3

> cd 28-01192eec59b3

> cat w1_slave

- Tải file code python để đo nhiệt độ.

> wget https://bitbucket.org/MattHawkinsUK/rpispy-misc/raw/master/python/ds18b20.py

- Mở file vừa tải chỉnh sửa lại id.

> sudo nano ds18b20.py

- Sau đó save lại và chạy:

> sudo python ds18b20.py
