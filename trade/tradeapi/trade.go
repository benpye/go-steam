/*
Wrapper around the HTTP trading API for type safety 'n' stuff.
*/
package tradeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/benpye/go-steam/community"
	"github.com/benpye/go-steam/economy/inventory"
	"github.com/benpye/go-steam/netutil"
	"github.com/benpye/go-steam/steamid"
)

const tradeURL = "https://steamcommunity.com/trade/%d/"

type Trade struct {
	client *http.Client
	other  steamid.SteamID

	LogPos  uint // not automatically updated
	Version uint // Incremented for each item change by Steam; not automatically updated.

	// the `sessionid` cookie is sent as a parameter/POST data for CSRF protection.
	sessionID string
	baseURL   string
}

// Creates a new Trade based on the given cookies `sessionid`, `steamLogin`, `steamLoginSecure` and the trade partner's Steam ID.
func New(sessionId, steamLogin, steamLoginSecure string, other steamid.SteamID) *Trade {
	client := new(http.Client)
	client.Timeout = 10 * time.Second

	t := &Trade{
		client:    client,
		other:     other,
		sessionID: sessionId,
		baseURL:   fmt.Sprintf(tradeURL, other),
		Version:   1,
	}
	community.SetCookies(t.client, sessionId, steamLogin, steamLoginSecure)
	return t
}

type Main struct {
	PartnerOnProbation bool
}

var onProbationRegex = regexp.MustCompile(`var g_bTradePartnerProbation = (\w+);`)

// Fetches the main HTML page and parses it. Thread-safe.
func (t *Trade) GetMain() (*Main, error) {
	resp, err := t.client.Get(t.baseURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	match := onProbationRegex.FindSubmatch(body)
	if len(match) == 0 {
		return nil, errors.New("tradeapi.GetMain: Could not find probation info")
	}

	return &Main{
		string(match[1]) == "true",
	}, nil
}

// Ajax POSTs to an API endpoint that should return a status
func (t *Trade) postWithStatus(url string, data map[string]string) (*Status, error) {
	status := new(Status)

	req := netutil.NewPostForm(url, netutil.ToURLValues(data))
	// Tales of Madness and Pain, Episode 1: If you forget this, Steam will return an error
	// saying "missing required parameter", even though they are all there. IT WAS JUST THE HEADER, ARGH!
	req.Header.Add("Referer", t.baseURL)

	resp, err := t.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(status)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func (t *Trade) GetStatus() (*Status, error) {
	return t.postWithStatus(t.baseURL+"tradestatus/", map[string]string{
		"sessionid": t.sessionID,
		"logpos":    strconv.FormatUint(uint64(t.LogPos), 10),
		"version":   strconv.FormatUint(uint64(t.Version), 10),
	})
}

// Thread-safe.
func (t *Trade) GetForeignInventory(contextID uint64, appID uint32, start *uint) (*inventory.PartialInventory, error) {
	data := map[string]string{
		"sessionid": t.sessionID,
		"steamid":   fmt.Sprintf("%d", t.other),
		"contextid": strconv.FormatUint(contextID, 10),
		"appid":     strconv.FormatUint(uint64(appID), 10),
	}
	if start != nil {
		data["start"] = strconv.FormatUint(uint64(*start), 10)
	}

	req, err := http.NewRequest("GET", t.baseURL+"foreigninventory?"+netutil.ToURLValues(data).Encode(), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Referer", t.baseURL)

	return inventory.DoInventoryRequest(t.client, req)
}

// Thread-safe.
func (t *Trade) GetOwnInventory(contextID uint64, appID uint32) (*inventory.Inventory, error) {
	return inventory.GetOwnInventory(t.client, contextID, appID)
}

func (t *Trade) Chat(message string) (*Status, error) {
	return t.postWithStatus(t.baseURL+"chat", map[string]string{
		"sessionid": t.sessionID,
		"logpos":    strconv.FormatUint(uint64(t.LogPos), 10),
		"version":   strconv.FormatUint(uint64(t.Version), 10),
		"message":   message,
	})
}

func (t *Trade) AddItem(slot uint, itemID, contextID uint64, appID uint32) (*Status, error) {
	return t.postWithStatus(t.baseURL+"additem", map[string]string{
		"sessionid": t.sessionID,
		"slot":      strconv.FormatUint(uint64(slot), 10),
		"itemid":    strconv.FormatUint(itemID, 10),
		"contextid": strconv.FormatUint(contextID, 10),
		"appid":     strconv.FormatUint(uint64(appID), 10),
	})
}

func (t *Trade) RemoveItem(slot uint, itemID, contextID uint64, appID uint32) (*Status, error) {
	return t.postWithStatus(t.baseURL+"removeitem", map[string]string{
		"sessionid": t.sessionID,
		"slot":      strconv.FormatUint(uint64(slot), 10),
		"itemid":    strconv.FormatUint(itemID, 10),
		"contextid": strconv.FormatUint(contextID, 10),
		"appid":     strconv.FormatUint(uint64(appID), 10),
	})
}

func (t *Trade) SetCurrency(amount uint, currencyID, contextID uint64, appID uint32) (*Status, error) {
	return t.postWithStatus(t.baseURL+"setcurrency", map[string]string{
		"sessionid":  t.sessionID,
		"amount":     strconv.FormatUint(uint64(amount), 10),
		"currencyid": strconv.FormatUint(uint64(currencyID), 10),
		"contextid":  strconv.FormatUint(contextID, 10),
		"appid":      strconv.FormatUint(uint64(appID), 10),
	})
}

func (t *Trade) SetReady(ready bool) (*Status, error) {
	return t.postWithStatus(t.baseURL+"toggleready", map[string]string{
		"sessionid": t.sessionID,
		"version":   strconv.FormatUint(uint64(t.Version), 10),
		"ready":     fmt.Sprint(ready),
	})
}

func (t *Trade) Confirm() (*Status, error) {
	return t.postWithStatus(t.baseURL+"confirm", map[string]string{
		"sessionid": t.sessionID,
		"version":   strconv.FormatUint(uint64(t.Version), 10),
	})
}

func (t *Trade) Cancel() (*Status, error) {
	return t.postWithStatus(t.baseURL+"cancel", map[string]string{
		"sessionid": t.sessionID,
	})
}

func isSuccess(v interface{}) bool {
	if m, ok := v.(map[string]interface{}); ok {
		return m["success"] == true
	}
	return false
}
