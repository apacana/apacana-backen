package helper

import (
	"fmt"
	"github.com/apacana/apacana-api/biz/dal/mysql"
	"math/rand"
	"time"
)

var pointTypeByName = map[string]mysql.PointType{
	"search_point": mysql.PointType_SEARCH,
	"agoda_hotel":  mysql.PointType_AGODA_HOTEL,
}

var nameByPointType = map[mysql.PointType]string{}

var randomColorList = make([]string, 0)

func init() {
	for mp, n := range pointTypeByName {
		nameByPointType[n] = mp
	}
	randomColorList = []string{
		"#3F51B5",
		"#303F9F",
		"#FF4081",
		"#33B5E5",
		"#AA66CC",
		"#99CC00",
		"#FFBB33",
		"#FF4444",
		"#0099CC",
		"#9933CC",
		"#669900",
		"#FF8800",
		"#CC0000",
	}
}

func GetPointTypeByName(name string) (mysql.PointType, error) {
	if value, ok := pointTypeByName[name]; !ok {
		return mysql.PointType_UNKNOW, fmt.Errorf("unknown point type name: %s", name)
	} else {
		return value, nil
	}
}

func GetNameByPointType(pointType mysql.PointType) (string, error) {
	if value, ok := nameByPointType[pointType]; !ok {
		return "", fmt.Errorf("unknown point type: %d", pointType)
	} else {
		return value, nil
	}
}

func CreateRandomColor() string {
	rand.Seed(time.Now().UnixNano())
	return randomColorList[rand.Intn(len(randomColorList))]
}
