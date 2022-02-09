package serverremote

import "fmt"

type serverremote struct {
	Url            string
	Tag            string
	UpdateInterval string
	OptParser      bool
	Enable         bool
	ImgUrl         string
}

func ServerRemoteInit() *serverremote {
	svrte := &serverremote{
		Url:            "http://yourselfbooklink.com",
		Tag:            "MergeClash",
		UpdateInterval: "172800",
		OptParser:      true,
		Enable:         true,
	}
	return svrte
}

func ServerRemote() string {
	svrte := ServerRemoteInit()
	return fmt.Sprintf("%s, tag=%s, update-interval=%s, opt-parser=%v, enabled=%v", svrte.Url, svrte.Tag, svrte.UpdateInterval, svrte.OptParser, svrte.Enable)
}
