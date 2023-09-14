# pointSwitch
毕业设计-转辙机监控系统开发



# 模块介绍

## global

加载全局配置,环境 --> 生成全局变量

- Logger : 全局日志文件记录
- Port : lora模块端口, 通过串口收发消息
- DB : gorm.DB数据库
- Addrs : int数组, 记录所有配置的终端lora地址, 0-65535
- States : map[int]int 记录所有配置的lora地址信息
- MDis: map[int]\[]int 记录终端lora传递的测量数据
- MTimes : map[int]\[]string  记录所有终端传递距离数据的时间

## controller

前端请求控制层, 用来与模型层进行交互, 完成后端逻辑

> dis

- get : 展示所有配置的地址的 dis 界面
- post : 获取前端传递的地址**(单个)**,  通过lora获取数据, 更新到全局常量中, 并返回前端

> display

- get : 展示所有配置的地址的已存储





## routers

