### Tham khảo
> https://www.youtube.com/watch?v=IBNZLREqBxg \
> https://en.wikipedia.org/wiki/Zram \
> http://apkmenet.blogspot.com/2014/12/huong-dan-tao-zram-e-tang-ram-cho.html \
> https://viblo.asia/p/toi-uu-hoa-su-dung-bo-nho-trong-he-thong-linux-maGK7DwAZj2

`ZRAM (ramzswap) sẽ nén các ứng dụng chạy ẩn (các quá trình chạy nền ) tới một phân vùng có tên là zram trên bộ nhớ máy. Khi mà lượng RAM thật đã bị đầy không thể lưu trữ được các quá trình chạy nền . Thông thường quá trình này sẽ bị diệt ngay khi bộ nhớ RAM thật bì đầy nhưng nếu bạn kích hoạt ZRAM thì các ứng dụng đó sẽ được nén và chuyển tới phân vùng ZRAM . Khi cần mở lại ứng dụng đó thì đã có sẵn trên ZRAM rồi và máy sẽ chạy nhanh hơn`

### Sử dụng trên Raspberry:
[Link](https://github.com/novaspirit/rpi_zram)
> sudo wget -O /usr/bin/zram.sh https://raw.githubusercontent.com/novaspirit/rpi_zram/master/zram.sh \
> sudo chmod +x /usr/bin/zram.sh \
> sudo nano /etc/rc.local
- Thêm nội dung sau vào trước dòng **exit 0**:
```
 # Use ZRam
 /usr/bin/zram.sh &
```

### Một số lệnh khác
- Xem các đĩa zram đang được sử dụng:
> sudo swapon -s
hoặc dùng lệnh:
> cat /proc/swaps
- Xem thông tin việc sử dụng zram. Ví dụ *zram0*
> sudo zramctl zram0



