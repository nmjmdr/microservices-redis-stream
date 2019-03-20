package status

import (
	"fmt"
	"net/http"
	//redis "github.com/go-redis/redis"
)

// Handle - health check end point
// Can be enhanced to collect and report status of various parameters
func Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "redis:6379",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })

	// err := client.Set("key", "value", 0).Err()
	// if err != nil {
	// 	fmt.Fprintf(w, `Error`)
	// 	return
	// }

	// val, err := client.Get("key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	fmt.Fprintf(w, "OK")
}
