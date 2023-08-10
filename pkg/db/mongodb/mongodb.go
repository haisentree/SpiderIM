package DBmongo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
}

func (m *MongoDB) Init_mongodb() {
	// 设置客户端连接配置
	username := "root"
	password := "qwer.1234@TYUI"
	clientOptions := options.Client().ApplyURI("mongodb://192.168.45.128:27017")
	a := options.Credential{
		Username: username,
		Password: password,
	}
	clientOptions.SetAuth(a)

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	m.Client = client
}

// 增加一条client消息
func (m *MongoDB) Insert_client_msg() {
	c := m.Client.Database("test").Collection("client")
	d := bson.M{"client": 28}
	if _, err := c.InsertOne(context.Background(), d); err != nil {
		log.Println("inser client msg error")
	}
}

// 更新client数据
func (m *MongoDB) Update_client_msg() {
	c := m.Client.Database("test").Collection("client")
	filter := bson.M{"client": 28}
	// update := bson.M{"$set": bson.M{"client": 23}}
	// update := bson.M{"$push": bson.M{"min_seq": 1}}
	// update := bson.M{"$push": bson.M{"msg":bson.M{"recv_id":23,"min_seq":1,"max_seq":32}}}
	update := bson.M{"$push":bson.M{"data":bson.M{"msg3":bson.A{bson.M{"seq":4,"content":"hello44"}}}}}

	

	if _, err := c.UpdateOne(context.Background(), filter, update); err != nil {
		log.Println("inser client msg error")
	}
}

func (m *MongoDB) Find_client_msg() {
	c := m.Client.Database("test").Collection("client")
	filter := bson.M{"client": 22}
	var result bson.M
	if err := c.FindOne(context.Background(), filter).Decode(&result); err != nil {
		log.Println("find client msg error")
	}
	log.Println(result)

}

// 根据seq查询client消息

// 增加一条group消息

// 根据seq查询group消息
