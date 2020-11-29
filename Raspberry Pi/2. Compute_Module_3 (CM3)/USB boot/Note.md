### USB Boot 
- USB Boot này là để giúp cho ta có thể đọc được Flash của con CM3+.
- Từ đó ta có thể cài đặt hệ điều hành Raspbian lên CM3+.
- Hoặc có thể đọc đc ổ boot và roofs trên Linux (nếu đã cài đặt HĐH thành công).
[Link](https://www.raspberrypi.org/documentation/hardware/computemodule/cm-emmc-flashing.md)


### Install và Sử dụng RPiboot trên Ubuntu
> sudo apt install git
> git clone --depth=1 https://github.com/raspberrypi/usbboot
> cd usbboot
> sudo apt install libusb-1.0-0-dev
> make
> sudo ./rpiboot
- Để mỗi lần sử dụng **rpiboot** không cần nhất thiết phải vào thư mục **usbboot** thì ta chỉ cần copy file app rpiboot sang thư mục **/usr/bin**. Mỗi lần dùng thì chỉ cần dùng lệnh: `sudo rpiboot`
- Còn thư mục usbboot phải giữ lại chứ đừng xóa nó đi, nếu xóa đi thì sẽ không dùng đc rpiboot nữa.


