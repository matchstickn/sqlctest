package server

// "reflect"
// "github.com/matchstickn/sqlctest/assets/db"

type trickId struct {
	Id int32 `json:"id"`
}

// func BodyToCreateTrick(trick db.Trick) (db.CreateTrickParams, error) {
// 	newTrick := db.CreateTrickParams{}

// 	if trick.Name.String != newTrick.Name.String {
// 		newTrick.Name.String = trick.Name.String
// 		newTrick.Name.Valid = true
// 	}

// 	if trick.Style.Int32 != newTrick.Style.Int32 {
// 		newTrick.Style.Int32 = trick.Style.Int32
// 		newTrick.Style.Valid = true
// 	}

// 	if trick.Power.Bool != newTrick.Power.Bool {
// 		newTrick.Power.Bool = trick.Power.Bool
// 		newTrick.Power.Valid = true
// 	}

// 	return newTrick, nil
// }
