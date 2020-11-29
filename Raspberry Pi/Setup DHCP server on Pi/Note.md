## Cac buoc thu hien:

### B1: sudo apt-get install isc-dhcp-server -y

### B2: sudo nano /etc/default/isc-dhcp-server
`INTERFACESv4="eth0"`

### B3: sudo nano /etc/dhcp/dhcpd.conf
```
subnet 192.168.10.0 netmask 255.255.255.0 {
  range 192.168.10.10 192.168.10.50;
  option routers 192.168.10.100;
}
```

### B4: sudo ifconfig eth0 192.168.10.10

### B5: sudo service isc-dhcp-server restart

### B6: sudo service isc-dhcp-server status

### B7: sudo nano /etc/network/interfaces
```
auto eth0
iface eth0 inet static
  address 192.168.10.100
  netmask 255.255.255.0
  gateway 192.168.10.100
```
