Link:
> https://www.digitalocean.com/community/questions/ubuntu-cron-job-reboot-every-4-hours
> https://crontab.guru/#*/2_*_*_*_*


### Auto Reboot sau mỗi 8 tiếng
> sudo crontab -e
- Thêm nội dung
`0 */8 * * * /sbin/shutdown -r now`
- Sau đó save và thoát ra mà thấy có dòng *crontab: installing new crontab* tức là mình đã thiết lập crontab thành công.


### Test auto reboot after 2 minute
*/2 * * * * /sbin/shutdown -r now


### Auto Reboot At 01:00 on every day-of-week from Monday through Sunday.
[Link](https://crontab.guru/#0_1_*_*_1-7)
0 1 * * 1-7 /sbin/shutdown -r now





### List user's crontab
> sudo crontab -l

