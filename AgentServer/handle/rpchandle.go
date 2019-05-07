package handle

import (
	"GoServer/Common/srpc"
	pb "GoServer/Share/Proto/dbrpc"
	"context"
	"errors"
	"sync"
	"time"

	log "github.com/cihub/seelog"
)

var hwdMtx sync.Mutex
var funcHwd pb.DBClient

func SetFuncHwd(c pb.DBClient) {
	hwdMtx.Lock()
	defer hwdMtx.Unlock()

	funcHwd = c
}

func GetFuncHwd() pb.DBClient {
	hwdMtx.Lock()
	defer hwdMtx.Unlock()

	return funcHwd
}

const HEART_BEAT_TIME_OUT = 30

func ConnectRpc(name string, addr string) error {

	hwd := srpc.NewSrpcClient(name)
	err := hwd.StartRpcClient(addr)
	if err != nil {
		log.Debugf("connect rpc server error : address %v", addr)
		return errors.New("connect rpc server error")
	}

	SetFuncHwd(pb.NewDBClient(hwd.Conn))

	r, err := GetFuncHwd().TestDBServer(context.Background(), &pb.MsgTestDBRequest{Test: " Redis Test "})
	if err != nil {
		log.Debugf("Request Redis Test Error %v", err)
	} else {
		log.Debugf("Request test rpc success : %v", r)
	}

	for {

		// the function of HandleHeartBeat is block
		HandleHeartBeat(name)

		time.Sleep(time.Second)

		log.Debug("ReConnect Rpc Server : ", hwd.Address)
		err = hwd.ReConnect()
		if err != nil {
			log.Debugf("ReConnect Rpc Server[%v] Error [%v] ", hwd.Address, err)
			continue
		}
		SetFuncHwd(pb.NewDBClient(hwd.Conn))
	}
	return nil
}

func HandleHeartBeat(peerid string) error {
	// first to register
	r2, err := GetFuncHwd().Register(context.Background(), &pb.MsgRegisterReq{Peerid: peerid})
	if err != nil {
		log.Error("The agent register db Error!")
		return errors.New("register db service error!")
	} else {
		if r2.Result == 0 {
			log.Debugf("The agent register db successful!")
		} else {
			return errors.New("register db service error!")
		}
	}

	ticker := time.NewTicker(time.Second)
	timeout := time.NewTimer(HEART_BEAT_TIME_OUT * time.Second)
	defer timeout.Stop()

	putStream, err := funcHwd.HeartBeat(context.Background())
	if err != nil {
		return err
	}
AGENT_RPC_HEART_BEAT:
	for {
		select {
		case <-ticker.C:
			err := putStream.Send(&pb.MsgHeartBeatReq{Peerid: peerid})
			if err != nil {
				log.Error("HeartBeat Error, Reconnect to db rpc")
				break AGENT_RPC_HEART_BEAT
			} else {
				timeout.Reset(HEART_BEAT_TIME_OUT * time.Second)
			}
		case <-timeout.C:
			log.Error("HeartBeat Timeout, Reconnect to db rpc")
			break AGENT_RPC_HEART_BEAT
		}
	}
	return nil
}
