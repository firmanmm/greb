package simple

import (
	"encoding/json"
	"net/http"

	greb "github.com/firmanmm/greb"
)

type Simple struct {
	ID            int     `json:"-" validate:"required"`
	Name          string  `json:"-" validate:"required"`
	Weight        float64 `json:"weight"`
	IsAlive       bool    `json:"-"`
	Authorization string  `json:"-" validate:"required"`
	SessionID     string  `json:"-"`
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
	x.Authorization, err = greb.BindString(req, "x-authorization", greb.BIND_TYPE_HEADER)
	if err != nil {
		return err
	}
	x.SessionID, err = greb.BindString(req, "SessionID", greb.BIND_TYPE_COOKIE)
	if err != nil {
		return err
	}
	if err := greb.Validate(x); err != nil {
		return err
	}
	return nil
}
