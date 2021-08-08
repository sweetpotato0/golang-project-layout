package register

import (
    "fmt"
    "layout/configs"
    "net"
    "time"

    "github.com/hashicorp/consul/api"
)

// Registry .
type Registry struct {
    Addr                           string
    Name                           string   // 服务名称
    Tag                            []string // // 服务端口
    Port                           int
    DeregisterCriticalServiceAfter time.Duration
    Interval                       time.Duration
}

// NewRegistry create a registry.
func NewRegistry(c *configs.Configuration) *Registry {
    return &Registry{
        Addr:                           c.Consul.Addr,
        Tag:                            []string{},
        DeregisterCriticalServiceAfter: time.Duration(1) * time.Minute,
        Interval:                       time.Duration(10) * time.Second,
    }
}

// RegistryOption is to add option to the registry.
type RegistryOption func(*Registry)

// Port set the port.
func Port(port int) RegistryOption {
    return func(r *Registry) {
        r.Port = port
    }
}

// Name set the port.
func Name(name string) RegistryOption {
    return func(r *Registry) {
        r.Name = name
    }
}

// Tag set the port.
func Tag(tag string) RegistryOption {
    return func(r *Registry) {
        r.Tag = append(r.Tag, tag)
    }
}

// Register .
func (r *Registry) Register() error {

    config := api.DefaultConfig()
    config.Address = r.Addr
    client, err := api.NewClient(config)
    if err != nil {
        return err
    }
    agent := client.Agent()

    IP := localIP()
    reg := &api.AgentServiceRegistration{
        ID:      fmt.Sprintf("%v-%v-%v", r.Name, IP, r.Port), // 服务节点的名称
        Name:    r.Name,
        Tags:    r.Tag,
        Port:    r.Port,
        Address: IP, // 服务 IP
        Check: &api.AgentServiceCheck{
            Interval:                       r.Interval.String(), // 健康检查间隔
            HTTP:                           fmt.Sprintf("http://%s/v1/agent/checks", r.Addr),
            DeregisterCriticalServiceAfter: r.DeregisterCriticalServiceAfter.String(),
            // GRPC:     fmt.Sprintf("%s/%v", r.Addr, r.Name),
        },
    }
    if err := agent.ServiceRegister(reg); err != nil {
        return err
    }

    return nil
}

func localIP() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, address := range addrs {
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}
