/certificate
add name=ca-template days-valid=3650 common-name=hapac2.local key-usage=key-cert-sign,crl-sign
add name=server-template days-valid=3650 common-name=hapac2.local

/certificate
sign ca-template name=root-ca
:delay 3s
sign ca=root-ca server-template name=server
:delay 3s

/certificate
set root-ca trusted=yes
set server trusted=yes

/ip service
set www-ssl certificate=server disabled=no

# curl -k -u admin:password -X POST https://10.1.1.1/rest/system/script/run --data '{".id":"*1"}' -H "content-type: application/json"