package utils

import "fmt"

type ImgTable struct {
	single Items
	ten Items
	twelve Items
	sixty Items
	twohun Items
	fivhun Items
	overTwoThou Items
}
type Items struct {
	scale string
	variants []string
} 

var basePath = "https://pwaa-result-img.s3.ap-northeast-2.amazonaws.com"
var imgTable = ImgTable{
	single: Items{
		scale: "1",
		variants: []string{"caramel"},
	}, 
	ten: Items{
		scale: "2000",
		variants:[]string{"ice_cream"},
	},
	twelve: Items{
		scale: "20000",
		variants: []string{"candle"},
	},
	sixty: Items{
		scale: "60000",
		variants: [] string{"shoes"},
	},
	twohun: Items{
		scale: "200000",
		variants: []string{"game_machine"},
	},
	fivhun: Items{
		scale: "500000",
		variants: []string{"cell_phone"},
	},
	overTwoThou: Items{
		scale: "2000000",
		variants: []string{"luxury_watch"},
	},
}

func priceToPresentValue(price int) Items {
	fmt.Print(price)
	switch {
		case 0 >= price || price < 2000:
			return imgTable.single
		case 2000 >= price || price < 20999:
			return imgTable.ten
		case 21000 >= price || price < 60999:
			return imgTable.twelve
		case 61000 >= price || price < 200999:
			return imgTable.sixty
		case 201000 >= price || price < 500999:
			return imgTable.twohun
		case 501000 >= price || price < 2000999:
			return imgTable.fivhun
		default: 
			return imgTable.overTwoThou
	}
}
func getRandomIndex(length int) int {
	return 0
}

func ImgPathStr(price int) string {
	var presentValue = priceToPresentValue(price)
	var index = getRandomIndex(len(presentValue.variants))
	var item = presentValue.variants[index]
	return basePath + "/"+ presentValue.scale + "/" + item + ".png"
}