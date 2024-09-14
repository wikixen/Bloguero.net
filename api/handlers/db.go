package handlers

import (
	conf "github.com/wikixen/blogapp/database/config"
)

// The only purpose of this file to create a DB var so that handlers can edit the DB 
var db = conf.CreateDB()