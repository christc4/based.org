<h1>Stuff that sucks</h1>
<small>Copied from <a href="https://suckless.org/sucks/">https://suckless.org/sucks</a></small>
<p>See the <a href="//suckless.org/philosophy">philosophy</a> page about what applies
to this page.</p>
<p>Bigger topics that suck: <a href="//suckless.org/sucks/systemd">systemd</a>,
<a href="//suckless.org/sucks/web">the web</a></p>
<h2>Libraries</h2>
<p>These libraries are broken/considered harmful and should not be used
if it's possible to avoid them. If you use them, consider looking for
alternatives.</p>
<ul>
<li><p><a href="http://library.gnome.org/devel/glib/">glib</a> - implements C++ STL on top of C
(because C++ sucks so much, let's reinvent it!), adding lots of useless data
types for <a href="http://library.gnome.org/devel/glib/unstable/glib-Basic-Types.html">&quot;portability&quot; and &quot;readability&quot;
reasons</a>.
even worse, it is not possible to write robust applications using glib, since
it <a href="https://bugzilla.gnome.org/show_bug.cgi?id=674446">aborts in out-of-memory situations</a>.
glib usage is required to write gtk+ and gnome applications, but is also used when common
functionality is needed (e.g. hashlists, base64 decoder, etc). it is not suited
at all for static linking due to its huge size and the authors explicitly state
that <a href="https://bugzilla.gnome.org/show_bug.cgi?id=768215#c16">&quot;static linking is not supported&quot;</a>.</p>
<p>Alternatives: <a href="https://github.com/atheme/libmowgli-2">libmowgli</a>,
<a href="https://github.com/rofl0r/libulz">libulz</a>,
BSD <a href="https://man.openbsd.org/queue">queue.h</a>/<a href="https://man.openbsd.org/tree">tree.h</a> macros.</p>
</li>
<li><p><a href="http://gmplib.org/">GMP</a> - GNU's bignum/arbitrary precision
library. Quite bloated, slow and <a href="https://gmplib.org/repo/gmp/file/tip/memory.c#l105">calls abort() on failed
malloc</a></p>
<p>Alternatives: <a href="http://www.libtom.net/LibTomMath/">libtommath</a>,
<a href="http://www.libtom.net/TomsFastMath/">TomsFastMath</a>,
<a href="https://github.com/creachadair/imath">imath</a>,
<a href="//libs.suckless.org/libzahl">libzahl</a> (WIP),
<a href="https://github.com/suiginsoft/hebimath">hebimath</a> (WIP)</p>
</li>
</ul>
<h2>Build Systems</h2>
<ul>
<li><a href="http://www.cmake.org/">cmake</a> (written in C++) - so huge and bloated,
compilation takes longer than compiling GCC (!). It's not even possible
to create freestanding Makefiles, since the generated Makefiles call
back into the cmake binary itself. Usage of cmake requires learning a
new custom scripting language with very limited expressiveness. Its
major selling point is the existence of a clicky-click GUI for windows
users.</li>
<li><a href="https://code.google.com/p/waf/">waf</a> and
<a href="http://www.scons.org/">scons</a> (both written in Python) - waf code is
dropped into the compilee's build tree, so it does not benefit from
updated versions and bugfixes.</li>
</ul>
<p>As these build systems are often used to compile C programs, one has to
set up a C++ compiler or Python interpreter respectively just in order
to be able to build some C code.</p>
<p>Alternatives:
<a href="http://pubs.opengroup.org/onlinepubs/9699919799/utilities/make.html">make</a>,
<a href="http://doc.cat-v.org/plan_9/4th_edition/papers/mk">mk</a></p>
<h2>Version Control Systems</h2>
<ul>
<li><a href="https://subversion.apache.org/">subversion</a> - Teaches developers to
think of version control in a harmful and terrible way, centralized,
ugly code, conceptionally broken in a lot of terms. &quot;Centralized&quot; is
said to be one of the main benefits for &quot;enterprise&quot; applications,
however, there is no benefit at all compared to decentralized version
control systems like git. There is no copy-on-write, branching
essentially will create a 1:1 copy of the full tree you have under
version control, making feature-branches and temporary changes to your
code a painful mess. It is slow, encourages people to come up with weird
workarounds just to get their work done, and the only thing enterprisey
about it is that it just sucks.</li>
</ul>
<h2>Programs</h2>
<p>There are many broken X programs. Go bug the developers of these
broken programs to fix them. Here are some of the main causes of this
brokenness:</p>
<ul>
<li>The program <strong>assumes a specific window management model</strong>,
e.g. assumes you are using a WIMP-window manager like those
found in KDE or Gnome. This assumption breaks the <a href="http://tronche.com/gui/x/icccm/">ICCCM
conventions</a>.</li>
<li>The application uses a <strong>fixed size</strong> - this limitation does not fit
into the world of tiling window managers very well, and can also be seen
as breaking the ICCCM conventions, because a fixed sized window assumes
a specific window management model as well (though the ICCCM does not
forbid fixed-size windows). In any case, the ICCCM requests that clients
accept any size the window manager proposes to them.</li>
<li>The program is based on strange <strong>non-standard window manager
hints</strong> that only work properly with a window manager supporting these
extensions - this simply breaks the ICCCM as well. E.g. trash icon
programs.</li>
<li>The program does not conform to ICCCM due to some <strong>missing or
improperly set hints</strong>.</li>
</ul>
<p>If you still need some program which expects a floating WM, use it in
floating mode.</p>
<h2>Documentation</h2>
<p>Somewhen GNU tried to make the world a bit more miserable by inventing
<a href="https://www.gnu.org/software/texinfo/">texinfo</a>. The result is that
in 2019 man pages are still used and the documentation of GNU tools
requires you to run <code>info $application</code>. The info browser is awkward and
unintuitive and the reason why no one gets further than finding 'q' to
quit it.</p>
<p>Look at GNU tools how to not handle documentation.</p>
<p>Talking about the suck in enforced HTML documentation, which forces
you to open up a 1 Gb of RAM wasting web browser, just to see some
eye-candy, which could have been described in the source with some easy
way to jump to that line in the source code, is not worth the time.</p>
<p>The suckless way is to have a short usage and a descriptive manpage. The
complete details are in the source.</p>
<p>Alternatives: roff, <a href="https://mandoc.bsd.lv/">mdoc</a>.</p>
<h2>C Compilers</h2>
<ul>
<li><a href="http://gcc.gnu.org/">GCC</a>: as of 2016 it is now written in C++ and so
complete suck. Why can't a compiler just be a simple binary doing its work
instead of adding path dependencies deep into the system?</li>
<li><a href="http://clang.llvm.org/">Clang</a> is written in C++. If you don't
believe that it sucks, try to build clang by hand.</li>
</ul>
<p>Alternatives: see the Compilers section of the <a href="../rocks">/rocks/</a> page.</p>
<h2>See also</h2>
<p>The <a href="http://harmful.cat-v.org/software/">list of harmful software</a> at
<a href="http://cat-v.org">cat-v.org</a>.<br>
<a href="systemd_sucks">SystemD</a> and the <a href="web_sucks">web</a> sucks.</p>
