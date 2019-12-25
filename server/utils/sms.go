package utils

import (
	"fmt"

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

func init() {
	qcloudSecretConfig = LoadJSON("conf/qcloud.key")
}

func SendSMS(params string) {
	credential := common.NewCredential(
		qcloudSecretConfig.SecretId,
		qcloudSecretConfig.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "ap-beijing", cpf)

	request := sms.NewSendSmsRequest()

	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}
