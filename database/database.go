package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB sql.DB

func InitDB()error{
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("PORT"),os.Getenv("PGUSER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
   db,err:=  sql.Open("postgres", psqlInfo)
      if err != nil {
		return err
	  }
	err2 := db.Ping()
	if err2 != nil {
			return err2
		}
    DB =* db
	 return nil
}