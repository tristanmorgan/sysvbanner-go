package banner_test

import (
	"bytes"
	"testing"

	"github.com/tristanmorgan/sysvbanner/banner"
)

type bufferWriter struct {
	buf []byte
}

func (writer *bufferWriter) Write(p []byte) (int, error) {
	writer.buf = append(writer.buf, p...)
	return len(p), nil
}

var table = []struct {
	name string
	code []string
}{
	{
		"ABCDEFGH", []string{
			" ###  ####   ###  ####  ##### #####  #### #   # ",
			"#   # #   # #   # #   # #     #     #     #   # ",
			"##### ####  #     #   # ###   ###   #  ## ##### ",
			"#   # #   # #     #   # #     #     #   # #   # ",
			"#   # #   # #     #   # #     #     #   # #   # ",
			"#   # #   # #   # #   # #     #     #   # #   # ",
			"#   # ####   ###  ####  ##### #      ###  #   # ",
			"                                                ",
			"",
			"",
		},
	}, {
		"IJKLMNOP", []string{
			" ###      # #   # #     #   # #   #  ###  ####  ",
			"  #       # #  #  #     ## ## ##  # #   # #   # ",
			"  #       # ###   #     # # # # # # #   # ####  ",
			"  #       # #  #  #     #   # #  ## #   # #     ",
			"  #       # #   # #     #   # #   # #   # #     ",
			"  #   #   # #   # #     #   # #   # #   # #     ",
			" ###   ###  #   # ##### #   # #   #  ###  #     ",
			"                                                ",
			"",
			"",
		},
	}, {
		"QRSTUVWX", []string{
			" ###  ####   #### ##### #   # #   # #   # #   # ",
			"#   # #   # #       #   #   # #   # #   #  # #  ",
			"#   # ####   ###    #   #   # #   # #   #   #   ",
			"#   # #   #     #   #   #   # #   # #   #  # #  ",
			"# # # #   #     #   #   #   #  # #  # # # #   # ",
			"#  #  #   # #   #   #   #   #  # #  ## ## #   # ",
			" ## # #   #  ###    #    ###    #   #   # #   # ",
			"                                                ",
			"",
			"",
		},
	}, {
		"YZabcdef", []string{
			"#   # #####       #               #          ## ",
			" # #      #       #               #         #   ",
			"  #      #   ###  # ##   ###   ## #  ###   #### ",
			"  #     #       # ##  # #   # #  ## #   #   #   ",
			"  #    #     #### #   # #     #   # #####   #   ",
			"  #   #     #   # #   # #   # #   # #       #   ",
			"  #   #####  #### ####   ###   ####  ####   #   ",
			"                                                ",
			"",
			"",
		},
	}, {
		"ghijklmn", []string{
			"      #       #       # #      ##               ",
			"      #                 #       #               ",
			" #### # ##   ##       # #  #    #   ## #  ####  ",
			"#   # ##  #   #       # # #     #   # # # #   # ",
			"#   # #   #   #       # ##      #   # # # #   # ",
			" #### #   #   #   #   # # #     #   #   # #   # ",
			"    # #   #    ## #   # #  #     ## #   # #   # ",
			"####               ###                          ",
			"",
			"",
		},
	}, {
		"opqrstuv", []string{
			"                                                ",
			"                                #               ",
			" ###  # ##   ## # # ##   ####  ###  #   # #   # ",
			"#   # ##  # #  ## ##  # #       #   #   # #   # ",
			"#   # #   # #   # #      ###    #   #   # #   # ",
			"#   # ####   #### #         #   #   #   #  # #  ",
			" ###  #         # #     ####     ##  ####   #   ",
			"      #         #                               ",
			"",
			"",
		},
	}, {
		"wxyz0123", []string{
			"                         ###    #    ###   ###  ",
			"                        #   #  ##   #   # #   # ",
			"#   # #   # #   # ##### #  ##   #       #     # ",
			"#   #  # #  #   #    #  # # #   #     ##    ##  ",
			"# # #   #   #   #   #   ##  #   #    #        # ",
			"# # #  # #   ####  #    #   #   #   #   # #   # ",
			" #### #   #     # #####  ###  ##### #####  ###  ",
			"            ####                                ",
			"",
			"",
		},
	}, {
		"56789[|]", []string{
			"#####   ##  #####  ###   ###   ###    #    ###  ",
			"#      #    #   # #   # #   #  #      #      #  ",
			"####  #         # #   # #   #  #      #      #  ",
			"    # ####     #   ###   ####  #      #      #  ",
			"    # #   #   #   #   #     #  #      #      #  ",
			"#   # #   #   #   #   #    #   #      #      #  ",
			" ###   ###    #    ###   ##    ###    #    ###  ",
			"                                                ",
			"",
			"",
		},
	},
}

func TestBanner(t *testing.T) {
	for _, v := range table {
		t.Run(v.name, func(t *testing.T) {
			writer := &bufferWriter{}
			banner.Banner(v.name, writer)

			for i, actual := range bytes.Split(writer.buf, []byte("\n")) {
				if !bytes.Equal([]byte(v.code[i]), actual) {
					t.Errorf("\nwant: '%v'\n got: '%v'\n", v.code[i], string(actual))
				}
			}
		})
	}
}
