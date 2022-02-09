package dns

type dns struct {
	Tag     string   `conf:"no-ipv6"`
	Server  []string `conf:"server"`
	Address string   `conf:"address"`
}

func DnsInit() *dns {
	ds := &dns{
		Tag: "no-ipv6",
		Server: []string{
			"119.29.29.29",
			"223.5.5.5",
			"1.2.4.8",
			// "114.114.114.114",
			// "8.8.8.8",
			"202.141.162.123",

			"/dl.google.com/119.29.29.29",
			"/dl.l.google.com/119.29.29.29",
			"/update.googleapis.com/119.29.29.29",
			"/*.dl.playstation.net/119.29.29.29",

			"/falogin.cn/system",
			"/melogin.cn/system",
			"/tplinklogin.net/system",
			"/tplogin.cn/system",

			"/aaplimg.com/119.29.29.29",
			"/apple.com/119.29.29.29",
			"/apple.com.cn/119.29.29.29",
			"/apple-mapkit.com/119.29.29.29",
			"/icloud.com/119.29.29.29",
			"/icloud.com.cn/119.29.29.29",
			"/mzstatic.com/119.29.29.29",
			"/*.aaplimg.com/119.29.29.29",
			"/*.apple.com/119.29.29.29",
			"/*.apple-cloudkit.com/119.29.29.29",
			"/*.apple-mapkit.com/119.29.29.29",
			"/*.cdn-apple.com/119.29.29.29",
			"/*.icloud.com/119.29.29.29",
			"/*.icloud.com.cn/119.29.29.29",
			"/*.mzstatic.com/119.29.29.29",
			"/*apple.com.akadns.net/119.29.29.29",
			"/*.apple.com.edgekey.net/119.29.29.29",
			"/*.apple.com.edgekey.net.globalredir.akadns.net/119.29.29.29",
			"/*icloud.com.akadns.net/119.29.29.29",

			"/*.taobao.com/223.5.5.5",
			"/*.tmall.com/223.5.5.5",
			"/*.alipay.com/223.5.5.5",
			"/*.alicdn.com/223.5.5.5",
			"/*.aliyun.com/223.5.5.5",

			"/*.jd.com/119.28.28.28",
			"/*.qq.com/119.28.28.28",
			"/*.tencent.com/119.28.28.28",
			"/*.weixin.com/119.28.28.28",

			"/*.bilibili.com/119.29.29.29",
			"/hdslb.com/119.29.29.29",
			"/*.163.com/119.29.29.29",
			"/*.126.com/119.29.29.29",
			"/*.126.net/119.29.29.29",
			"/*.127.net/119.29.29.29",
			"/*.netease.com/119.29.29.29",
			"/*.mi.com/119.29.29.29",
			"/*.xiaomi.com/119.29.29.29",
		},
		Address: "/mtalk.google.com/108.177.125.188",
	}
	return ds
}
