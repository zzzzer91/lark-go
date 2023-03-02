# lark go client

usage:

```go
appID, appSecret := "", ""
larkService := lark.NewService(appID, appSecret)
larkService.SendMsgByOpenID("openID", lark.NewTextMsg("hello"))
```