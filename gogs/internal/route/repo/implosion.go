package repo

import (
	"gogs.io/gogs/internal/context"
)

const (
	REPO_IMPLOSION = "repo/implosion"
)

func CheckImplosion() {

}

func Implosion(c *context.Context) {
	c.Data["PageIsImplosion"] = true

	c.Success(REPO_IMPLOSION)
}
