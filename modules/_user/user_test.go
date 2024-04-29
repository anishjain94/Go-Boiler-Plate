package user

import (
	"context"
	"fmt"
	"go-boiler-plate/infra/database"
	"go-boiler-plate/infra/environment"
	"go-boiler-plate/util"
	"net/http"
	"testing"
)

func TestGetUser(t *testing.T) {
	fmt.Println("cache 5")

	environment.InitializeEnvs()
	database.InitializeGorm()

	ctx := context.Background()

	user, err := GetUser(&ctx)
	util.AssertError(err, http.StatusInternalServerError, "user not found")
	fmt.Println(*user)

}

func Test(t *testing.T) {

	fmt.Println("ca")
	ch := make(chan int)

	// Receive operation on unbuffered channel
	go func() {
		value := <-ch
		fmt.Println("Received value:", value)
	}()

	// Send operation on the same unbuffered channel
	ch <- 10
	fmt.Println("Sent value:", 10)
}
