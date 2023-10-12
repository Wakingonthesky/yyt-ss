package services

import (
	"fmt"
	"math/rand"
	"time"
	"ytt-societyservice/config"

	unisms "github.com/apistd/uni-go-sdk/sms"
)

func Code(info string) error {
	accessKeyId := config.C.GetString("AccessKeyID")

	fmt.Println(accessKeyId)

	client := unisms.NewClient(accessKeyId)

	rand.Seed(time.Now().UnixNano())

	vcode := fmt.Sprintf("%06v", rand.Int31n(1000000))

	message := unisms.BuildMessage()

	message.SetTo(info)
	message.SetSignature("重邮天文协会")
	message.SetTemplateId("pub_verif_basic2")
	message.SetTemplateData(map[string]string{"code": vcode})

	_, err := client.Send(message)

	return err
}
