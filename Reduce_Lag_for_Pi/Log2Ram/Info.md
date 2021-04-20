- Theo như tìm hiểu thì **log2ram** nó giúp mình ghi log trên *Ram* thay vì ghi trực tiếp trên *sdcard*.
- Khi ghi log trên ram vậy thì sẽ hạn chế việc **hư hỏng thẻ nhớ**.

### Tham khảo
> https://mcuoneclipse.com/2019/04/01/log2ram-extending-sd-card-lifetime-for-raspberry-pi-lorawan-gateway/ \
> https://www.ubuntupit.com/how-to-write-log-files-in-ram-using-log2ram-in-linux/

## Setup
[Tham khảo Git](https://github.com/azlux/log2ram)
- Thực hiện các lệnh sau:
```
 sudo su
 echo "deb http://packages.azlux.fr/debian/ buster main" | sudo tee /etc/apt/sources.list.d/azlux.list
 wget -qO - https://azlux.fr/repo.gpg.key | sudo apt-key add -
 apt update
 apt install log2ram
```
- Mình chỉnh lại *SIZE* thành **128 MByte** trong file:
> nano /etc/log2ram.conf
- Sau đó dùng lệnh **mount** hoặc **reboot** lại Pi.
- Kiểm tra lại bằng lệnh **df -h**

### Note
- Mình không biết **log2ram** có hoạt động hay ko? Vì trong quá trình ghi log thì mình có *reboot* lại, và kiểm tra lại **syslog** thì vẫn có ghi lại log cho đến khi reboot.
- Còn khi rút nguồn đột ngột thì vẫn có log cho đến lúc mất nguồn. 
`Mình không biết là log2ram hoạt động tốt hay việc ghi log vẫn ghi vào sdcard như thông thường`.

### Disable Log2Ram
> sudo systemctl disable log2ram.service 
