package mq

type Channel string

func (p Channel) String() string {
	return string(p)
}

func (p Channel) GoString() string {
	return p.String()
}
