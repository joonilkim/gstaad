package cognito

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cogp "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type TokenPayload struct {
	IdToken      string
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
	TokenType    string
}

type CognitoClient interface {
	Signup(email, pw string) error
	Login(*TokenPayload, string, string) error
	Unregister(token string) error
}

type client struct {
	provider   *cogp.CognitoIdentityProvider
	clientID   string
	userPoolID string
}

func NewCognitoClient(sess *session.Session) CognitoClient {
	c := &client{}
	c.provider = cogp.New(sess)
	mustGetenv("AWS_REGION")
	c.clientID = mustGetenv("COGNITO_CLIENT_ID")
	c.userPoolID = mustGetenv("COGNITO_USERPOOL_ID")
	return c
}

func (c *client) confirm(email string) error {
	p := &cogp.AdminConfirmSignUpInput{
		UserPoolId: aws.String(c.userPoolID),
		Username:   aws.String(email),
	}

	_, er := c.provider.AdminConfirmSignUp(p)
	return er
}

func (c *client) Signup(email, pw string) error {
	p := &cogp.SignUpInput{
		ClientId: aws.String(c.clientID),
		Username: aws.String(email),
		Password: aws.String(pw),
	}
	_, er := c.provider.SignUp(p)
	if er != nil {
		return er
	}
	return c.confirm(email)
}

func (c *client) Unregister(token string) error {
	p := &cogp.DeleteUserInput{
		AccessToken: aws.String(token),
	}
	_, er := c.provider.DeleteUser(p)
	return er
}

func (c *client) Login(tk *TokenPayload, email, pw string) error {
	p := &cogp.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(email),
			"PASSWORD": aws.String(pw),
		},
		ClientId: aws.String(c.clientID),
	}

	r, er := c.provider.InitiateAuth(p)
	if er != nil {
		return er
	}

	tk.IdToken = *r.AuthenticationResult.IdToken
	tk.AccessToken = *r.AuthenticationResult.AccessToken
	tk.RefreshToken = *r.AuthenticationResult.RefreshToken
	tk.ExpiresIn = *r.AuthenticationResult.ExpiresIn
	tk.TokenType = *r.AuthenticationResult.TokenType
	return nil
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", k))
	}
	return v
}
