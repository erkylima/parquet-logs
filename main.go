package main

import (
	"github.com/erkylima/parquet-logs/pkg/api"
	_ "github.com/marcboeker/go-duckdb"
)

func main() {
	api := api.NewApi()

	api.InitLogsRouter()
	err := api.InitGin(0)
	if err != nil {
		panic(err)
	}

}
