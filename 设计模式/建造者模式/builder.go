package builder

import (
	"fmt"
)

const (
	MaxToTal = 10
	MaxIdle  = 9
	MinIdle  = 1
)

type ResourcePoolCfg struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

type ResourcePoolCfgBuilder struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

func (this *ResourcePoolCfgBuilder) SetName(name string) error {
	if name == "" {
		return fmt.Errorf("name is empty")
	}
	this.name = name
	return nil
}

func (this *ResourcePoolCfgBuilder) SetMinIdle(minIdle int) error {
	if minIdle < 0 {
		return fmt.Errorf("min idle cannot <= 0, %d", minIdle)
	}
	return nil
}

func (this *ResourcePoolCfgBuilder) SetMaxTotal(maxTotal int) error {
	if maxTotal <= 0 {
		return fmt.Errorf("max tatal cannot <= 0, input: %d", maxTotal)
	}
	this.maxTotal = maxTotal
	return nil
}

func (this *ResourcePoolCfgBuilder) Build() (*ResourcePoolCfg, error) {
	if this.name == "" {
		return nil, fmt.Errorf("name is empty")
	}

	if this.minIdle == 0 {
		this.minIdle = MinIdle
	}

	if this.maxIdle == 0 {
		this.maxIdle = MaxIdle
	}

	if this.maxTotal == 0 {
		this.maxTotal = MaxToTal
	}

	if this.maxTotal < this.maxIdle {
		return nil, fmt.Errorf("max total(%d) cannot < max idle(%d)", this.maxTotal, this.maxIdle)
	}

	if this.minIdle > this.maxIdle {
		return nil, fmt.Errorf("max idle(%d) cannot < min idle(%d)", this.maxIdle, this.minIdle)
	}

	return &ResourcePoolCfg{
		name:     this.name,
		maxTotal: this.maxTotal,
		maxIdle:  this.maxIdle,
		minIdle:  this.minIdle,
	}, nil
}

type ResourcePoolCigOption struct {
	maxTotal int
	maxIdle  int
	minIdle  int
}

type ResourcePoolCfgOptFunc func(option *ResourcePoolCigOption)

func NewResourcePoolCfg(name string, opts ...ResourcePoolCfgOptFunc) (*ResourcePoolCfg, error) {
	if name == "" {
		return nil, fmt.Errorf("name can not empty")
	}

	option := &ResourcePoolCigOption{
		maxTotal: 10,
		maxIdle:  9,
		minIdle:  1,
	}

	for _, opt := range opts {
		opt(option)
	}

	if option.maxTotal < 0 || option.maxIdle < 0 || option.minIdle < 0 {
		return nil, fmt.Errorf("args err, option:%v", option)
	}

	if option.maxTotal < option.maxIdle || option.minIdle > option.maxIdle {
		return nil, fmt.Errorf("args err, option: %v", option)
	}

	return &ResourcePoolCfg{
		name:     name,
		maxTotal: option.maxTotal,
		maxIdle:  option.maxIdle,
		minIdle:  option.minIdle,
	}, nil
}
