package patterns

import "fmt"

// интерфейс комманды
type Command interface {
	Execute()
}

// приемник
type Light struct{}

func (l *Light) On() {
	fmt.Println("Свет вкл")
}

func (l *Light) Off() {
	fmt.Println("Свет выкл")
}

// конкретная комманда
type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

// конкретная комманды
type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

// вызов
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func main() {
	light := &Light{}
	lightOnCommand := NewLightOnCommand(light)
	lightOffCommand := NewLightOffCommand(light)

	remoteControl := &RemoteControl{}

	remoteControl.SetCommand(lightOnCommand)
	remoteControl.PressButton()

	remoteControl.SetCommand(lightOffCommand)
	remoteControl.PressButton()
}
