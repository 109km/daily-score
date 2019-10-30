package models

import (
	"strings"
	"time"
)


// Get all ads
func GetAllAds() []*Ad {
	var AdList []*Ad
	qs := OrmInstance.QueryTable("ad")
	qs.All(&AdList)
	return AdList
}

func GetAdsByDate(date string) []*Ad{
	var AdList []*Ad
	qs := OrmInstance.QueryTable("ad")
	qs.Filter("submit_time__exact", date).All(&AdList)
	return AdList
}

func AddOneAd(q1 []string, q2 string, q3 string, name string,phone string) (aid int64, err error) {
	var ad Ad
	
	ad.Q1 = strings.Join(q1,",")
	ad.Q2 = q2
	ad.Q3 = q3
	ad.Phone = phone
	ad.Name = name

	date := time.Now().Format("2006-01-02 15:04:05")
	ad.SubmitTime = date[0:10]

	id, err := OrmInstance.Insert(&ad)
	if err == nil {
		return id, nil
	}
	return 0, err

}
