package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/erkylima/parquet-logs/internal/logs/domains"
	"github.com/erkylima/parquet-logs/internal/shared/models"
	_ "github.com/marcboeker/go-duckdb"
)

type parquetLogConnection struct {
	Ctx        context.Context
	collection *sql.DB
}

func NewParquetLogConnection(dbName, fileName, auditFileName string) (*parquetLogConnection, error) {
	db, err := sql.Open(dbName, fileName)
	ctx := context.Background()
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	exec, err := db.ExecContext(ctx, "ATTACH '"+dbName+"'")
	if err != nil {
		return nil, err
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affected == 0 {
		parquet, err := db.ExecContext(ctx, `CREATE TABLE parquet-logs (
			id VARCHAR(255) PRIMARY KEY, 
			name VARCHAR(255) NOT NULL,
			dataObject TEXT,
			description VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			eventDate VARCHAR(255) NOT NULL,
			eventTicks VARCHAR(255) NOT NULL,
			idClient VARCHAR(255) NOT NULL,
			idProfile VARCHAR(255) NOT NULL,        
			identifier VARCHAR(255) NOT NULL,
			operation VARCHAR(255) NOT NULL,
			processName VARCHAR(255) NOT NULL,
			tableName VARCHAR(255) NOT NULL,
			userId VARCHAR(255) NOT NULL,
			username VARCHAR(255) NOT NULL
			)`)
		if err != nil {
			fmt.Printf(err.Error())
		} else {
			fmt.Println(parquet.LastInsertId())
		}
	}

	return &parquetLogConnection{
		Ctx:        ctx,
		collection: db,
	}, nil
}

func (mc *parquetLogConnection) Create(entity *domains.Log) (string, error) {
	insert := entity.InsertString()
	_, err := mc.collection.Exec("INSERT INTO parquet-logs VALUES (" + insert + ")")
	if err != nil {
		fmt.Println(err)
	}

	return "", nil
}

func (mc *parquetLogConnection) List(filters models.Filter) ([]domains.Log, error) {

	var entities []domains.Log
	querySelect := "SELECT * FROM parquet-logs" + filters.QueryToSqlWhere() + ";"
	query, err := mc.collection.Query(querySelect)
	if err != nil {
		return nil, err
	}
	for query.Next() {
		var entity domains.Log
		if err := query.Scan(&entity.ID, &entity.Name, &entity.DataObject, &entity.Description, &entity.Email, &entity.EventDate, &entity.EventTicks, &entity.IDClient, &entity.IDProfile, &entity.Identifier, &entity.Operation, &entity.ProcessName, &entity.TableName, &entity.UserID, &entity.Username); err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	// defer cursor.Close(mc.Ctx)
	// for cursor.Next(mc.Ctx) {
	// 	var entity T
	// 	if err := cursor.Decode(&entity); err != nil {
	// 		return nil, err
	// 	}
	// 	entities = append(entities, entity)
	// }
	return entities, nil
}
