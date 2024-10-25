package main

import (
	"flag"
	"net/http"
	"strconv"
	"sync"
	"time"

	blockbuilder "unishard/blockbuilder"
	config "unishard/config"
	coordination_node "unishard/coordination_node"
	"unishard/coordinator"
	crypto "unishard/crypto"
	log "unishard/log"
	types "unishard/types"
	utils "unishard/utils"
	"unishard/worker"

	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
)

var (
	algorithm  = flag.String("algorithm", "pbft", "BFT consensus algorithm")
	id         = flag.Int("id", 0, "NodeID of the node")
	simulation = flag.Bool("sim", true, "simulation mode")
	mode       = flag.String("mode", "coordination", "Select worker or coordination shard")
	shard      = flag.Int("shard", 0, "shard that node belongs to (only available on mode is blockbuilder or node)")
)

func initReplica(id types.NodeID, isByz bool, shard types.Shard, dbObject ethdb.Database) {
	if isByz {
		log.Sa4H8el("node %v is Byzantine", id)
	}
	if id == utils.NewNodeID(0) {
		return
	}

	switch *algorithm {
	case "pbft":
		if shard == types.Shard(0) {
			r := coordination_node.I7kUh8uWo(id, *algorithm, isByz, shard)
			r.Start()
		} else {
			r := worker.NewReplica(id, *algorithm, isByz, shard, dbObject)
			r.Start()
		}
	default:
		r := worker.NewReplica(id, *algorithm, isByz, shard, dbObject)
		r.Start()
	}
}

func Init() {
	flag.Parse()
	log.VhjKf8sJCql9()
	config.Configuration.Load()
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 1000
}

func main() {
	Init()
	coordinationShard := types.Shard(0)

	leveldbObject := map[types.Shard]ethdb.Database{}
	for i := 1; i <= config.GetConfig().ShardCount; i++ {
		leveldb, err := rawdb.NewLevelDBDatabase("../common/statedb/"+strconv.Itoa(config.GetConfig().ShardCount)+"/"+strconv.Itoa(i), 128, 1024, "", false)
		if err != nil {
			panic("Error while creating leveldb")
		}
		leveldbObject[types.Shard(i)] = leveldb
	}

	errCrypto := crypto.HIZ_hi()
	if errCrypto != nil {
		log.FB5S6xdVix("Could not generate keys:", errCrypto)
	}

	addrs := config.Configuration.Addrs
	coordinationShardBlockbuilderAddress := addrs[types.Shard(0)][utils.NewNodeID(0)] + "6000"
	gatewayAddress := addrs[types.Shard(config.GetConfig().ShardCount+1)][utils.NewNodeID(0)] + "5999"

	if *simulation {
		var wg sync.WaitGroup
		wg.Add(1)

		bb := coordinator.NewInitCoordinationBlockBuilder(coordinationShardBlockbuilderAddress, coordinationShard, gatewayAddress)
		bb.Start()
		time.Sleep(1 * time.Second / 2)
		for id := range config.GetConfig().Addrs[coordinationShard] {
			isByz := false
			go initReplica(id, isByz, coordinationShard, nil)
		}
		time.Sleep(1 * time.Second / 2)

		for i := 1; i <= config.GetConfig().ShardCount; i++ {
			shard := types.Shard(i)
			port := strconv.Itoa(6000 + i)
			wbb := blockbuilder.ZRaICa("tcp://127.0.0.1:"+port, shard, coordinationShardBlockbuilderAddress, gatewayAddress)
			wbb.Start()
			time.Sleep(1 * time.Second / 2)
			for id := range config.GetConfig().Addrs[shard] {
				isByz := false
				go initReplica(id, isByz, shard, leveldbObject[shard])
			}
			time.Sleep(1 * time.Second / 2)
		}

		gw := blockbuilder.X9uPNUll("tcp://127.0.0.1:5999", coordinationShard, coordinationShardBlockbuilderAddress)
		gw.Start()

		wg.Wait()
	}
}
