package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Message struct {
	Wxid               string `json:"wxid"`
	Content            string `json:"content"`
	ToUser             string `json:"toUser"`
	Msgid              uint64 `json:"msgid"`
	OriginMsg          string `json:"originMsg"`
	ChatRoomSourceWxid string `json:"chatRoomSourceWxid"`
	MsgSource          string `json:"msgSource"`
	Type               uint32 `json:"type"`
	DisplayMsg         string `json:"displayMsg"`
	ImgData            string `json:"imgData"`
}

type PublicMessage struct {
	Data  []map[string]string `json:"data"`
	Total int                 `json:"total"`
	Wxid  string              `json:"wxid"`
}

func wsClient() {
	// /ws/generalMsg
	// /ws/publicMsg
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws/publicMsg"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("ReadMessage:", err)
			return
		}
		log.Printf("recv: %s", message)

		// var public_message_map map[string]interface{}
		// var pm PublicMessage

		// json.Unmarshal(public_message, &public_message_map)
		// json.Unmarshal(message, &pm)

		// log.Printf("content: %s\n", public_message_map["data"].([]interface{})[0].(map[string]interface{})["Content"])
		// log.Printf("content: %s\n", pm.Data[0]["Content"])
	}
}

func httpServer() {
	g := gin.Default()
	g.POST("/callback", func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			log.Printf("GetRawData faild: %s\n", err)
			return
		}
		log.Printf("data: %s\n", data)

		// var msg Message
		// err := c.BindJSON(&msg)
		// if err != nil {
		// 	log.Printf("bind json faild: %s\n", err)
		// 	return
		// }
		// log.Printf("msg: %+v\n", msg)
	})

	g.Run(*addr)
}

func escapeQuotes(s string) string {
	var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")
	return quoteEscaper.Replace(s)
}

func sendFormImg() {
	var client http.Client

	// 要上传的文件
	file, _ := os.Open(*img_path)
	defer file.Close()

	// 设置body数据并写入缓冲区
	bodyBuff := bytes.NewBufferString("") //bodyBuff := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuff)

	// math/rand
	// _ = bodyWriter.SetBoundary(fmt.Sprintf("-----------------------------%d", rand.Int()))
	// 加入图片二进制
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, escapeQuotes("image"), escapeQuotes(filepath.Base(file.Name()))))
	// h.Set("Content-Disposition", fmt.Sprintf(`form-data; name=image; filename="%s"`, escapeQuotes(filepath.Base(file.Name()))))
	h.Set("Content-Type", "image/jpg")
	part, err := bodyWriter.CreatePart(h)
	if err != nil {
		log.Fatalf("WriteField faild: %s\n", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		log.Fatalf("WriteField faild: %s\n", err)
	}
	// 其他字段
	err = bodyWriter.WriteField("wxid", *wxid)
	if err != nil {
		log.Fatalf("WriteField faild: %s\n", err)
	}

	// err = bodyWriter.WriteField("clear", "false")
	// if err != nil {
	// 	log.Fatalf("CreatePart faild: %s\n", err)
	// }

	// err = bodyWriter.WriteField("path", "")
	// if err != nil {
	// 	log.Fatalf("CreatePart faild: %s\n", err)
	// }

	// 填充boundary结尾
	bodyWriter.Close()

	// 组合创建数据包
	req, err := http.NewRequest("POST", *addr+"/api/sendimgmsg", bodyBuff)
	if err != nil {
		log.Fatalf("NewRequest faild: %s\n", err)
	}

	req.ContentLength = int64(bodyBuff.Len())
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())

	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("client Do faild: %s\n", err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("ReadAll faild: %s\n", err)
	}

	fmt.Printf("form-img send msg response: %s\n", data)
}

func sendJsonImg() {
	var client http.Client

	type ImgInfo struct {
		Wxid      string `json:"wxid"`
		Path      string `json:"path"`
		Image     []byte `json:"image"`
		ImageName string `json:"imageName"`
		Clear     bool   `json:"clear"`
	}

	data, err := ioutil.ReadFile(*img_path)
	if err != nil {
		log.Fatalf("read file faild: %s\n", err)
	}

	ii := ImgInfo{
		Wxid: *wxid,
		// Path: "D:\\xxxxxxxxxxxxxxxx",
		Image:     data,
		ImageName: "1.jpg",
		Clear:     false,
	}

	j_data, err := json.Marshal(ii)
	if err != nil {
		log.Fatalf("json Marshal faild: %s\n", err)
	}

	log.Printf("json data: %s\n", j_data)

	req, err := http.NewRequest("POST", *addr+"/api/sendimgmsg", bytes.NewReader(j_data))
	if err != nil {
		log.Fatalf("NewRequest faild: %s\n", err)
	}
	// req.ContentLength = int64(bodyBuff.Len())
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("client Do faild: %s\n", err)
	}

	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("ReadAll faild: %s\n", err)
	}

	fmt.Printf("json-img send msg response: %s\n", data)
}

func sendFormFile() {
	var client http.Client

	// 要上传的文件
	file, _ := os.Open(*file_path)
	defer file.Close()

	// 设置body数据并写入缓冲区
	bodyBuff := bytes.NewBufferString("") //bodyBuff := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuff)

	// math/rand
	// _ = bodyWriter.SetBoundary(fmt.Sprintf("-----------------------------%d", rand.Int()))
	// 加入图片二进制
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, escapeQuotes("file"), escapeQuotes(filepath.Base(file.Name()))))
	h.Set("Content-Type", "text/plain")
	part, err := bodyWriter.CreatePart(h)
	if err != nil {
		log.Fatalf("WriteField faild: %s\n", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		log.Fatalf("WriteField faild: %s\n", err)
	}
	// 其他字段
	err = bodyWriter.WriteField("wxid", *wxid)
	if err != nil {
		log.Fatalf("WriteField faild: %s\n", err)
	}

	err = bodyWriter.WriteField("path", "D:\\temp\\10.txt")
	if err != nil {
		log.Fatalf("WriteField faild: %s\n", err)
	}

	// 填充boundary结尾
	bodyWriter.Close()

	// 组合创建数据包
	req, err := http.NewRequest("POST", *addr+"/api/sendfilemsg", bodyBuff)
	if err != nil {
		log.Fatalf("NewRequest faild: %s\n", err)
	}

	req.ContentLength = int64(bodyBuff.Len())
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())

	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("client Do faild: %s\n", err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("ReadAll faild: %s\n", err)
	}

	fmt.Printf("form-file send msg response: %s\n", data)
}

