package parser

import (
	"fmt"
	"strings"
	"yuanmultx/dns"
	"yuanmultx/general"
	"yuanmultx/policy"
)

func getNewLine() {
	fmt.Println()
}
func parserSlice(value []string) strings.Builder {
	size := len(value)
	var content strings.Builder
	content.WriteString(value[0])
	for i := 1; i < size; i++ {
		content.WriteString(",")
		content.WriteString(value[i])
	}
	return content
}
func GeneralParser() {
	gerl := general.GeneralInit()
	fmt.Printf("[general]\n")
	fmt.Printf("server_check_url=")
	content := parserSlice(gerl.Server_check_url)
	fmt.Println(content.String())
	fmt.Printf("network_check_url=")
	content = parserSlice(gerl.Network_check_url)
	fmt.Println(content.String())
	fmt.Printf("dns_exclusion_list=")
	content = parserSlice(gerl.Dns_exclusion_list)
	fmt.Println(content.String())
	fmt.Printf("geo_location_checker=")
	content = parserSlice(gerl.Geo_location_checker)
	fmt.Println(content.String())
	fmt.Printf("resource_parser_url=")
	content = parserSlice(gerl.Resource_parser_url)
	fmt.Println(content.String())
	fmt.Printf("excluded_routes=")
	content = parserSlice(gerl.Excluded_routes)
	fmt.Println(content.String())
	fmt.Printf("fallback_udp_policy=")
	fmt.Println(gerl.Fallback_udp_policy)
	fmt.Printf("profile_img_url=")
	fmt.Println(gerl.Profile_img_url)
	getNewLine()
}
func DnsParser() {
	ds := dns.DnsInit()
	fmt.Printf("[dns]\n")
	fmt.Println(ds.Tag)
	fmt.Printf("address=%v\n", ds.Address)
	for _, svr := range ds.Server {
		fmt.Printf("server=%v\n", svr)
	}
	getNewLine()
}
func PolicyParser() {
	fmt.Printf("[policy]\n")
	ply := policy.PolicyInit()
	for _, grp := range ply.Group {
		fmt.Printf("static=%v,", grp.Name)
		for _, opt := range grp.Option {
			fmt.Printf("%v,", opt)
		}
		if len(grp.ServerTagRegex) != 0 {
			fmt.Printf("server-tag-regex=%v,", grp.ServerTagRegex)
		}
		fmt.Printf("img-url=%v\n", grp.ImgUrl)

	}
	for _, rul := range ply.Rule {
		fmt.Printf("url-latency-benchmark=%v,server-tag-regex=%v,check-interval=%v,tolerance=%v,img-url=%v\n", rul.UrlLatencyBenchmark, rul.ServerTagRegex, rul.CheckInterval, rul.Tolerance, rul.ImgUrl)
	}
	getNewLine()
}
