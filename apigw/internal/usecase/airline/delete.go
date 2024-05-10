package airline

import "context"

func (u *useCase) Delete(ctx context.Context, ID string) error {
	return u.airlineRepository.Delete(ctx, ID)
}
