package country

import (
	"context"
)

func (u *useCase) Delete(ctx context.Context, ID string) error {
	return u.countryRepository.Delete(ctx, ID)
}
