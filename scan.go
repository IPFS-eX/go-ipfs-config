package config

type Scan struct {
	NetworkId       int      // scan network id
	Bootstrap       []string // bootstrap scan peer host
	ListenAddr      string   // local listen addr of scan network
	Port            int      // listen port of scan network
	IgnoreIPv6      bool     // ignore IPv6 address
	IgnoreLocalHost bool     // ignore localhost, such as 127.xxx 192.xxx
	MaxNumWant      int      // maximum number of response peers for requesting file
}
