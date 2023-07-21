package rpcMsgGateway

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	mysql "FlyIM/pkg/db/mysql"
	mysqlBaseApi "FlyIM/pkg/db/mysql/base_api"
	"FlyIM/pkg/db/redis"
	pbMsgGateway "FlyIM/pkg/proto/base_api/auth"
)


type rpcMsg struct {
	rpcPort int
}

func NewRpcAuthServer(port int) *rpcAuth {
	return &rpcAuth{
		rpcPort: 8848,
	}
}

func ReceiveMessage(context.Context, *MessageReq) (*MessageResp, error) {

}