### Theo dõi nhiệt độ của Raspberry
[Link](https://linuxhint.com/raspberry_pi_temperature_monitor/)

- Dùng lệnh đơn giản sau:
> vcgencmd measure_temp


### Để tự động chạy file my-pi-temp.sh và lưu lại thì thực hiện các bước sau:
- Tạo 1 file để lưu lại nhiệt độ: 
> touch my-pi-temp.log
- Sử dụng **crontab** để cứ 5 phút sẽ chạy 1 lần:
> crontab -e
- Thêm nội dung sau vào cuối dòng:
`*/5 * * * * nohup /home/pi/my-pi-temp.sh >> /home/pi/my-pi-temp.log 2>&1 &`







