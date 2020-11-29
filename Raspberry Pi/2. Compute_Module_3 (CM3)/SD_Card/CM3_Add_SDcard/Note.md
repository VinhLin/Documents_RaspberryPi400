Link:
> http://docs.dinhviso.com.vn/s/bmtu2qrokmst95o4g9e0/software/d/bptgslbokmss4cjb6vag/itracking-3?currentPageId=&source=itracking%25203
> https://ralimtek.com/raspberry%20pi/electronics/software/raspberry_pi_secondary_sd_card/
> https://github.com/raspberrypi/linux/blob/rpi-4.9.y/arch/arm/boot/dts/overlays/sdio-overlay.dts
> https://www.raspberrypi.org/forums/viewtopic.php?p=1219308#p1219308
> https://www.raspberrypi.org/documentation/hardware/computemodule/cm-emmc-flashing.md



### Boot Rapbian cho CM3 bằng flash, còn thẻ nhớ (sd card) thì chứa dữ liệu.

### SDIO pin : 34-39
### Các bước thực hiện:
- Chép 2 file *sdio34.dtbo* và *sdio34.dts* sang cho Pi.

- Gõ lệnh **sudo nano /boot/config.txt** và thêm nội dung:
> dtoverlay=sdio34,poll_once=off,bus_width=1,sdio_overclock=20

- Copy sdio34.dtbo sang /boot/overlays, lệnh:
> sudo cp sdio34.dtbo /boot/overlays/

- Gõ lệnh **sudo nano /etc/rc.local** và thêm nội dung:
> sudo mount /dev/mmcblk1p1 /media/sdcardsecond/

- Tạo thư mục /media/sdcardsecond/
> sudo mkdir /media/sdcardsecond/

- Sau đó **reboot** lại Pi. Thế là xong.	







