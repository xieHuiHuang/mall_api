# 电商系统-微服务 mall_api

##### 一、开发环境安装：
1.1 项目开发工具及使用技术栈
+ sdk:go version go1.18.4 windows/amd64
+ 开发工具：goland2021.3.3
+ 技术栈 go、gin、gorm、grpc、consul、nacos、docker
+ 数据库：mysql8、redis5.0

1.2 代码git clone
+ git clone仓库代码：git clone git@github.com:xiehuihuang/mall_api.git
+ go mod tidy

1.3 docker安装mysql
+ 下载镜像：docker pull mysql:8.0.22
+ 查看镜像：docker images
+ 通过镜像启动：
> ```shell
  > docker run -p 3306:3306 --name mysql -v $PWD/data/developer/mysql/conf:/etc/mysql/conf.d -v  $PWD/data/developer/mysql/logs:/logs -v  $PWD/data/developer/mysql/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:8.0.22
  > ```
+ 重启容器：docker restart mysql
+ 查看镜像：docker ps -a
+ 停止正在运行镜像：docker stop mysql
+ 删除镜像：docker rm mysql

1.4 docker安装redis
+ 下载镜像：docker pull redis:latest
+ 查看镜像：docker images
+ 通过镜像启动：
> ```shell
  > docker run -p 6379:6379 --name redis -d redis:latest --requirepass "123456"
  > ```
+ 重启容器：docker restart redis
+ 查看镜像：docker ps -a

1.5 docker安装consul服务
+ 拉取镜像：docker pull consul:latest
+ 安装
> ```shell
  > docker run -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600/udp consul consul agent -dev -client=0.0.0.0
  > docker container update --restart=always 容器id 
  > ```
+ 浏览器访问：http://127.0.0.1:8500

1.6 docker安装nacos配置中心服务
+ 下载安装
> ```shell
  > docker run --name nacos-standalone -e MODE=standalone -e JVM_XMS=512m -e JVM_XMX=512m -e JVM_XMN=256m -p 8848:8848 -d nacos/nacos-server:latest
  > ```
+ 访问： http://127.0.0.1:8848/nacos
+ 登录：账号密码都是nacos

##### 二、用户微服务（user_web）：
1 用户服务
+ 用户注册接口
+ 用户登录接口
+ 获取图片验证码
+ 阿里云短信验证码
+ 用户列表查看
+ 检查密码
+ 登录权限校验
  + jwt
+ 通过api获取nacos的配置和nacos的配置更新
+ gin集成nacos
+ gin从consul中同步服务进行负载均衡

2 go grpc生成go文件
+ protoc -I . user.proto --go_out=plugins=grpc:.

3 nacos配置服务中心: user_web.json 配置信息如下
```json
{
  "name": "user_web",
  "host": "127.0.0.1",
  "port": 8021,
  "tags": ["user_web", "web"],
  "user_srv":{
    "name": "user_srv"
  },
  "jwt": {
    "key": "IOohnrP5!%MAZYYogKgKOP&iUA9^Xm%h"
  },
  "sms":{
    "key": "**************",
    "secrect": "*****************",
    "expire": 300
  },
  "redis": {
    "host": "127.0.0.1",
    "port": 6378,
    "password": "123456",
    "db1": 0,
    "db2": 14,
    "expire": 60
  },
  "consul": {
    "host": "127.0.0.1",
    "port": 8500
  }
}
```

##### 三、商品微服务（goods_web）：
1 商品服务
+ 商品模块（列表分页接口、新建商品、获取商品详情接口、商品删除/更新接口）
+ 商品分类接口
+ 商品轮播图接口
+ 品牌列表接口
+ 品牌分类接口
+ 用户权限验证
    + jwt
+ 通过api获取nacos的配置和nacos的配置更新
+ gin集成nacos
+ gin从consul中同步服务进行负载均衡

2 go grpc生成go文件
+ protoc -I . goods.proto --go_out=plugins=grpc:.

3 nacos配置服务中心: goods_srv.json 配置信息如下
```json
{
  "name": "goods_web",
  "host": "127.0.0.1",
  "port": 8022,
  "tags": ["goods_web", "web"],
  "goods_srv":{
    "name": "goods_srv"
  },
  "jwt": {
    "key": "IOohnrP5!%MAZYYogKgKOP&iUA9^Xm%h"
  },
  "consul": {
    "host": "127.0.0.1",
    "port": 8500
  }
}
```
#### 四 运行：
+ go  run main.go




