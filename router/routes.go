package router

var routes = Routes{
	"GET": &Route{
		Table: Routes{
			"schemas": &Route{
				Table: Routes{
					"$schema": &Route{
						Table: Routes{
							"archives": &Route{
								Funcs: Routes{
									"year": &Route{Check: IsYear,
										Funcs: Routes{
											"month": &Route{Check: IsMonth,
												Funcs: Routes{
													"day": &Route{HandlerFunc: exampleHandler, Check: IsDay}}}}}}}}}}},
			"testing": &Route{
				Table: Routes{
					"hello": &Route{
						Table: Routes{
							"world": &Route{HandlerFunc: exampleHandler}}}}}},
		Funcs: Routes{
			"day": &Route{Check: IsDay,
				Table: Routes{
					"classes": &Route{
						Table: Routes{
							"go": &Route{HandlerFunc: exampleHandler}}}}}}},
}
