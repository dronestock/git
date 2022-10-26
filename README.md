# git

Drone持续集成Git插件，功能

- 内置`Github`加速
- 同时支持`Push`和`Pull`
- 支持`SSH`无密码连接
- 支持`Tag`打包
- 多仓库支持
- 支持通用环境变量

## 使用

非常简单，只需要在`.drone.yaml`里增加配置

### 拉代码

使用本插件拉代码的最直接的原因是，在某些环境不好的机器上（比如我家的网络：成都电信宽带）就经常出现`Github`访问不了的问题

```yaml
clone:
  disable: true


steps:
  - name: 拉代码
    image: dronestock/git
    pull: always
    settings:
      verbose: true
      debug: true
```

### 推代码

推代码建议使用`SSH`而不是密码方式，更简单省事

```yaml
steps:
  - name: 推送Dart版本
    image: storezhang/git
    pull: always
    settings:
      remote: git@gitea.com:xxx/yyy.git
      folder: $${DART}
      force: true
      commit_message: ${DRONE_COMMIT_MESSAGE}
      tag: $${VERSION}
      ssh_key:
        from_secret: gitea_ssh_key
```


## 捐助

![支持宝](https://github.com/storezhang/donate/raw/master/alipay-small.jpg)
![微信](https://github.com/storezhang/donate/raw/master/weipay-small.jpg)

## 感谢Jetbrains

本项目通过`Jetbrains开源许可IDE`编写源代码，特此感谢
[![Jetbrains图标](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.png)](https://www.jetbrains.com/?from=dronestock/git)
