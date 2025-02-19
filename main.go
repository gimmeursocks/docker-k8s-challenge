package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "github.com/go-redis/redis/v8"
    "github.com/gorilla/mux"
    "context"
    "strconv"
)

var ctx = context.Background()

// Initialize Redis client
var rdb = redis.NewClient(&redis.Options{
    Addr:     fmt.Sprintf(os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")),
    Password: "", // No password set
    DB:       0,  // Default DB
})

func db_checker() {
	err := rdb.Ping(ctx).Err()  
	if err != nil {
		fmt.Println("Redis is not connected:", err)
		os.Exit(1)
	} else {
		fmt.Println("Redis is successfully connected!")
	}
}

func hello(w http.ResponseWriter, r *http.Request) {

	// Logging in app stdout
	fmt.Println("Wow! A new visit")

    // Get the current visit count
    visits, err := rdb.Get(ctx, "visits").Result()
    if err != nil {
		// Check the Redis server
		db_checker()
        visits = "0"
    }

    // Increment visit count
    newVisits, err := strconv.Atoi(visits) // Convert string to integer
    if err != nil {
        fmt.Println("Error converting visits to integer:", err)
        newVisits = 0
    }
    newVisits++

    // Store the updated visit count
    rdb.Set(ctx, "visits", strconv.Itoa(newVisits), 0)

    // Respond with the visit count
    fmt.Fprintf(w, "Hello, World! You are visitor number %d.", newVisits)
}

func main() {
	// Check the Redis server
	db_checker()

	// Start the Go App server
    router := mux.NewRouter()
    router.HandleFunc("/", hello)

    fmt.Println("Server is running...")
    log.Fatal(http.ListenAndServe(":" + os.Getenv("APP_PORT"), router))
}
