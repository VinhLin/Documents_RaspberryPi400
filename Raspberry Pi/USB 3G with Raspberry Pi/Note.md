### Link:
> https://www.sbprojects.net/projects/raspberrypi/mobile.php \
> https://kipalog.com/posts/Raspberry-Pi-ket-noi-Internet-su-dung-USB-3G-Viettel---D6601 \
> https://www.thefanclub.co.za/how-to/how-setup-usb-3g-modem-raspberry-pi-using-usbmodeswitch-and-wvdial

## Cac buoc thuc hien:
### B1: Cai dat 
> sudo apt-get install ppp usb-modeswitch wvdial

### B2: Kiem tra cong USB
> ls /dev/ttyUSB*

### B3: Chep file wvdial.conf
> sudo nano /etc/wvdial.conf

### B4: sudo nano /etc/network/interfaces
- Them vao dong:
```
auto ppp0
iface ppp0 inet wvdial
```

### B5: Ket noi
> sudo wvdial 3gviettel
- Neu thay led tren usb3g, tuc la no dang thuc hien ket noi.

### B6: Kiem tra ta co the dung lenh 
- De xem co ket noi ppp0 ko
> ifconfig	
- De kiem tra co ket noi dc ko
> ping 8.8.8.8









