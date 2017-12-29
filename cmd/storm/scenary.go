package main

import (
	"log"
	"time"
)

type Scenary struct {
	Name  string
	Steps []Step
}

func NewScenary(name string) (scenary *Scenary) {
	scenary = new(Scenary)
	scenary.Name = name

	return scenary
}

func (scenary *Scenary) AddStep(step Step) error {
	scenary.Steps = append(scenary.Steps, step)
	return nil
}

func (scenary *Scenary) Run(context *Context) error {
	defer context.Config.wg.Done()
	var elapsed time.Duration

	for _, v := range scenary.Steps {
		stepStart := time.Now()
		duration := time.Since(stepStart)
		elapsed += duration
		log.Printf("STEP %s took %s", v.StepId(context), duration)

		if err := v.Run(context); err != nil {
			context.Stats <- Stats{
				Id:       scenary.Name,
				Duration: elapsed,
				Status:   false,
			}

			return err
		}
	}

	context.Stats <- Stats{
		Id:       scenary.Name,
		Duration: elapsed,
		Status:   true,
	}

	return nil
}
