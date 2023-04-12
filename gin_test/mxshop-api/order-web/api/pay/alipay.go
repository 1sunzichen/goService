package pay

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"go.uber.org/zap"
	"gopro/gin_test/mxshop-api/order-web/global"
	"gopro/gin_test/mxshop-api/order-web/proto"
	"net/http"
	"strconv"
)
func main(){
	appId:="2021000122675806"
	privateKey:="MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCDw1ka+t4vT0E7OTUZqVklDC6P8AvXevzWVP36u6N9j6+S3apUAqqvSVNsGt0OVwosk1gRACgs6JCazir5irJsKJ85xB4O0Q+DNeCBRwjAZCXANKYsk+34WwB6yEJr/KYCdnLJ5ftXVZWmb+8fdS4g11HNpsbzxAIQR/Eqj6bRUSfs9crbzxjNyLG0TQVCmV1d1jTztph1utUF9eOCb58mdujcgyM/rSx8AZC0DY8NrsAd1rKCEN1Ot6t+REawoTfgbab1aAx5QS95PZwOGNFSjTlbgvABlbsEDLsAIP/fVpbW3ZaYpTc+odAaEDOboPqYZezhsmecRa9sVUItXrF3AgMBAAECggEAHytD1dUIYCqgZLEYtZRQ6SmjjhsbwgJu8wix9/ERMK+fud9D3pGu7L4sxMLqDe1bE8ZbK9JcrprpMiWZFuTPZjSJzfFtptWflMtW613xXQmTwI7zzFHGFlz4lRLwW3ktkCGS792+gh9VwkRyTX+7xLsKt9o+8AUq880A6K+Ip7TwwnaIQNzjStzNt2zMamo4WZ/9LuN/89DMonKhFA6FXyW27PT9zluWjf48tVEiRS/49S7y6Y52Fs0aNeMDX5EImqZXg+NAk6sWHpcI3gR3uxlmLxk+8vpWqj6fsAUG8E6GJWGFUcC1G251wb9zIzleWrJXeyWBVq3vfxuBlKw6kQKBgQDnbDCsR6SMBkcUtgQMCgeZzeK/gaYF8BfgRNIkL6VPQfmbH9VkseIa86EJq6L2U8qEYSwSr+L1c4J79vF6eWw0G1rHc5lzyribzI5ymzBJjTSdrgKXPTUe8TmnfoO4702RyEik+6Yq/Q8O8/uzTnkyMMbpyzBUUPq1qAWuh1TOyQKBgQCRwak3n1OdcesOJxR7RfenEER1N2ZG7+bAMJCGn4hRv1m3kub/pGBdBS9sRdSiCLydzgGyBsRPs6f3hFS1e/oLUksQWTjZdpwZQMUIs+5DtjOdTeagh1MGgZriQQ2ypV9kNL+q7vMUDEbTLWzH/UpeLQ4GjC3zUnA+EElsLgJePwKBgQCea0p6dOSoUhfQjrUAhNElMXJ6REcho3TEunfb+52/PtuenFEZCEhOyN5BX3RECaIFsvtXo33LJpJ5R9eQTpSKqvses/yk7m4ngQU2YRPSFc6h7h+p5mV51AnypcGIFJDWLfPEtNvQa8EmLFDuMtb2S7uvCcAAyBqHxgh1rACbOQKBgQCD6yYvXHl+D2OxvHcSF0JMpzF+cXSVEX3kRlAYN/1WF3yo5EFD8M7ygcXpFc6cFJI5tQDd0rgMdsq3/8H3O80UQBgGJOqKD4q6ZF+wP8GO8TIH1kC8252uTtESo9Q08u3CMOekWn4QkAfuC7ffzYRodhiynl7cUama0nzRd1bXWwKBgQDKAijC5QiIO55cAXmdeIjoTJ0G3TUckKl0ZSoTJHTP+n4JVrMMyQAeAnP36sKwvpwnN4QoQBEuYDsvNa8ZFcGga5hnPZ9VFpYkAJEJAlW5I4jsGeiYHgkciOig2S2ORfRwVDPz2JLqXDTBhxYcV8bldfZ+6mJTOBRTViRvP04MTw=="
	alipayPubKey:="MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAmSzJIB3B0wlu46m8tNjlm9h3cF+8fExKQYBvi5HOeUbuZdnlpOqSWYxYxZmVo+oDlnLi5J6cboXqm2fF2PhwbgZI8oWb1Hqpfi9RildqOr2AZXs57/zF7sWXzx6J6WDnt3eBvm4JL3DVRI87PTC1qn7vjtGPa2swSqUR/fwHYPZaTXf359a+Ic2MkIIab+USx0B+P3Qos9P5ayIOuQDrJesjTMAOAbqRweww7QJTTPbewrTP+DCQ6w6EHZPz288IDKV0DgUA+OSbPDlxhR0wzWki6KQQpqzLkjSpZj/Nfj2WzKOQGmtfSZolxFzmFyIJ03vbiYu6WbXDZoj/PoVu2wIDAQAB"
	var client, err= alipay.New(appId, privateKey, false)
	if err!=nil{
		panic(err)
	}
	err=client.LoadAliPayPublicKey(alipayPubKey)
	if err!=nil{
		panic(err)
	}

	var p = alipay.TradePagePay{}
	p.NotifyURL = "https://3x20u71045.zicp.fun/o/v1/pay/alipay/notify"
	p.ReturnURL = "127.0.0.1:8089"
	p.Subject = "子宸生鲜"
	p.OutTradeNo = "传递一个唯一单号"
	p.TotalAmount = "10.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err:= client.TradePagePay(p)
	if err != nil {
		fmt.Println(err)
	}

	var payURL = url.String()
	fmt.Println(payURL)
}
func ConnAliPay()(*alipay.Client,error){
	var aliInfo=global.ServerConfig.AliPayInfo
	fmt.Println(aliInfo.AppID,"aliInfo.AppID")
	var client, err= alipay.New(aliInfo.AppID, aliInfo.PrivateKey, false)
	if err!=nil{
		//fmt.Println(err)
		zap.S().Errorw("支付宝失败")
		//ctx.JSON(http.StatusInternalServerError,gin.H{
		//	"msg":err.Error(),
		//})
		return nil,err
		//panic(err)
	}
	err=client.LoadAliPayPublicKey(aliInfo.AliPublicKey)
	if err!=nil{
		fmt.Println(err)
		zap.S().Errorw("支付宝失败")
		//ctx.JSON(http.StatusInternalServerError,gin.H{
		//	"msg":err.Error(),
		//})
		return nil,err
		//panic(err)
	}
	return  client,nil
}
func GenPayUrl(rsp *proto.OrderInfoResponse)(string,error){
	var aliInfo=global.ServerConfig.AliPayInfo
	client,err:=ConnAliPay()
	if err!=nil{
		return "",err
	}
	var p = alipay.TradePagePay{}
	p.NotifyURL = aliInfo.NotifyUrl
		//"https://3x20u71045.zicp.fun/o/v1/pay/alipay/notify"
	p.ReturnURL = aliInfo.ReturnUrl
	p.Subject = "子宸生鲜-"+rsp.OrderSn
	p.OutTradeNo = rsp.OrderSn
	p.TotalAmount = strconv.FormatFloat(float64(rsp.Total),'f',2,64)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err:= client.TradePagePay(p)
	if err != nil {
		//fmt.Println(err)
		zap.S().Errorw("支付宝失败")

		//ctx.JSON(http.StatusInternalServerError,gin.H{
		//	"msg":err.Error(),
		//})
		return "",err
	}

	var payURL = url.String()
	return payURL,nil
}
func Notify(c *gin.Context){
	//支付宝回调通知
	//var aliInfo=global.ServerConfig.AliPayInfo
	client,err:=ConnAliPay()
	if err!=nil{
		//return "",err

	}
	 noti,err:=client.GetTradeNotification(c.Request)
	if err!=nil{
		fmt.Println("交易状态为：",noti.TradeStatus)
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":"不合法的通知",
		})
	}
	_,err=global.OrderClient.UpdateOrderStatus(context.Background(),&proto.OrderStatus{
		OrderSn: noti.OutTradeNo,
		Status: string(noti.TradeStatus),

	})
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{})
	}
	c.String(http.StatusOK,"success")


}