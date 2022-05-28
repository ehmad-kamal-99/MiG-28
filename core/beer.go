package core

// BeerStorage defines the interface required for handling beer storage related operations.
type BeerStorage interface {
	Add(beer *backend_svc_template.Beer) (string, error)
	Get(id string) (*backend_svc_template.Beer, error)
	Edit(beer *backend_svc_template.Beer) error
	List() ([]*backend_svc_template.Beer, error)
	Delete(id string) error
}

type beer struct {
	bd BeerStorage
}

func NewBeer(bd BeerStorage) *beer {
	return &beer{bd: bd}
}

func (b *beer) Add(beer *backend_svc_template.Beer) (*backend_svc_template.Beer, error) {
	id, err := b.bd.Add(beer)
	if err != nil {
		return nil, err
	}

	beer.ID = id

	return beer, nil
}

func (b *beer) Get(id string) (*backend_svc_template.Beer, error) {
	return b.bd.Get(id)
}

func (b *beer) Edit(beer *backend_svc_template.Beer) (*backend_svc_template.Beer, error) {
	if err := b.bd.Edit(beer); err != nil {
		return nil, err
	}

	return beer, nil
}

func (b *beer) List() ([]*backend_svc_template.Beer, error) {
	return b.bd.List()
}

func (b *beer) Delete(id string) error {
	return b.bd.Delete(id)
}
