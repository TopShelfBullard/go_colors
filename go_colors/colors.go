package colors

func Red(s string) string {
	return "\033[31m" + s + "\033[0m"
}

func Green(s string) string {
	return "\033[32m" + s + "\033[0m"
}

func Yellow(s string) string {
	return "\033[33m" + s + "\033[0m"
}

func Purple(s string) string {
	return "\033[34m" + s + "\033[0m"
}

func Violet(s string) string {
	return "\033[35m" + s + "\033[0m"
}

func Cyan(s string) string {
	return "\033[36m" + s + "\033[0m"
}

func BrightRed(s string) string {
	return "\033[91m" + s + "\033[0m"
}

func BrightGreen(s string) string {
	return "\033[92m" + s + "\033[0m"
}

func BrightYellow(s string) string {
	return "\033[93m" + s + "\033[0m"
}

func BrightPurple(s string) string {
	return "\033[94m" + s + "\033[0m"
}

func BrightViolet(s string) string {
	return "\033[95m" + s + "\033[0m"
}

func BrightCyan(s string) string {
	return "\033[96m" + s + "\033[0m"
}
