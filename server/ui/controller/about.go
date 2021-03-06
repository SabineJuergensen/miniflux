// Copyright 2017 Frédéric Guillot. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/miniflux/miniflux/server/core"
	"github.com/miniflux/miniflux/version"
)

// AboutPage shows the about page.
func (c *Controller) AboutPage(ctx *core.Context, request *core.Request, response *core.Response) {
	args, err := c.getCommonTemplateArgs(ctx)
	if err != nil {
		response.HTML().ServerError(err)
		return
	}

	response.HTML().Render("about", args.Merge(tplParams{
		"version":    version.Version,
		"build_date": version.BuildDate,
		"menu":       "settings",
	}))
}
