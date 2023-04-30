package lark_event

func IsChallengeStage(tp string) bool {
	return tp == MsgTypeChallengeFlag
}

func NewEventSubscriptionMessageChallengeResponse(challenge string) *EventSubscriptionMessageChallengeResponse {
	return &EventSubscriptionMessageChallengeResponse{
		Challenge: challenge,
	}
}
