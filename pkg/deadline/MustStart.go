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

func MustStart(ctx context.Context, duration time.Duration, f func(ctx context.Context)) {
	err := Start(ctx, duration, f)
	if err != nil {
		panic(err)
	}
}
