package twmerge

func getBreaks(groupId string) map[string]ClassPart {
	return map[string]ClassPart{
		"auto": ClassPart{
			NextPart:     map[string]ClassPart{},
			Validators:   []ClassGroupValidator{},
			ClassGroupId: groupId,
		},
		"avoid": ClassPart{
			NextPart:     make(map[string]ClassPart),
			Validators:   []ClassGroupValidator{},
			ClassGroupId: groupId,
		},
		"all": ClassPart{
			NextPart:     map[string]ClassPart{},
			Validators:   []ClassGroupValidator{},
			ClassGroupId: groupId,
		},
		"page": ClassPart{
			NextPart:     map[string]ClassPart{},
			Validators:   []ClassGroupValidator{},
			ClassGroupId: groupId,
		},
		"left": ClassPart{
			NextPart:     map[string]ClassPart{},
			Validators:   []ClassGroupValidator{},
			ClassGroupId: groupId,
		},
		"right": ClassPart{
			NextPart:     map[string]ClassPart{},
			Validators:   []ClassGroupValidator{},
			ClassGroupId: groupId,
		},
		"column": ClassPart{
			NextPart:     map[string]ClassPart{},
			Validators:   []ClassGroupValidator{},
			ClassGroupId: groupId,
		},
	}
}

