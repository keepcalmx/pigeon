# pigeon 

一个基于websocket的聊天系统，使用mysql+go+vue3开发。 [Live Demo](http://43.138.217.251)

![image](https://user-images.githubusercontent.com/100503668/201953166-834dda54-8f55-488d-b69f-ee006594f791.png)

### 已实现功能
- [x] 注册&登录
- [x] 在线/离线状态推送
- [x] 私聊&群聊

### 🗒️TODO
- [ ] 创建群组
- [ ] 个人资料修改
- [ ] 增加缓存优化查询
- [ ] 支持发送图片和表情
- [ ] 分布式部署支持

### ⚙️部署

克隆本项目到本地
```shell
git clone https://github.com/keepcalmx/pigeon.git
cd pigeon
```
修改前端配置文件里的VITE_PROD_URL为自己环境的ip（本地调试可忽略）
```shell
vim front-end/.env
```
构建
```shell
docker compose up
```
