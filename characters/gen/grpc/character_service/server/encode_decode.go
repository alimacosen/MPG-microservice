// Code generated by goa v3.11.3, DO NOT EDIT.
//
// CharacterService gRPC server encoders and decoders
//
// Command:
// $ goa gen mpg/characters/design

package server

import (
	"context"
	characterservice "mpg/characters/gen/character_service"
	character_servicepb "mpg/characters/gen/grpc/character_service/pb"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeCreateCharacterResponse encodes responses from the "CharacterService"
// service "createCharacter" endpoint.
func EncodeCreateCharacterResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	result, ok := v.(*characterservice.Character)
	if !ok {
		return nil, goagrpc.ErrInvalidType("CharacterService", "createCharacter", "*characterservice.Character", v)
	}
	resp := NewProtoCreateCharacterResponse(result)
	return resp, nil
}

// DecodeCreateCharacterRequest decodes requests sent to "CharacterService"
// service "createCharacter" endpoint.
func DecodeCreateCharacterRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *character_servicepb.CreateCharacterRequest
		ok      bool
	)
	{
		if message, ok = v.(*character_servicepb.CreateCharacterRequest); !ok {
			return nil, goagrpc.ErrInvalidType("CharacterService", "createCharacter", "*character_servicepb.CreateCharacterRequest", v)
		}
	}
	var payload *characterservice.CreateCharacterPayload
	{
		payload = NewCreateCharacterPayload(message)
	}
	return payload, nil
}

// EncodeGetCharacterResponse encodes responses from the "CharacterService"
// service "getCharacter" endpoint.
func EncodeGetCharacterResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	result, ok := v.(*characterservice.Character)
	if !ok {
		return nil, goagrpc.ErrInvalidType("CharacterService", "getCharacter", "*characterservice.Character", v)
	}
	resp := NewProtoGetCharacterResponse(result)
	return resp, nil
}

// DecodeGetCharacterRequest decodes requests sent to "CharacterService"
// service "getCharacter" endpoint.
func DecodeGetCharacterRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *character_servicepb.GetCharacterRequest
		ok      bool
	)
	{
		if message, ok = v.(*character_servicepb.GetCharacterRequest); !ok {
			return nil, goagrpc.ErrInvalidType("CharacterService", "getCharacter", "*character_servicepb.GetCharacterRequest", v)
		}
	}
	var payload *characterservice.GetCharacterPayload
	{
		payload = NewGetCharacterPayload(message)
	}
	return payload, nil
}

// EncodeUpdateCharacterResponse encodes responses from the "CharacterService"
// service "updateCharacter" endpoint.
func EncodeUpdateCharacterResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	result, ok := v.(int)
	if !ok {
		return nil, goagrpc.ErrInvalidType("CharacterService", "updateCharacter", "int", v)
	}
	resp := NewProtoUpdateCharacterResponse(result)
	return resp, nil
}

// DecodeUpdateCharacterRequest decodes requests sent to "CharacterService"
// service "updateCharacter" endpoint.
func DecodeUpdateCharacterRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *character_servicepb.UpdateCharacterRequest
		ok      bool
	)
	{
		if message, ok = v.(*character_servicepb.UpdateCharacterRequest); !ok {
			return nil, goagrpc.ErrInvalidType("CharacterService", "updateCharacter", "*character_servicepb.UpdateCharacterRequest", v)
		}
	}
	var payload *characterservice.UpdateCharacterPayload
	{
		payload = NewUpdateCharacterPayload(message)
	}
	return payload, nil
}

// EncodeDeleteCharacterResponse encodes responses from the "CharacterService"
// service "deleteCharacter" endpoint.
func EncodeDeleteCharacterResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	result, ok := v.(int)
	if !ok {
		return nil, goagrpc.ErrInvalidType("CharacterService", "deleteCharacter", "int", v)
	}
	resp := NewProtoDeleteCharacterResponse(result)
	return resp, nil
}

// DecodeDeleteCharacterRequest decodes requests sent to "CharacterService"
// service "deleteCharacter" endpoint.
func DecodeDeleteCharacterRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *character_servicepb.DeleteCharacterRequest
		ok      bool
	)
	{
		if message, ok = v.(*character_servicepb.DeleteCharacterRequest); !ok {
			return nil, goagrpc.ErrInvalidType("CharacterService", "deleteCharacter", "*character_servicepb.DeleteCharacterRequest", v)
		}
	}
	var payload *characterservice.DeleteCharacterPayload
	{
		payload = NewDeleteCharacterPayload(message)
	}
	return payload, nil
}
