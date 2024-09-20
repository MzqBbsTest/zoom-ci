# 项目名称：集群管理工具

## 简介

本项目是一个基于开源项目[zoom-ci](https://github.com/zoom-ci/zoom-ci)的二次开发，旨在方便管理和监控集群环境。通过提供简洁易用的界面和强大的功能，用户可以轻松地对集群进行部署、配置、监控和维护。

## 目录

- [背景](#背景)
- [新增功能](#新增功能)
- [安装](#安装)
- [许可证](#许可证)
- [致谢](#致谢)

## 背景

本项目是一个基于开源项目[zoom-ci](https://github.com/zoom-ci/zoom-ci)的二次开发，旨在方便管理集群环境。通过提供简洁易用的界面和强大的功能，用户可以轻松地对集群进行部署、配置和维护。


## 新增功能

- 集群管理
- 服务器管理
- WebSSH
- 屏蔽部署功能

## 安装

1、下载[最新版本release包](https://github.com/zoom-ci/zoom-ci/releases),并将其拷贝到任意目录（比如：~/zoom_workspace）并执行;

```shell
$ ./zoom-v1.0.3-darwin-amd64   # 这里以mac 64位版为例 

 _____________________________
       ___                    
         /                    
 -------/-----__----__---_--_-
       /    /   ) /   ) / /  )
 ____(_____(___/_(___/_/_/__/_
     /   Make CI/CD Easier.  
 (_ /                         


Service:              zoom
Version:              v1.0.3
Config Loaded:        ./zoom.ini
Log:                  stdout
Mail Enable:          0
HTTP Service:         :7002
Start Running...
```

2、打开浏览器，访问 `http://localhost:7002` (出现下图界面)，配置数据库及管理员信息，安装完成。
<p style="margin: 20px 0 40px 0;">
  <img height="500"  src="https://zoom-ci.github.io/docs/assets/img/zoom-install.png" />
</p>




## LICENSE

本项目采用 MIT 开源授权许可证，完整的授权说明已放置在 LICENSE 文件中。

## 致谢
特别感谢 [zoom-ci](https://github.com/zoom-ci/zoom-ci) 的支持和贡献。
