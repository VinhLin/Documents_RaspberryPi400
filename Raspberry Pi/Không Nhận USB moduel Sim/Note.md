- Hầu hết các Driver của các thiết bị Linux thường đã đc nhúng sẵn trên HĐH Linux rồi [Hình0](https://tldp.org/LDP/tlk/sources/sources.html). 

- Do đó nếu bình thường khi cắm thiết bị vào trên HĐH Linux mà không nhận đc, thì ta chỉ cần thực hiện kiểm tra phiên bản linux hiện tại và upgrade lên phiên bản cao hơn thì sẽ nhận đc thiết bị. [Link](https://technicalustad.com/fix-usb-device-not-recognized-in-raspberry-pi/)
> uname -a
> sudo apt-get upgrade

## Trường hợp mình gặp phải trên con module sim EG91 là: nhận thiếu driver.
- Mặc dù đã upgrade phiên bản mới nhất cho linux, dùng lệnh **lsusb** thì vẫn thấy module sim.
- Nhưng khi dùng lệnh **ls -lh /dev/** thì lại không thấy cổng USB của moduel sim.
`Thường thì khi nhận đc module sim thì nó sẽ thấy cổng usb /dev/ttyUSB0 -> /dev/ttyUSB3 (hoặc /dev/ttyACM0 -> /dev/ttyACM3)`
- Dùng lệnh **dmesg** để kiểm tra lần nữa thì đúng là nó có nhận đc driver module sim EG91 (như hình 1).
- Điều này có nghĩa là linux có nhận driver nhưng nhận chưa đủ khiến cho việc mount vào usb bị lỗi, **nguyên nhân có thể là do con module sim này mới linux chưa cập nhật đủ driver**. 
- Theo như trang [Tham khảo](https://bacnh.com/quectel-linux-usb-drivers-troubleshooting/) này thì ta khắc phục như sau: 

```
 Bước 1: Kiểm tra **Vendor ID** và **Product ID** bằng lệnh:
 > lsusb
 Bước 2: `echo "$VID $PID" > /sys/bus/usb-serial/drivers/option1/new_id`. Ví dụ module sim EG91 có VID là 2c7c và PID là 0191.
 > sudo su
 > modprobe option
 > echo "2c7c 0191" > /sys/bus/usb-serial/drivers/option1/new_id 
 Bước 3: Kiểm tra lại lần nữa bằng lệnh:
 > dmesg |grep tty*
```

### Lưu ý:
- Hiện tại dựa vào các trang tham khảo bên dưới, tính chỉnh sửa luôn bên trong **Patch kernel drivers** nhưng chưa đc.
[Tham khảo 1](https://en.dlyang.me/install-linux-usb-driver-for-quectel-bg96/)
[Tham khảo 2](https://lore.kernel.org/patchwork/patch/853074/)
[Tham khảo 3](https://hackersgrid.com/2019/04/unable-to-locate-package-linux-headers.html)
[Tham khảo 4](https://www.raspberrypi.org/forums/viewtopic.php?t=250664)
[Tham khảo 5](https://github.com/bacnh85/Quectel_Linux_USB_Driver)
[Tham khảo 6](https://github.com/notro/rpi-source/wiki)
- Do đó mình sẽ đặt các lệnh mount usb vào file **rc.local**
> sudo nano /etc/rc.local
> modprobe option
> echo "2c7c 0191" > /sys/bus/usb-serial/drivers/option1/new_id 














