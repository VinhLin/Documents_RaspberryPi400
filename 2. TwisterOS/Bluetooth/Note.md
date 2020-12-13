### Tham khảo
> https://quantrimang.com/cach-thiet-lap-wi-fi-va-bluetooth-tren-raspberry-pi-3-156116 \
> https://www.youtube.com/watch?v=tGQKjeeNNSw

## Thiết lập 
> sudo apt-get upgrade \
> sudo apt-get install -y bluez pulseaudio-module-bluetooth python-gobject python-gobject-2 bluez-tools sox libsox-fmt-all
- Sau đó reboot lại.
- Lúc này trên thanh công cụ có biểu tượng **Bluetooth**.
- Nhấn vào -> **Devices**. 
- Tiếp tục chọn **Seach** để tìm thiết bị Bluetooth. Thiết bị mình muốn kết nối là **BluetoothAudioV2**.
- Nhấn vào thiết bị và chọn **Create pairing with the device** (biểu tượng vòng tròn màu xanh có chữ ở giữa).
- Sau đó nhấn chọn **Setup...** -> **Audio Sink** -> **Next**
- Vậy là đã connect đc. Bước cuối là nhấn vào **biểu tượng âm thanh** trên thanh công cụ.
- Chọn **BluetoothAudioV2** để nó truyền âm thanh qua bluetooth thay vì qua *Built-in Audio Analog Mono*.
- Mở một bài hát để kiểm tra kết quả.

### Khắc phục lỗi "no connect to pulse audio server"
[Tham khảo](https://unix.stackexchange.com/questions/445386/pulseaudio-server-connection-failure-connection-refused-debian-stretch)
```
 # clean and reinstall pulseaudio
 sudo apt-get remove --purge alsa-base pulseaudio
 sudo apt-get install alsa-base pulseaudio
 sudo apt-get -f install && sudo apt-get -y autoremove && sudo apt-get autoclean && sudo apt-get clean && sudo sync && echo 3 | sudo tee /proc/sys/vm/drop_caches
 # fixes user folder permissions
 sudo chown -R $USER:$USER $HOME/
 # then reboot
 sudo reboot
```
- Lúc này kiểm tra lại biểu tượng âm thanh thì nó không còn lỗi trên nữa.