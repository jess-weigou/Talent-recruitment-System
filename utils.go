package main

type SuccessReturn struct {
    Msg   string      `json:"msg"`
    Data  interface{} `json:"data"`
    Error int         `json:"error"`
}
type ErrorReturn struct {
    Msg   string      `json:"msg"`
    Error int         `json:"error"`
}
