# ams
Automatic monitoring system

# 模块规范（etcd存储路径）

## 主页注册地址

> /ams/main/index/

## 页面模块（key）

以配置管理为例,以模块名作为key：

> config

## 页面模块值（value）

配置的值以：Name(显示的名称):HTML(html bootstrap可以显示的效果为主，至少包含<a href="#"></a>)

> 配置管理:<a href="/config"><button class="btn btn-success">配置管理</button></a>