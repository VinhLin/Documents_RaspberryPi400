- Tình trạng gặp phải trên bản **Ubuntu 20.20 ARM64** này thì phần audio bị lỗi.
- Vào Youtube thì không thấy phát âm thanh, mặc dù đã mở max loa, có cắm dây loa vào cổng HDMI thì vẫn không nghe thấy tiếng.

### Khắc phục
[Tham khảo 1](https://www.raspberrypi.org/forums/viewtopic.php?t=289126)
[Tham khảo 2](https://waldorf.waveform.org.uk/2020/ubuntu-desktops-on-the-pi.html)
- Thực hiện lệnh:
> sudo nano /etc/pulse/default.pa
- Add thêm 2 dòng sau:

```
load-module module-udev-detect tsched=0
set-default-sink 0
```
- Sau đó test lại trên Youtube thì mọi thứ OK.

### Note
- Ban đầu khắc phục bằng cách trên thì đã test **OK**. Nhưng hôm sau mở lên lại thì nó không còn âm thanh như trước nữa.
- Đã thử reboot, rút nguồn cắm lại nhưng vẫn không.
- Có thử seach trên gg thì cũng đã thực hiện cách khác nhưng vẫn không đc.
#### Tham khảo
[Link](https://ubuntu.pkgs.org/20.10/ubuntu-universe-arm64/pipewire-audio-client-libraries_0.3.10-4_arm64.deb.html)
> sudo apt-get update \
> sudo apt-get install pipewire-audio-client-libraries
#### Sửa lỗi
- Nguyên nhân gây ra tình trạng này, mình nghĩ là do **trên con Pi này có đến 2 cổng microHDMI, mà ban đầu mình đã out âm thanh được trên cổng 1, nhưng sau đó mình rút ra cắm cho cổng 2 thì sẽ không có âm thanh**.
- Cách khắc phục của mình là: Vào `/etc/pulse/default.pa` xóa đi nội dung 2 dòng đã thêm vào. Sau đó reboot lại Pi.
- Tiếp tục vào */etc/pulse/default.pa* thêm lại nội dung 2 dòng đã xóa, sau đó tiếp tục reboot Pi lại.
- Mình kiểm tra lại thì khi xem Youtube âm thanh đã có lại.

