package data

import (
	"database/sql"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenerateCouponCode(r *http.Request) string {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	con, err := sql.Open("mysql", user+":"+password+"@tcp(localhost:5000)/"+database)
	checkError(err)
	defer con.Close()

	r.PostFormValue("")
	if err := r.ParseForm(); err != nil {
		// handle error
	}
	data := make(map[string]string)
	for key, values := range r.PostForm {
		data[key] = strings.Join(values, " ")
	}
	log.Println("POST params were:", data)
	rand.Seed(time.Now().UnixNano())
	var i int
	if length, present := data["code_length"]; present {
		i, _ = strconv.Atoi(length)
	} else {
		i = 10
	}
	coupon_code := randSeq(i)
	log.Println("Generated a Coupon Code:", coupon_code)

	rows, err := con.Query("INSERT INTO `coupon` (code, expiry_date, availability_count, product_id, promo_type, discount_fixed, discount_variable) VALUES (?, ?, ?, ?, ?, ?, ?)", coupon_code, data["expiry_date"], data["availability_count"], data["product_id"], data["promo_type"], data["discount_fixed"], data["discount_variable"])
	checkError(err)
	log.Println(rows)
	return coupon_code
}

func GetCouponDetails(r *http.Request) map[string]string {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	con, err := sql.Open("mysql", user+":"+password+"@tcp(localhost:5000)/"+database)
	checkError(err)
	defer con.Close()

	log.Println("GET params were:", r.URL.Query())
	code := r.URL.Query().Get("code")
	rows, err := con.Query("SELECT * FROM `coupon` WHERE code=?", code)
	cols, _ := rows.Columns()

	data := make(map[string]string)
	if rows.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		rows.Scan(columnPointers...)
		for i, colName := range cols {
			data[colName] = columns[i]
		}
	}
	return data
}

func UpdateCouponDetails(r *http.Request) string {
	return "Successfully Updated"
}

func ValidateCoupon(r *http.Request) bool {
	coupon_details := GetCouponDetails(r)
	if coupon_details["valid"] == "false" {
		return false
	}
	availability_count, _ := strconv.Atoi(coupon_details["availability_count"])
	if availability_count < 1 {
		return false
	}
	// expiry_date, _ := time.Parse(time.RFC822, coupon_details["expiry_date"])
	// if expiry_date < time.Now(){
	// 	return false
	// }
	return true
}

func ConsumeCoupon(r *http.Request) string {
	if !ValidateCoupon(r) {
		return "Coupon Invalid"
	}

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	con, err := sql.Open("mysql", user+":"+password+"@tcp(localhost:5000)/"+database)
	checkError(err)
	defer con.Close()

	r.PostFormValue("")
	if err := r.ParseForm(); err != nil {
		// handle error
	}
	data := make(map[string]string)
	for key, values := range r.PostForm {
		data[key] = strings.Join(values, " ")
	}
	log.Println("POST params were:", data)

	coupon_details := GetCouponDetails(r)
	availability_count, _ := strconv.Atoi(coupon_details["availability_count"])
	// updated_availability_count := strconv.Itoa(availability_count-1)
	coupon_id := coupon_details["id"]
	log.Println("Updating the availability count of", coupon_details["code"], "to", availability_count-1)
	rows1, err := con.Query("UPDATE `coupon` SET availability_count=? WHERE id=?", availability_count-1, coupon_id)
	checkError(err)
	log.Println(rows1)
	rows2, err := con.Query("INSERT INTO `order` (coupon_id, client_id, status) VALUES (?, ?, ?)", coupon_id, rand.Intn(100), "Success")
	checkError(err)
	log.Println(rows2)

	return "Successfully Consumed"
}
