- Sau khi đọc tham khảo các tài liệu về cách tạo và sử dụng **Ram Disk** thì mình đã thử thực hiện thành công trên con Pi.
[Tham khảo 1](https://www.jamescoyle.net/how-to/943-create-a-ram-disk-in-linux)
[Tham khảo 2](https://www.linuxbabe.com/command-line/create-ramdisk-linux)

### Thực hiện
> sudo mkdir /mnt/ramdisk \
> sudo chmod 777 /mnt/ramdisk/ \
> sudo mount -t tmpfs -o size=30m tmpfs /mnt/ramdisk
- Sau đó dùng lệnh **df -h** thì ta sẽ thấy xuất hiện phân vùng *tmpfs* có dung lượng là *30MB*
- Để có thể tự khởi tạo ramdisk này mỗi khi khởi động Pi thì thực hiện sửa trong file */etc/fstab*
> sudo nano /etc/fstab \
- Thêm nội dung sau:
`tmpfs       /mnt/ramdisk tmpfs   defaults,noatime,nosuid,mode=0777,size=30m   0 0`

### Test
- Mình sẽ thử dùng record lại video từ camera và lưu lại vào ramdisk này.
> gst-launch-1.0 -e rtspsrc location="rtsp://192.168.10.117:554/user=admin_password=tlJwpbo6_channel=1_stream=1.sdp?real_stream" ! rtph264depay ! h264parse ! mp4mux ! filesink buffer-mode=0 buffer-size=1048576 location=/mnt/ramdisk/test.mp4
- Kiểm tra trên **df -h** thì sẽ thấy đc phần dung lượng.
- Nếu mà mình không backup data có trong **/mnt/ramdisk** thì khi reboot lại Pi, data có trong */mnt/ramdisk* sẽ mất. Vì bản chất của Ram là bộ nhớ lưu trữ bay hơi (sẽ mất dữ liệu khi khởi động lại).



