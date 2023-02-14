package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

type SysAuthorityResponse struct {
	Authority system.SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	OldAuthorityId string              `json:"oldAuthorityId"`
	Authority      system.SysAuthority `json:"authority"`
}
