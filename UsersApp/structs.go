package UsersApp


type UserCreateRequest struct {
	Pubkey string `json:"pubKey"                    validate:"required,nip05"`
}

