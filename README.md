# SpiderIM

# 一、功能演示

> 之前录制了个视频专门演示功能的，搞不见了，现在重新搞一下，演示一下几个发送的消息。

**使用的技术**

- mysql，mongodb
- redis
- rabbitmq
- grpc
- websokcet

## 模块说明

### base_api

- `/create_client`：base_api这个模块是用来编写client相关信息，连接Websocket server的是client ID，一个client发送信息给另一个clinet，通过ClinetID标识。（这个接口规范在业务创建一个user的时候，创建一个client和消息系统交互）

- `/create_client_to_message`：一个client首次给另一个client发送消息，在数据库创建一个会话，会创建如下数据。（规范发生在业务添加好友成功的时候，调用该接口）

client29要给client30，会创建两条数据，也就是各自的消息队列是独立的，min_seq和max_seq是同步的，这样编写同步消息逻辑更加简单。一条消息被存储了两份，也不会造成太大的压力。

（qq中自己这里删除一条消息，对方那里没有删除，这就是两个人的消息是独立存储的）

![image-20250102164047116](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102164047116.png)

这里直接来看一message，可以看到数据存储了两条，拉去消息直接搜索client_to_message_id即可。

![image-20250102170304526](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102170304526.png)

- `/create_collect_to_message`：这是用来创建群聊id的接口。（这个在业务层创建群聊的时候创建的）

用户向群组发送消息，collectID作为接受方。因此，一条群消息只存储了一次，拉去群消息时候根据collectID就可以了。

![image-20250102165449840](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102165449840.png)

看一下存储的群消息，搜索collect_to_message_id即可，会根据seq_id来排序消息。

![image-20250102170507368](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102170507368.png)

```go
func main() {
	router := gin.New()

	client := router.Group("/client")
	{
		client.POST("/create_client", baseAPIClient.CreateClient)
		client.POST("/create_client_to_message", baseAPIClient.CreateClientToMessage)
		client.POST("/create_collect_to_message", baseAPIClient.CreateCollectToMessage)
	}
	router.Run(":8080")
}

```

### msg_gateway

该模块会启动一个Websocket  Server用来接受客户端的连接。（这里有个安全隐患，没有对client进行校验。如果攻击者知道client_id和消息格式,就可以连接随机连接WSS进行发送消息）

:joy:我设置了uuid，首次连接需要提供uuid，如果client_uuid也被别人知道了就不安全了。

```go
type WServer struct {
	wsAddr         string
	wsMaxConnNum   int
	wsUpGrader     *websocket.Upgrader
	wsClientToConn map[uint64]map[uint8]*WSClient
}

type WSClient struct {
	*websocket.Conn
	platformID uint8
	clientID   uint64
	clinetType uint8
}
```

说明一下字段：

**客户端类型：**(有误)

- 1：内部客户端。也就是msg_relay模块
- 2：普通用户客户端。

**平台类型：**

- 0：内部客户端。msg_relay模块
- 1：Android
- 2：IOS
- 3：PC
- 4：Web

**消息类型：**

- 1：普通消息。一对一
- 2：群发消息。
- 3：转发消息。message上传的消息，不经过后端处理，直接转发。（msg_relay模块）

---



这个模块在client连接server时候会校验平台id、client_type是否合法，client_id、client_uuid是否存在，然后就是读取消息了，这里解释一下消息规范。

**解析消息**`ws.msgParse(conn, message)`

```go
func (ws *WServer) msgParse(conn *WSClient, binaryMsg []byte) {
	m := pkgMessage.CommonMsg{}
	json.Unmarshal(binaryMsg, &m)

	if err := Validate.Struct(m); err != nil {
		log.Println("validate error:", err)
		return
	}
	switch m.MessageType {
	case pkgMessage.Single_Common_Message_Request:
		log.Println("msg prase single common message")
		ws.parseSingleCommMsg(conn, &m)
	case pkgMessage.Single_Relay_Message_Request:
		log.Println("serverMsg")
		ws.parseSingleRelayMsg(conn, &m)
	case pkgMessage.Group_Common_Message_Request:
		log.Println("group msg")
		ws.parseGroupCommMsg(conn, &m)
	case pkgMessage.Group_List_Message_Request:
		ws.parseGroupListMsg(conn, &m)
	case pkgMessage.Control_Pull_Client_Message:
		ws.parsePullClientMsg(conn, &m)
	case pkgMessage.Control_Pull_Collect_Message:
		ws.parsePullCollectMsg(conn, &m)
	case pkgMessage.Control_Get_Client_Max_Seq:
		ws.parseGetClientMaxSeq(conn, &m)
	case pkgMessage.Control_Get_Collect_Max_Seq:
		ws.parseGetCollectMaxSeq(conn, &m)
	case pkgMessage.Control_Get_Client_Status:
		ws.parseGetClientStatus(conn, &m)
	default:
		log.Println("clientType error")
	}
}
```

消息类型做了备注。说明一下特殊的**群发消息**。

