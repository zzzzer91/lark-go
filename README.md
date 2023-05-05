# lark go client

An unofficial SDK for Lark (飞书).

## Usage

```go
appID, appSecret := "", ""
larkService := lark.NewService(appID, appSecret, time.Second*5)
larkService.SendMsgByOpenID("openID", lark.NewTextMsg("hello"))
```

## References

- [oapi-sdk-go](https://github.com/larksuite/oapi-sdk-go)
- [lark](https://github.com/chyroc/lark)
- [feishu2md](https://github.com/Wsine/feishu2md)
