package sns

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/jsii-runtime-go"
	"github.com/cevixe/cdk/module"
	"github.com/cevixe/cdk/naming"
)

type Filter = map[string]awssns.SubscriptionFilter

type SubProps struct {
	Topic    awssns.Topic       `field:"required"`
	Function awslambda.Function `field:"required"`
	Filters  *[]*Filter         `field:"optional"`
	Queue    awssqs.Queue       `field:"optional"`
}

func NewSubscriptions(mod module.Module, alias string, props *SubProps) *[]awssns.Subscription {
	if props.Queue == nil {
		return newDirectSubscriptions(mod, alias, props)
	} else {
		return newQueueSubscriptions(mod, alias, props)
	}
}

func newSubscriptionName(mod module.Module, alias string, idx int) naming.Name {

	alias = fmt.Sprintf("%s%02d", alias, idx)
	return naming.NewName(mod, naming.ResType_SNSSubscription, alias)
}

func newDirectSubscriptions(mod module.Module, alias string, props *SubProps) *[]awssns.Subscription {

	subs := make([]awssns.Subscription, 0)

	if props.Filters != nil && len(*props.Filters) > 0 {
		for idx, filter := range *props.Filters {
			name := newSubscriptionName(mod, alias, idx)
			sub := awssns.NewSubscription(
				mod.Resource(),
				name.Logical(),
				&awssns.SubscriptionProps{
					Topic:        props.Topic,
					Endpoint:     props.Function.FunctionArn(),
					Protocol:     awssns.SubscriptionProtocol_LAMBDA,
					FilterPolicy: filter,
				},
			)
			subs = append(subs, sub)
		}
	} else {
		name := newSubscriptionName(mod, alias, 0)
		sub := awssns.NewSubscription(
			mod.Resource(),
			name.Logical(),
			&awssns.SubscriptionProps{
				Topic:    props.Topic,
				Endpoint: props.Function.FunctionArn(),
				Protocol: awssns.SubscriptionProtocol_LAMBDA,
			},
		)
		subs = append(subs, sub)
	}

	return &subs
}

func newQueueSubscriptions(mod module.Module, alias string, props *SubProps) *[]awssns.Subscription {

	source := awslambdaeventsources.NewSqsEventSource(
		props.Queue,
		&awslambdaeventsources.SqsEventSourceProps{
			BatchSize: jsii.Number(1),
			Enabled:   jsii.Bool(true),
		},
	)
	props.Function.AddEventSource(source)

	subs := make([]awssns.Subscription, 0)

	if props.Filters != nil && len(*props.Filters) > 0 {
		for idx, filter := range *props.Filters {
			name := newSubscriptionName(mod, alias, idx)
			sub := awssns.NewSubscription(
				mod.Resource(),
				name.Logical(),
				&awssns.SubscriptionProps{
					Topic:        props.Topic,
					Endpoint:     props.Queue.QueueArn(),
					Protocol:     awssns.SubscriptionProtocol_SQS,
					FilterPolicy: filter,
				},
			)
			subs = append(subs, sub)
		}
	} else {
		name := newSubscriptionName(mod, alias, 0)
		sub := awssns.NewSubscription(
			mod.Resource(),
			name.Logical(),
			&awssns.SubscriptionProps{
				Topic:    props.Topic,
				Endpoint: props.Queue.QueueArn(),
				Protocol: awssns.SubscriptionProtocol_SQS,
			},
		)
		subs = append(subs, sub)
	}

	return &subs
}
