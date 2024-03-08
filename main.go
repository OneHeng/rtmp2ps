package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"rtmp2ps/config"
	"rtmp2ps/forward"
	log "rtmp2ps/logger"
	"rtmp2ps/rtmp"
)

func main() {
	log.LogInit("debug")
	
	configFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
		return
	}
	
	defer configFile.Close()
	
	buf, _ := ioutil.ReadAll(configFile)
	
	var conf config.Config
	err = json.Unmarshal(buf, &conf)
	if err != nil {
		log.Error("json Unmarshal failed, err:", err)
		return 
	}
	
	if conf.RtmpSvrConfig.ListenAddr != "" {
		svr := rtmp.NewRtmpSvr(conf.RtmpSvrConfig.ListenAddr)
		go func() {
			err = svr.Start()
			if err != nil {
				log.Error("rtmp svr start failed, err:", err)
			}
		}()
	}
	
	if conf.MediaSvrConfig.ForwardAddr != "" {
		conn := forward.NewConn(conf.MediaSvrConfig)
		err = conn.StartForward()
		if err != nil {
			log.Error("forward failed, err:", err)
		}
	}
	
}