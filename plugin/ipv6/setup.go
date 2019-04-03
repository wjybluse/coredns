package ipv6

import (
	"strconv"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"

	"github.com/mholt/caddy"
)

func init() {
	caddy.RegisterPlugin("ipv6", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	enable := ipv6Parse(c)
	P := IPv6{Disable: !enable}
	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		P.Next = next
		return P
	})
	return nil
}

func ipv6Parse(c *caddy.Controller) bool {
	//default ipv6 is in enable state
	for c.Next() {
		args := c.RemainingArgs()
		if len(args) == 0 {
			return true
		}
		flag, err := strconv.ParseBool(args[0])
		if err != nil {
			//if args is invalid, return ipv6 supported
			return true
		}
		return flag
	}
	return true
}
