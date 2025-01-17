package relayer

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

//nolint:gochecknoglobals // Promauto metrics are global by nature
var (
	bufferLen = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "worker",
		Name:      "buffer_length",
		Help:      "The length of the async send worker activeBuffer per destination chain. Alert if too high",
	}, []string{"chain"})

	mempoolLen = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "worker",
		Name:      "mempool_length",
		Help:      "The length of the mempool per destination chain. Alert if too high",
	}, []string{"chain"})

	workerResets = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "relayer",
		Subsystem: "worker",
		Name:      "reset_total",
		Help:      "The total number of times the worker has reset by destination chain. Alert if too high",
	}, []string{"chain"})

	submissionTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "relayer",
		Subsystem: "worker",
		Name:      "submission_total",
		Help:      "The total number of submissions to destination chain from a specific source chain",
	}, []string{"src_chain", "dst_chain"})

	msgTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "relayer",
		Subsystem: "worker",
		Name:      "msg_total",
		Help:      "The total number of messages submitted to a destination chain from a specific source chain",
	}, []string{"src_chain", "dst_chain"})

	emitCursor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "monitor",
		Name:      "emit_cursor",
		Help:      "The latest emitted cursor on a source chain for a specific destination chain",
	}, []string{"src_chain", "dst_chain"})

	submitCursor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "monitor",
		Name:      "submit_cursor",
		Help:      "The latest submitted cursor on a destination chain for a specific source chain",
	}, []string{"src_chain", "dst_chain"})

	accountBalance = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "monitor",
		Name:      "account_balance_ether",
		Help:      "The balance of the relayer account on a specific chain in ether",
	}, []string{"chain"})

	accountNonce = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "monitor",
		Name:      "account_nonce",
		Help:      "The nonce of the relayer account on a specific chain",
	}, []string{"chain"})
)
