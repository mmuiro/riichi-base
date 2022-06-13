package waits

type WaitKind int

const (
	Ryanmen WaitKind = iota
	Kanchan
	Penchan
	Shanpon
	Tanki
	KokushiSingle
	KokushiThirteen
	JunseiChuuren
)

/* although Chuuren Poutou fits into standard hand partitions,
the Junsei variant depends on all the waits of the hand;
hence, we give it a unique wait to indicate that. */
