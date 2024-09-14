package responses

type LoginRes struct{
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}