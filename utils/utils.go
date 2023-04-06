package utils

import (
	"fmt"
	"io"
	"math/rand"
	"t_chat/structs"
	"time"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func PrintRandomLogo(dest io.Writer) {
	switch seededRand.Intn(4) {
	case 0:
		fmt.Fprintf(dest, structs.Onion_logo, structs.ColorGreen, structs.ColorWhite,
			structs.ColorPurple, structs.ColorWhite,
			structs.ColorPurple, structs.ColorWhite,
			structs.ColorPurple, structs.ColorWhite,
			structs.ColorPurple, structs.ColorWhite,
			structs.ColorPurple, structs.ColorWhite,
			structs.ColorPurple, structs.ColorWhite,
			structs.ColorPurple, structs.ColorWhite,
			structs.ColorPurple, structs.ColorWhite,
			structs.ColorPurple, structs.ColorWhite,
			structs.ColorPurple, structs.ColorWhite,
			structs.ColorGreen, structs.ColorNone,
		)
	case 1:
		fmt.Fprintf(dest, structs.Logo_variant_1, structs.ColorLightBlue, structs.ColorNone)
	case 2:
		fmt.Fprintf(dest, structs.Logo_variant_2, structs.ColorPurple, structs.ColorNone)
	case 3:
		fmt.Fprintf(dest, structs.Logo_variant_3, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey,
			structs.ColorGreen, structs.ColorGrey,
			structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey,
			structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey,
			structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey,
			structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey,
			structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey,
			structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey,
			structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey, structs.ColorGreen, structs.ColorGrey,
			structs.ColorNone,
		)
	}
}
