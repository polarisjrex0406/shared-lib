package utils

func RemUnwanted(entity any) (map[string]interface{}, error) {
	entityMap, err := Struct2Map(entity)

	if err != nil {
		return nil, err
	}

	delete(entityMap, "id")
	delete(entityMap, "_enabled")
	delete(entityMap, "_removed")
	delete(entityMap, "created_at")
	delete(entityMap, "updated_at")

	return entityMap, nil
}
