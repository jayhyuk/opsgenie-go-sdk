package alertsv2

type Report struct {
	AckTime        int64 `json:"ackTime,omitempty"`
	CloseTime      int64 `json:"closeTime,omitempty"`
	AcknowledgedBy string `json:"acknowledgedBy,omitempty"`
	ClosedBy       string `json:"closedBy,omitempty"`
}