// This is horrible code. I'm sorry. I wanted to get the package working without writing the code to generate the config. Now that it is working I plan to writing it.
func MakeDefaultConfig() *TwMergeConfig {
	return &TwMergeConfig{
		ModifierSeparator: ':',
		ClassSeparator:    '-',
		ImportantModifier: '!',
		PostfixModifier:   '/',
		MaxCacheSize:      1000,
		// Prefix:            "",
		// theme:             TwTheme{},
		ConflictingClassGroups: ConflictingClassGroups{
			"overflow":         {"overflow-x", "overflow-y"},
			"overscroll":       {"overscroll-x", "overscroll-y"},
			"inset":            {"inset-x", "inset-y", "start", "end", "top", "right", "bottom", "left"},
			"inset-x":          {"right", "left"},
			"inset-y":          {"top", "bottom"},
			"flex":             {"basis", "grow", "shrink"},
			"gap":              {"gap-x", "gap-y"},
			"p":                {"px", "py", "ps", "pe", "pt", "pr", "pb", "pl"},
			"px":               {"pr", "pl"},
			"py":               {"pt", "pb"},
			"m":                {"mx", "my", "ms", "me", "mt", "mr", "mb", "ml"},
			"mx":               {"mr", "ml"},
			"my":               {"mt", "mb"},
			"size":             {"w", "h"},
			"font-size":        {"leading"},
			"fvn-normal":       {"fvn-ordinal", "fvn-slashed-zero", "fvn-figure", "fvn-spacing", "fvn-fraction"},
			"fvn-ordinal":      {"fvn-normal"},
			"fvn-slashed-zero": {"fvn-normal"},
			"fvn-figure":       {"fvn-normal"},
			"fvn-spacing":      {"fvn-normal"},
			"fvn-fraction":     {"fvn-normal"},
			"line-clamp":       {"display", "overflow"},
			"rounded":          {"rounded-s", "rounded-e", "rounded-t", "rounded-r", "rounded-b", "rounded-l", "rounded-ss", "rounded-se", "rounded-ee", "rounded-es", "rounded-tl", "rounded-tr", "rounded-br", "rounded-bl"},
			"rounded-s":        {"rounded-ss", "rounded-es"},
			"rounded-e":        {"rounded-se", "rounded-ee"},
			"rounded-t":        {"rounded-tl", "rounded-tr"},
			"rounded-r":        {"rounded-tr", "rounded-br"},
			"rounded-b":        {"rounded-br", "rounded-bl"},
			"rounded-l":        {"rounded-tl", "rounded-bl"},
			"border-spacing":   {"border-spacing-x", "border-spacing-y"},
			"border-w":         {"border-w-s", "border-w-e", "border-w-t", "border-w-r", "border-w-b", "border-w-l"},
			"border-w-x":       {"border-w-r", "border-w-l"},
			"border-w-y":       {"border-w-t", "border-w-b"},
			"border-color":     {"border-color-t", "border-color-r", "border-color-b", "border-color-l"},
			"border-color-x":   {"border-color-r", "border-color-l"},
			"border-color-y":   {"border-color-t", "border-color-b"},
			"scroll-m":         {"scroll-mx", "scroll-my", "scroll-ms", "scroll-me", "scroll-mt", "scroll-mr", "scroll-mb", "scroll-ml"},
			"scroll-mx":        {"scroll-mr", "scroll-ml"},
			"scroll-my":        {"scroll-mt", "scroll-mb"},
			"scroll-p":         {"scroll-px", "scroll-py", "scroll-ps", "scroll-pe", "scroll-pt", "scroll-pr", "scroll-pb", "scroll-pl"},
			"scroll-px":        {"scroll-pr", "scroll-pl"},
			"scroll-py":        {"scroll-pt", "scroll-pb"},
			"touch":            {"touch-x", "touch-y", "touch-pz"},
			"touch-x":          {"touch"},
			"touch-y":          {"touch"},
			"touch-pz":         {"touch"},
		},
		ClassGroups: ClassPart{
			NextPart: map[string]ClassPart{
				/**
				 * Aspect Ratio
				 * @see https://tailwindcss.com/docs/aspect-ratio
				 */
				"aspect": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": ClassPart{
							ClassGroupId: "aspect",
						},
						"square": ClassPart{
							ClassGroupId: "aspect",
						},
						"video": ClassPart{
							ClassGroupId: "aspect",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsArbitraryValue,
							ClassGroupId: "aspect",
						},
					},
				},
				/**
				 * Container
				 * @see https://tailwindcss.com/docs/container
				 */
				"container": ClassPart{
					NextPart:     map[string]ClassPart{},
					ClassGroupId: "container",
				},

				/**
				 * Columns
				 * @see https://tailwindcss.com/docs/columns
				 */
				"columns": ClassPart{
					NextPart: map[string]ClassPart{},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsTshirtSize,
							ClassGroupId: "columns",
						},
					},
				},

				"break": ClassPart{
					NextPart: map[string]ClassPart{

						/**
						 * Break After
						 * @see https://tailwindcss.com/docs/break-after
						 */
						"after": ClassPart{
							NextPart: getBreaks("break-after"),
						},

						/** Break Before @see https://tailwindcss.com/docs/break-before
						 */
						"before": ClassPart{
							NextPart: getBreaks("break-before"),
						},

						/**
						 * Break Inside
						 * @see https://tailwindcss.com/docs/break-inside
						 */
						"inside": ClassPart{
							NextPart: map[string]ClassPart{
								"auto": ClassPart{
									ClassGroupId: "break-inside",
								},
								"avoid": ClassPart{
									NextPart: map[string]ClassPart{
										"page": ClassPart{
											ClassGroupId: "break-inside",
										},
										"column": ClassPart{
											ClassGroupId: "break-inside",
										},
									},
									ClassGroupId: "break-inside",
								},
							},
						},

						/**
						 * Word Break
						 * @see https://tailwindcss.com/docs/word-break
						 */

						"normal": ClassPart{
							ClassGroupId: "break",
						},
						"words": ClassPart{
							ClassGroupId: "break",
						},
						"all": ClassPart{
							ClassGroupId: "break",
						},
						"keep": ClassPart{
							ClassGroupId: "break",
						},
					},
					Validators: []ClassGroupValidator{},
				},

				"box": ClassPart{
					NextPart: map[string]ClassPart{
						/**
						 * Box Sizing
						 * @see https://tailwindcss.com/docs/box-sizing
						 */

						"border": ClassPart{
							ClassGroupId: "box",
						},
						"content": ClassPart{
							ClassGroupId: "box",
						},

						/**
						 * Box Decoration Break
						 * @see https://tailwindcss.com/docs/box-decoration-break
						 */

						"decoration": ClassPart{
							NextPart: map[string]ClassPart{
								"slice": ClassPart{
									ClassGroupId: "box-decoration"},
								"clone": ClassPart{
									ClassGroupId: "box-decoration",
								},
							},
						},
					},
				},

				/**
				 * Display
				 * @see https://tailwindcss.com/docs/display
				 */

				"block": {
					ClassGroupId: "display",
				},
				"inline": {
					NextPart: map[string]ClassPart{
						"block": {ClassGroupId: "display"},
						"flex":  {ClassGroupId: "display"},
						"grid":  {ClassGroupId: "display"},
						"table": {ClassGroupId: "display"},
					},
					ClassGroupId: "display",
				},
				"flex": {
					NextPart: map[string]ClassPart{
						"row": ClassPart{
							NextPart: map[string]ClassPart{
								"reverse": ClassPart{
									ClassGroupId: "flex-direction",
								},
							},
							ClassGroupId: "flex-direction",
						},
						"col": ClassPart{
							NextPart: map[string]ClassPart{
								"reverse": ClassPart{
									ClassGroupId: "flex-direction",
								},
							},
							ClassGroupId: "flex-direction",
						},
						"wrap": ClassPart{
							NextPart: map[string]ClassPart{
								"reverse": ClassPart{
									ClassGroupId: "flex-wrap",
								},
							},
							ClassGroupId: "flex-wrap",
						},
						"nowrap": {
							ClassGroupId: "flex-wrap",
						},
						"1": {
							ClassGroupId: "flex",
						},
						"auto": {
							ClassGroupId: "flex",
						},
						"initial": {
							ClassGroupId: "flex",
						},
						"none": {
							ClassGroupId: "flex",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "flex",
						},
					},
					ClassGroupId: "display",
				},
				"table": {
					NextPart: map[string]ClassPart{
						"caption": {
							ClassGroupId: "display",
						},
						"cell": {
							ClassGroupId: "display",
						},
						"column": {
							NextPart: map[string]ClassPart{
								"group": {
									ClassGroupId: "display",
								},
							},
							ClassGroupId: "display",
						},
						"footer": {
							NextPart: map[string]ClassPart{
								"group": {
									ClassGroupId: "display",
								},
							},
						},
						"header": {
							NextPart: map[string]ClassPart{
								"group": {
									ClassGroupId: "display",
								},
							},
						},
						"row": {
							NextPart: map[string]ClassPart{
								"group": {
									ClassGroupId: "display",
								},
							},
							ClassGroupId: "display",
						},
						"auto": {
							ClassGroupId: "table-layout",
						},
						"fixed": {
							ClassGroupId: "table-layout",
						},
					},
					ClassGroupId: "display",
				},
				"flow": {
					NextPart: map[string]ClassPart{"root": {ClassGroupId: "display"}},
				},
				"grid": ClassPart{
					NextPart: map[string]ClassPart{
						"cols": {
							Validators: []ClassGroupValidator{
								{
									Fn:           IsAny,
									ClassGroupId: "grid-cols",
								},
							},
						},
						"rows": {
							Validators: []ClassGroupValidator{
								{
									Fn:           IsAny,
									ClassGroupId: "grid-rows",
								},
							},
						},
						"flow": {
							NextPart: map[string]ClassPart{
								"row": {
									NextPart: map[string]ClassPart{
										"dense": {
											ClassGroupId: "grid-flow",
										},
									},
									ClassGroupId: "grid-flow",
								},
								"col": {
									NextPart: map[string]ClassPart{
										"dense": {
											ClassGroupId: "grid-flow",
										},
									},
									ClassGroupId: "grid-flow",
								},
								"dense": {
									ClassGroupId: "grid-flow",
								},
							},
						},
					},
					Validators:   []ClassGroupValidator{},
					ClassGroupId: "display",
				},
				"contents": {ClassGroupId: "display"},
				"list": ClassPart{
					NextPart: map[string]ClassPart{
						"item": {
							ClassGroupId: "display",
						},
						"image": {
							NextPart: map[string]ClassPart{
								"none": {
									ClassGroupId: "list-image",
								},
							},
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "list-image",
								},
							},
						},
						"none": {
							ClassGroupId: "list-style-type",
						},
						"disc": {
							ClassGroupId: "list-style-type",
						},
						"decimal": {
							ClassGroupId: "list-style-type",
						},
						"inside": {
							ClassGroupId: "list-style-position",
						},
						"outside": {
							ClassGroupId: "list-style-position",
						},
					},
					Validators: []ClassGroupValidator{
						{
							// fn : TODO: You need to provide the function implementation here
							ClassGroupId: "list-style-type",
						},
					},
				},
				"hidden": {ClassGroupId: "display"},
				"float": ClassPart{
					NextPart: map[string]ClassPart{
						"right": {
							ClassGroupId: "float",
						},
						"left": {
							ClassGroupId: "float",
						},
						"none": {
							ClassGroupId: "float",
						},
						"start": {
							ClassGroupId: "float",
						},
						"end": {
							ClassGroupId: "float",
						},
					},
				},
				"clear": ClassPart{
					NextPart: map[string]ClassPart{
						"left": {
							ClassGroupId: "clear",
						},
						"right": {
							ClassGroupId: "clear",
						},
						"both": {
							ClassGroupId: "clear",
						},
						"none": {
							ClassGroupId: "clear",
						},
						"start": {
							ClassGroupId: "clear",
						},
						"end": {
							ClassGroupId: "clear",
						},
					},
				},
				"isolate": {ClassGroupId: "isolation"},
				"isolation": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": {
							ClassGroupId: "isolation",
						},
					},
				},
				"object": ClassPart{
					NextPart: map[string]ClassPart{
						"contain": {
							ClassGroupId: "object-fit",
						},
						"cover": {
							ClassGroupId: "object-fit",
						},
						"fill": {
							ClassGroupId: "object-fit",
						},
						"none": {
							ClassGroupId: "object-fit",
						},
						"scale": {
							NextPart: map[string]ClassPart{
								"down": {
									ClassGroupId: "object-fit",
								},
							},
						},
						"bottom": {
							ClassGroupId: "object-position",
						},
						"center": {
							ClassGroupId: "object-position",
						},
						"left": {
							NextPart: map[string]ClassPart{
								"bottom": {
									ClassGroupId: "object-position",
								},
								"top": {
									ClassGroupId: "object-position",
								},
							},
						},
						"right": {
							NextPart: map[string]ClassPart{
								"bottom": {
									ClassGroupId: "object-position",
								},
								"top": {
									ClassGroupId: "object-position",
								},
							},
						},
						"top": {
							ClassGroupId: "object-position",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "object-position",
						},
					},
				},

				"overflow": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": {
							ClassGroupId: "overflow",
						},
						"hidden": {
							ClassGroupId: "overflow",
						},
						"clip": {
							ClassGroupId: "overflow",
						},
						"visible": {
							ClassGroupId: "overflow",
						},
						"scroll": {
							ClassGroupId: "overflow",
						},
						"x": {
							NextPart: map[string]ClassPart{
								"auto": {
									ClassGroupId: "overflow-x",
								},
								"hidden": {
									ClassGroupId: "overflow-x",
								},
								"clip": {
									ClassGroupId: "overflow-x",
								},
								"visible": {
									ClassGroupId: "overflow-x",
								},
								"scroll": {
									ClassGroupId: "overflow-x",
								},
							},
						},
						"y": {
							NextPart: map[string]ClassPart{
								"auto": {
									ClassGroupId: "overflow-y",
								},
								"hidden": {
									ClassGroupId: "overflow-y",
								},
								"clip": {
									ClassGroupId: "overflow-y",
								},
								"visible": {
									ClassGroupId: "overflow-y",
								},
								"scroll": {
									ClassGroupId: "overflow-y",
								},
							},
						},
					},
				},
				"overscroll": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": {
							ClassGroupId: "overscroll",
						},
						"contain": {
							ClassGroupId: "overscroll",
						},
						"none": {
							ClassGroupId: "overscroll",
						},
						"x": {
							NextPart: map[string]ClassPart{
								"auto": {
									ClassGroupId: "overscroll-x",
								},
								"contain": {
									ClassGroupId: "overscroll-x",
								},
								"none": {
									ClassGroupId: "overscroll-x",
								},
							},
						},
						"y": {
							NextPart: map[string]ClassPart{
								"auto": {
									ClassGroupId: "overscroll-y",
								},
								"contain": {
									ClassGroupId: "overscroll-y",
								},
								"none": {
									ClassGroupId: "overscroll-y",
								},
							},
						},
					},
					Validators: []ClassGroupValidator{},
				},

				"static": ClassPart{
					ClassGroupId: "position",
				},
				"fixed": ClassPart{
					ClassGroupId: "position",
				},
				"absolute": ClassPart{
					ClassGroupId: "position",
				},
				"relative": ClassPart{
					ClassGroupId: "position",
				},
				"sticky": ClassPart{
					ClassGroupId: "position",
				},

				"inset": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": {
							ClassGroupId: "inset",
						},
						"x": {
							NextPart: map[string]ClassPart{
								"auto": {
									ClassGroupId: "inset-x",
								},
							},
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "inset-x",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "inset-x",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "inset-x",
								},
							},
						},
						"y": {
							NextPart: map[string]ClassPart{
								"auto": {
									ClassGroupId: "inset-y",
								},
							},
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "inset-y",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "inset-y",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "inset-y",
								},
							},
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "inset",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "inset",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "inset",
						},
					},
				},
				"start": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": {
							ClassGroupId: "start",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "start",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "start",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "start",
						},
					},
				},
				"end": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": {
							ClassGroupId: "end",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "end",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "end",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "end",
						},
					},
				},
				"top": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": {
							ClassGroupId: "top",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "top",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "top",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "top",
						},
					},
				},
				"right": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": {
							ClassGroupId: "right",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "right",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "right",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "right",
						},
					},
				},
				"bottom": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": {
							ClassGroupId: "bottom",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "bottom",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "bottom",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "bottom",
						},
					},
				},
				"left": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": {
							ClassGroupId: "left",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "left",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "left",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "left",
						},
					},
				},
				"visible": ClassPart{
					ClassGroupId: "visibility",
				},
				"invisible": ClassPart{
					ClassGroupId: "visibility",
				},
				"collapse": ClassPart{
					ClassGroupId: "visibility",
				},
				"z": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": {
							ClassGroupId: "z",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsInteger,
							ClassGroupId: "z",
						},
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "z",
						},
					},
				},
				"basis": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": {
							ClassGroupId: "basis",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "basis",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "basis",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "basis",
						},
					},
				},
				"grow": ClassPart{
					NextPart: map[string]ClassPart{
						"0": {
							ClassGroupId: "grow",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "grow",
						},
					},
					ClassGroupId: "grow",
				},
				"shrink": ClassPart{
					NextPart: map[string]ClassPart{
						"0": {
							ClassGroupId: "shrink",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "shrink",
						},
					},
					ClassGroupId: "shrink",
				},
				"order": ClassPart{
					NextPart: map[string]ClassPart{
						"first": {
							ClassGroupId: "order",
						},
						"last": {
							ClassGroupId: "order",
						},
						"none": {
							ClassGroupId: "order",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsInteger,
							ClassGroupId: "order",
						},
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "order",
						},
					},
				},
				"col": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": {
							NextPart:     map[string]ClassPart{},
							Validators:   []ClassGroupValidator{},
							ClassGroupId: "col-start-end",
						},
						"span": {
							NextPart: map[string]ClassPart{
								"full": {
									NextPart:     map[string]ClassPart{},
									Validators:   []ClassGroupValidator{},
									ClassGroupId: "col-start-end",
								},
							},
							Validators: []ClassGroupValidator{
								{
									Fn:           IsInteger,
									ClassGroupId: "col-start-end",
								},
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "col-start-end",
								},
							},
						},
						"start": {
							NextPart: map[string]ClassPart{
								"auto": {
									NextPart:     map[string]ClassPart{},
									Validators:   []ClassGroupValidator{},
									ClassGroupId: "col-start",
								},
							},
							Validators: []ClassGroupValidator{
								{
									Fn:           IsNumber,
									ClassGroupId: "col-start",
								},
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "col-start",
								},
							},
						},
						"end": {
							NextPart: map[string]ClassPart{
								"auto": {
									NextPart:     map[string]ClassPart{},
									Validators:   []ClassGroupValidator{},
									ClassGroupId: "col-end",
								},
							},
							Validators: []ClassGroupValidator{
								{
									Fn:           IsNumber,
									ClassGroupId: "col-end",
								},
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "col-end",
								},
							},
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "col-start-end",
						},
					},
				},
				"row": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": {
							NextPart:     map[string]ClassPart{},
							Validators:   []ClassGroupValidator{},
							ClassGroupId: "row-start-end",
						},
						"span": {
							NextPart: map[string]ClassPart{},
							Validators: []ClassGroupValidator{
								{
									Fn:           IsInteger,
									ClassGroupId: "row-start-end",
								},
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "row-start-end",
								},
							},
						},
						"start": {
							NextPart: map[string]ClassPart{
								"auto": {
									NextPart:     map[string]ClassPart{},
									Validators:   []ClassGroupValidator{},
									ClassGroupId: "row-start",
								},
							},
							Validators: []ClassGroupValidator{
								{
									Fn:           IsNumber,
									ClassGroupId: "row-start",
								},
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "row-start",
								},
							},
						},
						"end": {
							NextPart: map[string]ClassPart{
								"auto": {
									NextPart:     map[string]ClassPart{},
									Validators:   []ClassGroupValidator{},
									ClassGroupId: "row-end",
								},
							},
							Validators: []ClassGroupValidator{
								{
									Fn:           IsNumber,
									ClassGroupId: "row-end",
								},
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "row-end",
								},
							},
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "row-start-end",
						},
					},
				},
				"auto": ClassPart{
					NextPart: map[string]ClassPart{
						"cols": {
							NextPart: map[string]ClassPart{
								"auto": {
									NextPart:     map[string]ClassPart{},
									Validators:   []ClassGroupValidator{},
									ClassGroupId: "auto-cols",
								},
								"min": {
									NextPart:     map[string]ClassPart{},
									Validators:   []ClassGroupValidator{},
									ClassGroupId: "auto-cols",
								},
								"max": {
									NextPart:     map[string]ClassPart{},
									Validators:   []ClassGroupValidator{},
									ClassGroupId: "auto-cols",
								},
								"fr": {
									NextPart:     map[string]ClassPart{},
									Validators:   []ClassGroupValidator{},
									ClassGroupId: "auto-cols",
								},
							},
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "auto-cols",
								},
							},
						},
						"rows": {
							NextPart: map[string]ClassPart{
								"auto": {
									NextPart:     map[string]ClassPart{},
									Validators:   []ClassGroupValidator{},
									ClassGroupId: "auto-rows",
								},
								"min": {
									NextPart:     map[string]ClassPart{},
									Validators:   []ClassGroupValidator{},
									ClassGroupId: "auto-rows",
								},
								"max": {
									NextPart:     map[string]ClassPart{},
									Validators:   []ClassGroupValidator{},
									ClassGroupId: "auto-rows",
								},
								"fr": {
									NextPart:     map[string]ClassPart{},
									Validators:   []ClassGroupValidator{},
									ClassGroupId: "auto-rows",
								},
							},
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "auto-rows",
								},
							},
						},
					},
					Validators: []ClassGroupValidator{},
				},
				"gap": ClassPart{
					NextPart: map[string]ClassPart{
						"x": {
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "gap-x",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "gap-x",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "gap-x",
								},
							},
						},
						"y": {
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "gap-y",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "gap-y",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "gap-y",
								},
							},
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "gap",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "gap",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "gap",
						},
					},
				},
				"justify": ClassPart{
					NextPart: map[string]ClassPart{
						"normal": ClassPart{
							ClassGroupId: "justify-content",
						},
						"start": ClassPart{
							ClassGroupId: "justify-content",
						},
						"end": ClassPart{
							ClassGroupId: "justify-content",
						},
						"center": ClassPart{
							ClassGroupId: "justify-content",
						},
						"between": ClassPart{
							ClassGroupId: "justify-content",
						},
						"around": ClassPart{
							ClassGroupId: "justify-content",
						},
						"evenly": ClassPart{
							ClassGroupId: "justify-content",
						},
						"stretch": ClassPart{
							ClassGroupId: "justify-content",
						},
						"items": ClassPart{
							NextPart: map[string]ClassPart{
								"start": ClassPart{
									ClassGroupId: "justify-items",
								},
								"end": ClassPart{
									ClassGroupId: "justify-items",
								},
								"center": ClassPart{
									ClassGroupId: "justify-items",
								},
								"stretch": ClassPart{
									ClassGroupId: "justify-items",
								},
							},
						},
						"self": ClassPart{
							NextPart: map[string]ClassPart{
								"auto": ClassPart{
									ClassGroupId: "justify-self",
								},
								"start": ClassPart{
									ClassGroupId: "justify-self",
								},
								"end": ClassPart{
									ClassGroupId: "justify-self",
								},
								"center": ClassPart{
									ClassGroupId: "justify-self",
								},
								"stretch": ClassPart{
									ClassGroupId: "justify-self",
								},
							},
						},
					},
				},
				"content": ClassPart{
					NextPart: map[string]ClassPart{
						"normal": ClassPart{
							ClassGroupId: "align-content",
						},
						"start": ClassPart{
							ClassGroupId: "align-content",
						},
						"end": ClassPart{
							ClassGroupId: "align-content",
						},
						"center": ClassPart{
							ClassGroupId: "align-content",
						},
						"between": ClassPart{
							ClassGroupId: "align-content",
						},
						"around": ClassPart{
							ClassGroupId: "align-content",
						},
						"evenly": ClassPart{
							ClassGroupId: "align-content",
						},
						"stretch": ClassPart{
							ClassGroupId: "align-content",
						},
						"baseline": ClassPart{
							ClassGroupId: "align-content",
						},
						"none": ClassPart{
							ClassGroupId: "content",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "content",
						},
					},
				},
				"items": ClassPart{
					NextPart: map[string]ClassPart{
						"start": ClassPart{
							ClassGroupId: "align-items",
						},
						"end": ClassPart{
							ClassGroupId: "align-items",
						},
						"center": ClassPart{
							ClassGroupId: "align-items",
						},
						"baseline": ClassPart{
							ClassGroupId: "align-items",
						},
						"stretch": ClassPart{
							ClassGroupId: "align-items",
						},
					},
				},
				"self": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": ClassPart{
							ClassGroupId: "align-self",
						},
						"start": ClassPart{
							ClassGroupId: "align-self",
						},
						"end": ClassPart{
							ClassGroupId: "align-self",
						},
						"center": ClassPart{
							ClassGroupId: "align-self",
						},
						"stretch": ClassPart{
							ClassGroupId: "align-self",
						},
						"baseline": ClassPart{
							ClassGroupId: "align-self",
						},
					},
				},
				"place": ClassPart{
					NextPart: map[string]ClassPart{
						"content": ClassPart{
							NextPart: map[string]ClassPart{
								"start": ClassPart{
									ClassGroupId: "place-content",
								},
								"end": ClassPart{
									ClassGroupId: "place-content",
								},
								"center": ClassPart{
									ClassGroupId: "place-content",
								},
								"between": ClassPart{
									ClassGroupId: "place-content",
								},
								"around": ClassPart{
									ClassGroupId: "place-content",
								},
								"evenly": ClassPart{
									ClassGroupId: "place-content",
								},
								"stretch": ClassPart{
									ClassGroupId: "place-content",
								},
								"baseline": ClassPart{
									ClassGroupId: "place-content",
								},
							},
						},
						"items": ClassPart{
							NextPart: map[string]ClassPart{
								"start": ClassPart{
									ClassGroupId: "place-items",
								},
								"end": ClassPart{
									ClassGroupId: "place-items",
								},
								"center": ClassPart{
									ClassGroupId: "place-items",
								},
								"baseline": ClassPart{
									ClassGroupId: "place-items",
								},
								"stretch": ClassPart{
									ClassGroupId: "place-items",
								},
							},
						},
						"self": ClassPart{
							NextPart: map[string]ClassPart{
								"auto": ClassPart{
									ClassGroupId: "place-self",
								},
								"start": ClassPart{
									ClassGroupId: "place-self",
								},
								"end": ClassPart{
									ClassGroupId: "place-self",
								},
								"center": ClassPart{
									ClassGroupId: "place-self",
								},
								"stretch": ClassPart{
									ClassGroupId: "place-self",
								},
							},
						},
					},
				},
				"p": ClassPart{
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "p",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "p",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "p",
						},
					},
				},
				"px": ClassPart{
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "px",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "px",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "px",
						},
					},
				},
				"py": ClassPart{
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "py",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "py",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "py",
						},
					},
				},
				"ps": ClassPart{
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "ps",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "ps",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "ps",
						},
					},
				},
				"pe": ClassPart{
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "pe",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "pe",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "pe",
						},
					},
				},
				"pt": ClassPart{
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "pt",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "pt",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "pt",
						},
					},
				},
				"pr": ClassPart{
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "pr",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "pr",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "pr",
						},
					},
				},
				"pb": ClassPart{
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "pb",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "pb",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "pb",
						},
					},
				},
				"pl": ClassPart{
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "pl",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "pl",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "pl",
						},
					},
				},
				"m": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": ClassPart{
							Validators:   []ClassGroupValidator{},
							ClassGroupId: "m",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "m",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "m",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "m",
						},
					},
				},
				"mx": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": ClassPart{
							Validators:   []ClassGroupValidator{},
							ClassGroupId: "mx",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "mx",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "mx",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "mx",
						},
					},
				},
				"my": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": ClassPart{
							Validators:   []ClassGroupValidator{},
							ClassGroupId: "my",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "my",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "my",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "my",
						},
					},
				},
				"ms": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": ClassPart{
							Validators:   []ClassGroupValidator{},
							ClassGroupId: "ms",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "ms",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "ms",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "ms",
						},
					},
				},
				"me": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": ClassPart{
							Validators:   []ClassGroupValidator{},
							ClassGroupId: "me",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "me",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "me",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "me",
						},
					},
				},
				"mt": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": ClassPart{
							Validators:   []ClassGroupValidator{},
							ClassGroupId: "mt",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "mt",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "mt",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "mt",
						},
					},
				},
				"mr": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": ClassPart{
							Validators:   []ClassGroupValidator{},
							ClassGroupId: "mr",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "mr",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "mr",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "mr",
						},
					},
				},
				"mb": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": ClassPart{
							Validators:   []ClassGroupValidator{},
							ClassGroupId: "mb",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "mb",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "mb",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "mb",
						},
					},
				},
				"ml": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": ClassPart{
							Validators:   []ClassGroupValidator{},
							ClassGroupId: "ml",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "ml",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "ml",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "ml",
						},
					},
				},
				"space": ClassPart{
					NextPart: map[string]ClassPart{
						"x": ClassPart{
							NextPart: map[string]ClassPart{
								"reverse": ClassPart{
									Validators:   []ClassGroupValidator{},
									ClassGroupId: "space-x-reverse",
								},
							},
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "space-x",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "space-x",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "space-x",
								},
							},
						},
						"y": ClassPart{
							NextPart: map[string]ClassPart{
								"reverse": ClassPart{
									ClassGroupId: "space-y-reverse",
								},
							},
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "space-y",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "space-y",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "space-y",
								},
							},
						},
					},
					Validators: []ClassGroupValidator{},
				},
				"w": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": ClassPart{
							ClassGroupId: "w",
						},
						"min": ClassPart{
							ClassGroupId: "w",
						},
						"max": ClassPart{
							ClassGroupId: "w",
						},
						"fit": ClassPart{
							ClassGroupId: "w",
						},
						"svw": ClassPart{
							ClassGroupId: "w",
						},
						"lvw": ClassPart{
							ClassGroupId: "w",
						},
						"dvw": ClassPart{
							ClassGroupId: "w",
						},
					},
					Validators: []ClassGroupValidator{
						{
							Fn:           IsArbitraryValue,
							ClassGroupId: "w",
						},
						{
							Fn:           IsLength,
							ClassGroupId: "w",
						},
						{
							Fn:           IsArbitraryLength,
							ClassGroupId: "w",
						},
					},
				},
				"min": ClassPart{
					NextPart: map[string]ClassPart{
						"w": ClassPart{
							NextPart: map[string]ClassPart{
								"min": ClassPart{
									ClassGroupId: "min-w",
								},
								"max": ClassPart{
									ClassGroupId: "min-w",
								},
								"fit": ClassPart{
									ClassGroupId: "min-w",
								},
							},
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "min-w",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "min-w",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "min-w",
								},
							},
						},
						"h": ClassPart{
							NextPart: map[string]ClassPart{
								"min": ClassPart{
									ClassGroupId: "min-h",
								},
								"max": ClassPart{
									ClassGroupId: "min-h",
								},
								"fit": ClassPart{
									ClassGroupId: "min-h",
								},
								"svh": ClassPart{
									ClassGroupId: "min-h",
								},
								"lvh": ClassPart{
									ClassGroupId: "min-h",
								},
								"dvh": ClassPart{
									ClassGroupId: "min-h",
								},
							},
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "min-h",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "min-h",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "min-h",
								},
							},
						},
					},
					Validators: []ClassGroupValidator{},
				},
				"max": ClassPart{
					NextPart: map[string]ClassPart{
						"w": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "max-w",
								},
								"full": ClassPart{
									ClassGroupId: "max-w",
								},
								"min": ClassPart{
									ClassGroupId: "max-w",
								},
								"max": ClassPart{
									ClassGroupId: "max-w",
								},
								"fit": ClassPart{
									ClassGroupId: "max-w",
								},
								"prose": ClassPart{
									ClassGroupId: "max-w",
								},
								"screen": ClassPart{
									Validators: []ClassGroupValidator{
										ClassGroupValidator{
											Fn:           IsTshirtSize,
											ClassGroupId: "max-w",
										},
									},
									ClassGroupId: "max-w",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "max-w",
								},
								ClassGroupValidator{
									Fn:           IsLength,
									ClassGroupId: "max-w",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryLength,
									ClassGroupId: "max-w",
								},
								ClassGroupValidator{
									Fn:           IsTshirtSize,
									ClassGroupId: "max-w",
								},
							},
							ClassGroupId: "max-w",
						},
						"h": ClassPart{
							NextPart: map[string]ClassPart{
								"min": ClassPart{
									ClassGroupId: "max-h",
								},
								"max": ClassPart{
									ClassGroupId: "max-h",
								},
								"fit": ClassPart{
									ClassGroupId: "max-h",
								},
								"svh": ClassPart{
									ClassGroupId: "max-h",
								},
								"lvh": ClassPart{
									ClassGroupId: "max-h",
								},
								"dvh": ClassPart{
									ClassGroupId: "max-h",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "max-h",
								},
								ClassGroupValidator{
									Fn:           IsLength,
									ClassGroupId: "max-h",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryLength,
									ClassGroupId: "max-h",
								},
							},
							ClassGroupId: "max-h",
						},
					},
				},
				"h": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": ClassPart{
							ClassGroupId: "h",
						},
						"min": ClassPart{
							ClassGroupId: "h",
						},
						"max": ClassPart{
							ClassGroupId: "h",
						},
						"fit": ClassPart{
							ClassGroupId: "h",
						},
						"svh": ClassPart{
							ClassGroupId: "h",
						},
						"lvh": ClassPart{
							ClassGroupId: "h",
						},
						"dvh": ClassPart{
							ClassGroupId: "h",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsArbitraryValue,
							ClassGroupId: "h",
						},
						ClassGroupValidator{
							Fn:           IsLength,
							ClassGroupId: "h",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryLength,
							ClassGroupId: "h",
						},
					},
					ClassGroupId: "h",
				},
				"size": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": ClassPart{
							ClassGroupId: "size",
						},
						"min": ClassPart{
							ClassGroupId: "size",
						},
						"max": ClassPart{
							ClassGroupId: "size",
						},
						"fit": ClassPart{
							ClassGroupId: "size",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsArbitraryValue,
							ClassGroupId: "size",
						},
						ClassGroupValidator{
							Fn:           IsLength,
							ClassGroupId: "size",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryLength,
							ClassGroupId: "size",
						},
					},
					ClassGroupId: "size",
				},
				"text": ClassPart{
					NextPart: map[string]ClassPart{
						"base": ClassPart{
							ClassGroupId: "font-size",
						},
						"left": ClassPart{
							ClassGroupId: "text-alignment",
						},
						"center": ClassPart{
							ClassGroupId: "text-alignment",
						},
						"right": ClassPart{
							ClassGroupId: "text-alignment",
						},
						"justify": ClassPart{
							ClassGroupId: "text-alignment",
						},
						"start": ClassPart{
							ClassGroupId: "text-alignment",
						},
						"end": ClassPart{
							ClassGroupId: "text-alignment",
						},
						"opacity": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsNumber,
									ClassGroupId: "text-opacity",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryNumber,
									ClassGroupId: "text-opacity",
								},
							},
							ClassGroupId: "text-opacity",
						},
						"ellipsis": ClassPart{
							ClassGroupId: "text-overflow",
						},
						"clip": ClassPart{
							ClassGroupId: "text-overflow",
						},
						"wrap": ClassPart{
							ClassGroupId: "text-wrap",
						},
						"nowrap": ClassPart{
							ClassGroupId: "text-wrap",
						},
						"balance": ClassPart{
							ClassGroupId: "text-wrap",
						},
						"pretty": ClassPart{
							ClassGroupId: "text-wrap",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsTshirtSize,
							ClassGroupId: "font-size",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryLength,
							ClassGroupId: "font-size",
						},
						ClassGroupValidator{
							Fn:           IsAny,
							ClassGroupId: "text-color",
						},
					},
				},
				"antialiased": ClassPart{
					ClassGroupId: "font-smoothing",
				},
				"subpixel": ClassPart{
					NextPart: map[string]ClassPart{
						"antialiased": ClassPart{
							ClassGroupId: "font-smoothing",
						},
					},
				},
				"italic": ClassPart{
					ClassGroupId: "font-style",
				},
				"not": ClassPart{
					NextPart: map[string]ClassPart{
						"italic": ClassPart{
							ClassGroupId: "font-style",
						},
						"sr": ClassPart{
							NextPart: map[string]ClassPart{
								"only": ClassPart{
									ClassGroupId: "sr",
								},
							},
						},
					},
				},
				"font": ClassPart{
					NextPart: map[string]ClassPart{
						"thin": ClassPart{
							ClassGroupId: "font-weight",
						},
						"extralight": ClassPart{
							ClassGroupId: "font-weight",
						},
						"light": ClassPart{
							ClassGroupId: "font-weight",
						},
						"normal": ClassPart{
							ClassGroupId: "font-weight",
						},
						"medium": ClassPart{
							ClassGroupId: "font-weight",
						},
						"semibold": ClassPart{
							ClassGroupId: "font-weight",
						},
						"bold": ClassPart{
							ClassGroupId: "font-weight",
						},
						"extrabold": ClassPart{
							ClassGroupId: "font-weight",
						},
						"black": ClassPart{
							ClassGroupId: "font-weight",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsArbitraryNumber,
							ClassGroupId: "font-weight",
						},
						ClassGroupValidator{
							Fn:           IsAny,
							ClassGroupId: "font-family",
						},
					},
				},
				"normal": ClassPart{
					NextPart: map[string]ClassPart{
						"nums": ClassPart{
							ClassGroupId: "fvn-normal",
						},
						"case": ClassPart{
							ClassGroupId: "text-transform",
						},
					},
				},
				"ordinal": ClassPart{
					ClassGroupId: "fvn-ordinal",
				},
				"slashed": ClassPart{
					NextPart: map[string]ClassPart{
						"zero": ClassPart{
							ClassGroupId: "fvn-slashed-zero",
						},
					},
				},
				"lining": ClassPart{
					NextPart: map[string]ClassPart{
						"nums": ClassPart{
							ClassGroupId: "fvn-figure",
						},
					},
				},
				"oldstyle": ClassPart{
					NextPart: map[string]ClassPart{
						"nums": ClassPart{
							ClassGroupId: "fvn-figure",
						},
					},
				},
				"proportional": ClassPart{
					NextPart: map[string]ClassPart{
						"nums": ClassPart{
							ClassGroupId: "fvn-spacing",
						},
					},
				},
				"tabular": ClassPart{
					NextPart: map[string]ClassPart{
						"nums": ClassPart{
							ClassGroupId: "fvn-spacing",
						},
					},
				},
				"diagonal": ClassPart{
					NextPart: map[string]ClassPart{
						"fractions": ClassPart{
							ClassGroupId: "fvn-fraction",
						},
					},
				},
				"stacked": ClassPart{
					NextPart: map[string]ClassPart{
						"fractons": ClassPart{
							ClassGroupId: "fvn-fraction",
						},
					},
				},
				"tracking": ClassPart{
					NextPart: map[string]ClassPart{
						"tighter": ClassPart{
							ClassGroupId: "tracking",
						},
						"tight": ClassPart{
							ClassGroupId: "tracking",
						},
						"normal": ClassPart{
							ClassGroupId: "tracking",
						},
						"wide": ClassPart{
							ClassGroupId: "tracking",
						},
						"wider": ClassPart{
							ClassGroupId: "tracking",
						},
						"widest": ClassPart{
							ClassGroupId: "tracking",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsArbitraryValue,
							ClassGroupId: "tracking",
						},
					},
				},
				"line": ClassPart{
					NextPart: map[string]ClassPart{
						"clamp": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "line-clamp",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsNumber,
									ClassGroupId: "line-clamp",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryNumber,
									ClassGroupId: "line-clamp",
								},
							},
						},
						"through": ClassPart{
							ClassGroupId: "text-decoration",
						},
					},
				},
				"leading": ClassPart{
					NextPart: map[string]ClassPart{
						"none": ClassPart{
							ClassGroupId: "leading",
						},
						"tight": ClassPart{
							ClassGroupId: "leading",
						},
						"snug": ClassPart{
							ClassGroupId: "leading",
						},
						"normal": ClassPart{
							ClassGroupId: "leading",
						},
						"relaxed": ClassPart{
							ClassGroupId: "leading",
						},
						"loose": ClassPart{
							ClassGroupId: "leading",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsLength,
							ClassGroupId: "leading",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryValue,
							ClassGroupId: "leading",
						},
					},
				},
				"placeholder": ClassPart{
					NextPart: map[string]ClassPart{
						"opacity": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsNumber,
									ClassGroupId: "placeholder-opacity",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryNumber,
									ClassGroupId: "placeholder-opacity",
								},
							},
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsAny,
							ClassGroupId: "placeholder-color",
						},
					},
				},
				"underline": ClassPart{
					NextPart: map[string]ClassPart{
						"offset": ClassPart{
							NextPart: map[string]ClassPart{
								"auto": ClassPart{
									ClassGroupId: "underline-offset",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsLength,
									ClassGroupId: "underline-offset",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "underline-offset",
								},
							},
						},
					},
					ClassGroupId: "text-decoration",
				},
				"overline": ClassPart{
					ClassGroupId: "text-decoration",
				},
				"no": ClassPart{
					NextPart: map[string]ClassPart{
						"underline": ClassPart{
							ClassGroupId: "text-decoration",
						},
					},
				},
				"decoration": ClassPart{
					NextPart: map[string]ClassPart{
						"solid": ClassPart{
							ClassGroupId: "text-decoration-style",
						},
						"dashed": ClassPart{
							ClassGroupId: "text-decoration-style",
						},
						"dotted": ClassPart{
							ClassGroupId: "text-decoration-style",
						},
						"double": ClassPart{
							ClassGroupId: "text-decoration-style",
						},
						"none": ClassPart{
							ClassGroupId: "text-decoration-style",
						},
						"wavy": ClassPart{
							ClassGroupId: "text-decoration-style",
						},
						"auto": ClassPart{
							ClassGroupId: "text-decoration-thickness",
						},
						"from": ClassPart{
							NextPart: map[string]ClassPart{
								"font": ClassPart{
									ClassGroupId: "text-decoration-thickness",
								},
							},
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsLength,
							ClassGroupId: "text-decoration-thickness",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryLength,
							ClassGroupId: "text-decoration-thickness",
						},
						ClassGroupValidator{
							Fn:           IsAny,
							ClassGroupId: "text-decoration-color",
						},
					},
					ClassGroupId: "",
				},
				"uppercase": ClassPart{
					ClassGroupId: "text-transform",
				},
				"lowercase": ClassPart{
					ClassGroupId: "text-transform",
				},
				"capitalize": ClassPart{
					ClassGroupId: "text-transform",
				},
				"truncate": ClassPart{
					ClassGroupId: "text-overflow",
				},
				"indent": ClassPart{
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsArbitraryValue,
							ClassGroupId: "indent",
						},
						ClassGroupValidator{
							Fn:           IsLength,
							ClassGroupId: "indent",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryLength,
							ClassGroupId: "indent",
						},
					},
				},
				"align": ClassPart{
					NextPart: map[string]ClassPart{
						"baseline": ClassPart{
							ClassGroupId: "vertical-align",
						},
						"top": ClassPart{
							ClassGroupId: "vertical-align",
						},
						"middle": ClassPart{
							ClassGroupId: "vertical-align",
						},
						"bottom": ClassPart{
							ClassGroupId: "vertical-align",
						},
						"text": ClassPart{
							NextPart: map[string]ClassPart{
								"top": ClassPart{
									ClassGroupId: "vertical-align",
								},
								"bottom": ClassPart{
									ClassGroupId: "vertical-align",
								},
							},
						},
						"sub": ClassPart{
							ClassGroupId: "vertical-align",
						},
						"super": ClassPart{
							ClassGroupId: "vertical-align",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsArbitraryValue,
							ClassGroupId: "vertical-align",
						},
					},
				},
				"whitespace": ClassPart{
					NextPart: map[string]ClassPart{
						"normal": ClassPart{
							ClassGroupId: "whitespace",
						},
						"nowrap": ClassPart{
							ClassGroupId: "whitespace",
						},
						"pre": ClassPart{
							NextPart: map[string]ClassPart{
								"line": ClassPart{
									ClassGroupId: "whitespace",
								},
								"wrap": ClassPart{
									ClassGroupId: "whitespace",
								},
							},
							ClassGroupId: "whitespace",
						},
						"break": ClassPart{
							NextPart: map[string]ClassPart{
								"spaces": ClassPart{
									ClassGroupId: "whitespace",
								},
							},
							ClassGroupId: "",
						},
					},
				},
				"hyphens": ClassPart{
					NextPart: map[string]ClassPart{
						"none": ClassPart{
							ClassGroupId: "hyphens",
						},
						"manual": ClassPart{
							ClassGroupId: "hyphens",
						},
						"auto": ClassPart{
							ClassGroupId: "hyphens",
						},
					},
				},
				"bg": ClassPart{
					NextPart: map[string]ClassPart{
						"fixed": ClassPart{
							ClassGroupId: "bg-attachment",
						},
						"local": ClassPart{
							ClassGroupId: "bg-attachment",
						},
						"scroll": ClassPart{
							ClassGroupId: "bg-attachment",
						},
						"clip": ClassPart{
							NextPart: map[string]ClassPart{
								"border": ClassPart{
									ClassGroupId: "bg-clip",
								},
								"padding": ClassPart{
									ClassGroupId: "bg-clip",
								},
								"content": ClassPart{
									ClassGroupId: "bg-clip",
								},
								"text": ClassPart{
									ClassGroupId: "bg-clip",
								},
							},
						},
						"opacity": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsNumber,
									ClassGroupId: "bg-opacity",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryNumber,
									ClassGroupId: "bg-opacity",
								},
							},
						},
						"origin": ClassPart{
							NextPart: map[string]ClassPart{
								"border": ClassPart{
									ClassGroupId: "bg-origin",
								},
								"padding": ClassPart{
									ClassGroupId: "bg-origin",
								},
								"content": ClassPart{
									ClassGroupId: "bg-origin",
								},
							},
						},
						"bottom": ClassPart{
							ClassGroupId: "bg-position",
						},
						"center": ClassPart{
							ClassGroupId: "bg-position",
						},
						"left": ClassPart{
							NextPart: map[string]ClassPart{
								"bottom": ClassPart{
									ClassGroupId: "bg-position",
								},
								"top": ClassPart{
									ClassGroupId: "bg-position",
								},
							},
							ClassGroupId: "bg-position",
						},
						"right": ClassPart{
							NextPart: map[string]ClassPart{
								"bottom": ClassPart{
									ClassGroupId: "bg-position",
								},
								"top": ClassPart{
									ClassGroupId: "bg-position",
								},
							},
							ClassGroupId: "bg-position",
						},
						"top": ClassPart{
							ClassGroupId: "bg-position",
						},
						"no": ClassPart{
							NextPart: map[string]ClassPart{
								"repeat": ClassPart{
									ClassGroupId: "bg-repeat",
								},
							},
						},
						"repeat": ClassPart{
							NextPart: map[string]ClassPart{
								"x": ClassPart{
									ClassGroupId: "bg-repeat",
								},
								"y": ClassPart{
									ClassGroupId: "bg-repeat",
								},
								"round": ClassPart{
									ClassGroupId: "bg-repeat",
								},
								"space": ClassPart{
									ClassGroupId: "bg-repeat",
								},
							},
							ClassGroupId: "bg-repeat",
						},
						"auto": ClassPart{
							ClassGroupId: "bg-size",
						},
						"cover": ClassPart{
							ClassGroupId: "bg-size",
						},
						"contain": ClassPart{
							ClassGroupId: "bg-size",
						},
						"none": ClassPart{
							ClassGroupId: "bg-image",
						},
						"gradient": ClassPart{
							NextPart: map[string]ClassPart{
								"to": ClassPart{
									NextPart: map[string]ClassPart{
										"t": ClassPart{
											ClassGroupId: "bg-image",
										},
										"tr": ClassPart{
											ClassGroupId: "bg-image",
										},
										"r": ClassPart{
											ClassGroupId: "bg-image",
										},
										"br": ClassPart{
											ClassGroupId: "bg-image",
										},
										"b": ClassPart{
											ClassGroupId: "bg-image",
										},
										"bl": ClassPart{
											ClassGroupId: "bg-image",
										},
										"l": ClassPart{
											ClassGroupId: "bg-image",
										},
										"tl": ClassPart{
											ClassGroupId: "bg-image",
										},
									},
								},
							},
						},
						"blend": ClassPart{
							NextPart: map[string]ClassPart{
								"normal": ClassPart{
									ClassGroupId: "bg-blend",
								},
								"multiply": ClassPart{
									ClassGroupId: "bg-blend",
								},
								"screen": ClassPart{
									ClassGroupId: "bg-blend",
								},
								"overlay": ClassPart{
									ClassGroupId: "bg-blend",
								},
								"darken": ClassPart{
									ClassGroupId: "bg-blend",
								},
								"lighten": ClassPart{
									ClassGroupId: "bg-blend",
								},
								"color": ClassPart{
									NextPart: map[string]ClassPart{
										"dodge": ClassPart{
											ClassGroupId: "bg-blend",
										},
										"burn": ClassPart{
											ClassGroupId: "bg-blend",
										},
									},
								},
								"hard": ClassPart{
									NextPart: map[string]ClassPart{
										"light": ClassPart{
											ClassGroupId: "bg-blend",
										},
									},
								},
								"soft": ClassPart{
									NextPart: map[string]ClassPart{
										"light": ClassPart{
											ClassGroupId: "bg-blend",
										},
									},
								},
								"difference": ClassPart{
									ClassGroupId: "bg-blend",
								},
								"exclusion": ClassPart{
									ClassGroupId: "bg-blend",
								},
								"hue": ClassPart{
									ClassGroupId: "bg-blend",
								},
								"saturation": ClassPart{
									ClassGroupId: "bg-blend",
								},
								"luminosity": ClassPart{
									ClassGroupId: "bg-blend",
								},
								"plus": ClassPart{
									NextPart: map[string]ClassPart{
										"lighter": ClassPart{
											ClassGroupId: "bg-blend",
										},
									},
								},
							},
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsArbitraryPosition,
							ClassGroupId: "bg-position",
						},
						ClassGroupValidator{
							Fn:           IsArbitrarySize,
							ClassGroupId: "bg-size",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryImage,
							ClassGroupId: "bg-image",
						},
						ClassGroupValidator{
							Fn:           IsAny,
							ClassGroupId: "bg-color",
						},
					},
				},
				"from": ClassPart{
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsPercent,
							ClassGroupId: "gradient-from-pos",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryLength,
							ClassGroupId: "gradient-from-pos",
						},
						ClassGroupValidator{
							Fn:           IsAny,
							ClassGroupId: "gradient-from",
						},
					},
				},
				"via": ClassPart{
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsPercent,
							ClassGroupId: "gradient-via-pos",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryLength,
							ClassGroupId: "gradient-via-pos",
						},
						ClassGroupValidator{
							Fn:           IsAny,
							ClassGroupId: "gradient-via",
						},
					},
				},
				"to": ClassPart{
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsPercent,
							ClassGroupId: "gradient-to-pos",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryLength,
							ClassGroupId: "gradient-to-pos",
						},
						ClassGroupValidator{
							Fn:           IsAny,
							ClassGroupId: "gradient-to",
						},
					},
				},
				"rounded": ClassPart{
					NextPart: map[string]ClassPart{
						"none": ClassPart{
							ClassGroupId: "rounded",
						},
						"full": ClassPart{
							ClassGroupId: "rounded",
						},
						"s": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "rounded-s",
								},
								"full": ClassPart{
									ClassGroupId: "rounded-s",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsTshirtSize,
									ClassGroupId: "rounded-s",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "rounded-s",
								},
							},
							ClassGroupId: "rounded-s",
						},
						"e": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "rounded-e",
								},
								"full": ClassPart{
									ClassGroupId: "rounded-e",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsTshirtSize,
									ClassGroupId: "rounded-e",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "rounded-e",
								},
							},
							ClassGroupId: "rounded-e",
						},
						"t": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "rounded-t",
								},
								"full": ClassPart{
									ClassGroupId: "rounded-t",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsTshirtSize,
									ClassGroupId: "rounded-t",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "rounded-t",
								},
							},
							ClassGroupId: "rounded-t",
						},
						"r": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "rounded-r",
								},
								"full": ClassPart{
									ClassGroupId: "rounded-r",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsTshirtSize,
									ClassGroupId: "rounded-r",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "rounded-r",
								},
							},
							ClassGroupId: "rounded-r",
						},
						"b": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "rounded-b",
								},
								"full": ClassPart{
									ClassGroupId: "rounded-b",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsTshirtSize,
									ClassGroupId: "rounded-b",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "rounded-b",
								},
							},
							ClassGroupId: "rounded-b",
						},
						"l": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "rounded-l",
								},
								"full": ClassPart{
									ClassGroupId: "rounded-l",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsTshirtSize,
									ClassGroupId: "rounded-l",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "rounded-l",
								},
							},
							ClassGroupId: "rounded-l",
						},
						"ss": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "rounded-ss",
								},
								"full": ClassPart{
									ClassGroupId: "rounded-ss",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsTshirtSize,
									ClassGroupId: "rounded-ss",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "rounded-ss",
								},
							},
							ClassGroupId: "rounded-ss",
						},
						"se": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "rounded-se",
								},
								"full": ClassPart{
									ClassGroupId: "rounded-se",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsTshirtSize,
									ClassGroupId: "rounded-se",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "rounded-se",
								},
							},
							ClassGroupId: "rounded-se",
						},
						"ee": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "rounded-ee",
								},
								"full": ClassPart{
									ClassGroupId: "rounded-ee",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsTshirtSize,
									ClassGroupId: "rounded-ee",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "rounded-ee",
								},
							},
							ClassGroupId: "rounded-ee",
						},
						"es": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "rounded-es",
								},
								"full": ClassPart{
									ClassGroupId: "rounded-es",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsTshirtSize,
									ClassGroupId: "rounded-es",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "rounded-es",
								},
							},
							ClassGroupId: "rounded-es",
						},
						"tl": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "rounded-tl",
								},
								"full": ClassPart{
									ClassGroupId: "rounded-tl",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsTshirtSize,
									ClassGroupId: "rounded-tl",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "rounded-tl",
								},
							},
							ClassGroupId: "rounded-tl",
						},
						"tr": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "rounded-tr",
								},
								"full": ClassPart{
									ClassGroupId: "rounded-tr",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsTshirtSize,
									ClassGroupId: "rounded-tr",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "rounded-tr",
								},
							},
							ClassGroupId: "rounded-tr",
						},
						"br": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "rounded-br",
								},
								"full": ClassPart{
									ClassGroupId: "rounded-br",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsTshirtSize,
									ClassGroupId: "rounded-br",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "rounded-br",
								},
							},
							ClassGroupId: "rounded-br",
						},
						"bl": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "rounded-bl",
								},
								"full": ClassPart{
									ClassGroupId: "rounded-bl",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsTshirtSize,
									ClassGroupId: "rounded-bl",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "rounded-bl",
								},
							},
							ClassGroupId: "rounded-bl",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsTshirtSize,
							ClassGroupId: "rounded",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryValue,
							ClassGroupId: "rounded",
						},
					},
					ClassGroupId: "rounded",
				},
				"border": ClassPart{
					NextPart: map[string]ClassPart{
						"x": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsLength,
									ClassGroupId: "border-w-x",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryLength,
									ClassGroupId: "border-w-x",
								},
								ClassGroupValidator{
									Fn:           IsAny,
									ClassGroupId: "border-color-x",
								},
							},
							ClassGroupId: "border-w-x",
						},
						"y": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsLength,
									ClassGroupId: "border-w-y",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryLength,
									ClassGroupId: "border-w-y",
								},
								ClassGroupValidator{
									Fn:           IsAny,
									ClassGroupId: "border-color-y",
								},
							},
							ClassGroupId: "border-w-y",
						},
						"s": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsLength,
									ClassGroupId: "border-w-s",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryLength,
									ClassGroupId: "border-w-s",
								},
							},
							ClassGroupId: "border-w-s",
						},
						"e": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsLength,
									ClassGroupId: "border-w-e",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryLength,
									ClassGroupId: "border-w-e",
								},
							},
							ClassGroupId: "border-w-e",
						},
						"t": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsLength,
									ClassGroupId: "border-w-t",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryLength,
									ClassGroupId: "border-w-t",
								},
								ClassGroupValidator{
									Fn:           IsAny,
									ClassGroupId: "border-color-t",
								},
							},
							ClassGroupId: "border-w-t",
						},
						"r": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsLength,
									ClassGroupId: "border-w-r",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryLength,
									ClassGroupId: "border-w-r",
								},
								ClassGroupValidator{
									Fn:           IsAny,
									ClassGroupId: "border-color-r",
								},
							},
							ClassGroupId: "border-w-r",
						},
						"b": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsLength,
									ClassGroupId: "border-w-b",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryLength,
									ClassGroupId: "border-w-b",
								},
								ClassGroupValidator{
									Fn:           IsAny,
									ClassGroupId: "border-color-b",
								},
							},
							ClassGroupId: "border-w-b",
						},
						"l": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsLength,
									ClassGroupId: "border-w-l",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryLength,
									ClassGroupId: "border-w-l",
								},
								ClassGroupValidator{
									Fn:           IsAny,
									ClassGroupId: "border-color-l",
								},
							},
							ClassGroupId: "border-w-l",
						},
						"opacity": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsNumber,
									ClassGroupId: "border-opacity",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryNumber,
									ClassGroupId: "border-opacity",
								},
							},
							ClassGroupId: "border-opacity",
						},
						"solid": ClassPart{
							ClassGroupId: "border-style",
						},
						"dashed": ClassPart{
							ClassGroupId: "border-style",
						},
						"dotted": ClassPart{
							ClassGroupId: "border-style",
						},
						"double": ClassPart{
							ClassGroupId: "border-style",
						},
						"none": ClassPart{
							ClassGroupId: "border-style",
						},
						"hidden": ClassPart{
							ClassGroupId: "border-style",
						},
						"collapse": ClassPart{
							ClassGroupId: "border-collapse",
						},
						"separate": ClassPart{
							ClassGroupId: "border-collapse",
						},
						"spacing": ClassPart{
							NextPart: map[string]ClassPart{
								"x": ClassPart{
									Validators: []ClassGroupValidator{
										ClassGroupValidator{
											Fn:           IsArbitraryValue,
											ClassGroupId: "border-spacing-x",
										},
										ClassGroupValidator{
											Fn:           IsLength,
											ClassGroupId: "border-spacing-x",
										},
										ClassGroupValidator{
											Fn:           IsArbitraryLength,
											ClassGroupId: "border-spacing-x",
										},
									},
								},
								"y": ClassPart{
									Validators: []ClassGroupValidator{
										ClassGroupValidator{
											Fn:           IsArbitraryValue,
											ClassGroupId: "border-spacing-y",
										},
										ClassGroupValidator{
											Fn:           IsLength,
											ClassGroupId: "border-spacing-y",
										},
										ClassGroupValidator{
											Fn:           IsArbitraryLength,
											ClassGroupId: "border-spacing-y",
										},
									},
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "border-spacing",
								},
								ClassGroupValidator{
									Fn:           IsLength,
									ClassGroupId: "border-spacing",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryLength,
									ClassGroupId: "border-spacing",
								},
							},
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsLength,
							ClassGroupId: "border-w",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryLength,
							ClassGroupId: "border-w",
						},
						ClassGroupValidator{
							Fn:           IsAny,
							ClassGroupId: "border-color",
						},
					},
					ClassGroupId: "border-w",
				},
				"divide": ClassPart{
					NextPart: map[string]ClassPart{
						"x": ClassPart{
							NextPart: map[string]ClassPart{
								"reverse": ClassPart{
									ClassGroupId: "divide-x-reverse",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsLength,
									ClassGroupId: "divide-x",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryLength,
									ClassGroupId: "divide-x",
								},
							},
							ClassGroupId: "divide-x",
						},
						"y": ClassPart{
							NextPart: map[string]ClassPart{
								"reverse": ClassPart{
									ClassGroupId: "divide-y-reverse",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsLength,
									ClassGroupId: "divide-y",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryLength,
									ClassGroupId: "divide-y",
								},
							},
							ClassGroupId: "divide-y",
						},
						"opacity": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsNumber,
									ClassGroupId: "divide-opacity",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryNumber,
									ClassGroupId: "divide-opacity",
								},
							},
						},
						"solid": ClassPart{
							ClassGroupId: "divide-style",
						},
						"dashed": ClassPart{
							ClassGroupId: "divide-style",
						},
						"dotted": ClassPart{
							ClassGroupId: "divide-style",
						},
						"double": ClassPart{
							ClassGroupId: "divide-style",
						},
						"none": ClassPart{
							ClassGroupId: "divide-style",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsAny,
							ClassGroupId: "divide-color",
						},
					},
				},
				"outline": ClassPart{
					NextPart: map[string]ClassPart{
						"solid": ClassPart{
							ClassGroupId: "outline-style",
						},
						"dashed": ClassPart{
							ClassGroupId: "outline-style",
						},
						"dotted": ClassPart{
							ClassGroupId: "outline-style",
						},
						"double": ClassPart{
							ClassGroupId: "outline-style",
						},
						"none": ClassPart{
							ClassGroupId: "outline-style",
						},
						"offset": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsLength,
									ClassGroupId: "outline-offset",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "outline-offset",
								},
							},
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsLength,
							ClassGroupId: "outline-w",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryLength,
							ClassGroupId: "outline-w",
						},
						ClassGroupValidator{
							Fn:           IsAny,
							ClassGroupId: "outline-color",
						},
					},
					ClassGroupId: "outline-style",
				},
				"ring": ClassPart{
					NextPart: map[string]ClassPart{
						"inset": ClassPart{
							ClassGroupId: "ring-w-inset",
						},
						"opacity": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsNumber,
									ClassGroupId: "ring-opacity",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryNumber,
									ClassGroupId: "ring-opacity",
								},
							},
						},
						"offset": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsLength,
									ClassGroupId: "ring-offset-w",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryLength,
									ClassGroupId: "ring-offset-w",
								},
								ClassGroupValidator{
									Fn:           IsAny,
									ClassGroupId: "ring-offset-color",
								},
							},
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsLength,
							ClassGroupId: "ring-w",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryLength,
							ClassGroupId: "ring-w",
						},
						ClassGroupValidator{
							Fn:           IsAny,
							ClassGroupId: "ring-color",
						},
					},
					ClassGroupId: "ring-w",
				},

				"shadow": ClassPart{
					NextPart: map[string]ClassPart{
						"inner": ClassPart{
							ClassGroupId: "shadow",
						},
						"none": ClassPart{
							ClassGroupId: "shadow",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsTshirtSize,
							ClassGroupId: "shadow",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryShadow,
							ClassGroupId: "shadow",
						},
						ClassGroupValidator{
							Fn:           IsAny,
							ClassGroupId: "shadow-color",
						},
					},
					ClassGroupId: "shadow",
				},
				"opacity": ClassPart{
					NextPart: map[string]ClassPart{},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsNumber,
							ClassGroupId: "opacity",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryNumber,
							ClassGroupId: "opacity",
						},
					},
				},

				"mix": ClassPart{
					NextPart: map[string]ClassPart{
						"blend": ClassPart{
							NextPart: map[string]ClassPart{
								"normal": ClassPart{
									ClassGroupId: "mix-blend",
								},
								"multiply": ClassPart{
									ClassGroupId: "mix-blend",
								},
								"screen": ClassPart{
									ClassGroupId: "mix-blend",
								},
								"overlay": ClassPart{
									ClassGroupId: "mix-blend",
								},
								"darken": ClassPart{
									ClassGroupId: "mix-blend",
								},
								"lighten": ClassPart{
									ClassGroupId: "mix-blend",
								},
								"color": ClassPart{
									NextPart: map[string]ClassPart{
										"dodge": ClassPart{
											ClassGroupId: "mix-blend",
										},
										"burn": ClassPart{
											ClassGroupId: "mix-blend",
										},
									},
									ClassGroupId: "mix-blend",
								},
								"hard": ClassPart{
									NextPart: map[string]ClassPart{
										"light": ClassPart{
											ClassGroupId: "mix-blend",
										},
									},
								},
								"soft": ClassPart{
									NextPart: map[string]ClassPart{
										"light": ClassPart{
											ClassGroupId: "mix-blend",
										},
									},
								},
								"difference": ClassPart{
									ClassGroupId: "mix-blend",
								},
								"exclusion": ClassPart{
									ClassGroupId: "mix-blend",
								},
								"hue": ClassPart{
									ClassGroupId: "mix-blend",
								},
								"saturation": ClassPart{
									ClassGroupId: "mix-blend",
								},
								"luminosity": ClassPart{
									ClassGroupId: "mix-blend",
								},
								"plus": ClassPart{
									NextPart: map[string]ClassPart{
										"lighter": ClassPart{
											ClassGroupId: "mix-blend",
										},
									},
								},
							},
						},
					},
				},
				"filter": ClassPart{
					NextPart: map[string]ClassPart{
						"none": ClassPart{
							ClassGroupId: "filter",
						},
					},
					ClassGroupId: "filter",
				},
				"blur": ClassPart{
					NextPart: map[string]ClassPart{
						"none": ClassPart{
							ClassGroupId: "blur",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsTshirtSize,
							ClassGroupId: "blur",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryValue,
							ClassGroupId: "blur",
						},
					},
					ClassGroupId: "blur",
				},

				"brightness": ClassPart{
					NextPart: map[string]ClassPart{},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsNumber,
							ClassGroupId: "brightness",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryNumber,
							ClassGroupId: "brightness",
						},
					},
					ClassGroupId: "brightness",
				},

				"contrast": ClassPart{
					NextPart: map[string]ClassPart{},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsNumber,
							ClassGroupId: "contrast",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryNumber,
							ClassGroupId: "contrast",
						},
					},
					ClassGroupId: "contrast",
				},

				"drop": ClassPart{
					NextPart: map[string]ClassPart{
						"shadow": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "drop-shadow",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsTshirtSize,
									ClassGroupId: "drop-shadow",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "drop-shadow",
								},
							},
							ClassGroupId: "drop-shadow",
						},
					},
				},

				"grayscale": ClassPart{
					NextPart: map[string]ClassPart{
						"0": ClassPart{
							ClassGroupId: "grayscale",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsArbitraryValue,
							ClassGroupId: "grayscale",
						},
					},
					ClassGroupId: "grayscale",
				},

				"hue": ClassPart{
					NextPart: map[string]ClassPart{
						"rotate": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsNumber,
									ClassGroupId: "hue-rotate",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "hue-rotate",
								},
							},
						},
					},
				},
				"invert": ClassPart{
					NextPart: map[string]ClassPart{
						"0": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "invert",
								},
							},
							ClassGroupId: "invert",
						},
					},
					ClassGroupId: "invert",
				},

				"saturate": ClassPart{
					NextPart: map[string]ClassPart{},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsNumber,
							ClassGroupId: "saturate",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryNumber,
							ClassGroupId: "saturate",
						},
					},
					ClassGroupId: "saturate",
				},

				"sepia": ClassPart{
					NextPart: map[string]ClassPart{
						"0": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "sepia",
								},
							},
							ClassGroupId: "sepia",
						},
					},
					ClassGroupId: "sepia",
				},

				"backdrop": ClassPart{
					NextPart: map[string]ClassPart{
						"filter": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "backdrop-filter",
								},
							},
							ClassGroupId: "backdrop-filter",
						},
						"blur": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "backdrop-blur",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsTshirtSize,
									ClassGroupId: "backdrop-blur",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "backdrop-blur",
								},
							},
							ClassGroupId: "backdrop-blur",
						},
						"brightness": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsNumber,
									ClassGroupId: "backdrop-brightness",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryNumber,
									ClassGroupId: "backdrop-brightness",
								},
							},
							ClassGroupId: "backdrop-brightness",
						},
						"contrast": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsNumber,
									ClassGroupId: "backdrop-contrast",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryNumber,
									ClassGroupId: "backdrop-contrast",
								},
							},
							ClassGroupId: "backdrop-contrast",
						},
						"grayscale": ClassPart{
							NextPart: map[string]ClassPart{
								"0": ClassPart{
									ClassGroupId: "backdrop-grayscale",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "backdrop-grayscale",
								},
							},
							ClassGroupId: "backdrop-grayscale",
						},
						"hue": ClassPart{
							NextPart: map[string]ClassPart{
								"rotate": ClassPart{
									Validators: []ClassGroupValidator{
										ClassGroupValidator{
											Fn:           IsNumber,
											ClassGroupId: "backdrop-hue-rotate",
										},
										ClassGroupValidator{
											Fn:           IsArbitraryValue,
											ClassGroupId: "backdrop-hue-rotate",
										},
									},
								},
							},
							Validators: []ClassGroupValidator{},
						},
						"invert": ClassPart{
							NextPart: map[string]ClassPart{
								"0": ClassPart{
									Validators: []ClassGroupValidator{
										ClassGroupValidator{
											Fn:           IsArbitraryValue,
											ClassGroupId: "backdrop-invert",
										},
									},
									ClassGroupId: "backdrop-invert",
								},
							},
							Validators:   []ClassGroupValidator{},
							ClassGroupId: "backdrop-invert",
						},
						"opacity": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsNumber,
									ClassGroupId: "backdrop-opacity",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryNumber,
									ClassGroupId: "backdrop-opacity",
								},
							},
							ClassGroupId: "backdrop-opacity",
						},
						"saturate": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsNumber,
									ClassGroupId: "backdrop-saturate",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryNumber,
									ClassGroupId: "backdrop-saturate",
								},
							},
							ClassGroupId: "backdrop-saturate",
						},
						"sepia": ClassPart{
							NextPart: map[string]ClassPart{
								"0": ClassPart{
									Validators: []ClassGroupValidator{
										ClassGroupValidator{
											Fn:           IsArbitraryValue,
											ClassGroupId: "backdrop-sepia",
										},
									},
									ClassGroupId: "backdrop-sepia",
								},
							},
							ClassGroupId: "backdrop-sepia",
						},
					},
				},
				"caption": ClassPart{
					NextPart: map[string]ClassPart{
						"top": ClassPart{
							ClassGroupId: "caption",
						},
						"bottom": ClassPart{
							ClassGroupId: "caption",
						},
					},
					Validators: []ClassGroupValidator{},
				},
				"transition": ClassPart{
					NextPart: map[string]ClassPart{
						"none": ClassPart{
							ClassGroupId: "transition",
						},
						"all": ClassPart{
							ClassGroupId: "transition",
						},
						"colors": ClassPart{
							ClassGroupId: "transition",
						},
						"opacity": ClassPart{
							ClassGroupId: "transition",
						},
						"shadow": ClassPart{
							ClassGroupId: "transition",
						},
						"transform": ClassPart{
							ClassGroupId: "transition",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsArbitraryValue,
							ClassGroupId: "transition",
						},
					},
					ClassGroupId: "transition",
				},
				"duration": ClassPart{
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsNumber,
							ClassGroupId: "duration",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryValue,
							ClassGroupId: "duration",
						},
					},
				},
				"ease": ClassPart{
					NextPart: map[string]ClassPart{
						"linear": ClassPart{
							ClassGroupId: "ease",
						},
						"in": ClassPart{
							NextPart: map[string]ClassPart{
								"out": ClassPart{
									ClassGroupId: "ease",
								},
							},
							ClassGroupId: "ease",
						},
						"out": ClassPart{
							ClassGroupId: "ease",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsArbitraryValue,
							ClassGroupId: "ease",
						},
					},
				},
				"delay": ClassPart{
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsNumber,
							ClassGroupId: "delay",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryValue,
							ClassGroupId: "delay",
						},
					},
				},
				"animate": ClassPart{
					NextPart: map[string]ClassPart{
						"none": ClassPart{
							ClassGroupId: "animate",
						},
						"spin": ClassPart{
							ClassGroupId: "animate",
						},
						"ping": ClassPart{
							ClassGroupId: "animate",
						},
						"pulse": ClassPart{
							ClassGroupId: "animate",
						},
						"bounce": ClassPart{
							ClassGroupId: "animate",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsArbitraryValue,
							ClassGroupId: "animate",
						},
					},
				},
				"transform": ClassPart{
					NextPart: map[string]ClassPart{
						"gpu": ClassPart{
							ClassGroupId: "transform",
						},
						"none": ClassPart{
							ClassGroupId: "transform",
						},
					},
					ClassGroupId: "transform",
				},
				"scale": ClassPart{
					NextPart: map[string]ClassPart{
						"x": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsNumber,
									ClassGroupId: "scale-x",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryNumber,
									ClassGroupId: "scale-x",
								},
							},
						},
						"y": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsNumber,
									ClassGroupId: "scale-y",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryNumber,
									ClassGroupId: "scale-y",
								},
							},
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsNumber,
							ClassGroupId: "scale",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryNumber,
							ClassGroupId: "scale",
						},
					},
				},
				"rotate": ClassPart{
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsInteger,
							ClassGroupId: "rotate",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryValue,
							ClassGroupId: "rotate",
						},
					},
				},
				"translate": ClassPart{
					NextPart: map[string]ClassPart{
						"x": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "translate-x",
								},
								ClassGroupValidator{
									Fn:           IsLength,
									ClassGroupId: "translate-x",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryLength,
									ClassGroupId: "translate-x",
								},
							},
						},
						"y": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "translate-y",
								},
								ClassGroupValidator{
									Fn:           IsLength,
									ClassGroupId: "translate-y",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryLength,
									ClassGroupId: "translate-y",
								},
							},
						},
					},
				},
				"skew": ClassPart{
					NextPart: map[string]ClassPart{
						"x": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsNumber,
									ClassGroupId: "skew-x",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "skew-x",
								},
							},
						},
						"y": ClassPart{
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsNumber,
									ClassGroupId: "skew-y",
								},
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "skew-y",
								},
							},
						},
					},
				},
				"origin": ClassPart{
					NextPart: map[string]ClassPart{
						"center": ClassPart{
							ClassGroupId: "transform-origin",
						},
						"top": ClassPart{
							NextPart: map[string]ClassPart{
								"right": ClassPart{
									ClassGroupId: "transform-origin",
								},
								"left": ClassPart{
									ClassGroupId: "transform-origin",
								},
							},
							ClassGroupId: "transform-origin",
						},
						"right": ClassPart{
							ClassGroupId: "transform-origin",
						},
						"bottom": ClassPart{
							NextPart: map[string]ClassPart{
								"right": ClassPart{
									ClassGroupId: "transform-origin",
								},
								"left": ClassPart{
									ClassGroupId: "transform-origin",
								},
							},
							ClassGroupId: "transform-origin",
						},
						"left": ClassPart{
							ClassGroupId: "transform-origin",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsArbitraryValue,
							ClassGroupId: "transform-origin",
						},
					},
				},
				"accent": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": ClassPart{
							NextPart:     map[string]ClassPart{},
							Validators:   []ClassGroupValidator{},
							ClassGroupId: "accent",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsAny,
							ClassGroupId: "accent",
						},
					},
				},
				"appearance": ClassPart{
					NextPart: map[string]ClassPart{
						"none": ClassPart{
							ClassGroupId: "appearance",
						},
						"auto": ClassPart{
							ClassGroupId: "appearance",
						},
					},
				},
				"cursor": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": ClassPart{
							ClassGroupId: "cursor",
						},
						"default": ClassPart{
							ClassGroupId: "cursor",
						},
						"pointer": ClassPart{
							ClassGroupId: "cursor",
						},
						"wait": ClassPart{
							ClassGroupId: "cursor",
						},
						"text": ClassPart{
							ClassGroupId: "cursor",
						},
						"move": ClassPart{
							ClassGroupId: "cursor",
						},
						"help": ClassPart{
							ClassGroupId: "cursor",
						},
						"not": ClassPart{
							NextPart: map[string]ClassPart{
								"allowed": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
						"none": ClassPart{
							ClassGroupId: "cursor",
						},
						"context": ClassPart{
							NextPart: map[string]ClassPart{
								"menu": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
						"progress": ClassPart{
							ClassGroupId: "cursor",
						},
						"cell": ClassPart{
							ClassGroupId: "cursor",
						},
						"crosshair": ClassPart{
							ClassGroupId: "cursor",
						},
						"vertical": ClassPart{
							NextPart: map[string]ClassPart{
								"text": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
						"alias": ClassPart{
							ClassGroupId: "cursor",
						},
						"copy": ClassPart{
							ClassGroupId: "cursor",
						},
						"no": ClassPart{
							NextPart: map[string]ClassPart{
								"drop": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
						"grab": ClassPart{
							ClassGroupId: "cursor",
						},
						"grabbing": ClassPart{
							ClassGroupId: "cursor",
						},
						"all": ClassPart{
							NextPart: map[string]ClassPart{
								"scroll": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
						"col": ClassPart{
							NextPart: map[string]ClassPart{
								"resize": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
						"row": ClassPart{
							NextPart: map[string]ClassPart{
								"resize": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
						"n": ClassPart{
							NextPart: map[string]ClassPart{
								"resize": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
						"e": ClassPart{
							NextPart: map[string]ClassPart{
								"resize": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
						"s": ClassPart{
							NextPart: map[string]ClassPart{
								"resize": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
						"w": ClassPart{
							NextPart: map[string]ClassPart{
								"resize": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
						"ne": ClassPart{
							NextPart: map[string]ClassPart{
								"resize": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
						"nw": ClassPart{
							NextPart: map[string]ClassPart{
								"resize": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
						"se": ClassPart{
							NextPart: map[string]ClassPart{
								"resize": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
						"sw": ClassPart{
							NextPart: map[string]ClassPart{
								"resize": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
						"ew": ClassPart{
							NextPart: map[string]ClassPart{
								"resize": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
						"ns": ClassPart{
							NextPart: map[string]ClassPart{
								"resize": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
						"nesw": ClassPart{
							NextPart: map[string]ClassPart{
								"resize": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
						"nwse": ClassPart{
							NextPart: map[string]ClassPart{
								"resize": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
						"zoom": ClassPart{
							NextPart: map[string]ClassPart{
								"in": ClassPart{
									ClassGroupId: "cursor",
								},
								"out": ClassPart{
									ClassGroupId: "cursor",
								},
							},
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsArbitraryValue,
							ClassGroupId: "cursor",
						},
					},
				},
				"caret": ClassPart{
					NextPart: map[string]ClassPart{},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsAny,
							ClassGroupId: "caret-color",
						},
					},
				},
				"pointer": ClassPart{
					NextPart: map[string]ClassPart{
						"events": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "pointer-events",
								},
								"auto": ClassPart{
									ClassGroupId: "pointer-events",
								},
							},
						},
					},
				},
				"resize": ClassPart{
					NextPart: map[string]ClassPart{
						"none": ClassPart{
							ClassGroupId: "resize",
						},
						"y": ClassPart{
							ClassGroupId: "resize",
						},
						"x": ClassPart{
							ClassGroupId: "resize",
						},
					},
					ClassGroupId: "resize",
				},
				"scroll": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": ClassPart{
							ClassGroupId: "scroll-behavior",
						},
						"smooth": ClassPart{
							ClassGroupId: "scroll-behavior",
						},
						"m": ClassPart{
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "scroll-m",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "scroll-m",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "scroll-m",
								},
							},
						},
						"mx": ClassPart{
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "scroll-mx",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "scroll-mx",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "scroll-mx",
								},
							},
						},
						"my": ClassPart{
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "scroll-my",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "scroll-my",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "scroll-my",
								},
							},
						},
						"ms": ClassPart{
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "scroll-ms",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "scroll-ms",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "scroll-ms",
								},
							},
						},
						"me": ClassPart{
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "scroll-me",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "scroll-me",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "scroll-me",
								},
							},
						},
						"mt": ClassPart{
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "scroll-mt",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "scroll-mt",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "scroll-mt",
								},
							},
						},
						"mr": ClassPart{
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "scroll-mr",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "scroll-mr",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "scroll-mr",
								},
							},
						},
						"mb": ClassPart{
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "scroll-mb",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "scroll-mb",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "scroll-mb",
								},
							},
						},
						"ml": ClassPart{
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "scroll-ml",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "scroll-ml",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "scroll-ml",
								},
							},
						},
						"p": ClassPart{
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "scroll-p",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "scroll-p",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "scroll-p",
								},
							},
						},
						"px": ClassPart{
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "scroll-px",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "scroll-px",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "scroll-px",
								},
							},
						},
						"py": ClassPart{
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "scroll-py",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "scroll-py",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "scroll-py",
								},
							},
						},
						"ps": ClassPart{
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "scroll-ps",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "scroll-ps",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "scroll-ps",
								},
							},
						},
						"pe": ClassPart{
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "scroll-pe",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "scroll-pe",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "scroll-pe",
								},
							},
						},
						"pt": ClassPart{
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "scroll-pt",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "scroll-pt",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "scroll-pt",
								},
							},
						},
						"pr": ClassPart{
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "scroll-pr",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "scroll-pr",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "scroll-pr",
								},
							},
						},
						"pb": ClassPart{
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "scroll-pb",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "scroll-pb",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "scroll-pb",
								},
							},
						},
						"pl": ClassPart{
							Validators: []ClassGroupValidator{
								{
									Fn:           IsArbitraryValue,
									ClassGroupId: "scroll-pl",
								},
								{
									Fn:           IsLength,
									ClassGroupId: "scroll-pl",
								},
								{
									Fn:           IsArbitraryLength,
									ClassGroupId: "scroll-pl",
								},
							},
						},
					},
				},
				"snap": ClassPart{
					NextPart: map[string]ClassPart{
						"start": ClassPart{
							ClassGroupId: "snap-align",
						},
						"end": ClassPart{
							ClassGroupId: "snap-align",
						},
						"center": ClassPart{
							ClassGroupId: "snap-align",
						},
						"align": ClassPart{
							NextPart: map[string]ClassPart{
								"none": ClassPart{
									ClassGroupId: "snap-align",
								},
							},
						},
						"normal": ClassPart{
							ClassGroupId: "snap-stop",
						},
						"always": ClassPart{
							ClassGroupId: "snap-stop",
						},
						"none": ClassPart{
							ClassGroupId: "snap-type",
						},
						"x": ClassPart{
							ClassGroupId: "snap-type",
						},
						"y": ClassPart{
							ClassGroupId: "snap-type",
						},
						"both": ClassPart{
							ClassGroupId: "snap-type",
						},
						"mandatory": ClassPart{
							ClassGroupId: "snap-strictness",
						},
						"proximity": ClassPart{
							ClassGroupId: "snap-strictness",
						},
					},
				},
				"touch": ClassPart{
					NextPart: map[string]ClassPart{
						"auto": ClassPart{
							ClassGroupId: "touch",
						},
						"none": ClassPart{
							ClassGroupId: "touch",
						},
						"manipulation": ClassPart{
							ClassGroupId: "touch",
						},
						"pan": ClassPart{
							NextPart: map[string]ClassPart{
								"x": ClassPart{
									ClassGroupId: "touch-x",
								},
								"left": ClassPart{
									ClassGroupId: "touch-x",
								},
								"right": ClassPart{
									ClassGroupId: "touch-x",
								},
								"y": ClassPart{
									ClassGroupId: "touch-y",
								},
								"up": ClassPart{
									ClassGroupId: "touch-y",
								},
								"down": ClassPart{
									ClassGroupId: "touch-y",
								},
							},
						},
						"pinch": ClassPart{
							NextPart: map[string]ClassPart{
								"zoom": ClassPart{
									ClassGroupId: "touch-pz",
								},
							},
						},
					},
				},
				"select": ClassPart{
					NextPart: map[string]ClassPart{
						"none": ClassPart{
							ClassGroupId: "select",
						},
						"text": ClassPart{
							ClassGroupId: "select",
						},
						"all": ClassPart{
							ClassGroupId: "select",
						},
						"auto": ClassPart{
							ClassGroupId: "select",
						},
					},
					Validators: []ClassGroupValidator{},
				},
				"will": ClassPart{
					NextPart: map[string]ClassPart{
						"change": ClassPart{
							NextPart: map[string]ClassPart{
								"auto": ClassPart{
									ClassGroupId: "will-change",
								},
								"scroll": ClassPart{
									ClassGroupId: "will-change",
								},
								"contents": ClassPart{
									ClassGroupId: "will-change",
								},
								"transform": ClassPart{
									ClassGroupId: "will-change",
								},
							},
							Validators: []ClassGroupValidator{
								ClassGroupValidator{
									Fn:           IsArbitraryValue,
									ClassGroupId: "will-change",
								},
							},
						},
					},
				},
				"fill": ClassPart{
					NextPart: map[string]ClassPart{
						"none": ClassPart{
							ClassGroupId: "fill",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsAny,
							ClassGroupId: "fill",
						},
					},
				},
				"stroke": ClassPart{
					NextPart: map[string]ClassPart{
						"none": ClassPart{
							ClassGroupId: "stroke",
						},
					},
					Validators: []ClassGroupValidator{
						ClassGroupValidator{
							Fn:           IsLength,
							ClassGroupId: "stroke-w",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryLength,
							ClassGroupId: "stroke-w",
						},
						ClassGroupValidator{
							Fn:           IsArbitraryNumber,
							ClassGroupId: "stroke-w",
						},
						ClassGroupValidator{
							Fn:           IsAny,
							ClassGroupId: "stroke",
						},
					},
				},
				"sr": ClassPart{
					NextPart: map[string]ClassPart{
						"only": ClassPart{
							ClassGroupId: "sr",
						},
					},
				},
				"forced": ClassPart{
					NextPart: map[string]ClassPart{
						"color": ClassPart{
							NextPart: map[string]ClassPart{
								"adjust": ClassPart{
									NextPart: map[string]ClassPart{
										"auto": ClassPart{
											ClassGroupId: "forced-color-adjust",
										},
										"none": ClassPart{
											ClassGroupId: "forced-color-adjust",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
