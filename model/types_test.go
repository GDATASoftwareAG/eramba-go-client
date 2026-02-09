package model_test

import (
	"testing"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func Test_ErambaViewLink(t *testing.T) {
	type args struct {
		base string
		tool string
		id   int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "third party",
			args: args{
				base: "https://test.com",
				tool: "third-party",
				id:   10,
			},
			want: "https://test.com/third-party/view/ThirdParty/10?sort%5Bcreated%5D=desc",
		},
		{
			name: "security-policy-reviews",
			args: args{
				base: "https://test.com",
				tool: "security-policy-reviews",
				id:   10,
			},
			want: "https://test.com/security-policy-reviews/view/SecurityPolicyReviews/10?sort%5Bcreated%5D=desc",
		},
		{
			name: "risks",
			args: args{
				base: "https://test.com",
				tool: "risks",
				id:   10,
			},
			want: "https://test.com/risks/view/Risks/10?sort%5Bcreated%5D=desc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := model.ErambaViewLink(tt.args.base, tt.args.tool, tt.args.id); got != tt.want {
				t.Errorf("ErambaViewLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
