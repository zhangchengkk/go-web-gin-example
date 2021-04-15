## 自用 简易golang web framework

- gin 
- mysql - xorm 
- redis接入限流 
- 阿里云短信服务 
- 微信支付 
- 支付宝支付

```$xslt
|-api 接口，按用户角色分包。格式对应router
|-service 业务逻辑
|-common 通用代码、工具类
  |-httpCommonRes 通用接口返回数据结构
|-models xorm对应的持久对象
|-config 
  |-properties 通用变量设置
  |-...xxx.go 一些启动初始化的服务
|-resource 一些静态资源，例如https证书、支付吧证书
|-router
  |-handler 拦截器（包含tokne验证、redis滑动时间限流）
  |-...modules按api角色分模块的路由注册表
|-main.go 根据env（dev/pro）初始化config中服务、
          sync xorm models、
          跨域设置、
          pro下https加载
```

