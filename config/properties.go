package config

//部署时切pro
//const Env = "pro"
//const dbUser = "root"
//const dbPwd = "123456"
//const dbIp = "127.0.0.1:3306"
//const dbName = "schema"

//开发
const Env = "dev"
const dbUser = "root"
const dbPwd = "123456"
const dbIp = "127.0.0.1:3306"
const dbName = "test"


/* 日志输出路径 */
const LogPath = "/var/log/go-web-gin-example"

/* https证书加载路径 */
const HttpsCertPem = "resource/https.key"
const HttpsCertKey = "resource/https.pem"

/* 支付宝验证url */
const AliPayNotifyUrl = "https://xxx.xxx.xxx/alipay/notifyurl"
const AliPayReturnUrl = "https://xxx.xxx.xxx/alipay/returnurl"

/* 微信支付通知接口 */
const NotifyUrl = "https://xxx.xxx.xxx/wxpay/notifyUrl"




