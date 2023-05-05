package handler

import (
	"context"
	"log"
	characterservice "mpg/characters/gen/character_service"
	"mpg/characters/internal/injections"
	"mpg/characters/internal/model"
	inventoryservice "mpg/inventories/gen/inventory_service"
)

type CharacterHandler struct {
	logger *log.Logger
	instances *injections.Instances
}

type CreateInventoryPayload struct {
	UserId string `json:"user_id,omitempty"`
}

type CreateInventoryResult inventoryservice.Inventory

func NewCharacterHandler(logger *log.Logger) *CharacterHandler {
	instances, _ := injections.NewInstances(logger)
	return &CharacterHandler{logger: logger, instances: instances}
}

func (c *CharacterHandler) CreateCharacter(ctx context.Context, p *characterservice.CreateCharacterPayload) (res *characterservice.Character, err error) {
	name := p.Name
	if len(name) == 0 {
		return nil, characterservice.CreateInvalidArgs("Name can not be an empty string")
	}
	description := *p.Description

	svc := c.instances.GetSvc()

	characterPreliminary := &model.Character{
		ID:          "",
		Name:        name,
		Description: description,
		Health:      100,
		Experience:  0,
		InventoryID: "",
	}

	cp, err := svc.Create(ctx, characterPreliminary)
	if err != nil {
		return nil, err
	}

	cp, err = c.createInventory(ctx, cp)
	if err != nil {
		return nil, err
	}

	res = convert(cp)
	return
}

func (c *CharacterHandler) createInventory(ctx context.Context, cp *model.Character) (*model.Character, error) {
	inventoryClient := c.instances.GetInventoryClient()
	createInventoryRpc := inventoryClient.CreateInventory()

	createInventoryRes, err := createInventoryRpc(ctx, &inventoryservice.CreateInventoryPayload{UserID: cp.ID})
	if err != nil {
		return nil, err
	}

	InventoryId := createInventoryRes.(*inventoryservice.Inventory).ID

	svc := c.instances.GetSvc()
	cnt, err := svc.UpdateInventoryId(ctx, cp.ID, InventoryId)
	if err != nil {
		return nil, err
	}

	if cnt != 1 {

	}
	cp.InventoryID = InventoryId

	return cp, nil
}


func (c *CharacterHandler) GetCharacter (ctx context.Context, p *characterservice.GetCharacterPayload) (res *characterservice.Character, err error) {
	id := p.ID
	if len(id) == 0 {
		return nil, characterservice.GetInvalidArgs("Id can not be an empty string")
	}

	svc := c.instances.GetSvc()

	cp, err := svc.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	res = convert(cp)
	return
}

func (c *CharacterHandler) UpdateCharacter(ctx context.Context, p *characterservice.UpdateCharacterPayload) (res int, err error) {
	id := p.ID
	if p.Name != nil && len(*p.Name) == 0 {
		return 0, characterservice.UpdateInvalidArgs("Name can not be an empty string")
	}
	updateFields := assembleUpdateFields(p)

	if updateFields.Health < 0 || updateFields.Experience < 0 {
		return 0, characterservice.UpdateInvalidArgs("health and experience must be non-negative integers")
	}

	svc := c.instances.GetSvc()

	modifiedCnt, err := svc.Update(ctx, id, updateFields)
	if err != nil {
		return 0, err
	}

	res = modifiedCnt
	return
}

func (c *CharacterHandler) DeleteCharacter (ctx context.Context, p *characterservice.DeleteCharacterPayload) (res int, err error) {
	id := p.ID
	if len(id) == 0 {
		return 0, characterservice.DeleteInvalidArgs("Id can not be an empty string")
	}

	svc := c.instances.GetSvc()

	deletedCnt, err := svc.Delete(ctx, id)
	if err != nil {
		return 0, err
	}

	return deletedCnt, nil
}

func convert(c *model.Character) *characterservice.Character {
	return &characterservice.Character{
		ID: c.ID,
		Name: c.Name,
		Description: c.Description,
		Health: c.Health,
		Experience: c.Experience,
		InventoryID: c.InventoryID,
	}
}

func assembleUpdateFields (p *characterservice.UpdateCharacterPayload) *model.UpdateFields {
	updateFields := &model.UpdateFields{}

	if p.Name != nil {
		updateFields.Name = *p.Name
	}

	if p.Description != nil {
		updateFields.Description = *p.Description
	}

	if p.Health != nil {
		updateFields.Health = *p.Health
	}

	if p.Experience != nil {
		updateFields.Experience = *p.Experience
	}

	return updateFields
}