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

	queryPut := `
	UPDATE "expo"."tbl_score"
	SET "error"=(error+$1)/2, "time"=("time"+$2)/2, "score"=score+$3, difficulty = NULL
	WHERE ("expo"."tbl_score"."username" = $4);`

	rows, err := tx.Exec(queryPut, inputModel.Error, inputModel.Time, inputModel.Score, inputModel.Username)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return returnStatus, err.Error()
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return returnStatus, err.Error()
	}

	if rowsAffected < 1 {
		queryPost := `INSERT INTO "expo"."tbl_score" 
				("username", "error", "time", "score", "difficulty") VALUES 
				($1, $2, $3, $4, $5);`
		_, err = tx.Exec(queryPost, inputModel.Username, inputModel.Error, inputModel.Time, inputModel.Score, inputModel.Difficulty)
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
			return returnStatus, err.Error()
		}
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
	WHERE createdt <= '2023-12-01 16:00:00'
  ORDER BY "score" DESC, "time" ASC, error ASC, createdt ASC, username ASC;`

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
