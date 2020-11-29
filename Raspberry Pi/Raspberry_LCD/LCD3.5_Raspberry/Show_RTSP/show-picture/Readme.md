## Mục đích
`Hiện thị hình ảnh của camera lên LCD để từ đó chỉnh lại góc camera khi lắp đặt thiết bị ItrackingPro.`

- Sử dụng **OMXPlayer** để show streamURI camera lên camera.
- Phần code **show-picture** được viết dựa trên nền code **vision-photo**.

### Set IP tĩnh
- Con Pi đc set ip tĩnh là *192.168.10.100*. Vì các camera đều đc set là *192.168.10.*
> sudo nano /etc/network/interfaces

```
auto eth0
iface eth0 inet static 
  address 192.168.10.100
  netmask 255.255.255.0
  gateway 192.168.10.100

```

### Giảm thời gian khởi động Pi
[Tham Khảo](https://www.myhelpfulguides.com/2018/10/20/how-improve-raspberry-pi-boot-time-raspbian-lite/)
> sudo nano /boot/cmdline.txt
- Thêm **quiet** trước *rootwait*
- Sau đó reboot lại Pi.

### Chép file service và app qua Pi

```
 sudo scp showpicture.service pi@192.168.10.100:
 sudo scp show_picture_1.0.1 pi@192.168.10.100:
```

### Service 
> sudo mv showpicture.service /lib/systemd/system \
> sudo systemctl daemon-reload \
> sudo systemctl enable showpicture.service 
### App show_picture
> sudo cp show_picture_1.0.1 /usr/bin/show_picture

- Cuối cùng là **reboot** lại Pi.








