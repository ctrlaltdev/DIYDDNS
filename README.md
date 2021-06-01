DIY Script for Dynamic DNS using Cloudflare APIs

## INIT

You'll need your Cloudflare API Key (in your profile, on cloudflare website)

```sh
DIYDDNS -init
```
and provide your cloudflare email and api key when prompted

## RUN IT LIKE YOU MEAN IT

Now that you're set, you have to make that run: `all hail the mighty cron`

You should be able to use something like that:
```
* * * * * DIYDDNS -fqdn sub.domain.tld >> /dev/null
```
