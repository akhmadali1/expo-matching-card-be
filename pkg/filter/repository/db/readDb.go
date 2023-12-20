package filter_db

import (
	"fmt"
	openDb "match_card/config/db"
	filter_domain "match_card/pkg/filter/domain"
)

func Create(inputModel filter_domain.CreateUserRequest) (statusData bool, message string) {
	db := openDb.OpenConnection()
	defer db.Close()

	returnStatus := false
	msg := "no data"

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return returnStatus, err.Error()
	}

	queryPost := `INSERT INTO "expo"."tbl_filter" 
				("username", "time", "score") VALUES 
				($1, $2, $3);`
	_, err = tx.Exec(queryPost, inputModel.Username, inputModel.Time, inputModel.Score)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return returnStatus, err.Error()
	}
	tx.Commit()

	returnStatus = true
	return returnStatus, msg
}

func GetAllScore() (returnDb []filter_domain.ScoreResponse, statusData bool, message string) {
	db := openDb.OpenConnection()

	var returnDatas []filter_domain.ScoreResponse

	returnStatus := false
	queryGet := `SELECT
	"id",
	"username",
	"time",
	"score",
	"createdt"
  FROM
	"expo"."tbl_filter"
	WHERE createdt <= '2023-12-20 14:00:00'
  ORDER BY "score" DESC, "time" ASC, createdt ASC, username ASC;`

	row, err := db.Query(queryGet)
	if err != nil {
		fmt.Println(err)
		return returnDatas, returnStatus, err.Error()
	}

	defer db.Close()
	defer row.Close()
	msg := "No data available"

	for row.Next() {
		var returnData filter_domain.ScoreResponse
		row.Scan(
			&returnData.Id,
			&returnData.Username,
			&returnData.Time,
			&returnData.Score,
			&returnData.Createdt)
		msg = "Success Get Data"
		returnDatas = append(returnDatas, returnData)
	}
	returnStatus = true
	return returnDatas, returnStatus, msg

}
