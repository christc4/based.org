# Basic bond

compiled with XeLaTeX

	\documentclass[12pt,letterpaper,landscape] {letter}
	\usepackage{fontspec}
	\setmainfont{Arial}
	\thispagestyle{empty}
	\pagestyle{empty}
	\hoffset = 0pt
	\voffset = 0pt
	\oddsidemargin = 0pt
	\topmargin = 0pt
	\headheight = 0pt
	\headsep = 0pt
	\marginparsep = 0pt
	\marginparwidth = 0pt
	\footskip = 0pt
	\usepackage[pass,paperwidth=2.5in,paperheight=6.5in]{geometry}
	\usepackage[latin1]{inputenc}
	\usepackage{amsmath}
	\usepackage[tmargin=1cm]{geometry}
	\usepackage{tabto,ulem,booktabs}
	\usepackage{textcomp}
	\usepackage{amsfonts}
	\def\dotsign{\xleaders\hbox to .2em{\d{}}\hfill\d{}}
	\usepackage{amssymb}
	\begin{document}
	\begin{tabular}{ p{4.15in} l }
	Guatemala 05 de Septiembre del 2,012 & £10 \\ 
	\multicolumn{2}{ l}{Civillian Bond} \\
	\multicolumn{2}{ l }{Valid after 20/10/22} \\
	\end{tabular} 
	\makebox[.5\linewidth][r]{Signed: }\dotsign\smallskip\\
	\hspace*{.5\linewidth}Mr Christ Amlai\\
	\end{document}
