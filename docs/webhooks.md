#### Webhooks: Quick reference

Ref: [Telegram Bot API: Getting Updates](https://core.telegram.org/bots/api#getting-updates)

##### Getting webhook info

```
curl -X GET "https://api.telegram.org/bot{BOT_TOKEN}/getWebhookInfo"
```


##### Setting webhook 

```
curl -X POST "https://api.telegram.org/bot{BOT_TOKEN}/setWebhook" \
     -d "url=https://yourserver.com/bot" \
     -d "allowed_updates=[\"message\",\"callback_query\"]"
```
