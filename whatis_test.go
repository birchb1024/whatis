package whatis


WhatIs('42') // integer
WhatIs([3, 4, 5, 6, 7]) // A monotonically increasing integer sequence

WhatIs([	["Id", "Name"],
	[1, "Abe"],
	[2, "Alice"],
	[3, "Amy"]] ) // |A Table of peoples' names, sorted alpabetically
	WhatIs("http://www.example.com/") // A URL
	WhatIs("etoin.shrdlu@gmail.com") // An email address
	WhatIs("The computer is encredibly fast accurate and stupid, Stuart G. Walesh") // A quotation of Stuart G. Walesh
	WhatIs("La computadora es increíblemente rápida, precisa y estúpida.") // Spanish Text


