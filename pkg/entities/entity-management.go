package entities

import (
	"github.com/Zalaxci/space-gopher/pkg/components"
)

type EntityCategory struct {
	Name       string
	Components map[components.ComponentName][]*components.Component
}

func createEntityCategory(
	name string,
	entityComponents map[components.ComponentName]components.Component,
	entityCount uint16,
	modifyEachComponent func(compName components.ComponentName, entityID uint16, comp *components.Component),
) *EntityCategory {
	// Create a new entity category and set its name
	entityCat := EntityCategory{}
	entityCat.Name = name
	// Create a map of ComponentNames to a list (slice) of components for each entity
	entityCat.Components = make(map[components.ComponentName][]*components.Component)
	// Iterate over all provided components and needed entities
	var i uint16
	for compName, comp := range entityComponents {
		for i = 0; i < entityCount; i++ {
			// For each entity, store a copy of the component in a new variable
			currComponent := comp
			// Run the "WhenCreated" function to initialize the component, and modify the component according the function passed as a parameter
			currComponent.WhenCreated()
			modifyEachComponent(compName, i, &currComponent)
			// Finally, append the component to the list
			entityCat.Components[compName] = append(entityCat.Components[compName], &currComponent)
		}
	}
	return &entityCat
}
func (entityCat *EntityCategory) AddEntity(
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
		newComp.WhenCreated()
		// Add it to the list of components for every entity
		entityCat.Components[compName] = append(entityCat.Components[compName], &newComp)
	}
}
func (entityCat *EntityCategory) DestroyAll() {
	var i, entityCount uint16
	for compName := range entityCat.Components {
		entityCount = uint16(len(entityCat.Components[compName]))
		for i = 0; i < entityCount; i++ {
			(*entityCat.Components[compName][i]).WhenDeleted()
		}
		entityCat.Components[compName] = nil
	}
}
