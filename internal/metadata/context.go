package metadata

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ContextValues(ctx context.Context) (
	string,
	*sqlx.DB,
) {
	serviceName := ctx.Value("serviceName").(string)
	dbConnection := ctx.Value("dbConnection").(*sqlx.DB)

	return serviceName, dbConnection
}

func SetContextValues(
	ctx context.Context,
	serviceName string,
	dbConnection *sqlx.DB,
) context.Context {
	ctx = context.WithValue(ctx, "serviceName", serviceName)
	ctx = context.WithValue(ctx, "dbConnection", dbConnection)

	return ctx
}

func ContextServiceName(ctx context.Context) string {
	serviceName := ctx.Value("serviceName").(string)
	return serviceName
}

func ContextDbConnection(ctx context.Context) *sqlx.DB {
	dbConnection := ctx.Value("dbConnection").(*sqlx.DB)
	return dbConnection
}
