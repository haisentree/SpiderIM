package MsgGateway

import (
	DBMysql "SpiderIM/pkg/db/mysql"
	DBModel "SpiderIM/pkg/db/mysql/model"
	pbMsgGateway "SpiderIM/pkg/proto/msg_gateway"
	"log"

	"google.golang.org/grpc"

	"github.com/go-playground/validator/v10"
)

var (
	MsgGatewaySrvClient pbMsgGateway.MsgGatewayClient
	MysqlDB             DBMysql.MysqlDB
	Validate            *validator.Validate
)

func init() {
	SrvClient_Init()
	DBMysql_Init()
	Validate_Init()
}

func SrvClient_Init() {
	conn, err := grpc.Dial("0.0.0.0:8849", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	msgGatewaySrvClient := pbMsgGateway.NewMsgGatewayClient(conn)
	MsgGatewaySrvClient = msgGatewaySrvClient
}

func DBMysql_Init() {
	MysqlDB.InitMysqlDB()
	MysqlDB.DB.AutoMigrate(&DBModel.Client{})
}

func Validate_Init() {
	v := validator.New()
	Validate = v
}
