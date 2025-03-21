// Copyright 2018 The Cloudprober Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This binary implements a standalone ResourceDiscovery service (RDS) client.
package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"flag"

	"google.golang.org/protobuf/proto"

	tlsconfigpb "github.com/cloudprober/cloudprober/common/tlsconfig/proto"
	"github.com/cloudprober/cloudprober/internal/rds/client"
	configpb "github.com/cloudprober/cloudprober/internal/rds/client/proto"
	pb "github.com/cloudprober/cloudprober/internal/rds/proto"
	"github.com/cloudprober/cloudprober/logger"
)

var (
	rdsServer      = flag.String("rds_server", "", "gRPC server address")
	rdsServerCert  = flag.String("rds_server_cert", "", "gRPC server cert to use.")
	certServerName = flag.String("cert_server_name", "", "Server name for the cert, defaults to project name.")
	provider       = flag.String("provider", "gcp", "Resource provider")
	resType        = flag.String("resource_type", "gce_instances", "Resource type")
	publicIP       = flag.Bool("public_ip", false, "Public IP instead of private")
	project        = flag.String("project", "", "GCP project")
	filtersF       = flag.String("filters", "", "Comma separated list of filters, e.g. name=ig-us-central1-a-.*")
)

func main() {
	flag.Parse()

	if *project == "" {
		log.Fatal("--project is a required paramater")
	}

	serverName := *certServerName
	if serverName == "" {
		serverName = *project

		if strings.Contains(serverName, ":") {
			parts := strings.SplitN(serverName, ":", 2)
			serverName = parts[1] + "." + parts[0]
		}
	}

	c := &configpb.ClientConf{
		ServerOptions: &configpb.ClientConf_ServerOptions{
			ServerAddress: rdsServer,
		},
	}

	if *rdsServerCert != "" {
		c.ServerOptions.TlsConfig = &tlsconfigpb.TLSConfig{
			CaCertFile: rdsServerCert,
			ServerName: proto.String(serverName),
		}
	}

	c.Request = &pb.ListResourcesRequest{
		Provider:     proto.String(*provider),
		ResourcePath: proto.String(fmt.Sprintf("%s/%s", *resType, *project)),
	}

	if *publicIP {
		c.Request.IpConfig = &pb.IPConfig{
			IpType: pb.IPConfig_PUBLIC.Enum(),
		}
	}

	for _, f := range strings.Split(*filtersF, ",") {
		if f == "" {
			continue
		}
		fParts := strings.SplitN(f, "=", 2)
		if len(fParts) != 2 {
			log.Fatalf("bad filter in --filters flag (%s): %s", *filtersF, f)
		}
		c.Request.Filter = append(c.Request.Filter, &pb.Filter{
			Key:   proto.String(fParts[0]),
			Value: proto.String(fParts[1]),
		})
	}

	tgts, err := client.New(c, nil, &logger.Logger{})
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Printf("%s\n", time.Now())

		for _, ep := range tgts.ListEndpoints() {
			ip, _ := tgts.Resolve(ep.Name, 4)
			fmt.Printf("%s\t%s\n", ep.Name, ip.String())
		}
		time.Sleep(5 * time.Second)
	}
}
