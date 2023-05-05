package model

import (
	characterservice "mpg/characters/gen/character_service"
)

type InventoryFields struct {
	InventoryId string `bson:"inventoryid,omitempty"`
}

type UpdateFields struct {
	Name        string `bson:"name,omitempty"`
	Description string `bson:"description,omitempty"`
	Health     int     `bson:"health,omitempty"`
	Experience int     `bson:"experience,omitempty"`
}

type Character characterservice.Character

func (c *Character) GetId() string {
	return c.ID
}

func (c *Character) SetId(id string)  {
	c.ID = id
}

func (c *Character) GetName() string {
	return c.Name
}

func (c *Character) SetName(name string)  {
	c.Name = name
}

func (c *Character) GetDescription() string {
	return c.Description
}

func (c *Character) SetDescription(description string)  {
	c.Description = description
}

func (c *Character) GetHealth() int {
	return c.Health
}

func (c *Character) SetHealth(health int)  {
	c.Health = health
}