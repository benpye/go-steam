package steam

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/benpye/go-steam/netutil"
)

// GetServers queries the Steam web API for the CM server list
func GetServers(cellID, maxCount int) ([]*netutil.PortAddr, error) {
	client := new(http.Client)

	url := fmt.Sprintf("https://api.steampowered.com/ISteamDirectory/GetCMList/v1/?cellId=%d", cellID)
	if maxCount > 0 {
		url = url + fmt.Sprintf("&maxcount=%d", maxCount)
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := struct {
		Response struct {
			ServerList           []string
			ServerListWebsockets []string `json:"serverlist_websockets"`
			Result               uint32
			Message              string
		}
	}{}

	if err = json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}

	if r.Response.Result != 1 {
		return nil, fmt.Errorf("failed to get steam directory, result: %v, message: %v", r.Response.Result, r.Response.Message)
	}

	if len(r.Response.ServerList) == 0 {
		return nil, fmt.Errorf("steam returned zero servers for steam directory request")
	}

	servers := make([]*netutil.PortAddr, 0, len(r.Response.ServerList))

	for _, server := range r.Response.ServerList {
		addr := netutil.ParsePortAddr(server)
		if addr != nil {
			servers = append(servers, addr)
		}
	}

	return servers, nil
}

// GetServersDNS queries the domain 'cm0.steampowered.com' for a set of fallback servers
func GetServersDNS() ([]*netutil.PortAddr, error) {
	iplist, err := net.LookupIP("cm0.steampowered.com")
	if err != nil {
		return nil, err
	}

	servers := make([]*netutil.PortAddr, 0, len(iplist))

	for _, ip := range iplist {
		servers = append(servers, &netutil.PortAddr{
			IP:   ip,
			Port: 27017,
		})
	}

	return servers, nil
}
