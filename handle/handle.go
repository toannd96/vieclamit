package handle

import (
	"fmt"

	"vieclamit/models"
	"vieclamit/repository"
)

const (
	collection = "vieclamit"
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

// SearchJobByLocation serach all job by location
func (h *Handle) SearchJobByLocation(location string) (*models.Recruitments, error) {
	recruitments, err := h.Repo.FindByLocation(location, collection)
	if err != nil {
		return nil, err
	}
	// for _, recruitment := range *recruitments {
	// 	fmt.Println(recruitment.Location)
	// }
	return recruitments, nil
}

// SearchJobBySkill serach all job by skill
func (h *Handle) SearchJobBySkill(skill string) (*models.Recruitments, error) {
	recruitments, err := h.Repo.FindBySkill(skill, collection)
	if err != nil {
		return nil, err
	}
	// for _, recruitment := range *recruitments {
	// 	fmt.Println(recruitment.Title, recruitment.Company)
	// }
	return recruitments, nil
}

// SearchJobByCompany serach all job by company
func (h *Handle) SearchJobByCompany(company string) (*models.Recruitments, error) {
	recruitments, err := h.Repo.FindByCompany(company, collection)
	if err != nil {
		return nil, err
	}
	// for _, recruitment := range *recruitments {
	// 	fmt.Println(recruitment.Company)
	// }
	return recruitments, nil
}
