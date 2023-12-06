package auth_db

import (
	"fmt"
	openDb "match_card/config/db"
	auth_domain "match_card/pkg/auth/domain"
)

func Create(inputModel auth_domain.CreateUserRequest) (statusData bool, message string) {
	db := openDb.OpenConnection()
	defer db.Close()

	returnStatus := false
	msg := "no data"

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return returnStatus, err.Error()
	}

	queryPost := `INSERT INTO "expo"."tbl_score" 
				("username", "error", "time", "score", "difficulty") VALUES 
				($1, $2, $3, $4, $5);`
	_, err = tx.Exec(queryPost, inputModel.Username, inputModel.Error, inputModel.Time, inputModel.Score, inputModel.Difficulty)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return returnStatus, err.Error()
	}
	tx.Commit()

	returnStatus = true
	return returnStatus, msg
}

func GetAllScore() (returnDb []auth_domain.ScoreResponse, statusData bool, message string) {
	db := openDb.OpenConnection()

	var returnDatas []auth_domain.ScoreResponse

	returnStatus := false
	queryGet := `SELECT
	"id",
	"username",
	"error",
	"time",
	"score",
	"difficulty",
	"createdt"
  FROM
	"expo"."tbl_score"
	WHERE createdt <= '2023-12-06 17:00:00'
  ORDER BY "score" DESC, "time" ASC, error ASC, username ASC, createdt ASC;`

	row, err := db.Query(queryGet)
	if err != nil {
		fmt.Println(err)
		return returnDatas, returnStatus, err.Error()
	}

	defer db.Close()
	defer row.Close()
	msg := "No data available"

	for row.Next() {
		var returnData auth_domain.ScoreResponse
		row.Scan(
			&returnData.Id,
			&returnData.Username,
			&returnData.Error,
			&returnData.Time,
			&returnData.Score,
			&returnData.Difficulty,
			&returnData.Createdt)
		msg = "Success Get Data"
		returnDatas = append(returnDatas, returnData)
	}
	returnStatus = true
	return returnDatas, returnStatus, msg

}
