package main

import (
	"testing"
)

func TestBuildTree(t *testing.T) {
	tests := map[string]struct {
		values []int
		want   *Tree
	}{
		"empty tree": {
			values: []int{},
			want:   &Tree{},
		},
		"sub-4 tree": {
			values: []int{1, 2, 3, 4},
			want: &Tree{
				root: &Node{
					value: 1,
					right: &Node{
						value: 2,
						right: &Node{
							value: 3,
							right: &Node{
								value: 4,
							},
						},
					},
				},
			},
		},
		"unbalanced incrementing tree": {
			values: []int{1, 2, 3, 4, 5, 6},
			want: &Tree{
				root: &Node{
					value: 1,
					right: &Node{
						value: 2,
						right: &Node{
							value: 3,
							right: &Node{
								value: 4,
								right: &Node{
									value: 5,
									right: &Node{
										value: 6,
									},
								},
							},
						},
					},
				},
			},
		},
		"initially balanced tree": {
			values: []int{4, 2, 6, 1, 3, 5, 7, 8},
			want: &Tree{
				root: &Node{
					value: 4,
					right: &Node{
						value: 6,
						right: &Node{
							value: 7,
							right: &Node{
								value: 8,
							},
						},
						left: &Node{
							value: 5,
						},
					},
					left: &Node{
						value: 2,
						right: &Node{
							value: 3,
						},
						left: &Node{
							value: 1,
						},
					},
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := BuildTree(tt.values); !TreesAreEqual(got, tt.want) {
				t.Errorf("BuildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
