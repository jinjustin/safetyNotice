package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	//"image"
	//"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	//"reflect"

	"github.com/jasonlvhit/gocron"
	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/jinjustin/linebot/storemessage"
)

//LineMessage is a message that incoming from webhook
type LineMessage struct {
	Destination string `json:"destination"`
	Events      []struct {
		ReplyToken string `json:"replyToken"`
		Type       string `json:"type"`
		Timestamp  int64  `json:"timestamp"`
		Source     struct {
			Type   string `json:"type"`
			UserID string `json:"userId"`
		} `json:"source"`
		Message struct {
			ID   string `json:"id"`
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"events"`
}

//ReplyMessage is a event ที่จะส่งไปให้ Line Replytoken
type ReplyMessage struct {
	ReplyToken string `json:"replyToken"`
	Messages   []Text `json:"messages"`
}

//Text is a struct that represent a text
type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

//Profile is a struct that represent user profile
type Profile struct {
	UserID        string `json:"userID"`
	DisplayName   string `json:"displayName"`
	PictureURL    string `json:"pictureURL"`
	StatusMessage string `json:"statusMessage"`
}

//ImageInQue is a struct that use to represent image and time to send that image.
type ImageInQue struct {
	Time      string `json:"time"`
	ImageLink string `json:"imageLink"`
}

//ChannelToken is a token from line developer channel
var ChannelToken = "wXBBVwTFuP/RewWyGJEQGKQG5fuG+NwMhkElufZsGk9TVgOHCm2Fdg8/MkVhgFlZ1QtOtSV4FRck/+2pQM6JFQcuIQ+43fOIHJzbCKIX3eukzXxc7TLaBnOQL5to/DAo3hwfHaM/4KaYLYpjXpXXEgdB04t89/1O/w1cDnyilFU="

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	e.GET("/test", func(c echo.Context) error {
		p := readUsersFile()
		fmt.Println(p)
		return c.String(http.StatusOK, "ok")
	})

	e.POST("/webhook", func(c echo.Context) error {

		client := &http.Client{}
		bot, err := linebot.New("0dcc21819b9f37988e430c810cf1c849", ChannelToken, linebot.WithHTTPClient(client))
		if err != nil {
			return err
		}

		Line := new(LineMessage)
		if err := c.Bind(Line); err != nil {
			log.Println("err")
			return c.String(http.StatusOK, "error")
		}

		for _, event := range Line.Events {
			if event.Type == "message" {
				if event.Message.Type == "text" {
					if event.Message.Text == "Admin" {
						users := readUsersFile()
						p := getProfile(event.Source.UserID)
						check := true
						for _, user := range users {
							if user.UserID == p.UserID {
								check = false
							}
						}
						if check {
							users = append(users, p)
						}
						err := writeUsersFile(users)
						if err != nil {
							return err
						}
						// Unmarshal JSON
						flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(storemessage.InstructionMessage))
						if err != nil {
							log.Println(err)
						}
						// New Flex Message
						flexMessage := linebot.NewFlexMessage("คำสั่งผู้ดูแลระบบ", flexContainer)
						// Reply Message
						bot.ReplyMessage(event.ReplyToken, flexMessage).Do()
					} else {
						if authentication(event.Source.UserID) {
							words := strings.Fields(event.Message.Text)
							switch words[0] {
							case "ADD":
								result := checkTimeCorrection(words[1])
								if result == "" {
									i := ImageInQue{
										Time:      words[1] + ":00",
										ImageLink: words[2],
									}
									imageInQues := readImagesFile()
									imageInQues = append(imageInQues, i)
									sortedImage := sortImage(imageInQues)
									err := writeImageInQue(sortedImage)
									if err != nil {
										return err
									}
									message := "เพิ่มกำหนดการส่งรูปภาพที่เวลา " + words[1] + " สำเร็จ"
									textMessage := linebot.NewTextMessage(message)
									bot.ReplyMessage(event.ReplyToken, textMessage).Do()
								} else {
									textMessage := linebot.NewTextMessage(result)
									bot.ReplyMessage(event.ReplyToken, textMessage).Do()
								}
							case "DELETE":
								result := checkTimeCorrection(words[1])
								if result == "" {
									imageInQues := readImagesFile()
									delete := words[1] + ":00"
									check := true
									for num, image := range imageInQues {
										if image.Time == delete {
											removedImage := removeImage(imageInQues, num)
											sortedImage := sortImage(removedImage)
											writeImageInQue(sortedImage)
											message := "ลบกำหนดการส่งรูปภาพที่เวลา " + words[1] + " สำเร็จ"
											textMessage := linebot.NewTextMessage(message)
											bot.ReplyMessage(event.ReplyToken, textMessage).Do()
										}
									}
									if check {
										textMessage := linebot.NewTextMessage("ไม่มีรูปภาพที่มีกำหนดการส่งในเวลานี้")
										bot.ReplyMessage(event.ReplyToken, textMessage).Do()
									}
								} else {
									textMessage := linebot.NewTextMessage(result)
									bot.ReplyMessage(event.ReplyToken, textMessage).Do()
								}
							case "TEST":
								result := checkTimeCorrection(words[1])
								if result == "" {
									imageInQues := readImagesFile()
									check := true
									for _, image := range imageInQues {
										if image.Time == words[1]+":00" {
											check = false
											imageMessage := linebot.NewImageMessage(image.ImageLink, image.ImageLink)
											bot.ReplyMessage(event.ReplyToken, imageMessage).Do()
										}
									}
									if check {
										textMessage := linebot.NewTextMessage("ไม่มีรูปภาพที่มีกำหนดการส่งในเวลานี้")
										bot.ReplyMessage(event.ReplyToken, textMessage).Do()
									}
								} else {
									textMessage := linebot.NewTextMessage(result)
									bot.ReplyMessage(event.ReplyToken, textMessage).Do()
								}
							case "GETALL":
								imageInQues := readImagesFile()
								times := "กำหนดการส่งรูปภาพทั้งหมด"
								for _, image := range imageInQues {
									times = times + "\n"
									times = times + image.Time
								}
								textMessage := linebot.NewTextMessage(times)
								bot.ReplyMessage(event.ReplyToken, textMessage).Do()
							default:
								// Unmarshal JSON
								flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(storemessage.InstructionMessage))
								if err != nil {
									log.Println(err)
								}
								// New Flex Message
								flexMessage := linebot.NewFlexMessage("คำสั่งผู้ดูแลระบบ", flexContainer)
								// Reply Message
								bot.ReplyMessage(event.ReplyToken, flexMessage).Do()
							}
						}
					}
				}
			}
		}

		return c.String(http.StatusOK, "ok")
	})

	go executeCronJob()
	e.Logger.Fatal(e.Start(":1323"))

}

func replyMessageLine(Message ReplyMessage) error {
	value, _ := json.Marshal(Message)

	url := "https://api.line.me/v2/bot/message/reply"

	var jsonStr = []byte(value)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+ChannelToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	return err
}

func getProfile(userID string) Profile {

	url := "https://api.line.me/v2/bot/profile/" + userID

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+ChannelToken)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var profile Profile
	if err := json.Unmarshal(body, &profile); err != nil {
		panic(err)
	}
	return profile
}

func getMessageContent(messageID string) []byte {

	url := "https://api-data.line.me/v2/bot/message/" + messageID + "/content"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+ChannelToken)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func readUsersFile() []Profile {
	// Open our jsonFile
	jsonFile, err := os.Open("users.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		panic(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var profiles []Profile

	json.Unmarshal(byteValue, &profiles)

	return profiles
}

func readImagesFile() []ImageInQue {
	// Open our jsonFile
	jsonFile, err := os.Open("image.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		panic(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var imageInQue []ImageInQue

	json.Unmarshal(byteValue, &imageInQue)

	return imageInQue
}

func writeUsersFile(users []Profile) error {
	file, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("users.json", file, 0644)
	if err != nil {
		return err
	}

	return nil
}

func writeImageInQue(image []ImageInQue) error {
	file, err := json.MarshalIndent(image, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("image.json", file, 0644)
	if err != nil {
		return err
	}

	return nil
}

func authentication(userID string) bool {
	p := getProfile(userID)

	users := readUsersFile()

	check := false

	for _, user := range users {
		if p.UserID == user.UserID {
			check = true
		}
	}

	return check
}

func removeImage(image []ImageInQue, i int) []ImageInQue {
	image[len(image)-1], image[i] = image[i], image[len(image)-1]
	return image[:len(image)-1]
}

func sortImage(image []ImageInQue) []ImageInQue {
	for num1, i := range image {
		for num2, j := range image {
			check := false

			hourI, _ := strconv.Atoi(i.Time[0:2])

			hourJ, _ := strconv.Atoi(j.Time[0:2])

			minuteI, _ := strconv.Atoi(i.Time[3:5])

			minuteJ, _ := strconv.Atoi(j.Time[3:5])

			if hourI < hourJ {
				check = true
			} else if minuteI < minuteJ && hourI == hourJ {
				check = true
			}

			if check {
				image[num1], image[num2] = image[num2], image[num1]
			}
		}
	}
	return image
}

func checkTimeCorrection(time string) string {

	if len(time) != 5 {
		return "รูปแบบของเวลาไม่ถูกต้อง กรุณาใส่เวลาในรูปแบบ HH:MM เช่น 08:00 หรือ 17:30 เป็นต้น"
	}

	hour, err := strconv.Atoi(time[0:2])
	if err != nil {
		return err.Error()
	}

	if hour > 24 {
		return "กรุณาเลือกเวลาตั้งแต่ 00:00 นาฬิกา ถึง 23:59 นาฬิกา"
	}

	minute, err := strconv.Atoi(time[3:5])
	if err != nil {
		return err.Error()
	}

	if minute > 59 {
		return "กรุณาเลือกเวลาตั้งแต่ 00:00 นาฬิกา ถึง 23:59 นาฬิกา"
	}

	return ""
}

func myTask() {

	img := readImagesFile()

	currentTime := time.Now()
	timeStampString := currentTime.Format("2006-01-02 15:04:05")
	layOut := "2006-01-02 15:04:05"
	timeStamp, err := time.Parse(layOut, timeStampString)
	if err != nil {
		fmt.Println(err)
	}
	hr, min, _ := timeStamp.Clock()

	var strHr string

	if hr < 10 {
		strHr = "0" + strconv.Itoa(hr)
	}else{
		strHr = strconv.Itoa(hr)
	}

	t := strHr + ":" + strconv.Itoa(min) + ":00"

	fmt.Println(t)

	for _, i := range img{
		if i.Time == t{
			client := &http.Client{}
			bot, err := linebot.New("0dcc21819b9f37988e430c810cf1c849", ChannelToken, linebot.WithHTTPClient(client))
			if err != nil {
				panic(err)
			}
			imageMessage := linebot.NewImageMessage(i.ImageLink, i.ImageLink)
			_ , err = bot.BroadcastMessage(imageMessage).Do()
			if err != nil{
				panic(err)
			}
		}
	}
}

func executeCronJob() {
	gocron.Every(60).Second().Do(myTask)
	<-gocron.Start()
}

/*if event.Message.Type == "image"{

	fmt.Println(event.Message)

	imgByte := getMessageContent(event.Message.ID)

	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		return err
	}
	out, _ := os.Create("./temp.jpeg")
	defer out.Close()

	err = jpeg.Encode(out, img, nil)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Success")
	}
}*/
