- Tình trạng là khi cắm ổ cứng ngoài vào TwisterOS thì mặc dù mình thấy nó có nhận đc ổ cứng, nhưng click vào thì thư mục trống.
- Trên Raspbian thì vẫn bình thường.
- Khi mình thử dùng lệnh **mount** thì nó báo lỗi
`Mount is denied because the NTFS volume is already exclusively opened`

### Cách khắc phục 
[Tham khảo 1](https://www.raspberrypi.org/documentation/configuration/external-storage.md)
[Tham khảo 2](https://documentation.storj.io/resources/faq/linux-static-mount)
- Cài đặt 1 số tools trước:
> sudo apt update \
> sudo apt install exfat-fuse \
> sudo apt install ntfs-3g
- Dùng lệnh `sudo blkid` để lấy **UUID** của ổ cứng.
- Sau đó thực hiện:
> sudo mkdir /mnt/mydisk \
> sudo nano /etc/fstab
- Dựa trên nội dung sau và chỉnh sửa thích hợp để thêm vào
`UUID=<your HD UUID> /mnt/<directory name> <FSTYPE> defaults 0 2`
- Ví dụ ở đây mình có **UUID của ổ cứng là 7414D30514D2C96E**, thư mục sẽ mount là **/mnt/mydisk**, ổ cứng mình định dạng là **NTFS**. 
- Đây sẽ là nội dung cần thêm vào:
> UUID=7414D30514D2C96E /mnt/mydisk ntfs defaults 0 2
- Sau đó reboot lại máy, kiểm tra lại trong thư mục **/mnt/mydisk** thì mình sẽ thấy dữ liệu trong ổ cứng.
- Dùng lệnh này để kiểm tra lại:
> sudo mount -a
