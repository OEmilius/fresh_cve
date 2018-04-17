# fresh_cve
Periodically get fresh cve from cve.circl.lu &amp; access.redhat.com

## Installation
go get github.com/OEmilius/fresh_cve

## How it work
- read config.json
- init cache
- load old cache
- init LOADER
- start web server
- periodically check fresh cve
