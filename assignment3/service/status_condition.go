package service

func KondisiWater(water int) string {
	var waterStatus string
	if water <= 5 {
		waterStatus = "Aman"
	} else if water >= 6 && water <= 8 {
		waterStatus = "Siaga"
	} else {
		waterStatus = "Bahaya"
	}
	return waterStatus
}

func KondisiWind(wind int) string {
	var windStatus string
	if wind <= 6 {
		windStatus = "Aman"
	} else if wind >= 7 && wind <= 15 {
		windStatus = "Siaga"
	} else {
		windStatus = "Bahaya"
	}
	return windStatus
}
