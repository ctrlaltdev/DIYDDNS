DIY Script for Dynamic DNS using Cloudflare APIs

## INIT

You'll need your Cloudflare API Key (in your profile, on cloudflare website)

Copy `.env.sample` to `.env`:
```bash
$ cp .env.sample .env
```

Edit .env and fill with the email you use on Clouflare and your API KEY.

## RUN IT LIKE YOU MEAN IT

Now that you're set, you have to make that run: `all hail the mighty cron`

You should be able to use something like that:
```
* * * * * DIYDDNS >> /dev/null
```
