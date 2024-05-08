package country

import (
	"context"
)

func (u *useCase) Delete(ctx context.Context, ID string) error {
	err := u.countryRepository.Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}
