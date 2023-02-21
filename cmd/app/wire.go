//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/MrHenri/meuPet/internal/entities"
	"github.com/MrHenri/meuPet/internal/events"
	"github.com/MrHenri/meuPet/internal/infra/database"
	"github.com/MrHenri/meuPet/internal/usecase"
	pkgEvents "github.com/MrHenri/meuPet/pkg/events"
	"github.com/google/wire"
)

var setUserRepositoryDependency = wire.NewSet(
	database.NewUserRepository,
	wire.Bind(new(entities.UserRepositoryInterface), new(*database.UserRepository)),
)

var setUserEventDependency = wire.NewSet(
	events.NewUserEvent,
	wire.Bind(new(pkgEvents.EventInterface), new(*events.UserEvent)),
)

var setEventDispatcherDependency = wire.NewSet(
	pkgEvents.NewEventDispatcher,
	events.NewUserEvent,
	wire.Bind(new(pkgEvents.EventDispatcherInterface), new(*pkgEvents.EventDispatcher)),
	wire.Bind(new(pkgEvents.EventInterface), new(*events.UserEvent)),
)

// TODO: Event
func NewUserUseCase(db *sql.DB /*eventDispatcher pkgEvents.EventDispatcherInterface*/) *usecase.UserUseCase {
	wire.Build(
		setUserRepositoryDependency,
		// setUserEventDependency,
		usecase.NewUserUseCase,
	)
	return &usecase.UserUseCase{}
}
