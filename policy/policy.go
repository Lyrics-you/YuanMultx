package policy

type group struct {
	Name           string   `conf:"static"`
	Option         []string `conf:"static"`
	ServerTagRegex string   `conf:"server-tag-regex"`
	ImgUrl         string   `conf:"img-url"`
}
type rule struct {
	UrlLatencyBenchmark string `conf:"url-latency-benchmark"`
	ServerTagRegex      string `conf:"server-tag-regex="`
	CheckInterval       uint32 `conf:"check-interval"`
	Tolerance           uint32 `conf:"tolerance"`
	ImgUrl              string `conf:"img-url"`
}
type policy struct {
	Group []group
	Rule  []rule
}

func PolicyInit() *policy {
	ply := &policy{
		Group: []group{
			{Name: "NetEaseMusic",
				Option:         []string{"direct", "proxy"},
				ServerTagRegex: "(?=.*(music|𝐌𝐮𝐬𝐢𝐜|Unbolck|网易云|云音乐|Music|Netease|🎶|解锁))",
				ImgUrl:         "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Netease_Music.png"},
			{Name: "漏网之鱼",
				Option: []string{"direct", "best", "proxy"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Final.png"},
			{Name: "国内网站",
				Option: []string{"direct", "proxy"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Domestic.png"},
			{Name: "境外网站",
				Option: []string{"best", "proxy", "direct"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Global.png"},
			{Name: "国内流媒体",
				Option: []string{"direct", "proxy"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/DomesticMedia.png"},
			{Name: "国外流媒体",
				Option: []string{"best", "proxy", "direct"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/ForeignMedia.png"},
			{Name: "数字货币",
				Option: []string{"best", "proxy", "direct"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Cryptocurrency_3.png"},
			{Name: "Instagram",
				Option: []string{"best", "proxy", "🇭🇰Hong Kong", "🇨🇳Taiwan", "🇯🇵Japan", "🇰🇷Korea", "🇸🇬Singapore", "direct"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Instagram.png"},
			{Name: "TikTok",
				Option: []string{"direct", "proxy", "🇺🇸United States", "🇯🇵Japan", "🇰🇷Korea"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/TikTok_1.png"},
			{Name: "Youtube",
				Option: []string{"best", "direct", "🇭🇰Hong Kong", "🇯🇵Japan"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/YouTube.png"},
			{Name: "Pornhub",
				Option: []string{"best", "direct", "proxy", "🇭🇰Hong Kong", "🇨🇳Taiwan", "🇯🇵Japan", "🇰🇷Korea", "🇸🇬Singapore", "🇺🇸United States"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Pornhub_1.png"},
			{Name: "Spotify",
				Option: []string{"best", "direct", "proxy", "🇭🇰Hong Kong", "🇨🇳Taiwan", "🇯🇵Japan", "🇰🇷Korea", "🇸🇬Singapore", "🇺🇸United States"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Spotify.png"},
			{Name: "Netflix",
				Option: []string{"best", "direct", "proxy", "🇭🇰Hong Kong", "🇨🇳Taiwan", "🇯🇵Japan", "🇰🇷Korea", "🇸🇬Singapore", "🇺🇸United States"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Netflix_Letter.png"},
			{Name: "Spotify",
				Option: []string{"best", "direct", "proxy", "🇭🇰Hong Kong", "🇨🇳Taiwan", "🇯🇵Japan", "🇰🇷Korea", "🇸🇬Singapore", "🇺🇸United States"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Spotify.png"},
			{Name: "Telegram",
				Option: []string{"best", "direct", "proxy", "🇭🇰Hong Kong", "🇨🇳Taiwan", "🇯🇵Japan", "🇰🇷Korea", "🇸🇬Singapore", "🇺🇸United States"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Telegram.png"},
			{Name: "Google",
				Option: []string{"best", "proxy", "direct", "🇭🇰Hong Kong", "🇨🇳Taiwan", "🇯🇵Japan", "🇰🇷Korea", "🇸🇬Singapore", "🇺🇸United States"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Google_Search.png"},
			{Name: "Steam",
				Option: []string{"🇭🇰Hong Kong", "direct", "proxy"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Steam.png"},
			{Name: "PayPal",
				Option: []string{"direct", "proxy", "🇺🇸United States"},
				ImgUrl: "=https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/PayPal.png"},
			{Name: "Advertising",
				Option: []string{"direct", "reject"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Advertising.png"},
			{Name: "Apple Update",
				Option: []string{"direct", "reject"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Apple_Update.png"},
			{Name: "Apple",
				Option: []string{"direct", "proxy", "🇺🇸United States"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Apple.png"},
			{Name: "Microsoft",
				Option: []string{"direct", "proxy"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Microsoft.png"},
			{Name: "Speedtest",
				Option: []string{"direct", "proxy"},
				ImgUrl: "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Speedtest.png"},
		},
		Rule: []rule{
			{UrlLatencyBenchmark: "best",
				ServerTagRegex: "^((?!(music|𝐌𝐮𝐬𝐢𝐜|Unbolck|网易云|云音乐|Music|Netease|🎶|专线|手游|游戏|(?i)IPLC|IEPL|game)).)*$",
				CheckInterval:  1800,
				Tolerance:      10,
				ImgUrl:         "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Auto.png",
			},
			{UrlLatencyBenchmark: "专线",
				ServerTagRegex: "(手游|游戏|专线|(?i)IPLC|IEPL|game)",
				CheckInterval:  1800,
				Tolerance:      10,
				ImgUrl:         "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Server.png",
			},
			{UrlLatencyBenchmark: "🇭🇰Hong Kong",
				ServerTagRegex: "(?=.*(香港|HK|(?i)Hong))^((?!(专线|手游|游戏|(?i)IPLC|IEPL|game)).)*$",
				CheckInterval:  1800,
				Tolerance:      10,
				ImgUrl:         "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Hong_Kong.png",
			},
			{UrlLatencyBenchmark: "🇨🇳Taiwan",
				ServerTagRegex: "(?=.*(台湾|台灣|TW|(?i)Taiwan))^((?!(专线|手游|游戏|(?i)IPLC|IEPL|game)).)*$",
				CheckInterval:  1800,
				Tolerance:      10,
				ImgUrl:         "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Taiwan.png",
			},
			{UrlLatencyBenchmark: "🇯🇵Japan",
				ServerTagRegex: "(?=.*(日本|JP|(?i)Japan))^((?!(专线|手游|游戏|(?i)IPLC|IEPL|game)).)*$",
				CheckInterval:  1800,
				Tolerance:      10,
				ImgUrl:         "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Japan.png",
			},
			{UrlLatencyBenchmark: "🇰🇷Korea",
				ServerTagRegex: "(?=.*(韩国|韓國|南朝鲜|KR|(?i)Korean))^((?!(专线|手游|游戏|(?i)IPLC|IEPL|game)).)*$",
				CheckInterval:  1800,
				Tolerance:      10,
				ImgUrl:         "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Korea.png",
			},
			{UrlLatencyBenchmark: "🇸🇬Singapore",
				ServerTagRegex: "(?=.*(新加坡|狮城|SG|(?i)Singapore))^((?!(专线|手游|游戏|(?i)IPLC|IEPL|game)).)*$",
				CheckInterval:  1800,
				Tolerance:      10,
				ImgUrl:         "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/Singapore.png",
			},
			{UrlLatencyBenchmark: "🇺🇸United States",
				ServerTagRegex: "(?=.*(美国|美國|US|(?i)States|American))^((?!(专线|手游|游戏|网易云|云音乐|🎶(?i)IPLC|IEPL|game|music|𝐌𝐮𝐬𝐢𝐜|Unbolck|Music|Netease)).)*$",
				CheckInterval:  1800,
				Tolerance:      10,
				ImgUrl:         "https://cdn.jsdelivr.net/gh/Koolson/Qure@master/IconSet/Color/United_States.png",
			},
		},
	}
	return ply

}
