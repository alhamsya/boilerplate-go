package repositories

import "context"

type HelpersRepo interface {
	Tes(ctx context.Context) bool
}
