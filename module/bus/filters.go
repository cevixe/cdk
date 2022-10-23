package bus

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/jsii-runtime-go"
)

type Filter = map[string]awssns.SubscriptionFilter

func NewFilter(kind string, types ...string) *Filter {

	filter := Filter{
		"kind": awssns.SubscriptionFilter_StringFilter(&awssns.StringConditions{
			Allowlist: &[]*string{jsii.String(kind)},
		}),
	}

	if len(types) > 0 {
		allowedItems := make([]*string, 0)
		for _, item := range types {
			allowedItems = append(allowedItems, jsii.String(item))
		}
		subfilter := awssns.SubscriptionFilter_StringFilter(&awssns.StringConditions{
			Allowlist: &allowedItems,
		})
		filter["type"] = subfilter
	}

	return &filter
}
