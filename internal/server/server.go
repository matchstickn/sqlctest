package server

import (
	"database/sql"
	"fmt"

	"github.com/matchstickn/sqlctest/assets/db"
)

type trickId struct {
	Id int32 `json:"id"`
}

func BodyToCreateTrick(trick db.Trick) (db.CreateTrickParams, error) {
	name, style, power := sql.NullString{}, sql.NullInt32{}, sql.NullBool{}
	if err := name.Scan(trick.Name.String); err != nil {
		return db.CreateTrickParams{}, err
	}

	if err := style.Scan(trick.Style.Int32); err != nil {
		return db.CreateTrickParams{}, err
	}

	if err := power.Scan(trick.Power.Bool); err != nil {
		return db.CreateTrickParams{}, err
	}

	trickParams := db.CreateTrickParams{
		Name:  name,
		Style: style,
		Power: power,
	}
	fmt.Println(trickParams)
	return trickParams, nil
}
