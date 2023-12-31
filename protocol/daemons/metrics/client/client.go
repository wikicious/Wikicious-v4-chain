package client

import (
	"context"
	"time"

	gometrics "github.com/armon/go-metrics"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/telemetry"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/dydxprotocol/v4-chain/protocol/lib/metrics"
)

var (
	// 30 minutes
	METRICS_DAEMON_LOOP_DELAY_MS       uint32 = 30 * 60 * 1000
	METRICS_DAEMON_LOOP_DELAY_DURATION        = time.Duration(METRICS_DAEMON_LOOP_DELAY_MS)
)

// Start begins a job that periodically:
// 1) Emits metrics about app version and git commit.
func Start(
	ctx context.Context,
	logger log.Logger,
) {
	ticker := time.NewTicker(time.Duration(METRICS_DAEMON_LOOP_DELAY_DURATION * time.Millisecond))
	defer ticker.Stop()
	for ; true; <-ticker.C {
		RunMetricsDaemonTaskLoop(
			ctx,
			logger,
		)
	}
}

// RunMetricsDaemonTaskLoop contains the logic to emit metrics about the application running.
func RunMetricsDaemonTaskLoop(
	ctx context.Context,
	logger log.Logger,
) {
	// Report out app version and git commit.
	version := version.NewInfo()
	telemetry.SetGaugeWithLabels(
		[]string{metrics.AppInfo},
		1,
		[]gometrics.Label{
			metrics.GetLabelForStringValue(metrics.AppVersion, version.Version),
			metrics.GetLabelForStringValue(metrics.GitCommit, version.GitCommit),
		},
	)
	logger.Info(
		"App version",
		metrics.AppVersion,
		version.Version,
		metrics.GitCommit,
		version.GitCommit,
	)
}
