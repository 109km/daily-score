package models

import (
	"strings"
)

var (
	AdList []*Ad
)

// Get all ads
func GetAllAds() []*Ad {
	qs := OrmInstance.QueryTable("ad")
	qs.All(&AdList)
	return AdList
}

func AddOneAd(q1 []string, q2 string, q3 string, name string,phone string) (aid int64, err error) {
	var ad Ad
	
	ad.Q1 = strings.Join(q1,",")
	ad.Q2 = q2
	ad.Q3 = q3
	ad.Phone = phone
	ad.Name = name

	id, err := OrmInstance.Insert(&ad)
	if err == nil {
		return id, nil
	}
	return 0, err

}
