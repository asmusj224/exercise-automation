package workers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/asmusj224/exercise-automation/database"
	"github.com/asmusj224/exercise-automation/services"
	"github.com/robfig/cron/v3"
)

type Workers interface {
	Start()
}

type workers struct{}

func NewWorkers() Workers {
	return &workers{}
}

func (*workers) Start() {
	store := services.NewStore(database.DB)
	c := cron.New(cron.WithSeconds())
	emailService := services.NewEmailService()
	// 0 0 6 ? * MON-FRI
	c.AddFunc("@every 1m", func() {
		ctx := context.Background()
		workout, _ := store.GetRandomExerciseWorkout(ctx)
		var exercises []services.Exercise
		json.Unmarshal(workout.Exercises, &exercises)
		subject := "Exercise Plan"
		email_body := ""
		for _, exercise := range exercises {
			email_body += exercise.Name + "\n"
		}
		_, err := emailService.SendEmail("jeffrey.asmus88@gmail.com", subject, email_body)
		if err != nil {
			log.Println(err.Error())
		}
	})
	c.Start()
}
