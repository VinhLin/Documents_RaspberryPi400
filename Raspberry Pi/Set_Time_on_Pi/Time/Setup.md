- Để set time trên Raspberry thì có 2 cách.
- Sử dụng **date** hoặc **timedatectl**


## Cách 1: Dùng date
[Document](https://www.hivelocity.net/kb/how-to-change-date-time-zone-on-linux-server/)
### Example
> sudo su \
> date -s "23 June 2013 18:00:00"


## Cách 2: Dùng timedatectl
[Doc](https://www.maketecheasier.com/timedatectl-control-system-time-date-linux/)
### Example
> sudo timedatectl set-time "2014-11-08 06:40:00" \

