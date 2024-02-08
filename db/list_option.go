package db

import (
	"github.com/lazygophers/log"
	"github.com/lazygophers/utils/common"
	"strconv"
	"strings"
)

type Options struct {
	Key   uint8  `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

const (
	defaultOffset = 0
	defaultLimit  = 20

	maxLimit = 1000
)

type Page struct {
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`

	Total uint64 `json:"total,omitempty"`
}

type ListOption struct {
	Offset uint64 `json:"offset,omitempty"`
	Limit  uint64 `json:"limit,omitempty"`

	Options []*Options `json:"options,omitempty"`

	ShowTotal bool `json:"show_total,omitempty"`
}

func NewListOption() *ListOption {
	return &ListOption{
		Offset: defaultOffset,
		Limit:  defaultLimit,
	}
}

func (p *ListOption) SetOffset(offset uint64) *ListOption {
	p.Offset = offset
	return p
}

func (p *ListOption) SetLimit(limit uint64) *ListOption {
	p.Limit = limit
	return p
}

func (p *ListOption) SetShowTotal(showTotal ...bool) *ListOption {
	if len(showTotal) > 0 {
		p.ShowTotal = showTotal[0]
	} else {
		p.ShowTotal = true
	}
	return p
}

func (p *ListOption) SetOptions(options ...*Options) *ListOption {
	p.Options = options
	return p
}

func (p *ListOption) AddOptions(options ...*Options) *ListOption {
	p.Options = append(p.Options, options...)
	return p
}

func (p *ListOption) AddOption(key uint8, value string) *ListOption {
	return p.AddOptions(&Options{
		Key:   key,
		Value: value,
	})
}

func (p *ListOption) Clone() *ListOption {
	return &ListOption{
		Offset:    p.Offset,
		Limit:     p.Limit,
		Options:   p.Options,
		ShowTotal: p.ShowTotal,
	}
}

func (p *ListOption) Processor() *ListOptionProcessor {
	return NewListOptionProcessor(p.Clone())
}

func (p *ListOption) Page() *Page {
	return &Page{
		Offset: p.Offset,
		Limit:  p.Limit,
	}
}

type ListOptionProcessor struct {
	ListOption *ListOption

	Handler map[uint8]func(string) error
}

func NewListOptionProcessor(option *ListOption) *ListOptionProcessor {
	return &ListOptionProcessor{
		ListOption: option,
		Handler:    make(map[uint8]func(string) error, len(option.Options)),
	}
}

func (p *ListOptionProcessor) String(key uint8, logic func(value string) error) *ListOptionProcessor {
	p.Handler[key] = logic
	return p
}

func (p *ListOptionProcessor) Int(key uint8, logic func(value int) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		val, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		return logic(val)
	})
}

func (p *ListOptionProcessor) Int8(key uint8, logic func(value int8) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		val, err := strconv.ParseInt(value, 10, 8)
		if err != nil {
			return err
		}
		return logic(int8(val))
	})
}

func (p *ListOptionProcessor) Int16(key uint8, logic func(value int16) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		val, err := strconv.ParseInt(value, 10, 16)
		if err != nil {
			return err
		}
		return logic(int16(val))
	})
}

func (p *ListOptionProcessor) Int32(key uint8, logic func(value int32) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		val, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return err
		}
		return logic(int32(val))
	})
}

func (p *ListOptionProcessor) Int64(key uint8, logic func(value int64) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		val, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		return logic(val)
	})
}

func (p *ListOptionProcessor) Uint(key uint8, logic func(value uint) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		val, err := strconv.ParseUint(value, 10, 0)
		if err != nil {
			return err
		}
		return logic(uint(val))
	})
}

func (p *ListOptionProcessor) Uint8(key uint8, logic func(value uint8) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		val, err := strconv.ParseUint(value, 10, 8)
		if err != nil {
			return err
		}
		return logic(uint8(val))
	})
}

