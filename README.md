DIY Script for Dynamic DNS using Cloudflare APIs

## INSTALLATION

### HomeBrew (only for macOS and linux amd64 and arm64)

```sh
brew install ctrlaltdev/tap/diyddns
```
or
```sh
brew tap ctrlaltdev/tap
brew install diyddns
```

### Easy Shell Script

```sh
curl -fSsL https://git.io/JG0Il | sh
```

It will prompt you for your OS and ARCH to download and install the right version - it will require sudo to install the binary to /usr/local/bin

### DIY Shell Script

```sh
echo -n "What is your OS? [darwin/linux] "
read os
echo -n "What is your ARCH? [amd64/arm64/armv7/armv6/armv5] "
read arch

curl -o DIYDDNS-$os-$arch.tar.gz -sL https://github.com/ctrlaltdev/DIYDDNS/releases/download/v2.0.0/DIYDDNS-$os-$arch.tar.gz
curl -o DIYDDNS-$os-$arch.tar.gz.sha256 -sL https://github.com/ctrlaltdev/DIYDDNS/releases/download/v2.0.0/DIYDDNS-$os-$arch.tar.gz.sha256
sha256sum -c DIYDDNS-$os-$arch.tar.gz.sha256

tar xzf DIYDDNS-$os-$arch.tar.gz

rm DIYDDNS-$os-$arch.tar.gz*

sudo mv DIYDDNS /usr/local/bin/

echo "\nDIYDDNS INSTALLED\n"

DIYDDNS -h
```

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
