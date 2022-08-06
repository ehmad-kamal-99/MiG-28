package server

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	domain "backend-svc-template"
	"backend-svc-template/middleware"
	"backend-svc-template/mocks"
)

type mockService struct {
	bs *mocks.MockBeerService
}

type args struct {
	ctx *gin.Context
	rr  *httptest.ResponseRecorder
}

func init() {
	binding.Validator = new(middleware.DefaultValidator)
}

func Test_beer_add(t *testing.T) {
	// represents successful input and output for add beer api call.
	nbrb := newBeerReq{
		Name:  "Lager",
		Brand: "Carlsberg",
	}
	nbi := &domain.Beer{
		Name:  "Lager",
		Brand: "Carlsberg",
	}
	nbo := &domain.Beer{
		ID:    "1122",
		Name:  "Lager",
		Brand: "Carlsberg",
	}

	// represents invalid input for add beer api call.
	ibrb := newBeerReq{
		Name:  "Lager",
		Brand: "",
	}

	ms := mockService{
		bs: mocks.NewMockBeerService(gomock.NewController(t)),
	}

	b := beer{
		bs: ms.bs,
	}

	tests := []struct {
		name        string
		args        args
		wantStatus  int
		mockService func(ms mockService)
	}{
		{
			name:       "success - add beer",
			args:       mkTestContext("POST", nbrb),
			wantStatus: 201,
			mockService: func(ms mockService) {
				ms.bs.EXPECT().Add(nbi).Return(nbo, nil)
			},
		},
		{
			name:        "failure - missing beer name",
			args:        mkTestContext("POST", ibrb),
			wantStatus:  400,
			mockService: func(ms mockService) {},
		},
		{
			name:        "failure - empty request body",
			args:        mkTestContext("POST", nil),
			wantStatus:  400,
			mockService: func(ms mockService) {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockService(ms)

			b.add(tt.args.ctx)

			assert.EqualValues(t, tt.wantStatus, tt.args.rr.Code)
		})
	}
}

func mkTestContext(method string, obj interface{}) args {
	rr := httptest.NewRecorder()

	data, _ := json.Marshal(obj)

	ctx, _ := gin.CreateTestContext(rr)

	ctx.Request, _ = http.NewRequestWithContext(context.TODO(), method, "", bytes.NewBuffer(data))
	ctx.Request.Header.Set("Content-Type", "application/json")

	return args{
		ctx: ctx,
		rr:  rr,
	}
}

func Test_beer_delete(t *testing.T) {
	type fields struct {
		bs BeerService
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &beer{
				bs: tt.fields.bs,
			}
			b.delete(tt.args.ctx)
		})
	}
}

func Test_beer_edit(t *testing.T) {
	type fields struct {
		bs BeerService
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &beer{
				bs: tt.fields.bs,
			}
			b.edit(tt.args.ctx)
		})
	}
}

func Test_beer_get(t *testing.T) {
	type fields struct {
		bs BeerService
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &beer{
				bs: tt.fields.bs,
			}
			b.get(tt.args.ctx)
		})
	}
}

func Test_beer_list(t *testing.T) {
	type fields struct {
		bs BeerService
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &beer{
				bs: tt.fields.bs,
			}
			b.list(tt.args.ctx)
		})
	}
}

func Test_review_add(t *testing.T) {
	type fields struct {
		rs ReviewService
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &review{
				rs: tt.fields.rs,
			}
			r.add(tt.args.ctx)
		})
	}
}

func Test_review_delete(t *testing.T) {
	type fields struct {
		rs ReviewService
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &review{
				rs: tt.fields.rs,
			}
			r.delete(tt.args.ctx)
		})
	}
}

func Test_review_list(t *testing.T) {
	type fields struct {
		rs ReviewService
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &review{
				rs: tt.fields.rs,
			}
			r.list(tt.args.ctx)
		})
	}
}
