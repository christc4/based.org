# Groff me macros for writing formal letters

## Compilation

	$ groff -Tps -me example.me > example.ps
	$ ps2pdf example.ps example.pdf

## Source example

	.wh 2c hd \" Top margin (header) set to 2cm
	.wh -2c fo \" Bottom margin (footer) set to 2cm
	.po 2c \" Right margin set to 2cm
	.ll 17c \" Line length set to 17 cm
	.in 10c \" Address block indented 13 cm
	.sz 13 \" Font size
	John Smith,
	.br
	Buckingham Palace,
	.br
	London, SW1A 1AAU,
	.br
	United Kingdom
	.br
	john@smith.com
	.br
	+44 12345678900
	.sp
	.in 0 \" Address block indent removed
	10 Downing Street
	.br
	London, SW1A 2AA
	.br
	United Kingdom
	.br
	Dear Sir,

	Could you please move your car?

	Thank you.

	Best regards,
	.sp 6 \" Space left for signature
	John Smith
