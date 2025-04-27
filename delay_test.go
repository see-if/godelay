package godelay

import (
	"log"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	log.Println("start ...")
	//获取实例
	delay := GetEntry()
	//加入延时队列
	delay.Schedule(func() {
		log.Println("id_001  exec ...")
	}, 3*time.Second, "id_001")

	if time.Now().Year() == 2014 {
		//取消队列
		delay.Cancel("id_001")
	}
	//保持运行
	delay.Run()
}
