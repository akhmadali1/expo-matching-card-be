package logger_controller

import (
	"database/sql"
	"fmt"
	"log"
)

func InsertLogger(db *sql.DB, moduleId int, userId int, userName string, activity string, jsonBefore string, jsonAfter string) error {

	insertLog := `INSERT INTO "tbl_mlog" (moduleid, userid, username, activity, jsonbefore, jsonafter) VALUES ('%d', '%d', '%s' ,'%s', '%s' ,'%s')`
	responsequeryLog := fmt.Sprintf(insertLog, moduleId, userId, userName, activity, jsonBefore, jsonAfter)

	_, err := db.Exec(responsequeryLog)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
