// go test -v . -run Testify                                                                                                                                                                     ─╯

package theory

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTestifyExample(t *testing.T) {
	// проверяет, что myCompare() возвращает true
	assert.True(t, true)

	// сравнивает числа, строки
	assert.Equal(t, 3+3, 6)

	// сравнивает два JSON-объекта
	assert.JSONEq(t, `{"name": "Alice", "role": "Admin"}`,
		`{"role": "Admin", "name": "Alice"}`)

	// проверяет, что объект не nil, "", false, 0 и что длина слайса не равна 0
	assert.NotEmpty(t, []int{1, 2, 3})

	// сравнивает элементы в двух массивах, если неважен их порядок
	assert.ElementsMatch(t, []int{1, 2, 3, 4}, []int{2, 3, 4, 1})

	// проверяет, что err равна nil
	assert.NoError(t, nil)
}

func TestTestifyPerson_FullName(t *testing.T) {
	type fields struct {
		FirstName string
		LastName  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "simple test",
			fields: fields{
				FirstName: "Misha",
				LastName:  "Popov",
			},
			want: "Misha Popov",
		},
		{
			name: "long name",
			fields: fields{
				FirstName: "Pablo Diego KHoze Frantsisko de Paula KHuan" +
					" Nepomukeno Krispin Krispiano de la Santisima Trinidad Ruiz",
				LastName: "Picasso",
			},
			want: "Pablo Diego KHoze Frantsisko de Paula KHuan Nepomukeno" +
				" Krispin Krispiano de la Santisima Trinidad Ruiz Picasso",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Person{
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
			}
			v := u.FullName()
			assert.Equal(t, tt.want, v)
		})
	}
}

func TestTestifyFamily_AddNew(t *testing.T) {
	type newPerson struct {
		r Relationship
		p Person
	}
	tests := []struct {
		name           string
		existedMembers map[Relationship]Person
		newPerson      newPerson
		wantErr        bool
	}{
		{
			name: "add father",
			existedMembers: map[Relationship]Person{
				Mother: {
					FirstName: "Maria",
					LastName:  "Popova",
					Age:       36,
				},
			},
			newPerson: newPerson{
				r: Father,
				p: Person{
					FirstName: "Misha",
					LastName:  "Popov",
					Age:       42,
				},
			},
			wantErr: false,
		},
		{
			name: "catch error",
			existedMembers: map[Relationship]Person{
				Father: {
					FirstName: "Misha",
					LastName:  "Popov",
					Age:       42,
				},
			},
			newPerson: newPerson{
				r: Father,
				p: Person{
					FirstName: "Ken",
					LastName:  "Gymsohn",
					Age:       32,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Family{
				Members: tt.existedMembers,
			}
			err := f.AddNew(tt.newPerson.r, tt.newPerson.p)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Contains(t, f.Members, tt.newPerson.r)
		})
	}
}
