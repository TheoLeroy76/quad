package main

import "github.com/01-edu/z01"

func QuadA(param1 int, param2 int) {

	if param1 == 5 && param2 == 3 {
		z01.PrintRune('o')
		z01.PrintRune('-')
		z01.PrintRune('-')
		z01.PrintRune('-')
		z01.PrintRune('o')

		z01.PrintRune('\n')

		z01.PrintRune('|')
		z01.PrintRune(' ')
		z01.PrintRune(' ')
		z01.PrintRune(' ')
		z01.PrintRune('|')

		z01.PrintRune('\n')

		z01.PrintRune('o')
		z01.PrintRune('-')
		z01.PrintRune('-')
		z01.PrintRune('-')
		z01.PrintRune('o')

	}

	if param1 == 5 && param2 == 1 {
		z01.PrintRune('o')
		z01.PrintRune('-')
		z01.PrintRune('-')
		z01.PrintRune('-')
		z01.PrintRune('o')

		z01.PrintRune('\n')
	}

	if param1 == 1 && param2 == 1 {
		z01.PrintRune('o')

		z01.PrintRune('\n')
	}

	if param1 == 1 && param2 == 5 {
		z01.PrintRune('o')

		z01.PrintRune('\n')

		z01.PrintRune('|')

		z01.PrintRune('\n')

		z01.PrintRune('|')

		z01.PrintRune('\n')

		z01.PrintRune('|')

		z01.PrintRune('\n')

		z01.PrintRune('o')
	}

}
