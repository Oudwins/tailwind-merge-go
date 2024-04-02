package cache

type ICache interface {
	Get(string) string
	Set(string, string)
}
