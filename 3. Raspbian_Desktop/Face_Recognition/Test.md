### Link
> https://www.youtube.com/watch?v=CObCmwFnq7k \
> https://tutorial.cytron.io/2020/11/13/face-recognition-using-opencv-on-raspberry-pi-400/
- Phần tham khảo *Youtube tiếng Việt*, mình thấy cách làm của nó cũng tương tự như link trên. Được cái là nó giải thích ý nghĩa code trên file luôn.

- Mình test trên con Raspberry CM3+ (bản Desktop). 
- Dùng Camera IP, nên sẽ phải cần sửa lại 1 chút trong code.
### Cách thực hiện test
[Tham khảo cách train](https://www.tomshardware.com/how-to/raspberry-pi-facial-recognition)
- Mình cần cài đặt 1 số thứ cho Pi:
> sudo apt install cmake build-essential pkg-config git \
> sudo apt install libjpeg-dev libtiff-dev libjasper-dev libpng-dev libwebp-dev libopenexr-dev \
> sudo apt install libavcodec-dev libavformat-dev libswscale-dev libv4l-dev libxvidcore-dev libx264-dev libdc1394-22-dev libgstreamer-plugins-base1.0-dev libgstreamer1.0-dev \
> sudo apt install libgtk-3-dev libqtgui4 libqtwebkit4 libqt4-test python3-pyqt5 \
> sudo apt install libatlas-base-dev liblapacke-dev gfortran \
> sudo apt install libhdf5-dev libhdf5-103 \
> sudo apt install python3-dev python3-pip python3-numpy
- Cài đặt OpenCV cho Pi bằng **pip**. Nhớ lưu ý là phải cài **bản 3.4.6.27** vì nếu cài *bản 4* trở đi thì code bị lỗi
> pip3 install opencv-contrib-python==3.4.6.27
- Git clone:
> git clone https://github.com/carolinedunn/facial_recognition
- Tạo thư mục tên mình (ví dụ như tạo thư mục tên là Vinh) để chuẩn bị cho việc **train**
> cd facial_recognition/dataset \
> mkdir Vinh \
> cd ..
- Sửa lại 1 chút trong code **headshots.py** như hình *headshots.png*
- Sửa phần **name** cho đúng thư mục tên mình cần train, sửa lại đường dẫn camera vì mình dùng *camera IP* (ví dụ ở đây là: rtsp://192.168.10.117:554/user=admin_password=tlJwpbo6_channel=1_stream=1.sdp?real_stream)
- Chạy code thì dùng lệnh `python3 headshots.py`. **Nhớ nhấn dấu cách để cho có thể chụp lại ảnh và lưu lại trong thư mục Vinh**.
- Vẫn trong thư mục *facial_recognition* bằng sẽ chạy code train:
> python3 train_model.py
- Khi chạy code mình bị lỗi **No module named imutils**, lỗi này là do mình chưa cài đặt thư viện **imutils**, mình chỉ cần cài đặt là ok.
> pip3 install imutils
- Chạy lại train. Sau khi train xong, bước cuối cùng là test lại với code **facial_req.py**, sửa lại phần **VideoStream** như hình *facial_req.png*
- Chạy code: `python3 facial_req.py`
#### Mình đã thực hiện được thành công như trong video. File facial_recognition là file code mình thực hiện trên con Pi.


### Note
- Độ trễ khi thực hiện đoạn code này chỉ có **vài giây**.
- Dùng htop để kiểm tra thì thấy gần như sử dụng **full CPU** (400%)

