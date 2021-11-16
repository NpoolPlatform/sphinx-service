package approval

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

// 审核通过交易
func ApproveTransaction(transactionIDInsite string) (isApproved bool, err error) {
	ctx := context.Background()
	// 获取交易信息
	info, err := db.Client().Transaction.Query().
		Where(
			transaction.And(
				transaction.TransactionIDChain(transactionIDInsite),
			),
		).First(ctx)
	if err != nil {
		err = status.Error(codes.Internal, "approval db query failed")
		return
	}
	// 检查交易状态
	if info.Status != transaction.StatusPendingReview {
		if info.Status != transaction.StatusPendingProcess {
			err = status.Error(codes.AlreadyExists, "transaction already processed, current status: "+info.Status.String())
			return
		}
		logger.Sugar().Warn("repeated approval request, transactionIDInsite: " + transactionIDInsite)
	} else if info.Mutex {
		// return
		logger.Sugar().Warn("approval request already in process, transactionIDInsite: " + transactionIDInsite)
	}
	// 加锁
	info, err = info.Update().
		SetMutex(true).
		Save(ctx)
	if err != nil {
		err = status.Error(codes.Internal, "approval update db failed")
		return
	}
	// 进行审核
	approvalAllPassed := true
	if approvalAllPassed {
		// 审核通过
		_, err = info.Update().
			SetMutex(false).
			SetStatus(transaction.StatusPendingProcess).
			SetUpdatetimeUtc(int(time.Now().UTC().Unix())).
			Save(ctx)
		if err != nil {
			err = status.Error(codes.Internal, "approval update db failed, mutex still locked")
			return
		}
		isApproved = true
	} else {
		// 审核拒绝
		_, err = info.Update().
			SetMutex(false).
			SetStatus(transaction.StatusRejected).
			SetUpdatetimeUtc(int(time.Now().UTC().Unix())).
			Save(ctx)
		if err != nil {
			err = status.Error(codes.Internal, "approval update db failed, mutex still locked")
			return
		}
	}
	// nothing else, so collect error and return here
	return isApproved, err
}
