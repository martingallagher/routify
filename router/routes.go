package router

var routes = Routes{
	"GET": &Route{
		Children: Routes{
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
			"static": &Route{
				Children: Routes{
					"a": &Route{
						Children: Routes{
							"b": &Route{
								Children: Routes{
									"c": &Route{
										Children: Routes{
											"d": &Route{
												Children: Routes{
													"e": &Route{
														Children: Routes{
															"f": &Route{
																Children: Routes{
																	"g": &Route{
																		Children: Routes{
																			"h": &Route{
																				Children: Routes{
																					"i": &Route{
																						Children: Routes{
																							"j": &Route{
																								Children: Routes{
																									"k": &Route{
																										Children: Routes{
																											"l": &Route{
																												Children: Routes{
																													"m": &Route{
																														Children: Routes{
																															"n": &Route{
																																Children: Routes{
																																	"o": &Route{
																																		Children: Routes{
																																			"p": &Route{
																																				Children: Routes{
																																					"q": &Route{
																																						Children: Routes{
																																							"r": &Route{
																																								Children: Routes{
																																									"s": &Route{
																																										Children: Routes{
																																											"t": &Route{
																																												Children: Routes{
																																													"u": &Route{
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
			"schemas": &Route{
				Child: &Route{
					Param: "schema",
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
			"testing": &Route{
				Children: Routes{
					"hello": &Route{
						Children: Routes{
							"world": &Route{
								HandlerFunc: exampleHandler,
							},
						},
					},
				},
			},
		},
	},
}
