# YuanMultx
Provides a route for Quantumultx profile association.

自用：在自己的服务器上使用路由用于为Quantumultx的配置关联。

修改后可以在IOS端直接更新，无需重新配置，直接从服务器拉取配置。

## 使用

### 1.添加server_remote信息


使用前需要自己修改自己的服务器配置，在serverremote/serverremote.go中

```go
svrte := &serverremote{
    Url:            "http://yourselfbooklink.com",
    Tag:            "MergeClash",
    UpdateInterval: "172800",
    OptParser:      true,
    Enable:         true,
}
```

### 2.使用端口和路由

运行后，默认使用`本地地址:39790/getconfig`可以获取默认Quantumultx配置。

或者通过`/getconfig?name=configname`方式获取配置

```txt
127.0.0.1:39790/getconfig?name=YuanMultx
```

默认的配置文件有：

Limbopro\Yatta\Orz-3\Sabrina以及YuanMultx，YuanMultx参考前三者。

## 说明

开始设想给Quantumultx所有配置项都进行一个结构化的配置，然后输出自己的配置内容。结构化的配置可以用于在不同的配置间解析并聚合。

由于自己的时间不充足以及关于Quantumultx的配置理解不够，这个计划推后，后续可能会完成。

考虑到配置文件更新非频繁，完成计划较后。

订阅连接可以使用MergeLink的连接。