一个用户向一个群组发送消息，在业务层面，会提取发送放sendID，也就是clientID，接收方是一个数组，业务层面查取群组里有那些clientID，存储到接收方数组中，这就是Group_List_Message_Request消息，这条消息会被存储到collect_to_message表中。这样就存储了一条信息，也不用把接受方数组存储到数据库。

msg_relay模块获取到了Group_List_Message_Request消息，会将该消息分割转换成Group_Common_Message_Request消息，然后发送给对应的WSS，所以说消息3只需要转发不需要存储。

```go
// 消息类型
const (
	Single_Common_Message_Request = 1 // client对client发送的单条消息
	Single_Relay_Message_Request  = 2 // msg_relay从MQ中读取，转发给client
	Group_Common_Message_Request  = 3 // 该条消息用于存储，不需要转发
	Group_List_Message_Request    = 4 // 该条消息包含recv数组，需要进行转发消息,自带seq
	Control_Pull_Client_Message   = 5 // 首次登录，刷新时候，同步与数据库中的消息
	Control_Pull_Collect_Message  = 6
	Control_Get_Client_Max_Seq    = 7 // 获取seq，与本都seq对比
	Control_Get_Collect_Max_Seq   = 8
	Control_Get_Client_Status     = 9 // 获取client的在线状态
)
```

### msg_relay

该模块从MQ中读取消息，进行存储到数据库和转发消息。

为什么需要该模块？一是功能解耦。二是WSS连接的客户端有限，后期会有多个WSS，如果client_1连接这WSS_1，需要连接这WSS_2上的client_2发送消息，不能直接发送，那就需要msg_relay进行转发，msg_relay连接这WSS_1和WSS_2。

（该功能还没开发，relay如何知道client_2连接在那个WSS上，当然是通过查询redis）

### rpc模块

这个模块不多说了，当时为了使用grpc体现技术来做的这模块，其实没必要。后面开发也会先将该模块移除，等达到那个业务级别的再开发，初期开发rpc只会拖慢开发进度和增加复杂度。

## 功能演示

### 注册

**1.项目启动**

首先要确保mysql、mongodb、rabbitmq、redis服务正常。我服务全部安装在linux虚拟机中，项目也是在虚拟中启动。

总共分为五个要启动，先启动rpc模块，在启动base_api和msg_gateway，最后启动msg_relay

没有写配置文件，端口号在代码里面固定了，按照步骤创建数据启动应该就没啥问题。

![image-20250102171049585](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102171049585.png)

![image-20250102171110847](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102171110847.png)

![image-20250102171553260](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102171553260.png)

![image-20250102171624156](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102171624156.png)

![image-20250102171650885](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102171650885.png)

**2.注册两个普通client**

![image-20250102185909381](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102185909381.png)

```go
{
    "secret_key":"dsfrserererst",
    "client_type":1
}
```

这个secret_key字段目前在代码里是固定的

```json
{
    "clientID": 43,
    "clientUUID": "e1f555af-46ee-4acc-97c1-7ff97a87bae3",
    "errCode": 200,
    "errMsg": "success"
}

{
    "clientID": 44,
    "clientUUID": "b48bc214-2498-48a3-9d0c-2d52e5dbf5d1",
    "errCode": 200,
    "errMsg": "success"
}
```

**3.注册一个msg_relay客户端**

这个消息中的client_type是无效的，不会根据该字段的值创建对应客户端类型。当初的设计逻辑是要根据secret_key，因为普通的secret_key是不能创建内部客户端。

所以这里需要先创建一个普通client，然后再数据库把类型值改成0，并且把对应的数据填写再msg_relay中的代码中。

```go
wsConn, _, err := websocket.DefaultDialer.Dial("ws://192.168.45.128:8848/ws?clientID=37&clientUUID=fbeaedca-ca5b-4cf9-a53d-e5f1e5b59b82&platformID=0", nil)
```

![image-20250102191041920](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102191041920.png)

**4.创建一个client_to_message**

这个步骤在忘记做了，导致我在演示两个client发送单独消息的时候，直接错误崩溃了。

![image-20250102195041358](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102195041358.png)

```json
{
	"client_id": 43,
	"recv_id": 44
}
```

见鬼，还要下面再交换一下，不做下面这个步骤，还会出现上面错误，因为数据库中值存储一条对应的消息。

![image-20250102195530688](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102195530688.png)

```
{
	"client_id": 44,
	"recv_id": 43
}
```

![image-20250102195753375](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102195753375.png)

---

嗯，还真的见鬼

![image-20250102195956034](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102195956034.png)

好吧，我把这五个模块重启了一下，能够正常发送消息了。

这个需要交换一下创建两个消息体，放了一年半，我自己开发的自己都搞忘记了。

**4.创建一个collect**

这里不需要携带数据，后期开发需要要携带一个密钥的

![image-20250102191835549](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102191835549.png)

```json
{
    "collect_to_message": 6,
    "errCode": 200,
    "errMsg": "success"
}
```

---

### 消息

1. client_43给client_44发送消息hello1

首先要再把client的注册信息填入到params中才能连接上

![image-20250102192810799](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102192810799.png)

