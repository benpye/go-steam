package steam

import (
	"time"

	"github.com/benpye/go-steam/protocol"
	"github.com/benpye/go-steam/protocol/steamlang"
)

type License struct {
	PackageID           uint32
	TimeCreated         time.Time
	TimeNextProcess     time.Time
	MinuteLimit         int32
	MinutesUsed         int32
	PaymentMethod       steamlang.EPaymentMethod
	Flags               steamlang.ELicenseFlags
	PurchaseCountryCode string
	LicenseType         steamlang.ELicenseType
	TerritoryCode       int32
	ChangeNumber        int32
	// TODO: There are some additional fields here SteamKit doesn't handle
	//       but their meaning isn't immediately obvious
}

type LicenseListEvent struct {
	Result   steamlang.EResult
	Licenses []License
}

func (e *LicenseListEvent) GetJobID() protocol.JobID {
	return 0
}

type FreeLicenseEvent struct {
	JobID           protocol.JobID
	Result          steamlang.EResult
	GrantedPackages []uint32
	GrantedApps     []uint32
}

func (e *FreeLicenseEvent) GetJobID() protocol.JobID {
	return e.JobID
}

type AppOwnershipTicketEvent struct {
	JobID  protocol.JobID
	Result steamlang.EResult
	AppID  uint32
	Ticket []byte
}

func (e *AppOwnershipTicketEvent) GetJobID() protocol.JobID {
	return e.JobID
}

type DepotKeyEvent struct {
	JobID    protocol.JobID
	Result   steamlang.EResult
	DepotID  uint32
	DepotKey []byte
}

func (e *DepotKeyEvent) GetJobID() protocol.JobID {
	return e.JobID
}

type GameConnectTokensEvent struct {
	MaxTokensToKeep uint32
	Tokens          [][]byte
}

func (e *GameConnectTokensEvent) GetJobID() protocol.JobID {
	return 0
}
