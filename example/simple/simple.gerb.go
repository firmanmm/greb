package simple

import (
	"encoding/json"
	"net/http"

	greb "github.com/firmanmm/greb"
)

type Simple struct {
	ID      int     `json:"-" validate:"required"`
	Name    string  `json:"-" validate:"required"`
	Weight  float64 `json:"Weight"`
	IsAlive bool    `json:"-"`
}

func (x *Simple) BindRequest(req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(x); err != nil {
		return err
	}
	var err error
	x.ID, err = greb.BindInt(req, "ID", greb.BIND_TYPE_QUERY)
	if err != nil {
		return err
	}
	x.Name, err = greb.BindString(req, "Name", greb.BIND_TYPE_FORM)
	if err != nil {
		return err
	}
	x.IsAlive, err = greb.BindBool(req, "IsAlive", greb.BIND_TYPE_FORM)
	if err != nil {
		return err
	}
	if err := greb.Validate(x); err != nil {
		return err
	}
	return nil
}
