package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseLinkDetails(t *testing.T) {
	tests := []struct {
		path     string
		width    int
		height   int
		imageUrl string
		fileName string
	}{
		{
			"/300/200/www.audubon.org/sites/default/files/a1_1902_16_barred-owl_sandra_rothenberg_kk.jpg",
			300,
			200,
			"www.audubon.org/sites/default/files/a1_1902_16_barred-owl_sandra_rothenberg_kk.jpg",
			"a1_1902_16_barred-owl_sandra_rothenberg_kk.jpg",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.path, func(t *testing.T) {
			t.Parallel()
			data, err := ParseLinkDetails(tc.path)
			require.NoError(t, err)
			require.Equal(t, tc.width, data.width)
			require.Equal(t, tc.height, data.height)
			require.Equal(t, tc.imageUrl, data.originImageURL)
			require.Equal(t, tc.fileName, data.originImageName)
		})
	}
}