关于params中包含考虑send_id和planform_id，这是首次连接存储在WSS上的客户端信息，但是消息中也包含了send_id和planform_id。理论上应该是从客户端信息中提取这两个字段，但我没看代码，不确定之前写的。

演示中，消息格式错误，导致msg_gateway直接崩溃停止，后面要加上对应错误的处理逻辑。

两个客户端连接WSS后，发送下面消息。（这个WS通讯，Message没有限定json格式嵌套形式，这里也是踩坑试出来下面格式的）

```json
{
    "message_type": 1,
    "send_id":43,
    "platform_id": 2,
    "data":"{\"recv_id\":44,\"content\":\"msg 2\"}"
}
```

再redis中可以看到存储的状态信息是正确的

![image-20250102200709426](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102200709426.png)

可以看到发送的消息成功

![image-20250102200758984](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102200758984.png)

![image-20250102200829319](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102200829319.png)

看一下终端的数据流向

![image-20250102200953728](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102200953728.png)

![image-20250102201002108](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102201002108.png)

数据库存储

![image-20250102201104340](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102201104340.png)

后面的消息就不演示一下功能，存储和数据流向就不解释了

---

**2.群发消息**

- 使用client_32向client_43,client_44发送消息

![image-20250102201433844](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102201433844.png)

- 成功接受

![image-20250102201540572](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102201540572.png)

![image-20250102201558662](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102201558662.png)

### 控制消息

1. **获取client状态**（未完成）

![image-20250102201948954](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102201948954.png)

2. 获取最大序列（没有开发wss返回，终端返回）

（这里的没有对发送信息进行身份验证，比如B不是A的好友，但是可以通过获取B的状态，这个和业务相关，来确定是否要限制）

![image-20250102202149964](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102202149964.png)

![image-20250102202310989](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102202310989.png)

3. 拉去客户端消息（后面演示使用的是之前的client）

![image-20250102202649977](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102202649977.png)

![image-20250102202701472](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102202701472.png)

4. 获取群消息最大序列

![image-20250102202824762](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102202824762.png)

5. 拉去群消息（拉去消息的seq_id是起始信息id，有客户端本地提供，直接拉取到最大，后期开发可能需要分组拉取）

![image-20250102203115816](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102203115816.png)

![image-20250102202941448](https://raw.githubusercontent.com/haisentree/imageBed/main/image2024/image-20250102202941448.png)

# 二、小结

> 代码实现就不去说明了，比较麻烦，这个是项目初期，功能简单代码也容易看懂。这个项目是一年半前开发的，当时开发完自己就去备考了。之前是个私有仓库，因为里面没有设置的生产和开发配置文件，现在开源总结是为了找工作中演示一下。
>
> 后面如果时间就去规范开发一下这个项目。

**功能保证**

**1.消息同步**

用户之前的消息使用**读扩散**，同一个消息存储两份。群消息使用**写扩散**，消息存储一份。

**2.消息可靠性保证**

可靠性就是向TCP那样发送确认的形式。层数越多越麻烦，要根据业务等级来确认消息是否要保证可靠。

client->msg_gateway->消息rpc->mq->msg_relay

这个确认机制，可以实现客户端消息已读的功能。

mq的确认机制实现过，看的相关数据java接口比较完善和资料多，golang的我找了有一些时间才找到。

**3.应对海量用户**

用户量起来了，首先要有多个WSS，client连接的时候通过WSS连接中心，来分配连接那个WSS。这种情况下就没必要再使用grpc了，去掉这一层。

WSS和msg_relay都是可以水平扩展到，web接口的压力很小，应为一般只会再业务层面创建用户、添加好友，创建群组的时候调用。那么性能瓶颈就很可能出现在MQ这里了，因为这里只有一个mq，就不用考虑信息同步问题了，wss的消息在mq中汇总。如果要使用在不同机器上使用mq完成功能，需要深入学习一下。

**4.消息有序性**

使用seq，好像不是很难。

**5.数据存储**

目前的逻辑好像是，msg_relay从mq读取消息，完成存储到mongodb，mysql中，在发送，发送到msg_gateway收到其确认消息，则对mq中的消息消费完毕。

太复杂了，之前还考虑先存储再转发会给消息带来延迟。这里要对比出mongodb，mysql的优势，如果mongodb因为是非结构化，存储和读取速度比mysql快，那就先只存储到mongodb中，比如聊天信息云端保存15天，再mongodb中存储15天，用一个定时任务来将超过15天的数据存储到mysql中归档。

这样做后期的数据统计，只从mysql中统计，最新15天无法统计，如果要统计最新15天的还要从mongodb中读取。而且，后期开发坑你就不需要统计分析功能，非要统计mongodb中就可以统计，这15天之后的数据可以不要。预计后期开发客户端数据是端到端加密的，所以说更加不需要存储再mysql中。

现在的消息主要粗出再msyql中，mongodb用得不熟。

---

**后期工作**

1. 完善安全和权限机制
2. 规范消息格式，完善消息功能

```json
{
    "status":200,
    "message":"success",
    "data":{
        
    }
}
```

3. 实现wss水平扩展
4. 删除rpc模块
5. 确认存储流程
6. 规范配置文件和日志
