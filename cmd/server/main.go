package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

	"github.com/minhkhiemm/example-go/config/database/pg"
	"github.com/minhkhiemm/example-go/endpoints"
	serviceHttp "github.com/minhkhiemm/example-go/http"
	"github.com/minhkhiemm/example-go/service"
	accountService "github.com/minhkhiemm/example-go/service/account"
	detailService "github.com/minhkhiemm/example-go/service/detail"
	drinkService "github.com/minhkhiemm/example-go/service/drink"
	orderService "github.com/minhkhiemm/example-go/service/order"
)

func main() {
	// setup env on local
	if os.Getenv("ENV") == "local" {
		err := godotenv.Load()
		if err != nil {
			panic(fmt.Sprintf("failed to load .env by errors: %v", err))
		}
	}

	// setup addrr
	httpAddr := ":" + os.Getenv("PORT")

	// setup log
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// setup service
	var (
		pgDB, closeDB = pg.New(os.Getenv("PG_DATASOURCE"))
		s             = service.Service{
			OrderService: service.Compose(
				orderService.NewPGService(pgDB),
			).(orderService.Service),
			AccountService: service.Compose(
				accountService.NewPGService(pgDB),
			).(accountService.Service),
			DetailService: service.Compose(
				detailService.NewPGService(pgDB),
			).(detailService.Service),
			DrinkService: service.Compose(
				drinkService.NewPGService(pgDB),
			).(drinkService.Service),
		}
	)
	defer closeDB()

	var h http.Handler
	{
		h = serviceHttp.NewHTTPHandler(
			endpoints.MakeServerEndpoints(s),
			logger,
			os.Getenv("ENV") == "local",
		)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", httpAddr)
		errs <- http.ListenAndServe(httpAddr, h)
	}()

	logger.Log("exit", <-errs)
}
