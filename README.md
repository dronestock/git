# git

`Drone`持续集成`Git`插件，功能

- 内置`Github`加速
- 同时支持`Push`和`Pull`
- 支持`SSH`无密码连接
- 支持`Tag`打包
- 多仓库支持
- 支持通用环境变量

## 使用

非常简单，只需要在`.drone.yml`里增加配置

```yaml
clone:
  disable: true


steps:
  - name: 代码
    image: ccr.ccs.tencentyun.com/dronestock/git
```

更多使用教程，请参考[使用文档](https://www.dronestock.tech/plugin/stock/git)

## 交流

![微信群](https://www.dronestock.tech/communication/wxwork.jpg)

## 捐助

![支持宝](https://github.com/storezhang/donate/raw/master/alipay-small.jpg)
![微信](https://github.com/storezhang/donate/raw/master/weipay-small.jpg)

## 感谢`Jetbrains`

本项目通过`Jetbrains开源许可IDE`编写源代码，特此感谢

[![Jetbrains图标](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg)](https://www.jetbrains.com/?from=dronestock/cos)
