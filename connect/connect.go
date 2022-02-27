package relampago_connect

import (
	"strconv"
	"time"

	"github.com/rdbell/relampago"
	"github.com/rdbell/relampago/lnd"
	"github.com/rdbell/relampago/sparko"
	"github.com/rdbell/relampago/void"
)

type LightningBackendSettings struct {
	BackendType    string `envconfig:"LIGHTNING_BACKEND_TYPE" json:"lightning_backend_type"`
	ConnectTimeout string `envconfig:"LIGHTNING_CONNECT_TIMEOUT" json:"lightning_connect_timeout" default:"15"`

	SparkoURL   string `envconfig:"SPARKO_URL" json:"sparko_url"`
	SparkoToken string `envconfig:"SPARKO_TOKEN" json:"sparko_token"`

	LNDHost         string `envconfig:"LND_HOST" json:"lnd_host"`
	LNDCertPath     string `envconfig:"LND_CERT_PATH" json:"lnd_cert_path"`
	LNDMacaroonPath string `envconfig:"LND_MACAROON_PATH" json:"lnd_macaroon_path"`
}

func Connect(lbs LightningBackendSettings) (relampago.Wallet, error) {
	connectTimeout, err := strconv.Atoi(lbs.ConnectTimeout)
	if err != nil {
		return nil, err
	}

	// start lightning backend
	switch lbs.BackendType {
	case "lndrest":
	case "lndgrpc":
		return lnd.Start(lnd.Params{
			Host:           lbs.LNDHost,
			CertPath:       lbs.LNDCertPath,
			MacaroonPath:   lbs.LNDMacaroonPath,
			ConnectTimeout: time.Duration(connectTimeout) * time.Second,
		})
	case "eclair":
	case "clightning":
	case "sparko":
		return sparko.Start(sparko.Params{
			Host:           lbs.SparkoURL,
			Key:            lbs.SparkoToken,
			ConnectTimeout: time.Duration(connectTimeout) * time.Second,
		})
	case "lnbits":
	case "lnpay":
	case "zebedee":
	}

	// use void wallet that does nothing
	return void.Start()
}
