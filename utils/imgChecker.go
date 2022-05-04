package utils

type ImgTable struct {
	single Items
	ten Items
	twelve Items
	sixty Items
	twohun Items
	fivhun Items
	overTwoThou Items
}
type Items = []string

var basePath = "img source base string"
var imgTable = ImgTable{
	single: Items{"caramel"}, 
	ten: Items{"ice_cream"},
	twelve: Items{"candle"},
	sixty: Items{"shoeses"},
	twohun: Items{"game_machine"},
	fivhun: Items{"cell_phone"},
	overTwoThou: Items{"luxury_watch"},
}

func priceToPresentValue(price int) Items {
	switch {
		case 0 <= price || price < 2000:
			return imgTable.single
		case 2000 <= price || price < 20999:
			return imgTable.ten
		case 21000 <= price || price < 60999:
			return imgTable.twelve
		case 61000 <= price || price < 200999:
			return imgTable.sixty
		case 201000 <= price || price < 500999:
			return imgTable.twohun
		case 501000 <= price || price < 2000999:
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
	var index = getRandomIndex(len(presentValue))
	var item = presentValue[index]
	return basePath + item + ".png"
}