package adapter

import (
	"fmt"
	"github.com/mesh-operator/pkg/adapter/events"
)

// All events comes from adapter needs to be handle by various event handler.
type EventHandler interface {
	// Initializer
	Init()

	// AddService you should handle the event described that a service has been created
	AddService(event events.ServiceEvent, configuratorFinder func(s string) *events.ConfiguratorConfig)
	// DeleteService you should handle the event describe that a service has been removed
	DeleteService(event events.ServiceEvent)
	// AddInstance you should handle the event described that an instance has been registered
	AddInstance(event events.ServiceEvent, configuratorFinder func(s string) *events.ConfiguratorConfig)
	// AddInstance you should handle the event described that an instance has been registered
	ReplaceInstances(event events.ServiceEvent, configuratorFinder func(s string) *events.ConfiguratorConfig)
	// AddInstance you should handle the event describe that an instance has been unregistered
	DeleteInstance(event events.ServiceEvent)

	// AddConfigEntry you should handle the event depicted that a dynamic configuration has been added
	AddConfigEntry(event *events.ConfigEvent, cachedServiceFinder func(s string) *events.Service)
	// ChangeConfigEntry you should handle the event depicted that a dynamic configuration has been changed
	ChangeConfigEntry(event *events.ConfigEvent, cachedServiceFinder func(s string) *events.Service)
	// DeleteConfigEntry you should handle the event depicted that a dynamic configuration has been deleted
	DeleteConfigEntry(event *events.ConfigEvent, cachedServiceFinder func(s string) *events.Service)
}

// SimpleEventHandler Using printing the event's information as a simple handling logic.
type SimpleEventHandler struct {
	Name string
}

func (seh *SimpleEventHandler) Init() {}

func (seh *SimpleEventHandler) AddService(e events.ServiceEvent, configuratorFinder func(s string) *events.ConfiguratorConfig) {
	fmt.Printf("Simple event handler: Adding a service\n%v\n", e.Service)
}

func (seh *SimpleEventHandler) DeleteService(e events.ServiceEvent) {
	fmt.Printf("Simple event handler: Deleting a service\n%v\n", e.Service)
}

func (seh *SimpleEventHandler) AddInstance(e events.ServiceEvent, configuratorFinder func(s string) *events.ConfiguratorConfig) {
	fmt.Printf("Simple event handler: Adding an instance\n%v\n", e.Instance)
}

func (seh *SimpleEventHandler) ReplaceInstances(e events.ServiceEvent, configuratorFinder func(s string) *events.ConfiguratorConfig) {
	fmt.Printf("Simple event handler: Replacing these instances\n%v\n", e.Instances)
}

func (seh *SimpleEventHandler) DeleteInstance(e events.ServiceEvent) {
	fmt.Printf("Simple event handler: Deleting an instance\n%v\n", e.Instance)
}

func (seh *SimpleEventHandler) AddConfigEntry(e *events.ConfigEvent, cachedServiceFinder func(s string) *events.Service) {
	fmt.Printf("Simple event handler: adding a configuration\n%v\n", e.Path)
}

func (seh *SimpleEventHandler) ChangeConfigEntry(e *events.ConfigEvent, cachedServiceFinder func(s string) *events.Service) {
	fmt.Printf("Simple event handler: change a configuration\n%v\n", e.Path)
}

func (seh *SimpleEventHandler) DeleteConfigEntry(e *events.ConfigEvent, cachedServiceFinder func(s string) *events.Service) {
	fmt.Printf("Simple event handler: delete a configuration\n%v\n", e.Path)
}
