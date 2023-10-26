package parser

import (
	"fmt"
	"testing"
)

func TestParseGPT3Response(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     *GPT3Response
		wantErr  bool
	}{
		{
			name: "file file",
			response: `yaml_file: (if applicable)
file_name: (if applicable)
command_to_run: test
explanation: test`,
			want: &GPT3Response{
				YamlFile:     "(if applicable)",
				FileName:     "(if applicable)",
				CommandToRun: "test",
				Explanation:  "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := ParseGPT3Response(tt.response)
			if (err != nil) != tt.wantErr {
				fmt.Println(err)
				t.Errorf("ParseGPT3Response() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.YamlFile != tt.want.YamlFile {
				t.Errorf("ParseGPT3Response() got = %v, want %v", got, tt.want)
			}
			if got.FileName != tt.want.FileName {
				t.Errorf("ParseGPT3Response() got = %v, want %v", got, tt.want)
			}
			if got.CommandToRun != tt.want.CommandToRun {
				t.Errorf("ParseGPT3Response() got = %v, want %v", got, tt.want)
			}
			if got.Explanation != tt.want.Explanation {
				t.Errorf("ParseGPT3Response() got = %v, want %v", got, tt.want)
			}
		})
	}
}
