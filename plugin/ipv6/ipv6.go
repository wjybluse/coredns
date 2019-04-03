// Package IPv6 implements a plugin that returns details about the resolving
// querying it.
package ipv6

import (
	"context"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/request"

	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/miekg/dns"
)

var log = clog.NewWithPlugin("ipv6")

// IPv6 is plugin to disable ipv6 address
type IPv6 struct {
	Disable bool
	Next    plugin.Handler
}

// ServeDNS implements the plugin.Handler interface.
func (ip IPv6) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}
	if state.QType() == 0x1c && ip.Disable {
		//if disable ipv6 parser, the next handler will be set as nil
		return plugin.NextOrFailure(ip.Name(), nil, ctx, w, r)
	}
	return plugin.NextOrFailure(ip.Name(), ip.Next, ctx, w, r)
}

// Name implements the Handler interface.
func (ip IPv6) Name() string { return "ipv6" }
