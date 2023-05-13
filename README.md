# lark go client

An unofficial Lark SDK that contains some commonly used APIs. It also includes some useful tools, such as converting Lark docx to markdown.

## Basic Usage

```go
appID, appSecret := "", ""
larkService := lark.NewService(appID, appSecret, time.Second*5)

// send msg
larkService.SendMsgByOpenID("openID", lark.NewTextMsg("hello"))
```

## References

- [oapi-sdk-go](https://github.com/larksuite/oapi-sdk-go)
- [lark](https://github.com/chyroc/lark)
- [feishu2md](https://github.com/Wsine/feishu2md)
