// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package deadline

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeadline(t *testing.T) {
	x := false
	d, err := New(2*time.Second, func(ctx context.Context) {
		x = true
	})
	require.NoError(t, err)
	require.NotNil(t, d)
	assert.Equal(t, false, d.Started())
	assert.Equal(t, false, d.Cancelled())
	assert.Equal(t, false, d.Finished())
	assert.False(t, x)
	err = d.Start(context.Background())
	require.NoError(t, err)
	assert.Equal(t, true, d.Started())
	assert.Equal(t, false, d.Cancelled())
	assert.Equal(t, false, d.Finished())
	assert.False(t, x)
	time.Sleep(4 * time.Second) // sleep to wait until deadline is finished
	assert.Equal(t, true, d.Started())
	assert.Equal(t, false, d.Cancelled())
	assert.Equal(t, true, d.Finished())
	assert.True(t, x)
	err = d.Start(context.Background())
	require.Error(t, err)
}

func TestDeadlineCancel(t *testing.T) {
	x := false
	d, err := New(4*time.Second, func(ctx context.Context) {
		x = true
	})
	require.NoError(t, err)
	require.NotNil(t, d)
	err = d.Start(context.Background())
	require.NoError(t, err)
	d.Cancel()
	assert.Equal(t, true, d.Started())
	assert.Equal(t, true, d.Cancelled())
	assert.Equal(t, false, d.Finished())
	assert.False(t, x)
	time.Sleep(8 * time.Second)
	assert.Equal(t, true, d.Started())
	assert.Equal(t, true, d.Cancelled())
	assert.Equal(t, true, d.Finished())
	assert.False(t, x)
}

func TestDeadlineInvalidDuration(t *testing.T) {
	d, err := New(0*time.Second, func(ctx context.Context) {})
	require.Error(t, err)
	require.Nil(t, d)
}
