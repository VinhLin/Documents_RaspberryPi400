### Mình sử dụng UpSwift để thực hiện OTA updates.
[ThamKhao1](https://www.upswift.io/post/raspberry-pi-ota-update-in-2020-method)
[ThamKhao2](https://medium.com/@eitan.chud/how-to-implement-ota-updates-for-iot-and-connected-devices-82624cb6a8bf)

- Trước hết đăng ký tài khoản trên UpSwift 
```
  login: linducvinh@gmail.com 
  pass: linducvinh1412
```
- Với mỗi tài khoản free thì được đăng ký tối đa 3 thiết bị.
- Cài đặt trên Pi (như trong Remote Raspberry/Remote Online/UpSwift)
- Lên trang **https://dashboard.upswift.io/updates/**, chọn **New Micro Update**
- Sau đó điền các thông tin cần thiết
Ví dụ như cần update file code vision_transfer cho raspberry trên tất cả các thiết bị thì ta phải thực hiện:
```
> Apply micro update on: Production | All Devices
> Micro update version: 1.0.04
> Execute before update: systemctl stop visiontransfer.service
**Trước khi cập nhật cần stop service visiontransfer**
> Files: *File code vision_transfer cần cập nhật*
> Destination Path: /usr/bin
**Đường dẫn đến file code thực thi trên Raspberry**
> Execute after update: systemctl restart visiontransfer.service
**Sau khi cập nhật thì restart service visiontransfer**
> Deploy update
```
- Nó sẽ hiện tất cả các thiết bị và kết quả update thành công (Succeed) hay thất bại (Errors).

#### Note: Ngoài ra trong quá trình update, mình có thử rút nguồn Pi, sau đó cắm nguồn lại và có mạng internet thì quá trình update vẫn diễn ra và hoàn thành.

