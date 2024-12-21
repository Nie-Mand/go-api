package db

import "errors"

type DB struct {
	DBConfig
}

func New(cfgs ...DBConfig) (*DB, error) {
	if len(cfgs) == 0 {
		cfgs = append(cfgs, NewDefaultDBConfig())
	}

	if len(cfgs) > 1 {
		return nil, errors.New("too many configurations provided")
	}

	return &DB{
		DBConfig: cfgs[0],
	}, nil
}

func (d DB) Close() {
}
