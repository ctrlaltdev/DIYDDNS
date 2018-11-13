[![Bash](https://img.shields.io/github/license/ctrlaltdev/DIYDDNS.svg?style=for-the-badge)](https:github.com/ctrlaltdev/DIYDDNS/blob/master/LICENSE)
![Bash](https://img.shields.io/badge/_-SH-4EAA25.svg?style=for-the-badge)

DIY Script for Dynamic DNS using Cloudflare APIs

## INIT

You'll need your Cloudflare API Key (in your profile, on cloudflare website)

Copy `.env.sample` to `.env`:
```bash
$ cp .env.sample .env
```

Edit .env and fill with the email you use on Clouflare and your API KEY.

### Get your DNS zone ID: 

```bash
$ . .env; \
  curl -X GET "https://api.cloudflare.com/client/v4/zones" \
    -H "X-Auth-Email: $EMAIL" \
    -H "X-Auth-Key: $CLOUDFLARE_API_KEY" \
    -H "Content-Type: application/json" \
    | python -m json.tool
```

Find the right DNS zone, and copy the id to the ZONE_ID var in `.env`.

### Get your DNS entry ID:

```bash
$ . .env; \
  curl -X GET "https://api.cloudflare.com/client/v4/zones/$ZONE_ID/dns_records" \
    -H "X-Auth-Email: $EMAIL" \
    -H "X-Auth-Key: $CLOUDFLARE_API_KEY" \
    -H "Content-Type: application/json" \
    | python -m json.tool
```

Find the right DNS entry, and copy the id and the name to DNS_ID and DNS_NAME vars in `.env`.

BAM, you're good to go.

## RUN IT LIKE YOU MEAN IT

Now that you're set, you have to make that run: `all hail the mighty cron`

You should be able to use something like that:
```
* * * * * /bin/bash /path/to/DIYDDNS.sh >> /dev/null
```