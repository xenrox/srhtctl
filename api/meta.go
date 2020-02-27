package api

type metaUserStruct struct {
	userStruct
	UsePGPKey *int `json:"use_pgp_key"`
}