func (p *ListOptionProcessor) Uint16(key uint8, logic func(value uint16) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		val, err := strconv.ParseUint(value, 10, 16)
		if err != nil {
			return err
		}
		return logic(uint16(val))
	})
}

func (p *ListOptionProcessor) Uint32(key uint8, logic func(value uint32) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		val, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return err
		}
		return logic(uint32(val))
	})
}

func (p *ListOptionProcessor) Uint64(key uint8, logic func(value uint64) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		val, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		return logic(val)
	})
}

func (p *ListOptionProcessor) Float32(key uint8, logic func(value float32) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		val, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return err
		}
		return logic(float32(val))
	})
}

func (p *ListOptionProcessor) Float64(key uint8, logic func(value float64) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		return logic(val)
	})
}

func (p *ListOptionProcessor) Bool(key uint8, logic func(value bool) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		switch strings.ToLower(value) {
		case "true", "1", "yes", "y", "on", "enable", "enabled", "ok":
			return logic(true)
		case "false", "0", "no", "n", "off", "disable", "disabled", "cancel":
			return logic(false)
		default:
			return common.ErrInvalid("invalid bool value: " + value)
		}
	})
}

func (p *ListOptionProcessor) Timestamp(key uint8, logic func(value int64) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		val, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		return logic(val)
	})
}

func (p *ListOptionProcessor) TimestampRange(key uint8, logic func(start, end int64) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		timestamps := strings.Split(value, ",")
		if len(timestamps) != 2 {
			return common.ErrInvalid("invalid timestamp value: " + value)
		}

		start, err := strconv.ParseInt(timestamps[0], 10, 64)
		if err != nil {
			return err
		}

		end, err := strconv.ParseInt(timestamps[1], 10, 64)
		if err != nil {
			return err
		}

		return logic(start, end)
	})
}

func (p *ListOptionProcessor) BetweenTimestamp(key uint8, logic func(start, end int64) error) *ListOptionProcessor {
	return p.TimestampRange(key, logic)
}

func (p *ListOptionProcessor) StringSlice(key uint8, logic func(value []string) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		return logic(strings.Split(value, ","))
	})
}

func (p *ListOptionProcessor) IntSlice(key uint8, logic func(value []int) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		var values []int
		for _, val := range strings.Split(value, ",") {
			val, err := strconv.Atoi(val)
			if err != nil {
				return err
			}
			values = append(values, val)
		}
		return logic(values)
	})
}

func (p *ListOptionProcessor) Int8Slice(key uint8, logic func(value []int8) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		var values []int8
		for _, val := range strings.Split(value, ",") {
			val, err := strconv.ParseInt(val, 10, 8)
			if err != nil {
				return err
			}
			values = append(values, int8(val))
		}
		return logic(values)
	})
}

func (p *ListOptionProcessor) Int16Slice(key uint8, logic func(value []int16) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		var values []int16
		for _, val := range strings.Split(value, ",") {
			val, err := strconv.ParseInt(val, 10, 16)
			if err != nil {
				return err
			}
			values = append(values, int16(val))
		}
		return logic(values)
	})
}

func (p *ListOptionProcessor) Int32Slice(key uint8, logic func(value []int32) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		var values []int32
		for _, val := range strings.Split(value, ",") {
			val, err := strconv.ParseInt(val, 10, 32)
			if err != nil {
				return err
			}
			values = append(values, int32(val))
		}
		return logic(values)
	})
}

func (p *ListOptionProcessor) Int64Slice(key uint8, logic func(value []int64) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		var values []int64
		for _, val := range strings.Split(value, ",") {
			val, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return err
			}
			values = append(values, val)
		}
		return logic(values)
	})
}

func (p *ListOptionProcessor) UintSlice(key uint8, logic func(value []uint) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		var values []uint
		for _, val := range strings.Split(value, ",") {
			val, err := strconv.ParseUint(val, 10, 0)
			if err != nil {
				return err
			}
			values = append(values, uint(val))
		}
		return logic(values)
	})
}

