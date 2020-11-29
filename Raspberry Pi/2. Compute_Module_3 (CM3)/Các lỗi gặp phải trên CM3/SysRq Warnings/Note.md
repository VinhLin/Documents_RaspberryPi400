### Đây là lỗi thiết bị tự reset sau 1 khoảng thời gian hoạt động.

Link tham khảo:
> https://www.raspberrypi.org/forums/viewtopic.php?t=73701
> https://support.cumulusnetworks.com/hc/en-us/articles/204031098-SysRq-and-Cumulus-Linux


- Dựa theo link tham khảo trên thì nguyên nhân dẫn đến thiết bị tự reset là khi ta sử dụng chân RX,TX để truyền nhận dữ liệu thì chưa disabled the console serial port nên mới bị.

- Ở đây mình làm theo cả 2 hướng khắc phục luôn cho chắc ăn:
#### Hướng 1:
	> sudo su
	> echo 0 > /proc/sys/kernel/sysrq
#### Hướng 2:
	> sudo nano /boot/config.txt 
Add thêm dòng enable_uart=1 vào cuối hàng.
	> sudo systemctl mask serial-getty@ttyama0.service
	> sudo nano /boot/cmdline.txt
Xóa đi phần nội dung “console=serial0,115200”
	> sudo reboot

- Sau khi thực hiện các bước trên và reboot lại thì mình thấy nó không còn xuất hiện hiện tượng tự reset thiết bị nữa và trên app gps (code itracking_pro) không thấy bị tình trạng read only file system xảy ra trên /home/pi.





