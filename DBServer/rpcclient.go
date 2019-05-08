package main

import (
	"time"

	log "github.com/cihub/seelog"
)

type RpcClient struct {
	Peerid string   // the peerid of connect service
	Hbchan chan int //heart beat channel
	Status int      // the status of peer
}

const HEART_BEAT_TIME_OUT = 30 // 30 second

type RpcClientMgr struct {
	clis map[string]*RpcClient
}

var RpcCliMgr *RpcClientMgr

func init() {
	RpcCliMgr = &RpcClientMgr{
		clis: make(map[string]*RpcClient),
	}
}

func (this *RpcClientMgr) NewClient(id string) (*RpcClient, error) {
	newcli := &RpcClient{
		Peerid: id,
		Hbchan: make(chan int),
	}

	if _, ok := this.clis[id]; !ok {
		this.clis[id] = newcli
	} else {
		log.Errorf("client peerid[%v] already exsit .. exchange?", id)
		this.clis[id].Close()
		delete(this.clis, id)
		this.clis[id] = newcli
	}

	return newcli, nil
}

func (this *RpcClientMgr) GetClient(id string) *RpcClient {
	if cli, ok := this.clis[id]; ok {
		return cli
	}
	return nil
}

func (this *RpcClientMgr) RomoveClient(id string) {
	if cli, ok := this.clis[id]; ok {
		cli.Close()
		delete(this.clis, id)
	}
}

func (this *RpcClient) Close() {
	close(this.Hbchan)
}

func (this *RpcClient) HandleHeartBeat() {
	ticker := time.NewTimer(HEART_BEAT_TIME_OUT * time.Second)
	defer ticker.Stop()
LBL_handleHeartBeat_loop:
	for {
		select {
		case _, ok := <-this.Hbchan:
			log.Debugf("recv heartbeat from client peer[%v]", this.Peerid)
			if !ok {
				log.Infof("peer[%v] heart beat check stopped forcibly", this.Peerid)
				break LBL_handleHeartBeat_loop
			}
			ticker.Reset(HEART_BEAT_TIME_OUT * time.Second)
		case <-ticker.C:
			log.Debugf("haven't received peer[%v] heartbeat for 30 seconds, so break", this.Peerid)
			break LBL_handleHeartBeat_loop
		}
	}
}
