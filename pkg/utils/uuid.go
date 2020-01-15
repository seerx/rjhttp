package utils

import (
	"strings"

	"github.com/google/uuid"
)

//UUID 生成 uuid
func UUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
	//u := uuid.Must(uuid.NewV4(), nil)
	//return strings.Replace(u.String(), "-", "", -1)
}
