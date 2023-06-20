package sdk

type BaseRequest struct {
	//Sessionkey string `json:"sessionkey"`
	IsJson  bool   `json:"retJson"  form:"retJson"` //是否返回Json
	Botoken string `json:"botoken"  form:"botoken"` //机器人Token
}
