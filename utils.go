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

func MakeSuccessReturn(data interface{})(int ,interface{})  {
    return 200,SuccessReturn{
        Msg: "success",
        Data: data,
        Error: 0,
    }
}
func MakeErrorReturn(status int,code int ,msg string)(int ,interface{})  {
    return status,ErrorReturn{
        Msg: msg,
        Error: code,
    }
}