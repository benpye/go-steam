package steam

import (
	"time"

	"github.com/benpye/go-steam/protocol"
	"github.com/benpye/go-steam/protocol/protobuf/steam"
	"github.com/benpye/go-steam/protocol/steamlang"
)

type App struct {
	client *Client
}

// RequestFreeLicense requests a free license for a list of app IDs.
func (a *App) RequestFreeLicense(apps []uint32) *AsyncJob {
	body := new(steam.CMsgClientRequestFreeLicense)
	msg := protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientRequestFreeLicense, body)

	body.Appids = apps

	job := a.client.JobManager.NewJob()
	msg.SetSourceJobID(job.JobID)
	a.client.Write(msg)

	return job
}

// GetAppOwnershipTicket gets the app ownership ticket for an app.
func (a *App) GetAppOwnershipTicket(appID uint32) *AsyncJob {
	body := new(steam.CMsgClientGetAppOwnershipTicket)
	msg := protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientGetAppOwnershipTicket, body)

	body.AppId = &appID

	job := a.client.JobManager.NewJob()
	msg.SetSourceJobID(job.JobID)
	a.client.Write(msg)

	return job
}

// GetDepotDecryptionKey returns the decryption key for a specific depot.
func (a *App) GetDepotDecryptionKey(depotID uint32, appID uint32) *AsyncJob {
	body := new(steam.CMsgClientGetDepotDecryptionKey)
	msg := protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientGetDepotDecryptionKey, body)

	body.DepotId = &depotID
	body.AppId = &appID

	job := a.client.JobManager.NewJob()
	msg.SetSourceJobID(job.JobID)
	a.client.Write(msg)

	return job
}

func (a *App) HandlePacket(packet *protocol.Packet) {
	switch packet.EMsg {
	case steamlang.EMsg_ClientLicenseList:
		a.handleClientLicenseList(packet)
	case steamlang.EMsg_ClientRequestFreeLicenseResponse:
		a.handleClientRequestFreeLicenseResponse(packet)
	case steamlang.EMsg_ClientGameConnectTokens:
		a.handleClientGameConnectTokens(packet)
	case steamlang.EMsg_ClientGetAppOwnershipTicketResponse:
		a.handleClientGetAppOwnershipTicketResponse(packet)
	case steamlang.EMsg_ClientGetDepotDecryptionKeyResponse:
		a.handleClientGetDepotDecryptionKeyResponse(packet)
	}
}

func (a *App) handleClientLicenseList(packet *protocol.Packet) {
	msg := new(steam.CMsgClientLicenseList)
	packet.ReadProtoMsg(msg)

	list := msg.GetLicenses()
	licenseList := make([]License, 0, len(list))

	for _, license := range list {
		licenseList = append(licenseList, License{
			PackageID:           license.GetPackageId(),
			TimeCreated:         time.Unix(int64(license.GetTimeCreated()), 0),
			TimeNextProcess:     time.Unix(int64(license.GetTimeNextProcess()), 0),
			MinuteLimit:         license.GetMinuteLimit(),
			MinutesUsed:         license.GetMinutesUsed(),
			PaymentMethod:       steamlang.EPaymentMethod(license.GetPaymentMethod()),
			Flags:               steamlang.ELicenseFlags(license.GetFlags()),
			PurchaseCountryCode: license.GetPurchaseCountryCode(),
			LicenseType:         steamlang.ELicenseType(license.GetLicenseType()),
			TerritoryCode:       license.GetTerritoryCode(),
			ChangeNumber:        license.GetChangeNumber(),
		})
	}

	a.client.Emit(&LicenseListEvent{
		Result:   steamlang.EResult(msg.GetEresult()),
		Licenses: licenseList,
	})
}

func (a *App) handleClientRequestFreeLicenseResponse(packet *protocol.Packet) {
	msg := new(steam.CMsgClientRequestFreeLicenseResponse)
	packet.ReadProtoMsg(msg)

	a.client.Emit(&FreeLicenseEvent{
		JobID:           packet.TargetJobID,
		Result:          steamlang.EResult(msg.GetEresult()),
		GrantedPackages: msg.GetGrantedPackageids(),
		GrantedApps:     msg.GetGrantedAppids(),
	})
}

func (a *App) handleClientGameConnectTokens(packet *protocol.Packet) {
	msg := new(steam.CMsgClientGameConnectTokens)
	packet.ReadProtoMsg(msg)

	a.client.Emit(&GameConnectTokensEvent{
		MaxTokensToKeep: msg.GetMaxTokensToKeep(),
		Tokens:          msg.GetTokens(),
	})
}

func (a *App) handleClientGetAppOwnershipTicketResponse(packet *protocol.Packet) {
	msg := new(steam.CMsgClientGetAppOwnershipTicketResponse)
	packet.ReadProtoMsg(msg)

	a.client.Emit(&AppOwnershipTicketEvent{
		JobID:  packet.TargetJobID,
		Result: steamlang.EResult(msg.GetEresult()),
		AppID:  msg.GetAppId(),
		Ticket: msg.GetTicket(),
	})
}

func (a *App) handleClientGetDepotDecryptionKeyResponse(packet *protocol.Packet) {
	msg := new(steam.CMsgClientGetDepotDecryptionKeyResponse)
	packet.ReadProtoMsg(msg)

	a.client.Emit(&DepotKeyEvent{
		JobID:    packet.TargetJobID,
		Result:   steamlang.EResult(msg.GetEresult()),
		DepotID:  msg.GetDepotId(),
		DepotKey: msg.GetDepotEncryptionKey(),
	})
}
