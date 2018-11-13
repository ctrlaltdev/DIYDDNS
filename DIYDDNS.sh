#!/bin/sh

. ./.env

DIGIP=$(dig +short $DNS_NAME)
WANIP=$(dig +short myip.opendns.com @resolver1.opendns.com)

if [ "$DIGIP" != "$WANIP" ]
then
  curl -X PUT "https://api.cloudflare.com/client/v4/zones/$ZONE_ID/dns_records/$DNS_ID" \
    -H "X-Auth-Email: $EMAIL" \
    -H "X-Auth-Key: $CLOUDFLARE_API_KEY" \
    -H "Content-Type: application/json" \
    --data '{"type":"A","name":"'"$DNS_NAME"'","content":"'"$WANIP"'","proxied":"false","ttl":"1"}' \
    --silent
fi
