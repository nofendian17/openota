package state

import "context"

func (u *useCase) Delete(ctx context.Context, ID string) error {
	return u.stateRepository.Delete(ctx, ID)
}
