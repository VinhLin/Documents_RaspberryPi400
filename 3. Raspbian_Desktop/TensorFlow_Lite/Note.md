- Khi nội dung vào file code python mà nó báo lỗi `inconsistent use of tabs and spaces in indentation` thì cần phải xem code của người ta viết là dùng **tab** hay **space** để thụt vào cho phù hợp.

### Tham khảo thêm về TensorFlow Lite
> https://github.com/google/mediapipe \
> https://github.com/margaretmz/awesome-tensorflow-lite \
> https://google.github.io/mediapipe/solutions/models \
> https://github.com/google/mediapipe/tree/master/mediapipe/models \
> https://www.tensorflow.org/lite/guide/python \
> https://github.com/tensorflow/models/tree/master/official \
> https://blog.tensorflow.org/search?label=TensorFlow+Lite&max-results=20 \


### TensorFlow Lite setup Ubuntu 20.04
> https://www.andreaamico.eu/other/2020/07/27/tf-setup.html

### TensorFlow Lite Converter
> https://github.com/tensorflow/tensorflow/blob/master/tensorflow/lite/g3doc/r1/convert/cmdline_examples.md \


-----------------------------------------------------------------------------------------------------------------
> cd tflite1/ \
> source tflite1-env/bin/activate \
> python3 TFLite_detection_image.py --modeldir= --graph= --labels=  --image=

python3 TFLite_detection_image.py --modeldir=Sample_TFLite_model/ --graph=detect.tflite --labels=labelmap.txt --threshold=2 --image=test1.jpg
python3 TFLite_detection_image.py --modeldir=Sample_TFLite_model/ --graph=detectiontest.tflite --labels=labelmaptest.txt --threshold=0 --image=1603509447.jpg 

python3 TFLite_detection_stream.py --modeldir=models/ --streamurl="rtsp://192.168.10.71:554/user=admin_password=tlJwpbo6_channel=1_stream=1.sdp?real_stream" --graph=face_landmark.tflite --labels=face_detection_front_labelmap.txt --resolution=704x576


