package handle

import (
	"fmt"

	"vieclamit/repository"
)

// Handle struct
type Handle struct {
	Repo repository.Repository
}

// CheckJobDeadlineExpired check expired job deadline
func (h *Handle) CheckJobDeadlineExpired() error {
	count, err := h.Repo.Delete("vieclamit")
	if err != nil {
		return err
	}
	fmt.Printf("check job deadline time expired, removed %d\n", count)
	return nil
}
