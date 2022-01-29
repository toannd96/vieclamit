package handle

import (
	"fmt"
	"vieclamit/repository"
)

type Handle struct {
	Repo repository.Repository
}

func (h *Handle) CheckJobDeadlineExpired() error {
	count, err := h.Repo.Delete("vieclamit")
	if err != nil {
		return err
	}
	fmt.Printf("check job deadline time expired, removed %d\n", count)
	return nil
}
