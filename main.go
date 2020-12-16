package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	data "github.com/prajwalnayak7/Coupon-Management-System/data"
)

type key int

const (
	requestIDKey key = 0
)

var (
	listenAddr string
	healthy    int32
)

func connectToDatabase(user, password, database string){
	con, err := sql.Open("mysql", user+":"+password+"@tcp(localhost:5000)/"+database)
	if err != nil {
        log.Fatal(err)
    }
	defer con.Close()
}

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	flag.StringVar(&listenAddr, "listen-addr", ":5555", "server listen address")
	flag.Parse()

	logger := log.New(os.Stdout, "HTTP: ", log.LstdFlags)
	logger.Println("🔥 Server is starting...")

	user:=os.Getenv("MYSQL_USER")
	password:=os.Getenv("MYSQL_PASSWORD")
	database:=os.Getenv("MYSQL_DATABASE")
	connectToDatabase(user, password, database)
	logger.Println("Database", database,"connected successfully as user", user)

	router := http.NewServeMux()
	router.Handle("/coupon/", coupon())
	router.Handle("/coupon/consume/", consume())
	router.Handle("/ping", ping())

	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	server := &http.Server{
		Addr:         listenAddr,
		Handler:      tracing(nextRequestID)(logging(logger)(router)),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		logger.Println("Server is shutting down...")
		atomic.StoreInt32(&healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	logger.Println("Server is ready to handle requests at", listenAddr)
	atomic.StoreInt32(&healthy, 1)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

	<-done
	logger.Println("Server stopped")
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
				jsonData, _ := json.Marshal(data.GetCouponDetails(r))
				fmt.Fprintln(w, string(jsonData))
			case "POST":
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, data.GenerateCouponCode(r))
			case "PUT":
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, data.UpdateCouponDetails(r))
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
				fmt.Fprintln(w, data.ValidateCoupon(r))
			case "POST":
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, data.ConsumeCoupon(r))
			default:
				fmt.Fprintln(w, "Sorry, only GET and POST methods are supported.")
		}
	})
}

func ping() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(&healthy) == 1 {
			fmt.Fprintln(w, "Pong :)")
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintln(w, "Server Down :(")
	})
}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}
				logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
