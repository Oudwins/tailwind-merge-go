package twmerge

import (
	"sort"
	"strings"
	"testing"
)

func TestTailwindMerge(t *testing.T) {
	tt := []struct {
		in  string
		out string
	}{
		// handles arbitrary property conflicts correctly
		{
			in:  "[paint-order:markers] [paint-order:normal]",
			out: "[paint-order:normal]",
		},
		// handles arbitrary property conflicts with modifiers correctly
		{
			in:  "[paint-order:markers] hover:[paint-order:normal]",
			out: "[paint-order:markers] hover:[paint-order:normal]",
		}, {
			in:  "hover:[paint-order:markers] hover:[paint-order:normal]",
			out: "hover:[paint-order:normal]",
		}, {
			in:  "hover:focus:[paint-order:markers] focus:hover:[paint-order:normal]",
			out: "focus:hover:[paint-order:normal]",
		},
		// handles complex arbitrary property conflicts correctly
		{
			in:  "[-unknown-prop:::123:::] [-unknown-prop:url(https://hi.com)]",
			out: "[-unknown-prop:url(https://hi.com)]",
		},
		// handles important modifier correctly
		{
			in:  "![some:prop] [some:other]",
			out: "![some:prop] [some:other]",
		}, {
			in:  "![some:prop] [some:other] [some:one] ![some:another]",
			out: "[some:one] ![some:another]",
		},
		// handles simple conflicts with arbitrary values correctly
		{
			in:  "m-[2px] m-[10px]",
			out: "m-[10px]",
		}, {
			in:  "z-20 z-[99]",
			out: "z-[99]",
		}, {
			in:  "my-[2px] m-[10rem]",
			out: "m-[10rem]",
		}, {
			in:  "cursor-pointer cursor-[grab]",
			out: "cursor-[grab]",
		}, {
			in:  "m-[2px] m-[calc(100%-var(--arbitrary))]",
			out: "m-[calc(100%-var(--arbitrary))]",
		}, {
			in:  "m-[2px] m-[length:var(--mystery-var)]",
			out: "m-[length:var(--mystery-var)]",
		}, {
			in:  "opacity-10 opacity-[0.025]",
			out: "opacity-[0.025]",
		}, {
			in:  "scale-75 scale-[1.7]",
			out: "scale-[1.7]",
		}, {
			in:  "brightness-90 brightness-[1.75]",
			out: "brightness-[1.75]",
		}, {
			in:  "min-h-[0.5px] min-h-[0]",
			out: "min-h-[0]",
		}, {
			in:  "text-[0.5px] text-[color:0]",
			out: "text-[0.5px] text-[color:0]",
		}, {
			in:  "text-[0.5px] text-[--my-0]",
			out: "text-[0.5px] text-[--my-0]",
		},
		// handles arbitrary length conflicts with labels and modifiers correctly
		{
			in:  "hover:m-[2px] hover:m-[length:var(--c)]",
			out: "hover:m-[length:var(--c)]",
		}, {
			in:  "hover:focus:m-[2px] focus:hover:m-[length:var(--c)]",
			out: "focus:hover:m-[length:var(--c)]",
		}, {
			in:  "border-b border-[color:rgb(var(--color-gray-500-rgb)/50%))]",
			out: "border-b border-[color:rgb(var(--color-gray-500-rgb)/50%))]",
		}, {
			in:  "border-[color:rgb(var(--color-gray-500-rgb)/50%))] border-b",
			out: "border-[color:rgb(var(--color-gray-500-rgb)/50%))] border-b",
		},
		// handles complex arbitrary value conflicts correctly
		{
			in:  "grid-rows-[1fr,auto] grid-rows-2",
			out: "grid-rows-2",
		}, {
			in:  "grid-rows-[repeat(20,minmax(0,1fr))] grid-rows-3",
			out: "grid-rows-3",
		},
		// handles ambiguous arbitrary values correctly
		{
			in:  "mt-2 mt-[calc(theme(fontSize.4xl)/1.125)]",
			out: "mt-[calc(theme(fontSize.4xl)/1.125)]",
		}, {
			in:  "p-2 p-[calc(theme(fontSize.4xl)/1.125)_10px]",
			out: "p-[calc(theme(fontSize.4xl)/1.125)_10px]",
		}, {
			in:  "mt-2 mt-[length:theme(someScale.someValue)]",
			out: "mt-[length:theme(someScale.someValue)]",
		}, {
			in:  "mt-2 mt-[theme(someScale.someValue)]",
			out: "mt-[theme(someScale.someValue)]",
		}, {
			in:  "text-2xl text-[length:theme(someScale.someValue)]",
			out: "text-[length:theme(someScale.someValue)]",
		}, {
			in:  "text-2xl text-[calc(theme(fontSize.4xl)/1.125)]",
			out: "text-[calc(theme(fontSize.4xl)/1.125)]",
		}, {
			in:  "bg-cover bg-[percentage:30%] bg-[length:200px_100px]",
			out: "bg-[length:200px_100px]",
		},
		// basic arbitrary variants
		{
			in:  "[&>*]:underline [&>*]:line-through",
			out: "[&>*]:line-through",
		}, {
			in:  "[&>*]:underline [&>*]:line-through [&_div]:line-through",
			out: "[&>*]:line-through [&_div]:line-through",
		}, {
			in:  "supports-[display:grid]:flex supports-[display:grid]:grid",
			out: "supports-[display:grid]:grid",
		},
		// arbitrary variants with modifiers
		{
			in:  "dark:lg:hover:[&>*]:underline dark:lg:hover:[&>*]:line-through",
			out: "dark:lg:hover:[&>*]:line-through",
		}, {
			in:  "dark:lg:hover:[&>*]:underline dark:hover:lg:[&>*]:line-through",
			out: "dark:hover:lg:[&>*]:line-through",
		}, {
			in:  "hover:[&>*]:underline [&>*]:hover:line-through",
			out: "hover:[&>*]:underline [&>*]:hover:line-through",
		},
		// arbitrary variants with attribute selectors
		{
			in:  "[&[data-open]]:underline [&[data-open]]:line-through",
			out: "[&[data-open]]:line-through",
		},
		// multiple arbitrary variants
		{
			in:  "[&>*]:[&_div]:underline [&>*]:[&_div]:line-through",
			out: "[&>*]:[&_div]:line-through",
		}, {
			in:  "[&>*]:[&_div]:underline [&_div]:[&>*]:line-through",
			out: "[&>*]:[&_div]:underline [&_div]:[&>*]:line-through",
		},
		// arbitrary variants with arbitrary properties
		{
			in:  "[&>*]:[color:red] [&>*]:[color:blue]",
			out: "[&>*]:[color:blue]",
		},
		// merges classes from same group correctly
		{
			in:  "overflow-x-auto overflow-x-hidden",
			out: "overflow-x-hidden",
		}, {
			in:  "basis-full basis-auto",
			out: "basis-auto",
		}, {
			in:  "w-full w-fit",
			out: "w-fit",
		}, {
			in:  "overflow-x-auto overflow-x-hidden overflow-x-scroll",
			out: "overflow-x-scroll",
		}, {
			in:  "overflow-x-auto hover:overflow-x-hidden overflow-x-scroll",
			out: "hover:overflow-x-hidden overflow-x-scroll",
		}, {
			in:  "col-span-1 col-span-full",
			out: "col-span-full",
		},
		// merges classes from Font Variant Numeric section correctly
		{
			in:  "lining-nums tabular-nums diagonal-fractions",
			out: "lining-nums tabular-nums diagonal-fractions",
		}, {
			in:  "normal-nums tabular-nums diagonal-fractions",
			out: "tabular-nums diagonal-fractions",
		}, {
			in:  "tabular-nums diagonal-fractions normal-nums",
			out: "normal-nums",
		}, {
			in:  "tabular-nums proportional-nums",
			out: "proportional-nums",
		},
		// handles color conflicts properly
		{
			in:  "bg-grey-5 bg-hotpink",
			out: "bg-hotpink",
		}, {
			in:  "hover:bg-grey-5 hover:bg-hotpink",
			out: "hover:bg-hotpink",
		}, {
			in:  "stroke-[hsl(350_80%_0%)] stroke-[10px]",
			out: "stroke-[hsl(350_80%_0%)] stroke-[10px]",
		},
		// handles conflicts across class groups correctly
		{
			in:  "inset-1 inset-x-1",
			out: "inset-1 inset-x-1",
		}, {
			in:  "inset-x-1 inset-1",
			out: "inset-1",
		}, {
			in:  "inset-x-1 left-1 inset-1",
			out: "inset-1",
		}, {
			in:  "inset-x-1 inset-1 left-1",
			out: "inset-1 left-1",
		}, {
			in:  "inset-x-1 right-1 inset-1",
			out: "inset-1",
		}, {
			in:  "inset-x-1 right-1 inset-x-1",
			out: "inset-x-1",
		}, {
			in:  "inset-x-1 right-1 inset-y-1",
			out: "inset-x-1 right-1 inset-y-1",
		}, {
			in:  "right-1 inset-x-1 inset-y-1",
			out: "inset-x-1 inset-y-1",
		}, {
			in:  "inset-x-1 hover:left-1 inset-1",
			out: "hover:left-1 inset-1",
		},
		// ring and shadow classes do not create conflict
		{
			in:  "ring shadow",
			out: "ring shadow",
		}, {
			in:  "ring-2 shadow-md",
			out: "ring-2 shadow-md",
		}, {
			in:  "shadow ring",
			out: "shadow ring",
		}, {
			in:  "shadow-md ring-2",
			out: "shadow-md ring-2",
		},
		// touch classes do create conflicts correctly
		{
			in:  "touch-pan-x touch-pan-right",
			out: "touch-pan-right",
		}, {
			in:  "touch-none touch-pan-x",
			out: "touch-pan-x",
		}, {
			in:  "touch-pan-x touch-none",
			out: "touch-none",
		}, {
			in:  "touch-pan-x touch-pan-y touch-pinch-zoom",
			out: "touch-pan-x touch-pan-y touch-pinch-zoom",
		}, {
			in:  "touch-manipulation touch-pan-x touch-pan-y touch-pinch-zoom",
			out: "touch-pan-x touch-pan-y touch-pinch-zoom",
		}, {
			in:  "touch-pan-x touch-pan-y touch-pinch-zoom touch-auto",
			out: "touch-auto",
		},
		// line-clamp classes do create conflicts correctly
		{
			in:  "overflow-auto inline line-clamp-1",
			out: "line-clamp-1",
		}, {
			in:  "line-clamp-1 overflow-auto inline",
			out: "line-clamp-1 overflow-auto inline",
		},
		// merges content utilities correctly
		{
			in:  "content-['hello'] content-[attr(data-content)]",
			out: "content-[attr(data-content)]",
		},
		// merges tailwind classes with important modifier correctly
		{
			in:  "!font-medium !font-bold",
			out: "!font-bold",
		}, {
			in:  "!font-medium !font-bold font-thin",
			out: "!font-bold font-thin",
		}, {
			in:  "!right-2 !-inset-x-px",
			out: "!-inset-x-px",
		}, {
			in:  "focus:!inline focus:!block",
			out: "focus:!block",
		},
		// conflicts across prefix modifiers
		{
			in:  "hover:block hover:inline",
			out: "hover:inline",
		}, {
			in:  "hover:block hover:focus:inline",
			out: "hover:block hover:focus:inline",
		}, {
			in:  "hover:block hover:focus:inline focus:hover:inline",
			out: "hover:block focus:hover:inline",
		}, {
			in:  "focus-within:inline focus-within:block",
			out: "focus-within:block",
		},
		// conflicts across postfix modifiers
		{
			in:  "text-lg/7 text-lg/8",
			out: "text-lg/8",
		}, {
			in:  "text-lg/none leading-9",
			out: "text-lg/none leading-9",
		}, {
			in:  "leading-9 text-lg/none",
			out: "text-lg/none",
		}, {
			in:  "w-full w-1/2",
			out: "w-1/2",
		},
		// handles negative value conflicts correctly
		{
			in:  "-m-2 -m-5",
			out: "-m-5",
		}, {
			in:  "-top-12 -top-2000",
			out: "-top-2000",
		},
		// handles conflicts between positive and negative values correctly
		{
			in:  "-m-2 m-auto",
			out: "m-auto",
		}, {
			in:  "top-12 -top-69",
			out: "-top-69",
		},
		// handles conflicts across groups with negative values correctly
		{
			in:  "-right-1 inset-x-1",
			out: "inset-x-1",
		}, {
			in:  "hover:focus:-right-1 focus:hover:inset-x-1",
			out: "focus:hover:inset-x-1",
		},
		// merges non-conflicting classes correctly
		{
			in:  "border-t border-white/10",
			out: "border-t border-white/10",
		}, {
			in:  "border-t border-white",
			out: "border-t border-white",
		}, {
			in:  "text-3.5xl text-black",
			out: "text-3.5xl text-black",
		},
		// does not alter non-tailwind classes
		{
			in:  "non-tailwind-class inline block",
			out: "non-tailwind-class block",
		}, {
			in:  "inline block inline-1",
			out: "block inline-1",
		}, {
			in:  "inline block i-inline",
			out: "block i-inline",
		}, {
			in:  "focus:inline focus:block focus:inline-1",
			out: "focus:block focus:inline-1",
		},
		// merges classes with per-side border colors correctly
		{
			in:  "border-t-some-blue border-t-other-blue",
			out: "border-t-other-blue",
		}, {
			in:  "border-t-some-blue border-some-blue",
			out: "border-some-blue",
		},
		// handles pseudo variants conflicts properly
		{
			in:  "empty:p-2 empty:p-3",
			out: "empty:p-3",
		}, {
			in:  "hover:empty:p-2 hover:empty:p-3",
			out: "hover:empty:p-3",
		}, {
			in:  "read-only:p-2 read-only:p-3",
			out: "read-only:p-3",
		},
		// handles pseudo variant group conflicts properly
		{
			in:  "group-empty:p-2 group-empty:p-3",
			out: "group-empty:p-3",
		}, {
			in:  "peer-empty:p-2 peer-empty:p-3",
			out: "peer-empty:p-3",
		}, {
			in:  "group-empty:p-2 peer-empty:p-3",
			out: "group-empty:p-2 peer-empty:p-3",
		}, {
			in:  "hover:group-empty:p-2 hover:group-empty:p-3",
			out: "hover:group-empty:p-3",
		}, {
			in:  "group-read-only:p-2 group-read-only:p-3",
			out: "group-read-only:p-3",
		},
		// merges standalone classes from same group correctly
		{
			in:  "inline block",
			out: "block",
		}, {
			in:  "hover:block hover:inline",
			out: "hover:inline",
		}, {
			in:  "hover:block hover:block",
			out: "hover:block",
		}, {
			in:  "inline hover:inline focus:inline hover:block hover:focus:block",
			out: "inline focus:inline hover:block hover:focus:block",
		}, {
			in:  "underline line-through",
			out: "line-through",
		}, {
			in:  "line-through no-underline",
			out: "no-underline",
		},
		// supports Tailwind CSS v3.3 features
		{
			in:  "text-red text-lg/7 text-lg/8",
			out: "text-red text-lg/8",
		}, {
			in:  "hyphens-auto hyphens-manual",
			out: "hyphens-manual",
		}, {
			in:  "from-0% from-red",
			out: "from-0% from-red",
		}, {
			in:  "caption-top caption-bottom",
			out: "caption-bottom",
		}, {
			in:  "line-clamp-2 line-clamp-none line-clamp-[10]",
			out: "line-clamp-[10]",
		}, {
			in:  "delay-150 delay-0 duration-150 duration-0",
			out: "delay-0 duration-0",
		}, {
			in:  "justify-normal justify-center justify-stretch",
			out: "justify-stretch",
		}, {
			in:  "content-normal content-center content-stretch",
			out: "content-stretch",
		}, {
			in:  "whitespace-nowrap whitespace-break-spaces",
			out: "whitespace-break-spaces",
		},
		// supports Tailwind CSS v3.4 features
		{
			in:  "h-svh h-dvh w-svw w-dvw",
			out: "h-dvh w-dvw",
		}, {
			in:  "text-wrap text-pretty",
			out: "text-pretty",
		}, {
			in:  "w-5 h-3 size-10 w-12",
			out: "size-10 w-12",
		}, {
			in:  "grid-cols-2 grid-cols-subgrid grid-rows-5 grid-rows-subgrid",
			out: "grid-cols-subgrid grid-rows-subgrid",
		}, {
			in:  "min-w-0 min-w-50 min-w-px max-w-0 max-w-50 max-w-px",
			out: "min-w-px max-w-px",
		}, {
			in:  "forced-color-adjust-none forced-color-adjust-auto",
			out: "forced-color-adjust-auto",
		}, {
			in:  "appearance-none appearance-auto",
			out: "appearance-auto",
		}, {
			in:  "float-start float-end clear-start clear-end",
			out: "float-end clear-end",
		}, {
			in:  "*:p-10 *:p-20 hover:*:p-10 hover:*:p-20",
			out: "*:p-20 hover:*:p-20",
		},
		// twMerge
		{
			in:  "mix-blend-normal mix-blend-multiply",
			out: "mix-blend-multiply",
		}, {
			in:  "h-10 h-min",
			out: "h-min",
		}, {
			in:  "stroke-black stroke-1",
			out: "stroke-black stroke-1",
		}, {
			in:  "stroke-2 stroke-[3]",
			out: "stroke-[3]",
		}, {
			in:  "outline-black outline-1",
			out: "outline-black outline-1",
		}, {
			in:  "grayscale-0 grayscale-[50%]",
			out: "grayscale-[50%]",
		}, {
			in:  "grow grow-[2]",
			out: "grow-[2]",
		},
		// {
		// 	in:  "grow', [null, false, [['grow-[2]']]]",
		// 	out: "grow-[2]",
		// },
		// CUSTOM TESTS
		{
			// test case where there is modifier & maybePostfix which causes maybePostfix to be beyond size of baseClass
			in:  "hover:bg-red-500/90",
			out: "hover:bg-red-500/90",
		},
		{
			in:  "group-has-[[data-sidebar=menu-action]]/menu-item:pr-8 group-has-[[data-sidebar=menu-action]]/menu-item:pr-6",
			out: "group-has-[[data-sidebar=menu-action]]/menu-item:pr-6",
		},
	}

	for _, tc := range tt {
		got := Merge(tc.in)
		if areStringsEqual(got, tc.out) == false {
			t.Errorf("twMerge failed -> | in: %v | %v != %v", tc.in, got, tc.out)
		} /* else {
			// t.Log("twMerge passed -> | in: ", tc.in, " | out: ", got, " | expected: ", tc.out)
		} */
	}
}

func areStringsEqual(s1, s2 string) bool {
	// Split each string into individual parts
	parts1 := strings.Split(s1, " ")
	parts2 := strings.Split(s2, " ")

	// Sort the parts
	sort.Strings(parts1)
	sort.Strings(parts2)

	// Compare the sorted parts
	return strings.Join(parts1, " ") == strings.Join(parts2, " ")
}
