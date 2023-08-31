package netxd_dal_interfaces


type TransactionInterface interface {
	TransferMoney(from int64, to int64, amount int64)(string, error)
}