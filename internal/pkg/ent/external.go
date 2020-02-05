package ent

import (
	"database/sql"
	entsql "github.com/facebookincubator/ent/dialect/sql"
)

func (c *Client) DB() *sql.DB {
	return c.driver.(*entsql.Driver).DB()
}