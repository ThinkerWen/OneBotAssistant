<div align="center">
    <img src="https://www.hive-net.cn/Assets/SiteGlobal/Hive_blank.png" width="200" alt="HeroPower"/>
    <h1>OneBotAssistant</h1>

[![PackageVersion](https://img.shields.io/badge/Code-Github-red)](https://github.com/ThinkerWen/OneBotAssistant)
[![PackageVersion](https://img.shields.io/badge/Go-1.22.1-blue)](https://go.dev/dl/)
[![PackageVersion](https://img.shields.io/badge/OneBot-11-brightgreen)](https://github.com/botuniverse/onebot-11)
[![PackageVersion](https://img.shields.io/badge/download-release-blue)](https://github.com/ThinkerWen/OneBotAssistant/releases)
</div>

----

## 介绍

本程序包含5种实用功能，使用 [配置文件](#配置文件) 和 Q群内指令 来一件控制开启关闭，功能如下：
1. [自动回复](#1自动回复)
2. [敏感词检测](#2敏感词检测)
3. [聊天机器人](#3聊天机器人)
4. [大学生网课搜题](#4大学生网课搜题)
5. [王者荣耀战力查询](#5王者荣耀战力查询)

功能详细和指令介绍请点击：[功能介绍](#功能)

本项目基于[OneBot-11](https://github.com/botuniverse/onebot-11)开发，项目代码采用Go语言开发，OneBot的Go语言SDK：[glib-onebot](https://github.com/ThinkerWen/glib-onebot)。

用户安装请先安装OneBot环境，如：[NapCatQQ](https://github.com/NapNeko/NapCatQQ)。

**如果对您有帮助可以点个Star，谢谢**！
****

## 安装

### Docker(推荐)
```shell
docker run -d \
  --restart=unless-stopped \
  -v ./application.yaml:/app/application.yaml \
  -v ./auto_reply_config.yaml:/app/auto_reply_config.yaml \
  -v ./one_bot_assistant.db:/app/one_bot_assistant.db \
  --name="OneBotAssistant" \
  designerwang/one-bot-assistant:latest
```

### linux
```bash
# 下载可执行文件
wget 'https://github.com/ThinkerWen/OneBotAssistant/releases/download/v0.1.0/OneBotAssistant.linux.amd64'
# 运行该文件
./OneBotAssistant.linux.amd64
```

### macOS & windows
1.从 [releases](https://github.com/ThinkerWen/OneBotAssistant/releases) 下载最新版本，对应你系统的可执行文件：

2.运行下载的可执行文件


****

## 功能

**功能的开启、关闭和参数配置见：[配置文件](#配置文件)**

### 1.自动回复
运行程序后会在当前目录自动创建auto_reply_config.yaml文件，其中包含自动回复的触发词、回复内容、作用范围。

当群聊消息中匹配到触发词后，Bot会自动回复对应的内容。触发词可以在QQ发送“添加自动回复”然后按流程添加。

自动回复的触发词中星号“*”代表模糊匹配，例如：
1. “你好”：全匹配，只能由“你好”触发
2. “\*你好”：后缀匹配，可以由“你好”、“你好啊”...触发
3. “你好*”和“\*你好\*”以此类推。

#### Bot触发指令：

| 命令     | 功能                 |
|--------|--------------------|
| 添加自动回复 | 用于添加新的自动回复，发送后开启流程 |
| 开启自动回复 | 发送后开启本群的自动回复       |
| 关闭自动回复 | 发送后关闭本群的自动回复       |

<br>

### 2.敏感词检测
敏感词检测默认检测不文明词汇，具有 反政,色情,反动,言语不当的词汇，包含两种检测模式：
1. 用户自定义检测
2. 在线API检测（免费）

敏感词检测为模糊匹配，即当消息文本中包含敏感词汇，Bot就会自动进行撤回（禁言）
#### Bot触发指令：
| 命令      | 功能                |
|---------|-------------------|
| 添加敏感词   | 用于添加新的敏感词，发送后开启流程 |
| 开启敏感词检测 | 发送后开启本群的敏感词检测     |
| 关闭敏感词检测 | 发送后关闭本群的敏感词检测     |

<br>

### 3.聊天机器人
聊天机器人使用 [茉莉云-机器人API](https://mlyai.com/) 其中包含许多已有的小功能，用户去官网申请API-KEY和API-SECRET后，添加到配置文件中即可（每日有免费额度）。
#### Bot触发指令：
| 命令      | 功能                |
|---------|-------------------|
| 开启聊天机器人 | 发送后开启本群的聊天机器人     |
| 关闭聊天机器人 | 发送后关闭本群的聊天机器人     |

<br>

### 4.大学生网课搜题
此功能用于查询大学生网课的答案，在Q群中发送网课题目，Bot会自动回复对应答案。（见[OnlineCourseAPI](https://github.com/ThinkerWen/OnlineCourseAPI)）
#### Bot触发指令：
| 命令     | 功能                          |
|--------|-----------------------------|
| 搜题 [题目]  | 发送后Bot会回复题目的答案（例：搜题 我国的国体是） |
| 搜题帮助   | 发送后Bot会回复搜题功能使用方法           |
| 开启网课搜题 | 发送后开启本群的网课搜题                |
| 关闭网课搜题 | 发送后关闭本群的网课搜题                |

<br>

### 5.王者荣耀战力查询
本功能用于王者荣耀英雄的最低战力地区，在农批玩家中较受欢迎
#### Bot触发指令：
| 命令                | 功能                                |
|-------------------|-----------------------------------|
| 查战力 [英雄名] [服务器大区] | 发送后Bot会回复对应英雄的最低战力（例：查战力 李白 安卓QQ） |
| 查战力帮助             | 发送后Bot会回复查战力功能使用方法                |
| 开启战力查询            | 发送后开启本群的战力查询                      |
| 关闭战力查询            | 发送后关闭本群的战力查询                      |
****

## 配置文件
程序的配置文件名为`application.yaml`会在第一次运行时自动创建，无需手动创建，以下是配置文件对应介绍
```yaml
api_url: ws://127.0.0.1:3001  # OneBot服务的地址
proxy: "" # 本程序中外部网络请求的代理，默认无代理（敏感词检测的API会用到）
hosts:    # 插件的管理员，所有设置指令只能由管理员触发
  - 296854007
auto_reply:   # 自动回复
  enable: true
  groups:     # 开启自动回复的群号（后续方法以此类推）
    - 123456
hero_power:   # 王者荣耀战力
  enable: true
  groups:
    - 123456
molly:        # 聊天机器人
  enable: true
  api_key: ""     # 茉莉云api_key
  api_secret: ""  # 茉莉云api_secret
  qq: 296854007   # 聊天机器人QQ
  name: 小风       # 茉莉云机器人名称
  groups:
    - 123456
online_course:    # 大学生网课搜题
  enable: true
  groups:
    - 123456
  limit: 1        # 网课答案展示限制条数
  token: free     # API token，默认free，有免费额度
sensitive:        # 敏感词检测
  enable: true
  alert_times: 3  # 警告次数，撤回敏感超过警告次数后则会禁言
  shut_seconds: 60  # 禁言时长（秒）
  groups:
    - 123456
```

****

## 提供帮助

#### 交流Q群：[103172845](https://h5.qun.qq.com/h5/jump-page/index.html?sid=2&isQim=false&src_type=internal&version=1&uin=103172845&card_type=group&source=qrcode&jump_from=&auth=&authSig=ssD9NFl2r5rHhGL4SvyIF56kSJi33zxFu2LqZ0XvUUGIZN3CyhCanNyji7cNXAwo&source_id=2_40001)