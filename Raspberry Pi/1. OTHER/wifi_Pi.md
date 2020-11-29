### Tham khảo
> https://raspberrypi.vn/cai-dat-wifi-tren-raspberry-6843.pi

### Command
> sudo nano /etc/wpa_supplicant/wpa_supplicant.conf
- Thêm nội dụng. Với **ssid** là tên wifi; **psk** là pass wifi.
```
network={
 ssid="MyNetworkSSID"
 psk="Pa55w0rd1234"
}
```
