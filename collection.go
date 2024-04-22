package rredis

type collectionType string

const (
	HASHMAP collectionType = "hashmap"
)

type collection interface {
	get(key string) (entry, bool)
	set(entry)
	delete(key string)
	exist(key string) bool
}

type entry struct {
	k string
	v string
}

/**
 * TODO: Optimize using BigCache technique
 */
type hashmap struct {
	entries map[string]entry
}

func newHashmap() *hashmap {
	return &hashmap{entries: make(map[string]entry)}
}

func (h *hashmap) get(key string) (entry, bool) {
	if entry, ok := h.entries[key]; ok {
		return entry, true
	}
	return *new(entry), false
}

func (h *hashmap) set(e entry) {
	h.entries[e.k] = e
}

func (h *hashmap) delete(key string) {
	delete(h.entries, key)
}

func (h *hashmap) exist(key string) bool {
	_, ok := h.entries[key]
	return ok
}
