### Tham khảo
> http://docs.dinhviso.com.vn/s/bmtu2qrokmst95o4g9e0/software/d/bo01slbokmst95o4ge10/raspberrypi-hl-l2321d-printer \
> https://raspberrytips.com/install-printer-raspberry-pi/ \
> https://askubuntu.com/questions/101629/windows-7-cant-find-cups-printer-shared-from-ubuntu \
> https://pronoy.in/raspberrypi/cups/linux/printers/2020/10/08/cups-printer-server-on-the-raspberry-pi.html

## Setup
### Cài đặt
```
 apt-cache search brlaser
 apt-cache policy printer-driver-brlaser
 sudo apt install libcupsimage2-dev libcups2-dev cmake cups
 git clone https://github.com/pdewacht/brlaser.git
 cd brlaser
 cmake .
 make
 sudo make install
 sudo cupsctl --remote-admin --remote-any --share-printers
```
### Thiết lập
- Nhớ cắm dây USB kết nối giữa máy in với Pi.
- Vào [localhost](http://localhost:631/admin) này để thiết lập
- Sau khi vào trang chọn **Administration** -> **Add Printer**
- Đăng nhập (login và pass tương tự như kết nối vào Raspberry) (như Hình_A) 
```
 login: pi
 Pass: ldv1413212
```
- Tiếp tục thực hiện như theo trang `https://raspberrytips.com/install-printer-raspberry-pi/`
- Sau khi thực hiện sẽ đc kết quả như Hình_C.
- Vậy là xong. Mỗi lần cần in thì nhấn vào **Print** rồi chọn **Brother_HL-L2320D_series**, sau đó thiết lập những thông tin cần thiết, như:
`Copies là số bản cần in`
`Long side hoặc Two-sided là để chọn chức năng in 2 mặt`



