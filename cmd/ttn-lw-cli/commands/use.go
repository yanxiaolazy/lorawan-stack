// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcmiddleware/discover"
	"gopkg.in/yaml.v2"
)

const configFileName = ".ttn-lw-cli.yml"

func sliceContains(slice []string, val string) bool {
	for _, v := range slice {
		if val == v {
			return true
		}
	}
	return false
}

func useCommandWriteConfigFlags() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{}
	flagSet.Bool("user", false, "Write config file in user config directory")
	flagSet.Bool("overwrite", false, "Overwrite existing config files")
	return flagSet
}

func getConfigFileDestinationPath(base string, user bool, overwrite bool) (string, error) {
	fileName := base
	if user {
		dir, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}
		if err = os.MkdirAll(dir, 0755); err != nil {
			return "", err
		}
		fileName = filepath.Join(dir, base)
	}
	_, err := os.Stat(fileName)
	if !os.IsNotExist(err) && !overwrite {
		logger.Warnf("%s already exists. Use --overwrite", fileName)
		return "", errFileExists.WithAttributes("file", fileName)
	}
	return fileName, nil
}

var (
	errNoHost          = errors.DefineInvalidArgument("no_host", "no host set")
	errFileExists      = errors.DefineAlreadyExists("file_exists", "`{file}` exists")
	errFailWrite       = errors.DefinePermissionDenied("fail_write", "failed to write `{file}`")
	errNoCluster       = errors.DefineInvalidArgument("no_cluster", "no cluster specified")
	errInvalidTenantID = errors.DefineInvalidArgument("invalid_tenant_id", "invalid tenant ID `{tenant_id}`")

	knownCommunityClusters = []string{"eu1", "nam1", "au1"}
	knownCloudClusters     = []string{"eu1", "nam1", "au1"}
	identifierRegex        = regexp.MustCompile(`^[a-z][a-z0-9-]{2,}$`)

	useCloudCommand = &cobra.Command{
		Use:               fmt.Sprintf("use-cloud --cluster %v --tenant tenant-id", knownCloudClusters),
		Aliases:           []string{"generate-configuration-cloud", "generate-cfg-cloud"},
		Short:             "Generate client configuration for The Things Stack Cloud",
		PersistentPreRunE: preRun(),
		RunE: func(cmd *cobra.Command, args []string) error {
			cluster, _ := cmd.Flags().GetString("cluster")
			tenant, _ := cmd.Flags().GetString("tenant")

			cluster = strings.TrimSuffix(strings.TrimSpace(cluster), ".cloud.thethings.industries")
			if cluster == "" {
				logger.Warnf("Please specify the The Things Stack Cloud cluster you want to connect to. Known cluster names: %v", knownCloudClusters)
				logger.Warn("For more information, see https://console.cloud.thethings.industries/")
				cmd.Usage()
				return errNoCluster.New()
			} else if !sliceContains(knownCommunityClusters, cluster) {
				logger.Warnf("Unknown cluster name '%s'! Your configuration may be incorrect!", cluster)
				logger.Warnf("Unless you know what you are doing, please use one of the following cluster names: %v", knownCloudClusters)
				logger.Warn("For more information, see https://console.cloud.thethings.industries/")
			}
			if !identifierRegex.MatchString(tenant) {
				logger.Warnf("Invalid tenant identifier '%s'! Your configuration is incorrect!", tenant)
				logger.Warn("For more information, see https://console.cloud.thethings.industries/")
				return errInvalidTenantID.WithAttributes("tenant_id", tenant)
			}

			insecure := false
			host := fmt.Sprintf("%s.%s.cloud.thethings.industries", tenant, cluster)
			grpcPort := discover.DefaultPorts[!insecure]
			config := MakeDefaultConfig(fmt.Sprintf("%s:%d", host, grpcPort), fmt.Sprintf("https://%s.eu1.cloud.thethings.industries/oauth", tenant), insecure)
			config.IdentityServerGRPCAddress = fmt.Sprintf("%s.eu1.cloud.thethings.industries:%d", tenant, grpcPort)
			config.CredentialsID = host

			b, err := yaml.Marshal(config)
			if err != nil {
				return err
			}
			user, _ := cmd.Flags().GetBool("user")
			overwrite, _ := cmd.Flags().GetBool("overwrite")
			configFile, err := getConfigFileDestinationPath(configFileName, user, overwrite)
			if err != nil {
				return err
			}
			if err := ioutil.WriteFile(configFile, b, 0644); err != nil {
				return err
			}
			logger.Infof("Config file for %s written in %s", host, configFile)
			return nil
		},
	}
	useCommunityCommand = &cobra.Command{
		Use:               fmt.Sprintf("use-community %v", knownCommunityClusters),
		Aliases:           []string{"generate-configuration-community", "generate-cfg-community"},
		Short:             "Generate client configuration for The Things Stack Community Edition",
		PersistentPreRunE: preRun(),
		RunE: func(cmd *cobra.Command, args []string) error {
			var cluster string
			switch {
			case len(args) == 1:
				cluster = args[0]
			case len(args) > 1:
				logger.Warnf("Received multiple arguments, considering '%s' as cluster name", args[0])
				cluster = args[0]
			}

			cluster = strings.TrimSuffix(strings.TrimSpace(cluster), ".cloud.thethings.network")
			if cluster == "" {
				logger.Warnf("Please specify the community cluster you want to connect to. Known cluster names: %v", knownCommunityClusters)
				logger.Warn("For more information, see https://console.cloud.thethings.network/")
				cmd.Usage()
				return errNoCluster.New()
			} else if !sliceContains(knownCommunityClusters, cluster) {
				logger.Warnf("Unknown cluster name '%s'! Your configuration may be incorrect!", cluster)
				logger.Warnf("Unless you know what you are doing, please use one of the following cluster names: %v", knownCommunityClusters)
			}
			insecure := false
			host := fmt.Sprintf("%s.cloud.thethings.network", cluster)
			grpcPort := discover.DefaultPorts[!insecure]
			config := MakeDefaultConfig(fmt.Sprintf("%s:%d", host, grpcPort), "https://eu1.cloud.thethings.network/oauth", insecure)
			config.IdentityServerGRPCAddress = fmt.Sprintf("eu1.cloud.thethings.network:%d", grpcPort)
			config.CredentialsID = host

			b, err := yaml.Marshal(config)
			if err != nil {
				return err
			}
			user, _ := cmd.Flags().GetBool("user")
			overwrite, _ := cmd.Flags().GetBool("overwrite")
			configFile, err := getConfigFileDestinationPath(configFileName, user, overwrite)
			if err != nil {
				return err
			}
			if err := ioutil.WriteFile(configFile, b, 0644); err != nil {
				return err
			}
			logger.Infof("Config file for %s written in %s", host, configFile)
			return nil
		},
	}
	useCommand = &cobra.Command{
		Use:               "use",
		Aliases:           []string{"generate-configuration", "generate-cfg"},
		Example:           `use localhost --insecure --oauth-server-address http://localhost:1885/oauth`,
		Short:             "Generate client configuration for The Things Stack deployments",
		PersistentPreRunE: preRun(),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errNoHost
			}
			insecure, _ := cmd.Flags().GetBool("insecure")
			fetchCA, _ := cmd.Flags().GetBool("fetch-ca")
			user, _ := cmd.Flags().GetBool("user")
			overwrite, _ := cmd.Flags().GetBool("overwrite")

			host := args[0]
			grpcPort, _ := cmd.Flags().GetInt("grpc-port")
			if grpcPort == 0 {
				grpcPort = discover.DefaultPorts[!insecure]
			}
			grpcServerAddress, err := discover.DefaultPort(host, grpcPort)
			if err != nil {
				return err
			}
			oauthServerAddress, _ := cmd.Flags().GetString("oauth-server-address")
			if oauthServerAddress == "" {
				oauthServerBaseAddress, err := discover.DefaultURL(host, discover.DefaultHTTPPorts[!insecure], !insecure)
				if err != nil {
					return err
				}
				oauthServerAddress = oauthServerBaseAddress + "/oauth"
			}
			conf := MakeDefaultConfig(grpcServerAddress, oauthServerAddress, insecure)
			conf.CredentialsID = host

			// Get CA certificate from server
			if !insecure && fetchCA {
				h := md5.New()
				io.WriteString(h, host)
				caFileName := fmt.Sprintf("ca.%s.pem", hex.EncodeToString(h.Sum(nil))[:6])
				caFile, err := getConfigFileDestinationPath(caFileName, user, overwrite)
				if err != nil {
					return err
				}

				logger.Infof("Will retrieve certificate from %s", conf.NetworkServerGRPCAddress)
				conn, err := tls.Dial("tcp", conf.NetworkServerGRPCAddress, &tls.Config{InsecureSkipVerify: true})
				if err != nil {
					return err
				}
				defer conn.Close()

				buf := &bytes.Buffer{}
				for _, cert := range conn.ConnectionState().PeerCertificates {
					if !cert.IsCA {
						continue
					}
					if err = pem.Encode(buf, &pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw}); err != nil {
						logger.Warnf("Could not retrieve certificate: %s", err)
					}
				}
				if err = ioutil.WriteFile(caFile, buf.Bytes(), 0644); err != nil {
					return errFailWrite.WithCause(err).WithAttributes("file", caFile)
				}
				logger.Infof("CA file for %s written in %s", host, caFile)
				abs, err := filepath.Abs(caFile)
				if err != nil {
					return err
				}
				conf.CA = abs
			}

			b, err := yaml.Marshal(conf)
			if err != nil {
				return err
			}
			configFile, err := getConfigFileDestinationPath(configFileName, user, overwrite)
			if err != nil {
				return err
			}
			if err = ioutil.WriteFile(configFile, b, 0644); err != nil {
				return errFailWrite.WithCause(err).WithAttributes("file", configFile)
			}
			logger.Infof("Config file for %s written in %s", host, configFile)
			return nil
		},
	}
)

