# lark go client

An unofficial SDK for Lark (飞书).

usage:

```go
appID, appSecret := "", ""
larkService := lark.NewService(appID, appSecret, time.Second*5)
larkService.SendMsgByOpenID("openID", lark.NewTextMsg("hello"))
```