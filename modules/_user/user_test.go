package user

import (
	"context"
	"fmt"
	"go-boiler-plate/infra/database"
	"go-boiler-plate/infra/environment"
	"go-boiler-plate/util"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/go-logr/stdr"
	"go.opentelemetry.io/otel"
)

func TestGetUser(t *testing.T) {
	fmt.Println("cache 2")

	logger := stdr.New(log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile))
	otel.SetLogger(logger)

	environment.InitializeEnvs()
	database.InitializeGorm()

	ctx := context.Background()

	user, err := GetUser(&ctx)
	util.AssertError(err, http.StatusInternalServerError, "user not found")

	fmt.Printf("%+v", user)
}
