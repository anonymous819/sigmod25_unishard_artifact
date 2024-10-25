package measure

import (
	"math"
	"time"
	blockchain "unishard/blockchain"
	config "unishard/config"
	log "unishard/log"
	message "unishard/message"
	types "unishard/types"
	utils "unishard/utils"

	"github.com/ethereum/go-ethereum/common"
)

type PBFTMeasure struct {
	ConflictTransaction            map[common.Hash]types.BlockHeight
	StartTime                      time.Time
	TotalCommittedTx               int
	TotalCommittedLocalTx          int
	TotalCommittedCrossTx          int
	TotalLocalLatency              int64
	TotalWBBWaitingTime            int64
	TotalWorkerConsensusTime       int64
	TotalNetwork1                  int64
	TotalCBBWaitingTime            int64
	TotalCoordinationConsensusTime int64
	TotalNetwork2                  int64
	TotalWBBWaitingTime2           int64
	TotalWorkerConsensusTime2      int64
	TotalLocalWBBWaitingTime       int64
	TotalLocalWorkerConsensusTime  int64
}

func NewPBFTMeasure() *PBFTMeasure {
	return &PBFTMeasure{
		ConflictTransaction:            make(map[common.Hash]types.BlockHeight),
		StartTime:                      time.Time{},
		TotalCommittedTx:               0,
		TotalCommittedLocalTx:          0,
		TotalCommittedCrossTx:          0,
		TotalLocalLatency:              0,
		TotalWBBWaitingTime:            0,
		TotalWorkerConsensusTime:       0,
		TotalNetwork1:                  0,
		TotalCBBWaitingTime:            0,
		TotalCoordinationConsensusTime: 0,
		TotalNetwork2:                  0,
		TotalWBBWaitingTime2:           0,
		TotalWorkerConsensusTime2:      0,
		TotalLocalWBBWaitingTime:       0,
		TotalLocalWorkerConsensusTime:  0,
	}
}

