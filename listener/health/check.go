// Package health for listeners
package health

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// CheckGroup represents a struct that checks on whether our service is ready to receive requests or not
type CheckGroup struct {
	build string
	log   *zerolog.Logger
	port  string
}

// NewCheckGroup creates a new instance of CheckGroup
func NewCheckGroup(log *zerolog.Logger, build string, port string) CheckGroup {
	return CheckGroup{
		build: build,
		log:   log,
		port:  port,
	}
}

// Liveness returns simple status info if the service is alive. If the
// app is deployed to a Kubernetes cluster, it will also return pod, node, and
// namespace details via the Downward API. The Kubernetes environment variables
// need to be set within your Pod/Deployment manifest.
func (cg *CheckGroup) Liveness(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		host = "unavailable"
	}

	info := struct {
		Status    string `json:"status,omitempty"`
		Build     string `json:"build,omitempty"`
		Host      string `json:"host,omitempty"`
		Pod       string `json:"pod,omitempty"`
		PodIP     string `json:"podIP,omitempty"`
		Node      string `json:"node,omitempty"`
		Namespace string `json:"namespace,omitempty"`
	}{
		Status:    "up",
		Build:     cg.build,
		Host:      host,
		Pod:       os.Getenv("KUBERNETES_PODNAME"),
		PodIP:     os.Getenv("KUBERNETES_NAMESPACE_POD_IP"),
		Node:      os.Getenv("KUBERNETES_NODENAME"),
		Namespace: os.Getenv("KUBERNETES_NAMESPACE"),
	}

	resp := struct {
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
	}{
		Success: true,
		Data:    info,
	}

	w.WriteHeader(http.StatusOK)

	//nolint:golint,errcheck,gosec
	json.NewEncoder(w).Encode(resp)
}

// StartListenerServer starts the http server
func (cg *CheckGroup) StartListenerServer(ctx context.Context) error {
	cg.log.Info().Msg("Starting http server at port " + cg.port)

	http.HandleFunc("/liveness", cg.Liveness)
	server := &http.Server{
		Addr:              ":" + cg.port,
		ReadHeaderTimeout: 3 * time.Second,
	}

	// Open port so we can return any errors immediately if this fails.
	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		cg.log.Err(err).Str("listen_port", server.Addr).Msg("error while trying to listen to port for liveness checks")
		return err
	}

	go func() {
		<-ctx.Done()
		err = server.Shutdown(context.TODO())
		if err != nil {
			cg.log.Err(err).Msg("error while shutting down http server for liveness check")
		}
	}()

	go func() {
		err = server.Serve(listener)
		if err != nil {
			cg.log.Err(err).Msg("error trying to server requests for liveness checks")
		}
	}()

	return err
}