func init() {
	useCommand.Flags().Bool("insecure", defaultInsecure, "Connect without TLS")
	useCommand.Flags().String("host", defaultClusterHost, "Server host name")
	useCommand.Flags().String("oauth-server-address", "", "OAuth Server address")
	useCommand.Flags().Bool("fetch-ca", false, "Connect to server and retrieve CA")
	useCommand.Flags().Int("grpc-port", 0, "")
	useCommand.Flags().AddFlagSet(useCommandWriteConfigFlags())
	useCommand.SetUsageTemplate(usageTemplateWithoutGlobalFlags)
	Root.AddCommand(useCommand)

	useCloudCommand.Flags().String("cluster", "", "The Things Stack Cloud cluster to use")
	useCloudCommand.Flags().String("tenant", "", "The Things Stack Cloud tenant to use")
	useCloudCommand.Flags().AddFlagSet(useCommandWriteConfigFlags())
	useCloudCommand.SetUsageTemplate(usageTemplateWithoutGlobalFlags)
	Root.AddCommand(useCloudCommand)

	useCommunityCommand.SetUsageTemplate(usageTemplateWithoutGlobalFlags)
	useCommunityCommand.Flags().AddFlagSet(useCommandWriteConfigFlags())
	Root.AddCommand(useCommunityCommand)
}
