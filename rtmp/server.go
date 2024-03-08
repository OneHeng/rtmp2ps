package rtmp

import (
	log "rtmp2ps/logger"
	vdkrtmp "github.com/deepch/vdk/format/rtmp"
)
type RtmpSvr struct {
	listenAddr string
	svr *vdkrtmp.Server
}

func NewRtmpSvr(listenAddr string) *RtmpSvr{
	return &RtmpSvr{
		listenAddr: listenAddr,
		svr: &vdkrtmp.Server{
			Addr: listenAddr,
		},
	}
}

func (s *RtmpSvr) Start() error {
	log.Info("start rtmp svr, listenaddr:", s.listenAddr)
	
	s.svr.ListenAndServe()
	return nil
}