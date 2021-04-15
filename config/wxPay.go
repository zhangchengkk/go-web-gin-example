package config

import (
	"context"
	"crypto/x509"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"net/http"
	"time"
)

// 商户证书序列号（CA api证书同时包含商户私钥文件，下方）
const mchCertificateSerialNumber = ""
// 商户私钥文件路径
const privateKeyPath = "cert/wxPay/xxxxx_key.pem"
// 微信支付平台证书文件路径
const wechatCertificatePath= "cert/wxPay/wechatpay_xxxxxxx.pem"
//商户id
const MchId = "xxxxx"

var WxPayCtx context.Context
var WxPayClient *core.Client

func initWxPay(){
	// 加载商户私钥
	var err error
	privateKey, err := utils.LoadPrivateKeyWithPath(privateKeyPath)
	if err != nil {
		fmt.Fprintf(gin.DefaultWriter,"load private err:%s",err)
	}
	// 加载微信支付平台证书
	wechatPayCertificate, err := utils.LoadCertificateWithPath(wechatCertificatePath)
	if err != nil {
		fmt.Fprintf(gin.DefaultWriter,"load certificate err:%s",err)
	}
	// 增加客户端配置
	WxPayCtx = context.Background()
	opts := []option.ClientOption{
		option.WithMerchant(MchId, mchCertificateSerialNumber , privateKey), // 设置商户信息，用于生成签名信息
		option.WithWechatPay([]*x509.Certificate{wechatPayCertificate}),  // 设置微信支付平台证书信息，对回包进行校验
		option.WithHTTPClient(&http.Client{}),  // 可以不设置
		option.WithTimeout(2 * time.Second),    // 自行进行超时时间配置
		option.WithHeader(&http.Header{}),      // 可以自行设置Header
	}
	WxPayClient, err = core.NewClient(WxPayCtx, opts...)
	if err != nil {
		//fmt.Println("new wechat pay client err:%s", err.Error())
		fmt.Fprintf(gin.DefaultWriter, err.Error())
	}

}
