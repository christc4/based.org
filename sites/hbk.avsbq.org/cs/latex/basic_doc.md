# Basic doc

	\documentclass{article}
	\usepackage{geometry, fontspec, hyperref}
	\geometry{a4paper, left=20mm, right=20mm, top=10mm}
	\setmainfont{Arial}
	\makeatletter
	\renewcommand{\maketitle}{\bgroup\setlength{\parindent}{0pt}
	\begin{flushleft}
	  \textbf{\@title}\\
	  \@author\\
	    \@date
	\end{flushleft}\egroup
	}
	\makeatother
	\title{Wadham College written work}
	\author{Christ Amlai}
	\date{\today}
	\begin{document}
	\maketitle
	\fbox{\begin{minipage}{24em}
	    \textbf{Written Work 2: A brief summary of the changes to: cut, fit and fabric in mens suiting}
	\end{minipage}}
	\\\\
	The...
	\end{document}
