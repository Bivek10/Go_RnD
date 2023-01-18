package repository

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewUserRepository),
	fx.Provide(NewPlantRepository),
	fx.Provide(NewQuizRepository),
	fx.Provide(NewQuestionRepository),
)
