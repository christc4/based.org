<small>Copied from http://suckless.org/rocks/</small>
<h1>Stuff that rocks</h1>
<p>Software on this page should have ideas similar to the suckless
<a href="//suckless.org/philosophy/">philosophy</a>. The software must be FOSS licensed.</p>
<h2>Libraries</h2>
<p>This section is for small, usable development libraries, which can be used for
writing software that sucks less. These should preferably be under the MIT/X
consortium or BSD licenses, WTFPL, or public domain, or alternatively LGPL,
because it makes them legally compatible with other suckless projects.</p>
<h3>Libc Implementations</h3>
<ul>
<li><a href="https://musl.libc.org/">musl</a> - standard C library that attempts to be even smaller than uClibc.
See also: <a href="https://www.etalabs.net/compare_libcs.html">embedded libc comparison</a></li>
<li><a href="https://uclibc.org/">uClibc</a> - strives to be a minimalist C library suitable for embedded computing.</li>
</ul>
<h3>Compilers</h3>
<ul>
<li><a href="https://git.sr.ht/~mcf/cproc">cproc</a> - frontend for C11 for QBE.</li>
<li><a href="https://github.com/larmel/lacc">lacc</a> - A simple, self-hosting C compiler.</li>
<li><a href="https://c9x.me/compile/">qbe</a> - compiler backend targeting AMD64 (Linux and OSX), ARM64, and RISC-V64.</li>
<li><a href="https://www.simple-cc.org/">scc</a> - Simple C Compiler.</li>
<li><a href="https://bellard.org/tcc/">tcc</a> - Tiny C Compiler <a href="https://repo.or.cz/w/tinycc.git">git repo</a></li>
</ul>
<h3>Compression</h3>
<ul>
<li><a href="http://oldhome.schmorp.de/marc/liblzf.html">liblzf</a> - very fast, legally unencumbered compression library.</li>
<li><a href="https://sortix.org/libz/">libz</a> - ABI and API compatible zlib fork with focus on correctness, proper portability, auditability, simplification and opportunities to optimize (<a href="https://zlib.net/zlib_license.html">zlib license</a>).</li>
<li><a href="https://lzip.nongnu.org">Lzip</a> - Properly designed data compressor outperforming gzip and bzip2 (GPLv2+).</li>
<li><a href="https://github.com/richgel999/miniz">miniz</a> - single C-file reimplementation of zlib (MIT License).</li>
<li><a href="https://tukaani.org/xz/embedded.html">xz embedded</a> - lightweight decompressor for the xz LZMA compressor (public domain).</li>
<li><a href="https://zlib.net/">zlib</a> - the &quot;standard&quot; compression/decompression library, quite small, used in many applications (<a href="https://zlib.net/zlib_license.html">zlib license</a>).</li>
</ul>
<h3>GUI</h3>
<ul>
<li><a href="https://github.com/Immediate-Mode-UI/Nuklear">Nuklear</a> - A single-header ANSI C immediate mode cross-platform GUI library.</li>
</ul>
<h3>Cryptography</h3>
<ul>
<li><a href="https://www.libressl.org/">LibreSSL</a> - LibreSSL is a version of the TLS/crypto stack forked from OpenSSL in 2014.</li>
<li><a href="https://monocypher.org/">Monocypher</a> - Small and freestanding cryptographic
library with relatively good performance written in C99.</li>
</ul>
<h3>Miscellaneous</h3>
<ul>
<li><a href="https://github.com/alfredh/baresip">baresip</a> - barebone SIP stack (BSD licence).</li>
<li><a href="http://libccv.org/">ccv</a> - C-based/Cached/Core Computer Vision Library, A Modern Computer Vision Library.</li>
<li><a href="http://software.schmorp.de/pkg/libev.html">libev</a> - high performance event-loop modelled after libevent but much smaller (dual licensed under 2-clause BSD and GPL).</li>
<li><a href="https://mandoc.bsd.lv/">mandoc</a> - The mandoc UNIX manpage compiler toolset.</li>
<li><a href="https://www.uninformativ.de/git/nhkd/file/README.html">nhkd</a> - nano hot key daemon.</li>
<li><a href="https://www.pjsip.org/">pjsip</a> - open-source SIP stack (GPL).</li>
<li><a href="https://gitlab.com/drummyfish/raycastlib">raycastlib</a> - Single-header public domain C raycasting renderer with no dependencies, using only integer math.</li>
<li><a href="https://git.2f30.org/sdhcp/">sdhcp</a> - IPv4 DHCP client.</li>
<li><a href="https://gitlab.com/drummyfish/small3dlib">small3dlib</a> - Single-header public domain C software 3D renderer with no dependencies, using only integer math.</li>
<li><a href="https://github.com/nothings/stb">stb</a> - single-file public domain (or MIT licensed) libraries for C/C++: various image, font, audio loaders etc.</li>
<li><a href="https://github.com/baskerville/sxhkd">sxhkd</a> - Simple X hotkey daemon</li>
<li><a href="https://github.com/termbox/termbox">termbox</a> - simple ncurses-like library for creating TUIs.</li>
<li><a href="https://github.com/C-Chads/tinygl">TinyGL</a> - free software-only subset of OpenGL 1.1 implemented in C</li>
</ul>
<h3>Distros</h3>
<ul>
<li><a href="https://morpheus.2f30.org/">morpheus</a> - A statically linked musl-based Linux distro. Unmaintained currently.</li>
<li><a href="https://github.com/oasislinux/oasis">oasis</a> - A statically linked small Linux system.</li>
<li><a href="https://sxmo.org/">sxmo</a>: simple X Mobile, Sxmo is a minimalist environment
for Linux mobile devices, such as the pinephone.</li>
</ul>
<h2>Programs</h2>
<p>Some programs work well with dwm. Some general rule of thumb for judging if a
program is usable are:</p>
<ul>
<li>The program will accept any window size.</li>
<li>The program consists of a single window (there are no nested windows, such as in Xpdf).</li>
</ul>
<p>This covers most console-based programs and programs from
<a href="https://9fans.github.io/plan9port/">plan9port</a>.</p>
<h3>Audio Players</h3>
<ul>
<li><a href="https://cmus.github.io/">C* Music Player</a> - aka cmus, an ncurses music player that can be controlled with a UNIX socket.</li>
<li><a href="https://moc.daper.net/">moc</a> - console audio player for Linux/UNIX.</li>
<li><a href="https://www.musicpd.org/">mpd</a> - A client/server based music player with console and graphical front-ends.</li>
<li><a href="https://www.mpg123.de/">mpg123</a> - A console MPG player which doesn't use auto*hell or extra libraries.</li>
<li><a href="http://mpg321.sourceforge.net">mpg321</a> - A command-line MP3 player.</li>
<li><a href="https://sr.ht/~q3cpma/mus/">mus</a> - Modular daemon/client CLI player consuming a plaintext album directory queue</li>
<li><a href="https://www.xiph.org/">vorbis-tools</a> (Ogg/FLAC) - Command-line tools to play Ogg and FLAC files.</li>
</ul>
<h3>BitTorrent clients</h3>
<ul>
<li><a href="https://github.com/btpd/btpd">btpd</a> - The BitTorrent Protocol Daemon.</li>
</ul>
<h3>Feed aggregators / newsreaders</h3>
<ul>
<li><a href="https://codeberg.org/grisha/newsraft">newsraft</a> - feed reader with ncurses user interface.</li>
<li><a href="https://codemadness.org/sfeed.html">sfeed</a> - RSS and Atom parser, CLI programs and <a href="https://codemadness.org/sfeed_curses.html">sfeed_curses UI</a>.</li>
<li><a href="https://github.com/kouya/snownews/">snownews</a> - Text mode RSS newsreader for Linux and Unix.</li>
<li><a href="http://r-36.net/scm/zs/">zs</a> - Zeitungsschau RSS to email converter.</li>
</ul>
<h3>File managers</h3>
<ul>
<li><a href="https://github.com/gokcehan/lf">lf</a> - ranger-like file manager written in Go (WIP but usable).</li>
<li><a href="https://midnight-commander.org/">mc</a> - Midnight Commander is a free cross-platform orthodox file manager.</li>
<li><a href="https://github.com/jarun/nnn">nnn</a> - Nnn's Not Noice, a fork of noice with more features.</li>
<li><a href="https://git.2f30.org/noice/">noice</a> - Small and portable file browser.</li>
<li><a href="https://ranger.github.io/">ranger</a> - ranger is a file manager with vi key bindings written in python but with an interface that rocks.</li>
<li><a href="https://lecram.github.io/p/rover/">rover</a> - Simple file browser for the terminal.</li>
<li><a href="https://github.com/afify/sfm/">sfm</a> - sfm is a simple file manager for unix-like systems.</li>
</ul>
<h3>Git</h3>
<ul>
<li><a href="https://git.codemadness.org/stagit/">stagit</a> - static git page generator (HTML).</li>
<li><a href="https://git.codemadness.org/stagit-gopher/">stagit-gopher</a> - static git page generator for gopher (geomyidae .gph pages).</li>
<li><a href="https://git.milotier.net/stagit-gemini/">stagit-gemini</a> - static git page generator for gemini (gemtext .gmi pages).</li>
</ul>
<h3>Gopher clients</h3>
<ul>
<li><a href="https://lynx.invisible-island.net/">Lynx</a> - text-mode web browser, supports the Gopher protocol.</li>
<li><a href="gopher://bitreich.org/1/releases/sacc">sacc</a> - terminal gopher client by <a href="//suckless.org/people/quinq/">quinq</a>.</li>
</ul>
<h3>IRC clients</h3>
<ul>
<li><a href="https://git.causal.agency/catgirl">catgirl</a> - is a TLS-only terminal IRC client.</li>
<li><a href="https://github.com/nmeum/hii">hii</a> - A file-based IRC client, inspired by ii, with a focus on easing <a href="https://github.com/nmeum/insomnia">frontend development</a></li>
<li><a href="//tools.suckless.org/ii/">ii</a> - A FIFO based IRC client which is part of the suckless.org project.</li>
<li><a href="https://c9x.me/irc/">irc.c from mpu</a> - A minimal curses IRC client.</li>
<li><a href="https://swtch.com/irc/">Irc</a> - IRC client for ACME.</li>
<li><a href="http://plan9.bell-labs.com/sources/contrib/fgb/rc/ircrc">ircrc</a> - An rc-based IRC client similar to ircc. Needs minor modification to run on UNIX.</li>
<li><a href="https://irssi.org/">irssi</a> - Text-mode IRC client.</li>
<li><a href="https://github.com/aaronNGi/jj">jj</a> - A rework of ii that offloads all the IRC protocol parsing to an awk script, permitting triggers.</li>
<li><a href="https://github.com/mcpcpc/kirc">kirc</a> - A tiny IRC client written in POSIX C99.</li>
<li><a href="https://github.com/ec429/quIRC">quIRC</a> - a lightweight console IRC client.</li>
<li><a href="//tools.suckless.org/sic/">sic</a> - Another suckless.org IRC client. Similar to ircc.</li>
</ul>
<h3>Image viewers</h3>
<ul>
<li><a href="https://feh.finalrewind.org/">feh</a> - A highly customizable image viewer that can also set a desktop background image.</li>
<li><a href="https://sr.ht/~exec64/imv/">imv</a> - Simple X11/Wayland Image Viewer. Depends on SDL2 and FreeImage.</li>
<li><a href="https://9fans.github.io/plan9port/">jpg/gif/bmp/png</a> - Simple programs from Plan 9 to display images in no-frills windows. Included with <a href="https://9fans.github.io/plan9port/">plan9port</a>.</li>
<li><a href="https://git.codemadness.org/lel/file/README.html">lel</a> - Farbfeld image viewer.</li>
<li><a href="https://www.johnhawthorn.com/meh/">meh</a> - image viewer using raw XLib, libjpeg, libpng and libgif.</li>
<li><a href="https://9fans.github.io/plan9port/man/man1/page.html">page</a> - Plan 9's image/document viewer program. Included with <a href="https://9fans.github.io/plan9port/">plan9port</a>.</li>
<li><a href="https://spiegl.de/qiv/">qiv</a> - Quick Image Viewer.</li>
<li><a href="https://github.com/muennich/sxiv">sxiv</a> - Simple (or small or suckless) X Image Viewer. Depends on xlib and imlib2. Unmaintained currently.</li>
<li><a href="https://codeberg.org/nsxiv/nsxiv/">nsxiv</a> - Neo Simple X Image Viewer, a fork of now unmaintained sxiv. Depends on xlib and imlib2.</li>
<li><a href="http://web.aanet.com.au/gwg/xli.html">xli</a></li>
<li><a href="https://github.com/stoeckmann/xwallpaper">xwallpaper</a> - Minimalistic wallpaper utility</li>
<li><a href="http://sourceforge.net/projects/xzgv/">xzgv</a></li>
</ul>
<h3>Instant-Messaging clients</h3>
<ul>
<li><a href="https://www.bitlbee.org/main.php/news.r.html">bitlbee</a> - A program to translate IM protocols to IRC.</li>
<li><a href="https://www.centerim.org/index.php/Main_Page">CenterIM</a> - A centericq fork.</li>
<li><a href="http://cybione.org/~irssi-xmpp/">irssi-xmpp</a> - An irssi plugin to connect to the Jabber network.</li>
<li><a href="http://mcabber.com/">mcabber</a> - A console jabber client.</li>
<li><a href="https://profanity-im.github.io/">profanity</a> - A console XMPP client written in C. OMEMO/GPG/OTR supported.</li>
<li><a href="http://ysmv7.sourceforge.net/">ysm</a> - A console-based ICQ client.</li>
</ul>
<h3>Mail clients</h3>
<ul>
<li><a href="https://sr.ht/~rjarry/aerc">aerc</a> - email client for your terminal.</li>
<li><a href="https://github.com/nicm/fdm">fdm</a> - fetch and deliver mail.</li>
<li><a href="http://heirloom.sourceforge.net/mailx.html">heirloom-mailx</a> - A mail client based on the original Berkeley Mail 8.1 with many new features.</li>
<li><a href="https://isync.sourceforge.net/">isync</a> - MDA that can work better for IMAP, also known as mbsync.</li>
<li><a href="https://github.com/chneukirchen/mblaze">mblaze</a> - UNIX utilities to deal with Maildir.</li>
<li><a href="http://marmaro.de/prog/mmh/">mmh</a> - Meillo's mail handler (mmh), a fork of nmh.</li>
<li><a href="https://marlam.de/mpop/">mpop</a> - a POP3 mail client. It is much <a href="https://marlam.de/mpop/documentation">faster and simpler than fetchmail</a>.</li>
<li><a href="http://www.mutt.org/">mutt</a> - A ncurses mail client for fetching, reading and sending email.</li>
<li><a href="https://www.nongnu.org/nmh/">nmh</a> - Message Handling System.</li>
<li><a href="https://9fans.github.io/usr/local/plan9/mail/">plan9port/Mail</a> - A mail client for acme. Included with <a href="https://9fans.github.io/plan9port/">plan9port</a>.</li>
<li><a href="https://www.sdaoden.eu/code.html">s-nail</a> - An improved heirloom-mailx.</li>
</ul>
<h3>Media players</h3>
<ul>
<li><a href="https://www.ffmpeg.org/ffplay.html">ffplay</a> - FFplay is a very simple and portable media player. It came with ffmpeg which depends on mplayer.</li>
<li><a href="http://www.mplayerhq.hu/design7/news.html">mplayer</a></li>
<li><a href="https://mpv.io/">mpv</a> - a free, open-source and cross-platform media player.</li>
</ul>
<h3>Notifications</h3>
<ul>
<li><a href="https://github.com/dudik/herbe">herbe</a> - Daemon-less notifications without D-Bus. Minimal, lightweight &amp; written in C. You can use <a href="https://github.com/Sweets/tiramisu">tiramisu</a> to call herbe.</li>
<li><a href="https://github.com/Sweets/tiramisu">tiramisu</a> - Tiramisu is a notification daemon based on dunst that outputs notifications to STDOUT in order to allow the user to process notifications any way they prefer, like in your dwm statusbar for example.</li>
</ul>
<h3>Password managers</h3>
<ul>
<li><a href="https://www.nongnu.org/oath-toolkit/oathtool.1.html">oathtool</a> - Open AuTHentication (OATH) one-time password tool.</li>
<li><a href="https://github.com/ritze/pinentry-dmenu">pinentry-dmenu</a> - pinentry-dmenu is a pinentry program with the charm of dmenu. It is a nice interface to <a href="https://www.passwordstore.org/">pass</a></li>
<li><a href="https://www.passwordstore.org/">pass</a> - &quot;The standard UNIX password manager&quot;.</li>
<li><a href="https://github.com/BioBox/spm">spm</a> - Simple password manager (actively maintained fork of nmeum's tpm).</li>
<li><a href="https://github.com/nmeum/tpm/">tpm</a> - Tiny password manager.</li>
</ul>
<h3>PDF viewers</h3>
<ul>
<li><a href="https://mupdf.com/">mupdf</a> - a lightweight PDF viewer written in C. It supports PDF, XPS, EPUB, XHTML, CBZ, PNG, JPEG, GIF and TIFF.</li>
<li><a href="https://pwmt.org/projects/zathura/">zathura</a> - an extensible document viewer/shell which can support CBZ, DJVU, PS, EPUB (with mupdf) and PDF (with mupdf or poppler).</li>
</ul>
<h3>Shells</h3>
<ul>
<li><a href="http://gondor.apana.org.au/~herbert/dash/">dash</a> - DASH is a POSIX-compliant implementation of /bin/sh that aims to be as small as possible.</li>
<li><a href="https://www.mirbsd.org/mksh.htm">mksh</a> - MirBSD Korn Shell, an actively developed free implementation of the Korn Shell programming language
and a successor to the Public Domain Korn Shell (pdksh).</li>
<li><a href="https://github.com/ibara/oksh">oksh</a> - Portable OpenBSD ksh.</li>
<li><a href="https://yash.osdn.jp">yash</a> - Yet another shell, is intended to be the most POSIX-compliant shell in the world while supporting features for daily interactive and scripting use.</li>
</ul>
<h3>Text editors</h3>
<ul>
<li><a href="http://acme.cat-v.org">acme</a> - Rob Pike's framing text editor for Plan 9. Included in <a href="https://9fans.github.io/plan9port/">plan9port</a>.</li>
<li><a href="https://9fans.github.io/plan9port/man/man1/ed.html">ed</a> - <a href="https://www.gnu.org/fun/jokes/ed-msg.html">ED IS THE STANDARD TEXT EDITOR!</a></li>
<li><a href="http://www.radare.org/r/">ired</a> - A minimalist hexadecimal editor and bindiffer for p9, w32 and *nix.</li>
<li><a href="https://homepage.boetes.org/software/mg/">mg</a> - A portable version of mg maintained by the OpenBSD team.</li>
<li><a href="https://github.com/adsr/mle">mle</a> - A small, flexible console text editor.</li>
<li><a href="https://www.nano-editor.org/">nano</a> - A pico clone - this is small simple code and easy to use.</li>
<li><a href="https://github.com/aligrudi/neatvi">neatvi</a> - A minimal vi implementation supporting bidirectional UTF-8</li>
<li><a href="https://github.com/kyx0r/nextvi">nextvi</a> - A continuation of neatvi development with more features.</li>
<li><a href="https://sites.google.com/a/bostic.com/keithbostic/vi">nvi</a> - A small, multiple file vi-alike.</li>
<li><a href="https://micro-editor.github.io/">micro</a> - A terminal text editor, written in go with common key bindings like ctrl-c to copy and ctrl-v to paste.</li>
<li><a href="http://sam.cat-v.org/">sam</a> - An editor by Rob Pike with inspiration from ed.</li>
<li><a href="https://sim.ssnf.xyz/">sim</a> - The sim text editor. Based on vim and sam.</li>
<li><a href="http://ex-vi.sourceforge.net/">traditional vi</a> - A fixed version of the original vi.</li>
<li><a href="https://www.vim.org">vim</a> (With the GUI, use <code>:set go+=c</code> to kill popup dialogs). It can be compiled to be as minimal as possible (see vim-tiny in Debian repos).</li>
<li><a href="https://github.com/martanne/vis">vis</a> - A modern, legacy free, simple yet efficient vim-like editor.</li>
<li><a href="http://www.cse.yorku.ca/~oz/wily/">wily</a> - An acme clone for POSIX.</li>
</ul>
<h3>Text processing</h3>
<ul>
<li><a href="https://github.com/skeeto/scratch/tree/master/csvquote">csvquote</a> - tool to
encode troublesome CSV characters such that unix tools can correctly process
them. SIMD optimized by default, portable C fallback can be enabled at build
time.</li>
<li><a href="https://codemadness.org/json2tsv.html">json2tsv</a> - a JSON to TAB-Separated Value (TSV) converter and stand-alone JSON parser code.</li>
<li><a href="https://github.com/mity/md4c">md4c</a> - markdown to html converter.  Fast,
<a href="https://commonmark.org/">CommonMark</a> compliant, offers extensions which can
be switched on/off via cli args.</li>
</ul>
<h3>Utilities / miscellaneous</h3>
<ul>
<li><a href="https://www.brain-dump.org/projects/abduco/">abduco</a> - session {at,de}tach support.</li>
<li><a href="https://www.brain-dump.org/projects/dvtm/">dvtm</a> - dynamic virtual terminal manager.</li>
<li><a href="https://eradman.com/entrproject/">entr</a> - Run arbitrary commands when files change.</li>
<li><a href="https://gitlab.com/mhdy/mrandr">mrandr</a> - Simple and lightweight monitor (display) profile manager written in POSIX shell.</li>
<li><a href="https://github.com/deadpixi/mtm">mtm</a> - Micro Terminal Multiplexer.</li>
<li><a href="https://github.com/chneukirchen/nq">nq</a> - UNIX command line queue utility.</li>
<li><a href="https://www.ivarch.com/programs/pv.shtml">pv</a> - tool for monitoring the progress of data through a pipeline.</li>
<li><a href="https://github.com/p-gen/smenu">smenu</a> - A powerful and versatile CLI selection tool for interactive or scripting use.</li>
<li><a href="https://github.com/bitsmanent/snore">snore</a> - Sleep with feedback.</li>
<li><a href="https://github.com/yt-dlp/yt-dlp">yt-dlp</a> - youtube-dl fork, download videos from youtube.com or other video platforms.</li>
<li><a href="http://zbar.sourceforge.net/">zbar</a> - software suite for reading bar codes from various sources, such as video streams, image files.</li>
</ul>
<h3>Web browsers</h3>
<ul>
<li><a href="http://lab-fgb.com/abaco/">Abaco</a> - A graphical web browser for Plan 9 that supports most of HTML 4.01.</li>
<li><a href="https://www.dillo.org/">Dillo</a> - Blazingly fast graphical browser with basic CSS support.</li>
<li><a href="http://edbrowse.org/">edbrowse</a> - ed-alike webbrowser. Unfortunately it is not the standard web browser.</li>
<li><a href="http://elinks.or.cz/">ELinks</a> -  Feature-rich fork of Links, but without the graphical mode.</li>
<li><a href="http://links.twibright.com/">Links</a> - A text-based browser that also has a simple graphical mode.</li>
<li><a href="http://lynx.invisible-island.net/">Lynx</a> - text-mode web browser, supports the Gopher protocol.</li>
<li><a href="http://www.netsurf-browser.org/">netsurf</a> - Very fast graphical browser with custom engine that has near-perfect CSS support.</li>
<li><a href="//surf.suckless.org/">surf</a> - WebKit2 browser that can be embedded into other applications and uses dmenu as the address bar.</li>
<li><a href="https://github.com/cmcaine/tridactyl">Tridactyl</a> - A WebExtension that tries to make Firefox Vim like. Unlike many similar extensions it works with firefox 57+.</li>
<li><a href="https://www.uzbl.org/">uzbl</a> - Web interface tools which adhere to the UNIX philosophy.</li>
<li><a href="http://vimperator.org/">Vimperator</a> - An extension to make Firefox keyboard friendly and Vim-like. (Doesn't work with Firefox 57+).</li>
<li><a href="http://w3m.sourceforge.net/">w3m</a> - Text-based browser which can view images in the terminal (X11 hack).</li>
</ul>
<h3>X11</h3>
<ul>
<li><a href="https://github.com/arnoldrobbins/9menu/">9menu</a> - A menu program based on the Blit-style menus so prevalent in Plan 9.</li>
<li><a href="//tools.suckless.org/dmenu/">dmenu</a> - dmenu is a dynamic menu for X.</li>
<li><a href="https://github.com/lolilolicon/FFcast">ffcast</a> - Simple screencasting.</li>
<li><a href="http://www.muquit.com/muquit/software/grabc/grabc.html">grabc</a> - Grab the color of some pixel in X11.</li>
<li><a href="http://www.semicomplete.com/projects/keynav/">keynav</a> - A new way for keyboard selection.</li>
<li><a href="https://github.com/wmutils/core">wmutils</a> - A set of utilities for managing windows. Can be used on their own or to augment a WM itself. Support for tiling present.</li>
<li><a href="https://github.com/jcs/xbanish">xbanish</a> - Hide the cursor while typing.</li>
<li><a href="https://github.com/astrand/xclip">xclip</a> - A tool for controlling the X11 clipboard.</li>
<li><a href="http://www.semicomplete.com/projects/xdotool/">xdotool</a> - A tool for scripting X11 actions.</li>
<li><a href="https://github.com/baskerville/xdo">xdo</a> - Similar to xdotool but it's smaller and has a much cleaner codebase.</li>
<li><a href="http://r-36.net/scm/xzoom/">xzoom</a> - A simple zoom application.</li>
</ul>
<h2>Daemons</h2>
<p>This section is dedicated to daemons that are related to the UNIX philosophy.</p>
<h3>Web Servers</h3>
<ul>
<li><a href="https://unix4lyfe.org/darkhttpd/">darkhttpd</a> - Darkhttpd is a simple, fast HTTP 1.1 web server for static content.</li>
<li><a href="https://github.com/troglobit/merecat/">merecat</a> - thttpd with vhost support.</li>
<li><a href="http://acme.com/software/mini_httpd/">mini_httpd</a> - By the thttpd developers.  Even smaller subset of thttpd with support for .htpasswd, CGI, dirlisting, HTTP errors and SSL, nothing more or less than that.</li>
<li><a href="https://bsd.plumbing/">OpenBSD httpd/relayd</a> - Designed to combat the feature creep of modern webservers. Boasts concise configuration, with sane defaults.</li>
<li><a href="//tools.suckless.org/quark/">quark</a> - suckless simple static web server.</li>
<li><a href="http://acme.com/software/thttpd/">thttpd</a> - Outperforms many bigger httpds. Old, still works very well!</li>
</ul>
<h3>Gopher servers</h3>
<ul>
<li><a href="http://r-36.net/scm/geomyidae/">geomyidae</a> - small gopher-daemon written by <a href="//suckless.org/people/20h/">20h</a>.</li>
</ul>
<h3>Misc daemons</h3>
<ul>
<li><a href="https://git.2f30.org/scron/">scron</a> - simple crond.</li>
</ul>
