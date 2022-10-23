package bus

import (
	"log"

	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/service/sns"
)

type BusProps struct {
	Type BusType `field:"required"`
}

func NewBus(mod module.Module, alias string, props *BusProps) Bus {
	switch props.Type {
	case BusType_Standard:
		return newStandardBus(mod, alias, props)
	case BusType_Advanced:
		return newAdvancedBus(mod, alias, props)
	default:
		log.Fatal("unknown bus type")
	}
	return nil
}

func newStandardBus(mod module.Module, alias string, props *BusProps) Bus {

	topic := sns.NewTopic(
		mod,
		alias,
		&sns.TopicProps{Type: sns.TopicType_Standard},
	)

	return &busImpl{
		module:   mod,
		name:     alias,
		typ:      props.Type,
		resource: topic,
	}
}

func newAdvancedBus(mod module.Module, alias string, props *BusProps) Bus {

	topic := sns.NewTopic(
		mod,
		alias,
		&sns.TopicProps{Type: sns.TopicType_FIFO},
	)

	return &busImpl{
		module:   mod,
		name:     alias,
		typ:      props.Type,
		resource: topic,
	}
}
