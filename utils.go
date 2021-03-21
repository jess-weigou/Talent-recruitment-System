package main

type Data struct {
    Token string `json:"token"`
}
type SuccessReturn struct {
    Status bool `json:"status"`
    Msg   string      `json:"msg"`
    TokenData  Data   `json:"data"`
}
type ErrorReturn struct {
    Status bool    `json:"status"`
    Msg   string      `json:"msg"`
    TokenData  Data   `json:"data"`
}

func MakeSuccessReturn(data string)(int ,interface{})  {
    return 200,SuccessReturn{
        Status: true,
        Msg: "success",
        TokenData: Data{Token: data},
    }
}
func MakeErrorReturn(msg string)(int ,interface{})  {
    return 200,ErrorReturn{
        Status: false,
        Msg: msg,
        TokenData: Data{Token: ""},
    }
}