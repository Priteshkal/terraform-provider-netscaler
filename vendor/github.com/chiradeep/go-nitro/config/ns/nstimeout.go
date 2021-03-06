package ns

type Nstimeout struct {
	Anyclient          int `json:"anyclient,omitempty"`
	Anyserver          int `json:"anyserver,omitempty"`
	Client             int `json:"client,omitempty"`
	Halfclose          int `json:"halfclose,omitempty"`
	Httpclient         int `json:"httpclient,omitempty"`
	Httpserver         int `json:"httpserver,omitempty"`
	Newconnidletimeout int `json:"newconnidletimeout,omitempty"`
	Nontcpzombie       int `json:"nontcpzombie,omitempty"`
	Reducedfintimeout  int `json:"reducedfintimeout,omitempty"`
	Server             int `json:"server,omitempty"`
	Tcpclient          int `json:"tcpclient,omitempty"`
	Tcpserver          int `json:"tcpserver,omitempty"`
	Zombie             int `json:"zombie,omitempty"`
}
