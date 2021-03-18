package storemessage

import (

)

//ExampleMessage is a message that use for demo
var ExampleMessage = `{
	"type": "bubble",
	"header": {
	  "type": "box",
	  "layout": "horizontal",
	  "contents": [
		{
		  "type": "text",
		  "text": "ปลอดภัยไว้ก่อน",
		  "weight": "bold",
		  "size": "lg",
		  "color": "#0A0A0AFF",
		  "align": "start",
		  "contents": []
		}
	  ]
	},
	"hero": {
	  "type": "image",
	  "url": "https://www.friendlygl.com/wp-content/uploads/2019/02/shutterstock_1.jpg",
	  "size": "full",
	  "aspectRatio": "20:13",
	  "aspectMode": "cover",
	  "action": {
		"type": "uri",
		"label": "Action",
		"uri": "https://linecorp.com/"
	  }
	},
	"body": {
	  "type": "box",
	  "layout": "horizontal",
	  "spacing": "md",
	  "contents": [
		{
		  "type": "box",
		  "layout": "vertical",
		  "flex": 2,
		  "contents": [
			{
			  "type": "text",
			  "text": "1. ตรวจสอบยางรถยนต์",
			  "weight": "bold",
			  "size": "md",
			  "flex": 1,
			  "gravity": "top",
			  "contents": []
			},
			{
			  "type": "separator"
			},
			{
			  "type": "text",
			  "text": "2. เติมน้ำมัน",
			  "weight": "bold",
			  "size": "md",
			  "flex": 2,
			  "gravity": "center",
			  "contents": []
			},
			{
			  "type": "separator"
			},
			{
			  "type": "text",
			  "text": "3. เช็คแบรก",
			  "weight": "bold",
			  "size": "md",
			  "flex": 2,
			  "gravity": "center",
			  "contents": []
			},
			{
			  "type": "separator"
			},
			{
			  "type": "text",
			  "text": "4. เมาไม่ขับ",
			  "weight": "bold",
			  "size": "md",
			  "color": "#030303FF",
			  "flex": 1,
			  "gravity": "bottom",
			  "contents": []
			}
		  ]
		}
	  ]
	},
	"footer": {
	  "type": "box",
	  "layout": "horizontal",
	  "contents": [
		{
		  "type": "button",
		  "action": {
			"type": "message",
			"label": "ตรวจสอบแล้ว",
			"text": "ตรวจสอบแล้ว"
		  },
		  "flex": 10,
		  "color": "#76DA23FF",
		  "style": "primary"
		}
	  ]
	}
  }`

//ExampleMessage2 is a message that use for demo
var ExampleMessage2 = `{
	"type": "bubble",
	"header": {
	  "type": "box",
	  "layout": "vertical",
	  "flex": 0,
	  "contents": [
		{
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "ปลอดภัยไว้ก่อน",
			  "weight": "bold",
			  "size": "xl",
			  "color": "#090707FF",
			  "contents": []
			}
		  ]
		}
	  ]
	},
	"hero": {
	  "type": "image",
	  "url": "https://www.friendlygl.com/wp-content/uploads/2019/02/shutterstock_1.jpg",
	  "size": "full",
	  "aspectRatio": "20:13",
	  "aspectMode": "cover",
	  "action": {
		"type": "uri",
		"label": "Action",
		"uri": "https://linecorp.com/"
	  }
	},
	"body": {
	  "type": "box",
	  "layout": "vertical",
	  "spacing": "md",
	  "contents": [
		{
		  "type": "text",
		  "text": "ขั้นตอนปฏิบัติก่อนออกเดินทาง",
		  "weight": "bold",
		  "size": "lg",
		  "gravity": "center",
		  "wrap": true,
		  "contents": []
		},
		{
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "sm",
		  "margin": "lg",
		  "contents": [
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "1.",
				  "weight": "bold",
				  "size": "md",
				  "color": "#090606FF",
				  "flex": 1,
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "ตรวจลมยางรถยนต์",
				  "weight": "bold",
				  "size": "md",
				  "color": "#050505FF",
				  "flex": 4,
				  "wrap": true,
				  "contents": []
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "2.",
				  "weight": "bold",
				  "size": "md",
				  "color": "#010101FF",
				  "flex": 1,
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "เติมน้ำมันให้เต็มถัง",
				  "weight": "bold",
				  "size": "md",
				  "color": "#090505FF",
				  "flex": 4,
				  "wrap": true,
				  "contents": []
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "3.",
				  "weight": "bold",
				  "size": "md",
				  "color": "#070606FF",
				  "flex": 1,
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "เช็คผ้าเบรก",
				  "weight": "bold",
				  "size": "md",
				  "color": "#000000FF",
				  "flex": 4,
				  "wrap": true,
				  "contents": []
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "contents": [
				{
				  "type": "text",
				  "text": "4.",
				  "weight": "bold",
				  "size": "md",
				  "color": "#0A0606FF",
				  "flex": 1,
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "เมาไม่ขับ",
				  "weight": "bold",
				  "size": "md",
				  "color": "#010101FF",
				  "flex": 4,
				  "contents": []
				}
			  ]
			}
		  ]
		}
	  ]
	}
  }`

