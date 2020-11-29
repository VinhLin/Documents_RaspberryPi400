- Đây là các bước mình đã làm trên con Pi về việc can thiệp vào **kernel source** để khắc phục việc không nhận được cổng usb của nodule sim. 
- Mặc dù đã khắc phục bằng tay bằng 2 lệnh dưới, nhưng mình vẫn thử can thiệp vào source linux để config lại thế nào.
 
```
 modprobe option
 echo "2c7c 0191" > /sys/bus/usb-serial/drivers/option1/new_id 
```
- Mặc dù đã thử nhưng vẫn thất bại, ở đây mình note lại các bước đã làm trên raspberry.

## Thực hiện
### Lấy source linux
#### Cách 1: 
> sudo apt-get install raspberrypi-kernel
> sudo apt-get install raspberrypi-kernel-headers
> sudo su
> apt update -y && apt upgrade -y && apt dist-upgrade
> sudo reboot
> sudo apt-get install linux-headers-$(uname -r)
- Thất bại. Nếu lấy source thành công thì nó sẽ có trong thư mục **/usr/src**.
#### Cách 2:
- Dựa theo [link](https://github.com/notro/rpi-source/wiki)
- Cái này thì làm thành công, thực hiện 
> sudo apt-get install git bc bison flex libssl-dev
> sudo apt-get install libncurses5-dev make
> sudo apt-get install ncurses-dev
> sudo wget https://raw.githubusercontent.com/notro/rpi-source/master/rpi-source -O /usr/local/bin/rpi-source && sudo chmod +x /usr/local/bin/rpi-source && /usr/local/bin/rpi-source -q --tag-update
> rpi-source

### Thực hiện thay đổi 
[ThamKhao1](https://www.raspberrypi.org/forums/viewtopic.php?t=188337)
[ThamKhao2](https://www.raspberrypi.org/documentation/linux/kernel/building.md)
> cd linux/drivers/usb/serial
> nano option.c
- Tìm phần Quectel có trong file, thêm tương tự như cho EG91 với PID là **0x0191**
> cd
> cd linux
> KERNEL=kernel7
> make ARCH=arm CROSS_COMPILE=arm-linux-gnueabihf- bcm2709_defconfig
> make prepare
> make menuconfig
make

make ARCH=arm CROSS_COMPILE=arm-linux-gnueabihf- menuconfig

export ARCH=arm
export CROSS_COMPILE=arm-linux-gnueabihf-
make bcm2709_defconfig
make menuconfig

sudo make -C /lib/modules/`uname -r`/build M=`pwd`/drivers/usb/class obj-m=cdc-acm.o modules







