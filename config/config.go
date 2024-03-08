package config

type Config struct {
	RtmpSvrConfig RtmpSvrConfig `json:"rtmpsvr_config"`
	MediaSvrConfig MediaSvrConfig `json:"mediasvr_config"`
}

// rtmp服务配置
type RtmpSvrConfig struct {
	ListenAddr string `json:"listen_addr"`
}

// 转推媒体服务配置
type MediaSvrConfig struct {
	ForwardAddr string `json:"forward_addr"` // 转推媒体服务器地址
	Network string `json:"network"` // tcp/udp
}