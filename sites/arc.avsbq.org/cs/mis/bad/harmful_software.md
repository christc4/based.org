<small>Copied from http://harmful.cat-v.org/software/</small>
<h1>All software sucks</h1>

<p>&ldquo;<em>And folks, let&rsquo;s be honest. Sturgeon was an optimist. Way more than 90% of code is crap.</em>&rdquo; &ndash; Al viro</p>

<p>&ldquo;<em>There are two ways of constructing a software design: One way is to make it so simple that there are obviously no deficiencies and the other way is to make it so complicated that there are no obvious deficiencies.</em>&rdquo; &ndash; C.A.R. Hoare, The 1980 ACM Turing Award Lecture</p>

<p>&ldquo;<em>One of my most productive days was throwing away 1000 lines of code.</em>&rdquo; &ndash; <a href="http://genius.cat-v.org/ken-thompson/">Ken Thompson</a></p>

<p>&ldquo;<em>..At first I hoped that such a technically unsound project would collapse but I
soon realized it was doomed to success. Almost anything in software can be
implemented, sold, and even used given enough determination. There is nothing a
mere scientist can say that will stand against the flood of a hundred million
dollars. But there is one quality that cannot be purchased in this way -and
that is reliability. The price of reliability is the pursuit of the utmost
simplicity. It is a price which the very rich find most hard to pay.</em>&rdquo; &ndash; C.A.R. Hoare</p>

<p>&ldquo;<em>It&rsquo;s a curious thing about our industry: not only do we not learn from our mistakes, we also don&rsquo;t learn from our successes.</em>&rdquo;  &ndash; Keith Braithwaite</p>

<ul>
<li><a href="http://quotes.cat-v.org/programming/">More software and programming quotes</a>.</li>
</ul>


<h2>Harmful stuff</h2>

<p>This is a work in progress, much is missing.</p>

