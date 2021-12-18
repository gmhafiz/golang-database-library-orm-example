package sqlboiler

import (
	"context"
	"fmt"
)

func (r *database) ListM2M(ctx context.Context) (interface{}, error) {
	return nil, fmt.Errorf("not implemented. See https://github.com/volatiletech/sqlboiler/issues/756#issuecomment-663869392")
}
