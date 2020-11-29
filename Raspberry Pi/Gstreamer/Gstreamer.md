### Setup on Raspberry
> sudo apt-get update
> sudo apt-get upgrade
> sudo apt-get -y install libgstreamer1.0-dev
> sudo apt-get -y install libgstreamer-plugins-base1.0-dev
> sudo apt-get -y install libglib2.0-dev
> sudo apt-get -y install gstreamer1.0-libav
> sudo apt-get -y install gstreamer1.0-tools
> sudo apt-get -y install gstreamer1.0-plugins-good 
> sudo apt-get -y install gstreamer1.0-plugins-bad
> sudo apt-get -y install gstreamer1.0-plugins-ugly
> sudo apt-get -y install libgstreamer1.0-dev libgstreamer-plugins-base1.0-dev


### Display all installed packages with gstreamer in the package name
> dpkg -l | grep gstreamer

### Version 
> gst-launch-1.0 --gst-version

Link:
> https://askubuntu.com/questions/769413/how-do-i-check-gstreamer-version


## Install on Snap
- Gstreamer có thể đc cài đặt bằng (mình chưa thử cài). Do ban đầu mình đã cài bằng cách trên rồi.
[Tham khảo](https://snapcraft.io/install/gstreamer/debian)
### Install Snap On Debian 9 (Stretch) and newer.
> sudo apt update \
> sudo apt install snapd \
> sudo snap install core
### Install gstreamer
> sudo snap install gstreamer --edge