<table>
<tr><th>Harmful things</th><th>Less harmful alternatives</th></tr>
<tr><td>SGML, <a href="xml/">XML</a>, YAML.</td><td>JSON, CSV, <a href="http://man.cat-v.org/plan_9/6/ndb">ndb(6)</a>, plain(UTF-8) text.</td></tr>
<tr><td>NFS, SMB, AFS, WebDAV.</td><td><a href="http://9p.cat-v.org">9p</a>.</td></tr>
<tr><td><a href="c++/">C++</a>, <a href="java/">Java</a>, Vala, D, Python, <a href="ruby/">Ruby</a>.</td><td>C, <a href="http://go-lang.cat-v.org">Go</a>, <a href="http://limbo.cat-v.org">Limbo</a>.</td></tr>
<tr><td>pthreads (PoSix threads).</td><td>CSP-style concurrency: <a href="http://go-lang.cat-v.org">Go</a>, <a href="http://man.cat-v.org/plan_9/2/thread">libthread</a>, libtask, <a href="http://limbo.cat-v.org">Limbo</a>, Erlang, ...</td></tr>
<tr><td>Perl, Ruby.</td><td><a href="http://rc.cat-v.org">rc</a>, <a href="http://man.cat-v.org/plan_9/1/awk">awk</a>.</td></tr>
<tr><td>PCRE</td><td><a href="http://doc.cat-v.org/bell_labs/structural_regexps/">Structural Regular Expressions</a> or plain classic/extended <a href="http://man.cat-v.org/plan_9/6/regexp">regexps</a> (as used in awk, sed, grep, etc.).</td></tr>
<tr><td>Bash, tcsh, zsh.</td><td><a href="http://rc.cat-v.org">rc</a>, OpenBSD's pdksh, ash/dash.</td></tr>
<tr><td>GNU Coreutils.</td><td>Plan 9 from User Space.</td></tr>
<tr><td>GNU Screen.</td><td>tmux.</td></tr>
<tr><td>GNU info.</td><td><a href="http://man.cat-v.org">Man pages</a>.</td></tr>
<tr><td><a href="GCC">GCC</a>.</td><td><a href="http://man.cat-v.org/plan_9/1/2c">8c</a>, tcc.</td></tr>
<tr><td>glibc.</td><td>musl, ucLibc, DietLibc.</td></tr>
<tr><td>GNU autoconf/automake, CMake, imake, scons, waf.</td><td><a href="http://doc.cat-v.org/plan_9/4th_edition/papers/mk">mk</a>, or plain old portable makefiles.</td></tr>
<tr><td>Glib.</td><td>libc (see above), <a href="http://man.cat-v.org/p9p/3/">p9p's C libraries</a>.</td></tr>
<tr><td>GTK, Qt, VxWindows.</td><td>Tk, textual interfaces.</td></tr>
<tr><td>Vim, Emacs, nano, Eclipse, ...</td><td><a href="http://acme.cat-v.org">Acme</a>, <a href="http://sam.cat-v.org">Sam</a>, <a href="http://man.cat-v.org/plan_9/1/ed">ed</a>.</td></tr>
<tr><td><a href="http://benlynn.blogspot.com/2011/02/utf-8-good-utf-16-bad_07.html">UTF-16</a>, UTF-32, Latin-1, other encodings.</td><td><a href="http://doc.cat-v.org/bell_labs/utf-8_history">UTF-8</a>.</td>.</tr>
<tr><td>iSCSI, FCoE.</td><td>AoE (ATA over Ethernet).</td></tr>
<tr><td>PAM (Pluggable Authentication Modules).</td><td><a href="http://doc.cat-v.org/plan_9/4th_edition/papers/auth">Factotum</a>.</td></tr>
<tr><td><a href="xml/xmpp/">Jabber and XMPP</a>.</td><td>IRC (for IM), STOMP (for general distributed messaging).</td></tr>
<tr><td>IMAP.</td><td>SMAP(Simple Mail Access Protocol).</td></tr>
<tr><td>SQL databases.</td><td>Tutorial D, pq, BigTable, plain old hierarchical filesystems.</td></tr>
<tr><td><a href="svn/">Subversion, aka svn</a>.</td><td>Git, Mercurial (aka hg); hell, even CVS or plain old tarballs would be better than svn.</td></tr>
<tr><td>FreeBSD, NetBSD, Solaris.</td><td>OpenBSD.</td></tr>
<tr><td>Apache, lighttpd.</td><td>thttpd, OpenBSD's fork of apache 1.3, nginx, or best of all: don't use HTTP.</td></tr>
<tr><td>SVG</td><td>PS(PostScript)</td></tr>
<tr><td>PDF</td><td>PS(PostScript), DjVu.</td></tr>
<tr><td>EPUB</td><td>DjVu.</td></tr>
<tr><td>ALSA</td><td>OSS4</td></tr>
<tr><td>GPL, LGPL, Apache Software License, MPL, CC.</td><td>ISC, MIT/X, BSD, CC0, public domain.</td></tr>
<tr><td>head</td><td><a href="http://man.cat-v.org/unix_8th/1/sed">sed 11q</a></td></tr>
</table>


<p>Harmful things that are so superfluous and useless that require no alternative:</p>

<ul>
<li>Unicode BoM in UTF-8.</li>
<li>IDN.</li>
<li>Dynamic linking, in particular in Unix systems.</li>
<li>PoSix locales.</li>
<li>Adobe Flash.</li>
<li>GNOME and KDE.</li>
<li>Boost.</li>
<li><a href="http://archive.eiffel.com/doc/manuals/technology/bmarticles/uml/page.html">UML</a>.</li>
</ul>


<p><strong>Note</strong>: At the moment a detailed rationale is not provided for most of this, so figuring out why some things are considered more or less harmful than others is left as an exercise for the reader. Here is a hint: <em>complexity</em> is the bane of all software, <em>simplicity</em> is the most important quality. See also: <a href="http://doc.cat-v.org/programming/worse_is_better">Worse is Better by Richard Gabriel</a>.</p>

<h2>Features</h2>

<p>JWZ&rsquo;s Law of Software Envelopment: &ldquo;<em>Every program attempts to expand until it can read mail. Those programs which cannot so expand are replaced by ones which can.</em>&rdquo;</p>

<p>Most so called &lsquo;features&rsquo; are either useless or worse than useless.</p>

<ul>
<li><a href="http://weblogs.mozillazine.org/roc/archives/2010/06/not_implementin.html">Not Implementing Features Is Hard</a> - By Robert O'Callahan.</li>
</ul>


<h2>Frameworks</h2>

<ul>
<li><a href="http://discuss.joelonsoftware.com/default.asp?joel.3.219431">Why I Hate Frameworks</a> - By Benji Smith.</li>
</ul>
