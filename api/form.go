package api

type ProductForm struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

func (v *ProductForm) Validate() []string {
	errs := []string{}
	if v.ID < 1 {
		errs = append(errs, "ID can't be empty")
	}

	if len(v.Title) < 1 {
		errs = append(errs, "Title can't be empty")
	}

	if len(v.Author) < 1 {
		errs = append(errs, "Athour can't be empty")
	}

	if v.Price < 0 {
		errs = append(errs, "Price can't be empty")
	}

	return errs
}
