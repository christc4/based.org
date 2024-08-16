# Postage stuff

Compiled with pdfLaTeX

	\documentclass[a4paper, 18pt]{article}
	\usepackage{array}
	\usepackage[margin=5em]{geometry}
	\usepackage{parskip}
	\usepackage{titlesec}
	\usepackage{pdflscape}  % or lscape for DVI output
	\titleformat{\section}{\normalfont\fontsize{14}{22}\selectfont\bfseries}{\thesection}{1em}{\vspace*{1em}\MakeUppercase}
	\begin{document}
	\begin{landscape}
	\pagenumbering{gobble}
	\section*{To: \\\\ JOHN SMITH}
	\section*{Address: \\\\ Block 8 Lot 51, Flanders St, Phase 6 Peach Meadows, \\\\ 1776, Wormy, Worm City, \\\\ Sealand \\\\ Telephone: +74 585 585 5924, email: a@gmail.com}
	\newpage
	\section*{FROM:\\\\Paul Allen\\\\ ADDRESS: \\\\ Flat 984 Bauline Court, 68 Wyke Road, \\\\ E5 358, Manchester, \\\\ UNITED KINGDOM \\\\ TELEPHONE: +44 6578 785556, email: dawn@dawn.com}
	\end{landscape}
	\end{document}
