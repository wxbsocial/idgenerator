package models

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const POSITIVE_SIGN_MASK = int64(^uint64(0) >> 1)

type Generator struct {
	*Config
	TimeGetter
	offsetTimestamp    int64
	offsetDataCenterId int64
	offsetNodeId       int64
	signMask           int64
	maxSequence        int64
	buffer             chan int64
}

func NewGenerator(config *Config, timeGetter TimeGetter) *Generator {

	var (
		offsetNodeId       = config.bitsSequence
		offsetDataCenterId = offsetNodeId + config.bitsNodeId
		offsetTimestamp    = offsetDataCenterId + config.bitsDataCenterId
		maxSequence        = int64(1) << config.bitsSequence
	)

	generator := Generator{
		Config:             config,
		TimeGetter:         timeGetter,
		offsetNodeId:       offsetNodeId,
		offsetDataCenterId: offsetDataCenterId,
		offsetTimestamp:    offsetTimestamp,
		maxSequence:        maxSequence,
		signMask:           POSITIVE_SIGN_MASK,
		buffer:             make(chan int64, config.bufferSize),
	}

	// termChan := make(chan os.Signal)
	// signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	// ctx, cancelFunc := context.WithCancel(context.Background())

	go generator.produce()
	// <-termChan

	return &generator

}

func (gen *Generator) Get() (int64, error) {
	value, ok := <-gen.buffer
	if !ok {
		return 0, errors.New("buffer is closed")
	}

	return value, nil
}

func (gen *Generator) produce() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	latestId := int64(0)

	for {

		id := gen.genIdWithoutSeq(gen.Now())

		if id == latestId {
			continue
		}

		for i := int64(0); i < gen.maxSequence; i++ {

			select {
			case gen.buffer <- (id | i):
			case <-sigs:
				{
					log.Println("exit beause receive term sigs")
					os.Exit(0)
				}
			}

		}

		latestId = id

	}
}

func (gen *Generator) genIdWithoutSeq(nowTs int64) int64 {

	nowTs -= gen.baseTimestamp

	id := nowTs<<gen.offsetTimestamp | gen.dataCenterId<<gen.offsetDataCenterId | gen.nodeId<<gen.offsetNodeId

	return id & gen.signMask
}
