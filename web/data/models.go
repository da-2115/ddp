// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package data

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type MemberGender string

const (
	MemberGenderMale   MemberGender = "Male"
	MemberGenderFemale MemberGender = "Female"
)

func (e *MemberGender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = MemberGender(s)
	case string:
		*e = MemberGender(s)
	default:
		return fmt.Errorf("unsupported scan type for MemberGender: %T", src)
	}
	return nil
}

type NullMemberGender struct {
	MemberGender MemberGender `json:"member_gender"`
	Valid        bool         `json:"valid"` // Valid is true if MemberGender is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullMemberGender) Scan(value interface{}) error {
	if value == nil {
		ns.MemberGender, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.MemberGender.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullMemberGender) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.MemberGender), nil
}

type RoundClass string

const (
	RoundClassUnder14 RoundClass = "Under14"
	RoundClassUnder16 RoundClass = "Under16"
	RoundClassUnder18 RoundClass = "Under18"
	RoundClassUnder21 RoundClass = "Under21"
	RoundClassOpen    RoundClass = "Open"
	RoundClass50Plus  RoundClass = "50Plus"
	RoundClass60Plus  RoundClass = "60Plus"
	RoundClass70Plus  RoundClass = "70Plus"
)

func (e *RoundClass) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = RoundClass(s)
	case string:
		*e = RoundClass(s)
	default:
		return fmt.Errorf("unsupported scan type for RoundClass: %T", src)
	}
	return nil
}

type NullRoundClass struct {
	RoundClass RoundClass `json:"round_class"`
	Valid      bool       `json:"valid"` // Valid is true if RoundClass is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRoundClass) Scan(value interface{}) error {
	if value == nil {
		ns.RoundClass, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.RoundClass.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRoundClass) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.RoundClass), nil
}

type RoundDivision string

const (
	RoundDivisionRecurve         RoundDivision = "Recurve"
	RoundDivisionCompound        RoundDivision = "Compound"
	RoundDivisionRecurveBarebow  RoundDivision = "RecurveBarebow"
	RoundDivisionCompoundBarebow RoundDivision = "CompoundBarebow"
	RoundDivisionLongbow         RoundDivision = "Longbow"
)

func (e *RoundDivision) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = RoundDivision(s)
	case string:
		*e = RoundDivision(s)
	default:
		return fmt.Errorf("unsupported scan type for RoundDivision: %T", src)
	}
	return nil
}

type NullRoundDivision struct {
	RoundDivision RoundDivision `json:"round_division"`
	Valid         bool          `json:"valid"` // Valid is true if RoundDivision is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRoundDivision) Scan(value interface{}) error {
	if value == nil {
		ns.RoundDivision, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.RoundDivision.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRoundDivision) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.RoundDivision), nil
}

type RoundGender string

const (
	RoundGenderMale   RoundGender = "Male"
	RoundGenderFemale RoundGender = "Female"
)

func (e *RoundGender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = RoundGender(s)
	case string:
		*e = RoundGender(s)
	default:
		return fmt.Errorf("unsupported scan type for RoundGender: %T", src)
	}
	return nil
}

type NullRoundGender struct {
	RoundGender RoundGender `json:"round_gender"`
	Valid       bool        `json:"valid"` // Valid is true if RoundGender is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRoundGender) Scan(value interface{}) error {
	if value == nil {
		ns.RoundGender, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.RoundGender.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRoundGender) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.RoundGender), nil
}

type Championship struct {
	Championshipid   int32  `json:"championshipid"`
	Eventid          int32  `json:"eventid"`
	Championshipname string `json:"championshipname"`
}

type End struct {
	Endid              int32  `json:"endid"`
	Rangeid            int32  `json:"rangeid"`
	Archeryaustraliaid string `json:"archeryaustraliaid"`
	Finalscore         int32  `json:"finalscore"`
	Staged             bool   `json:"staged"`
}

type Event struct {
	Eventid   int32     `json:"eventid"`
	Eventname string    `json:"eventname"`
	Date      time.Time `json:"date"`
}

type Member struct {
	Archeryaustraliaid string       `json:"archeryaustraliaid"`
	Passwordhash       string       `json:"passwordhash"`
	Firstname          string       `json:"firstname"`
	Dateofbirth        time.Time    `json:"dateofbirth"`
	Gender             MemberGender `json:"gender"`
	Clubrecorder       bool         `json:"clubrecorder"`
}

type Practiceevent struct {
	Practiceid         int32  `json:"practiceid"`
	Eventid            int32  `json:"eventid"`
	Archeryaustraliaid string `json:"archeryaustraliaid"`
}

type Range struct {
	Rangeid    int32 `json:"rangeid"`
	Roundid    int32 `json:"roundid"`
	Distance   int32 `json:"distance"`
	Targetsize int32 `json:"targetsize"`
}

type Round struct {
	Roundid  int32         `json:"roundid"`
	Eventid  int32         `json:"eventid"`
	Class    RoundClass    `json:"class"`
	Division RoundDivision `json:"division"`
	Gender   RoundGender   `json:"gender"`
}

type Score struct {
	Scoreid     int32  `json:"scoreid"`
	Endid       int32  `json:"endid"`
	Arrownumber int32  `json:"arrownumber"`
	Score       string `json:"score"`
}
