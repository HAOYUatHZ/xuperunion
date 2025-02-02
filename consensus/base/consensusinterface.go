package base

import (
	"github.com/xuperchain/log15"
	"github.com/xuperchain/xuperunion/common/config"
	"github.com/xuperchain/xuperunion/pb"
)

// ConsensusInterface is the interface of consensus
type ConsensusInterface interface {
	Type() string
	Version() int64
	// 用于回滚或者重启时一些临时数据的恢复
	InitCurrent(block *pb.InternalBlock) error
	Configure(xlog log.Logger, cfg *config.NodeConfig, consCfg map[string]interface{},
		extParams map[string]interface{}) error
	// CompeteMaster 返回是否为矿工以及是否需要进行SyncBlock
	CompeteMaster(height int64) (bool, bool)
	CheckMinerMatch(header *pb.Header, in *pb.InternalBlock) (bool, error)
	// 开始挖矿前进行相应的处理
	ProcessBeforeMiner(timestamp int64) (map[string]interface{}, bool)
	// 用于确认块后进行相应的处理
	ProcessConfirmBlock(block *pb.InternalBlock) error

	// Get current core miner info
	GetCoreMiners() []*MinerInfo
}
