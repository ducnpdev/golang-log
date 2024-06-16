package rlogger

type ctxLogger string

var FieldsLog = []ctxLogger{LoggerType, LoggerSpanID, LoggerFlowID, LoggerTranID, LoggerReqID, LoggerAdapterID, LoggerPartnerID, LoggerWorkerID}

const (
	LoggerType      = ctxLogger("type")
	LoggerSpanID    = ctxLogger("spanID")
	LoggerFlowID    = ctxLogger("flowID")
	LoggerTranID    = ctxLogger("tranID")
	LoggerReqID     = ctxLogger("reid")
	LoggerAdapterID = ctxLogger("adapterID")
	LoggerPartnerID = ctxLogger("partnerID")
	LoggerWorkerID  = ctxLogger("workerID")
)
