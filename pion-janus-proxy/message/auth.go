package message

type UserToken struct {
	Id      string `json:"unique_order_code" mapstructure:"unique_order_code"`
	Token   string `json:"token" mapstructure:"token"`
	Refresh string `json:"refresh_token" mapstructure:"refresh_token"`
}

type AuthRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Type         string `json:"type"`
	Role         string `json:"role"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	DeviceId     string `json:"device_id"`
	DeviceName   string `json:"device_name"`
}
