package server

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2/internal/schema"
	"github.com/matchstickn/sqlctest/assets/db"
)

type trickId struct {
	Id int32 `json:"id"`
}

func SchemaRegisterConverterNulls(d *schema.Decoder) {
	d.RegisterConverter(sql.NullInt32{}, convertNullInt32)
	d.RegisterConverter(sql.NullBool{}, convertNullBool)
	d.RegisterConverter(sql.NullInt32{}, convertNullString)
}

func convertNullInt32(value interface{}) reflect.Value {
	var nv sql.NullInt32
	if err := nv.Scan(value); err != nil {
		return reflect.Value{}
	}
	return reflect.ValueOf(nv)
}

func convertNullBool(value interface{}) reflect.Value {
	var nv sql.NullBool
	if err := nv.Scan(value); err != nil {
		return reflect.Value{}
	}
	return reflect.ValueOf(nv)
}

func convertNullString(value interface{}) reflect.Value {
	var nv sql.NullString
	if err := nv.Scan(value); err != nil {
		return reflect.Value{}
	}
	return reflect.ValueOf(nv)
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
