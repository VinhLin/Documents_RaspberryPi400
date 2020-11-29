[Tham khảo](https://www.jamescoyle.net/knowledge/951-the-difference-between-a-tmpfs-and-ramfs-ram-disk)
`There are two file system types built into most modern Linux distributions which allow you to create a RAM based storage area which can be mounted and used link a normal folder.`

### What is a memory based file system (RAM disk)?
- A memory based file system is something which creates a storage area directly in a computers RAM as if it were a partition on a disk drive. As RAM is a volatile type of memory which means when the system is restarted or crashes the file system is lost along with all it’s data.
- The two main RAM based file system types in Linux are tmpfs and ramfs.
 * **ramfs** là loại hệ thống tệp cũ hơn và được thay thế phần lớn trong hầu hết các trường hợp bằng **tmpfs**.

### tmpfs
- tmpfs là một hệ thống tệp RAM mới hơn, khắc phục nhiều nhược điểm với ramfs. 
- Bạn có thể chỉ định giới hạn kích thước trong tmpfs, điều này sẽ gây ra lỗi "disk full" khi đạt đến giới hạn. Hành vi này giống hệt như một phân vùng của đĩa vật lý.
- Kích thước và dung lượng đã sử dụng trên phân vùng *tmpfs* cũng được hiển thị trong *df -h*

## Create a RAM disk in Linux use tmpfs
[Tham Khảo](https://www.jamescoyle.net/how-to/943-create-a-ram-disk-in-linux)
- Kiểm tra xem dung lượng *Ram* còn lại trên máy là bao nhiêu bằng lệnh:
> free -g
- Tạo một folder để mount Ram disk.
> mkdir /mnt/ramdisk
- Then use the mount command to create a RAM disk.
> mount -t [TYPE] -o size=[SIZE] [FSTYPE] [MOUNTPOINT]

```
 Substitute the following attirbutes for your own values:
   - [TYPE] is the type of RAM disk to use; either tmpfs or ramfs.
   - [SIZE] is the size to use for the file system. Remember that ramfs does not have a physical limit and is specified as a starting size.
   - [FSTYPE] is the type of RAM disk to use; either tmpfs, ramfs, ext4, etc.
```
- Example: `mount -t tmpfs -o size=512m tmpfs /mnt/ramdisk`
- Bạn có thể thêm mục mount vào **/etc/fstab** để tạo RAM disk sau khi khởi động lại. Tuy nhiên, hãy nhớ rằng dữ liệu sẽ biến mất mỗi khi máy được khởi động lại.
> nano /etc/fstab
> tmpfs /mnt/ramdisk tmpfs nodev,nosuid,noexec,nodiratime,size=1024M 0 0


