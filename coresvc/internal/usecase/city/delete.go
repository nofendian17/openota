package city

import "context"

func (u *useCase) Delete(ctx context.Context, ID string) error {
	return u.cityRepository.Delete(ctx, ID)
}
