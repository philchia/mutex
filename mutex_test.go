package mutex

import "testing"

func TestMutex_TryLock(t *testing.T) {
	type fields struct {
		state int32
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Lock a unlocked lock",
			fields: fields{
				state: 0,
			},
			want: true,
		},
		{
			name: "Lock a locked lock",
			fields: fields{
				state: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mutex{
				state: tt.fields.state,
			}
			if got := m.TryLock(); got != tt.want {
				t.Errorf("Mutex.TryLock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMutex_Lock(t *testing.T) {
	type fields struct {
		state int32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Lock a unlocked lock",
			fields: fields{
				state: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mutex{
				state: tt.fields.state,
			}
			m.Lock()
		})
	}
}

func TestMutex_Unlock(t *testing.T) {
	type fields struct {
		state int32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Unlock a locked lock",
			fields: fields{
				state: 1,
			},
		},
		{
			name: "Unlock a unlocked lock",
			fields: fields{
				state: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mutex{
				state: tt.fields.state,
			}
			m.Unlock()
		})
	}
}
