/ip firewall mangle
add action=passthrough chain=forward comment="Upload Global Counter" \
    disabled=yes src-address=10.1.1.0/24
add action=passthrough chain=forward comment="Download Global Counter" \
    disabled=yes dst-address=10.1.1.0/24
/system scheduler
add disabled=yes interval=5m name=ResetMangleCounters on-event=\
    ResetMangleCounter policy=\
    ftp,reboot,read,write,policy,test,password,sniff,sensitive,romon \
    start-date=nov/01/2020 start-time=00:00:00
/system script
add dont-require-permissions=no name=ResetMangleCounter owner=p policy=\
    ftp,reboot,read,write,policy,test,password,sniff,sensitive,romon source="/\
    log info (\"D \" . [/ip firewall mangle get [find where comment=\"Download\
    \_Global Counter\"] bytes])\r\
    \n/log info (\"U \" . [/ip firewall mangle get [find where comment=\"Uploa\
    d Global Counter\"] bytes])\r\
    \n/ip firewall mangle reset-counters-all"
#### Exclude below if you have no security concern
/ip cloud
set update-time=no