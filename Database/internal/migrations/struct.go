package migrations

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
		Name: "driver_status",
		Values: []string{
			"ACCEPT",
			"REJECT",
			"PENDING",
			"BLOCK",
		},
	},
	{
		Name: "email_status",
		Values: []string{
			"SENT",
			"DELIVERED",
			"DELIVERED_DELAYED",
			"FAILED",
		},
	},
	{
		Name: "otp_status",
		Values: []string{
			"SUCCESS",
			"REJECT",
			"PENDING",
			"FAILED",
		},
	},
}
