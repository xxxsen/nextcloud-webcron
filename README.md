nextcloud-webcron
===

nextcloud的webcron

使用方法:

```yaml
version: '2'
services:
  crontask:
    image: xxxsen/nextcloud-webcron:latest
    environment:
      - EXPRESSION=*/5 * * * *
      - URL=https://mynextcloud.com/cron.php
      - RUN_ON_START=true
      - LOG_LEVEL=debug
```

