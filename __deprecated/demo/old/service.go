package demo

import "gitlab.com/megatech/serverex/lib/util"

type IRunable interface {
	Run()
	Stop()
}

type IServiceInfo interface {
	ServiceID() string
	ServiceName() string
}

type ServiceInfo struct {
	ID, Name string
}

func NewServiceInfo(name string) ServiceInfo {
	return ServiceInfo{Name: name, ID: util.NewSnowflakeIDBase58()}
}

func (p *ServiceInfo) ServiceID() string {
	return p.ID
}

func (p *ServiceInfo) ServiceName() string {
	return p.Name
}