func (measure *PBFTMeasure) CalculateMeasurements(curWorkerBlockBlock *blockchain.WorkerBlock, shard types.Shard) message.Experiment {
	crossShardTransactionLatency := make([]*message.Bncgq6Ys, 0)
	crossLatencyDissection := make([]*message.SBXZ5rXiXHpc, 0)
	localLatencyDissection := make([]*message.SBXZ5rXiXHpc, 0)

	// calculate local latency and latency dissection
	for _, tx := range curWorkerBlockBlock.BlockData.CJpmTwa4y {
		wbbWaitingTime, workerConsensusTime, network1, cbbWaitingTime, coordinationConsensusTime, network2, wbbWaitingTime2, workerConsensusTime2 := CalculateLatencyDissection(tx)
		switch tx.IsCrossShardTx {
		case true:
			if utils.CalculateShardToSend([]common.Address{tx.To}, config.GetConfig().ShardCount)[0] == shard {
				measure.TotalCommittedTx++
				measure.TotalCommittedCrossTx++
				measure.TotalWBBWaitingTime += wbbWaitingTime
				measure.TotalWorkerConsensusTime += workerConsensusTime
				measure.TotalNetwork1 += network1
				measure.TotalCBBWaitingTime += cbbWaitingTime
				measure.TotalCoordinationConsensusTime += coordinationConsensusTime
				measure.TotalNetwork2 += network2
				measure.TotalWBBWaitingTime2 += wbbWaitingTime2
				measure.TotalWorkerConsensusTime2 += workerConsensusTime2
			}
			crossShardTransactionLatency = append(crossShardTransactionLatency, &message.Bncgq6Ys{
				W4puDs:      tx.Hash,
				DESc6onmvvK: time.Now().UnixMilli() - tx.Timestamp,
			})
			crossLatencyDissection = append(crossLatencyDissection, &message.SBXZ5rXiXHpc{
				JR5HajI:  tx.Hash,
				UCK4axEV: []int64{wbbWaitingTime, workerConsensusTime, network1, cbbWaitingTime, coordinationConsensusTime, network2, wbbWaitingTime2, workerConsensusTime2},
			})
		case false:
			measure.TotalCommittedTx++
			measure.TotalCommittedLocalTx++
			measure.TotalLocalWBBWaitingTime += wbbWaitingTime
			measure.TotalLocalWorkerConsensusTime += workerConsensusTime
			latency := time.Now().UnixMilli() - tx.Timestamp
			measure.TotalLocalLatency += latency
			localLatencyDissection = append(localLatencyDissection, &message.SBXZ5rXiXHpc{
				JR5HajI:  tx.Hash,
				UCK4axEV: []int64{wbbWaitingTime, workerConsensusTime},
			})
		}
	}

	// calculate TPS
	totalTPS := math.Round((float64(measure.TotalCommittedTx) / time.Since(measure.StartTime).Seconds()))
	localTPS := math.Round((float64(measure.TotalCommittedLocalTx) / time.Since(measure.StartTime).Seconds()))
	crossTPS := math.Round((float64(measure.TotalCommittedCrossTx) / time.Since(measure.StartTime).Seconds()))
	totalConflict := len(measure.ConflictTransaction)

	// calculate average latency
	avgLocalLatency := CalculateAverageTimeDifference(measure.TotalLocalLatency, measure.TotalCommittedLocalTx)
	avgWBBWaitingTime := CalculateAverageTimeDifference(measure.TotalWBBWaitingTime, measure.TotalCommittedCrossTx)
	avgWorkerConsensusTime := CalculateAverageTimeDifference(measure.TotalWorkerConsensusTime, measure.TotalCommittedCrossTx)
	avgNetwork1 := CalculateAverageTimeDifference(measure.TotalNetwork1, measure.TotalCommittedCrossTx)
	avgCBBWaitingTime := CalculateAverageTimeDifference(measure.TotalCBBWaitingTime, measure.TotalCommittedCrossTx)
	avgCoordinationConsensusTime := CalculateAverageTimeDifference(measure.TotalCoordinationConsensusTime, measure.TotalCommittedCrossTx)
	avgNetwork2 := CalculateAverageTimeDifference(measure.TotalNetwork2, measure.TotalCommittedCrossTx)
	avgWBBWaitingTime2 := CalculateAverageTimeDifference(measure.TotalWBBWaitingTime2, measure.TotalCommittedCrossTx)
	avgWorkerConsensusTime2 := CalculateAverageTimeDifference(measure.TotalWorkerConsensusTime2, measure.TotalCommittedCrossTx)
	avgLocalWBBWaitingTime := CalculateAverageTimeDifference(measure.TotalLocalWBBWaitingTime, measure.TotalCommittedLocalTx)
	avgLocalWorkerConsensusTime := CalculateAverageTimeDifference(measure.TotalLocalWorkerConsensusTime, measure.TotalCommittedLocalTx)

	consensusForISC := avgWorkerConsensusTime + avgCoordinationConsensusTime
	ISC := avgNetwork1 + avgNetwork2
	consensusForCommit := avgWorkerConsensusTime2
	waitingTime := avgWBBWaitingTime + avgCBBWaitingTime + avgWBBWaitingTime2
	total := consensusForISC + ISC + consensusForCommit + waitingTime

	log.B4DhS7wNt0JY("[%v] total TPS: %v, local TPS: %v, cross TPS: %v, local latency: %v, total conflict tx: %v, total committed tx: %v, committed tx num: %v", shard, totalTPS, localTPS, crossTPS, avgLocalLatency, totalConflict, measure.TotalCommittedTx, len(curWorkerBlockBlock.BlockData.CJpmTwa4y))
	log.B4DhS7wNt0JY("[%v] [Cross-shard latency dissection] Consensus for communication: %.3f, Communication: %.3f, Consensus for execution and commit: %.3f, Queueing and waiting time: %.3f, Total: %.3f", shard, consensusForISC, ISC, consensusForCommit, waitingTime, total)
	log.B4DhS7wNt0JY("[%v] [Local latency dissection] Consensus for execution and commit: %.3f, Queueing and waiting time: %.3f, Total: %.3f", shard, avgLocalWorkerConsensusTime, avgLocalWBBWaitingTime, avgLocalWBBWaitingTime+avgLocalWorkerConsensusTime)

	experiment := message.Experiment{
		Shard:        shard,
		Bkq028NruW:   crossShardTransactionLatency,
		RoQbttkP:     avgLocalWBBWaitingTime + avgLocalWorkerConsensusTime,
		W2g99STnJRA1: crossLatencyDissection,
		JUgObcOr:     localLatencyDissection,
	}

	calDuration := config.GetConfig().Benchmark.N / config.GetConfig().Benchmark.Throttle
	if !measure.StartTime.IsZero() && time.Since(measure.StartTime).Seconds() > float64(calDuration)-10 {
		experimentTransactionResult := message.R8SUVOg{
			RRTdrUMYO: measure.TotalCommittedTx,
			Zjb9_abe:  measure.TotalCommittedLocalTx,
			GV9IpLO9e: measure.TotalCommittedCrossTx,
			XRLlpce9L: time.Since(measure.StartTime).Seconds(),
		}
		experiment.Da8uC0Sys2u = experimentTransactionResult
		experiment.E2v01O_Ii = true
	}

	return experiment
}

func CalculateAverageTimeDifference(total int64, num int) float64 {
	return math.Round((float64(total)/1000)/float64(num)*1000) / 1000
}

func CalculateLatencyDissection(tx *message.Transaction) (wbbWaitingTime, workerConsensusTime, network1, cbbWaitingTime, coordinationConsensusTime, network2, wbbWaitingTime2, workerConsensusTime2 int64) {
	return tx.LatencyDissection.WBBWaitingTime - tx.Timestamp,
		tx.LatencyDissection.WorkerConsensusTime - tx.LatencyDissection.WBBWaitingTime,
		tx.LatencyDissection.Network1 - tx.LatencyDissection.WorkerConsensusTime,
		tx.LatencyDissection.CBBWaitingTIme - tx.LatencyDissection.Network1,
		tx.LatencyDissection.CoordinationConsensusTime - tx.LatencyDissection.CBBWaitingTIme,
		tx.LatencyDissection.Network2 - tx.LatencyDissection.CoordinationConsensusTime,
		tx.LatencyDissection.WBBWaitingTime2 - tx.LatencyDissection.Network2,
		tx.LatencyDissection.WorkerConsensusTime2 - tx.LatencyDissection.WBBWaitingTime2
}
