package main

import (
	"encoding/json"
	"errors"
)

type Error struct {
	ErrorCode int     `json:"error_code"`
	ErrorText string  `json:"error_text,omitempty"`
	CashBack  float64 `json:"cash_back,omitempty"`
	Change    float64 `json:"change,omitempty"`
}

const (
	E_NO_ERROR = iota
	E_PRODUCT_NOT_EXIST
	E_NOT_ENOUGH_PRODUCTS
	E_PRODUCT_ALREADY_EXISTS
	E_PARSE_PRODUCT_JSON
	E_PRODUCT_NAME_LENGTH
	E_PRODUCT_TITLE_LENGTH
)

var errorTexts = map[int]string{
	E_NO_ERROR:               "",
	E_PRODUCT_NOT_EXIST:      "Product not exists",
	E_NOT_ENOUGH_PRODUCTS:    "Not enough products",
	E_PRODUCT_ALREADY_EXISTS: "Product already exists",
	E_PARSE_PRODUCT_JSON:     "Malformed JSON data",
	E_PRODUCT_NAME_LENGTH:    "Product name can't be empty",
	E_PRODUCT_TITLE_LENGTH:   "Product title can't be empty",
}

func NewError(code int, cash float64) *Error {
	err := new(Error)
	if code == 0 {
		return err
	}
	err.ErrorCode = code
	err.ErrorText = errorTexts[code]
	err.CashBack = cash
	return err
}

func (e *Error) AsJSON() ([]byte, error) {
	data, err := json.Marshal(e)
	if err != nil {
		return nil, errors.New("Error marshalling to json: " + err.Error())
	}
	return data, nil
}
