package steam

import (
	"errors"
	"sort"
	"sync"
	"time"

	"github.com/benpye/go-steam/netutil"
)

// ServerState represents the known state of a CM server
type ServerState uint32

const (
	// ServerStateUnknown represents a CM that has not yet been connected to
	ServerStateUnknown ServerState = iota
	// ServerStateGood represents a CM that has been successfully connected to
	ServerStateGood
	// ServerStateBad represents a CM that has failed to connect or has been forcibly disconnected
	ServerStateBad
)

// ServerInfo represents a single CM server state
type ServerInfo struct {
	state       ServerState
	address     *netutil.PortAddr
	lastFailure time.Time
}

// ServerList is used to query for and maintain state about CM servers
type ServerList struct {
	servers     []ServerInfo
	initialized bool
	lock        sync.RWMutex
}

func (s *ServerList) populateServerList() error {
	// If the list is already populated then there is nothing to do
	if s.initialized {
		return nil
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	// TODO: provide ability to specify cell ID
	servers, err := GetServers(0, 16)

	if err != nil {
		servers, err = GetServersDNS()

		if err != nil {
			return err
		}
	}

	if len(servers) == 0 {
		return errors.New("no Steam CM servers returned")
	}

	s.servers = make([]ServerInfo, 0, len(servers))

	for idx := range servers {
		s.servers = append(s.servers, ServerInfo{
			state:       ServerStateUnknown,
			address:     servers[idx],
			lastFailure: time.Time{},
		})
	}

	s.initialized = true
	return nil
}

func (s *ServerList) resort() {
	sort.SliceStable(s.servers, func(i, j int) bool { return s.servers[i].lastFailure.Before(s.servers[j].lastFailure) })
}

// GetServerCandidate gets the next server that may be available
func (s *ServerList) GetServerCandidate() (*netutil.PortAddr, error) {
	err := s.populateServerList()
	if err != nil {
		return nil, err
	}

	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.servers[0].address, nil
}

// UpdateServerState updates the state of the provided server
func (s *ServerList) UpdateServerState(addr *netutil.PortAddr, state ServerState) error {
	// If we're not initialized then we must have gotten here because the user specified a server
	// We don't care to update in that case
	if !s.initialized {
		return nil
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	for idx, server := range s.servers {
		if server.address.Equal(addr) {
			needsSort := false

			if s.servers[idx].state != state {
				if state == ServerStateBad {
					s.servers[idx].lastFailure = time.Now()
					needsSort = true
				} else if !s.servers[idx].lastFailure.IsZero() {
					s.servers[idx].lastFailure = time.Time{}
					needsSort = true
				}
			}

			if needsSort {
				s.resort()
			}

			return nil
		}
	}

	return errors.New("address provided does not match server")
}

// UpdateList replaces the current server list - this is called when we recieve the CM server list
func (s *ServerList) UpdateList(serverList []*netutil.PortAddr) {
	s.lock.Lock()
	defer s.lock.Unlock()

	newServerList := make([]ServerInfo, 0, len(serverList))

	for _, addr := range serverList {
		server := ServerInfo{
			state:       ServerStateUnknown,
			address:     addr,
			lastFailure: time.Time{},
		}

		for _, oldServer := range s.servers {
			if oldServer.address.Equal(server.address) {
				server.lastFailure = oldServer.lastFailure
				server.state = oldServer.state
			}

			break
		}

		newServerList = append(newServerList, server)
	}

	if len(newServerList) > 0 {
		s.servers = newServerList
		s.resort()

		s.initialized = true
	}
}
