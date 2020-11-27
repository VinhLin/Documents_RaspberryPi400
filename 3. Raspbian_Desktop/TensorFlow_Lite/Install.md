# TensorFlow_Lite
- **TensorFlow_Lite** là bản rút gọn của *TensorFlow*, nó có thể cài đặt và chạy trên Pi, đỡ lag hơn rất nhiều.
[Video so sánh](https://www.youtube.com/watch?v=TiOKvOrYNII)
- Cài đặt **TensorFlow_Lite** trên *Raspberry Pi 3+*.
[Youtube](https://www.youtube.com/watch?v=aimSGOAUI8Y)
[Github](https://github.com/EdjeElectronics/TensorFlow-Lite-Object-Detection-on-Android-and-Raspberry-Pi)
[Note](https://github.com/EdjeElectronics/TensorFlow-Lite-Object-Detection-on-Android-and-Raspberry-Pi/blob/master/Raspberry_Pi_Guide.md)

### Tham khảo thêm
> https://www.tensorflow.org/lite/guide/python \
> https://coral.ai/docs/accelerator/get-started/#requirements

### Pip3 on Raspberry
`pip is installed by default in Raspberry Pi OS Desktop images.`
- Do mình đang test thử trên bản **Raspberry Lite** nên nó không có sẵn **pip3** và mình phải cài đặt thêm.
[Tham khảo 1](https://www.raspberrypi.org/documentation/linux/software/python.md)
[Tham khảo 2](https://gallaugher.com/makersnack-installing-circuitpython-on-a-raspberry-pi/)
> sudo apt-get install python3-pip 
- Upgrade pip3 on Pi
> sudo pip3 install --upgrade setuptools

### python3-venv
- Cài đặt thêm **python3-venv**. [Tham khảo](https://www.techcoil.com/blog/how-to-use-python-3-virtual-environments-to-run-python-3-applications-on-your-raspberry-pi/)
> sudo apt-get install python3-venv -y

## Install
> sudo apt-get update \
> sudo apt-get upgrade \
> sudo apt-get dist-upgrade
- Download repository và thực hiện một số lệnh:
> git clone https://github.com/EdjeElectronics/TensorFlow-Lite-Object-Detection-on-Android-and-Raspberry-Pi.git \
> mv TensorFlow-Lite-Object-Detection-on-Android-and-Raspberry-Pi tflite1 \
> cd tflite1 \
> sudo pip3 install virtualenv \
> python3 -m venv tflite1-env \
> source tflite1-env/bin/activate \
> bash get_pi_requirements.sh

## Run
- **Set up TensorFlow Lite detection model** thì có 2 cách. Một là sử dụng *TFLite model* do **Google** cung cấp; hai là *Sử dụng mô hình được đào tạo tùy chỉnh của riêng bạn*.
- Mình sẽ làm thử demo **Option 1: Using Google's sample TFLite model**
[Link](https://www.tensorflow.org/lite/models/object_detection/overview)
> wget https://storage.googleapis.com/download.tensorflow.org/models/tflite/coco_ssd_mobilenet_v1_1.0_quant_2018_06_29.zip \
> unzip coco_ssd_mobilenet_v1_1.0_quant_2018_06_29.zip -d Sample_TFLite_model
- Run the TensorFlow Lite model. Một số lệnh **TFLite_detection** trong thư mục:
#### Webcam
> python3 TFLite_detection_webcam.py --modeldir=Sample_TFLite_model
#### Stream
> python3 TFLite_detection_stream.py --modeldir=Sample_TFLite_model --streamurl="http://ipaddress:port/stream/video.mjpeg" --resolution=1920x1080
#### Video
> python3 TFLite_detection_video.py --modeldir=Sample_TFLite_model --video=test.mp4
#### Image
- Image Dir:
> python3 TFLite_detection_image.py --modeldir=Sample_TFLite_model --imagedir=squirrels 
- Image file: 
> python3 TFLite_detection_image.py --modeldir=Sample_TFLite_model --image=test1.jpg
 
 



