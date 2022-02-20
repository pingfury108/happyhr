package metadata

import (
	"happyhr/models"
)

type Result struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type QueryEmployee struct {
	ID           uint   `from:"id"`
	Name         string `from:"name"`
	SerialNumber uint   `from:"serial_number"`
}

type QueryResult struct {
	Msg  string            `json:"msg"`
	Code int               `json:"code"`
	Data []models.Employee `json:"data"`
}

type CreateResult struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type UpdateResult struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

type DeleteEmployee struct {
	ID uint `from:"id"`
}
