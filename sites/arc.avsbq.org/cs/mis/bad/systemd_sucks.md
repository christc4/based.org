<h1>systemd is the best example of Suck.</h1>
<p>There is a menace which is spreading like a disease throughout the Linux
world, it is called <a href="https://github.com/systemd/systemd">systemd</a>.</p>
<h2>Short Historical Summary</h2>
<p>systemd is a replacement for the standard init command, which normally runs
as process id 1 on initialisation of a UNIX bootup. There has been a movement,
especially around the <a href="http://www.redhat.com">Red Hat</a>-related developers to
copy <a href="http://www.microsoft.com/">Microsoft Windows</a> and all of its features.
Now this interpretation of how a userspace should look like is implemented and
was introduced with big criticism and change in the Open Source world into
many distributions. The <a href="https://wiki.debian.org/Debate/initsystem/systemd">debacle in Debian</a>
is the best example in how to not introduce such a changing technology into a
distribution.</p>
<h2>What PID 1 Should Do</h2>
<p>When your system boots up the kernel is executing a given binary in its known
namespace. To
see what are the only tasks the application running as pid 1 has to do, see
<a href="//git.suckless.org/sinit/">sinit</a>. Just wait for child process to reap
and run some other init scripts.</p>
<h2>The systemd Chronicles</h2>
<p>Now follows a collection of links related to all the features systemd tries to
enforce upon you as a Linux user, because »they know better«.</p>
<p>Please add all the links you can find! Generations after us have to know how
to not do software and system development! See the
<a href="//suckless.org/wiki">wiki instructions</a> on how to contribute to this
page.</p>
<h2>The Never Ending Part</h2>
<ul>
<li><p>Your link here.</p>
</li>
<li><p><a href="http://soggywizards.com/tips/code/os/initialization/systemd/good-bad-ugly/">The SystemD Platform: an Unbiased Critique</a></p>
</li>
</ul>
<h2>Part 4</h2>
<ul>
<li><a href="https://github.com/systemd/systemd/issues/2039">systemd breaks POSIX IPC</a></li>
<li><a href="https://knzl.de/systemd-removeipc">systemd-logind deletes your message queues</a></li>
<li><a href="https://github.com/systemd/systemd/issues/8579">systemd breaks the mouse(again)</a></li>
</ul>
<h2>Part 3</h2>
<p><a href="https://github.com/systemd/systemd/blob/76153ad45f09b6ae45464f2e03d3afefbb4b2afe/NEWS">NEWS revision 76153ad45f09b6ae45464f2e03d3afefbb4b2afe</a></p>
<ul>
<li><a href="https://github.com/systemd/systemd/blob/76153ad45f09b6ae45464f2e03d3afefbb4b2afe/NEWS#L1306">systemd does {,U}EFI bootload</a>
<ul>
<li>Should systemd's PID be changed from 1 to a negative, or imaginary,
number? It now exists before the kernel itself, during a bootup. 
See also <a href="https://wiki.archlinux.org/index.php/systemd-boot">systemd-boot</a>.</li>
</ul>
</li>
<li><a href="https://github.com/systemd/systemd/blob/76153ad45f09b6ae45464f2e03d3afefbb4b2afe/NEWS#L1233">systemd replaces sudo and su</a>
<ul>
<li>Please note the command name, machinectl and 
its <a href="https://www.freedesktop.org/software/systemd/man/machinectl.html">features at the manpage</a>. 
In exchange for a program which contains sudo, su and kill 
(and does some functions which historically ssh/telnet did), bare metal
users have a tons of bloat and a lot of things to disable, if even 
possible, useful only to people which deal with virtual machines.</li>
</ul>
</li>
<li><a href="https://github.com/systemd/systemd/blob/76153ad45f09b6ae45464f2e03d3afefbb4b2afe/NEWS#L1073">systemd-journald can do log-rotate</a>
<ul>
<li>Being journal files binaries written with easily corruptable
transactions, does this feature make the log unreadable at times?</li>
</ul>
</li>
<li><a href="https://github.com/systemd/systemd/blob/76153ad45f09b6ae45464f2e03d3afefbb4b2afe/NEWS#L708">Transient units</a>
<ul>
<li>Temporary services, because we love to reinvent procps, forking, 
nohup and lsof.</li>
</ul>
</li>
<li><a href="https://github.com/systemd/systemd/blob/76153ad45f09b6ae45464f2e03d3afefbb4b2afe/NEWS#L546">systemd does socat/netcat</a></li>
<li><a href="https://github.com/systemd/systemd/blob/76153ad45f09b6ae45464f2e03d3afefbb4b2afe/NEWS#L274">systemd-logind does sighup and nohup</a>
<ul>
<li>Logout is equivalent to shutting off the machine, so you will NOT 
have any running program after logout, unless you inform your init system.</li>
</ul>
</li>
<li><a href="https://github.com/systemd/systemd/blob/76153ad45f09b6ae45464f2e03d3afefbb4b2afe/NEWS#L391">systemd-nspawn can patch at will any kind of file in a container</a>
<ul>
<li>Paired with transient units and user escalation performable remotely,
this can mean that if you house VPS instances somewhere, your 
hosting provider has means and tools to spy, modify, delete any 
kind of content you store there. Encrypt everything, read your TOS.</li>
</ul>
</li>
<li><a href="https://github.com/systemd/systemd/blob/76153ad45f09b6ae45464f2e03d3afefbb4b2afe/NEWS#L425">systemd does UNIX nice</a></li>
</ul>
<p>from 2016-08-06</p>
<h2>Part 2</h2>
<ul>
<li><a href="http://bugzilla.redhat.com/show_bug.cgi?id=1350450">systemd locks down /etc and makes it read-only</a></li>
<li><a href="https://lists.dns-oarc.net/pipermail/dns-operations/2016-June/014964.html">systemd now does your DNS</a></li>
<li><a href="https://bugzilla.redhat.com/show_bug.cgi?id=817186">systemd hates when you adapt your system (graphics on other than vt1)</a></li>
</ul>
<h2>Part 1</h2>
<p>(If you are adding more parts, link to the right revision and not just the
plain file.)</p>
<p><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c">NEWS revision 3f80af378331748e1373d4c189a8e20cddb735c9</a></p>
<ul>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n2089">logind should wait</a>
<ul>
<li>systemd was introduced to decrease the boot up time. Now that they
do not understand all implications and dependencies, let us add some
artifical time we found out might work for the developers
laptops. More on this small world hypothesis of the systemd
developers below.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1852">screen brightness</a>
<ul>
<li>Screen brightness is something that should crash your boot up when
it is not working.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1763">hostnamed</a>
<ul>
<li>There really should be a process running which exposes the content
of a file. Complexity is without cost.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1742">seqnum removed</a>
<ul>
<li>The sequential ordering of requests was one reason why udevd was
introduced. Now remove it, because the developer laptops do not have
a problem anymore.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1703">floppy group removed</a>
<ul>
<li>Because we know what is right to know about groups. This is just one
example of the mass of group name dependencies systemd is adding.
See sinit for how to not need such dependencies.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1651">sysv removed</a>
<ul>
<li>We have won. Now remove all remains of our defeated enemy as
fast as we can. As said in the beginning of the systemd crusade
against the UNIX infidels: »You can patch it out.« It is no more
there.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1627">abnormal processes</a>
<ul>
<li>Now systemd is getting deep into philosophy. What is »abnormal«?
Well, let us just define it. There is no technical merit to accept
this.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1578">systemd-resolved</a>
<ul>
<li>Every configuration file needs its own process and service.</li>
<li>Symlinks are a good way to solve all world problems.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1571">new is better</a>
<ul>
<li>The systemd development process is flawed by always assuming »new is
best«.</li>
<li>Network configuration should be in my init process.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1500">remote pid 1</a>
<ul>
<li>»Everything will end up having a remote API.« I wonder when systemd
will understand MIME and e-mail.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1465">init does man</a>
<ul>
<li>My init process is too big, it needs its own file hierarchy and an
abstraction layer to find paths.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1401">factory reset</a>
<ul>
<li>Welcome to the Windows OEM world: Factory reset for Linux! Of course
it is in your init process.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1397">system runs runs</a>
<ul>
<li>Exactly. The predisposition of being able to call such a complex
command does not imply the running system. Let's check it again.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1306">clean up directories</a>
<ul>
<li>There is another monster in systemd, it handles tmp files.
There are just some cases before it was introduced to have to clean up
a directory in the file tree. Now there are hundreds. And easily
another case can be added! Of course your init process does that.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1259">firstboot</a>
<ul>
<li>»Interactive queries« pulls in many dependencies. Let us have it in
every installation out there on by default. Of course in pid 1.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1234">journald</a>
<ul>
<li>This is a bit longer. Because the systemd developers were not able
to contribute to any syslog daemon they had to write their own with
some binary format so the principle of being able to read your log
files after a critical crash was violated. To be nice invaders the old
systems were allowed to order the system log through the specified
mechanism. Now that they implemented our specifics, turn off the
neutral syslog delivery. You will see this pattern of »now that we
conquered your culture, obey« more often in systemd.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1224">systemd-terminal</a>
<ul>
<li>Why does the kernel have tty handling? So in serious situations you
will be able to debug it over the last standing PIN on your
motherboard. Let us remove this, run it in pid 1.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1216">networkd is your oppressor</a>
<ul>
<li>Premature optimisation of IP configurations always leads to misery.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1211">We do not understand broadcast</a>
<ul>
<li>With the growth of systemd in complexity and the new depending
software the implications of the added hacks are increasing.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1193">Timezone hack</a>
<ul>
<li>systemd is too complex for such a simple transaction with the
kernel. Do not inform the kernel and add another assumption which is
only documented in the changelog.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1168">the web is a session</a> 
<ul>
<li>For the future generations: In 2015 nearly everyone was trying to
connect all you do to the web. In the UNIX model it is just a client
which should be simple and integrate into the system. systemd is
doing it the other way around and does integrate itself into the
web.</li>
<li>The discussion about why my pid 1 is handling sessions is discussed
in other points.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1128">pid 1 does DNS</a>
<ul>
<li>This change adds another open door to your pid 1, adds caches, new
APIs and of course it will never fail and break systemd on your
initial boot.</li>
<li>They are planning (2015-07) to add multicast DNS to your pid 1.
Multicast DNS implies an enforced timeout to know if someone answered.
systemd was about speed, you know. Look at avahi in how to ruin this
concept for Linux. It has the same developer flaws as systemd: Too
many forced dependencies, you need to implement its API and when it is
taking over you are doomed.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1113">policykit</a>
<ul>
<li>There is a bus inside your pid 1 and if it crashes you are flawed.
Exactly, there are a thousand more cases of errors that could occur
and make your system unbootable. Instead of using a separation of
functions, add everything to a big bus.</li>
<li>Of course when you are using a misdesign like dbus you need to add
interactions over the bus to add features you forgot in the initial
design. Now let us have our pid 1 have to query for the permission
to boot.</li>
<li>»Interactive authentication« will produce you a graphical input form
on your serial line.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n1044">Calendar</a>
<ul>
<li>As you see, your pid 1 should handle your calendars and cron jobs
too.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n994">utmp should go</a>
<ul>
<li>We have taken over your culture, now die! Another flaw in the
systemd small world theory: When something is getting optional it
will be removed.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n964">password agent</a>
<ul>
<li>»Interactive authentication«</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n944">udev timeout</a>
<ul>
<li>Instead of patching the kernel to add a simple solution, add a hack.
Only the systemd developers tell you when it is allowed to wait or
sleep in userspace. The rest obey our orders!</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n924">systemd-pm</a>
<ul>
<li>Power management is required on boot up.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n898">user systemd units</a>
<ul>
<li>What can go wrong when you are adding more paths that are read,
parsed and executed?</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n888">hack the reload</a>
<ul>
<li>First systemd was adding »better features« like socket activation to
make developers use their mechanism for daemons. They hit the
proprietary wall of disgust with this changelog entry. systemd is
too big and you will lose your face if you change the misdesign. Now
add another hack because we can do it. Big empires fall too and
sadly have too many casualties when they are falling. :(</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n783">X11 in systemd</a>
<ul>
<li>Of course graphics were missing in pid 1.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n738">complexity is purity</a>
<ul>
<li>You will of course need PPPoE when you do parallel bootup. Every
1000 lines of code add one critical bug you never fix.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n700">gateway hostname</a>
<ul>
<li>We rule the world so we are above IETF and IANA. Now add our own
hostnames that of course won't add another assumption.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n689">no editor in systemd</a>
<ul>
<li>This one is a setback. Why is there no default editor in systemd in
case of factory reset?</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n610">8x ctrl + alt + del</a>
<ul>
<li>In systemd you press eight times Ctrl+Alt+Del to trigger reboot.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n594">privacy policy</a>
<ul>
<li>For the next generations: In 2015 privacy was a big issue because of
the mentioned hard-wiring between the web and software.
As you can see, every commit which adds some preparation for a 
feature adds another intepretation of what will be a major
assumption in a next release. If you
handle privacy you will have some features depending on that user
decision and of course the factory reset default value.</li>
<li>Why didn't they use XML for /etc/os-release?</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n502">fds cache</a>
<ul>
<li>We have talked about misdesign, too-big-to-fail and world
domination. This is the next example of a hack that is prone to
fail.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n380">umount -rf</a>
<ul>
<li>This is umount for dummies. Just do one thing – right.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n291">libudev will be orphaned</a>
<ul>
<li>With the advent of udevd there was a compatibility to its complexity
called libudev. X11 uses it to query the changing devices. And of
course make it a non-independent API in systemd. Why? You can guess
why: Defeating the infidels.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n281">fsck indirections</a>
<ul>
<li>When there was syslog fsck did output errors to your display as
easy as possible. Now add a hack to have this possible again.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n236">systemd-importd</a>
<ul>
<li>This is pure evil. Your pid 1 is now able to import complete system
images over the network and show them to you as your running system.
There is nothing that can go wrong.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n166">CGI for systemd</a>
<ul>
<li>The web thing has been discussed before.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n57">kdbus</a>
<ul>
<li>As of 2015-07-31 kdbus is not in the mainline Linux kernel. systemd
made kdbus non-optional in its release. The kernel maintainers are
still debating the kdbus ABI or possible alternatives, but if
systemd depends on the current state of kdbus the kernel maintainers
are faced with the hard decision to either break Fedora userspace or
accept the current kdbus proposal into the kernel with its security
and maintainability issues. This is the best example how
systemd is forcing you into decisions. Of course if you are  a mindless
bureaucrat it helps you to keep your job.</li>
</ul>
</li>
<li><a href="http://cgit.freedesktop.org/systemd/systemd/tree/NEWS?id=2d1ca11270e66777c90a449096203afebc37ec9c#n950">readahead removed</a>
<ul>
<li>The first thing swallowed in on Fedora was readahead. Now that (of
course!) everyone is using an SSD (at least the developers of
systemd do that) it can be removed. Why was it there? Is it possible
to make it a separate tool again?  There's no time for that, we are implementing
new features.</li>
</ul>
</li>
</ul>
<p>from 2015-07-31</p>
<h2>Practical systemd</h2>
<p>Here is what happens on a stock Arch Linux system, powered by systemd,
when a non-root user tries to restart the system:</p>
<pre><code>$ reboot
Failed to set wall message, ignoring: The name org.freedesktop.PolicyKit1 was not provided by any .service files
Failed to reboot system via logind: The name org.freedesktop.PolicyKit1 was not provided by any .service files
Failed to talk to init daemon.
</code></pre>
<p>In contrast, here is the equivalent error message on a system powered by
runit:</p>
<pre><code>$ reboot
init: fatal: unable to create /etc/runit/stopit: access denied
</code></pre>
<p>And on the oldest and best, Slackware:</p>
<pre><code>$ reboot
reboot: must be superuser.
</code></pre>
<p>systemd is driving &quot;just google the problem&quot; attitude, because how the hell are
you expected to troubleshoot this kind of error otherwise?</p>
<h2>Further Reading</h2>
<ul>
<li><a href="http://without-systemd.org/wiki/index.php/Main_Page">Without systemd</a></li>
<li><a href="http://without-systemd.org/wiki/index.php/Arguments_against_systemd">Arguments Against systemd</a></li>
</ul>
