# Simple letter

	\documentclass[10pt]{letter} % 10pt font size default, 11pt and 12pt are also possible
	\usepackage{geometry} % Required for adjusting page dimensions
	%\longindentation=0pt % Un-commenting this line will push the closing "Sincerely," to the left of the page
	\geometry{
		paper=a4paper, % Change to letterpaper for US letter
		top=3cm, % Top margin
		bottom=1.5cm, % Bottom margin
		left=4.5cm, % Left margin
		right=4.5cm, % Right margin
		%showframe, % Uncomment to show how the type block is set on the page
	}
	\usepackage{microtype} % Improve justification
	\opening{\textbf{Dear Sir/Madam,}}
	I look forward to your reply ;)
	\vspace{2\parskip} % Extra whitespace for aesthetics
	\closing{Sincerely, \\Christos}
	\vspace{2\parskip} % Extra whitespace for aesthetics
	\end{letter}
	\end{document}

