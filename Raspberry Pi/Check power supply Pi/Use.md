### Tham khảo
> https://bia.is/2019/02/02/raspberry-pi-check-your-power-supply/ \
> https://gist.github.com/maxme/d5f000c84a4313aa531288c35c3a8887 \
> https://www.raspberrypi.org/documentation/hardware/raspberrypi/power/README.md

### Use
> sudo apt-get install sysbench \
> sudo chmod 777 raspberry-power-supply-check.sh \
> ./raspberry-power-supply-check.sh 


### Log raspberry-power-supply-check.sh 
- Mình sẽ log lại như với **Temperature**.
> touch power-supply.log \
> crontab -e
- Thêm nội dung này vào:
```
 # Log Power Supply
 */5 * * * * /home/pi/raspberry-power-supply-check.sh >> /home/pi/power-supply.log 2>&1 &
```
- Trong file *raspberry-power-supply-check.sh* mình cần chỉnh sửa lại một chút:
  * Comment lại vòng lặp *while*.
  * Thêm phần `echo "$(date) @ $(hostname)"` để hiện ngày giờ.
- Save và reboot lại Pi để xem kết quả trong file **power-supply.log**.








