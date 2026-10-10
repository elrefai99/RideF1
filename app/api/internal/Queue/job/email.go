package JobQueue

import (
	"fmt"

	"github.com/elrefai99/RideF1/app/api/internal/Queue/model"
)

func SendEmail(jobs model.IJob) {
	fmt.Printf(
		"Sending email: %s\n",
		jobs.Payload,
	)
}
