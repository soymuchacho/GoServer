package handle

import (
	"GoServer/Common/srpc"
	pb "GoServer/Share/Proto/dbrpc"
	"context"
	"time"

	log "github.com/cihub/seelog"
)

func ReConnect() {

}

const HEART_BEAT_TIME_OUT = 30

func HandleHeartBeat(peerid string, hwd *srpc.SrpcClient) error {
	c := pb.NewDBClient(hwd.Conn)

	timeout := time.NewTimer(HEART_BEAT_TIME_OUT * time.Second)

	putStream, err := c.HeartBeat(context.Background())
	if err != nil {
		return err
	}
AGENT_RPC_HEART_BEAT:
	for {
		err := putStream.Send(&pb.MsgHeartBeatReq{Peerid: peerid})
		if err != nil {
			log.Error("HeartBeat Error, Reconnect to db rpc")
			hwd.ReConnect()
			break AGENT_RPC_HEART_BEAT
		}
		timeout.Reset(HEART_BEAT_TIME_OUT * time.Second)
		select {
		case <-ticker.C:
			log.Error("HeartBeat Timeout, Reconnect to db rpc")
			hwd.ReConnect()
		}
	}
	return nil
}
