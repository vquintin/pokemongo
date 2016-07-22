package auth

type Provider string
type Token string

type AuthInfo struct {
	Provider Provider
	Token    Token
}

type PokemonGoConnector interface {
	AuthInfo() (AuthInfo, error)
}
