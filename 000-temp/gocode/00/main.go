package main

import (
	"fmt"
	"strconv"
)

func main() {
	totalAgents := 5
	agents := make([]Agent,0)
	agents = append(agents, Agent{name: "Dee Fercloze", equipment: "none"})
	agents = append(agents, Agent{name: "Jake Boswell", equipment: "full"})
	agents = append(agents, Agent{name: "Jane Carter", equipment: "full"})
	agents = append(agents, Agent{name: "Max Miller", equipment: "full"})
	agents = append(agents, Agent{name: "Ty Walker", equipment: "full"})

	///code
	fmt.Println(agents[1])

	gearTable := createGearTable(agents, totalAgents)
	fmt.Println(gearTable)
}

func createGearTable(agents []Agent, totalAgents int) string {
	if len(agents) != totalAgents {
		return "ERROR: Wrong number of agents"
	}
	gearTable := "Operation Go: Agent Manifest\n"
	gearTable += "----------------------------"
	for num, agent := range agents {
		gearTable += "\n" + strconv.Itoa(num+1) + ". " + agent.name
		gearTable += ", Gear: " + agent.equipment
	}
	return gearTable
}

type Agent struct {
	name string
	equipment string
}