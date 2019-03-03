package database

import (
	"github.com/andrzejd-pl/SimpleRestBlogBackend/usage"
	"github.com/joho/godotenv"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetAllCategories(t *testing.T) {
	err := godotenv.Load("../.env")
	usage.CheckErr(err)
	tests := []struct {
		name string
		want []Category
	}{
		{
			"Category",
			[]Category{
				{1, "Programowanie", "Wpisy na temat programowania, projektów, itp."},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllCategories(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllCategories() = %v, want %v", got, tt.want)
			}
		})
	}
}
