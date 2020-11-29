Link tham khao:
> https://www.raspberrypi.org/documentation/configuration/external-storage.md



### Các bước cài đặt và thiết lập trên Pi:
	> sudo apt update
	> sudo apt install exfat-fuse
	> sudo apt install ntfs-3g
- Sau khi đã cài đặt xong các bước trên, ta cắm usb vào raspbeery và thực hiện lệnh:
	> sudo blkid
- Lúc này ta sẽ thấy được thông tin của con usb và ổ đĩa đang chứa nó.
- Bây giờ ta tiến hành tạo thư mục và mount:
	> sudo mkdir /media/usbcam
	> sudo mount /dev/sda1 /media/usbcam
- Ngoài ra để có thể tự động mount khi reboot lại pi thì ta thực hiện:
	> sudo nano /etc/rc.local
Chép nội dung sau trên dùng exit 0:	`sudo mount /dev/sda1 /media/usbcam`
Sau đó reboot lại pi và kiểm tra kết quả.




### Note:
- Nếu thư mục /media/usbcam này không trống (tức là có 1 file bất kì nào đó trong đó) thì quá trình mount sẽ thất bại.
- Khi đã mount thành công thì ta có thể dùng lệnh sau để kiểm tra xem liệu đã thực sự mount thành công chưa:
	> df -h
Lúc này ta sẽ thấy thông tin dòng cuối ví dụ như sau: 
`/dev/sda1        15G  4.7G  9.8G  33% /media/usbcam`




