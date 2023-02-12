package formRequest

type MailBoxFormRequest struct {
	Name       string `json:"name,omitempty" binding:"required"`
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	TlsEnabled bool   `json:"tls_enabled" `
	MaxSize    int8   `json:"max_size"`
}
