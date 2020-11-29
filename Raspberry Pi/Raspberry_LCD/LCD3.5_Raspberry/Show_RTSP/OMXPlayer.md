## OMXPlayer

### Tham khảo
> https://www.raspberrypi.org/documentation/raspbian/applications/omxplayer.md \
> https://elinux.org/Omxplayer#RTSP

### Cài đặt
> sudo apt-get install omxplayer


- Anh Toản nói có thể dùng **omxplayer** này để show video từ stream camera qua màn hình LCD.
- Mình đã thử và thật sự là đã thành công. Mình sử dụng bản Raspbian Lite và cài đặt **omxplayer**.
> omxplayer -o local rtsp://192.168.10.47:554/user=admin_password=tlJwpbo6_channel=1_stream=1.sdp?real_stream


### Thư mục show-picture
> https://git.dinhviso.com.vn/vinhlin/show-picture





