package service

import (
	"database-example/model"
	"strconv"

	events "github.com/SOA-Tim-5/common/common/saga/complete_encounter"
	saga "github.com/SOA-Tim-5/common/common/saga/messaging"
)

type CompleteEncounterOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewCompleteEncounterOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*CompleteEncounterOrchestrator, error) {
	o := &CompleteEncounterOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func Start(o *CompleteEncounterOrchestrator, order *model.TouristProgress, encounterId int64) error {
	event := &events.UpdateLevelCommand{
		Type: events.UpdateFollower,
		UpdateLevel: events.UpdateLevel{
			UserId:      strconv.FormatInt(order.UserId, 10),
			Level:       strconv.FormatInt(int64(order.Level), 10),
			EncounterId: encounterId,
		},
	}
	println("start of saga")
	println(event.Type)
	println("level" + event.UpdateLevel.Level)
	return o.commandPublisher.Publish(event)
}

func (o *CompleteEncounterOrchestrator) handle(reply *events.CompleteEncounterReply) {
	command := events.UpdateLevelCommand{UpdateLevel: reply.UpdateLevel}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *CompleteEncounterOrchestrator) nextCommandType(reply events.CompleteEncounterReplyType) events.CompleteEncounterCommandType {
	switch reply {
	case events.FollowerNotUpdated:
		return events.RollbackEncounter
	default:
		return events.UnknownCommand
	}
}
