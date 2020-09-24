package cmd

import (
	"log"
	"testing"
	//. "github.com/smartystreets/goconvey/convey"
	"github.com/uniplaces/carbon"
)

func TestCarbon(t *testing.T) {
	startTime,endTime := carbon.Now().StartOfMonth().Unix(),carbon.Now().EndOfMonth().Unix()
	var totalScore float64
	db.Table("c_bill b").Select("SUM(b.price)").Where("addtime BETWEEN ? AND ?", startTime, endTime).Scan(&totalScore)
	log.Printf("TOTAL %#v\n", totalScore)
}