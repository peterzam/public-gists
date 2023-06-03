docker run -it --rm --name certbot \
-v "/tmp/certbot/etc:/etc/letsencrypt" \
-v "/tmp/certbot/lib:/var/lib/letsencrypt" \
certbot/certbot certonly --manual --preferred-challenges dns