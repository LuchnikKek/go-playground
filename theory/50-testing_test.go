//	"table driven test name-values-want": {
//		"prefix": "tdtn",
//		"body": [
//			"func Test$1(t *testing.T) {",
//				"\ttestCases := []struct {",
//					"\t\tname   string",
//					"\t\tvalues []$2",
//					"\t\twant   $3",
//				"\t}{",
//					"\t\t$4",
//				"\t}",
//				"\tfor _, tC := range testCases {",
//					"\t\tt.Run(tC.desc, func(t *testing.T) {",
//						"\t\t\tif res := $1(test.values...); res != test.want {",
//							"\t\t\t\tt.Errorf(\"$1() = %d, want %d\", res, test.want)",
//						"\t\t\t}",
//					"\t\t})",
//				"\t}",
//			"}"
//		]
//	},
//
//	"table driven test case (name-values-want)": {
//		"prefix": "tdtc",
//		"body": [
//			"{",
//			"\tname:   \"$1\",",
//			"\tvalues: []$2{$3},",
//			"\twant:   $4,",
//			"}$5"
//		]
//	},
package theory

import "testing"

func TestFamily_AddNew(t *testing.T) {
	type testValue struct {
		Family   *Family
		Relation Relationship
		Person   Person
	}
	testCases := []struct {
		name  string
		value testValue
		want  error
	}{
		{
			name: "add to empty",
			value: testValue{
				Family:   &Family{},
				Relation: Father,
				Person: Person{
					FirstName: "Father",
					LastName:  "Mishi",
					Age:       57,
				},
			},
			want: nil,
		},
		{
			name: "add same person",
			value: testValue{
				Family: &Family{
					Members: map[Relationship]Person{
						Father: {
							FirstName: "FatherOne",
							LastName:  "Mishi",
							Age:       52,
						},
					},
				},
				Relation: Father,
				Person: Person{
					FirstName: "FatherTwo",
					LastName:  "Mishi",
					Age:       57,
				},
			},
			want: ErrRelationshipAlreadyExists,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			if err := tC.value.Family.AddNew(tC.value.Relation, tC.value.Person); err != tC.want {
				t.Errorf("AddNew() = \"%s\", want err \"%s\"", err, tC.want)
			}
		})
	}
}

func TestFamily_AddNewEmpty(t *testing.T) {
	fam := Family{}
	father := Person{
		FirstName: "FatherTwo",
		LastName:  "Mishi",
		Age:       57,
	}
	if err := fam.AddNew(Father, father); err != nil {
		t.Errorf("got error %v", err)
	}
}

func TestFamily_AddNewExisting(t *testing.T) {
	fam := Family{
		Members: map[Relationship]Person{
			Father: {
				FirstName: "FatherOne",
				LastName:  "Mishi",
				Age:       52,
			},
		},
	}
	father := Person{
		FirstName: "FatherTwo",
		LastName:  "Mishi",
		Age:       57,
	}
	if err := fam.AddNew(Father, father); err != ErrRelationshipAlreadyExists {
		t.Errorf("expected ErrRelationshipAlreadyExists, got %v", err)
	}
}

func TestFamily_AddNewRefactored(t *testing.T) {
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
			if (err != nil) != tt.wantErr {
				t.Errorf("AddNew() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
