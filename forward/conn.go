package forward

import "rtmp2ps/config"

type Conn struct {
	conf config.MediaSvrConfig
}

func NewConn(conf config.MediaSvrConfig) *Conn {
	return &Conn{}
}

func (c *Conn) StartForward() error {
	for {
		
	}
	
	return nil
}