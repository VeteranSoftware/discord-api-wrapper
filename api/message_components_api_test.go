/*
 * Copyright (c) 2022-2024. Veteran Software
 *
 *  Discord API Wrapper - A custom wrapper for the Discord REST API developed for a proprietary project.
 *
 *  This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
 *  License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License along with this program.
 *  If not, see <http://www.gnu.org/licenses/>.
 */

package api

import (
	"reflect"
	"testing"
)

const (
	quickBrownFox = "The quick brown fox jumps over the lazy dog"
	googleDotCom  = "https://google.com"
)

func TestComponentGetType(t *testing.T) {
	type fields struct {
		Type ComponentType
	}
	tests := []struct {
		name   string
		fields fields
		want   ComponentType
	}{
		{
			name:   "Action Row",
			fields: fields{Type: ComponentTypeActionRow},
			want:   ComponentTypeActionRow,
		},
		{
			name:   "Button",
			fields: fields{Type: ComponentTypeButton},
			want:   ComponentTypeButton,
		},
		{
			name:   "Select Menu",
			fields: fields{Type: ComponentTypeSelectMenu},
			want:   ComponentTypeSelectMenu,
		},
		{
			name:   "Text Input",
			fields: fields{Type: ComponentTypeTextInput},
			want:   ComponentTypeTextInput,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Type: tt.fields.Type,
			}
			if got := c.Type; got != tt.want {
				t.Errorf("Type = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentSetType(t *testing.T) {
	type fields struct {
		Type ComponentType
	}
	type args struct {
		t ComponentType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Component
	}{
		{
			name:   "Action Row",
			fields: fields{Type: ComponentTypeActionRow},
			args:   args{t: ComponentTypeActionRow},
			want: &Component{
				Type: ComponentTypeActionRow,
			},
		},
		{
			name:   "Button",
			fields: fields{Type: ComponentTypeButton},
			args:   args{t: ComponentTypeButton},
			want: &Component{
				Type: ComponentTypeButton,
			},
		},
		{
			name:   "Select Menu",
			fields: fields{Type: ComponentTypeSelectMenu},
			args:   args{t: ComponentTypeSelectMenu},
			want: &Component{
				Type: ComponentTypeSelectMenu,
			},
		},
		{
			name:   "Text Input",
			fields: fields{Type: ComponentTypeTextInput},
			args:   args{t: ComponentTypeTextInput},
			want: &Component{
				Type: ComponentTypeTextInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Type: tt.fields.Type,
			}
			if got := c.SetType(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentGetCustomID(t *testing.T) {
	type fields struct {
		CustomID string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Empty String Test",
			fields: fields{CustomID: ""},
			want:   "",
		},
		{
			name:   "Single Character Test",
			fields: fields{CustomID: "G"},
			want:   "G",
		},
		{
			name:   "Long String Test",
			fields: fields{CustomID: quickBrownFox},
			want:   quickBrownFox,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				CustomID: tt.fields.CustomID,
			}
			if got := c.CustomID; got != tt.want {
				t.Errorf("GetCustomID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentSetCustomID(t *testing.T) {
	type fields struct {
		CustomID string
	}
	type args struct {
		t string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Component
	}{
		{
			name:   "Empty String Test",
			fields: fields{CustomID: ""},
			args:   args{t: ""},
			want: &Component{
				CustomID: "",
			},
		},
		{
			name:   "Single Character Test",
			fields: fields{CustomID: "G"},
			args:   args{t: "G"},
			want: &Component{
				CustomID: "G",
			},
		},
		{
			name:   "Long String Test",
			fields: fields{CustomID: quickBrownFox},
			args:   args{t: quickBrownFox},
			want: &Component{
				CustomID: quickBrownFox,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				CustomID: tt.fields.CustomID,
			}
			if got := c.SetCustomID(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCustomID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentIsDisabled(t *testing.T) {
	type fields struct {
		Disabled bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Is Disabled",
			fields: fields{Disabled: true},
			want:   true,
		},
		{
			name:   "Is Not Disabled",
			fields: fields{Disabled: false},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Disabled: tt.fields.Disabled,
			}
			if got := c.IsDisabled(); got != tt.want {
				t.Errorf("IsDisabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentSetDisabled(t *testing.T) {
	type fields struct {
		Disabled bool
	}
	type args struct {
		d bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Component
	}{
		{
			name:   "Set Disabled",
			fields: fields{Disabled: true},
			args:   args{d: true},
			want: &Component{
				Disabled: true,
			},
		},
		{
			name:   "Set Enabled",
			fields: fields{Disabled: false},
			args:   args{d: false},
			want: &Component{
				Disabled: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Disabled: tt.fields.Disabled,
			}
			if got := c.SetDisabled(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDisabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentGetButtonStyle(t *testing.T) {
	type fields struct {
		Style any
	}
	tests := []struct {
		name   string
		fields fields
		want   ButtonStyle
	}{
		{
			name:   "Primary",
			fields: fields{Style: ButtonPrimary},
			want:   ButtonPrimary,
		},
		{
			name:   "Secondary",
			fields: fields{Style: ButtonSecondary},
			want:   ButtonSecondary,
		},
		{
			name:   "Success",
			fields: fields{Style: ButtonSuccess},
			want:   ButtonSuccess,
		},
		{
			name:   "Danger",
			fields: fields{Style: ButtonDanger},
			want:   ButtonDanger,
		},
		{
			name:   "Link",
			fields: fields{Style: ButtonLink},
			want:   ButtonLink,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Style: tt.fields.Style,
			}
			if got := c.Style.(ButtonStyle); got != tt.want {
				t.Errorf("GetButtonStyle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentSetButtonStyle(t *testing.T) {
	type fields struct {
		Style any
	}
	type args struct {
		s ButtonStyle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Component
	}{
		{
			name:   "Primary",
			fields: fields{Style: ButtonPrimary},
			args:   args{s: ButtonPrimary},
			want: &Component{
				Style: ButtonPrimary,
			},
		},
		{
			name:   "Secondary",
			fields: fields{Style: ButtonSecondary},
			args:   args{s: ButtonSecondary},
			want: &Component{
				Style: ButtonSecondary,
			},
		},
		{
			name:   "Success",
			fields: fields{Style: ButtonSuccess},
			args:   args{s: ButtonSuccess},
			want: &Component{
				Style: ButtonSuccess,
			},
		},
		{
			name:   "Danger",
			fields: fields{Style: ButtonDanger},
			args:   args{s: ButtonDanger},
			want: &Component{
				Style: ButtonDanger,
			},
		},
		{
			name:   "Link",
			fields: fields{Style: ButtonLink},
			args:   args{s: ButtonLink},
			want: &Component{
				Style: ButtonLink,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Style: tt.fields.Style,
			}
			if got := c.SetButtonStyle(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetButtonStyle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentGetTextInputStyle(t *testing.T) {
	type fields struct {
		Style any
	}
	tests := []struct {
		name   string
		fields fields
		want   TextInputStyle
	}{
		{
			name:   "Short",
			fields: fields{Style: TextInputShort},
			want:   TextInputShort,
		},
		{
			name:   "Paragraph",
			fields: fields{Style: TextInputParagraph},
			want:   TextInputParagraph,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Style: tt.fields.Style,
			}
			if got := c.Style.(TextInputStyle); got != tt.want {
				t.Errorf("Style.(TextInputStyle) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentSetTextInputStyle(t *testing.T) {
	type fields struct {
		Style any
	}
	type args struct {
		s TextInputStyle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Component
	}{
		{
			name:   "Short",
			fields: fields{Style: TextInputShort},
			args:   args{s: TextInputShort},
			want: &Component{
				Style: TextInputShort,
			},
		},
		{
			name:   "Paragraph",
			fields: fields{Style: TextInputParagraph},
			args:   args{s: TextInputParagraph},
			want: &Component{
				Style: TextInputParagraph,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Style: tt.fields.Style,
			}
			if got := c.SetTextInputStyle(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetButtonStyle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewComponent(t *testing.T) {
	tests := []struct {
		name string
		want *Component
	}{
		{
			name: "New Component",
			want: &Component{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewComponent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComponent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentGetEmoji(t *testing.T) {
	type fields struct {
		Emoji *Emoji
	}
	tests := []struct {
		name   string
		fields fields
		want   *Emoji
	}{
		{
			name: "Custom Emoji",
			fields: fields{Emoji: &Emoji{
				ID:       StringToSnowflake("941127649168871454"),
				Name:     "glitch",
				Animated: false,
			}},
			want: &Emoji{
				ID:       StringToSnowflake("941127649168871454"),
				Name:     "glitch",
				Animated: false,
			},
		},
		{
			name: "Unicode Emoji",
			fields: fields{Emoji: &Emoji{
				ID:       nil,
				Name:     "🔥",
				Animated: false,
			}},
			want: &Emoji{
				ID:       nil,
				Name:     "🔥",
				Animated: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Emoji: tt.fields.Emoji,
			}
			if got := c.Emoji; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Emoji = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentSetEmoji(t *testing.T) {
	type fields struct {
		Emoji *Emoji
	}
	type args struct {
		e *Emoji
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Component
	}{
		{
			name: "Custom Emoji",
			fields: fields{Emoji: &Emoji{
				ID:       StringToSnowflake("941127649168871454"),
				Name:     "glitch",
				Animated: false,
			}},
			args: args{e: &Emoji{
				ID:       StringToSnowflake("941127649168871454"),
				Name:     "glitch",
				Animated: false,
			}},
			want: &Component{
				Emoji: &Emoji{
					ID:       StringToSnowflake("941127649168871454"),
					Name:     "glitch",
					Animated: false,
				},
			},
		},
		{
			name: "Unicode Emoji",
			fields: fields{Emoji: &Emoji{
				ID:       nil,
				Name:     "🔥",
				Animated: false,
			}},
			args: args{e: &Emoji{
				ID:       nil,
				Name:     "🔥",
				Animated: false,
			}},
			want: &Component{
				Emoji: &Emoji{
					ID:       nil,
					Name:     "🔥",
					Animated: false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Emoji: tt.fields.Emoji,
			}
			if got := c.SetEmoji(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetEmoji() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentGetURL(t *testing.T) {
	type fields struct {
		URL string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "URL",
			fields: fields{URL: googleDotCom},
			want:   googleDotCom,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				URL: tt.fields.URL,
			}
			if got := c.URL; got != tt.want {
				t.Errorf("URL = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponentSetURL(t *testing.T) {
	type fields struct {
		URL string
	}
	type args struct {
		u string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Component
	}{
		{
			name:   "URL",
			fields: fields{URL: googleDotCom},
			args:   args{u: googleDotCom},
			want: &Component{
				URL: googleDotCom,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				URL: tt.fields.URL,
			}
			if got := c.SetURL(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
