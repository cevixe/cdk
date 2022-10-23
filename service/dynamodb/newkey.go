package dynamodb

type Key struct {
	PartitionKey *Attribute `field:"required"`
	SortKey      *Attribute `field:"optional"`
}

func NewKey(partitionKey *Attribute, sortKey *Attribute) *Key {
	return &Key{
		PartitionKey: partitionKey,
		SortKey:      sortKey,
	}
}
