# online prototype
前天和[baifengbai](https://github.com/baifengbai)聊天的时候，提到了搭建一套直播服务大概需要哪些技术，然后就大致构思了几个基础模块，但是其实我也没想好要做成啥样😂

![prototype](/demo.png)

## 结构内容
大致包括三个大模块：
- **推拉流模块**
    - RTMP 服务器，我自己为了方便找了一个 alifg/nginx-rtmp 镜像，在 docker 里跑的，随用随删
    - 推流 可以用专用软件直接推、ffmpeg 推或者 编写代码往推流地址上推
    - 拉流 因为我不会 APP 开发，对音视频编解码也不是很了解，就用了aliplayer 直接拉 flv 流了
- **web 模块** 主要用于网站构建，python 搭建原型感觉还是比较适合的，快速迭代
- **websocket 模块** Java 的 netty框架 或者 go 的gorilla/websocket模块都是比较适合的，高层次的 api，可以更专注于业务层开发。

然后网上搜了些现成的轮子，稍微改改，基本上一个原型就出来了。

## 使用流程

1. 开启 docker 容器，把rtmp 服务器跑起来。
```shell
docker pull alfg/nginx-rtmp
docker run -it -p 1935:1935 -p 8080:80 --rm 1ffba0323f13
```

2. 开启 web 服务器
```shell
cd /web
nnohup python3 server.py >> web.nohup.out &
```

3. 开启 websocket 服务器
```shell
cd websocket
go build server.go
nohup ./server >> websocket.nohup.out &
```

4. 推流
```shell
cd streamer/pushstreamer/target
java -jar pushstreamer-1.0-SNAPSHOT.jar rtmp://localhost:1935/stream/example 25
```

5. 开始测试
```shell
http://localhost:9999/live
```
