package general

type general struct {
	// QuanX会对server_check_url指定的网址进行相应测试，以确认节点的可用性
	Server_check_url []string `conf:"server_check_url"`
	// 网络检查
	Network_check_url []string `conf:"network_check_url"`
	// list中的域名将不使用fake-ip方式, 多个域名用“, ”连接。其它域名则全部采用 fake-ip及远程解析的模式
	Dns_exclusion_list []string `conf:"dns_exclusion_list"`
	// 用于获取及展示节点信息
	Geo_location_checker []string `conf:"geo_location_checker"`
	// 资源解析器，自定义各类远程资源的转换
	Resource_parser_url []string `conf:"resource_parser_url"`
	// 下列路径将不经过QuanX的处理
	Excluded_routes []string `conf:"excluded_routes"`
	// UDP策略
	Fallback_udp_policy string `conf:"fallback_udp_policy"`
	// profile图片
	Profile_img_url string `conf:"profile_img_url"`
}

func GeneralInit() *general {
	gerl := &general{
		Server_check_url:  []string{"http://cp.cloudflare.com/generate_204"},
		Network_check_url: []string{"http://www.qualcomm.cn/generate_204"},
		Dns_exclusion_list: []string{
			"*.cmpassport.com",
			"*.jegotrip.com.cn",
			"*.icitymobile.mobi",
			"id6.me",
			"*.pingan.com.cn",
			"*.cmbchina.com",
			"*.localnetwork.uop",
			"mfs.ykimg.com*.ttf",
			"*.icbc.com.cn",
		},
		Geo_location_checker: []string{
			"http://ip-api.com/json/?lang=zh-CN",
			"https://raw.githubusercontent.com/Orz-3/Orz-3/master/QuantumultX/IP.js",
			"http://api.live.bilibili.com/ip_service/v1/ip_service/get_ip_addr?",
			"https://cdn.jsdelivr.net/gh/KOP-XIAO/QuantumultX@master/Scripts/IP_bili_cn.js",
		},
		Resource_parser_url: []string{
			"https://cdn.jsdelivr.net/gh/KOP-XIAO/QuantumultX@master/Scripts/resource-parser.js",
		},
		Excluded_routes: []string{
			"239.255.255.250/32",
			"24.105.30.129/32",
			"185.60.112.157/32",
			"185.60.112.158/32",
			"182.162.132.1/32",
		},
		Fallback_udp_policy: "direct",
		Profile_img_url:     "https://raw.githubusercontent.com/KOP-XIAO/QuantumultX/master/img/dragonball/1.PNG",
	}
	return gerl
}
