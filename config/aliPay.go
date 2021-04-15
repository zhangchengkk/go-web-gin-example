package config

import (
"fmt"
"github.com/smartwalle/alipay/v3"
"os"
)

//appid
const AppId = "2021001165605729"
/*支付宝生成证书位置*/
const AppPublicCert = "cert/xxxxxxx"
const AliPayRootCert = "cert/xxxxxx"
const AliPayPublicCert = "cert/xxxxxx"
// 可选，支付宝提供给我们用于签名验证的公钥，通过支付宝管理后台获取
const AliPublicKey = ""
// 必须，上一步中使用 RSA签名验签工具 生成的私钥
const PrivateKey = ""


var AliPayClient *alipay.Client
func initAliPay() {
	var err error
	AliPayClient, err = alipay.New(AppId, PrivateKey, true)
	if err != nil {
		fmt.Println("初始化支付宝失败, 错误信息为", err)
	}
	err = AliPayClient.LoadAppPublicCertFromFile(AppPublicCert)
	if err != nil {
		fmt.Println("1初始化支付宝失败, 错误信息为", err)
	}
	err = AliPayClient.LoadAliPayRootCertFromFile(AliPayRootCert)
	if err != nil {
		fmt.Println("2初始化支付宝失败, 错误信息为", err)
	}
	err = AliPayClient.LoadAliPayPublicCertFromFile(AliPayPublicCert)
	if err != nil {
		fmt.Println("3初始化支付宝失败, 错误信息为", err)
	}
	if err != nil {
		fmt.Println("初始化支付宝失败, 错误信息为", err)
		os.Exit(-1)
	}
}
