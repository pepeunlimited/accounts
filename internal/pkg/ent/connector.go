package ent

import (
	"github.com/facebookincubator/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pepeunlimited/microservice-kit/misc"
	"github.com/pepeunlimited/microservice-kit/sqlz"
	"log"
	"strconv"
)


func NewEntClient() *Client {
	username 	:= misc.GetEnv(sqlz.MysqlUser, "root")
	password 	:= misc.GetEnv(sqlz.MysqlRootPassword, "root")
	host 		:= misc.GetEnv(sqlz.MysqlHost, "localhost")
	port, _ 	:= strconv.Atoi(misc.GetEnv(sqlz.MysqlPort, "3306"))
	database 	:= misc.GetEnv(sqlz.MysqlDatabase, "accounts")  // <- change this
	uri 		:= sqlz.MySQLURI(username, password, host, port, database)
	client, err := Open(dialect.MySQL, uri)
	if err != nil {
		log.Panic("failed to open MySQL connection err: " + err.Error())
	}
	return client
}