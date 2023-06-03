docker network create -d macvlan -o parent=eth0 \
  --subnet 10.1.1.0/24 \
  --gateway 10.1.1.1 \
  --ip-range 10.1.1.96/28 \
  --aux-address 'host=10.1.1.4' \
macvlan_net