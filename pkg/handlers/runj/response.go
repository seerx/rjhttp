package runj

import "github.com/seerx/runjson/pkg/rj"

type RjResponse struct {
	Error string      `json:"error"`
	Data  rj.Response `json:"data"`
}
