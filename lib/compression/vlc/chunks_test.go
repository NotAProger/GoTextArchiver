package vlc

import (
	"reflect"
	"testing"
)

func Test_splitByChuncks(t *testing.T) {
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
			name: "base case",
			args: args{
				bStr:      "0010000000000100100101100101",
				chunkSize: 8,
			},
			want: BinaryChunks{"00100000", "00000100", "10010110", "01010000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChuncks(tt.args.bStr, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChuncks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_ToString(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want string
	}{
		{
			name: "base test",
			bcs:  BinaryChunks{"00101111", "10000000", "00000001"},
			want: "001011111000000000000001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.ToString(); got != tt.want {
				t.Errorf("BinaryChunks.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBinChunks(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want BinaryChunks
	}{
		{
			name: "bse test",
			data: []byte{20, 30, 60, 18, 1, 255},
			want: BinaryChunks{"00010100", "00011110", "00111100", "00010010", "00000001", "11111111"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinChunks(tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
