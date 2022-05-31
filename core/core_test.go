package core

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	domain "backend-svc-template"
	"backend-svc-template/mocks"
)

type mockStorage struct {
	bs *mocks.MockBeerStorage
	rs *mocks.MockReviewStorage
}

func Test_beer_Add(t *testing.T) {
	ms := mockStorage{
		bs: mocks.NewMockBeerStorage(gomock.NewController(t)),
		rs: mocks.NewMockReviewStorage(gomock.NewController(t)),
	}

	beerSvc := NewBeer(ms.bs)

	type args struct {
		beer *domain.Beer
	}
	tests := []struct {
		name        string
		args        args
		want        *domain.Beer
		mockStorage func(storage mockStorage)
		wantErr     bool
	}{
		{
			name: "success - add beer",
			args: args{
				beer: &domain.Beer{
					Name:  "Lager",
					Brand: "Carlsberg",
				},
			},
			want: &domain.Beer{
				ID:    "1122",
				Name:  "Lager",
				Brand: "Carlsberg",
			},
			mockStorage: func(s mockStorage) {
				s.bs.EXPECT().Add(&domain.Beer{
					Name:  "Lager",
					Brand: "Carlsberg",
				}).Return("1122", nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockStorage(ms)

			got, err := beerSvc.Add(tt.args.beer)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_beer_Delete(t *testing.T) {
	type fields struct {
		bd BeerStorage
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &beer{
				bs: tt.fields.bd,
			}
			if err := b.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_beer_Edit(t *testing.T) {
	type fields struct {
		bd BeerStorage
	}
	type args struct {
		beer *domain.Beer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Beer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &beer{
				bs: tt.fields.bd,
			}
			got, err := b.Edit(tt.args.beer)
			if (err != nil) != tt.wantErr {
				t.Errorf("Edit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Edit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_beer_Get(t *testing.T) {
	type fields struct {
		bd BeerStorage
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Beer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &beer{
				bs: tt.fields.bd,
			}
			got, err := b.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_beer_List(t *testing.T) {
	type fields struct {
		bd BeerStorage
	}
	type args struct {
		filter []string
		sort   []string
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.Beer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &beer{
				bs: tt.fields.bd,
			}
			got, err := b.List(tt.args.filter, tt.args.sort, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_review_Add(t *testing.T) {
	type fields struct {
		rd ReviewStorage
	}
	type args struct {
		review *domain.Review
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Review
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &review{
				rd: tt.fields.rd,
			}
			got, err := rs.Add(tt.args.review)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_review_Delete(t *testing.T) {
	type fields struct {
		rd ReviewStorage
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &review{
				rd: tt.fields.rd,
			}
			if err := rs.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_review_List(t *testing.T) {
	type fields struct {
		rd ReviewStorage
	}
	type args struct {
		beerID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.Review
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &review{
				rd: tt.fields.rd,
			}
			got, err := rs.List(tt.args.beerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}
