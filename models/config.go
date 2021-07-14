package models

type Config struct {
	dataCenterId     int64
	nodeId           int64
	bitsDataCenterId int64
	bitsNodeId       int64
	bitsSequence     int64
	baseTimestamp    int64
	bufferSize       int
}

func NewConfig(
	dataCenterId int64,
	nodeId int64,
	bitsDataCenterId int64,
	bitsNodeId int64,
	bitsSequence int64,
	baseTimestamp int64,
	bufferSize int,

) *Config {

	if bitsDataCenterId+bitsNodeId+bitsSequence > 64 {

		panic("the sum of bitsDataCenterId,bitsNodeId and bitsSequence cannot be greater than 64")
	}

	return &Config{
		dataCenterId:     dataCenterId,
		nodeId:           nodeId,
		bitsDataCenterId: bitsDataCenterId,
		bitsNodeId:       bitsNodeId,
		bitsSequence:     bitsSequence,
		baseTimestamp:    baseTimestamp,
		bufferSize:       bufferSize,
	}
}
