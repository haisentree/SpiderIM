package baseAPIClient

// 该模块功能:
// 1.创建websocket client用户
// 2.client 发送本地每个client seq给ws服务器，然后对比数据库，将最新的client seq发送给用户。（全部发送？对比后部分发送？）
//	client收到消息后，发送ws消息，ws服务器解析后，读取mongodb的数据，发送给client。
