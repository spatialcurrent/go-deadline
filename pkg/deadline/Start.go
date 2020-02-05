// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package deadline

import (
	"context"
	"time"
)

func Start(ctx context.Context, duration time.Duration, f func(ctx context.Context)) error {
	d, err := New(duration, f)
	if err != nil {
		return err
	}
	err = d.Start(ctx)
	if err != nil {
		return err
	}
	return nil
}
