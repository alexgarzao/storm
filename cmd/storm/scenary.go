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
	scenarioStart := time.Now()

	for _, v := range scenary.Steps {
		stepStart := time.Now()
		err := v.Run(context)
		stepDuration := time.Since(stepStart)
		log.Printf("STEP %s took %s", v.StepID(context), stepDuration)

		context.Stats <- Stats{
			ScenarioID: scenary.Name,
			EndpointID: v.EndpointID(context),
			Duration:   stepDuration,
			MustStat:   v.MustStat(context),
			Status:     err == nil,
		}

		// TODO: nao daria para dar um break no if abaixo e setar status se err == nil?
		if err != nil {
			context.Stats <- Stats{
				ScenarioID: scenary.Name,
				EndpointID: "",
				Duration:   time.Since(scenarioStart),
				MustStat:   true,
				Status:     false,
			}

			return err
		}
	}

	context.Stats <- Stats{
		ScenarioID: scenary.Name,
		EndpointID: "",
		Duration:   time.Since(scenarioStart),
		MustStat:   true,
		Status:     true,
	}

	return nil
}
