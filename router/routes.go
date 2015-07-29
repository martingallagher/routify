package router

var routes = Routes{
	"PUT": &Route{
		Children: Routes{
			"schemas": &Route{
				Child: &Route{
					Param:       "schema",
					HandlerFunc: exampleHandler,
				},
			},
		},
	},
	"DELETE": &Route{
		Children: Routes{
			"schemas": &Route{
				Child: &Route{
					Param:       "schema",
					HandlerFunc: exampleHandler,
				},
			},
		},
	},
	"PATCH": &Route{
		Children: Routes{
			"schemas": &Route{
				Child: &Route{
					Param:       "schema",
					HandlerFunc: exampleHandler,
				},
			},
		},
	},
	"GET": &Route{
		Children: Routes{
			"schemas": &Route{
				Child: &Route{
					Param:       "schema",
					HandlerFunc: exampleHandler,
					Children: Routes{
						"archives": &Route{
							Child: &Route{
								Param: "year",
								Check: IsYear,
								Child: &Route{
									Param: "month",
									Check: IsMonth,
									Child: &Route{
										Param:       "day",
										Check:       IsDay,
										HandlerFunc: exampleHandler,
									},
								},
							},
						},
					},
				},
			},
			"testing/hello/world": &Route{
				HandlerFunc: exampleHandler,
			},
			"/": &Route{
				HandlerFunc: exampleHandler,
			},
			"nofunc": &Route{
				Child: &Route{
					Param: "a",
					Child: &Route{
						Param: "b",
						Child: &Route{
							Param: "c",
							Child: &Route{
								Param: "d",
								Child: &Route{
									Param: "e",
									Child: &Route{
										Param: "f",
										Child: &Route{
											Param: "g",
											Child: &Route{
												Param: "h",
												Child: &Route{
													Param: "i",
													Child: &Route{
														Param: "j",
														Child: &Route{
															Param: "k",
															Child: &Route{
																Param: "l",
																Child: &Route{
																	Param: "m",
																	Child: &Route{
																		Param: "n",
																		Child: &Route{
																			Param: "o",
																			Child: &Route{
																				Param: "p",
																				Child: &Route{
																					Param: "q",
																					Child: &Route{
																						Param: "r",
																						Child: &Route{
																							Param: "s",
																							Child: &Route{
																								Param: "t",
																								Child: &Route{
																									Param:       "u",
																									HandlerFunc: exampleHandler,
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"static/a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t/u": &Route{
				HandlerFunc: exampleHandler,
			},
		},
	},
	"POST": &Route{
		Children: Routes{
			"schemas": &Route{
				Child: &Route{
					Param:       "schema",
					HandlerFunc: exampleHandler,
				},
			},
		},
	},
}
