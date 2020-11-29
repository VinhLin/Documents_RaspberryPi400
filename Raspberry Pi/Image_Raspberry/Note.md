## Install Image Raspbian on Ubuntu

### Cách 1: Cách làm thông thường mình hay làm.
- Khi tạo ra một file **backup** image cho ItrackingV3 thì dung lượng của nó là hơn 7G.
- Sau đó trên Ubuntu mình chạy **pishrink.sh** để giảm dung lượng xuống còn 3.5G
- Rồi trên Windows dùng tool Win32DiskImager để cài đặt.

### Cách 2: Cài trực tiếp trên Ubuntu 
- Sử dụng một tool khác tương tự như Win32DiskImager trên Ubuntu.
- Và đó là [etcher](https://www.balena.io/etcher/) 
[Tham khảo 1](https://itsfoss.com/tutorial-how-to-install-raspberry-pi-os-raspbian-wheezy/)
[Tham khảo 2](https://magpi.raspberrypi.org/articles/pi-sd-etcher)
- Ngoài ra, anh Toản có sử dụng **raspberrypi_image_resize.sh** cũng có tác dụng là giảm dung lượng file, nhưng nó làm thêm việc là nén lại thành file **.gz** có dung lượng 1.5G (nó giống như nén **.rar**, nếu burn cái file .gz này ra thì mình vẫn có đc file .img có dung lượng là 3.5G)
- Bây giờ mình download **etcher**, giải nén nó và sử dụng thôi (ko cần phải cài đặt gì cả, nhấn chạy trực tiếp, còn ko thì chạy bằng lệnh). Tool này mình thử thì nó có thể chạy đc file *.gz*, không cần phải giải nén nó, rất tiện và mình còn tiết kiệm đc dung lượng image cho Pi.
- Cách sử dụng cũng tương tự như trên Win32DiskImager.
`Lưu ý là đối với con Compute Module thì nhớ chạy file rpiboot trước để có thể đọc đc fash của Pi.




