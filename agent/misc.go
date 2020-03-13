package main

type DeviceAttributes struct {
	ID         string `json:"id"`
	PrettyName string `json:"pretty_name"`
}

type AuthRequest struct {
	Identity   *DeviceIdentity   `json:"identity"`
	Attributes *DeviceAttributes `json:"attributes"`
	PublicKey  string            `json:"public_key"`
	TenantID   string            `json:"tenant_id"`
	Version    string            `json:"version"`
	Sessions   []string          `json:"sessions,omitempty"`
}

type AuthResponse struct {
	UID       string `json:"uid"`
	Token     string `json:"token"`
	Name      string `json:"name"`
	Namespace string `json: "namespace"`
}
