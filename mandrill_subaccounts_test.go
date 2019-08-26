package gochimp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMandrillSubaccounts(t *testing.T) {
	mandrill := testMandrillClient(t)
	r, e := mandrill.SubaccountAdd("test-id", "Test Name", "Test Notes", 0)
	require.NoError(t, e)
	require.Equal(t, "test-id", r.Id)
	require.Equal(t, "Test Name", r.Name)
	_, e = mandrill.SubaccountInfo("test-id")
	require.NoError(t, e)
	_, e = mandrill.SubaccountPause("test-id")
	require.NoError(t, e)
	_, e = mandrill.SubaccountResume("test-id")
	require.NoError(t, e)

	_, e = mandrill.SubaccountDelete("test-id")
	require.NoError(t, e)

	_, e = mandrill.SubaccountInfo("test-id")
	require.NoError(t, e)
}
