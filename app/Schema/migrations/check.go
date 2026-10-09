package migrations

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type EnumDefinition struct {
	Name   string
	Values []string
}

var Enums = []EnumDefinition{
	{
		Name: "user_status",
		Values: []string{
			"ACTIVE",
			"INACTIVE",
			"BLOCK",
			"BAN",
		},
	},
	{
		Name: "property_status",
		Values: []string{
			"ACTIVE",
			"PENDING",
			"DELETED",
			"STOP",
		},
	},
}

func EnumCheck(db *gorm.DB) error {
	for _, enum := range Enums {
		var exists bool
		err := db.Raw(
			`
               SELECT EXISTS (
                    SELECT 1 
                    FROM pg_type t 
                    JOIN pg_namespace n ON n.oid = t.typnamespace 
                    WHERE t.typname = ? AND n.nspname = current_schema() AND t.typtype = 'e'
               )`,
			enum.Name,
		).Scan(&exists).Error
		if err != nil {
			return err
		}
		if exists {
			continue
		}

		value := make([]string, 0, len(enum.Values))

		for _, val := range enum.Values {
			value = append(value, "'"+strings.ReplaceAll(val, "'", "''")+"'")
		}
		query := fmt.Sprintf("CREATE TYPE %s AS ENUM (%s)", quot(enum.Name), strings.Join(value, ", "))

		if err := db.Exec(query).Error; err != nil {
			return err
		}
	}
	return nil
}

func quot(name string) string {
	return `"` + strings.ReplaceAll(name, `"`, `""`) + `"`
}
