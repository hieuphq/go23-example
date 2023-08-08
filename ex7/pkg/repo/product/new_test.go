package product

import (
	"testing"

	"github.com/dwarvesf/go23/ex7/pkg/model"
	"github.com/dwarvesf/go23/ex7/pkg/repo/testutil"
)

func Test_pgRepo_CreateProduct(t *testing.T) {
	db := testutil.NewDB()

	type args struct {
		p model.Product
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Product
		wantErr bool
	}{
		{
			name: "create product 1",
			args: args{
				p: model.Product{
					Name:     "Create Product 1",
					SKU:      "create-SKU-1",
					Price:    100000000.00,
					Currency: "VND",
				},
			},
			want: &model.Product{
				Name:     "Create Product 1",
				SKU:      "create-SKU-1",
				Price:    100000000.00,
				Currency: "VND",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := pgRepo{
				DB: db,
			}
			got, err := r.CreateProduct(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgRepo.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if got.SKU != tt.want.SKU {
					t.Errorf("pgRepo.CreateProduct() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
