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

func AddNicoRemote() string {
	svrte := &serverremote{
		Url:            "",
		Tag:            "",
		UpdateInterval: "172800",
		OptParser:      true,
		Enable:         true,
	}
	return fmt.Sprintf("%s, tag=%s, update-interval=%s, opt-parser=%v, enabled=%v", svrte.Url, svrte.Tag, svrte.UpdateInterval, svrte.OptParser, svrte.Enable)
}

func AddFrogRemote() string {
	svrte := &serverremote{
		Url:            "",
		Tag:            "",
		UpdateInterval: "172800",
		OptParser:      true,
		Enable:         true,
	}
	return fmt.Sprintf("%s, tag=%s, update-interval=%s, opt-parser=%v, enabled=%v", svrte.Url, svrte.Tag, svrte.UpdateInterval, svrte.OptParser, svrte.Enable)
}

func ServerRemote() string {
	return AddNicoRemote() + "\n" + AddFrogRemote()
}
