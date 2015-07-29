package website

import "github.com/martingallagher/routify/router"

var Routes = router.Routes{
	"GET": &router.Route{
		Children: router.Routes{
			"/": &router.Route{
				HandlerFunc: index,
			},
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
		},
	},
}
