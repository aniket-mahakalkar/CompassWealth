package views

import "compass-wealth/enums"

type EditUserRole struct {
	ID   uint        `json:"id"`
	Role enums.Roles `json:"role"`
}
