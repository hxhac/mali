package request

type SysAuthorityBtnReq struct {
	AuthorityId string `json:"authorityId"`
	Selected    []uint `json:"selected"`
	MenuID      uint   `json:"menuID"`
}
