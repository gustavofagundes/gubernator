/*
Copyright 2018-2022 Mailgun Technologies Inc

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gubernator_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"testing"

	"github.com/gubernator-io/gubernator/v3"
	"github.com/mailgun/holster/v4/clock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/http2"
)

func spawnDaemon(t *testing.T, conf gubernator.DaemonConfig) *gubernator.Daemon {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), clock.Second*10)
	d, err := gubernator.SpawnDaemon(ctx, conf)
	cancel()
	require.NoError(t, err)
	d.SetPeers([]gubernator.PeerInfo{{HTTPAddress: conf.HTTPListenAddress, IsOwner: true}})
	return d
}

func makeRequest(t *testing.T, conf gubernator.DaemonConfig) error {
	t.Helper()

	client, err := gubernator.NewClient(gubernator.WithTLS(conf.ClientTLS(), conf.HTTPListenAddress))
	require.NoError(t, err)

	var resp gubernator.CheckRateLimitsResponse
	err = client.CheckRateLimits(context.Background(), &gubernator.CheckRateLimitsRequest{
		Requests: []*gubernator.RateLimitRequest{
			{
				Name:      "test_tls",
				UniqueKey: "account:995",
				Algorithm: gubernator.Algorithm_TOKEN_BUCKET,
				Duration:  gubernator.Second * 30,
				Limit:     100,
				Hits:      1,
			},
		},
	}, &resp)

	if err != nil {
		return err
	}
	rl := resp.Responses[0]
	assert.Equal(t, "", rl.Error)
	return nil
}

func TestSetupTLS(t *testing.T) {
	tests := []struct {
		tls  *gubernator.TLSConfig
		name string
	}{
		{
			name: "user provided certificates",
			tls: &gubernator.TLSConfig{
				CaFile:   "contrib/certs/ca.cert",
				CertFile: "contrib/certs/gubernator.pem",
				KeyFile:  "contrib/certs/gubernator.key",
			},
		},
		{
			name: "user provided certificate without IP SANs",
			tls: &gubernator.TLSConfig{
				CaFile:               "contrib/certs/ca.cert",
				CertFile:             "contrib/certs/gubernator_no_ip_san.pem",
				KeyFile:              "contrib/certs/gubernator_no_ip_san.key",
				ClientAuthServerName: "gubernator",
			},
		},
		{
			name: "auto tls",
			tls: &gubernator.TLSConfig{
				AutoTLS: true,
			},
		},
		{
			name: "generate server certs with user provided ca",
			tls: &gubernator.TLSConfig{
				CaFile:    "contrib/certs/ca.cert",
				CaKeyFile: "contrib/certs/ca.key",
				AutoTLS:   true,
			},
		},
		{
			name: "client auth enabled",
			tls: &gubernator.TLSConfig{
				CaFile:     "contrib/certs/ca.cert",
				CaKeyFile:  "contrib/certs/ca.key",
				AutoTLS:    true,
				ClientAuth: tls.RequireAndVerifyClientCert,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := gubernator.DaemonConfig{
				HTTPListenAddress: "127.0.0.1:9685",
				TLS:               tt.tls,
			}

			d := spawnDaemon(t, conf)

			client, err := gubernator.NewClient(gubernator.WithTLS(conf.ClientTLS(), conf.HTTPListenAddress))
			require.NoError(t, err)

			var resp gubernator.CheckRateLimitsResponse
			err = client.CheckRateLimits(context.Background(), &gubernator.CheckRateLimitsRequest{
				Requests: []*gubernator.RateLimitRequest{
					{
						Name:      "test_tls",
						UniqueKey: "account:995",
						Algorithm: gubernator.Algorithm_TOKEN_BUCKET,
						Duration:  gubernator.Second * 30,
						Limit:     100,
						Hits:      1,
					},
				},
			}, &resp)
			require.NoError(t, err)

			rl := resp.Responses[0]
			assert.Equal(t, "", rl.Error)
			assert.Equal(t, gubernator.Status_UNDER_LIMIT, rl.Status)
			assert.Equal(t, int64(99), rl.Remaining)
			d.Close(context.Background())
		})
	}
}

func TestSetupTLSSkipVerify(t *testing.T) {
	conf := gubernator.DaemonConfig{
		HTTPListenAddress: "127.0.0.1:9685",
		TLS: &gubernator.TLSConfig{
			CaFile:   "contrib/certs/ca.cert",
			CertFile: "contrib/certs/gubernator.pem",
			KeyFile:  "contrib/certs/gubernator.key",
		},
	}

	d := spawnDaemon(t, conf)
	defer d.Close(context.Background())

	tls := &gubernator.TLSConfig{
		AutoTLS:            true,
		InsecureSkipVerify: true,
	}

	err := gubernator.SetupTLS(tls)
	require.NoError(t, err)
	conf.TLS = tls

	err = makeRequest(t, conf)
	require.NoError(t, err)
}

func TestSetupTLSClientAuth(t *testing.T) {
	serverTLS := gubernator.TLSConfig{
		CaFile:           "contrib/certs/ca.cert",
		CertFile:         "contrib/certs/gubernator.pem",
		KeyFile:          "contrib/certs/gubernator.key",
		ClientAuth:       tls.RequireAndVerifyClientCert,
		ClientAuthCaFile: "contrib/certs/client-auth-ca.pem",
	}

	conf := gubernator.DaemonConfig{
		HTTPListenAddress: "127.0.0.1:9685",
		TLS:               &serverTLS,
	}

	d := spawnDaemon(t, conf)
	defer d.Close(context.Background())

	// Given generated client certs
	tls := &gubernator.TLSConfig{
		AutoTLS:            true,
		InsecureSkipVerify: true,
	}

	err := gubernator.SetupTLS(tls)
	require.NoError(t, err)
	conf.TLS = tls

	// Should not be allowed without a cert signed by the client CA
	err = makeRequest(t, conf)
	require.Error(t, err)
	// Error is different depending on golang version
	//assert.Contains(t, err.Error(), "tls: certificate required")

	// Given the client auth certs
	tls = &gubernator.TLSConfig{
		CertFile:           "contrib/certs/client-auth.pem",
		KeyFile:            "contrib/certs/client-auth.key",
		InsecureSkipVerify: true,
	}

	err = gubernator.SetupTLS(tls)
	require.NoError(t, err)
	conf.TLS = tls

	// Should be allowed to connect and make requests
	err = makeRequest(t, conf)
	require.NoError(t, err)
}

func TestTLSClusterWithClientAuthentication(t *testing.T) {
	serverTLS := gubernator.TLSConfig{
		CaFile:     "contrib/certs/ca.cert",
		CertFile:   "contrib/certs/gubernator.pem",
		KeyFile:    "contrib/certs/gubernator.key",
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	d1 := spawnDaemon(t, gubernator.DaemonConfig{
		HTTPListenAddress: "127.0.0.1:9685",
		TLS:               &serverTLS,
	})
	defer d1.Close(context.Background())

	d2 := spawnDaemon(t, gubernator.DaemonConfig{
		HTTPListenAddress: "127.0.0.1:9686",
		TLS:               &serverTLS,
	})
	defer d2.Close(context.Background())

	peers := []gubernator.PeerInfo{
		{
			HTTPAddress: d1.Listener.Addr().String(),
		},
		{
			HTTPAddress: d2.Listener.Addr().String(),
		},
	}
	d1.SetPeers(peers)
	d2.SetPeers(peers)

	// Should result in a remote call to d2
	err := makeRequest(t, d1.Config())
	require.NoError(t, err)

	config := d2.Config()
	client := &http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: config.ClientTLS(),
		},
	}

	resp, err := client.Get(fmt.Sprintf("https://%s/metrics", config.HTTPListenAddress))
	require.NoError(t, err)
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	// Should have called /v1/peer.forward on d2
	assert.Contains(t, string(b), `{path="`+gubernator.RPCPeerForward+`"} 1`)
}

func TestHTTPSClientAuth(t *testing.T) {
	conf := gubernator.DaemonConfig{
		HTTPListenAddress:       "127.0.0.1:9685",
		HTTPStatusListenAddress: "127.0.0.1:9686",
		TLS: &gubernator.TLSConfig{
			CaFile:     "contrib/certs/ca.cert",
			CertFile:   "contrib/certs/gubernator.pem",
			KeyFile:    "contrib/certs/gubernator.key",
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	d := spawnDaemon(t, conf)
	defer d.Close(context.Background())

	clientWithCert := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: conf.TLS.ClientTLS,
		},
	}

	clientWithoutCert := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: conf.TLS.ServerTLS.RootCAs,
			},
		},
	}

	reqCertRequired, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://%s/healthz", conf.HTTPListenAddress), nil)
	require.NoError(t, err)
	reqNoClientCertRequired, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://%s/healthz", conf.HTTPStatusListenAddress), nil)
	require.NoError(t, err)

	// Test that a client without a cert can access /v1/HealthCheck at status address
	resp, err := clientWithoutCert.Do(reqNoClientCertRequired)
	require.NoError(t, err)
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, `{"status":"healthy","peer_count":1}`, strings.ReplaceAll(string(b), " ", ""))

	// Verify we get an error when we try to access existing HTTPListenAddress without cert
	_, err = clientWithoutCert.Do(reqCertRequired) //nolint:all
	require.Error(t, err)
	// The error message is different depending on what version of golang is being used
	//assert.Contains(t, err.Error(), "remote error: tls: certificate required")

	// Check that with a valid client cert we can access /v1/healthz at existing HTTPListenAddress
	resp3, err := clientWithCert.Do(reqCertRequired)
	require.NoError(t, err)
	defer resp3.Body.Close()
	b, err = io.ReadAll(resp3.Body)
	require.NoError(t, err)
	assert.Equal(t, `{"status":"healthy","peer_count":1}`, strings.ReplaceAll(string(b), " ", ""))
}

// Ensure SpawnDaemon() setup peer TLS auth properly
func TestHTTPSDaemonPeerAuth(t *testing.T) {
	var daemons []*gubernator.Daemon
	var peers []gubernator.PeerInfo
	name := t.Name()

	for _, peer := range []gubernator.PeerInfo{
		{HTTPAddress: "127.0.0.1:9780"},
		{HTTPAddress: "127.0.0.1:9781"},
		{HTTPAddress: "127.0.0.1:9782"},
	} {
		ctx, cancel := context.WithTimeout(context.Background(), clock.Second*10)
		d, err := gubernator.SpawnDaemon(ctx, gubernator.DaemonConfig{
			Logger:            logrus.WithField("instance", peer.HTTPAddress),
			InstanceID:        peer.HTTPAddress,
			HTTPListenAddress: peer.HTTPAddress,
			AdvertiseAddress:  peer.HTTPAddress,
			DataCenter:        peer.DataCenter,
			Behaviors: gubernator.BehaviorConfig{
				// Suitable for testing but not production
				GlobalSyncWait: clock.Millisecond * 50,
				GlobalTimeout:  clock.Second * 5,
				BatchTimeout:   clock.Second * 5,
			},
			TLS: &gubernator.TLSConfig{
				CaFile:     "contrib/certs/ca.cert",
				CertFile:   "contrib/certs/gubernator.pem",
				KeyFile:    "contrib/certs/gubernator.key",
				ClientAuth: tls.RequireAndVerifyClientCert,
			},
		})
		cancel()
		require.NoError(t, err)
		peers = append(peers, d.PeerInfo)
		daemons = append(daemons, d)
	}

	for _, d := range daemons {
		d.SetPeers(peers)
	}

	defer func() {
		for _, d := range daemons {
			_ = d.Close(context.Background())
		}
	}()

	client, err := daemons[0].Client()
	require.NoError(t, err)

	for i := 0; i < 1_000; i++ {
		key := fmt.Sprintf("account:%08x", rand.Int())
		var resp gubernator.CheckRateLimitsResponse
		err := client.CheckRateLimits(context.Background(), &gubernator.CheckRateLimitsRequest{
			Requests: []*gubernator.RateLimitRequest{
				{
					Duration:  gubernator.Millisecond * 9000,
					Name:      name,
					UniqueKey: key,
					Limit:     20,
					Hits:      1,
				},
			},
		}, &resp)

		require.Nil(t, err)
		require.Equal(t, "", resp.Responses[0].Error)
	}

}