func (p *ListOptionProcessor) Uint8Slice(key uint8, logic func(value []uint8) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		var values []uint8
		for _, val := range strings.Split(value, ",") {
			val, err := strconv.ParseUint(val, 10, 8)
			if err != nil {
				return err
			}
			values = append(values, uint8(val))
		}
		return logic(values)
	})
}

func (p *ListOptionProcessor) Uint16Slice(key uint8, logic func(value []uint16) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		var values []uint16
		for _, val := range strings.Split(value, ",") {
			val, err := strconv.ParseUint(val, 10, 16)
			if err != nil {
				return err
			}
			values = append(values, uint16(val))
		}
		return logic(values)
	})
}

func (p *ListOptionProcessor) Uint32Slice(key uint8, logic func(value []uint32) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		var values []uint32
		for _, val := range strings.Split(value, ",") {
			val, err := strconv.ParseUint(val, 10, 32)
			if err != nil {
				return err
			}
			values = append(values, uint32(val))
		}
		return logic(values)
	})
}

func (p *ListOptionProcessor) Uint64Slice(key uint8, logic func(value []uint64) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		var values []uint64
		for _, val := range strings.Split(value, ",") {
			val, err := strconv.ParseUint(val, 10, 64)
			if err != nil {
				return err
			}
			values = append(values, val)
		}
		return logic(values)
	})
}

func (p *ListOptionProcessor) Float32Slice(key uint8, logic func(value []float32) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		var values []float32
		for _, val := range strings.Split(value, ",") {
			val, err := strconv.ParseFloat(val, 32)
			if err != nil {
				return err
			}
			values = append(values, float32(val))
		}
		return logic(values)
	})
}

func (p *ListOptionProcessor) Float64Slice(key uint8, logic func(value []float64) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		var values []float64
		for _, val := range strings.Split(value, ",") {
			val, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return err
			}
			values = append(values, val)
		}
		return logic(values)
	})
}

func (p *ListOptionProcessor) BoolSlice(key uint8, logic func(value []bool) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		var values []bool
		for _, val := range strings.Split(value, ",") {
			switch strings.ToLower(val) {
			case "true", "1", "yes", "y", "on", "enable", "enabled", "ok":
				values = append(values, true)
			case "false", "0", "no", "n", "off", "disable", "disabled", "cancel":
				values = append(values, false)
			default:
				return common.ErrInvalid("invalid bool value: " + val)
			}
		}
		return logic(values)
	})
}

func (p *ListOptionProcessor) Has(key uint8, logic func() error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		return logic()
	})
}

func (p *ListOptionProcessor) Order(key uint8, logic func(value string) error) *ListOptionProcessor {
	return p.String(key, func(value string) error {
		switch strings.ToLower(value) {
		case "desc", "descend", "descending":
			return logic(" desc")
		case "asc", "ascend", "ascending", "":
			fallthrough
		default:
			return logic(" asc")
		}
	})
}

func (p *ListOptionProcessor) Process(tx *Scoop, values any) (*Page, error) {
	for _, option := range p.ListOption.Options {
		if handler, ok := p.Handler[option.Key]; ok {
			err := handler(option.Value)
			if err != nil {
				log.Errorf("err:%v", err)
				return nil, err
			}
		}
	}

	if p.ListOption.Offset < 0 {
		p.ListOption.Offset = defaultOffset
	}

	if p.ListOption.Limit < 0 {
		p.ListOption.Limit = defaultLimit
	} else if p.ListOption.Limit > maxLimit {
		p.ListOption.Limit = maxLimit
	}

	tx = tx.Offset(p.ListOption.Offset).Limit(p.ListOption.Limit)

	page := &Page{
		Offset: p.ListOption.Offset,
		Limit:  p.ListOption.Limit,
	}

	tx.inc()
	defer tx.dec()

	err := tx.Find(values).Error
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	if p.ListOption.ShowTotal {
		page.Total, err = tx.Count()
		if err != nil {
			log.Errorf("err:%v", err)
			return nil, err
		}
	}

	return page, nil
}
