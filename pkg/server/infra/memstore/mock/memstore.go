package mock

type MemStoreService struct {
	OnGet func(key string) ([]byte, error)
	OnAdd func(key string, value []byte, sec int) error
}

func (s *MemStoreService) Get(key string) ([]byte, error) {
	return s.OnGet(key)
}

func (s *MemStoreService) Add(key string, value []byte, sec int) error {
	return s.OnAdd(key, value, sec)
}
