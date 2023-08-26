package dechline

import (
	"bytes"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

type service struct {
	isMock         bool
	isShowResponse bool
	isUseProxy     bool
	proxyUrl       string
}

func New(isMock bool, isShowResponse bool, isUseProxy bool, proxyUrl string) *service {
	return &service{
		isMock:         isMock,
		isShowResponse: isShowResponse,
		proxyUrl:       proxyUrl,
	}
}

func (rcv *service) SendLineNotifyMsg(msg string, token string) {
	client := resty.New()

	if rcv.isUseProxy {
		// Setting a Proxy URL and Port
		client.SetProxy(rcv.proxyUrl)
	}

	resp, err := client.R().
		// EnableTrace().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("Authorization", token).
		SetFormData(map[string]string{
			"message": msg,
		}).
		Post("https://notify-api.line.me/api/notify")

	if err != nil {
		fmt.Println(err)
	}

	// Explore response object
	if rcv.isShowResponse {
		fmt.Println("Response Info:")
		fmt.Println("  Error      :", err)
		fmt.Println("  Status Code:", resp.StatusCode())
	}

}

func (rcv *service) SendLineNotifyMsgList(msg []string, token string) {
	msgToSend := ""
	for _, v := range msg {
		msgToSend = msgToSend + v
	}

	if !rcv.isMock {
		rcv.SendLineNotifyMsg(msgToSend, token)
	} else {
		fmt.Println("--- Example Send Line ----")
		fmt.Println(msgToSend)
		fmt.Println("--- Example Send Line ----")
	}
}

func (rcv *service) SendLineNotifyMsgAndImage(msg string, fileName string, token string) {
	if !rcv.isMock {
		rcv.sendLineNotifyMsgAndImage(msg, fileName, token)
	} else {
		fmt.Println("--- Example Send Line ----")
		fmt.Println(msg)
		fmt.Println("--- Example Send Line ----")
	}
}

func (rcv *service) SendLineNotifyMsgAndImageSteam(msg string, profileImgBytes []byte, token string) {
	if !rcv.isMock {
		rcv.sendLineNotifyMsgAndImageSteam(msg, profileImgBytes, token)
	} else {
		fmt.Println("--- Example Send Line ----")
		fmt.Println(msg)
		fmt.Println("--- Example Send Line ----")
	}
}

func (rcv *service) sendLineNotifyMsgAndImage(msg string, fileName string, token string) {
	profileImgBytes, _ := os.ReadFile(fileName)

	client := resty.New()

	// Setting a Proxy URL and Port
	if rcv.isUseProxy {
		// client.SetProxy(cont.Proxy)
		client.SetProxy(rcv.proxyUrl)
	}

	resp, err := client.R().
		// EnableTrace().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("Authorization", token).
		SetFileReader("imageFile", "img.png", bytes.NewReader(profileImgBytes)).
		SetFormData(map[string]string{
			"message": msg,
		}).
		Post("https://notify-api.line.me/api/notify")

	if err != nil {
		fmt.Println(err)
	}

	// Explore response object
	if rcv.isShowResponse {
		fmt.Println("Response Info:")
		fmt.Println("  Error      :", err)
		fmt.Println("  Status Code:", resp.StatusCode())
	}

}

func (rcv *service) sendLineNotifyMsgAndImageSteam(msg string, profileImgBytes []byte, token string) {
	// profileImgBytes, _ := os.ReadFile(fileName)

	client := resty.New()

	// Setting a Proxy URL and Port
	if rcv.isUseProxy {
		// client.SetProxy(cont.Proxy)
		client.SetProxy(rcv.proxyUrl)
	}

	resp, err := client.R().
		// EnableTrace().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("Authorization", token).
		SetFileReader("imageFile", "img.png", bytes.NewReader(profileImgBytes)).
		SetFormData(map[string]string{
			"message": msg,
		}).
		Post("https://notify-api.line.me/api/notify")

	if err != nil {
		fmt.Println(err)
	}

	// Explore response object
	if rcv.isShowResponse {
		fmt.Println("Response Info:")
		fmt.Println("  Error      :", err)
		fmt.Println("  Status Code:", resp.StatusCode())
	}

}
