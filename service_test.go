/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*

	Test service start/stop here

*/

package iconfigcon

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	intf "github.com/untillpro/airs-istorage"
	"github.com/untillpro/godif/services"
)

func Test_StartStop(t *testing.T) {
	ctx, err := setUp(t)
	defer tearDown(ctx, t)
	require.Nil(t, err, err)

	/*
		Your tests here
	*/
}

func setUp(t *testing.T) (context.Context, error) {

	// Declare own service
	Declare(Service{Host: "localhost", Port: 800})

	intf.DeclareForTest()

	return services.ResolveAndStart()
}

func tearDown(ctx context.Context, t *testing.T) {
	services.StopAndReset(ctx)
}
