package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCart_AddItem(t *testing.T) {
	type fields struct {
		ID          string
		Total       int
		Subtotal    int
		CouponCode  string
		CouponValue int
		CouponType  string
		Items       []Items
	}
	type args struct {
		item Items
	}
	tests := []struct {
		name   string
		fields fields
		args   args

		wantErr error
	}{
		{
			name: "Success add item",
			fields: fields{
				ID:    "aas",
				Items: []Items{},
			},
			args: args{
				item: Items{
					ID:       "new",
					Price:    10,
					Quantity: 1,
				},
			},
		},
		{
			name: "Error duplicated item",
			fields: fields{
				ID: "aas",
				Items: []Items{
					Items{
						ID: "exist",
					},
				},
			},
			args: args{
				item: Items{
					ID:       "exist",
					Price:    10,
					Quantity: 1,
				},
			},
			wantErr: ErrDuplicatedItem,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cart{
				ID:          tt.fields.ID,
				Total:       tt.fields.Total,
				Subtotal:    tt.fields.Subtotal,
				CouponCode:  tt.fields.CouponCode,
				CouponValue: tt.fields.CouponValue,
				CouponType:  tt.fields.CouponType,
				Items:       tt.fields.Items,
			}
			if err := c.AddItem(tt.args.item); err != nil && err != tt.wantErr {
				t.Errorf("Cart.AddItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCart_RemoveItem(t *testing.T) {
	type fields struct {
		ID          string
		Total       int
		Subtotal    int
		CouponCode  string
		CouponValue int
		CouponType  string
		Items       []Items
	}
	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		expect  []Items
		wantErr error
	}{
		{
			name: "success to remove",
			fields: fields{
				ID: "cart",
				Items: []Items{
					Items{
						ID: "notremove",
					},
					Items{
						ID: "remove",
					},
				},
			},
			args: args{
				ID: "remove",
			},
			expect: []Items{
				Items{
					ID: "notremove",
				},
			},
		},
		{
			name: "success to remove",
			fields: fields{
				ID: "cart",
				Items: []Items{
					Items{
						ID: "notremove",
					},
					Items{
						ID: "remove",
					},
				},
			},
			args: args{
				ID: "different",
			},
			expect: []Items{
				Items{
					ID: "notremove",
				},
				Items{
					ID: "remove",
				},
			},
			wantErr: ErrItemNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cart{
				ID:          tt.fields.ID,
				Total:       tt.fields.Total,
				Subtotal:    tt.fields.Subtotal,
				CouponCode:  tt.fields.CouponCode,
				CouponValue: tt.fields.CouponValue,
				CouponType:  tt.fields.CouponType,
				Items:       tt.fields.Items,
			}
			if err := c.RemoveItem(tt.args.ID); err != nil && err != tt.wantErr {
				t.Errorf("Cart.RemoveItem() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.expect, c.Items)
		})
	}
}
