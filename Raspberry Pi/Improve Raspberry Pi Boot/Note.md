Link:
https://www.myhelpfulguides.com/2018/10/20/how-improve-raspberry-pi-boot-time-raspbian-lite/
https://raspberrypi.stackexchange.com/questions/79728/keep-hdmi-off-on-boot

### Command
sudo nano /boot/config.txt
sudo nano /boot/cmdline.txt


### Improve Boot
> sudo nano /boot/cmdline.txt \
- Add **quiet** trước *rootwait*
- Sau đó reboot lại Pi.
















