package config

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strconv"
	"time"

	"github.com/lncm/lnd-rpc/v0.10.0/lnrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/macaroon.v2"
)

type rpcCreds map[string]string

func (m rpcCreds) RequireTransportSecurity() bool { return true }
func (m rpcCreds) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return m, nil
}
func newCreds(bytes []byte) rpcCreds {
	creds := make(map[string]string)
	creds["macaroon"] = hex.EncodeToString(bytes)
	return creds
}

func getClient(hostname string, port int, tlsFile, macaroonFile string) lnrpc.LightningClient {
	macaroonBytes, err := ioutil.ReadFile(macaroonFile)
	if err != nil {
		panic(fmt.Sprintln("Cannot read macaroon file", err))
	}

	mac := &macaroon.Macaroon{}
	if err = mac.UnmarshalBinary(macaroonBytes); err != nil {
		panic(fmt.Sprintln("Cannot unmarshal macaroon", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	transportCredentials, err := credentials.NewClientTLSFromFile(tlsFile, hostname)
	if err != nil {
		panic(err)
	}

	fullHostname := fmt.Sprintf("%s:%d", hostname, port)

	connection, err := grpc.DialContext(ctx, fullHostname, []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(transportCredentials),
		grpc.WithPerRPCCredentials(newCreds(macaroonBytes)),
	}...)
	if err != nil {
		panic(fmt.Errorf("unable to connect to %s: %w", fullHostname, err))
	}

	return lnrpc.NewLightningClient(connection)
}

func NewLnConnection() (lnrpc.LightningClient, error) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	homeDir := usr.HomeDir
	lndir := os.Getenv("LN_DIR")
	lndDir := fmt.Sprintf("%s/%s", homeDir, lndir)

	lnTlsCert := os.Getenv("LN_TLS_CERT")
	lnMacaroon := os.Getenv("LN_MACAROON_FILE")
	lnHost := os.Getenv("LN_HOST")
	lnPort := os.Getenv("LN_PORT")

	p, _ := strconv.Atoi(lnPort)

	var (
		hostname     = lnHost
		port         = p
		tlsFile      = fmt.Sprintf("%s/%s", lndDir, lnTlsCert)
		macaroonFile = fmt.Sprintf("%s/%s", lndDir, lnMacaroon)
	)

	client := getClient(hostname, port, tlsFile, macaroonFile)

	return client, nil
}
