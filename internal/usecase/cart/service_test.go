package cart

import (
	"reflect"
	"testing"

	"github.com/julioshinoda/cart/entity"
	"github.com/julioshinoda/cart/mocks"
)

func TestService_AddItem(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		ID   string
		item entity.Items
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Cart
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				repo: &mocks.Repository{},
			},
			args: args{
				ID: "id",
				item: entity.Items{
					ID:       "some-id",
					Price:    12025,
					Quantity: 1,
				},
			},
			want: &entity.Cart{
				ID: "id",
				Items: []entity.Items{
					{
						ID:       "some-id",
						Price:    12025,
						Quantity: 1,
					},
				},
				Total:    12025,
				Subtotal: 12025,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: tt.fields.repo,
			}

			switch tt.name {
			case "success":
				s.repo.(*mocks.Repository).On("Get", tt.args.ID).Return(nil, entity.ErrNotFound)
				s.repo.(*mocks.Repository).On("Update", tt.want).Return(nil)
			}
			got, err := s.AddItem(tt.args.ID, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
