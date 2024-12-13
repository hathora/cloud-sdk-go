// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AppConfig struct {
	// Configure [player authentication](https://hathora.dev/docs/lobbies-and-matchmaking/auth-service) for your application. Use Hathora's built-in auth providers or use your own [custom authentication](https://hathora.dev/docs/lobbies-and-matchmaking/auth-service#custom-auth-provider).
	AuthConfiguration AuthConfiguration `json:"authConfiguration"`
	// Readable name for an application. Must be unique within an organization.
	AppName string `json:"appName"`
}

func (o *AppConfig) GetAuthConfiguration() AuthConfiguration {
	if o == nil {
		return AuthConfiguration{}
	}
	return o.AuthConfiguration
}

func (o *AppConfig) GetAppName() string {
	if o == nil {
		return ""
	}
	return o.AppName
}
