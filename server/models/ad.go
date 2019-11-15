package models

import (
	"strings"
	"time"
)

var MaxRowsLimit = 1000

// Get all ads
func GetAllAds(start int, limit int) (AdList []*Ad, total int64) {

	if limit > MaxRowsLimit {
		limit = MaxRowsLimit
	}
	total, _ = OrmInstance.QueryTable("ad").Count()
	qs := OrmInstance.QueryTable("ad").Limit(limit, start)
	qs.All(&AdList)
	return AdList, total
}

func GetAdsByDate(date string, start int, limit int) (AdList []*Ad, total int64) {

	if limit > MaxRowsLimit {
		limit = MaxRowsLimit
	}
	qs := OrmInstance.QueryTable("ad")
	total, _ = qs.Filter("submit_time__exact", date).Count()
	qs.Limit(limit, start).All(&AdList)
	return AdList, total
}

func AddOneAd(q1 []string, q2 string, q3 string, name string, phone string) (aid int64, err error) {
	var ad Ad

	ad.Q1 = strings.Join(q1, ",")
	ad.Q2 = q2
	ad.Q3 = q3
	ad.Phone = phone
	ad.Name = name

	date := time.Now().Format("2006-01-02 15:04:05")
	ad.SubmitTime = date[0:10]
	ad.LogTime = date

	created, id, err := OrmInstance.ReadOrCreate(&ad, "Phone")
	if err == nil {
		if created {
			return id, nil
		} else {
			return -1, nil
		}
	}
	return -2, err

}
