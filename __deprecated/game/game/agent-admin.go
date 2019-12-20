package game

import (
	"net"
	"sync"
)

type AgentAdmin struct {
	acc  AgentID
	info *agentMap
}

func NewAgentAdmin() *AgentAdmin {
	return &AgentAdmin{
		info: newAgentMap(),
	}
}

func (p *AgentAdmin) id() AgentID {
	p.acc = p.acc + 1
	return p.acc
}

func (p *AgentAdmin) RunAgentAsync(peer *net.TCPConn, stop <-chan bool, wg *sync.WaitGroup) {
	id := p.id()
	agent := NewAgent(id, peer, stop, wg, func() {
		p.info.Remove(id)
	})
	p.info.Add(id, agent)
	go agent.loop()
}

type agentInfo struct {
	agent *Agent
}

type agentMap struct {
	agents map[AgentID]*agentInfo
	sync.Mutex
}

func newAgentMap() *agentMap {
	return &agentMap{
		agents: map[AgentID]*agentInfo{},
	}
}

func (p *agentMap) Add(id AgentID, agent *Agent) {
	p.Lock()
	// if exist?
	p.agents[id] = &agentInfo{agent}
	p.Unlock()
}

func (p *agentMap) Remove(id AgentID) {
	p.Lock()
	// if not exist?
	delete(p.agents, id)
	p.Unlock()
}
