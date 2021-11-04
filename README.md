# Sphinx Service

### 交易服务

* [x] 编写proto(grpc)
* [x] 编写model(ent)
* [x] 更新项目框架到latest
* [ ] 接入Apollo并获取mysql数据库
* [ ] ent的非drop自动迁移
* [ ] 测试：make编译，确认proto正常使用
* [ ] 测试数据：填充CoinInfo
* [ ] 测试：实现CoinInfo接口并跑通，确认http请求给到grpc
* [ ] 功能：签名和代理服务接入点
* [ ] 阻塞：签名和代理服务功能实现
* [ ] 功能：实现传递接口，创建账户、余额、交易查询、状态查询
* [ ] 功能：实现转账接口，交易状态机
* [ ] 测试：节点/网络异常状态下，交易状态正常，具备恢复能力
* [ ] 确认：交易服务与签名/代理的连接安全性
* [ ] 测试：整体测试

### 钱包插件

* [x] 编写proto(grpc)
* [x] 编写model(skipped)
* [ ] 功能：余额查询、交易查询
* [ ] 功能：广播交易
* [ ] 功能：获取预签名信息
* [ ] 测试：单元测试
* [ ] 测试：连接上交易服务，并在其上测试

### 签名服务

* [x] 编写proto(grpc)
* [x] 编写model(gorm/skipped)
* [ ] model转为oss存储
* [ ] 框架：基本代码结构
* [ ] 功能：创建账户
* [ ] 功能：进行签名
* [ ] 测试：连接上交易服务，并在其上测试
* [ ] 后续完善：脚本签名鉴权

### 框架命令

* make init ```初始化仓库，创建go.mod```
* make verify ```验证开发环境与构建环境，检查code conduct```
* make verify-build ```编译目标```
* make test ```单元测试```
* make generate-docker-images ```生成docker镜像```
* make sphinx-service ```单独编译服务```
* make sphinx-service-image ```单独生成服务镜像```
* make deploy-to-k8s-cluster ```部署到k8s集群```

###  Models

* 交易服务(service)部分采用统一框架ent
* 签名服务(sign/keystore)部分暂用云端存储COS
  * 已用gorm编写结构，可参照type进行序列化后存储
* 插件服务(plugin/agent)无需存储数据，使用的基本都是公链功能

### 描述

- 运行在k8s
- 对外提供交易业务服务
- 依赖外部签名服务
- 框架版本11.02

### 最佳实践

* 每个服务只提供单一可执行文件，有利于docker镜像打包与k8s部署管理
* 每个服务提供http调试接口，通过curl获取调试信息
* 集群内服务间direct call调用通过服务发现获取目标地址进行调用
* 集群内服务间event call调用通过rabbitmq解耦
