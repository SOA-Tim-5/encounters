package handler

import (
	"database-example/model"
	"database-example/service"
	"strconv"

	events "github.com/SOA-Tim-5/common/common/saga/complete_encounter"
	saga "github.com/SOA-Tim-5/common/common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompleteEncounterCommandHandler struct {
	encounterService  *service.EncounterService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewCompleteEncounterCommandHandler(encounterService *service.EncounterService, publisher saga.Publisher, subscriber saga.Subscriber) (*CompleteEncounterCommandHandler, error) {
	o := &CompleteEncounterCommandHandler{
		encounterService:  encounterService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *CompleteEncounterCommandHandler) handle(command *events.UpdateLevelCommand) {
	id, err := primitive.ObjectIDFromHex(command.UpdateLevel.UserId)
	if err != nil {
		return
	}
	i, _ := strconv.ParseInt(id.String(), 10, 64)
	order := &model.EncounterInstance{Id: i}

	reply := events.CompleteEncounterReply{UpdateLevel: command.UpdateLevel}

	switch command.Type {
	case events.CompleteEncounter:
		reply.Type = events.EncounterCompleted
	case events.RollbackEncounter:
		handler.encounterService.RollbackCompletitionEncounter(order.Id)

		reply.Type = events.EncounterRolledBack
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
