// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package querytest

import (
	"context"
)

type Querier interface {
	User(ctx context.Context) ([]int64, error)
	UsersB(ctx context.Context, id []int64) *UsersBBatchResults
}

var _ Querier = (*Queries)(nil)
