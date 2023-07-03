package constants

type Order string

const (
	ASC  Order = "ASC"
	DESC Order = "DESC"
)

var ValidOrders = map[string]struct{}{
	string(ASC):  {},
	string(DESC): {},
}
