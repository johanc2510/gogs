package repo

import (
	"gogs.io/gogs/internal/context"
)

const (
	REPO_EXPLOSION = "repo/explosion"
)

func CheckExplosion() {

}

func Explosion(c *context.Context) {
	c.Data["PageIsExplosion"] = true

	c.Success(REPO_EXPLOSION)
}
