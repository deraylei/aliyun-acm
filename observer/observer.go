package observer

import (
	"github.com/xiaojiaoyu100/aliyun-acm/v2/config"
	"github.com/xiaojiaoyu100/aliyun-acm/v2/info"
)

type Observer struct {
	coll     map[info.Info]*config.Config
	h        Handler
	consumed bool
}

func New(ss ...Setting) (*Observer, error) {
	o := &Observer{}
	o.coll = make(map[info.Info]*config.Config)
	for _, s := range ss {
		if err := s(o); err != nil {
			return nil, err
		}
	}
	return o, nil
}

func (o *Observer) Ready() bool {
	var ret = true
	for _, c := range o.coll {
		if !c.Pulled {
			ret = false
			break
		}
	}
	return ret
}

func (o *Observer) Info() []info.Info {
	var ii []info.Info
	for i := range o.coll {
		ii = append(ii, i)
	}
	return ii
}

func (o *Observer) Handle() {
	if o.consumed {
		return
	}
	o.consumed = true
	o.h(o.coll)
}

func (o *Observer) UpdateInfo(i info.Info, conf *config.Config) {
	o.coll[i] = conf
}
