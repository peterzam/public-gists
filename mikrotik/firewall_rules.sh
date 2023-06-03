/interface bridge settings
set use-ip-firewall=yes
/ip neighbor discovery-settings
set discover-interface-list=!none
/ip cloud
set update-time=no
/ip firewall address-list
add address=10.1.1.3-10.1.1.254 list=allowed-client
/ip firewall filter
add action=accept chain=input comment="default configuartion" \
    connection-state=established,related disabled=yes
add action=accept chain=input disabled=yes src-address-list=allowed-client
add action=accept chain=input disabled=yes protocol=icmp
add action=drop chain=input disabled=yes
/ip firewall service-port
set ftp disabled=yes
/ip service
set telnet disabled=yes
set ftp disabled=yes
set www address=10.1.1.0/24
set ssh disabled=yes
set api disabled=yes
set winbox address=10.1.1.0/24
set api-ssl disabled=yes
/ip socks
set version=5
/system watchdog
set watchdog-timer=no
/tool bandwidth-server
set authenticate=no enabled=no
/tool mac-server
set allowed-interface-list=none
/tool mac-server mac-winbox
set allowed-interface-list=static
/tool mac-server ping
set enabled=no