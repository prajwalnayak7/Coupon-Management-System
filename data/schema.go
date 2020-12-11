package data

import (
	"time"
)

type coupon struct {
	id 					int64 `json:"id"`
	code 				string `json:"code"`
	expiryDate 			time.Time `json:"expiry_date"`
	availabilityCount 	int `json:"availability_count"`
	productId 			string `json:"product_id"`
	promoType			string `json:"promo_type"`
	discountFixed		float64 `json:"discount_fixed"`
	discountVariable	float64 `json:"discount_variable"`
	valid				bool `json:"valid"`
	createdAt 			time.Time `json:"created_at"`
	updatedAt 			time.Time `json:"updated_at"`
}

type order struct {
	id 			int64 `json:"id"`
	couponId 	int64 `json:"coupon_id"`
	clientId 	string `json:"client_id"`
	status 		string `json:"status"`
	createdAt 	time.Time `json:"created_at"`
}
