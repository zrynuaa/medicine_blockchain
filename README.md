# medicine_blockchain
## 		-- 基于CPABE06的Fabric医药分开联盟链

## 角色：服务节点、药店、医院

* 医院发布处方经过ABE加密上链（只发布化学名）；
* 药店获得能够解密的处方信息根据自己药店的药品信息，发出卖药信息加密上链；
* 服务节点根据用户身份，给出可以购买的药品信息，用户选择要购买的药品，购买信息加密上链。



## 运行

### 依赖库

Fabric的SDK（go语言）

https://github.com/thorweiyan/fabric_go_sdk



CPABE06的实现，基于RPC服务，Server部署在复旦内网服务器，运行系统需要能够连接内网

https://github.com/zrynuaa/cpabe06_client

CPABE实现参考文献[BSW06]：

*注：如果有需要获取CPABE server部分，请联系royzhang.go@gmail.com或者zry_nuaa@163.com*

## 步骤

### 1.依赖库准备

```
go get github.com/thorweiyan/fabric_go_sdk

go get github.com/zrynuaa/cpabe06_client
```

### 2.Step by Step

* 启动Fabric基础网络（需要dep ensure）

```
cd $GOPATH/src/gihub.com/thorweiyan/fabric_go_sdk

make restart
```

* 运行医药分开系统

```
cd $GOPATH/src/github.com/zrynuaa/medicine_blockchain

./start.sh

(如果出现端口被占用，先运行stop.sh)
```

* 前端界面

```
医院节点：
http://localhost:8880/html/hospital.html

药店1:
http://localhost:8881/html/store.html

药店2:
http://localhost:8882/html/store.html

药店3:
http://localhost:8883/html/store.html

服务节点:
http://localhost:8884/html/controller.html
```

### 3.Reference

[BSW06]  J. Bethencourt, A. Sahai, B. Waters. Ciphertext-Policy Attribute-Based Encryption, 2006