package baseAPIClient

import (
	pbBaseAPIClient "SpiderIM/pkg/proto/base_api/client"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

// 该模块功能:
// 1.创建websocket client用户
// 2.client 发送本地每个client seq给ws服务器，然后对比数据库，将最新的client seq发送给用户。（全部发送？对比后部分发送？）
//
//	client收到消息后，发送ws消息，ws服务器解析后，读取mongodb的数据，发送给client。

// 使用密钥来创建
// 这个接口只在创建user时候调用与client进行绑定
type CreateClientReq struct {
	SecretKey  string `json:"secret_key" binding:"required,max=32"`
	ClientType uint64 `json:"client_type" binding:"required"`
}

type CreateClientToMessageReq struct {
	ClientID uint64 `json:"client_id" binding:"required"`
	RecvID   uint64 `json:"recv_id" binding:"required"`
}

func CreateClient(c *gin.Context) {
	params := CreateClientReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(400, gin.H{"errCode": 400, "errMsg": "json data error"})
		return
	}

	req := &pbBaseAPIClient.CreateClientReq{
		SecretKey:  params.SecretKey,
		ClientType: uint32(params.ClientType),
	}

	resp, err := BaseAPIClientSrvClient.CreateClient(context.Background(), req)
	if err != nil {
		log.Println("rpc error!")
		c.JSON(200, gin.H{"errCode": 500, "errMsg": "rpc error"})
		return
	}
	// User存储clientID、clientUUID
	c.JSON(200, gin.H{"errCode": 200, "errMsg": "success", "clientID": resp.ClientID, "clientUUID": resp.ClientUUID})
}

func CreateClientToMessage(c *gin.Context) {
	params := CreateClientToMessageReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(400, gin.H{"errCode": 400, "errMsg": "json data error"})
		return
	}

	req := &pbBaseAPIClient.CreateClientToMessageReq{
		ClientID: params.ClientID,
		RecvID:   params.RecvID,
	}

	resp, err := BaseAPIClientSrvClient.CreateClientToMessage(context.Background(), req)
	if err != nil {
		log.Println("rpc error!")
		c.JSON(200, gin.H{"errCode": 500, "errMsg": "rpc error"})
		return
	}
	c.JSON(200, gin.H{"errCode": 200, "errMsg": "success", "client_to_message": resp.ClientToMsgID})
}

func CreateCollectToMessage(c *gin.Context) {

	req := &pbBaseAPIClient.CreateCollectToMessageReq{Create: true}

	resp, err := BaseAPIClientSrvClient.CreateCollectToMessage(context.Background(), req)
	if err != nil {
		log.Println("rpc error!")
		c.JSON(200, gin.H{"errCode": 500, "errMsg": "rpc error"})
		return
	}
	c.JSON(200, gin.H{"errCode": 200, "errMsg": "success", "collect_to_message": resp.CollectToMsgID})
}
