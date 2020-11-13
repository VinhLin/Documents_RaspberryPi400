- Trải nghiệm ban đầu thì **Không được hoàn hảo như mong đợi** cho lắm.
- Lý do vì có không thể cài đặt đc **VSCode** hoặc **Golang** bằng **Snap** cho **ARM64** được. Vả lại nếu mở nhiều **tag trên Firefox** thì sẽ gây ra hiện tượng rất **lag** khi sử dụng.
- Nhưng nếu chỉ nói về các tính **giải trí cơ bản** thì con Pi400 đáp ứng khá tốt.

```
 sudo apt-get update
 sudo apt-get upgrade
```
## Một số tools cài đặt ban đầu sau khi Install Ubuntu 20.10
> sudo apt install net-tools \
> sudo apt install screen \
> sudo apt-get install arp-scan \
> sudo apt-get install wvdial \
> sudo apt install speedtest-cli \
> sudo apt-get install openssh-server \
> sudo apt install git \
> sudo apt-get install ibus-unikey 
### Một số tool được cài bằng Snap
> sudo snap install mdview \
> sudo snap install htop
### Note
- Có rất nhiều tools không được cài bằng **Snap trên ARM64** như: code *(VSCode)*, go *(Golang)*, vlc, rambox
- Do gặp tình trạng lag nên mình bỏ ý định cài **VSCode** và **Golang**. Còn **Rambox** thì không cài được.
- Riêng **vlc** mặc dù không cài đc bằng snap, nhưng mình có thể cài bằng cách khác. [Tham khảo](https://linuxhint.com/best_10_video_players_linux/) trang này thì mình thực hiện lệnh sau:
> sudo apt-get install vlc qtwayland5
- Trải nghiệm **vlc** cho **movie full HD** thì **tuyệt**. Không giống như con Pi 3+ hồi trước, xem phim không nổi với nó.




