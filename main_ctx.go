package main

// import (
// 	"context"
// 	"crypto/rand"
// 	"encoding/hex"

// 	loggerRus "github.com/ducnpdev/golang-log/logrus"

// 	"github.com/google/uuid"
// )

// func init() {

// }

// var (
// 	lrus loggerRus.Logger
// )

// func GenLogID() string {
// 	bytes := make([]byte, 8)
// 	_, err := rand.Read(bytes)
// 	if err != nil {
// 		return ""
// 	}
// 	return hex.EncodeToString(bytes)
// }

// func main() {
// 	ctx := context.Background()
// 	lrus = loggerRus.New()

// 	partnerId := uuid.NewString()
// 	ctx = context.WithValue(ctx, loggerRus.LoggerReqID, partnerId)
// 	FollowLogic(ctx)
// }

// func FollowLogic(ctx context.Context) {
// 	lrus.DebugWithContext(ctx, "START FollowLogic")
// 	// --
// 	FollowLogic1(ctx)

// 	FollowLogic2(ctx)

// 	lrus.DebugWithContext(ctx, "END FollowLogic")

// }

// // START: FollowLogic1
// func FollowLogic1(ctx context.Context) {
// 	ctx = context.WithValue(ctx, loggerRus.LoggerFlowID, GenLogID())
// 	lrus.DebugWithContext(ctx, "START FollowLogic1")

// 	// usecase1

// 	// adapter 1

// 	lrus.DebugWithContext(ctx, "END FollowLogic1")
// }

// // END:

// // START: FollowLogic2
// func FollowLogic2(ctx context.Context) {
// 	ctx = context.WithValue(ctx, loggerRus.LoggerFlowID, GenLogID())
// 	lrus.DebugWithContext(ctx, "START FollowLogic2")

// 	lrus.DebugWithContext(ctx, "END FollowLogic2")

// }

// // END:

// // START: FollowLogic3
// func FollowLogic3(ctx context.Context) {
// 	ctxNew := context.WithValue(ctx, loggerRus.LoggerFlowID, GenLogID())
// 	lrus.DebugWithContext(ctxNew, "FollowLogic3")
// }

// // END:
