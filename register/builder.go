package register

// // Builder creates a resolver that will be used to watch name resolution updates.
// type Builder interface {
//     // Build creates a new resolver for the given target.
//     //
//     // gRPC dial calls Build synchronously, and fails if the returned error is
//     // not nil.
//     Build(target Target, cc ClientConn, opts BuildOptions) (Resolver, error)
//     // Scheme returns the scheme supported by this resolver.
//     // Scheme is defined at https://github.com/grpc/grpc/blob/master/doc/naming.md.
//     Scheme() string
// }
import (
    "layout/configs"

    "google.golang.org/grpc/resolver"
)

// Builder .
type Builder struct {
    c *configs.Configuration
}

// NewBuilder .
func NewBuilder(c *configs.Configuration) *Builder {
    return &Builder{
        c: c,
    }
}

// Build .
func (b *Builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

    r := NewResolver(b.c, cc, target.Endpoint, b.c.Consul.Addr)
    go r.watcher()

    return r, nil
}

// Scheme .
func (b *Builder) Scheme() string {
    return "consul"
}
