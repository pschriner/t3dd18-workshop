package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	lorem "github.com/drhodes/golorem"
	"github.com/pschriner/t3dd18-workshop/model"
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func main() {
	var st sonyflake.Settings
	st.StartTime = time.Now()
	sf = sonyflake.NewSonyflake(st)
	for index := 0; index < 10; index++ {
		i := model.Image{UID: generateuid(), Filename: "test.png"}
		crdate := time.Now()
		startdate := time.Now()
		startdate = startdate.AddDate(0, 0, rand.Intn(5)-5)
		enddate := startdate
		enddate = enddate.AddDate(0, 0, rand.Intn(5)+7)
		a := model.Article{UID: generateuid(), Title: lorem.Word(1, 10), Text: lorem.Paragraph(1, 10), Image: i, Tstamp: &crdate, Startdate: &startdate, Enddate: &enddate}
		b, _ := json.Marshal(a)
		fmt.Println(string(b) + "\n")
	}

}

func generateuid() int64 {
	uidcounter, _ := sf.NextID()
	return int64(uidcounter)
}
