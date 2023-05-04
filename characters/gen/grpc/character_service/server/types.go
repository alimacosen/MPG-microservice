// Code generated by goa v3.11.3, DO NOT EDIT.
//
// CharacterService gRPC server types
//
// Command:
// $ goa gen characters/design

package server

import (
	characterservice "characters/gen/character_service"
	character_servicepb "characters/gen/grpc/character_service/pb"
)

// NewCreateCharacterPayload builds the payload of the "createCharacter"
// endpoint of the "CharacterService" service from the gRPC request type.
func NewCreateCharacterPayload(message *character_servicepb.CreateCharacterRequest) *characterservice.CreateCharacterPayload {
	v := &characterservice.CreateCharacterPayload{
		Name:        message.Name,
		Description: message.Description,
	}
	return v
}

// NewProtoCreateCharacterResponse builds the gRPC response type from the
// result of the "createCharacter" endpoint of the "CharacterService" service.
func NewProtoCreateCharacterResponse(result *characterservice.Character) *character_servicepb.CreateCharacterResponse {
	message := &character_servicepb.CreateCharacterResponse{
		Id:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		Health:      int32(result.Health),
		Experience:  int32(result.Experience),
		InventoryId: result.InventoryID,
	}
	return message
}

// NewGetCharacterPayload builds the payload of the "getCharacter" endpoint of
// the "CharacterService" service from the gRPC request type.
func NewGetCharacterPayload(message *character_servicepb.GetCharacterRequest) *characterservice.GetCharacterPayload {
	v := &characterservice.GetCharacterPayload{
		ID: message.Id,
	}
	return v
}

// NewProtoGetCharacterResponse builds the gRPC response type from the result
// of the "getCharacter" endpoint of the "CharacterService" service.
func NewProtoGetCharacterResponse(result *characterservice.Character) *character_servicepb.GetCharacterResponse {
	message := &character_servicepb.GetCharacterResponse{
		Id:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		Health:      int32(result.Health),
		Experience:  int32(result.Experience),
		InventoryId: result.InventoryID,
	}
	return message
}

// NewUpdateCharacterAttributesPayload builds the payload of the
// "updateCharacterAttributes" endpoint of the "CharacterService" service from
// the gRPC request type.
func NewUpdateCharacterAttributesPayload(message *character_servicepb.UpdateCharacterAttributesRequest) *characterservice.UpdateCharacterAttributesPayload {
	v := &characterservice.UpdateCharacterAttributesPayload{
		ID:          message.Id,
		Name:        message.Name,
		Description: message.Description,
	}
	if message.Health != nil {
		health := int(*message.Health)
		v.Health = &health
	}
	if message.Experience != nil {
		experience := int(*message.Experience)
		v.Experience = &experience
	}
	return v
}

// NewProtoUpdateCharacterAttributesResponse builds the gRPC response type from
// the result of the "updateCharacterAttributes" endpoint of the
// "CharacterService" service.
func NewProtoUpdateCharacterAttributesResponse(result *characterservice.Character) *character_servicepb.UpdateCharacterAttributesResponse {
	message := &character_servicepb.UpdateCharacterAttributesResponse{
		Id:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		Health:      int32(result.Health),
		Experience:  int32(result.Experience),
		InventoryId: result.InventoryID,
	}
	return message
}

// NewDeleteCharacterPayload builds the payload of the "deleteCharacter"
// endpoint of the "CharacterService" service from the gRPC request type.
func NewDeleteCharacterPayload(message *character_servicepb.DeleteCharacterRequest) *characterservice.DeleteCharacterPayload {
	v := &characterservice.DeleteCharacterPayload{
		ID: message.Id,
	}
	return v
}

// NewProtoDeleteCharacterResponse builds the gRPC response type from the
// result of the "deleteCharacter" endpoint of the "CharacterService" service.
func NewProtoDeleteCharacterResponse(result *characterservice.Character) *character_servicepb.DeleteCharacterResponse {
	message := &character_servicepb.DeleteCharacterResponse{
		Id:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		Health:      int32(result.Health),
		Experience:  int32(result.Experience),
		InventoryId: result.InventoryID,
	}
	return message
}
