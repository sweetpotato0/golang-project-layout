package register

import (
    "fmt"
    "net"
    "strconv"
    "time"

    "layout/configs"
    "layout/pkg/log"

    "github.com/hashicorp/consul/api"
    "google.golang.org/grpc/resolver"
)

// Resolver .
type Resolver struct {
    addr    string
    cc      resolver.ClientConn
    timeout int
    name    string
    ticket  int
}

// NewResolver initialize an etcd client
func NewResolver(c *configs.Configuration, cc resolver.ClientConn, name, addr string) *Resolver {
    return &Resolver{
        addr:   addr,
        cc:     cc,
        ticket: c.Consul.Ticket,
        name:   name,
    }
}

// ResolveNow .
func (r Resolver) ResolveNow(rn resolver.ResolveNowOptions) {}

// Close closes the resolver.
func (r Resolver) Close() {}

func (r *Resolver) watcher() {

    config := api.DefaultConfig()
    config.Address = r.addr
    client, err := api.NewClient(config)
    if err != nil {
        msg := fmt.Sprintf("create consul client failed: %v", err)
        log.Errorf(msg)
        panic(fmt.Errorf(msg))
    }

    t := time.NewTicker(time.Duration(r.ticket) * time.Millisecond)
    for {
        select {
        case <-t.C:
        }
        services, _, err := client.Health().Service(r.name, "", true, &api.QueryOptions{})
        if err != nil {
            log.Errorf("retrieving instances from Consul failed: %v", err)
        }

        newAddrs := make([]resolver.Address, 0)
        for _, service := range services {
            addr := net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port))
            newAddrs = append(newAddrs, resolver.Address{
                Addr:       addr,
                ServerName: service.Service.Service,
            })
        }

        r.cc.UpdateState(resolver.State{Addresses: newAddrs})
    }
}
