# ams
Automatic monitoring system

# 模块规范（etcd存储路径）

## 主页注册地址

> /ams/main/index/

```
/ams  介绍:ams项目配置主路径
/ams/main  介绍:主页配置路径
/ams/main/index 介绍:主页网站跳转配置
/ams/main/backend 介绍:标签页管理
/ams/main/services 介绍::服务管理
/ams/main/ansible 介绍:配置管理机 提供服务器ip，ip api,ansible pub key等
/ams/main/ansible/ip 介绍:配置管理机 提供ip感知功能，用于多ip，复杂配置获取可访问ip的功能
```

## 页面模块（key）

以配置管理为例,以模块名作为key：

> config

## 页面模块值（value）

配置的值以：Name(显示的名称):HTML(html bootstrap可以显示的效果为主，至少包含<a href="#"></a>)

> 配置管理:<a href="/config"><button class="btn btn-success">配置管理</button></a>

## ansible 中控机

> ansible

提供ansible api ip pub_key等信息 供初始化服务调用