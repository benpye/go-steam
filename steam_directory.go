package steam

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/benpye/go-steam/netutil"
)

// Load initial server list from Steam Directory Web API.
// Call InitializeSteamDirectory() before Connect() to use
// steam directory server list instead of static one.
func InitializeSteamDirectory() error {
	return steamDirectoryCache.Initialize()
}

var steamDirectoryCache *steamDirectory = &steamDirectory{}

type steamDirectory struct {
	sync.RWMutex
	servers       []netutil.PortAddr
	isInitialized bool
}

// Get server list from steam directory and save it for later
func (sd *steamDirectory) Initialize() error {
	sd.Lock()
	defer sd.Unlock()

	// First try using web API to get CM list
	err := sd.getServerList()
	if err != nil {
		// Failing this attempt to use DNS
		err = sd.getServerListDNS()

		if err != nil {
			return err
		}
	}

	sd.isInitialized = true
	return nil
}

func (sd *steamDirectory) getServerList() error {
	client := new(http.Client)
	resp, err := client.Get(fmt.Sprintf("https://api.steampowered.com/ISteamDirectory/GetCMList/v1/?cellId=0"))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	r := struct {
		Response struct {
			ServerList []string
			Result     uint32
			Message    string
		}
	}{}
	if err = json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return err
	}
	if r.Response.Result != 1 {
		return fmt.Errorf("failed to get steam directory, result: %v, message: %v", r.Response.Result, r.Response.Message)
	}
	if len(r.Response.ServerList) == 0 {
		return fmt.Errorf("steam returned zero servers for steam directory request")
	}

	sd.servers = make([]netutil.PortAddr, 0, len(r.Response.ServerList))

	for _, server := range r.Response.ServerList {
		addr := netutil.ParsePortAddr(server)
		if addr != nil {
			sd.servers = append(sd.servers, *addr)
		}
	}

	return nil
}

func (sd *steamDirectory) getServerListDNS() error {
	iplist, err := net.LookupIP("cm0.steampowered.com")
	if err != nil {
		return err
	}

	sd.servers = make([]netutil.PortAddr, 0, len(iplist))

	for _, ip := range iplist {
		sd.servers = append(sd.servers, netutil.PortAddr{
			IP:   ip,
			Port: 27017,
		})
	}

	return nil
}

func (sd *steamDirectory) GetRandomCM() *netutil.PortAddr {
	sd.RLock()
	defer sd.RUnlock()
	if !sd.isInitialized {
		panic("steam directory is not initialized")
	}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &sd.servers[rng.Int31n(int32(len(sd.servers)))]
}

func (sd *steamDirectory) IsInitialized() bool {
	sd.RLock()
	defer sd.RUnlock()
	isInitialized := sd.isInitialized
	return isInitialized
}
