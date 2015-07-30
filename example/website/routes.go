package website

import "github.com/martingallagher/routify/router"

var Routes = &router.Router{
	Routes: router.Routes{
		"GET": &router.Route{
			Children: router.Routes{
				"hello": &router.Route{
					Child: &router.Route{
						Param:       "str",
						HandlerFunc: hello,
					},
				},
				"printnum": &router.Route{
					Child: &router.Route{
						Param:       "num",
						Check:       validateNumber,
						HandlerFunc: printnum,
					},
				},
				"/": &router.Route{
					HandlerFunc: index,
				},
			},
		},
	},
	Validators: router.Validators{
		"$num": validateNumber,
	},
}
