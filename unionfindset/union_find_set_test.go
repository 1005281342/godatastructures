package unionfindset

import (
	"reflect"
	"testing"
)

func TestConstructorDSU(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want DSU
	}{
		{name: "ok", args: args{n: 3}, want: DSU{0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConstructorDSU(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConstructorDSU() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDSU_Find(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		dsu  DSU
		args args
		want int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dsu.Find(tt.args.a); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDSU_Union(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		dsu  DSU
		args args
		want bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dsu.Union(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckRing(t *testing.T) {
	var (
		//存在环
		//edges = [][2]int{
		//	{0, 1}, {1, 2}, {1, 3},
		//	{2, 3}, {3, 4}, {2, 5},
		//}

		//不存在环
		edges = [][2]int{
			{0, 1}, {1, 2}, {1, 3},
			{3, 4}, {2, 5},
		}
		n = 6
	)

	var dsu = ConstructorDSU(n)

	for i := 0; i < len(edges); i++ {
		if !dsu.Union(edges[i][0], edges[i][1]) {
			t.Fatal("存在环")
		}
	}

	t.Log("不存在环")
}
