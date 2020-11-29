## Setup
[Tham khảo sản phẩm](https://hshop.vn/products/man-hinh-7-inch-hdmi-lcd-h-cam-ung-dien-dung-waveshare-co-vo-bao-ve)
- Để có thể sử dụng được màn hình LCD 7 Inch, thông qua cổng HDMI thì mình cần thực hiện một số thiết lập sau:
> sudo nano /boot/config.txt
- Thêm nội dung sau vào cuối dòng:

```
# Setup Display LCD
hdmi_force_hotplug=1
hdmi_group=2
hdmi_mode=87
config_hdmi_boost=10
hdmi_cvt 1024 600 60 6 0 0 0
```
## Note
- Ở đây mình không có thêm dòng `max_usb_current=1`, vì dòng này có tác dụng **max dòng điện output từ cổng USB**. Và từ cổng USB của Pi cấp nguồn cho LCD, nhưng nếu làm như vậy thì con Pi sẽ thường rất nóng. Do đó mình cấp **nguồn ngoài 5V** cho LCD, nên mình không cần khai báo lệnh *max_usb_current=1* trong file */boot/config.txt*.
- Tuy nhiên đối với màn hình cảm ứng, nếu cấp nguồn ngoài cho màn hình thì **chức năng cảm ứng** sẽ không dùng đc. Để dùng cảm ứng thì bắt buộc phải dùng đến cổng USB để cấp nguồn cho nó.

### Một số lệnh khác cho màn hình cảm ứng trong file /boot/config.txt
- Đảo chiều màn hình
> display_rotate=2 
- Cấp nguồn cho màn hình dùng cổng USB trên con Pi thì dùng thêm lệnh
> max_usb_current=1

### Tham khảo
> https://learn.pimoroni.com/tutorial/pi-lcd/getting-started-with-raspberry-pi-7-touchscreen-lcd \
> https://www.waveshare.com/wiki/7inch_HDMI_LCD_(H)_(with_case)



