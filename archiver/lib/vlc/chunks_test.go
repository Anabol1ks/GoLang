package vlc

import (
	"reflect"
	"testing"
)

func Test_splitByChunks(t *testing.T) {
	type args struct {
		bStr      string
		chunkSize int
	}

	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "base name",
			args: args{
				bStr:      "0010000000100010000100101010011011",
				chunkSize: 8,
			},
			want: BinaryChunks{"00100000", "00100010", "00010010", "10100110", "11000000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunks(tt.args.bStr, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_ToHex(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want HexChunks
	}{
		{
			name: "base test",
			bcs:  BinaryChunks{"0101111", "10000000"},
			want: HexChunks{"2F", "80"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.ToHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryChunks.ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHexChunks(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want HexChunks
	}{
		{
			name: "test hex",
			str:  "20 30 3C 18",
			want: HexChunks{"20", "30", "3C", "18"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHexChunks(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHexChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunk_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		hc   HexChunk
		want BinaryChunk
	}{
		{
			name: "tobinary test",
			hc:   HexChunk("2F"),
			want: BinaryChunk("00101111"),
		},
		{
			name: "tobinary test",
			hc:   HexChunk("80"),
			want: BinaryChunk("10000000"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hc.ToBinary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexChunk.ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunks_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		hcs  HexChunks
		want BinaryChunks
	}{
		{
			name: "base test",
			hcs:  HexChunks{"2F", "80"},
			want: BinaryChunks{"00101111", "10000000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hcs.ToBinary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexChunks.ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_Join(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want string
	}{
		{
			name: "base test",
			bcs:  BinaryChunks{"01001111", "10000000"},
			want: "0100111110000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.Join(); got != tt.want {
				t.Errorf("BinaryChunks.Join() = %v, want %v", got, tt.want)
			}
		})
	}
}