type FileMsg struct {
	Wxid     string `json:"wxid"`
	Path     string `json:"path,omitempty"`
	File     []byte `json:"file"`
	FileName string `json:"fileName"`
}

func sendJsonFile() {
	var client http.Client

	data, err := ioutil.ReadFile(*file_path)
	if err != nil {
		log.Fatalf("read file faild: %s\n", err)
	}

	fm := FileMsg{
		Wxid: *wxid,
		// Path: "D:\\xxxxxxxxxxxxxxxxxxx",
		File:     data,
		FileName: path.Base(filepath.ToSlash(*file_path)),
		// Clear: false,
	}

	j_data, err := json.Marshal(fm)
	if err != nil {
		log.Fatalf("json Marshal faild: %s\n", err)
	}

	fmt.Printf("j_data: %s\n", j_data)

	req, err := http.NewRequest("POST", *addr+"/api/sendfilemsg", bytes.NewReader(j_data))
	if err != nil {
		log.Fatalf("NewRequest faild: %s\n", err)
	}
	// req.ContentLength = int64(bodyBuff.Len())
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("client Do faild: %s\n", err)
	}

	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("ReadAll faild: %s\n", err)
	}

	fmt.Printf("json-file send msg response: %s\n", data)
}

type Account struct {
	CustomAccount       string `json:"customAccount"`
	Nickname            string `json:"nickname"`
	Note                string `json:"note"`
	Pinyin              string `json:"pinyin"`
	PinyinAll           string `json:"pinyinAll"`
	ProfilePicture      string `json:"profilePicture"`
	ProfilePictureSmall string `json:"profilePictureSmall"`
}

type AccountBody struct {
	Code    int     `json:"code"`
	Data    Account `json:"data"`
	Message string  `json:"message"`
}

func accountByWxid(wxid string) Account {
	var client http.Client
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/accountbywxid?wxid=%s", *addr, wxid), nil)
	if err != nil {
		log.Fatalf("json Marshal faild: %s\n", err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("client Do faild: %s\n", err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("ReadAll faild: %s\n", err)
	}

	ab := AccountBody{}
	err = json.Unmarshal(data, &ab)
	if err != nil {
		log.Fatalf("Unmarshal faild: %s\n", err)
	}

	return ab.Data
}

type Contact struct {
	CustomAccount       string `json:"customAccount"`
	EncryptName         string `json:"encryptName"`
	Nickname            string `json:"nickname"`
	Note                string `json:"note"`
	NotePinyin          string `json:"notePinyin"`
	NotePinyinAll       string `json:"notePinyinAll"`
	Pinyin              string `json:"pinyin"`
	PinyinAll           string `json:"pinyinAll"`
	ProfilePicture      string `json:"profilePicture"`
	ProfilePictureSmall string `json:"profilePictureSmall"`
	Reserved1           string `json:"reserved1"`
	Reserved2           string `json:"reserved2"`
	Type                string `json:"type"`
	VerifyFlag          string `json:"verifyFlag"`
	Wxid                string `json:"wxid"`
}

type ContactBody struct {
	Code int `json:"code"`
	Data struct {
		Contacts []Contact `json:"contacts"`
		Total    int       `json:"total"`
	} `json:"data"`
	Message string `json:"message"`
}

func getContacts() {
	var client http.Client

	req, err := http.NewRequest("GET", *addr+"/api/contacts", nil)
	if err != nil {
		log.Fatalf("json Marshal faild: %s\n", err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("client Do faild: %s\n", err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("ReadAll faild: %s\n", err)
	}

	cb := ContactBody{}
	err = json.Unmarshal(data, &cb)
	if err != nil {
		log.Fatalf("Unmarshal faild: %s\n", err)
	}

	fmt.Printf("data: %+v\n", cb)

	for i := 0; i < cb.Data.Total; i++ {
		c := accountByWxid(cb.Data.Contacts[i].Wxid)
		cb.Data.Contacts[i].ProfilePicture = c.ProfilePicture
		cb.Data.Contacts[i].ProfilePictureSmall = c.ProfilePictureSmall
	}

	fmt.Printf("%+v\n", cb.Data.Contacts)
}

var addr = flag.String("addr", "localhost:8080", "Http service address")
var mode = flag.String("mode", "ws", "Select the startup mode. The optional values are ws, http, form-img, json-img, form-file, json-file and contacts")
var img_path = flag.String("img", "../public/1.jpg", "Specify image path when sending image messages")
var file_path = flag.String("file", "../public/1.txt", "Send file message specifying file path")
var wxid = flag.String("wxid", "", "Send message recipient's wxid")

func main() {
	flag.Parse()
	log.SetFlags(0)

	switch *mode {
	case "ws":
		wsClient()
	case "http":
		httpServer()
	case "form-img":
		sendFormImg()
	case "json-img":
		sendJsonImg()
	case "form-file":
		sendFormFile()
	case "json-file":
		sendJsonFile()
	case "contacts":
		getContacts()
	}
}
