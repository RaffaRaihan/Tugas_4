package main

import (
	"fmt"
	"sort"
)

func main() {
	// Mendefinisikan map yang berisi lirik lagu
	List_Stanza := map[int]string{
	// Stanza I
    1:  "Stanza I",
    2:  "Indonesia, tanah airku",
    3:  "Tanah tumpah darahku",
    4:  "Di sanalah aku berdiri",
    5:  "/n",
    6:  "Jadi pandu ibuku",
    7:  "Indonesia, kebangsaanku",
    8:  "Bangsa dan tanah airku",
    9:  "Marilah kita berseru",
    10: "/n",
    11: "Indonesia bersatu",
    12: "Hiduplah tanahku",
    13: "Hiduplah negeriku",
    14: "Bangsaku, rakyatku, semuanya",
    15: "/n",
    16: "Bangunlah jiwanya",
    17: "Bangunlah badannya",
    18: "Untuk Indonesia Raya",
    19: "Indonesia raya",
    20: "/n",
    21: "Merdeka, merdeka",
    22: "Hiduplah Indonesia raya",
    23: "/dn",

    // Stanza II
    24: "Stanza II",
    25: "Indonesia, tanah yang mulia",
    26: "Tanah kita yang kaya",
    27: "Di sanalah aku berdiri",
    28: "Untuk selama-lamanya",
    29: "/n",
    30: "Indonesia, tanah pusaka",
    31: "Pusaka kita semuanya",
    32: "Marilah kita mendoa",
    33: "Indonesia bahagia!",
    34: "/n",
    35: "Suburlah tanahnya",
    36: "Suburlah jiwanya",
    37: "Bangsanya, rakyatnya, semuanya",
    38: "Sadarlah hatinya",
    39: "/n",
    40: "Sadarlah budinya",
    41: "Untuk Indonesia Raya.",
    42: "Indonesia raya",
    43: "Merdeka, merdeka",
    44: "Hiduplah Indonesia raya",
    45: "/dn",

    // Stanza III
    46: "Stanza III",
    47: "Indonesia, tanah yang mulia",
    48: "Tanah kita yang kaya",
    49: "Di sanalah aku berdiri",
    50: "Untuk selama-lamanya",
    51: "/n",
    52: "Indonesia, tanah pusaka",
    53: "Pusaka kita semuanya",
    54: "Marilah kita mendoa",
    55: "Indonesia bahagia!",
    56: "/n",
    57: "Suburlah tanahnya",
    58: "Suburlah jiwanya",
    59: "Bangsanya, rakyatnya, semuanya",
    60: "Sadarlah hatinya",
    61: "/n",
    62: "Sadarlah budinya",
    63: "Untuk Indonesia Raya.",
    64: "Indonesia raya",
    65: "Merdeka, merdeka",
    66: "Hiduplah Indonesia raya",
}

	// Mengurutkan kunci dari map
	keys := make([]int, 0, len(List_Stanza))
	for k := range List_Stanza {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// Menampilkan lirik berdasarkan urutan kunci
	for _, k := range keys {
		if List_Stanza[k] == "/n" {
			fmt.Println() // Cetak baris baru
		} else if List_Stanza[k] == "/dn" {
			fmt.Println() // Cetak baris baru
			fmt.Println() // Cetak baris baru lagi
		} else {
			fmt.Println(List_Stanza[k]) // Cetak lirik
		}
	}
}