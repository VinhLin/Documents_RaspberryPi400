- Dựa trên [tài liệu](https://www.raspberrypi.org/documentation/configuration/uart.md) thì raspberry pi đều có 2 UART (ngoại trừ Pi 4 là có 6 UART).
- Trên con Compute Module 3+ (CM3+) cũng không ngoại lệ, cũng có 2 uart nhưng mặc định thì nó chỉ cho sử dụng 1 uart thôi, đó là **serial 0**.
- Muốn sử dụng uart thứ 2 thì cần phải khai báo nó.
[Tham khảo](https://www.raspberrypi.org/forums/viewtopic.php?t=194072)

### Thực hiện khai báo cả hai UART:
> sudo nano /boot/config.txt 
- Thêm nội dung sau:

```
 #UART0 aka ttyAMA0 pin 14,15
 enable_uart=1
 dtoverlay=pi3-disable-bt
 dtparam=uart0=on

 #UART1 aka ttyS0
 dtoverlay=uart1,txd1_pin=32,rxd1_pin=33
```
- Sau đó reboot pi, kiểm tra lại bằng lệnh **ls -lh /dev** ta sẽ thấy có đầy đủ cả 2 uart.














