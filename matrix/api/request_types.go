package api

type RegisterRequest struct {
	Authentication           *RegisterAuthenticationData `json:"auth,omitempty"`
	BindEmail                bool                        `json:"bind_email,omitempty"`
	Username                 string                      `json:"username,omitempty"`
	Password                 string                      `json:"password,omitempty"`
	DeviceId                 string                      `json:"device_id,omitempty"`
	InitialDeviceDisplayName string                      `json:"initial_device_display_name,omitempty"`
}

type RegisterAuthenticationData struct {
	Type    string `json:"type"`
	Session string `json:"session"`
}

const LoginTypePassword = "m.login.password"
const LoginTypeToken = "m.login.token"
type LoginRequest struct {
	Type string `json:"type"`
	Username string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	// ... and other parameters we don't care about
}