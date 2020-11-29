- Theo mặc định thì serial 0 sẽ ứng với UART 0 (hay còn gọi là **primary UART** ứng với chân **GPIO 14 (transmit)** and **GPIO 15 (receive)**)

### Để có thể sử dụng serial 0 trên CM3+ ta thực hiện:
> sudo raspi-config

```
Chọn Interfacing Options -> Serial
Chọn No //Disable serial login shell
Chọn Yes //Enable serial interface
```
- Sau đó nhấn OK để reboot lại pi
- Tiếp tục thực hiện:
> sudo systemctl disable hciuart

> sudo nano /boot/config.txt
- Thêm nội dung này vào

```
 enable_uart=1
 dtoverlay=disable-bt
```
- Thực hiện câu lệnh
> sudo systemctl mask serial-getty@ttyAMA0.service
- Mở file **/boot/cmdline.txt**
> sudo nano /boot/cmdline.txt
- Xóa dòng **console=serial0,115200** sau đó save lại.
- Reboot lại pi. Thế là xong


### Kiểm tra
- Cài đặt **screen** trên con pi.
- Dùng lệnh sau để kiêm tra xem có nhận đc dữ liệu từ uart hay không (ở đây là mình có data từ stm32 truyền tới pi thông qua uart).
> sudo screen /dev/ttyAMA0 115200