//ExampleMessage3 is a message use for demo
var ExampleMessage3 = `{
	"type": "bubble",
	"direction": "ltr",
	"header": {
	  "type": "box",
	  "layout": "vertical",
	  "contents": [
		{
		  "type": "text",
		  "text": "ปลอดภัยไว้ก่อน",
		  "weight": "bold",
		  "size": "xl",
		  "color": "#000000FF",
		  "align": "start",
		  "contents": []
		}
	  ]
	},
	"hero": {
	  "type": "image",
	  "url": "https://www.friendlygl.com/wp-content/uploads/2019/02/shutterstock_1.jpg",
	  "size": "full",
	  "aspectRatio": "1.51:1",
	  "aspectMode": "fit"
	},
	"body": {
	  "type": "box",
	  "layout": "vertical",
	  "contents": [
		{
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "กรุณาทำแบบสำรวจต่อไปนี้",
			  "size": "lg",
			  "flex": 1,
			  "contents": []
			},
			{
			  "type": "text",
			  "text": "ก่อนออกเดินทาง",
			  "size": "lg",
			  "contents": []
			}
		  ]
		}
	  ]
	},
	"footer": {
	  "type": "box",
	  "layout": "horizontal",
	  "contents": [
		{
		  "type": "button",
		  "action": {
			"type": "uri",
			"label": "ทำแบบสำรวจ",
			"uri": "https://linecorp.com"
		  }
		}
	  ]
	}
  }`

//InstructionMessage is a message use for demo
var InstructionMessage = `{
	"type": "bubble",
	"body": {
	  "type": "box",
	  "layout": "vertical",
	  "contents": [
		{
		  "type": "text",
		  "text": "คำสั่งผู้ดูแลระบบ",
		  "weight": "bold",
		  "size": "xl",
		  "contents": []
		},
		{
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "sm",
		  "margin": "lg",
		  "contents": [
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "ADD",
				  "weight": "bold",
				  "size": "md",
				  "color": "#070505FF",
				  "flex": 1,
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "ใช้สำหรับเพิ่มกำหนดการส่งรูปภาพลงในระบบ",
				  "size": "md",
				  "color": "#000000FF",
				  "flex": 3,
				  "wrap": true,
				  "contents": []
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "ตัวอย่าง",
				  "weight": "bold",
				  "size": "md",
				  "color": "#000000FF",
				  "flex": 1,
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "ADD 18:00 https://line.me/static/115d5539e2d10b8da66d31ce22e6bccd/84249/favicon.png",
				  "flex": 3,
				  "wrap": true,
				  "contents": []
				}
			  ]
			},
			{
			  "type": "separator",
			  "color": "#000000FF"
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "DELETE",
				  "weight": "bold",
				  "size": "md",
				  "color": "#000000FF",
				  "flex": 1,
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "ใช้สำหรับลบกำหนดการส่ง\nในเวลานั้น",
				  "size": "md",
				  "color": "#000000FF",
				  "flex": 3,
				  "wrap": true,
				  "contents": []
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "ตัวอย่าง",
				  "weight": "bold",
				  "size": "md",
				  "color": "#000000FF",
				  "flex": 1,
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "DELETE 18:00",
				  "size": "md",
				  "color": "#000000FF",
				  "flex": 3,
				  "contents": []
				}
			  ]
			},
			{
			  "type": "separator",
			  "color": "#000000FF"
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "TEST",
				  "weight": "bold",
				  "size": "md",
				  "color": "#000000FF",
				  "flex": 1,
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "ใช้สำหรับดูรุปภาพที่ระบบจะส่งในเวลานั้น",
				  "size": "md",
				  "color": "#000000FF",
				  "flex": 3,
				  "wrap": true,
				  "contents": []
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "ตัวอย่าง",
				  "weight": "bold",
				  "size": "md",
				  "color": "#000000FF",
				  "contents": []
				},
				{
				  "type": "text",
				  "text": "TEST 18:00",
				  "color": "#000000FF",
				  "flex": 3,
				  "contents": []
				}
			  ]
			}
		  ]
		},
		{
		  "type": "separator",
		  "color": "#000000FF"
		},
		{
		  "type": "box",
		  "layout": "baseline",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "text",
			  "text": "GETALL",
			  "weight": "bold",
			  "size": "md",
			  "color": "#000000FF",
			  "contents": []
			},
			{
			  "type": "text",
			  "text": "ใช้เพื่อดูกำหนดการส่งทั้งหมด",
			  "flex": 3,
			  "wrap": true,
			  "contents": []
			}
		  ]
		},
		{
		  "type": "box",
		  "layout": "baseline",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "text",
			  "text": "ตัวอย่าง",
			  "weight": "bold",
			  "size": "md",
			  "color": "#000000FF",
			  "contents": []
			},
			{
			  "type": "text",
			  "text": "GETALL",
			  "flex": 3,
			  "contents": []
			}
		  ]
		},
		{
		  "type": "separator"
		},
		{
		  "type": "box",
		  "layout": "baseline",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "text",
			  "text": "หมายเหตุ",
			  "weight": "bold",
			  "size": "md",
			  "color": "#000000FF",
			  "flex": 1,
			  "contents": []
			},
			{
			  "type": "text",
			  "text": "เวลาให้ส่งในรูปแบบ HH:MM ตั้งแต่ 00:00 ถึง 23:59",
			  "flex": 2,
			  "wrap": true,
			  "contents": []
			}
		  ]
		}
	  ]
	}
  }`