package validatecmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	internalcmd "package-operator.run/internal/cmd"
)

func TestValidateFolder(t *testing.T) {
	t.Parallel()

	scheme, err := internalcmd.NewScheme()
	require.NoError(t, err)

	cmd := NewCmd(internalcmd.NewValidate(scheme))
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.SetOut(stdout)
	cmd.SetErr(stderr)
	cmd.SetArgs([]string{"testdata"})

	require.NoError(t, cmd.Execute())
	require.Empty(t, stdout.String())
	require.Empty(t, stderr.String())
}
