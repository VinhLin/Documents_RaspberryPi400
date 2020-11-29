### Tham khảo khác
> https://raspberrypi.stackexchange.com/questions/43535/eth0-doesnt-keep-its-settings-after-reboot/50742 \
> https://www.raspberrypi.org/forums/viewtopic.php?t=172163
- Đã làm theo nhưng dường như không khả quan lắm, lúc có ip lúc lại mất.

## Fix theo hướng dẫn tại: https://samhobbs.co.uk/2013/11/fix-for-ethernet-connection-drop-on-raspberry-pi?page=1
> mkdir ~/bin \
> nano ~/bin/network-monitor.sh
- Thêm nội dung sau:

```
#!/bin/bash

LOGFILE=/home/pi/network-monitor.log

if ifconfig eth0 | grep -q "inet addr:" ;
then
        echo "$(date "+%m %d %Y %T") : Ethernet OK" >> $LOGFILE
else
        echo "$(date "+%m %d %Y %T") : Ethernet connection down! Attempting reconnection." >> $LOGFILE
        ifup --force eth0
        OUT=$? #save exit status of last command to decide what to do next
        if [ $OUT -eq 0 ] ; then
                STATE=$(ifconfig eth0 | grep "inet addr:")
                echo "$(date "+%m %d %Y %T") : Network connection reset. Current state is" $STATE >> $LOGFILE
        else
                echo "$(date "+%m %d %Y %T") : Failed to reset ethernet connection" >> $LOGFILE
        fi
fi
```
- Tiếp tục thực hiện thao tác:
> PATH=$PATH:~/bin \
> source ~/.bashrc \
> chmod +x ~/bin/network-monitor.sh
### Test thử 
> sudo ./bin/network-monitor.sh
- Dùng lệnh `tail -f ~/network-monitor.log` để đọc log xem Ethernet có OK không.
### Automating with cron
> sudo nano /etc/crontab
- Thêm nội dung sau vào cuối dòng:
`*/5 * * * * root bash /home/pi/bin/network-monitor.sh`
- Sau đó **reboot** lại Pi và kiểm tra file *network-monitor.log*








