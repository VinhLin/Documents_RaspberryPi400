> https://raspberrypi.stackexchange.com/questions/169/how-can-i-extend-the-life-of-my-sd-card

- Trên thiết bị ItrackingV3 thì có 2 bộ phận lưu trữ đc chia như sau:
 * **Fash** Do sử dụng Compute Module nên bộ nhớ fash sẽ là nơi lưu trữ *OS*, *app* (như app *visiontransfer*, *visiongps*, *visionphoto*,...)
 * **SD Card** Được sử dụng để lưu các file video (độ phân giải HD). Hiện tại lưu trữ video là yếu tố bắt buộc đối với các thiết bị giám sát hành trình.

- Tình trạng thường gặp là lỗi **Read only file system**, lỗi này gặp phải là do đọc/ghi (read/write) quá nhiều.
- Lỗi này thường xảy ra trên thẻ nhớ sd (sd card). `Note: Nếu lỗi này xảy ra trên fash là coi như thiết bị đi tong, do đó ta không nên đọc/ghi quá nhiều trong bộ nhớ Fash`.
- Cách khắc phục theo  kiểu *Workaround* tạm thời là **remount** lại theo lệnh:
> sudo su \
> mount -o remount,rw /media/sdcardsecond/

## Một số cách để tránh tình trạng Read only file system
- Theo [tài liệu tham khảo](https://raymii.org/s/blog/Broken_Corrupted_Raspberry_Pi_SD_Card.html) này thì có một số cách để hạn chế việc read/write và *fix sd card*. 
`Note: Các bước thực hiện thì mình đều làm trên Raspberry. Do thẻ nhớ đc cắm sẵn trên thiết bị.`
- Thực hiện một số lệnh sau, mục đích là để kiểm tra xem ổ đĩa nào đang chứa sdcard và nó đã *mount* chưa, nếu đã đc mount thì cần phải *umount* nó thì mới có thể tiếp tục. Ở đây sdcard của mình là **/dev/mmcblk1p1**
> df -h \
> sudo su \
> umount /dev/mmcblk1p1
### Sửa chữa hệ thống tập tin tự động
[Tham khảo](https://vi.joecomp.com/fsck-command-linux)
> fsck -p /dev/mmcblk1p1
### Disable swap
- Linux divides its physical RAM (random access memory) into chucks of memory called pages. Swapping is the process whereby a page of memory is copied to the preconfigured space on the hard disk, called swap space, to free up that page of memory. The combined sizes of the physical memory and the swap space is the amount of virtual memory available.
- **Swappig causes a lot of writes to the SD card**. You would want to turn it off to save writes. The downside of this is that when there is not enough RAM available the linux OOM killer will randomly kill processes to save RAM.
> dphys-swapfile swapoff \
> dphys-swapfile uninstall \
> update-rc.d dphys-swapfile remove \
`Không sử dụng Ram ảo sẽ giúp giảm thiểu đc số lần Read/Write`
### Set up an fsck at every boot
> tune2fs -c 1 /dev/mmcblk1p1
- Mỗi lần **boot** thiết bị thì nó sẽ tự động thực hiện việc sửa chữa thẻ  nhớ.
- Để dùng đc lệnh này lúc đó mình cần mount lại thẻ nhớ.

## Note: 
- Mình có thực hiện một số cách khác theo hướng dẫn và không hiểu lắm.
- Đặc biệt là về **badblocks**, thực hiện thì mất đến hơn **2 tiếng**, nếu trong quá  trình thực hiện mà bị ngắt kết nối là hư luôn thẻ nhớ. **Từ lợn lành thành lợn què**. Cho nên mình note lại để mai mốt còn **né** nó ra.
### Find the other superblocks
`The superblock contains information about the file system such as the file system type, size, status, information about other metadata structures, block counts, inode counts, supported features, maintenance information, and more and so on`
> mke2fs -n /dev/mmcblk1p1
### The following badblocks command will scan and report bad blocks for the device. It is a destructive write operation, you will lose your data.
> badblocks -o ./badblocks.list -w -s -v -b 4096 -c 16 /dev/mmcblk1p1
- `-o` to output the badblocks list to the file ./badblocks.list, `-w` for the write operation, `-s` to show progress, `-v` to be verbose, `-b` 4096 for the blocksize of 4K and `-c` 16 to test 16 blocks at once (default is 64).
`num_read_errors, num_write_errors, num_corruption_errors`





