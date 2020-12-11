- Mục đích là cùng hiện một đoạn video cùng lúc với nhau trên nhiều màn hình.
- Giả sử như mình có 2 con Pi và đang muốn show một đoạn video trên 2 màn hình khác nhau (một con Pi ứng với 1 màn hình LCD).
- Để làm đc điều này thì mình sẽ dùng đến **omxplayer-sync**.
### Tham khảo
> https://www.khm.de/~sievers/net/omxplayer-sync.html \
> https://github.com/turingmachine/omxplayer-sync \
> https://vimeo.com/137133716

## Setup trên cả 2 con Pi.
```
 sudo su
 apt-get remove omxplayer
 rm -rf /usr/bin/omxplayer /usr/bin/omxplayer.bin /usr/lib/omxplayer
 apt-get install libpcre3 fonts-freefont-ttf fbset libssh-4 python3-dbus
 wget https://github.com/magdesign/PocketVJ-CP-v3/raw/master/sync/omxplayer_0.3.7-git20170130-62fb580_armhf.deb
 dpkg -i omxplayer_0.3.7-git20170130~62fb580_armhf.deb
 wget -O /usr/bin/omxplayer-sync https://github.com/turingmachine/omxplayer-sync/raw/master/omxplayer-sync
 chmod 0755 /usr/bin/omxplayer-sync
 wget https://github.com/turingmachine/omxplayer-sync/raw/master/synctest.mp4
```
- Mình muốn con Pi nào đóng vai trò là *master* thì sẽ chạy lệnh:
> omxplayer-sync -muv synctest.mp4
- Con Pi còn lại sẽ là *slave*:
> omxplayer-sync -luv synctest.mp4
### Ở đây có một số chú ý là:
- **Filename** phải trùng nhau trên cả master và slave.
- Phải kết nối cáp LAN giữa các con Pi để có thể truyền data.
- Cái này chỉ truyền đc hình ảnh thôi.

### Fix lỗi.
[Tham khảo](https://github.com/turingmachine/omxplayer-sync/issues/36)
> sudo apt --fix-broken install
- Sau đó cài đặt lại **python3-dbus**
> sudo apt-get install python3-dbus

#### Note: Mình cài xong hết rồi, nhưng test thì không đc như mong muốn, cái này coi như tham khảo biết đâu sau này cần và sẽ làm đc thành công. 


