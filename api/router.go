package api

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func InitRouter() http.Handler {
	router := http.NewServeMux()
	router.Handle("/coupon/", coupon())
	router.Handle("/coupon/consume/", consume())
	fmt.Printf("%T", router)
	return router
}

func coupon() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/coupon/" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		switch r.Method {
		case "GET":
			w.WriteHeader(http.StatusOK)
			jsonData, _ := json.Marshal(GetCouponDetails(r))
			fmt.Fprintln(w, string(jsonData))
		case "POST":
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, GenerateCouponCode(r))
		case "PUT":
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, UpdateCouponDetails(r))
		default:
			fmt.Fprintln(w, "Sorry, only GET and POST methods are supported.")
		}
	})
}

func consume() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/coupon/consume/" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		switch r.Method {
		case "GET":
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, ValidateCoupon(r))
		case "POST":
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, ConsumeCoupon(r))
		default:
			fmt.Fprintln(w, "Sorry, only GET and POST methods are supported.")
		}
	})
}
