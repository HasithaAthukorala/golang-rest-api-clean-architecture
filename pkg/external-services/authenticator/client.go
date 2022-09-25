package authenticator

type AuthenticationClient interface {
	Authenticate(token string) bool
}

type authenticationClient struct {
	authenticatorName string
}

func New(authenticatorName string) AuthenticationClient {
	return &authenticationClient{authenticatorName: authenticatorName}
}

func (client *authenticationClient) Authenticate(token string) bool {
	// Execute an API call to the IDP here and verify the token.
	// For the demo purposes I'm using `abcde` as the token
	if token == "abcde" {
		return true
	}
	return false
}
