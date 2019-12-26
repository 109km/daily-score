package utils

import (
	"github.com/astaxie/beego/logs"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

type SecretConfig struct {
	SecretId  string `json:"SecretId"`
	SecretKey string `json:"SecretKey"`
}

var qcloudSecretConfig SecretConfig

var qcloudCPF *profile.ClientProfile
var qcloudClient *sms.Client

func init() {
	qcloudSecretConfig = LoadJSON("conf/qcloud.key")
	qcloudCPF = profile.NewClientProfile()
	qcloudCPF.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	credential := common.NewCredential(
		qcloudSecretConfig.SecretId,
		qcloudSecretConfig.SecretKey,
	)
	qcloudClient, _ = sms.NewClient(credential, "ap-beijing", qcloudCPF)
}

func SendSMS(params string) {

	request := sms.NewSendSmsRequest()
	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := qcloudClient.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		// fmt.Printf("An API error has returned: %s", err)
		logs.Error(-1, "TencentCloudSDKError: ", err.Error())
		return
	}
	if err != nil {
		logs.Error(-1, "utils.SendSMS.error:", err.Error())
		panic(err)
	}
	logs.Info("utils.SendSMS.response:", response.ToJsonString())
}

func GetMessageStatus() {

}
