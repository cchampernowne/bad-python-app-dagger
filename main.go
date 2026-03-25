package main

import (
	"context"
	"dagger/ci/internal/dagger"
)

type Ci struct{}

func (m *Ci) Scan(
	ctx context.Context,
	src *dagger.Directory,
	token *dagger.Secret,
) (string, error) {
	return dag.Container().
		From("semgrep/semgrep:latest").
		WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithSecretVariable("SEMGREP_APP_TOKEN", token).
		WithExec([]string{"semgrep", "ci"}).
		Stdout(ctx)
}
