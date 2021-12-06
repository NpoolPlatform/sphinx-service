package check

import "github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"

func State(state transaction.Status) bool {
	if state != transaction.StatusPendingReview &&
		state != transaction.StatusConfirm &&
		state != transaction.StatusRejected &&
		state != transaction.StatusPendingTransaction &&
		state != transaction.StatusDone {
		return false
	}

	return true
}
