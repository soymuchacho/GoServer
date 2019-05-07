package rpcapi

import (
	"GoServer/DBServer/rpcclient"
	pb "GoServer/Share/Proto/dbrpc"
	"context"

	log "github.com/cihub/seelog"
)

type DBRpcServer struct{}

func (s *DBRpcServer) TestDBServer(ctx context.Context, in *pb.MsgTestDBRequest) (*pb.MsgTestDBResponse, error) {
	return &pb.MsgTestDBResponse{Reply: "test ok" + in.Test}, nil
}

func (s *DBRpcServer) Register(ctx context.Context, in *pb.MsgRegisterReq) (*pb.MsgRegisterAck, error) {
	var ret int32 = 0
	cli, err := rpcclient.RpcCliMgr.NewClient(in.Peerid)
	if err != nil {
		ret = 1
	}
	go cli.HandleHeartBeat()
	return &pb.MsgRegisterAck{Result: ret}, err
}

func (s *DBRpcServer) HeartBeat(instream pb.DB_HeartBeatServer) error {
	var peerid string
	for {
		if note, err := instream.Recv(); err == nil {
			peerid = note.Peerid
			log.Debugf("recv heartbeat from peer[%v] ", note.Peerid)
			cli := rpcclient.RpcCliMgr.GetClient(note.Peerid)
			if cli == nil {
				log.Errorf("client[%v] is not exsit!!", note.Peerid)
				break
			} else {
				cli.Hbchan <- 1
			}
		} else {
			rpcclient.RpcCliMgr.RomoveClient(peerid)
			log.Debug("break ", err)
			break
		}
	}
	return nil
}
