package usersapp

type UserCreateRequest struct {
	Pubkey string `json:"pubKey"                    validate:"required,nip05"`
}
