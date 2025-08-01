// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CreateAppConfig struct {
	LoadBalancer *LoadBalancerConfig `json:"loadBalancer,omitempty"`
	// Configure [player authentication](https://hathora.dev/docs/backend-integrations/lobbies-and-matchmaking/auth-service) for your application. Use Hathora's built-in auth providers or use your own [custom authentication](https://hathora.dev/docs/lobbies-and-matchmaking/auth-service#custom-auth-provider).
	AuthConfiguration AuthConfiguration `json:"authConfiguration"`
	// Readable name for an application. Must be unique within an organization.
	AppName string `json:"appName"`
}

func (o *CreateAppConfig) GetLoadBalancer() *LoadBalancerConfig {
	if o == nil {
		return nil
	}
	return o.LoadBalancer
}

func (o *CreateAppConfig) GetAuthConfiguration() AuthConfiguration {
	if o == nil {
		return AuthConfiguration{}
	}
	return o.AuthConfiguration
}

func (o *CreateAppConfig) GetAppName() string {
	if o == nil {
		return ""
	}
	return o.AppName
}
