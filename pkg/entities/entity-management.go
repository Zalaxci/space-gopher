package entities

import (
	"github.com/Zalaxci/space-gopher/pkg/components"
	"github.com/veandco/go-sdl2/sdl"
)

type EntityCategory struct {
	Name       string
	Components map[components.ComponentName][]components.Component
}

func createEntityCategory(
	renderer *sdl.Renderer,
	name string,
	entityComponents map[components.ComponentName]components.Component,
	entityCount uint16,
	maxEntityCount uint16,
) (entityCat *EntityCategory) {
	// Set the entity's name
	entityCat.Name = name
	var i uint16
	// Iterate over all provided components
	for compName, comp := range entityComponents {
		// Run the "WhenCreated" function to initialize each one
		comp.WhenCreated(renderer)
		// Create a list (slice) of components for each entity. Store the list on the map of components with the component name as the key
		entityCat.Components[compName] = make([]components.Component, 0, maxEntityCount)
		// Add the (same) component for every entity
		for i = 0; i < entityCount; i++ {
			entityCat.Components[compName] = append(entityCat.Components[compName], comp)
		}
	}
	return
}
func (entityCat *EntityCategory) addEntity(
	renderer *sdl.Renderer,
	entityComponents map[components.ComponentName]components.Component,
) {
	// Iterate over all pre-existing components
	for compName := range entityCat.Components {
		// Ensure the corresponding new entity's component was included in the "entityComponents" parameter
		newComp, exists := entityComponents[compName]
		if !exists {
			panic("attempted to add entity to entity list " + entityCat.Name + " without providing all necessary components")
		}
		// Initialize the component
		newComp.WhenCreated(renderer)
		// Add it to the list of components for every entity
		entityCat.Components[compName] = append(entityCat.Components[compName], newComp)
	}
}
