// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type detectExecuteScanOptions struct {
	APIToken       string   `json:"apiToken,omitempty"`
	CodeLocation   string   `json:"codeLocation,omitempty"`
	ProjectName    string   `json:"projectName,omitempty"`
	ProjectVersion string   `json:"projectVersion,omitempty"`
	Scanners       []string `json:"scanners,omitempty"`
	ScanPaths      []string `json:"scanPaths,omitempty"`
	ScanProperties []string `json:"scanProperties,omitempty"`
	ServerURL      string   `json:"serverUrl,omitempty"`
}

// DetectExecuteScanCommand Executes Synopsis Detect scan
func DetectExecuteScanCommand() *cobra.Command {
	const STEP_NAME = "detectExecuteScan"

	metadata := detectExecuteScanMetadata()
	var stepConfig detectExecuteScanOptions
	var startTime time.Time

	var createDetectExecuteScanCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "Executes Synopsis Detect scan",
		Long:  `This step executes [Synopsis Detect](https://synopsys.atlassian.net/wiki/spaces/INTDOCS/pages/62423113/Synopsys+Detect) scans.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			startTime = time.Now()
			log.SetStepName(STEP_NAME)
			log.SetVerbose(GeneralConfig.Verbose)

			path, _ := os.Getwd()
			fatalHook := &log.FatalHook{CorrelationID: GeneralConfig.CorrelationID, Path: path}
			log.RegisterHook(fatalHook)

			err := PrepareConfig(cmd, &metadata, STEP_NAME, &stepConfig, config.OpenPiperFile)
			if err != nil {
				return err
			}
			log.RegisterSecret(stepConfig.APIToken)

			if len(GeneralConfig.HookConfig.SentryConfig.Dsn) > 0 {
				sentryHook := log.NewSentryHook(GeneralConfig.HookConfig.SentryConfig.Dsn, GeneralConfig.CorrelationID)
				log.RegisterHook(&sentryHook)
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetry.Send(&telemetryData)
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, STEP_NAME)
			detectExecuteScan(stepConfig, &telemetryData)
			telemetryData.ErrorCode = "0"
		},
	}

	addDetectExecuteScanFlags(createDetectExecuteScanCmd, &stepConfig)
	return createDetectExecuteScanCmd
}

func addDetectExecuteScanFlags(cmd *cobra.Command, stepConfig *detectExecuteScanOptions) {
	cmd.Flags().StringVar(&stepConfig.APIToken, "apiToken", os.Getenv("PIPER_apiToken"), "Api token to be used for connectivity with Synopsis Detect server.")
	cmd.Flags().StringVar(&stepConfig.CodeLocation, "codeLocation", os.Getenv("PIPER_codeLocation"), "An override for the name Detect will use for the scan file it creates.")
	cmd.Flags().StringVar(&stepConfig.ProjectName, "projectName", os.Getenv("PIPER_projectName"), "Name of the Synopsis Detect (formerly BlackDuck) project.")
	cmd.Flags().StringVar(&stepConfig.ProjectVersion, "projectVersion", os.Getenv("PIPER_projectVersion"), "Version of the Synopsis Detect (formerly BlackDuck) project.")
	cmd.Flags().StringSliceVar(&stepConfig.Scanners, "scanners", []string{"signature"}, "List of scanners to be used for Synopsis Detect (formerly BlackDuck) scan.")
	cmd.Flags().StringSliceVar(&stepConfig.ScanPaths, "scanPaths", []string{"."}, "List of paths which should be scanned by the Synopsis Detect (formerly BlackDuck) scan.")
	cmd.Flags().StringSliceVar(&stepConfig.ScanProperties, "scanProperties", []string{"--blackduck.signature.scanner.memory=4096", "--blackduck.timeout=6000", "--blackduck.trust.cert=true", "--detect.policy.check.fail.on.severities=BLOCKER,CRITICAL,MAJOR", "--detect.report.timeout=4800", "--logging.level.com.synopsys.integration=DEBUG"}, "Properties passed to the Synopsis Detect (formerly BlackDuck) scan. You can find details in the [Synopsis Detect documentation](https://synopsys.atlassian.net/wiki/spaces/INTDOCS/pages/622846/Using+Synopsys+Detect+Properties)")
	cmd.Flags().StringVar(&stepConfig.ServerURL, "serverUrl", os.Getenv("PIPER_serverUrl"), "Server url to the Synopsis Detect (formerly BlackDuck) Server.")

	cmd.MarkFlagRequired("apiToken")
	cmd.MarkFlagRequired("projectName")
	cmd.MarkFlagRequired("projectVersion")
}

// retrieve step metadata
func detectExecuteScanMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:    "detectExecuteScan",
			Aliases: []config.Alias{},
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:        "apiToken",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "detect/apiToken"}},
					},
					{
						Name:        "codeLocation",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "projectName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "detect/projectName"}},
					},
					{
						Name:        "projectVersion",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "detect/projectVersion"}},
					},
					{
						Name:        "scanners",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "detect/scanners"}},
					},
					{
						Name:        "scanPaths",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "detect/scanPaths"}},
					},
					{
						Name:        "scanProperties",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "detect/scanProperties"}},
					},
					{
						Name:        "serverUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "detect/serverUrl"}},
					},
				},
			},
		},
	}
	return theMetaData
}
