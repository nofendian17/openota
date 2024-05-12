package airport

import "context"

func (u *useCase) Delete(ctx context.Context, ID string) error {
	return u.airportRepository.Delete(ctx, ID)
}
