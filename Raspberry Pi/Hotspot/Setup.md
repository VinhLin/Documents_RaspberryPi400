## Các bước cài đặt:

> sudo apt update 

> sudo apt install hostapd

> sudo systemctl stop hostapd

> sudo apt-get install iptables-persistent
`Yes Yes`

> sudo nano /etc/dhcpcd.conf
```
    denyinterfaces eth0
    denyinterfaces wlan0
```

> sudo nano /etc/dhcp/dhcpd.conf
```
    subnet 192.168.2.0 netmask 255.255.255.0 {
      range 192.168.2.10 192.168.2.50;
      option routers 192.168.2.200;
    }
```
> sudo nano /etc/default/isc-dhcp-server
`INTERFACESv4="eth0 wlan0"`

> sudo nano /etc/hostapd/hostapd.conf
```
	interface=wlan0
	#driver=rtl871xdrv
	ssid=Pi_Home
	country_code=US
	hw_mode=g
	channel=6
	macaddr_acl=0
	auth_algs=1
	ignore_broadcast_ssid=0
	wpa=2
	wpa_passphrase=Camerahotspot
	wpa_key_mgmt=WPA-PSK
	wpa_pairwise=CCMP
	wpa_group_rekey=86400
	ieee80211n=1
	wme_enabled=1

```

> sudo nano /etc/default/hostapd
`DAEMON_CONF="/etc/hostapd/hostapd.conf"`


> sudo nano /etc/sysctl.conf
`net.ipv4.ip_forward=1`

> sudo sh -c "echo 1 > /proc/sys/net/ipv4/ip_forward"

> sudo iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE

> sudo netfilter-persistent save

> sudo iptables -A FORWARD -i eth0 -o wlan0 -m state --state RELATED,ESTABLISHED -j ACCEPT

> sudo iptables -A FORWARD -i wlan0 -o eth0 -j ACCEPT

> sudo sh -c "iptables-save > /etc/iptables/rules.v4"

> sudo nano /etc/network/interfaces
```
    auto wlan0
    iface wlan0 inet static
      address 192.168.2.200
      netmask 255.255.255.0
      gateway 192.168.2.200
```

> sudo nano /etc/wpa_supplicant/wpa_supplicant.conf
- Comment các thiết lập network wifi.

> sudo systemctl unmask hostapd

> sudo systemctl enable hostapd

> sudo systemctl restart hostapd

> sudo reboot

### Connect
- Kết nối với wifi Pi_Home thì pass là: Camerahotspot.
> ssh pi@192.168.2.200












