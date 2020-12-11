package data

import (
	"database/sql"
	// "fmt"
)

var DB *sql.DB


func GenerateCouponCode() string {
	return "GenerateCouponCode"
}

func GetCouponDetails() string {
	rows, err := DB.Query("SELECT COUNT(*) FROM coupon")
	if err != nil { /* error handling */}
	items := make([]*SomeStruct, 0, 10)
	var ida, idb uint
	for rows.Next() {
		err = rows.Scan(&ida, &idb)
		if err != nil { /* error handling */}
		items = append(items, &SomeStruct{ida, idb})
	}

	return "bam"
}

// func UpdateCouponDetails(){
// 	return "nothing"
// }

// func ValidateCoupon(){
// 	return "nothing"
// }

// func ConsumeCoupon(){
// 	return "nothing"
// }