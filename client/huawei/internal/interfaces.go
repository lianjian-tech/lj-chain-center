/*
 * Copyright (c) Huawei Technologies Co., Ltd. 2021-2021. All rights reserved.
 */

package contractapi

// Log for contract coder
type Log interface {
	// the Debug log interface
	Debug(msg string)
	Debugf(format string, args ...interface{})

	// the Info log interface
	Info(msg string)
	Infof(format string, args ...interface{})

	// the Warn log interface
	Warn(msg string)
	Warnf(format string, args ...interface{})

	// the Error log interface
	Error(msg string)
	Errorf(format string, args ...interface{})
}

// Stub is common contract stub.
type Stub interface {
	Log

	// Get the invoke function name
	FuncName() string

	// Get the invoke parameters
	Parameters() [][]byte

	// Get the tx id for this transaction
	// Deprecated: In the next version, this interface will be deleted.
	TxID() string

	// Get the chain id for this transaction
	ChainID() string

	// Get the contract id for this transaction
	ContractName() string
}

// SQLContract is the SQL smart contract interface .
type SQLContract interface {
	Init(stub SQLContractStub) ([]byte, error)
	Invoke(stub SQLContractStub) ([]byte, error)
}

// SQLContractStub is the SQL smart contract stub .
type SQLContractStub interface {
	Stub
	Query(query string, args ...interface{}) (Rows, error)
	Exec(query string, args ...interface{}) error
}

// Rows is the SQL rows interface .
type Rows interface {
	Next() (bool, error)
	Scan(dest ...interface{}) error
	Columns() ([]string, error)
	Close() error
}

// Contract is the leveldb interface .
type Contract interface {
	Init(stub ContractStub) ([]byte, error)
	Invoke(stub ContractStub) ([]byte, error)
}

// ContractStub is the leveldb stub interface .
type ContractStub interface {
	Stub

	// Get value by key from state DB
	GetKV(key string) ([]byte, error)

	// This action will only generate read write set,
	// the key value will not be put into stateDB until the transaction is validated.
	PutKV(key string, value []byte) error

	// The value should be base type or implement the Marshal(v interface{}) ([]byte, error)
	// and Unmarshal(data []byte, v interface{}) error interface
	PutKVCommon(key string, value interface{}) error

	// This action will only generate read write set, the key value will not be del until the transaction is validated.
	DelKV(key string) error

	// Get a K、V iterator for startKey and endKey. Users do not need to perceive the internal buffering process,
	// even if the result are too large
	// [startKey, endKey) for example: 11--13, you will get 11, 12, not including 13
	// for example: 11--11,you will get nothing
	GetIterator(startKey, endKey string) (Iterator, error)
	// Get history iterator of the key
	GetKeyHistoryIterator(key string) (HistoryIterator, error)

	// 以下两个接口，作为GenerateComKey、DivideComKey、GetComKeyIterator的替代方案
	// Save composite index for objectKey.
	// The mark is the common mark for the object index.
	// The attributes is the attributes for the object value
	// The objectKey is the key for the origin object
	// 直接在此接口内部，完成索引的生成和保存，相当于GenerateComKey和putKV的组合使用
	SaveComIndex(indexName string, attributes []string, objectKey string) error

	// Get the object value by composite index
	// We can get the index by mark and attributes, then we can get object key and value by index
	GetKVByComIndex(indexName string, attributes []string) (Iterator, error)

	// Delete the index for certain object key
	DelComIndexOneRow(indexName string, attributes []string, objectKey string) error
}

// Iterator  interface for user to get key and value one by one
type Iterator interface {

	// is there next k/v of the iterator
	Next() bool

	// read key from iterator
	Key() string

	// read value from iterator
	Value() []byte

	// after read key and value from iterator, the iterator should be closed
	Close()
}

// HistoryIterator  interface for user to get key history one by one
type HistoryIterator interface {
	Iterator

	// Version returns BlockNum & TxNum
	Version() (uint64, int32)

	// TxHash returns transaction hash
	TxHash() []byte

	// IsDeleted returns whether the key has been deleted
	IsDeleted() bool

	// Timestamp returns timestamp
	Timestamp() uint64
}

// ValueSerialization when using "PutKVCommon(key string, value interface{}) error",
// the parameter "value " should be implemented with ValueSerialization interface
type ValueSerialization interface {
	Marshal() ([]byte, error)
}
