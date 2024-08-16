ARCHIVED FROM 

https://pspodcasting.net/dan/blog/2019/plan9_desktop.html

accessed on friday thrid maay twenty twethy four

<h2>PLAN 9 DESKTOP GUIDE</h2>

<img src="thumbs/bunny_white.png" title="The Plan 9 mascot, glendy, sais it all: weird...">

<h3>INDEX</h3>
<ul>
<li><a href="#intro">What is Plan 9?</a></li>
<li><a href="#limitations">Limitations and Workarounds</a></li>
<ul>
<li><a href="#connecting">Connecting to Other Systems</a></li>
<ul>
<li><a href="#vnc">VNC</a></li>
<li><a href="#rdp">RDP</a></li>
<li><a href="#ssh">SSH</a></li>
<li><a href="#9p">9P</a></li>
<li><a href="#other_methods">Other methods</a></li>
</ul>
<li><a href="#porting">Porting Applications</a></li>
<li><a href="#emulation">Emulating other Operating Systems</a></li>
<li><a href="#virtualizing">Virtualizing other Operating Systems</a></li>
</ul>
<li><a href="#basics">Basics</a></li>
<ul>
<li><a href="#rio">Window Management</a></li>
<li><a href="#copy">Copy Pasting</a></li>
<li><a href="#programs">Essential Programs</a></li>
<li><a href="#text">Manipulating Text in the Terminal</a></li>
<li><a href="#acme">Acme - The Do It All Application</a></li>
<li><a href="#workspaces">Multiple Workspaces</a></li>
<li><a href="#tiling">Tiling Windows</a></li>
<li><a href="#plumbing">Plumbing</a></li>
</ul>
<li><a href="#sysadmin">System Administration</a></li>
<ul>
<li><a href="#sysadmin_basics">Basic System Administration</a></li>
<li><a href="#sysadmin_battery">Battery Monitoring</a></li>
<li><a href="#sysadmin_init">Configuring Startup and Shutdown</a></li>
<li><a href="#sysadmin_rio">Wallpapers, themes, etc...</a></li>
<li><a href="#sysadmin_language">Internationalization</a></li>
<li><a href="#sysadmin_users">User Management and Security</a></li>
<li><a href="#sysadmin_disks">Disk Management</a></li>
<li><a href="#sysadmin_backups">Backups</a></li>
<li><a href="#sysadmin_software">Package Management</a></li>
<li><a href="#sysadmin_filemng">File Management</a></li>
<li><a href="#sysadmin_unix">Tips for UNIX Sysadmins</a></li>
<li><a href="#sysadmin_cpu">Quick CPU+AUTH+Qemu+Drawterm HOWTO</a></li>
<ul>
<li><a href="#sysadmin_cpu_9front">9front</a></li>
<li><a href="#sysadmin_cpu_9legacy">9legacy</a></li>
<li><a href="#sysadmin_cpu_desktop">CPU+Rio desktop</a></li>
<li><a href="#sysadmin_cpu_terminals">CPU+PXE terminals</a></li>
</ul>
</ul>
<li><a href="#automation">Automation</a></li>
<ul>
<li><a href="#rc">Shell Scripting</a></li>
<li><a href="#rio_scripting">Rio Scripting</a></li>
<ul>
<li><a href="#script_scramble">Scrambling and Unscrambling a Rio Screen</a></li>
<li><a href="#script_max">max - Maximizing Windows</a></li>
<li><a href="#script_workspace">ws - Multiple Workspaces</a></li>
<li><a href="#script_tiling">tile - Tiling Window Manager</a></li>
</ul>
<li><a href="#acme_scripting">Acme Scripting</a></li>
<ul>
<li><a href="#acme_coffee">Coffee - Chill ASCII Animations</a></li>
<li><a href="#acme_slides">Slides - Acme Presentation Show</a></li>
<li><a href="#acme_chat">Chat - Simple Peer to Peer Chatting</a></li>
<li><a href="#acme_play">Play - Music Playlist in Acme</a></li>
</ul>
<li><a href="#web_scripting">Web Scripting</a></li>
<ul>
<li><a href="#9front_web_scripts">9front Web Scripts</a></li>
</ul>
</ul>
<li><a href="#development">Development</a></li>
<ul>
<li><a href="#version_control">Version Control</a></li>
<li><a href="#namespaces">Files and Namespaces</a></li>
</ul>
<li><a href="#web">The Web</a></li>
<ul>
<li><a href="#wifi">Wireless Network</a></li>
<li><a href="#browsing">Browsing The Web</a></li>
<li><a href="#download">Downloading</a></li>
<li><a href="#email">Email</a></li>
<li><a href="#chat">Chatting</a></li>
<li><a href="#werc">Running a Web Server</a></li>
</ul>
<li><a href="#media">Multimedia</a></li>
<ul>
<li><a href="#audio">Audio</a></li>
<li><a href="#video">Video</a></li>
<ul>
<li><a href="#youtube">Youtube</a></li>
</ul>
</ul>
<li><a href="#graphic">Graphics</a></li>
<ul>
<li><a href="#page">Viewing Images/Documents</a></li>
<ul>
<li><a href="#comic">Reading Comics</a></li>
</ul>
<li><a href="#paint">Creating Images</a></li>
<li><a href="#screenshot">Taking a Screenshot</a></li>
<ul>
<li><a href="#screencasting">Screencasting</a></li>
</ul>
</ul>
<li><a href="#peripherals">Peripherals</a></li>
<ul>
<li><a href="#memstick">USB sticks</a></li>
<li><a href="#cds">CD/DVD/BS's</a></li>
<li><a href="#printers">Printers</a></li>
</ul>
<li><a href="#games">Games and other Fun Stuff</a></li>
<ul>
<li><a href="#included_games">Included Games</a></li>
<li><a href="#included_emulators">Included Game Emulators</a></li>
<li><a href="#extra_games">3rd Party Games</a></li>
<li><a href="#edutainment">Edutainment</a></li>
<ul>
<li><a href="#edutainment_arithmetic">Arithmetic</a></li>
<li><a href="#edutainment_quiz">Quiz</a></li>
<li><a href="#edutainment_touchtype">Touchtype</a></li>

</ul>
<li><a href="#telnet_fun">Playing With Telnet</a></li>
<li><a href="#misc_fun">Miscellaneous Fun</a></li>
<li><a href="#os_fun">Obscure Operating Systems</a></li>
<ul>
<li><a href="#inferno">Inferno</a></li>
<li><a href="#unix_v8">UNIX V8</a></li>
</ul>
</ul>
<li><a href="#office">Office</a></li>
<ul>
<li><a href="#office_reading">Reading Office Documents</a></li>
<ul>
<li><a href="#epub">Reading Epubs</a></li>
</ul>
<li><a href="#office_writing">Writing Office Documents</a></li>
<ul>
<li><a href="#office_troff">Tweaking Troff Macros</a></li>
</ul>
<li><a href="#office_spell">Spellchecking</a></li>
<li><a href="#office_pim">PIM</a></li>
<ul>
<li><a href="#pim_2do">2do Lists</a></li>
<li><a href="#pim_que">Queues</a></li>
<li><a href="#pim_pass">Password Management</a></li>
<li><a href="#pim_account">Personal Accounting</a></li>
<li><a href="#pim_time">Time Management</a></li>
</ul>
<li><a href="#office_math">Math, Graphs and Units</a></li>
<li><a href="#office_sheet">Spreadsheets</a></li>
<li><a href="#office_db">Databases</a></li>
<ul>
<li><a href="#office_db_awk">Awk as a Database</a></li>
<li><a href="#office_db_ndb">Ndb as a Database</a></li>
</ul>
</ul>
<li><a href="#conclusion">Conclusion</a></li>
</ul>

<h2 id="intro">What is Plan 9..?</h2>

<p>
Briefly,
<i>Plan 9 from Bell Labs</i> is a computer operating system developed by the original UNIX design team.
After decades of work on Research UNIX in the late 80's,
the team decided to write a new operating system from scratch,
Plan 9 was finally released in 1992,
and a few years later they released yet another operating system called
<a href="https://github.com/inferno-os/inferno-os">Inferno</a>,
which share many of the same characteristics as its sister project.
These systems,
and variations thereof,
have more or less been in continual development since.
The history and design philosophy behind these operating systems,
is interesting,
but we will not talk about that here.
Instead,
this article will focus on the practical aspects of using Plan 9 as day-to-day desktop system.
</p>

<p>
Beware that prior exposure to UNIX is a double-edged sword.
There are similar sounding commands and conventions between the two platforms,
and Plan 9 does follow the UNIX philosophy
(much more so then UNIX in fact).
Nevertheless,
Plan 9 is <i>not</i> UNIX!
It is an operating system written entirely from scratch,
backwards compatibility was not a goal.
If you expect just another Ubuntu spin-off,
you will be very disappointed.
In fact,
lets be clear here:
You <i>will</i> be disappointed,
period.
Now with that disclaimer out of the way,
lets have some fun!
</p>

<p>
In 2002 the 4th edition of Plan 9 was released,
it was essentially a rolling release,
that continued to receive updated from Bell Labs until 2015,
when the project was officially discontinued.
In mid 2021 though,
Bell Labs gave ownership of all previous Plan 9 sources to the
<a href="https://p9f.org">Plan 9 foundation</a>.
The goal of this foundation is to continue the development of Plan 9,
but so far,
not much has happened.
There are several community forks around though,
two of them,
<a href="https://9legacy.org">9legacy</a>
and
<a href="https://9front.org">9front</a>,
sprang into existence around 2010.
If you want to use Plan 9 as a day-to-day desktop,
which will be the focus of this article,
I strongly recommend going with 9front.
It is likely the only candidate that will actually run on your physical hardware,
and it has many features that a modern user takes for granted,
such as auto-mounting USB sticks,
wifi support,
working audio,
video playback and <b>git</b>.
9front has an excellent
<a href="https://fqa.9front.org">fqa</a>
and
<a href="https://wiki.9front.org">community wiki</a>,
that do a far better job of presenting accurate information then I do
(be prepared for quirky humor though!).
Still,
it can be interesting to play with 9legacy too,
if only for historical curiosity,
so I will give some pointers in this article on "classic Plan 9"
(9legacy and the old 4th edition of Plan 9 are nearly identical),
where it differs significantly from 9front.
For classic Plan 9,
the Plan 9
<a href="https://9p.io/wiki/plan9/plan_9_wiki">wiki</a>
from Bell Labs is a better source of documentation then the 9front resources.
</p>

<h2 id="limitations">Limitations and workarounds</h2>

<p>
More then anything,
Plan 9 is a <i>simple</i> operating system.
The kernel is only 200,000 lines of code,
and the userland about a million.
In comparison the source code for the Firefox web browser is more than 24 million lines of code!
As you might imagine then,
there are no "modern" web browsers in Plan 9.
There are no office suits,
triple A games,
VOIP
or repositories of 30,000 pre-compiled packages.
Plan 9 is not for the faint of heart!
</p>

<p>
Of course there are workarounds for the above limitations,
here are a few suggestions:
</p>

<h3 id="connecting">Connecting to Other Systems</h3>

<h4 id="vnc">VNC</h4>

<p>
It is simple enough to connect to a remote UNIX/Windows machine from Plan 9 using VNC,
or vice versa
(I use the term "UNIX" broadly - it includes Mac,
Android, Linux, BSD, etc...).
From Plan 9 you can connect to a VNC server using <b>vncv</b>,
or run a VNC server with <b>vncs</b>
(there is little reason to run a VNC server on Plan 9 though,
use <b>drawterm</b>,
mentioned below,
instead).
</p>

<p>
For example,
assuming you have <b>tigervnc</b> installed on a UNIX machine,
with the ip address <b>192.168.0.1</b>,
and a desired VNC screen resolution of 1366 x 768 pixels:
You can run <b>vncserver -geometry 1366x768 :1</b>,
and give it a login password
(if you are not prompted for a password you may need to run <b>vncpasswd</b> first).
Now,
on the Plan 9 machine,
run the command <b>vncv 192.168.0.1:1</b>,
and login.
By default this will probably run a very basic <b>twm</b> desktop,
which makes many inexperienced users suspect that the desktop failed somehow.
You probably want to change <b>~/.vnc/xstartup</b>,
to run a fancier window manager.
To use <b>openbox</b> instead of <b>twm</b> for instance,
add this line to the file:
</p>

<p>
<b>
exec /usr/bin/openbox-session
</b>
</p>

<p>
You can choose whatever desktop you want here,
but beware that configuring <b>xstartup</b> gets <i>vastly</i> more complex if you use some bloated mess like Gnome or KDE.
</p>

<h4 id="rdp">RDP</h4>

<p>
It is possible to connect to a remote Windows machine using RDP,
see
<a href="http://shithub.us/covertusername/rd/HEAD/info.html">rd</a>
if you need that sort of thing.
</p>

<h4 id="ssh">SSH</h4>

<p>
9front ships with a working <b>ssh</b> and <b>sshfs</b> client
(<b>sshfs</b> mounts the remote file system in <b>/n/ssh</b>),
but classic Plan 9 has a very outdated version of <b>ssh</b>,
that in all likelihood will not
(or at least <i>should</i> not)
be able to connect to your UNIX machines.
</p>

<h4 id="9p">9P</h4>

<p>
It is in fact much easier to import Plan 9 technologies to foreign systems then vice versa,
and there are good solutions for working with Plan 9 from UNIX.
We will discuss technologies such as <b>plan9port</b> and <b>drawterm</b> later,
but for now,
lets talk about mounting the Plan 9 file system natively in UNIX using the 9P protocol.
There are various ways you can do this,
including mounting it directly,
in Linux at least,
like so:
<b>sudo mount -o rw -t 9p 192.168.0.1 /mnt</b>
(substitute the ip address for the Plan 9 machine you're using).
But you will probably get better results using one of the many 9P clients that's out there,
such as
<b>9pfuse</b> from the <b>plan9port</b> package,
or <a href="https://github.com/ftrvxmtrx/9pfs">9pfs</a>
You can use it like so: <b>9pfs 192.168.0.1 /mnt</b>,
assuming you have the right privileges.
</p>

<h4 id="other_methods">Other methods</h4>

<p>
There is some support for NFS and SMB in Plan 9
(see <b>nfs</b>(4) and <b>cifs</b>(4)),
but I don't recommend using NFS,
the Plan 9 client is <i>very</i> outdated.
Speaking of outdated,
you naturally have <b>ftpfs</b> and <b>telnet</b> as well.
</p>

<h3 id="porting">Porting applications</h3>

<p>
There exists a fairly good port of Plan 9 userland programs and services for UNIX,
called <a href="https://9fans.github.io/plan9port">Plan9Port</a> (or <b>p9p</b> for short - a more lightweight alternative is <a href="https://tools.suckless.org/9base"><b>9base</b></a>),
it is available in the repositories of most popular UNIX systems.
Once installed,
use the <b>9</b> command to run the Plan9Port programs rather then the UNIX counterparts,
eg. <b>9 acme</b>.
It does not fully replicate the Plan 9 experience of course,
but it does make UNIX less of a pain to use.
</p>

<p>
To run a full Plan 9 shell,
using Plan9Port commands instead of the UNIX equivalents,
either run <b>9 acme</b>,
execute <b>win</b> in it and run <b>9 rc</b>.
Or run <b>9 9term</b>,
then run <b>9 rc</b>.
You can configure your <b>~/.xinitrc</b> file to start the Plan 9 look-alike window manager,
with <b>exec 9 rio</b>,
and set up a very authentic looking Plan 9 desktop.
But there is little point in doing so,
unless you really want to hide the fact that UNIX is running in the background.
Plan9Port's <b>rio</b> only <i>looks</i> like the Plan 9 window manager,
but it doesn't have the same useful features,
and it is quite flaky to boot.
In my opinion there are far better native UNIX alternatives,
including the Plan9 inspired <b>wmii</b>/<b>dwm</b> window managers,
or variations thereof.
</p>

<p>
It is possible,
but <i>much</i> harder,
to go in the other direction.
Plan 9 has a UNIX compatibility suit of programs and libraries in <b>/bin/ape</b>,
such as <b>ape/sh</b>,
which gives you a <b>ksh</b> like UNIX shell
(run <b>vt</b> first to emulate a VT-100 terminal).
And <b>ape/cc</b> a POSIX compliant C compiler,
with corresponding UNIX-friendly libraries.
Plus a few other UNIX'y utilities.
This UNIX compatibility is old and quite unmaintained.
9front has its own semi-official portability layer called <b>npe</b>,
see the 9front
<a href="http://docs.9front.org/porting">porting guide</a>
for further tips.
</p>

<p>
Note however that simply having a UNIX shell,
does not mean that all your shell scripts will magically work.
Plan 9 has it's own version of <b>cat</b>,
<b>echo</b>,
<b>ls</b>,
<b>sed</b> and so on.
If your script uses these programs,
it needs to be adapted to use the Plan 9 versions of them.
As always,
read the man pages carefully
(no really - <i>read</i> them!).
</p>

<p>
Finally,
even though Plan 9 <s>has</s> had a very good POSIX compliance,
it's by no means certain that UNIX programs will compile.
Most will not.
The majority of UNIX software does not restrict themselves to POSIX alone,
but require large extensions.
Most of which are not supported.
For example,
Plan 9 does not have X (by default),
curses,
sockets,
numerical UID/GID's or links,
so any programs depending on such things needs to be patched and rewritten before they will work.
In practice only the simplest of programs can be ported with any reasonable amount of effort.
</p>

<h3 id="emulation">Emulating other operating systems</h3>

<a href="images/emulation.png"><img src="thumbs/emulation.png" title="A FreeBSD desktop with P9P acme and a drawterm running a FreeBSD desktop in Plan 9 with VNC"></a>

<p>
In a traditional Plan 9 network,
one or more CPU servers are providing file and authentication services to multiple diskless workstations,
called "terminals".
These terminals are desktop systems connected to the CPU server.
This is a bit confusing for UNIX users,
so in this article we will refer to a diskless workstation as a remote desktop,
and a window running a shell as a terminal,
as is the custom in UNIX.
If you have set up a CPU Server in Plan 9
(see section
<a href="https://fqa.9front.org/fqa7.html#7.5">7.5</a>
and
<a href="https://fqa.9front.org/fqa7.html#7.6">7.6</a>
in the 9front fqa -
see also
<a href="#sysadmin_cpu">Quick CPU+AUTH+Qemu+Drawterm HOWTO</a>
below),
either physically,
or on a virtual machine,
you can emulate a Plan 9 remote desktop on a UNIX/Windows machine with <a href="https://drawterm.9front.org"><b>drawterm</b></a>
(for classic Plan 9 use
<a href="https://swtch.com/drawterm">this link</a>).
<b>drawterm</b> works very well,
it also has access to the host file system under <b>/mnt/term</b>,
making it easy to work on files across operating systems.
</p>

<p>
There is a 3rd party port of
<a href="https://9p.io/sources/extra/X11.iso.bz2">X</a>
for Plan 9,
together with <b>linuxemu</b>,
it can be used to run Linux software natively
(see section
<a href="https://fqa.9front.org/fqa8.html#8.7.1">8.7.1</a>
in the 9front fqa).
This implementation is not perfect however,
it is old and tedious to work with,
and I have had little success with it myself.
</p>

<h3 id="virtualizing">Virtualizing other operating systems</h3>

<p>
<a href="images/oses.jpg"><img src="thumbs/oses.jpg" title="You can have a menagerie of OS's in Plan 9"></a>
</p>

<p>
There are many different virtualization solutions available for UNIX/Windows capable of running Plan 9,
such as <b>qemu</b> and <b>VirtualBox</b>.
Plan 9 has very limited hardware support,
especially if you want to use the classic versions of this operating system.
Virtualization is a practical way to eliminate such concerns.
</p>

<p>
9front also includes its own hypervisor
(see <a href="https://fqa.9front.org/fqa8.html#8.7.5">section 8.7.5</a> in their fqa),
<b>vmx</b>,
capable of running Linux,
OpenBSD,
allegedly Windows,
and plausibly other operating systems.
<i>PS:</i> You need modern Intel hardware for this to work.
</p>

<h2 id="basics">Basics</h2>

<p>
I assume you have already downloaded and installed Plan 9,
either on a physical machine or on a virtual one.
If not you can get the <a href="https://9front.org">9front iso</a>,
and follow the installation instructions in <a href="https://fqa.9front.org/fqa4.html">section 4</a> of their fqa.
Again,
this is not a guide for installing and configuring a Plan 9 system,
use the 9front fqa for that.
Our focus here is doing day-to-day tasks after the initial setup is done.
</p>

<p>
<i>PS:</i> This is also the subject of section 8 in the 9front fqa -
<a href="https://fqa.9front.org/fqa8.html">Using 9front</a>.
This article simply repeats and expands upon some of that content.
</p>

<p>
<i>PS:</i> If you want to install 9legacy,
it follows much the same steps as 9front,
but here are a couple of tips:
After hitting <b>Return</b> at the <b>"Location of archives [browse]:"</b> prompt,
you will see <b>/%</b>,
just type <b>exit</b> to continue the installation.
Choose <b>plan9</b> when asked to <b>"Enable boot method"</b>,
otherwise just follow the defaults and choose <b>"y"</b> at yes or no prompts.
Finally: when installing 9legacy in <b>qemu</b>,
be sure to set the virtual harddisk as the first disk drive,
eg. <b>qemu-system-x86 -m 2G -hda 9legacy.img</b>,
<i>do not</i> use <b>-hdd</b> or similar,
otherwise boot setup will fail during installation.
</p>

<h3 id="rio">Window Management</h3>

<p>
The window manager in Plan 9 is called <b>rio</b>,
it provides a remarkably clean and simple desktop,
somewhat akin to <b>twm</b> in UNIX.
Unlike <b>twm</b> though,
it doesn't look like crap by default,
and the source code is only 6000 lines of code,
which incidentally is also about the same size as Plan 9's graphical library,
libdraw.
In contrast <b>twm's</b> source is closer to 30,000 lines,
and the X Window System backend,
more then 8 million!
</p>

<p>
Window management is straight forward:
<b>rio</b> provides only one menu,
which you can access by right clicking the mouse on the desktop background.
Hold down the mouse button while you are selecting a menu option,
and release the mouse button only after you have made your choice.
To create a new window,
which is always a terminal,
choose <b>New</b>.
The mouse pointer changes to a cross.
Right click in a corner and drag the mouse,
a red rectangular box appears,
release the mouse button when the window has the size you want.
</p>

<p>
If you choose the <b>Delete</b> option in the <b>rio</b> menu,
the mouse pointer changes to a cross with a circle.
Right click on the window you wish to delete.
If you <b>Hide</b> a window,
it will appear in the <b>rio</b> menu,
select it from the menu to make it visible again.
</p>

<p>
You can also <b>Resize</b> and <b>Move</b> a window by using the <b>rio</b> menu,
but it's easier to click and drag the windows directly:
To resize a window,
left click the blue border and drag,
to move it,
right click and drag.
</p>

<p>
Right clicking in a terminal window will also bring up the <b>rio</b> menu,
but other programs will not necessarily do so.
If you need to access the window manager menu while running a fullscreen <b>acme</b> window for instance,
you must first shrink the window or move it out of the way,
and then right click the gray <b>rio</b> background.
By default there are no key-bindings to control <b>rio</b>,
you can only do so using the mouse
(<i>What?!?</i> Mouse actions are <i>required?!?</i> I know right, Plan 9 is so radical - take a look at the <a href="#workspaces">workspaces</a> section below though).
</p>

<h3 id="copy">Copy Pasting</h3>

<p>
In order to use Plan 9 effectively,
you need a 3-button mouse.
Such mice are quite common nowadays,
with the scroll wheel doubling as the middle mouse button
(for laptops I recommend ThinkPads).
The 3 mouse buttons,
and combinations of clicks,
are used throughout Plan 9 for manipulating text.
If you don't have a mouse with 3 buttons,
you can simulate the middle click by holding down the <b>Shift</b> key and right clicking.
But this will quickly become tedious,
so go out and buy a 3-button mouse ASAP.
</p>

<p>
You can select text in the normal way,
by left click and drag.
You can also double left click a word to select it.
If you double click the end of a line,
the whole line will be selected,
and if you double click a parenthesis,
or square bracket or some such delimiter,
the text inside these parenthesis will be selected.
</p>

<p>
To cut the selected text,
hold down the left button and click the middle mouse button.
To paste the text,
click the left button and while holding it down,
click the right button.
To "copy" text,
left click and middle click,
release the middle mouse button and click the right button.
Such combinations of mouse clicks are called mouse chording.
They are used very consistently in Plan 9 programs,
and feel intuitive enough once you get the hang of it.
</p>

<h3 id="programs">Essential Programs</h3>

<a href="images/apps.png"><img src="thumbs/apps.png" title="Some common Plan 9 applications"></a>

<p>
There are only a handful of programs in Plan 9,
they are simple to learn and work very well.
Some essential applications are:
</p>

<p>
<ul>
<li><b>rio</b> the window manager</li>
<li><b>rc</b> the shell</li>
<li><b>acme</b> a text editor, and more!</li>
<li><b>mothra</b> the web browser (use <b>abaco</B> in classic Plan 9)</li>
<li><b>page</b> a document/image viewer</li>
<li><b>play</b>, <b>zuke</b> music players (use <b>juke</b> in classic Plan 9)</li>
<li><b>stats</b> monitoring system load</li>
</ul>
</p>

<h3 id="text">Manipulating Text in the Terminal</h3>

<p>
You do not have to play around much in the Plan 9 terminal before you realize that it works quite differently from UNIX.
One surprise is that terminals do not auto scroll,
if you <b>cat</b> a very long file for instance,
it will just display the first screenful of text,
and wait for you to manually hit <b>PgDn</b> or the <b>down</b> arrow key.
This behavior is actually very convenient,
since it removes the need for pagers.
But sometimes it can cut against you.
If you're compiling software for instance,
the compilation will stop once the text has filled the screen,
and only continue if you manually scroll down.
Clearly,
this is not what you want!
Middle click the terminal window and select <b>scroll</b>,
it will now auto scroll,
just as UNIX terminals do.
You can go back to the default behavior again,
by middle clicking and selecting <b>noscroll</b>.
</p>

<p>
Another annoyance might be that there is no <b>Tab</b> auto-completion.
Don't worry,
use <b>Ctrl-f</b> instead,
it does much the same thing.
There is no advanced auto-completion of program names and flags,
like <b>zsh</b> and <b>fish</b> users might be accustomed to.
But this really isn't an issue since Plan 9 has virtually no programs or flags to speak of,
as you will discover soon enough.
</p>

<p>
The third thing you may notice is that the terminal text can be freely edited.
You can add any text anywhere and copy paste the text arbitrarily,
the Plan 9 terminal thus feels much more like a text editor then a UNIX terminal
(a consequence of this free-form text editing is that the mouse cursor has to be put at the end of the last line in order to execute a command with the <b>Return</b> key,
otherwise it will just add a literal newline to the text
- this is only mildly annoying once you get used to it).
What's the point of this novel design?
First of all it eliminates a host of special purpose programs that UNIX requires,
for example there is no <b>clear</b> command in Plan 9,
you just cut the text.
There is no <b>reset</b> or <b>readline</b> either,
as they are not needed.
Secondly,
once learned,
this behavior feels very intuitive.
Why shouldn't you be able to cut and paste text and freely sprinkle your terminal output with random comments?
Going back to a UNIX terminal,
after having spent some time in Plan 9,
really feels like leaving the 90's - and going back to the 70's
(fun tip:
check out <b>/bin/hold</b> to see how a basic text editor in Plan 9 can be written in just five lines of shell script!).
</p>

<p>
Lastly,
there is no history command in the Plan 9 terminal,
hitting the <b>up</b> arrow key on the keyboard will just move the pointer one line up,
like any text editor would.
- What else did you expect?
Relax though,
you can rerun the previous command with <b>""</b>
(<b>"</b> will reprint it).
</p>

<p>
Hang on!
The <i>command</i> <b>""</b>,
isn't double quotes used for quoting?!?
Not in Plan 9,
double quotes are just ordinary characters.
Whereas UNIX has three escape characters,
Plan 9 has only one,
the single quote
(well, OK, backslash is also used in <i>some</i> situations).
The UNIX command <b>"$message has a literal \$ and ' sign"</b>,
would be <b>''$message' has a literal $ and '' sign'</b> in Plan 9
(two single quotes within single quotes is interpreted as one literal single quote).
</p>

<p>
<i>PS:</i> <b>"</b> and <b>""</b> are actually shell scripts,
provided by 9front,
classic Plan 9 systems do not have these.
</p>

<p>
Back to our topic of rerunning commands,
note that the need to auto-complete text and rerun commands are much greater in UNIX then in Plan 9.
It is easy to copy paste text in the terminal,
so use that functionality for what it's worth!
You don't need to use insane syntax like <b>ls !$</b> to run <b>ls</b> with the previous arguments,
or <b>^foo^bar</b> to spell correct the last argument and rerun it.
Just type <b>ls</b> in the terminal and copy paste the previous arguments,
and if you need to spell correct the last argument,
then just do so,
copy paste the result when you're done.
There is also a full copy of the terminal text in <b>/dev/text</b>.
So the command <b>cat /dev/text &gt; transcript</b> is essentially the same as <b>script</b> in UNIX,
<b> &gt; /dev/text</b> is basically <b>clear</b>,
and the command <b>grep '^; ' /dev/text</b> the same as <b>history</b>
(assuming of course that your shell prompt is <b>; </b>).
Note that you can search this log for other things then just your previous commands,
and you can manipulate this data in many other interesting ways as well.
For example,
need to do advanced searching or manipulation of the shell history?
Just open <b>/dev/text</b> in a text editor,
eg. <b>sam /dev/text</b>.
</p>

<p>
But what if you want a system wide history log for all of your windows?
There is no such file in Plan 9,
but it's easy enough to make one.
For example,
the following script will save your command history to a central file.
Only unique commands are saved,
if we saved all of the text,
our central history file would grow extremely large.
For example,
it would be quite redundant to have ten thousand entries of <b>cd</b> in our history log,
not to mention hundreds of copies of the manpages and text files we have been reading.
</p>

<p>
<pre>
#!/bin/rc
# savehist - prune and save command history
# usage: savehist

# set some defaults
rfork ne
temp=/tmp/savehist-$pid
hist=$home/lib/text
touch $hist

# rewrite history
cat &lt;{grep '^; ' /dev/text} $hist | sort | uniq &gt; $temp
mv $temp $hist
</pre>
</p>

<p>
With this in place we can run <b>savehist</b> before <b>exit</b> to save our current history,
or we can wrap these steps into one by adding something like this to our
<b>$home/lib/profile</b>:
<b>fn quit{ savehist; exit }</b>
(<i>PS:</i> Don't call this function <b>exit</b> unless you really want a fork bomb!)
</p>

<p>
In addition to <b>/dev/text</b> you also have <b>/dev/snarf</b>,
which holds the "snarf" buffer,
the clipboard in Plan 9 speak
(if you want to write to the window, use <b>/dev/cons</b>).
All of these files refer to your current window,
if you want to use these files for a different window,
see the <a href="#rio_scripting">rio scripting</a> section below.
</p>

<p>
The graphical desktop runs "within" the text console in Plan 9,
so writing to the system console will actually print the text verbatim onto the screen.
For example,
running <b>sleep 600; echo Bug Me! &gt; '#c/cons'</b>
will send a fairly obtrusive notification to your screen in 10 minutes.
This can be a bit disconcerting for a beginner,
but it's easy to redirect such messages if you don't want them to clutter up your screen.
Just run <b>cat /dev/kprint</b> in a window and hide it.
See the
<a href="#rio_scripting">rio scripting</a> section below,
for some ideas on how to avoid or abuse this functionality further.
</p>

<h3 id="acme">Acme - The Do It All Application</h3>

<a href="images/acme.png"><img src="thumbs/acme.png" title="Acme running a Mail client to the left, a shell in the middle, and a file manager and text editor to the left"></a>

<p>
The <b>acme</b> text editor is arguably the main user application for Plan 9,
it doubles as the systems file manager,
terminal,
mail reader and more.
It can even be used as a fully fledged window manager,
by replacing <b>rio</b> with <b>acme</b> in your <b>$home/lib/profile</b>
(but I don't recommend it - you will not be able to run any other programs - then again, why would you want to?).
</p>

<p>
Let's do a whirlwind tour of <b>acme</b>:
The first blue row contains commands for the entire <b>acme</b> window,
such as <b>Exit</b>,
if you middle-click this button,
<b>acme</b> will exit.
<b>Dump</b> will create a file called <b>acme.dump</b>,
this can be used to save a particular window arrangement,
and restored with <b>acme -l acme.dump</b>.
<b>Putall</b> will save all modified text files.
</p>

<p>
If you middle-click <b>Newcol</b> a new column will appear.
The column has it's own row of commands,
in the second blue row.
<b>Delcol</b> will delete the column.
<b>Cut</b>,
<b>Paste</b> and
<b>Snarf</b> (eg. "Copy"),
will do text manipulations.
But it's easier to use mouse chords for this:
Left and middle-click to Cut,
Left and right-click to Paste,
and finally Left and middle-click,
then right-click to Snarf,
or Copy.
The mouse chords are awkward to explain,
but try it out,
it will feel very intuitive with a little practice.
</p>

<p>
Middle-clicking <b>New</b> will create a new window in
the column.
Again,
it too,
will have it's own row of commands.
<b>Del</b> will delete the window.
The window is initially empty,
try writing some random text into it.
You will see that a new command appears,
<b>Undo</b>
(it's meaning should be obvious).
After typing in some text,
you can also hit the <b>Esc</b> key to mark the recently added text,
hitting <b>Esc</b> again,
will cut the text.
How do we save our file?
First we need to give it a name:
Click on the far left side of the menu,
left of <b>Del</b>,
and type <b>/usr/glenda/testfile</b>
(<b>glenda</b> is the default user in Plan 9,
and <b>/usr/glenda</b> is the default home directory).
Yet another command will appear,
<b>Put</b>,
middle click it to save your work.
That was a <i>lot</i> of typing!
Isn't there an easier way to do this?
Sure,
remember that Plan 9 allows you to copy paste pretty much anything.
Find the directory you want in a terminal,
with <b>Ctrl-f</b> auto-completion and everything,
then print the directory name with <b>pwd</b>,
and just copy paste that into the <b>acme</b> window,
and append your new filename.
Easier yet,
run <b>touch testfile; B testfile</b> from a terminal and the file will be opened for you in <b>acme</b>.
</p>

<p>
By now you will have noticed a very unique feature of <b>acme</b>,
it's menus are pure text.
The "buttons" are just regular words.
To illustrate:
Type <b>Del</b> (case sensitive!) somewhere in the yellow text window,
then middle click it.
The window will disappear.
<b>Del</b> is just a command,
same as <b>echo</b> or <b>cat</b>.
Another test:
Type <b>echo hi there</b> and middle click,
and drag,
so that the red mark covers all three words.
<b>hi there</b> will be printed in a new window.
</p>

<p>
You can use the <b>Look</b> command to search for words in the window.
Type <b>monkey</b> a couple of times in the yellow text window,
now type <b>Look monkey</b> in the blue window menu,
and middle click and drag,
to mark the two words.
The first occurrence of <b>monkey</b> will be highlighted,
run the command again,
and the second occurrence will be highlighted,
and so on.
An easier method however would be to just right-click the word <b>monkey</b>,
anywhere in the text,
the next occurrence of the word will be highlighted,
and the mouse pointer will be moved there.
Just right-click again to see the next occurrence of the word,
and so on.
</p>

<p>
The <b>Zerox</b> command in the column menu will duplicate a window,
this is very useful if you are editing a long file,
and you need to see or edit different parts of the file at the same time,
any changes made in one window will appear in the other.
<b>Sort</b> will sort the column windows by name,
it does not sort the content of the windows.
To do that,
mark the text,
type <b>|sort</b> in the window menu,
and middle-click it.
As you can see,
you can freely use arbitrary Plan 9 commands to manipulate the text in <b>acme</b>.
</p>

<p>
If you want to do search and replace operations,
use the <b>Edit</b> command.
This command is a back end for the <b>sam</b> text editor,
which uses much the same text editing commands as <b>ed</b>
(which again is similar to <b>sed</b> or <b>vi</b>).
For example,
double click one of the <b>monkey</b> words to highlight it,
then type <b>Edit s/monkey/chimpanzee/</b>,
and middle click and drag to execute this command.
The highlighted word will be changed to <b>chimpanzee</b>.
To change all the occurrences of <b>monkey</b>,
type
<b>Edit ,s/monkey/chimpanzee/g</b>
(in <b>vi</b> this would be <b>:%s/monkey/chimpanzee/g</b>).
</p>

<p>
Side note:
Although the above <b>ed</b> style substitution works in <b>sam</b>,
<b>sam</b> is not a line-based editor like <b>ed</b>,
and a more proper <b>sam</b> command for the above would be:
<b>Edit ,x/monkey/c/chimpanzee/</b>
(that is: for each <b>/monkey/</b> change to <b>/chimpanzee/</b>).
To read the <b>sam</b> tutorial,
run: <b>page /sys/doc/sam/sam.tut.out</b>
</p>

<p>
<b>acme</b> lacks many built-in features that a UNIX user might expect,
but you can create much of this functionality simply by piping the text through standard utilities.
Here are some examples:
</p>

<p>
<ul>
<li><b>Edit =</b> print current line number</li>
<li><b>Edit ,|sort -r</b> reverse sort the file</li>
<li><b>Edit ,|grep -n .</b> add line numbers</li>
<li><b>Edit ,s/^.*: //g</b> remove line numbers</li>
<li><b>Edit s/^/    /g</b> indent text</li>
<li><b>Edit s/^    //g</b> unindent text</li>
<li><b>Edit s/^/#/g</b> comment out lines of code</li>
<li><b>Edit s/^#//g</b> uncomment lines of code</li>
<li><b>Edit ,|wc -c</b> file word count</li>
<li><b>Edit ,|fmt</b> nicely format the file</li>
<li><b>Edit ,|cb</b> beautify C source code</li>
<li><b>Edit s/./-/g</b> underline after copying a line</li>
<li><b>|tr A-Z a-z</b> lowercase text</li>
<li><b>|tr a-z A-Z</b> uppercase text</li>
<li><b>|tr a-zA-Z n-za-mN-ZA-M</b> rot13 text</li>
</ul>
</p>

<p>
Open a <b>New</b> window and type in the filename <b>/usr/glenda</b> to the far left,
then type <b>Get</b> to the far right,
right of <b>Look</b>,
and middle click it.
The contents of the <b>/usr/glenda</b> directory will fill the window.
If you right-click on a directory in this window,
the contents of that directory will be opened in a new window.
To do operations on files,
just type a command and execute;
for example type <b>rm</b> before <b>testfile</b>,
and middle click the two words to remove this file.
If you right-click a text file,
the contents of that file will be opened for editing in <b>acme</b>.
</p>

<p>
Exactly what happens when you right-click something in <b>acme</b>,
depends on the word you click.
For example clicking on the word <b>/usr/glenda/pictures/cirno.png</b>,
will open this picture in the image viewer <b>page</b>,
and clicking <b>jazz.mp3</b>,
will start playing the audio file with <b>play</b>.
Provided of course that the files in question exist on your system.
The last example also assumes that the <b>jazz.mp3</b> file is located in the same directory as the one you launched <b>acme</b> from,
if not you need to specify a correct file path.
The actual work of connecting the right words to the right programs is handled by <b>plumber</b>,
which we will talk about later,
but for now it's enough to know that right clicking a filename anywhere in <b>acme</b> will usually just "do the right thing"
(you'll note though that actions are evaluated for <i>words</i>,
not files).
</p>

<p>
Each window has a dark blue square to the far left of the menu,
you can click and drag this box to resize or move the window to another column.
The columns themselves also have a dark blue square,
click and drag this to resize or move the column.
</p>

<p>
You can also right-click on the dark blue window square,
to hide all the column windows except that one,
left-click on it to bring the windows back.
Left-clicking on the square will increase the window size a little,
middle-clicking will maximize the window.
</p>

<p>
Left-clicking on the scroll bar will scroll upwards,
right-clicking downwards.
Clicking towards the bottom of the scroll bar will scroll a lot,
clicking towards the top will only scroll a little.
Middle clicking will transport you directly to that portion of the file.
Play around and experiment with these mouse actions,
pretty soon you will get the hang of it.
Other Plan 9 applications with scroll bars work the same way
(in 9front at least).
</p>

<h3 id="workspaces">Multiple Workspaces</h3>

<a href="images/rio.png"><img src="thumbs/rio.png" title="five hidden rios inside rio inside rio inside rio inside rio..."></a>

<p>
<b>rio</b> does not have virtual workspaces per se,
but 9front includes some tools that let you set up a keyboard driven desktop with pseudo-workspaces.
For instance,
you can add the following snippet to your <b>$home/lib/profile</b>:
</p>

<p>
<pre>
fn workspaces{
</dev/kbdtap riow >/dev/kbdtap |[3] bar
}
</pre>
</p>

<p>
You can now run <b>workspaces</b> and switch to a new "workspace" with <b>Win</b>+<b>n</b>
(where <b>Win</b> is the Windows key next,
between the left <b>Ctrl</b> and <b>Alt</b> keys,
and <b>n</b> is a number between 0 and 9).
You can also move windows about with <b>Win</b>+<b>Arrow</b>,
or resize them with <b>Win</b>+<b>Shift</b>+<b>Arrow</b>
(see <b>riow</b>(1) and <b>bar</b>(1) for more info).
Classic Plan 9 does not have these tools,
although there is an old fork of <b>rio</b> called <b>rio-virtual</b>,
that does include virtual workspaces.
There are also ways to create such services without the 9front extensions:
You'll note that all windows in all <b>riow</b> "workspaces" are listed in the <b>rio</b> menu and can be unhidden.
This is because <b>riow</b> doesn't actually add workspaces as such,
but rather provides a clever way of hiding and unhiding groups of windows,
and thus gives you the illusion of workspaces.
For a similar,
but more simplistic,
way to do this see <a href="#script_workspace">the rio scripting</a> section below.
</p>

<p>
It is actually quite easy though to manually create pseudo-workspaces in <b>rio</b>:
Just create a new terminal window and run
<b>plumber ; rio</b>.
This will run a <b>rio</b> desktop in this window
(<b>plumber</b> is not required here,
but it will make sure that files automatically opened will only be opened in this isolated <b>rio</b> and not outside of it).
You can maximize this "virtual workspace" and do your work,
hide this window when you want to go back to your first workspace,
then switch back to it by selecting the hidden window in the <b>rio</b> menu.
You can have as many of these workspaces as you like,
and you can run <b>rio</b> inside <b>rio</b> inside <b>rio</b> ad infinitum...
To organize this mess a bit you can also manually label your workspaces.
Lets say you have 4 <b>rio</b> workspaces hidden in the background,
the <b>rio</b> menu will just list them as: <b>rio</b>, <b>rio</b>, <b>rio</b>, <b>rio</b>.
That's not very helpful.
By running <b>grep rio /mnt/wsys/wsys/*/label</b> you will see what window id these workspaces have.
You can then redefine their label,
eg. <b>echo -n workspace1 &gt; /mnt/wsys/wsys/3/label</b>.
The <b>rio</b> menu will now list this window as <b>"workspace1"</b>,
instead of <b>"rio"</b>.
</p>

<p>
Another simple workspace solution is <b>drawterm</b>.
Once a Plan 9 CPU server
(see section
<a href="https://fqa.9front.org/fqa7.html#7.5">7.5</a>
and
<a href="https://fqa.9front.org/fqa7.html#7.6">7.6</a>
in the 9front fqa,
and the
<a href="#sysadmin_cpu">Quick CPU+AUTH+Qemu+Drawterm HOWTO</a>
section below)
has been configured,
you can connect as many <b>drawterm</b> clients to it as you wish
(all of the workspace related tricks mentioned above will also work from within <b>drawterm</b>).
</p>

<h3 id="tiling">Tiling Windows</h3>

<a href="images/tiling.png"><img src="thumbs/tiling.png" title="Propper initial window placement and acme can reduce the need for automatic tiling"></a>

<p>
First of all,
<b>acme</b>
<i>is</i> a tiling window manager.
Just maximize the editor and do your stuff.
</p>

<p>
Secondly,
you can use your <b>rio</b> startup file,
<b>$home/bin/rc/riostart</b>,
to automatically set up a desktop that suits your needs.
For example,
if you have a 1366x768 screen,
you can add these instructions to add an <b>acme</b> window to the left half of the screen,
and a terminal window on the right half:
</p>

<p>
<pre>
window 0,0,683,768 acme
window 683,0,1366,768
</pre>
</p>

<p>
Unlike UNIX,
graphical programs executed in a Plan 9 terminal will not launch a new window,
rather,
the terminal will morph into this new program.
In other words,
running the PDF/Image viewer <b>page</b>,
or the web browser <b>mothra</b> in a terminal for instance,
will in no way effect window placement.
So having an initial window placement that works on your desktop,
will significantly reduce the need for automatic window tiling.
But if you need that,
take a look at the
<a href="#script_tiling">rio scripting</a> section below.
</p>

<h3 id="plumbing">Plumbing</h3>

<p>
We have already seen brief mentions of the Plan 9 plumber a few times in this guide,
but lets take a closer look.
The plumber is essentially a simple inter-process messaging system.
It lets you define a set of actions based on text patterns given to it.
For instance,
in the system wide plumber rules in
<b>/sys/lib/plumb/basic</b>,
you will find the following section:
</p>

<p>
<pre>
# open urls with web browser
type is text
data matches 'https?://[^ ]+'
plumb to web
plumb client window $browser
</pre>
</p>

<p>
This rule is very simple:
If the message is <b>text</b> (it's always text),
if it matches <b>"http://"</b> or <b>"https://"</b> something or other,
define it as <b>"web"</b> related,
and launch a new program,
<b>"$browser"</b>,
with the given text as arguments.
So in effect,
whenever an URL is sent to the plumber,
it opens it up in your default web browser.
So,
right clicking <b>http://9front.org</b> in <b>acme</b> will open up that web page,
likely in <b>mothra</b>.
You can also run the command <b>plumb http://9front.org</b> in a terminal,
for the same effect.
</p>

<p>
You can define your own rules too.
For example,
I wrote my own simple
<a href="#epub">Epub reader</a>,
and added these lines to
<b>$home/lib/plumbing</b>,
in order to always open Epub files with my custom reader:
</p>

<p>
<pre>
# open epubs with custom script
type is text
data matches '([a-zA-Z0-9_\-./]+).(epub|EPUB)'
arg isfile $0
plumb to image
plumb start window eread $file
</pre>
</p>

<p>
This rule adds a check to see if the given argument is an existing file,
if it is <b>$file</b> is set to this filename,
but the logic is otherwise much the same as the above URL rule.
Just make sure that your custom plumber rules end with the line
<b>include basic</b>,
otherwise you will loose all of the default system plumbing rules.
</p>

<p>
Plumbing rules are not restricted to file suffixes.
Suppose you are reading through several documents at the moment,
and you want to bookmark these to keep track of your reading progress.
The solution is simple,
write a database,
lets call it <b>$home/lib/bookmarks</b>,
with content similar to this:
</p>

<p>
<pre>
# work stuff
/usr/glenda/doc/papers/lengthy.pdf!123
/usr/glenda/doc/papers/plain.txt:206

# plan 9 stuff
/sys/doc/9.ps!3
/sys/doc/acme/acme.ps
acme(1)

# fun stuff
/usr/glenda/doc/books/peter_pan.txt:/Chapter 2/
/usr/glenda/music/podcasts/bsdnow/acdecc6a-f7b7-4d64-b64d-f7be713b78e2.mp3
</pre>
</p>

<p>
Right clicking on any of these lines in <b>acme</b>,
will open up the file with an appropriate program.
<b>page</b> for PDF's and postscript files,
<b>play</b> for audio files,
and plain text files directly in <b>acme</b>.
But the default plumbing rules allow you to be even more specific then that.
Piping something like <b>lengthy.pdf!123</b>,
will not only open the PDF in <b>page</b>,
but also on page 123.
Plain text files can also be addressed,
such as <b>plain.txt:206</b> for line 206 of that file,
or <b>peter_pan.txt:/Chapter 2/</b> to open up our Peter Pan book and look for the text string "Chapter 2".
Usually such textual plumbing rules are used when programming,
to open a source file on the offending line by right clicking a diagnostic message for instance,
but we can also used them to keep track of ourselves.
</p>

<p>
Speaking of which,
lets look at one more example of how we can modify plumbing rules to suit our workflow.
in the PIM section below,
I mention a script called
<a href="#pim_que"><b>que</b></a>,
which iterates over a list (a queue),
by printing the next line in the file whenever you run <b>que</b> on it.
Lets assume we have a list called <b>$home/lib/que/peterpan</b> with the following content:
</p>

<p>
<pre>
/usr/glenda/doc/books/peter_pan.txt:/Chapter 1/
/usr/glenda/doc/books/peter_pan.txt:/Chapter 2/
/usr/glenda/doc/books/peter_pan.txt:/Chapter 3/
...
</pre>
</p>

<p>
Now,
each time we run <b>que $home/lib/que/peterpan</b>,
it will tell us what chapter to read next in our book.
And sure enough,
we can right click this output in <b>acme</b> to open up the book in the right place
(since "Chapter x" contains whitespace we need to right click and drag to mark the whole line).
But that is <i>waaay</i> too much work for a lazy pants such as myself!
What I really want is just to add something like this to my bookmark database:
</p>

<p>
<pre>
/usr/glenda/lib/que/peterpan:que
</pre>
</p>

<p>
Right click this in <b>acme</b>,
and have it automatically call <b>que</b> and open up the right chapter for me.
As it turns out,
such automation is easy-peasy,
I just need to add this plumbing rule to my <b>$home/lib/plumbing</b>
(and update my rule set with the command: <b>cp $home/lib/plumbing /mnt/plumb/rules</b>):
</p>

<p>
<pre>
# plumb the next item in a queue file
type is text
data matches '([a-zA-Z0-9_-./]+)(:que)'
arg isfile $1
plumb to none
plumb start rc -c 'plumb `{que '$file'}'
</pre>
</p>

<p>
This rule checks if the plumber received <b>"something_something:que"</b>,
and that the first argument (excluding the <b>:que</b>) was a real file.
We are not interested in opening this file,
so we <b>plumb</b> it to <b>"none"</b>,
and then we run our shell command <b>plumb `{que $file}</b>.
Of course our queue doesn't need to be plain text chapters,
they could be PDF's with page numbers or sequential audio files in a podcast,
or what have you.
</p>

<p>
We can abuse the plumber in all kinds of fun and potentially destructive ways.
It basically allows you to define <i>any</i> text pattern and connect that to <i>any</i> command.
Even if you don't go bananas with this,
it is an eye opening experience to read <b>/sys/lib/plumb/basic</b> and realize just how simple inter-process messaging can be!
</p>

<h2 id="sysadmin">System Administration</h2>

<h3 id="sysadmin_basics">Basic System Administration</h3>

<p>
To shutdown or reboot a Plan 9 system,
you can use the <b>fshalt</b> and <b>reboot</b> commands.
The <b>fshalt</b> command only halts the file system,
but if you have enabled ACPI support,
by adding <b>*acpi=1</b> in <b>plan9.ini</b>
(see section <a href="https://fqa.9front.org/fqa9.html#9.2.3">9.2.3</a> in the fqa),
it will also power off the system on supported hardware
(in either case it is perfectly safe to turn off the machine using the power button once the file system is halted).
</p>

<p>
If you are using a remote Plan 9 desktop,
such as <b>drawterm</b>,
it is safe to just kill the application directly.
The remote desktop is stateless,
and thus shutting it down will in no way effect the host file system.
In fact,
the system is designed to run a CPU server 24/7,
connected to diskless clients where the users do their actual work.
Probably because of this design,
Plan 9 file systems do not try to recover from a power loss,
so don't pull the plug on your file server!
</p>

<p>
I had some CWFS file system corruptions on one of my old laptops.
None of the 9front devs I talked to had seen this problem,
so it might very well be hardware related.
In any event,
the following script solved my issue.
It's unlikely that you'll run into the same problem,
but I leave the script here since it has interesting educational value
(of course,
this is no a substitute for 
<a href="#sysadmin_backups">backups</a>):
</p>

<p>
<pre>
#!/bin/rc
# halt - paranoid shutdown procedures
# usage: halt [-r]
# bugs:  slow shutdown, plausibly overkill

rfork e
fn rtmp{
rm -rf /tmp/*
}
fn dump{
if (test -f /srv/hjfs.cmd) echo dump >> /srv/hjfs.cmd
if not echo dump >> /srv/cwfs.cmd
}
fn sync{
if (test -f /srv/hjfs.cmd) echo sync >> /srv/hjfs.cmd
if not echo sync >> /srv/cwfs.cmd
}
fn fsck{
if (test -f /srv/cwfs.cmd)
echo check ream >> /srv/cwfs.cmd
}
wait=120
if (~ $#* 0) cmd=/bin/fshalt
if not cmd=/bin/reboot

rtmp; dump; sync; sync; sync; fsck; sleep $wait; $cmd
</pre>
</p>

<p>
<i>Update:</i>
A new file system from the 9front developers is in the works,
<a href="https://orib.dev/gefs.pdf">GEFS</a>
(A Good Enough File System),
that tries to recover safely after a crash,
and it has some intriguing long term goals,
such as self-healing.
</p>

<p>
<i>PS:</i> <b>fshalt</b> does not work right in <b>qemu</b> if you use classic Plan 9,
such as 9legacy.
In such cases you should write your own shutdown script,
like so
(note: this is not an issue in 9front):
</p>

<p>
<pre>
#!/bin/rc
# halt - shutdown file server
# usage: halt
echo fsys main sync &gt;&gt;/srv/fscons
sleep 5
echo Its now safe to turn off your computer
echo fsys main halt &gt;&gt;/srv/fscons
</pre>
</p>

<p>
To monitor your remaining battery,
memory usage,
ethernet traffic,
system load and other resources,
you can use the <b>stats</b> and <b>memory</b> commands.
Simply <b>cat</b>'ing around in <b>/dev</b> will also provide much system information,
for instance <b>cat /dev/kmesg</b> is essentially equivalent to <b>dmesg</b> in UNIX.
There is also limited support for suspend and hibernate if you add the <b>apm0=</b> value to <b>plan9.ini</b>
(see section <a href="https://fqa.9front.org/fqa9.html#9.2.3">9.2.3</a> in the fqa and <b>apm</b>(8)).
Don't expect this to work though,
ACPI and APM is a
<a href="https://ftp.openbsd.org/pub/OpenBSD/songs/song45.mp3">hairy business</a>!
</p>

<p>
<i>PS:</i> <b>memory</b> is just a simple shell script in 9front that cat's <b>/dev/swap</b> and reformats the values in more human readable form,
classic Plan 9 systems do not have this script.
</p>

<h3 id="sysadmin_battery">Battery Monitoring</h3>

<p>
Speaking of not working,
battery monitoring usually doesn't in my experience
(to check if it works on your box,
just run <b>stats</b>,
right click and add battery).
And unless you are very lucky,
plugging in a headset will not automatically redirect audio output either.
I had both problems on my cheap Acer laptop
(note to self: only buy ThinkPads from now on).
The last issue will be revisited in the <a href="#audio">audio</a> section below,
as for battery monitoring,
a very simple workaround is to run
<b>sleep 7200; echo Warning: batteries about to go out! &gt; '#c/cons'</b>.
Assuming that your computer has 2 hours (7200 seconds) and 15 minutes of battery capacity,
and you run this command when you know that the machine is fully charged,
you will get notified 15 minutes before your battery runs out.
</p>

<p>
The main problem with this elegant solution,
is that it does not work at all if you expect to reboot your computer at some unknown point in the future.
I find that this is frequently the case when I am traveling,
and need battery monitoring the most.
So I need a way to start a 2 hour countdown that persists across reboots,
this script does the trick:
</p>

<p>
<pre>
#!/bin/rc
# batt - print estimated remaining battery power
# usage: [battery=min] batt [-]
#
# bug: the script doesn't actually know anything about your battery,
#      the user is required to run batt - initially to set a timer.

# set some defaults
rfork e
if(~ $battery "") battery=120 # hardware dependent
capa=$battery
batt=$home/lib/battery
stat='Battery at %p%% estimated remaining time: %r min'
mesg='Your battery is about to run out!'
ping=$home/media/music/samples/mario.mp3

# parse arguments
switch($#*){
case 0
if(! test -f $batt){
echo 'batt: countdown hasn''t started, run batt - first!' &gt;[1=2]
exit notstarted
}
used = `{cat $batt}
pros = `{echo 100 - ($used^00 / $capa) | hoc | sed 's/\..*//'}
remn = `{echo $capa - $used | hoc}
echo $stat | sed -e 's/%p/'$pros'/' -e 's/%r/'$remn'/' -e 's/%%/%/'
case 1
# -, start or continue a countdown
if(! test -f $batt) echo 0 &gt; $batt
while (sleep 60) {
date &gt; $home/lib/end
used = `{echo `{cat $batt} + 1 | hoc}
if (test $used -ge $capa) {
echo $mesg &gt;'#c/cons'
if(test -f $ping) play $ping &gt;[2]/dev/null
rm -f $batt
exit
}
echo $used &gt; $batt
}&amp;
case *
echo 'Usage: [battery=min] batt [-]' &gt;[1=2]
exit usage
}
</pre>
</p>

<p>
You'll note that this simple countdown script measures time in minutes (120, not 7200),
the main reason for this crude measurement of time is battery related,
if we counted every second,
the script would be 60 times harder on our battery.
Anyway,
using this script you can start a countdown when you know the battery is fully charged with the command <b>batt -</b> (or <b>battery=80 batt -</b> or whatever to set a countdown other then the default 120 minutes).
Once that daemon has started,
run <b>batt</b> to get an estimated remaining time of juice.
But here comes the clever part:
After a reboot run <b>batt -</b> to continue a battery countdown!
In fact,
you can fully automate this step by adding something like this to our <b>$home/lib/profile</b>:
</p>

<p>
<pre>
battery=80 # default battery capacity
if (test -d /mnt/term/dev){
# do drawterm stuff
...
}
if not {
# do non-drawterm stuff
if(test -f $home/lib/battery) batt -
...
}
</pre>
</p>

<p>
Don't let the boilerplate here scare you.
If you don't use <b>drawterm</b>,
just add <b>if(test -f $home/lib/battery) batt -</b>,
and you're done
(but you probably don't want to mess with battery stuff if you are using <b>drawterm</b>, for obvious reasons).
This command simply checks to see if the file that the <b>batt</b> daemon uses to measure the battery countdown exists.
Since our <b>batt</b> script removes this file once the countdown has expired,
it knows that an unfinished countdown was in progress before the last reboot,
and so it respawns the daemon.
This is also a convenient place to set your default battery capacity.
Of course,
you could just edit the <b>batt</b> script,
but if you are using this on multiple laptops,
setting such a value in <b>$home/lib/profile</b> might be more practical.
</p>

<p>
Finally,
to know when the laptop is done recharging from a depleted battery,
just measure the time it takes in Ubuntu,
or other suitable grandma distro,
and set an appropriate timer in Plan 9.
We could also wrap this up in a simple script that interrupts a battery countdown and cleans up the temp file:
</p>

<p>
<pre>
#!/bin/rc
# recharge - estimate when battery is recharged
# usage: recharge

slay batt | rc
rm -f $home/lib/battery
sleep 1800
echo 'Battery is fully charged!' &gt; '#c/cons'
</pre>
</p>

<p>
Our script is quite unintelligent,
but in my opinion it is a nice example of how you can create fairly useful and simple workarounds on UNIX-like operating systems,
even when they lack vital features.
</p>

<h3 id="sysadmin_init">Configuring Startup and Shutdown</h3>

<p>
Plan 9 has no <b>/etc</b> directory like UNIX,
instead it is configured through a small handful of files.
The most important of which is probably <b>$home/lib/profile</b>,
the user startup file.
This is where you customize your user specific settings,
it is somewhat analogous to <b>~/.profile</b> in UNIX,
but more important since desktop and shell are much more integrated in Plan 9.
Personally I like to add this line to my <b>lib/profile</b>:
<b>. $home/lib/aliases</b>,
which enables me to add custom aliases to this separate file,
while keeping only system related configurations in <b>lib/profile</b>.
But that is just a matter of taste.
</p>

<p>
Beware that the settings in <b>$home/lib/profile</b> needs to cater to very different situations!
Whether you are booting a CPU server,
a standalone "terminal",
or a diskless one,
or are logging in through a remote connection
(<b>rcpu</b> or <b>drawterm</b> for example),
they all read <b>lib/profile</b>,
but often need different customization's.
The moral is,
be careful when editing your profile,
hubris cause debris.
</p>

<p>
The kernel configuration is in the <b>plan9.ini</b> file,
which resides in a special boot partition.
To read the contents of this partition you must first run <b>9fs 9fat</b>
(for classic Plan 9 run <b>9fat:</b>),
you can then read this file in <b>/n/9fat/plan9.ini</b>
(note: like all Plan 9 commands this manipulates the namespace of your <i>current</i> process,
so you will not see this file in other processes).
It is by editing this file that you configure your system to run as a CPU server or terminal,
you may also need to tweak some hardware specific values here.
See <b>plan9.ini</b>(8) and
<a href="https://fqa.9front.org/fqa3.html">section 3</a>
of the fqa.
</p>

<p>
Network configuration is handled in <b>/lib/ndb/local</b>,
with additional related files in that directory.
But you don't need to mess around with this file if you just want to quickly
<a href="#web">connect to the internet</a> on a laptop
(see
<a href="https://fqa.9front.org/fqa6.html">section 6</a>
in the fqa).
Mail configuration is handled by a number of files in <b>/mail/lib</b>
(see
<a href="https://fqa.9front.org/fqa7.html#7.7">section 7.7</a>
in the fqa).
</p>

<p>
Lastly there is also a desktop specific startup file in <b>$home/bin/rc/riostart</b>,
which is useful for specifying what programs and windows to auto launch,
it is discussed in the <a href="#tiling">tiling windows</a> section of this article.
</p>

<h3 id="sysadmin_rio">Wallpapers, themes, etc...</h3>

<p>
<a href="images/riotheme.png"><img src="thumbs/riotheme.png" title="Rio can be extended and tweaked with 3rd party tools"></a>
</p>

<p>
The <b>rio</b> window manager is painstakingly crafted with love and care to look as boring as humanly possible.
This is important -
a distraction free environment is a productive environment.
But it is possible to install 3rd party patches that let you customize the <b>rio</b> theme and set a wallpaper,
if you really crave such frippery
(this will not work for classic Plan 9 however):
</p>

<h4>Installing theme and wallpaper patch for 9front rio:</h4>

<p>
<pre>
# install rio-themes
; <i>bind -ac /dist/plan9front /</i>
; <i>cd /sys/src/cmd/rio</i>
; <i>hget https://ftrv.se/_/9/patches/rio-themes.patch | patch -p5</i>
; <i>mk install</i>
; <i>reboot</i> # or otherwise restart rio

# write a theme, eg. in $home/lib/theme/rio.theme
# ps: wallpaper must be in the plan 9 image format,
# eg. jpg -9t &lt;pic_1920x1080.jpg &gt;$home/lib/1920x1080.img
rioback /usr/glenda/lib/1920x1080.img
back      f1f1f1
high      cccccc
border    999999
text      000000
htext     000000
title     000000
ltitle    bcbcbc
hold      000099
lhold     005dbb
palehold  4993dd
paletext  6f6f6f
size      000000
menubar   448844
menuback  eaffea
menuhigh  448844
menubord  88cc88
menutext  000000
menuhtext eaffea

# use your theme (add it to riostart if you want)
; <i>window 'cat $home/lib/theme/rio.theme &gt; /mnt/wsys/theme;
sleep 0.5;
grep softscreen /dev/vgactl &gt;&gt; /dev/vgactl;
echo hwblank off &gt;&gt; /dev/vgactl'</i>
</pre>
</p>

<h3 id="sysadmin_language">Internationalization</h3>

<p>
For better or worse,
computing is an <i>English</i> affair.
I'm sorry,
but if you want to program and use any operating system in any professional capacity,
you need to learn the English language.
Nearly all vital documentation,
and any defining works in programming,
computer science and computing history,
will be written in this language.
I don't mean to be unsympathetic here,
I am not a native English speaker myself,
so I know that this can be a tall order,
but that's just the way it is.
</p>

<p>
Having that said,
technically speaking,
Plan 9 does have very good internationalization support.
Of course,
all of the instructions given during installation,
and all of the available documentation is in English.
But the system itself supports most languages as everything is Unicode throughout.<a title="Although not *every* language is supported, some exotic languages need more then just basic Unicode">*</a>
So as long as you have the necessary fonts installed,
you can read and write any language
(well, languages that aren't written from left to right will require some work).
UTF-8 was in fact <i>invented</i> by the Plan 9 developers!
For example,
to write the Northern Norwegian sentence <i>"Æ e i Å æ å"</i>
(yes, this is a real sentence<a title="Loosely translated: I'm in Å too ('Å' can mean too/also/and/oh/river/a town in Lofoten (context is important in the Norwegian language))" href="">*</a>
),
type <b>Alt+Shift+a+e e i Alt+o+Shift+a Alt+a+e Alt+o+a</b>.
A list of the international characters available with the <b>Alt</b> key combo,
can be found in <b>/lib/keyboard</b>.
So to find out how to write a smiley face in Plan 9,
just type <b>grep ☺ /lib/keyboard</b>
(naturally the ☺ can be copy pasted),
and it will print:
</p>

<p>
<pre>
263A  :)          ☺    smiley face
</pre>
</p>

<p>
That is,
type <b>Alt+:+)</b> to produce the Unicode character <b>0x263A</b>,
aka a smiley face.
You can change the default US qwerty layout with the <b>kbmap</b> command,
right click on the layout you want,
then type <b>q</b> to quit.
To set this change permanently:
</p>

<p>
<pre>
# change dvorak to whatever layout you prefer
# setting layout in 9front:
; <i>9fs 9fat</i>
; <i>echo 'kbmap=dvorak' &gt;&gt; /n/9fat/plan9.ini</i>

# setting layout in classic Plan 9:
; <i>sam $home/lib/profile</i>
# add the following line somewhere near the top
; <i>cp /sys/lib/kbmap/dvorak /dev/kbmap</i>
</pre>
</p>

<h3 id="sysadmin_users">User Management and Security</h3>

<p>
To add a new user called <b>bob</b>,
that is a member of the email (<b>upas</b>) and admin groups
(<b>adm</b> for user administration,
<b>sys</b> for access to system files),
on a system using the hjfs file system type:
</p>

<p>
<pre>
# add user to the file server
; <i>echo newuser bob &gt;&gt; /srv/hjfs.cmd</i>
; <i>echo newuser upas +bob &gt;&gt; /srv/hjfs.cmd</i>
; <i>echo newuser adm +bob &gt;&gt; /srv/hjfs.cmd</i>
; <i>echo newuser sys +bob &gt;&gt; /srv/hjfs.cmd</i>

# add user to the auth server
; <i>auth/keyfs</i>
; <i>auth/changeuser bob</i>
; <i>auth/enable bob</i>
</pre>
</p>

<p>
If you are using the cwfs file system,
use <b>cwfs.cmd</b> instead of <b>hjfs.cmd</b>.
If you are using a classic Plan 9 system,
use <b>fscons</b>,
and the command <b>uname</b> rather then <b>newuser</b>,
but otherwise it's the same.
The very first thing Bob needs to do when he first logs in to the Plan 9 box,
is to type <b>/sys/lib/newuser</b>.
This will create an initial home directory with basic files such as a <b>lib/profile</b> and a <b>tmp</b> directory.
<i>Why doesn't the system do this by default?</i>
Consider it a security feature,
users who aren't able to type <b>/sys/lib/newuser</b>,
have only limited access to the system in order to protect the other users.
Btw, you may wish to add the new user to <b>secstore</b> as well
(see
<a href="https://fqa.9front.org/fqa7.html#7.4.3.1">section 7.4.3.1</a>
in the fqa).
</p>

<p>
Security in Plan 9 is built around an astute observation;
While it's the operating systems responsibility to secure the digital world
(ei. the network),
it is your responsibility,
as a physical being,
to provide physical security.
Like me,
being a scrawny nerd,
you may find that statement disconcerting.
Relax,
don't get buffed,
get smart:
For example,
if a Plan 9 network of multiple diskless terminals,
is serviced by a single file server,
that isn't also a CPU server;
The <i>only</i> practical way to compromise file security on that network,
is to gain physical access to the file server machine.
The sysadmin can lock this machine behind a server room door,
behind a death-ray enhanced mutant shark pool,
or whatever physical restraints his evil boss may fancy.
</p>

<p>
The user who boot's a machine has physical access to it.
This <b>hostowner</b> owns all the resources of that machine,
but how much power that gives him on the network depends entirely on how the network is configured.
A Plan 9 machine that isn't a CPU server,
cannot be logged into remotely,
a machine that isn't a file server,
cannot export its files,
and a machine that isn't an auth server,
cannot authenticate remote transactions.
In practice though,
a 9front user will typically set up his laptop as a self contained
<a href="#sysadmin_cpu">CPU+AUTH+File server</a>,
in which case the <b>hostowner</b> is nearly as powerful as the Almighty <b>root</b> in UNIX.
(although he must still show ostensible respect for file permissions)
Single-user "terminals" on the other hand,
where originally diskless,
and do not export any resources whatsoever.
Thus they have nothing to secure and Plan 9 will let anyone login to such a machine without a password.
This is not ideal today,
when a default Plan 9 installation provides a "terminal" with local disk storage.
There are a few ways to work around this issue:
1) Configure the system to run as a CPU+AUTH server,
which does require a password to login.
2) Configure the BIOS to set up a boot password.
3) 9front allows you to encrypt the harddisk,
requiring a passphrase to log in
(see
<a href="https://fqa.9front.org/fqa4.html#4.4">section 4.4</a>
in the fqa).
</p>

<p>
To demonstrate some multiuser shenanigans:
</p>

<p>
<pre>
# UNIX friendly aliases
fn su{
rcpu -u $*
}
fn chown{
chgrp -u $*
}

; <i>su bob</i>   # switch user on CPU server
...
ERROR ERROR ERROR # Oops, bobs profile is missconfigured
... 
; <i>echo allow &gt;&gt; /srv/hjfs.cmd</i> # fs hostowner: allow chown
; <i>chown glenda /usr/bob/lib/profile</i>
; <i>B /usr/bob/lib/profile</i>      # fix the problem
; <i>chown bob /usr/bob/lib/profile</i>
; <i>su bob</i>
</pre>
</p>

<h3 id="sysadmin_disks">Disk Management</h3>

<p>
There is no <b>df</b> command in Plan 9 for measuring disk usage,
but you can get that information in other ways.
On an hjfs file system run this command:
<b>echo df &gt;&gt; /srv/hjfs.cmd</b>.
On cwfs the method is a bit awkward:
<b>echo statw &gt;&gt; /srv/cwfs.cmd &amp;&amp; cat /srv/cwfs.cmd</b>,
will give you a bunch of statistics,
currently using 16 Kb file system blocks
(hit <b>Del</b> when you are done)
What you probably want is the last digit in the <b>wmax</b> line,
which will tell you how much percentage of the disk you are using
(the <b>cache</b> line here is also important,
the cache is only 1/5 the size of the main storage area,
but if it runs out of space - you will run into problems!).
Here is a crude <b>df</b> script for 9front that you may find useful:
</p>

<p>
<pre>
#!/bin/rc
# df - print disk usage on hjfs/cwfs
# usage: df

if (test -f /srv/hjfs.cmd) {
echo df &gt;&gt; /srv/hjfs.cmd
}
if not {
echo statw &gt;&gt; /srv/cwfs.cmd
dd -if /srv/cwfs.cmd -bs 1024 -count 21 -quiet 1 |
grep wmax | sed 's/.*\+//'
}
</pre>
</p>

<p>
I think the method is similar to this in classic Plan 9,
but I am not exactly sure how to do this
(feel free to drop me a line if you know how,
or detect any other faults in my article).
For individual files and folders you can of course use the trusty old <b>du</b> command to measure their size.
Here is a simple and handy script that lists the files and folders in your current directory sorted by disk usage:
</p>

<p>
<pre>
#!/bin/rc
# dus - disk usage summary for current dir
# usage: dus

du -s * | sort -nrk 1 | awk '{
if ($1 &gt; 1073741824) printf("%7.2f %s\t%s\n", $1/1073741824, "Tb", $2)
else if ($1 &gt; 1048576) printf("%7.2f %s\t%s\n", $1/1048576, "Gb", $2)
else if ($1 &gt; 1024) printf("%7.2f %s\t%s\n", $1/1024, "Mb", $2)
else printf("%7.2f %s\t%s\n", $1, "Kb", $2)
}'
</pre>
</p>

<p>
For an example of how to format a USB stick with FAT32 (ei. a DOS partition) and use it in Plan 9,
see the section about <a href="#memstick">USB sticks</a> below.
The process for creating a Plan 9 partition,
instead of FAT32,
is fairly similar.
Assuming the usb stick is called <b>sdUc59fd</b>,
here is how to format it with an hjfs file system:
</p>

<p>
<pre>
; <i>disk/fdisk -baw /dev/sdUc59fd/data</i>
; <i>disk/prep -bw -a fs /dev/sdUc59fd/plan9</i>
; <i>hjfs -r -f /dev/sdUc59fd/fs</i>
; <i>hjfs -n hjfsusb -f /dev/sdUc59fd/fs</i>
; <i>mount /srv/hjfsusb /n/hjfsusb</i>
; <i>touch /n/hjfsusb/anewfile</i>
</pre>
</p>

<p>
And here is how to do it using cwfs:
</p>

<p>
<pre>
; <i>disk/fdisk -baw /dev/sdUc59fd/data</i>
; <i>disk/prep -bw -a^(fscache fsworm other) /dev/sdUc59fd/plan9</i>
; <i>cwfs64x -n fsusb -f /dev/sdUc59fd/fscache -C -c</i>
config: <i>service cwfs</i>
config: <i>config /dev/sdUc59fd/fscache</i>
config: <i>noauth</i>
config: <i>filsys main c(/dev/sdUc59fd/fscache)(/dev/sdUc59fd/fsworm)</i>
config: <i>filsys dump o</i>
config: <i>filsys other (/dev/sdUc59fd/other)</i>
config: <i>ream other</i>
config: <i>ream main</i>
config: <i>end</i>
; <i>mount /srv/fsusb /n/fsusb</i>
; <i>touch /n/fsusb/anewfile</i>
</pre>
</p>

<p>
The above example is for 9front,
as for classic Plan 9 systems,
here is how you create a kfs file system:
</p>

<p>
<pre>
; <i>disk/fdisk -baw /dev/sdUc59fd/data</i>
; <i>disk/prep -a fs /dev/sdUc59fd/plan9</i>
; <i>disk/kfs -f /dev/sdUc59fd/fs</i>
; <i>mount -c /srv/kfs /n/kfs</i>
; <i>touch /n/kfs/anewfile</i>
</pre>
</p>

<h3 id="sysadmin_backups">Backups</h3>

<p>
Plan 9 file systems all have
<a href="#version_control">snapshot capabilities</a>,
so as long as the file system itself is in working order,
you can restore damaged or lost data without much hassle.
Of course,
there is a big <i>if</i> here:
The file system can get damaged,
and the machine it runs on can get damaged,
and the building it lies in can get damaged,
and the country it lies in can get damaged,
and the world it lies in... you get the picture.
So even if you have a super sophisticated ultra safe file system with all the trimmings,
it is <i>not</i> safe!
You should backup your data to an offsite location,
preferably two offsite locations:
If an intruder compromises the data at one site,
having two backups lets you verify which data is accurate and which is corrupt.
</p>

<p>
The trick to migrating from the concept of backups to the practice of it,
is two fold.
First,
backups must be takes <i>automatically</i>.
Doing backups manually ensures that they don't get done.
Secondly,
only backing up <i>essential</i> files will dramatically increase cost effectiveness.
If you are an organized individual,
just write a <b>proto</b>(2) file for your important files,
and schedule a regular <b>mkfs</b>(8) job with <b>cron</b>(8).
I however,
am not an organized individual.
My first problem is that I boot my laptop only semi-regularly,
so I need some easy way to schedule a job "at least" once a day/week/month;
If a weekly job hasn't been run for a week or more when I boot my box,
it needs to run again.
Here is a simple script that accomplishes this:
</p>

<p>
<pre>
#!/bin/rc
# schedule - run commands at scheduled intervals
# usage: schedule
# depend: window schedule in $home/bin/rc/riostart
#
# format: add commands to run in one of the following
# files in $home/lib; daily, weakly, monthly.

# set some defaults
rfork e
lock=$home/lib/lock
mkdir -p $lock
date=`{date}
datesec=`{date -n}
weekrun=Mon
daily=$home/lib/daily
weekly=$home/lib/weekly
monthly=$home/lib/monthly

# check monthly scripts
if(test -f $monthly){
lockfile=monthly_$date(2)^_$date(6)
if(! test -f $lock/$lockfile){
rm -f $lock/monthly_*
touch $lock/$lockfile
@{rc $monthly}
}
}

# check weekly scripts
if(test -f $weekly){
lockfile=weekly_$datesec
if(! test -f $lock/weekly_*) touch $lock/$lockfile
oldlockfile=`{ls -p $lock/weekly_*}
olddatesec=`{echo $oldlockfile | sed 's/weekly_//'}
oldweeksec=`{echo $olddatesec + 604800 | bc}
olddaysec=`{echo $olddatesec + 86400 | bc}
# by default run weekly scripts on a certain day,
# but make sure it runs at least once a week.
if(~ $date(1) $weekrun || test $datesec -gt $oldweeksec){
# also make sure it doesnt run twice in a single day
if(test $datesec -gt $olddaysec){
rm -f $lock/weekly_*
touch $lock/$lockfile
@{rc $weekly}
}
}
}

# check daily scripts
if(test -f $daily){
lockfile=daily_`{date -i}
if(! test -f $lock/$lockfile){
rm -f $lock/daily_*
touch $lock/$lockfile
@{rc $daily}
}
}

# respawn shell
rc
</pre>
</p>

<p>
The script works by writing "lock" files with dates attached whenever a scheduled job is executed.
If these dates are older then a day/week/month
(feel free to expand the script to include quarterly/semily/yearly run jobs if you wish),
the job is executed again and the lock files are updated.
Exactly how you want to run <b>schedule</b> depends on your needs and tastes,
but one suggestion is to add <b>window schedule</b> to <b>$home/bin/rc/riostart</b>.
</p>

<p>
Now,
to tackle my second problem:
Just as time management in my life is disorderly,
so are my files.
I know I have important stuff lying around somewhere that I need to backup,
but it's too much hassle finding out where.
Doing a full backup however is vastly inefficient,
since my home directory contains some
<a href="../2018/text_obsolete.html">non-textual</a> nonsense.
What I need is some quick way of saying backup everything,
<i>except</i> this and that.
Here is one suggestion:
</p>

<p>
<pre>
#!/bin/rc
# nom - no match, print all files except those given
# usage: nom files...

rfork ne
temp=/tmp/nom-$pid

fn sigexit{ rm -f $temp }
if(~ $* */*){
echo 'nom quitting: can''t handle ''/''s.' &gt;[1=2]
exit slash
}

ls -d $* &gt; $temp
ls | comm -23 - $temp
exit    # force file cleanup

#!/bin/rc
# backup - backup important files to offsite storage
# usage: backup

rfork ne

# backup semi-important files
mkdir -p /tmp/backup
fn copy{
mkdir -p $2
if (~ `{ls -ld $1} d*){
mkdir $2/$1
dircp $1 $2/$1
}
if not cp $1 $2
}
fn sigexit{ rm -rf /tmp/backup }
cd $home
for(file in `{nom bin doc games jw media pkg site tmp})
copy $file /tmp/backup
cd $home/bin
for(file in `{nom 386 amd64})
copy $file /tmp/backup/bin
cd $home/doc
for(file in `{nom books health os papers})
copy $file /tmp/backup/doc

backup=9front-^`{date -i}^.tar.gz
tar czf /tmp/backup /tmp/$backup
cd /tmp
# PS: The first whitespace in sed here is a tab
md5sum $backup | sed 's/    / /' &gt;&gt; CHECKSUM

# copy backup to offsite locations
fn sshcopy{
sshfs $1
if(! test -d /n/ssh/backup) {
echo Error: SSH failed!
exit ssh
}
cp /tmp/$backup /n/ssh/backup
cat /tmp/CHECKSUM &gt;&gt; /n/ssh/backup/CHECKSUM
}
sshcopy bkpserv1
sshcopy bkpserv2

rm -rf /tmp/$backup /tmp/backup
</pre>
</p>

<p>
Now the script here is very much tailored to my own idiosyncratic needs,
so <i>don't</i> just copy paste it!
For example,
I omit some big directories in <b>$home</b>,
such as <b>media</b>,
where I pub all of that non-textual mess,
and <b>site</b>
where I keep my web site.
I do copy <b>bin</b> and <b>doc</b>,
but only parts of them.
Clearly,
such details,
will not be relevant for your setup.
But I hope the example might inspire you to write a useful backup utility yourself.
With these tools in place,
I can just add <b>backup</b> to <b>$home/lib/weekly</b>,
and a weekly ~10 Mb backup of my ~10 Gb<a title="Apparently 99.9%+ of my disk usage is junk that I don't want to safeguard. Sounds about right" href="">*</a>
used diskspace is automatically taken,
if I happen to boot my laptop at least once a week.
Of course,
it's still useful to have a full <b>tar czf $home /n/ssh/backup/9front-full.tgz</b> backup lying around,
but running that command manually once or twice a year suffice for my needs.
</p>

<p>
PS: If you happen to be a ZFS user,
you may be yawning right about now.
ZFS does indeed have many fancy features that the Plan 9 file system lacks,
but in my humble opinion,
the practicality of these features are overrated.
For good data security you need two offsite backups even with ZFS,
and with such a setup,
additional data integrity and redundancy is somewhat overkill.
Data compression,
not to mention deduplication,
is even less relevant.
With Terabyte harddisks on commodity hardware nowadays we have infinite disk space,
infinite +50% extra is still infinite.
Besides,
if space were really such a premium,
redundancy would be evil.
In any event,
if you want self healing and all that jazz in Plan 9,
just backup your files to a UNIX machine using ZFS
(or better jet,
run Plan 9
<a href="#sysadmin_cpu">virtually</a>
from a UNIX machine using ZFS).
</p>

<p>
<pre>
ZFS primer for non-ZFS systems:

snapshots:     <a href="#version_control">yesterday</a>
integrity:     md5sum myfiles.tar.gz &gt;&gt; CHECKSUM
redundancy:    <a href="#ssh">cp myfiles.tar.gz /n/ssh/backup</a>
compression:   gzip myfile
encryption:    <a href="#pim_pass">auth/secstore -p myfile</a>
replication:   tar xzf myfiles.tar.gz
deduplication: &lt;buy a disk man&gt;
self healing:  tel mysysadmin
</pre>
</p>

<h3 id="sysadmin_software">Package Management</h3>

<p>
Plan 9 does not really have package management facilities in the sense that a UNIX user would expect.
The system is intended to be "fully-featured" (albeit minimalistic) and few 3rd party software exists,
those that do tend to be distributed as plain source code requiring the user to compile them manually.
It has been toyed with some package management solutions for Plan 9,
but for the most part Plan 9 users usually just compile what they need by hand.
Here are a few examples to demonstrate what "package management" may entail in Plan 9:
</p>

<p>
PS: When compiling software in a Plan 9 terminal,
remember to middle click the window and select <b>scroll</b>.
Otherwise the compilation will freeze once the output has reached the bottom of the window
(this is a "feature", not a bug).
</p>

<h4>Updating the 9front system - elaborately:</h4>

<p>
<pre>
; <i>sysupdate</i>    # download latest sources
; <i>cd /</i>         # rebuild system
; <i>. /sys/lib/rootstub</i>
; <i>cd /sys/src</i>
; <i>mk nuke</i>      # sometimes needed after library changes
; <i>mk install</i>
; <i>mk clean</i>
; <i>cd /sys/man</i>  # optionally rebuild documentation
; <i>mk</i>
; <i>cd /sys/doc</i>
; <i>mk</i>
; <i>mk html</i>
; <i>cd /sys/src/9/pc</i>
; <i>mk install</i>   # optionally rebuild (32-bit) kernel
; <i>9fs 9fat</i>
; <i>rm /n/9fat/9bootfat</i>
; <i>cp /386/9bootfat /n/9fat</i>
; <i>chmod +al /n/9fat/9bootfat</i>
; <i>cp /386/9pc /n/9fat</i>
; <i>sleep 10; reboot</i> # if you have installed a new kernel
</pre>
</p>

<p>
Of course,
you do not need to reinstall the kernel and rebuild the docs for every minor update,
usually all you need to do is:
</p>

<h4>Updating the 9front system - quickly:</h4>

<p>
<pre>
; <i>sysupdate</i>
; <i>cd /sys/src</i>
; <i>mk install</i>
</pre>
</p>

<h4>Install xscreensaver package from the 9front extras:</h4>

<p>
<pre>
; <i>cd /tmp</i>
; <i>9fs 9front</i>     # download package
; <i>tar xzf /n/extra/src/xscr.tgz</i>
; <i>cd xscr</i>        # compile programs and install them
; <i>mk</i>
; <i>for(f in 8.*){ mv $f $home/bin/$cputype/^`{echo $f | sed 's/8.//'} }</i>
</pre>
</p>

<h4>Install vim 7.1 port (old stuff):</h4>

<p>
<pre>
; <i>cd /tmp</i>
; <i>hget http://vmsplice.net/vim71src.tgz | gunzip -c | tar x</i>
; <i>cd vim71/src</i>
; <i>mk -f Make_plan9.mk install</i>
</pre>
</p>

<h4>Install the Bell-Labs port of perl (old stuff):</h4>

<p>
<pre>
; <i>9fs sources</i>    # download iso and mount it
; <i>bunzip2 &lt; /n/sources/extra/perl.iso.bz2 &gt; /tmp/perl.iso</i>
; <i>mount &lt;{9660srv -s &gt;[0=1]} /n/iso /tmp/perl.iso</i>
; <i>cp /n/iso/386/bin/perl $home/bin/386</i>    # install the binary
</pre>
</p>

<h4>Install lua from git.sr.ht:</h4>

<p>
<pre>
; <i>cd /tmp</i>
; <i>git/clone https://git.sr.ht/~kvik/lu9</i>
; <i>cd lu9</i>
; <i>mk pull</i>
; <i>mk install</i>
; <i>lu9 script.lua</i> # or interactively: lu9 -i
</pre>
</p>

<h4>Install <a href="https://www.t3x.org/s9book/index.html">Scheme from Empty Space</a>:</h4>

<p>
<pre>
; <i>cd /tmp</i>
; <i>git/clone https://github.com/bakul/s9fes</i>
; <i>cd s9fes</i>
; <i>mk</i>
; <i>mk inst</i>
; <i>s9</i>            # do some scheming
</pre>
</p>

<h4>Recompile 9front to amd64 and install golang:</h4>

<p>
<pre>
# go will only work on amd64 architecture, so if you are
# running 386, rebuilt to 64-bit first:
; <i>cd /</i>
; <i>. /sys/lib/rootstub</i>
; <i>cd /sys/src</i>
; <i>objtype=amd64 mk install</i>
; <i>cd /sys/src/9/pc64</i>        # build and install a 64-bit kernel
; <i>mk install</i>
; <i>9fs 9fat</i>
; <i>rm /n/9fat/9bootfat</i>
; <i>cp /386/9bootfat /n/9fat</i>
; <i>chmod +al /n/9fat/9bootfat</i>
; <i>cp /amd64/9pc64 /n/9fat</i>
; <i>sam /n/9fat/plan9.ini</i>    # make sure bootfile=9pc64 (not 9pc!)
; <i>sleep 10; reboot</i>         # reboot to a 64-bit system, download Go stuff

# now, lets build go, we will bootstrap the latest version
# of go from 9legacy, then use that to build the go source
# (these instructions quickly get outdated):
; <i>mkdir /sys/lib/go</i>
; <i>cd /sys/lib/go</i>
; <i>hget http://www.9legacy.org/download/go/go1.19.2-plan9-amd64-bootstrap.tbz |</i>
;     <i>bunzip2 -c | tar x</i>
; <i>hget https://golang.org/dl/go1.19.2.src.tar.gz |</i>
;     <i>gunzip -c  | tar x</i>
; <i>mv go amd64-1.19.2</i>
; <i>GOROOT_BOOTSTRAP=/sys/lib/go/go-plan9-amd64-bootstrap</i>
; <i>GOROOT=/sys/lib/go/amd64-1.19.2</i>
; <i>cd amd64-1.19.2/src</i>
; <i>make.rc</i>
; <i>bind -b $GOROOT/bin /bin</i>
# get some recent certificates
; <i>hget https://curl.haxx.se/ca/cacert.pem &gt; /sys/lib/tls/ca.pem</i>
; <i>go get golang.org/x/tools/cmd/godoc</i>

# to make the go environment permanent, add these
# instructions to your $home/lib/profile
; <i>GOROOT=/sys/lib/go/amd64-1.19.2</i>
; <i>bind -b $GOROOT/bin /bin</i>
</pre>
</p>

<p>
<i>PS:</i> In classic Plan 9,
you would run <b>replica/pull /dist/replica/network</b> to get the latest sources from Bell Labs,
and <b>9fs sources</b> to get the Bell Labs repository of contributory software listed under <b>/n/sources</b>.
Today however,
these resources are gone.
You can still mount a snapshot of the contrib repository in 9legacy by running the command
<b>srv -nqC tcp!9p.io sources /n/sources</b>,
the official 3rd party software from Bell Labs will be in <b>/n/sources/extra</b>,
while the repository of contributors are in <b>/n/sources/contrib</b>.
You can also manually mount the Bell Labs wiki from 9p.io like so:
<b>srv -m 'net!9p.io!wiki' wiki /mnt/wiki</b>,
you can then access the wiki by running <b>acme /acme/wiki/guide</B>,
and follow the instructions there
(in 9front accessing these resources are done with:
<b>9fs sources</b> and <b>9fs wiki</b>).
Note however that these old resources are in no way maintained,
so they are more of archaeological,
then practical,
interest.
Concerning 9front specific scripts and programs,
many of them may work just fine in 9legacy,
or any other classic Plan 9 system.
Feel free to try it out :)
</p>

<h3 id="sysadmin_filemng">File Management</h3>

<p>
The default "file manager" in Plan 9 is <b>acme</b>.
If you run <b>B path/mydir</b> for instance,
the contents of <b>mydir</b> will be listed in <b>acme</b>.
Right clicking on a directory here will list its contents,
clicking on a text file will open it up for editing,
and clicking a PDF or audio file will open it up in <b>page</b> or <b>play</b>,
and so on.
To do file operations,
just type in the commands and execute them,
eg. type and middle click <b>touch myfile</b> to create <b>myfile</b>.
</p>

<p>
You can of course use the shell to manage your files,
but there are a few differences between UNIX and Plan 9 that might trip you up.
For example,
you don't have <b>rmdir</b>,
just use <b>rm</b> to delete your directories.
Also there is no <b>cp -r</b>,
instead you have <b>dircp</b> that copies directories.
So,
if you need to copy <b>mydir</b> to <b>otherpath</b>,
you need to run <b>mkdir otherpath/mydir; dircp mydir otherpath/mydir</b>.
If you only want to copy the content of <b>mydir</b>,
not the directory itself however,
just run <b>dircp mydir otherpath</b>.
This may seem cumbersome to a UNIX user,
but it does actually have some benefits.
Beyond a simpler implementation,
the approach is unambiguous.
I do not know how many times I have run <b>cp -r mydir otherpath</b> in UNIX,
when I actually meant to run <b>cp -r mydir/* otherpath</b>
(ei. I only wanted to copy the <i>contents</i> of <b>mydir</b>).
In Plan 9 you don't have this problem.
</p>

<p>
Lastly,
if you really want a GUI,
there is a nice 3rd party file manager,
called
<a href="http://shithub.us/phil9/vdir/HEAD/info.html">vdir</a>.
It works much like the <b>acme</b> file manager,
you right click on things to open them up.
</p>

<p>
<a href="images/filemng.png"><img src="thumbs/filemng.png" title="Different approaches to file management"></a>
</p>

<h3 id="sysadmin_unix">Tips for UNIX Sysadmins</h3>

<p>
As the previous section illustrates,
there are some fundamental differences between UNIX utilities,
and Plan 9 equivalents.
A good UNIX to Plan 9 translation of various sysadmin commands are given
<a href="unix_to_plan9.html">here</a>.
You will note that many essential tools that *nix graybeards take for granted,
such as <b>find</b> or <b>top</b>,
are not available in Plan 9.
And naturally,
standard tools may not work as you expect either,
the shell does not <b>Tab</b> auto-complete,
<b>cp</b> does not copy recursively,
<b>ls</b> does not columnize its output,
and so on.
This can be very unsettling for seasoned UNIX veterans,
but don't panic,
the Plan 9 way of doing things will make sense if you give it time.
Incidentally,
<b>walk</b> (or even <b>du -a</b>) can be used as a lightweight alternative to <b>find</b>,
<b>pstree</b>, <b>memory</b> and <b>winwatch</b> should help you monitor your programs,
<b>Ctrl-f</b> auto-completes filenames in the shell,
and as we have seen,
<b>dircp</b> copies recursively and <b>lc</b> lists files in multiple columns.
You can usually reach your goal in Plan 9,
you just have to learn to walk a different path...
</p>

<h3 id="sysadmin_cpu">Quick CPU+AUTH+Qemu+Drawterm HOWTO</h3>

<p>
As mentioned at the onset,
the focus of this guide is on using Plan 9 as a day-to-day desktop,
not installation and configuration.
So I really didn't want to to do this...
but I suppose it's unavoidable.
The problem here is that there are just so many variables when setting up a Plan 9 CPU server.
For example,
do you run the machine on bare metal or virtually,
if virtually what virtual machine do you use,
in what operating system,
if Linux,
what distro...
And we haven't even begun to consider the many different ways you can configure Plan 9 itself!
What I present here then is just a <i>quick</i> howto.
I assume you want to install a Plan 9 CPU server in <b>qemu</b> on a Linux,
or other UNIX,
machine,
and that you go with all of the default options during the installation of Plan 9,
and that you say <b>"y"</b> to all yes or no questions.
I will not explain indepth the steps we take here,
and I gloss over details that are unimportant.
But if you follow the instructions carefully,
you will end up with a <b>drawterm</b> connected to a Plan 9 CPU+AUTH server running in <b>qemu</b>,
well...
at least on my machine ;)
</p>

<p>
To make this work,
we need to use some painfully detailed <b>qemu</b> flags,
so I recommend using the following wrapper script to launch <b>qemu</b>:
</p>

<p>
<pre>
#!/bin/sh
# 9qemu: wrapper script for launching Plan 9 in qemu
# usage: 9qemu disk [args...]

disk=$1 &amp;&amp; shift
if [ $(uname -s) = Linux ]; then
# non-linux systems may not have this
kvm=-enable-kvm
fi
flags="-net nic,model=virtio,macaddr=52:54:00:12:34:56 \
-net user,hostfwd=tcp::17010-:17010,hostfwd=tcp::17019-:17019,\
hostfwd=tcp::17020-:17020,hostfwd=tcp::12567-:567 \
-device virtio-scsi-pci,id=scsi -device scsi-hd,drive=vd0 \
-device sb16 -vga std -drive if=none,id=vd0,file=$disk"

qemu-system-x86_64 $kvm -m 2G $flags $*
</pre>
</p>

<h4 id="sysadmin_cpu_9front">9front</h4>

<p>
You have two file systems to choose from in 9front,
generally,
<b>hjfs</b> is simpler while <b>cwfs</b> is faster.
You <i>have to</i> use <b>hjfs</b> if your disk is smaller then around 30 Gb,
but if you prefer the default <b>cwfs</b> file system,
follow the steps that are commented out:
</p>

<p>
<pre>
# Step 0: install qemu and drawterm (9front edition)
$ <i>sudo apt install qemu</i> # adjust to suit your system
$ <i>firefox https://drawterm.9front.org</i> # download drawterm
$ <i>tar xzf drawterm-*.tar.gz</i>
$ <i>cd drawterm-*</i>
$ <i>CONF=linux386 make</i>    # adjust to suit your system
$ <i>cp drawterm $HOME/bin</i>

# Step 1: install 9front and reboot
$ <i>qemu-img create -f qcow2 9front.img 2G</i>
#qemu-img create -f qcow2 9front.img 30G # cwfs needs a big disk!
$ <i>9qemu 9front.img -cdrom 9front.iso -boot d</i> # use hjfs file system!
$ <i>9qemu 9front.img</i>

# Step 2: configure boot
; <i>9fs 9fat</i>
; <i>sam /n/9fat/plan9.ini</i>
# change bootargs and add this:
bootargs=local!/dev/sd00/fs -m 448 -A -a tcp!*!564
nobootprompt=local!/dev/sd00/fs -m 448 -A -a tcp!*!564
#bootargs=local!/dev/sd00/fscache -a tcp!*!564
# do not set nobootprompt yet for cwfs!
user=glenda
auth=10.0.2.15
cpu=10.0.2.15
authdom=virtual
service=cpu

# Step 3: write nvram and add user
; <i>auth/wrkey</i>
authid: <i>glenda</i>
authdom: <i>virtual</i>
secstore key: <i>******</i>
password: <i>******</i>

; <i>auth/keyfs</i>
; <i>auth/changeuser glenda</i>
password: <i>******</i>
post id: <i>glenda</i>

# Step 4: configure network
; <i>sam /lib/ndb/local</i>
# change last line and add this:
sys=cirno ether=525400123456 authdom=virtual auth=10.0.2.15 ip=10.0.2.15

ipnet=qemu ip=10.0.2.0 ipmask=255.255.255.0
ipgw=10.0.2.2
auth=10.0.2.15
authdom=virtual
fs=10.0.2.15
cpu=10.0.2.15
dns=8.8.8.8

# Step 5: configure startup
; <i>sam $home/lib/profile</i>
# add these lines at the end of the cpu section, before "case con":
if (test -d /mnt/term/dev) {
# cpu call from drawterm
webfs
plumber
rio -i riostart
}

# reboot
; <i>sleep 5; fshalt -r</i>

# Step Z: enable auth services for cwfs, you only need to do this if you
# used the cwfs file system rather then hjfs during installation (ps: you
# may want to set nobootprompt in plan9.ini after this):
#bootargs:local!/dev/sd00/fscache -c
#config: noauth
#config: noauth
#config: end

# Connecting to the server with drawterm:
$ <i>drawterm -a 'tcp!localhost!12567' -s localhost -h localhost -u glenda</i>
</pre>
</p>

<h4 id="sysadmin_cpu_9legacy">9legacy</h4>

<p>
As you will see,
setting up a 9legacy CPU+AUTH server is notably different from 9front.
Classic Plan 9 has also a few issues with <b>qemu</b>,
first of all,
Plan 9 from Bell Labs does not recognize the harddisk with this setup,
although 9legacy does.
The <b>fshalt</b> script in the original Plan 9 system does not work right in <b>qemu</b>,
which is why we make our own <b>halt</b> script in this example.
Finally,
graphics do not work with this setup.
This isn't a huge deal
(unless you hate <b>ed</b>),
since we can connect to the CPU server with a graphical <b>drawterm</b> once things have been configured.
However,
if you just want to quickly install 9legacy and play around in the desktop without <b>drawterm</b>,
run something like this instead:
<b>qemu-systex-x86_64 -m 2G -hda 9legacy.img</b>
<i>PS:</i> To avoid a naming conflict with the 9front <b>drawterm</b>,
we call the original version of drawterm "<b>9drawterm</b>".
</p>

<p>
<pre>
# Step 0: install qemu and drawterm (original plan9 edition)
$ <i>sudo apt install qemu</i> # adjust to suit your system
$ <i>firefox https://github.com/9fans/drawterm</i> # download drawterm
$ <i>unzip drawterm-master.zip</i>
$ <i>cd drawterm-master</i>
$ <i>CONF=unix make</i>
$ <i>cp drawterm $HOME/bin/9drawterm</i>

# Step 1: install 9legacy and reboot
$ <i>qemu-img create -f qcow2 9legacy.img 2G</i>
$ <i>9qemu 9legacy.img -cdrom 9legacy.iso -boot d</i>
# PS: choose /dev/sdD0/data as the distribution source, type exit at the
# /% prompt, and choose plan9 as the boot method.
$ <i>9qemu 9legacy.img</i>

# Step 2: Do some initial configurations
; <i>echo uname adm +glenda &gt;&gt;/srv/fscons</i>
; <i>cp /adm/timezone/GMT /adm/timezone/local</i> # adjust to suit your needs
; <i>mv /cfg/example /cfg/gnot</i>
; <i>echo ip/ipconfig &gt;&gt; /cfg/gnot/cpurc</i>
; <i>echo aux/listen -q -t /rc/bin/service.auth -d /rc/bin/service tcp &gt;&gt; /cfg/gnot/cpustart</i>
; <i>mv /rc/bin/service.auth/authsrv.tcp567 /rc/bin/service.auth/tcp567</i>
; <i>echo fsys main create /active/cron/glenda glenda glenda d775 &gt;&gt;/srv/fscons</i>
; <i>echo fsys main create /active/sys/log/cron glenda glenda a664 &gt;&gt;/srv/fscons</i>
; <i>ed /rc/bin/cpurc</i>
<i>g/^# auth/s/# (auth.+)/\1/</i>
<i>w</i>
<i>q</i>

# Step 3: configure network
; <i>ed /lib/ndb/local
$
a

sys=gnot ether=525400123456 authdom=virtual auth=10.0.2.15 ip=10.0.2.15

ipnet=qemu ip=10.0.2.0 ipmask=255.255.255.0
ipgw=10.0.2.2
auth=10.0.2.15
authdom=virtual
fs=10.0.2.15
cpu=10.0.2.15
dns=8.8.8.8
.
w
q</i>
; <i>ed /lib/ndb/auth
$
a
hostid=glenda
uid=!sys uid=!adm uid=*
.
w
q</i>

# Step 4: rebuild kernel
; <i>cd /sys/src/9/pc</i>
; <i>mk 'CONF=pccpuf'</i>
; <i>9fat:</i>
; <i>cp 9pccpuf /n/9fat</i>
; <i>ed /n/9fat/plan9.ini
/9pcf/s/9pcf/9pccpuf/
w
q</i>

# Step 5: Setup nvram and users
; <i>auth/wrkey</i>
authid: <i>glenda</i>
authdom: <i>virtual</i>
password: <i>******</i>

; <i>auth/keyfs</i>
; <i>auth/changeuser glenda</i>
password: <i>******</i>
post id: <i>glenda</i>

# Step 6: Halt system and reboot
# PS: the classic fshalt script doesn't work in qemu
; <i>ed /rc/bin/halt
a
#!/bin/rc
echo fsys main sync &gt;&gt;/srv/fscons
sleep 5
echo Its now safe to turn off your computer
echo fsys main halt &gt;&gt;/srv/fscons
.
w
q</i>
; <i>chmod +x /rc/bin/halt</i>
; <i>halt</i>
# click Machine -&gt; Reset in qemu when its safe to reboot

# Connecting to the server with (the original) drawterm:
$ <i>9drawterm -a 'tcp!localhost!12567' -s localhost -c localhost -u glenda</i>
</pre>
</p>

<h4 id="sysadmin_cpu_desktop">CPU+Rio desktop</h4>

<p>
By default a CPU server in Plan 9 does not run a graphical desktop,
the original intention was that this machine would service a number of diskless single-user remote desktops ("terminals") on the network.
If you set up your laptop as a self contained CPU+AUTH server however,
you almost certainly want to use it interactively!
To do so,
you can investigate the difference between <b>/bin/termrc</b> and <b>/bin/cpurc</b>,
the scripts that configure the system to run as either a "terminal" or a CPU server.
In 9front for instance,
you can add this to <b>/cfg/&lt;mymachine&gt;/cpustart</b> to enable a graphical desktop on the CPU server &lt;mymachine&gt;:
</p>

<p>
<pre>
aux/realemu
aux/vga -m vesa -l 1600x900x32 # screen dependent
bind -a '#m' /dev
aux/mouse ps2                  # mouse dependent

for(i in v m i f L A)          # add extra devices
bind -a '#'^$i /dev &gt;/dev/null &gt;[2=1]
rc $home/lib/profile           # regular user setup
</pre>
</p>

<p>
For 9legacy the specifics are a little different,
although you use the same method.
It is a mute point however,
since a 9legacy CPU server cannot run graphics in a virtual machine
(in my experiments at least),
and it is unlikely that you'll be able to run such a system on bare metal.
</p>

<h4 id="sysadmin_cpu_terminals">CPU+PXE terminals</h4>

<p>
Personal computing,
and other fads,
aside,
it is possible to run a Plan 9 network with multiple diskless workstations,
as God intended.
With minor tweaks you can follow the above instructions,
and install a CPU+AUTH+File server on real hardware.
Once that is up and running,
you only need a few additional tweaks to pxe boot diskless workstations.
First,
enable the <b>dhcp</b> and <b>tftp</b> daemons on the server,
by adding these lines to <b>/cfg/&lt;mymachine&gt;/cpurc</b>:
</p>

<p>
<pre>
ip/dhcpd
ip/tftpd
</pre>
</p>

<p>
Then configure the network to use these services,
by adding the following lines in the <b>ipnet</b> tuple in <b>/lib/ndb/local</b>:
</p>

<p>
<pre>
...
ipnet=qemu ip=10.0.2.0 ipmask 255.255.255.0
...
<b>dns=10.0.2.15</b>
<b>dnsdomain=qemu</b>
<b>tftp=10.0.2.15</b>

# add a line for each pxe booted client
sys=term1 dom=term1.qemu ether=8c1645bac636 ip=10.0.2.101 bootf=/386/9bootpxe
...
</pre>
</p>

<p>
We use the <b>authdom</b> and <b>dnsdomain</b> "qemu" here,
which is a rather daft name if we intend to a physical installation.
It doesn't actually matter what label we give it though,
as long as it uniquely identifies the auth server.
The <b>ether</b> line here is the MAC address of the diskless workstation we want to pxe boot.
Finally,
we must provide a <b>plan9.ini</b> file for our client in <b>/cfg/pxe/&lt;MAC address&gt;</b>
(/cfg/pxe/8c1645bac636 in our case):
</p>

<p>
<pre>
bootfile=/386/9pc
bootargs=tls
nobootprompt=tls
auth=10.0.2.15
fs=10.0.2.15
mouseport=ps2intellimouse
monitor=vesa
vgasize=1920x1080x32
*acpi=1
user=dan
</pre>
</p>

<p>
If we now reboot our server,
connect an ethernet cable to the client and configure its BIOS to boot via the network,
everything should work fine
(if not,
<a href="https://fqa.9front.org/fqa6.html#6.7">section 6.7</a>
in the 9front fqa,
might provide some help).
Of course we can mangle our network further in infinite ways:
We could run the CPU, AUTH and File server on separate machines,
and we can have more then one CPU/File server.
We could also do all of this virtually,
or a mix of virtual and bare metal configurations,
including using UNIX as an emulated 9P server.
Feel free to experiment!
</p>

<h2 id="automation">Automation</h2>

<p>
What is the fundamental value of a computer?
However controversial it may be to say so,
it is not watching skaters trip over themselves on Youtube,
or emailing cute cat photos to your colleagues.
The fundamental value of a computer is <i>automation</i>.
Just as a tractor allows you to plow a field with much less effort then a shovel would,
so a computer allows you to do your monthly accounting with much less effort then pen and paper.
So the question of how to use a computer efficiently and wisely,
boils down to programming it to do your chores.
Now I know what you are thinking,
but relax.
There is "programming",
and then there is <i>programming</i>,
we are only going to cover the first topic here,
and leave the latter for the professionals ;)
</p>

<h3 id="rc">Shell Scripting</h3>

<p>
Plan 9's shell, <b>rc</b> is heavily inspired by the classic UNIX shell, <b>sh</b> (the Bourne Shell).
Nevertheless it is a complete rewrite and behaves quite differently.
One obvious difference is the syntax.
The original UNIX shell was designed to mimic the syntax of a user-friendly programming language called ALGOL.
In retrospect this was undeniably a mistake.
<b>rc</b> however mimics the C syntax,
which makes a lot more sense,
since this is the programming language used elsewhere in the system.
</p>
<p>
Another big difference is that <b>sh</b> treats everything as a string,
support for arrays were added later.
This means that correct quoting is super important in the UNIX shell,
and arrays are clunky.
The Plan 9 shell on the other hand treats everything as a list,
so arrays are seamless.
Quoting is also simpler since there is only one escape character
(single quotes).
</p>

<p>
You will find several <b>rc</b> scripts in this article that demonstrate it's use,
but here is a short list of <b>sh</b> to <b>rc</b> translations
(like C, curly brackets in <b>rc</b> are somewhat optional):
</p>

<p>
<pre>
UNIX SHELL               PLAN 9 SHELL

mesg="Hello World"       mesg=(Hello World)
echo "$mesg's!"          echo $"mesg'''s!'
echo ${a}string          echo $"a^string
rm *.{mp3,ogg}           rm *.^(mp3 ogg)

echo date: `date`        echo date: `{date}
list=(`ls`)              list=`{ls}
echo 1st: ${list[0]}     echo 1st: $list(1)
echo all: ${list[@]}     echo all: $list
echo num: ${#list[@]}    echo num: $#list

echo 2&gt;/dev/null         echo &gt;[2]/dev/null
echo &gt;/dev/null 2&gt;&amp;1     echo &gt;/dev/null &gt;[2=1]

if [ "$1" = yes ]; then  if(~ $1 yes){
echo hi                  echo hi
else                     } if not {
echo bye &amp;&amp; exit 1       echo bye &amp;&amp; exit bye
fi                       }
echo err: $? pid: $$     echo err: $status pid: $pid

while true; do           while(){
(subproc)                @{subproc}
done                     }

for i in "$@"; do        for(i in $*){
echo ${i%.*}             echo $i | sed 's/\..*//'
echo $(($i + 1))         echo $i + 1 | hoc
let j++                  j=`{echo $j + 1 | hoc}
done                     }

case in "$@"; do         switch($*){
a) echo Abe          case a
;;                    echo Abe
b) echo Bob          case b
;;                    echo Bob
*) echo Who?         case *
;;                    echo Who?
esac                     }

alias l='ls -l'          fn l{ ls -l }
f(){                     fn f{
echo Funky!              echo Funky!
}                        }
</pre>
</p>

<p>
Many short scripts in this article are written as functions,
this is because I usually add them to a custom alias file,
as mentioned in the <a href="#sysadmin_init">configuring startup and shutdown</a>
section.
But you can easily rewrite these functions as standalone shell scripts if you want.
</p>

<h3 id="rio_scripting">Rio Scripting</h3>

<p>
The desktop in Plan 9 is fully scriptable,
and in true UNIX fashion,
you control it by using a file interface.
</p>

<p>
For example,
if you only have one window open,
and run the command <b>ls /dev/wsys/wsys</b>,
you should see something similar to this:
<b>/mnt/wsys/wsys/1/</b>
This tells you that there is only one window currently open,
which has the ID <b>1</b>.
</p>

<p>
Now run the command <b>echo new sam &gt; /mnt/wsys/wctl</b>,
this should open up a new <b>sam</b> window.
If you <b>ls</b> the <b>/mnt/wsys/wsys</b> directory again,
you should see two windows listed.
You can now delete the <b>sam</b> window with the command <b>echo delete &gt; /mnt/wsys/wsys/2/wctl</b>,
assuming that your <b>sam</b> window had the ID <b>2</b>.
To resize the first terminal window,
either run <b>echo resize -r 0 0 1360 1080 &gt; /mnt/wsys/wsys/1/wctl</b>,
or more simply <b>echo resize -r 0 0 1360 1080 &gt; /dev/wctl</b>.
</p>

<p>
<ul>
<li><b>/dev/wctl</b> window control file for the current window</li>
<li><b>/mnt/wsys/wctl</b> window control file for the system</li>
<li><b>/mnt/wsys/wsys/<i>n</i>/wctl</b> window control file for window <i>n</i></li>
</ul>
</p>

<p>
<b>rio</b> also provides other files that you can use to control its interface,
some of these are discussed in the
<a href="#text">manipulating text in the terminal</a>
and
<a href="#screenshot">taking a screenshot</a>
sections.
For all of these files,
the ones in <b>/dev</b> refer to your current window,
use <b>/mnt/wsys/wsys/<i>n</i>/</b> to manipulate another window.
Here is the full list of files that <b>rio</b> provides:
</p>

<p>
<ul>
<li><b>cons</b> the console</li>
<li><b>consctl</b> the console control file</li>
<li><b>kbd</b> raw keyboard events</li>
<li><b>cursor</b> appearance of the mouse cursor</li>
<li><b>label</b> the window label</li>
<li><b>mouse</b> raw mouse input</li>
<li><b>screen</b> image of the screen</li>
<li><b>snarf</b> the snarf buffer, or "clipboard"</li>
<li><b>text</b> copy of the window text</li>
<li><b>wctl</b> the window control file</li>
<li><b>wdir</b> the current working directory</li>
<li><b>winid</b> the window ID number</li>
<li><b>window</b> image of the window</li>
<li><b>wsys</b> a subdirectory containing the other windows in <b>rio</b></li>
</ul>
</p>

<p>
The fact that the window manager can be easily scripted with standard shell tools gives it enormous flexibility.
Just a quick example to wet your appetite:
The following command will print the window ID number for each window on the screen:
<b>for (win in /mnt/wsys/wsys/*) cat $win/winid &gt; $win/cons</b>
(if you only want to print ID's on <i>visible</i> windows use this command:
<b>for (win in /mnt/wsys/wsys/*) if (dd -if $win/wctl -bs 128 -count 1 -quiet 1 | grep -s visible) cat $win/winid &gt; $win/cons</b>)
<i>PS:</i> If you just want to quickly get the window id of some specific application,
say <b>acme</b>,
you can just <b>grep</b> for it:
<b>grep acme /mnt/wsys/wsys/*/label</b>.
</p>

<h4 id="script_scramble">Scrambling and Unscrambling a Rio Screen</h4>

<p>
As mentioned in the <a href="#text">manipulating text</a> section,
text written to the system console will appear directly on your screen.
This can be seriously annoying,
especially if you have buggy hardware,
which can make the kernel spam error messages that clutter upp your screen.
To automatically ignore all such messages,
you can add this line:
<b>window -hide -scroll cat /dev/kprint</b>
to <b>$home/bin/rc/riostart</b>.
You may also find the following script helpful,
it basically redraws the active <b>rio</b> windows,
and thus unscrambles the screen:
</p>

<p>
<pre>
#!/bin/rc
# unscramble - clear up a garbled rio screen
# usage: unscramble

rfork e
screensize=(`{echo $vgasize | sed 's/x/ /g'})
window -r 0 0 $screensize(1) $screensize(2) exit
for (win in /mnt/wsys/wsys/*) {
if(dd -if $win/wctl -bs 128 -count 1 -quiet 1 | grep -s visible){
echo hide &gt; $win/wctl
echo unhide &gt; $win/wctl
}
}
</pre>
</p>

<p>
<a href="images/scramble.png"><img src="thumbs/scramble.png" title="A de facto screen locker"></a>
</p>

<p>
Of course if you happen to be a weird cookie such as myself,
you may actually enjoy scrambling the screen on purpose.
We have already seen examples of how this feature provides an easy notification mechanism,
but you can abuse it in other ways as well.
For instance,
I have a <b>bat</b> script that draws an ASCII bat signal on the system console and plays the batman theme song.
It totally messes up the display,
and is a nice facepalmer if I happen to mistype <b>batt</b>,
mentioned in the <a href="#sysadm_battery">battery monitoring</a> section above.
Here is another,
more "useful" example.
The script draws a fast moving fullscreen <b>stats</b> display,
and then garbles it up at regular intervals.
I find it sufficiently newb repellent to work as a de facto screen locker.
A non-Plan 9 user (aka everyone) who sees such a screen,
will assume that the computer is horribly broken somehow and refuse to touch it with a ten foot pole.
Of course,
once you delete the <b>stats</b> window everything will return back to normal.
</p>

<p>
<pre>
#!/bin/rc
# scramble - garbles up a rio screen
# usage: scramble

rfork e
fn sigexit{ kill stats | rc }
screensize=(`{echo $vgasize | sed 's/x/ /g'})
window -r 0 0 $screensize(1) $screensize(2) stats -T 0.01 -cflmsz &amp;
while(sleep 3){
if (! ps | grep -s stats) exit
dd -if /dev/random -of '#c/cons' -bs 1024 -count 1 -quiet 1
}
</pre>
</p>

<h4 id="script_max">max - Maximizing Windows</h4>

<p>
The following script lets you maximize windows in various ways,
eg. <b>max</b> will make your current window fullscreen,
and <b>max u</b> will restore its previous dimensions.
<b>max r 2</b> will place window with ID 2 on the right half of the screen,
and so on,
enjoy!
</p>

<p>
<pre>
#!/bin/rc
# max - maximize windows
# usage: max [orientation] [winid]
#
# orientation can be: f (fullscreen), l (left), r (right), t (top), b (bottom),
# tl (top-left), tr (top-right), bl (bottom-left), br (bottom-right) or u
# (unmaximize), default is fullscreen.
#
# bugs: if you are maximizing another window, orientation is required
#       unmaximize is only useful right after maximizing a window.

# set some defaults
screensize=(0 0 `{echo $vgasize | awk -Fx '{ print $1, $2 }'})
if(~ $#windowsize 0)
windowsize=`{dd -if /dev/window -bs 1 -count 70 -quiet 1 |
awk '{ print $2, $3, $4, $5}'}
window=/dev/wctl
if(~ $#* 0) echo resize -r $screensize &gt; $window
if(~ $#* 2) window=/mnt/wsys/wsys/$2/wctl
if(test $#* -gt 2){
echo usage: max [orientation] [winid] &gt;[1=2]
exit
}

# maximize window
echo current &gt; $window
switch $1 {
case f
echo resize -r $screensize &gt; $window
case l
echo resize -r $screensize |
awk '{ printf("%s %s %d %d %d %d", $1, $2, $3, $4, $5/2, $6) }' &gt; $window
case r
echo resize -r $screensize |
awk '{ printf("%s %s %d %d %d %d", $1, $2, $5/2, $4, $5, $6) }' &gt; $window
case t
echo resize -r $screensize |
awk '{ printf("%s %s %d %d %d %d", $1, $2, $3, $4, $5, $6/2) }' &gt; $window
case b
echo resize -r $screensize |
awk '{ printf("%s %s %d %d %d %d", $1, $2, $3, $6/2, $5, $6) }' &gt; $window
case tl
echo resize -r $screensize |
awk '{ printf("%s %s %d %d %d %d", $1, $2, $3, $4, $5/2, $6/2) }' &gt; $window
case tr
echo resize -r $screensize |
awk '{ printf("%s %s %d %d %d %d", $1, $2, $5/2, $4, $5, $6/2) }' &gt; $window
case bl
echo resize -r $screensize |
awk '{ printf("%s %s %d %d %d %d", $1, $2, $3, $6/2, $5/2, $6) }' &gt; $window
case br
echo resize -r $screensize |
awk '{ printf("%s %s %d %d %d %d", $1, $2, $5/2, $6/2, $5, $6) }' &gt; $window
case u
echo resize -r $windowsize &gt; $window
windowsize=()
}
</pre>
</p>

<h4 id="script_workspace">ws - Multiple Workspaces</h4>

<p>
This script provides a virtual workspace-like service for <b>rio</b>.
You use it by typing <b>ws n</b>,
where <i>n</i> is an arbitrary workspace number.
The script works by registering which windows belongs to which "workspace",
and then automatically hides or unhides the correct windows as you "switch" between them.
Of course this is only a pseudo-virtual workspace,
all the windows are still available in the <b>rio</b> menu,
and plumbing a file in one "workspace" may open the file in a different "workspace".
I recommend <b>riow</b> mentioned in the
<a href="#workspaces">workspaces</a> section above,
for a better end user experience,
but the following script might provide some useful insights.
</p>

<p>
<pre>
#!/bin/rc
# ws - pseudo virtual workspaces for rio
# usage: ws n
#
# bugs: the ws workspaces are not isolated from each other, if you need
#       that open a fullscreen window in each ws workspace and run
#       plumber followed by rio in it. even then rio is still blissfully
#       unaware of such "workspaces".

# set some defaults
rfork ne
tmp=/tmp/ws
winid=`{cat /dev/winid}

# initialize 1st desktop on first run
if(! test -d $tmp){
mkdir -p $tmp
touch $tmp/1
echo 1 &gt; $tmp/currentws
ls -np /mnt/wsys/wsys &gt; $tmp/`{cat $tmp/currentws}
}

# update window lists
ls -np /mnt/wsys/wsys &gt; $tmp/riowindows
cat $tmp/[0-9]* | sort -n &gt; $tmp/wswindows
comm -23 $tmp/riowindows $tmp/wswindows &gt;&gt; $tmp/`{cat $tmp/currentws}
for(i in `{comm -13 $tmp/riowindows $tmp/wswindows}){
for(w in $tmp/[0-9]*) sed '/^'$i'$/d' $w &gt; $tmp/TMP &amp;&amp; mv $tmp/TMP $w
}
currentws=`{cat $tmp/currentws}
# no args: echo current ws (after updating windows) and exit
if(~ $#* 0){ echo $currentws &amp;&amp; exit }
touch $tmp/$1

# switch desktop
if(~ $1 $currentws){ echo this is workspace $1 &amp;&amp; exit }
for(i in `{cat $tmp/`{cat $tmp/currentws} | sed '/^'$winid'$/d' })
echo hide &gt; /mnt/wsys/wsys/$i^/wctl
echo $1 &gt; $tmp/currentws
for(i in `{cat $tmp/`{cat $tmp/currentws}}) echo unhide &gt; /mnt/wsys/wsys/$i^/wctl
echo hide &gt; /mnt/wsys/wsys/$winid^/wctl
</pre>
</p>

<h4 id="script_tiling">tile - Tiling Window Manager</h4>

<p>
<b>tile</b> will auto arrange your windows in a tiling fashion.
The algorithm is simple,
place one window on the left half of the screen,
then carve up the right half in even slices for the remaining windows.
The script is intentionally basic,
so feel free to expand or adjust it to suit your own needs.
</p>

<p>
<pre>
#!/bin/rc
# tile - tile windows
# usage: tile

# gather some information
rfork e
screensize=(0 0 `{echo $vgasize | awk -Fx '{print $1, $2}'})
windows=`{for (win in /mnt/wsys/wsys/*)
if(dd -if $win/wctl -bs 128 -count 1 -quiet 1|grep -s visible)
echo `{basename $win}
}
fn left{awk  '{printf("%s %s %d %d %d %d",$1,$2,   0,   0,$5/2,$6 )}'}
fn right{awk '{printf("%s %s %d %d %d %d",$1,$2,$5/2,'$b',$5, '$e')}'}

# auto tile windows
if(~ $#windows 1)
echo resize -r $screensize &gt; /mnt/wsys/wsys/$windows/wctl
if not {
echo current &gt; /mnt/wsys/wsys/$windows(1)^/wctl
echo resize -r $screensize | left &gt; /mnt/wsys/wsys/$windows(1)^/wctl
windows=`{echo $windows | sed 's/^[^ ]+ //'}    # shift windows
step=`{ echo $screensize(4) / $#windows | bc }
b=0; e=$step                # begin, end
for(i in $windows){
echo current &gt; /mnt/wsys/wsys/$i/wctl
echo resize -r $screensize | right &gt; /mnt/wsys/wsys/$i/wctl
b=`{ echo $b + $step | bc }
e=`{ echo $e + $step | bc }
}
}
</pre>
</p>

<h3 id="acme_scripting">Acme Scripting</h3>

<p>
In the above section several window manager scripts are demonstrated,
but if you middle click <b>tile</b>,
or any of the other window manager scripts,
in <b>acme</b>,
nothing will happen.
The reason for this is that the <i>namespace</i> of a terminal window,
and <b>acme</b>,
are different.
If you middle click <b>win</b> and look around in <b>/dev</b>
and <b>/mnt</b> you will see that these directories have different contents then the same directories in a regular terminal.
But don't fret,
you can ask <b>acme</b> to run a command using the namespace of the shell that invoked it with the <b>Local</b> command.
So middle clicking <b>Local tile</b> will tile your <b>rio</b> windows just fine
(to "middle click" two words you need to middle click and drag to select the text).
</p>

<p>
Another way to do this,
if you plan on using <b>tile</b> a lot in <b>acme</b>,
is to write a wrapper script for it in <b>/acme/bin/Tile</b>:
</p>


<p>
<pre>
#!/bin/rc
# Tile - wrapper for tile
tile $*
</pre>
</p>

<p>
As you can see,
this is just an ordinary shell script.
The only difference is that while <b>acme</b> binds <b>/acme/bin</b> to <b>/bin</b>,
other programs don't,
and hence files in <b>/acme/bin</b> are specific to this programs namespace.
When we execute <b>tile</b> here it will have a regular shell namespace,
and thus work as expected.
Note that our <b>acme</b> wrapper has a capital <b>T</b>,
to avoid a naming conflict with our <b>rio</b> tiling script.
There are other simple shell scripts you may want to add to <b>/acme/bin</b>.
In our introductory section on <a href="#acme">acme</a>,
we listed several examples of how you can do basic text editing operations using external tools.
It is easy to make <b>acme</b> commands out of these examples,
by adding shell scripts for them in <b>/acme/bin</b>.
For example,
we could write the following <b>t+</b> and <b>t-</b> scripts,
to indent and unindent text:
</p>


<p>
<pre>
#!/bin/rc
# t+ - indent text
sed 's/^/    /'

#!/bin/rc
# t- - unindent text
sed 's/^    //'
</pre>
</p>

<p>
We can now indent lines in <b>acme</b> by highlighting (left click and drag) the text and pipe it to our new script by middle clicking <b>|t+</b>.
Naturally we can just as easily write commands for commenting and uncommenting text,
for adding and removing line numbers to a file,
for upper and lower casing text,
for obfuscating text with <b>rot13</b>,
and so on.
But the real fun stars when you begin to script <b>acme</b> itself!
Like <b>rio</b>,
<b>acme</b> can be fully controlled by writing plain text to a set of control files.
Lets look at a couple of quick examples:
</p>

<h4 id="acme_coffee">Coffee - Chill ASCII Animations</h4>

<p>
Suppose we have created a series of ASCII coffee mugs in a directory,
our first artwork,
<b>$home/lib/animation/coffee/1</b>,
may look something like this:
</p>

<p>
<pre>
(

) (
( ) (
,---------.__
|         |  |
|         |  |
|         |  |
|         |=='
|         |
`---------'

COFFEE  BREAK
</pre>
</p>

<p>
The other files I will leave to your imagination,
but the point is that when they are displayed in rapid order,
an ASCII animation of a steaming coffee mug will be the result.
In a <b>rio</b> terminal,
we could write such an animation script like this:
</p>

<p>
<pre>
#!/bin/rc
# coffee - print ASCII animation of steaming coffee mug
# usage: coffee

play $home/music/samples/coffee.mp3 &gt;[2]/dev/null &amp;
while()
for(i in $home/lib/animation/coffee/*)
&gt; /dev/text &amp;&amp; cat $i &amp;&amp; sleep 1
</pre>
</p>

<p>
But this will not work in an <b>acme</b> <b>win</b> terminal,
since we don't have <b>/dev/text</b> in our namespace.
To clear the text in an <b>acme</b> window we need to write the command <b>Edit ,d</b>
(that is <b>:%d</b> for all you <b>vi</b> users out there),
select this command,
then click it with our middle mouse button.
Can we do this programmatically?
Sure:
</p>

<p>
<pre>
#!/bin/rc
echo -n Edit ,d   &gt; /dev/acme/body
echo -n /Edit ,d/ &gt; /dev/acme/addr
cat /dev/acme/addr | awk '{ print "MX", $1, $2 }' &gt; /dev/acme/event
</pre>
</p>

<p>
This essentially follows the three steps mentioned above.
The last line is the most cryptic.
What we are doing here is reading the address (the marked text),
which returns information like the beginning and end positions,
which we then feed to <b>awk</b>.
We also append the <b>"MX"</b> command to <b>event</b>,
which tells <b>acme</b> that a middle mouse button was "clicked" on this region of text.
</p>

<p>
Lets call this script <b>/acme/bin/clear</b>,
and lets add a script in <b>$home/bin/rc/clear</b> that clears a <b>rio</b> terminal:
</p>

<p>
<pre>
#!/bin/rc
# clear - clear up a rio terminal
# usage: clear (see also /acme/bin/clear)
&gt; /dev/text
</pre>
</p>

<p>
We can now adjust our coffee animation script so that it works in both the <b>rio</b> terminal and in <b>acme's</b> <b>win</b> terminal:
</p>

<p>
<pre>
while()
for(i in $home/lib/animation/coffee/*)
clear &amp;&amp; cat $i &amp;&amp; sleep 1
</pre>
</p>

<h4 id="acme_slides">Slides - Acme Presentation Show</h4>

<p>
Here is another simple example.
Suppose we have a directory of files called <b>1</b>, <b>2</b>, <b>3</b>...
each providing a slide in a textual slide show.
We could open the first slide by right clicking <b>1</b>.
Then manually editing the filename to <b>2</b>,
type <b>Get</b> and middle click it.
Annoyingly we would need to click <b>Get</b> twice,
since <b>acme</b> will warn us that loading this file will change the contents of our window.
Lets automates this:
</p>

<p>
<pre>
#!/bin/rc
# Slide[-+] - go back and forwards in a slide show
# usage: Slide[-+]
#
# bugs:  slides must be named 1, 2, 3...
#        to "install" the script copy it to /acme/bin/Slide^('' - +)
#        (that is to /acme/bin/Slide{,-,+} in UNIX speak)

switch($0){
case *Slide
ls `{pwd}
exit
case *Slide+
page=`{echo `{basename $%} + 1 | hoc}
if(! test -f $page) exit
case *Slide-
page=`{echo `{basename $%} - 1 | hoc}
if(! test -f $page) exit
case *
echo Error: bogus program name!
exit wrongname
}

echo 'name '`{pwd}^/$page'' &gt; /mnt/acme/$winid/ctl
echo clean &gt; /mnt/acme/$winid/ctl
echo get   &gt; /mnt/acme/$winid/ctl
</pre>
</p>

<p>
To install this script copy it to <b>/acme/bin/Slide</b> and make it executable,
then copy this script to <b>Slide+</b> and <b>Slide-</b> in the same location.
If we now open up one of our iterative slides in <b>acme</b>,
we can middle click <b>Slide+</b> to advance to the next slide,
<b>Slide-</b> to go back to the previous one,
or <b>Slide</b> to list our slides.
We can click <b>Slide+</b> or <b>Slide-</b> repeatedly,
the slide show will stop once we reach the end,
or the beginning,
respectively.
</p>

<p>
Our script contains a couple of special variables,
<b>$0</b> refer to the name of the program that is running.
The behavior of our program will change depending on what it's called,
if it's called <b>Slide+</b> it will advance the slide,
if it's called <b>Slide-</b> it will retreat the slide,
and so on.
<b>$%</b> is a variable particular to <b>acme</b>,
it refers to the filename in the tag of the current <b>acme</b> window.
The last three lines are simple enough,
change the filename,
tell <b>acme</b> not to bother us about contents changing,
and finally load the new file.
</p>

<h4 id="acme_chat">Chat - Simple Peer to Peer Chatting</h4>

<p>
Long before the days of modern chat protocols,
such as IRC,
the ancient UNIX systems came with a simple peer-to-peer chat program called <b>write</b>.
This program established a simple connection between two users,
and just wrote whatever the users had written verbatim to a common text window.
The text would be garbled if both users wrote simultaneously,
so it was customary for the user who had initiated the conversation to write first,
and end his input with (o),
for "over".
Then the other user would reply,
and end his input with (o).
And finally,
when the conversation had run its course,
a user would signal that he ended the conversation with (oo),
for "over and out".
</p>

<p>
Surprisingly enough,
you will actually find this 50 year old program on most modern UNIX boxes today,
even though nobody uses it.
Plan 9 however is a modern operating system for the 90's,
and thus do not include this archaic program.
But if you are feeling nostalgic,
it's trivial to implement it:
</p>

<p>
<pre>
; <i>touch /usr/chat &amp;&amp; chmod 666 /usr/chat</i>
; <i>tail -f /usr/chat &amp;</i>
; <i>while(){ read &gt;&gt; /usr/chat }</i>
</pre>
</p>

<p>
An arbitrary number of users can write the same commands,
and join the chat,
remote Plan 9 users too,
they just need to import the chat machines file system,
and they are good to go.
Whatever people write to the file will be printed verbatim to all that are viewing it.
But this solution is awkward.
For one,
the UNIX <b>write</b> command notified the user you wanted to talk to
(in the increasingly unlikely event that he worked on 
<a href="../2018/console_desktop.html">a text console</a>),
ours doesn't.
And there are some other rough edges besides.
Surely we can write a nicer <b>acme</b> client for this?
Let's start off by implementing a simple notification system;
we can do so in various ways,
but here is a quick suggestion:
</p>

<p>
<pre>
; <i>touch $home/lib/notify</i>
; <i>chmod 666 $home/lib/notify</i> # allow everyone read/write access
; <i>B $home/lib/profile</i>
...
# notify daemon (see statusmsg(8))
while(sleep 5){
if(test -s $home/lib/notify)
@{cat $home/lib/notify; sleep 5} | aux/statusmsg
&gt; $home/lib/notify
}&amp;

# usage: notify user message...
fn notify{
recv=$1 &amp;&amp; shift
if(test -w /usr/$recv/lib/notify)
echo $* &gt;&gt; /usr/$recv/lib/notify
}

# miscellaneous oldschool commands
fn write{
echo 'Use Chat in acme you Neanderthal!'
}
fn wall{
for(recv in `{who}) notify $recv $*
}
fn mesg{
if(~ $1 y) chmod 666 $home/lib/notify
if(~ $1 n) chmod 644 $home/lib/notify
}
fn finger{
whois $1
for(file in /usr/$1/lib/^(plan project))
if(test -f $file) cat $file
}

# clean up old chat logs
rm -f $home/lib/chat    
...
; <i>reboot</i>
</pre>
</p>

<p>
With a notification mechanism in place,
we can go ahead and write our <b>acme</b> Chat client
(we use the unintuitive variable <b>$recv</b> for our message receivers since <b>$user</b> is already taken,
it refers to <i>your</i> user).
We'll implement it as two commands,
<b>Chat</b> for connecting to a chat session,
and <b>Reply</b> for taking whatever we have written in the tag line,
and add it to the chat log.
</p>

<p>
<pre>
#!/bin/rc
# Chat - open a new chat window
# usage: Chat user...

# are we host or client?
rfork ne
if(~ $#* 0) exit
if(~ $#* 1 &amp;&amp; test -f /usr/$1/lib/chat) host=$1
if not host=$user
log=/usr/$host/lib/chat
tag=' ['$host'] Reply '
if(~ $host $user){
touch $log &amp;&amp; chmod 666 $log
for(recv in $*) notify $recv $user wants to Chat!
}

# set up chat window
id=`{awk '{ print $1 }' /dev/new/ctl}
for(cmd in nomenu cleartag scratch) echo $cmd &gt; /mnt/acme/$id/ctl
echo -n $"tag &gt; /mnt/acme/$id/tag
tail -f $log &gt; /mnt/acme/$id/body &gt;[2]/dev/null

#!/bin/rc
# Reply - write tag comments to a chat log
# usage: Reply comment...

rfork ne
tag=`{sed 's/.*(\[.+\] Reply).*/\1/' /mnt/acme/$winid/tag}
host=`{echo $tag | sed 's/.*\[(.+)\].*/\1/'}
reply=`{sed 's/.*Reply //' /mnt/acme/$winid/tag}

echo $user: $reply &gt;&gt; /usr/$host/lib/chat
echo cleartag &gt; /mnt/acme/$winid/ctl
echo ' '$"tag' ' &gt; /mnt/acme/$winid/tag
</pre>
</p>

<p>
The <b>Chat</b> program first determines if we are the chat host or not.
(it's the host that maintains the log and invites the guests)
Next,
the program spawns a new <b>acme</b> window by reading <b>/dev/new/ctl</b>,
the first argument returned when reading this file,
is the ID number of our newly spawned window.
Then we write a few commands to the control file of the new window,
specifying that it should have an empty tag line,
and that <b>acme</b> shouldn't warn us if the content of this window changes.
Finally we add the hosts name and the command <b>Reply</b> to the tag line,
and start listening for changes to the chat log,
which will be printed to the body of our new window.
</p>

<p>
To write something in the chat window,
just add your comment after the <b>Reply</b> command in the blue tag line,
and middle click <b>Reply</b> when you're done.
The comment will be added to the chat log,
prefixed with your user name,
and the tag line will be reset.
Our command contains some odd regex,
a <b>(.+)</b> sed argument would be <b>\(..*\)</b> in UNIX.
Plan 9 utilities all use the <b>egrep</b> like <b>regexp</b>(6) library for regular expressions.
Another detail:
<b>' '$"tag' '</b> is ugly,
but necessary to preserve whitespace correctly.
</p>

<p>
Naturally,
our <b>Chat</b> program is amazingly primitive;
it's method of choosing a host and cleaning up old chat logs is sloppy and it lacks many common features.
Compared to UNIX <b>write</b> however,
it's actually quite advanced;
We can chat with an arbitrary amount of people over the secure network protocol 9P,
the users are identified and can write simultaneously without garbling the output,
and we even have a GUI notification mechanism.
It's not a bad starting point,
but feel free to expand the code to suit your own needs :)
</p>

<h4 id="acme_play">Play - An Acme Music Player</h4>

<p>
In the
<a href="#audio">audio</a>
section below,
we list some simple shell functions for pausing,
resuming and skipping songs we are playing.
We can easily write some of these as shell scripts in <b>$home/bin/rc</b>,
and thus also use them in <b>acme</b>
(we could place them in <b>/acme/bin</b>,
if we <i>only</i> want to use them from <b>acme</b>):
</p>

<p>
<pre>
#!/bin/rc
# skip - skip a song that is playing
kill pcmconv | rc

#!/bin/rc
# pause - pause a song that is playing
stop pcmconv | rc

#!/bin/rc
# resume - resume a paused song
start pcmconv | rc

#!/bin/rc
# vol - adjust audio volume
# usage: vol n
echo master $1 $1 &gt; /dev/volume
</pre>
</p>

<p>
But these functions have two limitations,
first they do not show you a visual playlist,
neither do they allow you to move freely back and forth in the playlist,
you can only skip the current song and play the next one on the list.
If you think about it though,
<a href="#pim_que">que</a>
(a script mentioned later)
has the needed functionality for iterating over a playlist,
and <b>acme</b>,
being a text editor after all,
has the needed functionality to visualize and edit such a list.
It turns out that wrapping these things into a cohesive GUI is very easy.
Our <b>acme</b> <b>Play</b> command looks like this:
</p>

<p>
<pre>
#!/bin/rc
# Play - play a list of audio files
# usage: Play
# bugs:  must run command in a window with audio filenames

echo cleartag &gt; /mnt/acme/$winid/ctl
echo ' Play Quit Repeat ' &gt; /mnt/acme/$winid/tag

while(){
if(! test -d /mnt/acme/$winid) exit
if(! grep -s '&lt;--' /mnt/acme/$winid/body &amp;&amp;
grep -s Norepeat /mnt/acme/$winid/tag) exit
song=`{que $%}
echo clean &gt; /mnt/acme/$winid/ctl
echo get   &gt; /mnt/acme/$winid/ctl
play $song &gt; /dev/null &gt;[2=1]
}
</pre>
</p>

<p>
To use this program,
we must first write a playlist of audio files.
We can easily generate one,
by running something like:
<b>du -a $home/music/creedence | awk '/\.mp3/ { print $2 }' | sort &gt; $home/lib/playlist/creedence</b>
Lets assume we have opened a playlist in <b>acme</b> that looks like this:
</p>

<p>
<pre>
/usr/glenda/music/creedence/01_pagan_baby.mp3
/usr/glenda/music/creedence/02_sailors_lament.mp3
/usr/glenda/music/creedence/03_chameleon.mp3
...
</pre>
</p>

<p>
If we type <b>Play</b> now in the tag line and middle click it,
it will add the commands <b>Quit</b> and <b>Repeat</b> to our tag line,
and start playing <b>01_pagan_baby.mp3</b>,
and then update our playlist,
so that it looks like this:
</p>

<p>
<pre>
/usr/glenda/music/creedence/01_pagan_baby.mp3
/usr/glenda/music/creedence/02_sailors_lament.mp3&lt;--
/usr/glenda/music/creedence/03_chameleon.mp3
...
</pre>
</p>

<p>
When <b>01_pagan_baby.mp3</b> is finished playing,
<b>02_sailors_lament.mp3</b> will start playing,
and the <b>"&lt;--"</b> marker will move to <b>03_chameleon.mp3</b>,
the next song to be played,
and so on.
This will continue indefinitely,
repeating the playlist over and over again.
If we middle click <b>Repeat</b> however,
it will change to <b>Norepeat</b> and the playlist will only play once,
then stop.
(this relies on the convention that <b>que</b> removes the "&lt;--" marker once the queue is finished)
Finally,
just <b>Del</b>'ing this window does kind of work,
but the last song will continue to play until it is finished.
To gracefully quit both this window and the audio playing,
middle click <b>Quit</b>.
Here are our support scripts:
</p>

<p>
<pre>
#!/bin/rc
# Quit - quit Play
stop pcmconv  | rc
echo clean &gt; /mnt/acme/$winid/ctl
echo del   &gt; /mnt/acme/$winid/ctl
start pcmconv | rc
kill pcmconv  | rc

#!/bin/rc
# Repeat - toggle norepeat for Play
echo cleartag &gt; /mnt/acme/$winid/ctl
echo ' Play Quit Norepeat ' &gt; /mnt/acme/$winid/tag

#!/bin/rc
# Norepeat - toggle repeat for Play
echo cleartag &gt; /mnt/acme/$winid/ctl
echo ' Play Quit Repeat '   &gt; /mnt/acme/$winid/tag
</pre>
</p>

<p>
What is important to note here,
is that the playlist is just a plain text file.
So we can freely add or remove lines here as we see fit,
we can also freely move the <b>"&lt;--"</b> arrow to whatever line we want.
After we have middle clicked <b>Put</b> to save our changes,
the next song to be played will be whatever line our arrow is at.
So how do we shuffle our playlist?
If we use <b>acme</b> in Linux with Plan9Port,
we can just mark the playlist and middle click <b>|shuf</b>
(or <b>Edit ,|shuf</b> if it's a very long playlist).
But Plan 9 has no <b>shuf</b> command!
No matter,
we'll just make one.
This crude solution should suffice for our needs:
</p>

<p>
<pre>
#!/bin/rc
# shuf - shuffle input lines
# usage: shuf &lt; input &gt; output

ifs='
'
fn sigexit{ rm -f /tmp/shuf-$pid }
for(line in `{cat /fd/0}){
for(i in 1 2 3) awk '{ printf("%d", substr($2,19)) }' /dev/time
echo @@@$line
} &gt;&gt; /tmp/shuf-$pid
sort -n /tmp/shuf-$pid | sed 's/^[0-9]+@@@//' &gt; /fd/1
exit    # force clean up
</pre>
</p>

<p>
In this script we are using the last nanosecond of the computer clock to generate some fairly random numbers.
<b>/fd/0</b>, <b>/fd/1</b> and <b>/fd/2</b> are equivalent to
<b>/dev/stdin</b>, <b>/dev/stdout</b> and <b>/dev/stderr</b> in UNIX
(and <b>ifs</b> equivalent to <b>IFS</b> naturally -
Plan 9 is even more lower case oriented then UNIX).
The last <b>exit</b> here is ugly,
but sometimes necessary to force the <b>sigexit</b> trap to work
(Plan 9 pays homage to unreliable UNIX signals).
</p>

<p>
We have only provided a handful of crude scripts in this section,
but hopefully they illustrate how easy it is to expand <b>acme's</b> capabilities using only a handful of tiny shell scripts.
Who needs a bloated graphical toolkit anyway,
when you've got <b>acme</b>!
</p>

<h3 id="web_scripting">Web Scripting</h3>

<p>
As mentioned elsewhere,
Plan 9 networks are controlled through plain files.
This implementation is unusual,
and you should read <b>/sys/doc/net/net.ps</b> to familiarize yourself with the concept.
As with the notion that a desktop is controlled by writing text strings to files,
this idea may seem bizarre or even amusing at first.
But any smirk you may have quickly fades as the
<a href="#rio_scripting">rio scripting</a>
section describes how to develop advanced desktop features with simple shell scripts
(try enabling virtual workspaces in Windows 7 with CMD!).
I think the same will be true for web scripting.
Here is a fully-fledged <b>telnet</b> implementation just to wet your appetite:
</p>

<p>
<pre>
#!/bin/rc
clonefile=/net/tcp/clone
&lt;[4] $clonefile {
netdir=`{basename -d $clonefile} ^ / ^ `{cat /fd/4}
echo connect $1|$2 &gt;$netdir/ctl || exit 'cannot connect'
cat $netdir/data &amp;
cat &gt;$netdir/data
}
</pre>
</p>

<h4 id="9front_web_scripts">9front Web Scripts</h4>

<p>
9front ships with the IRC client <b>ircrc</b>,
the pastebin command <b>webpaste</b>,
and the <b>hget</b> and <b>hpost</b> commands.
All of these programs are shell scripts.
<b>hget</b> and <b>hpost</b> are somewhat like <b>wget</b> and <b>curl</b> in UNIX,
but they are only a hundred,
and two hundred,
line shell scripts,
respectively
(in contrast <b>wget</b> and <b>curl</b> are 300,000 lines of C each!).
We will not print their source code here,
but they are worth studying if you plan on writing web scripts in Plan 9 yourself.
As for <b>webpaste</b> it is just a few lines long,
and it is a good demonstration of how to transmit data to a web service
(it depends on <b>hpost</b>):
</p>

<p>
<pre>
#!/bin/rc
if(~ $#* 0)
file=/fd/0
if not
file=$1
hpost -u http://okturing.com -p / a_body@file submit:submit fake:fake a_func:add_post url: |
grep -e '\/body\"' |
sed 1q | sed 's/^.*href=\"//g; s/body\".*$/body/g'
</pre>
</p>

<p>
<i>PS:</i> Classic Plan 9 does not include the above mentioned scripts.
</p>

<p>
<b>ircrc</b> is also about two hundred lines of code,
and is well worth studying.
But here is a stripped down version to illustrate what is possible in Plan 9.
It is a very primitive IRC client that only supports a few IRC commands and hardcodes your nick and channel,
but it is a <i>working</i> IRC client nonetheless.
And at under 70 lines of shell that isn't bad at all:
</p>

<p>
<pre>
#!/bin/rc
rfork ne
server=irc.oftc.net
port=6667
realname=myrealname
target='#cat-v'
netdir=()
nick=mynick

fn sighup {
exit 'hang up'
}
fn sigint sigterm {
if (! ~ $#netdir 0)
echo QUIT : Leaving... &gt; $netdir/data
}
fn mshift {
shift
echo $*
}
fn etime {
date | awk '{print $4}' | awk -F ':' '{print "[" $1 ":" $2 "]"}'
}

fn work {
echo USER $user foo bar :$realname &gt; $netdir/data
echo NICK $nick &gt; $netdir/data
echo PRIVMSG 'nickserv :'identify $"pass &gt; $netdir/data
echo JOIN $target &gt; $netdir/data
while (cmd=`{read}) {
s=$status
if(~ $s *eof) {
echo QUIT : Leaving... &gt; $netdir/data
exit
}
msg=()
out=()
switch ($cmd(1)) {
case /j
if (~ $#cmd 2) {
target=$cmd(2)
msg = (JOIN `{mshift $cmd})
}
case /q
msg = `{mshift $cmd}
case /x
echo QUIT : Leaving... &gt; $netdir/data
exit
case /*
echo unknown command
case *
msg = 'PRIVMSG '^$target^' :'^$"cmd
out = '('^$target^')    ⇐    '^$"cmd
}
echo $msg &gt; $netdir/data
echo `{etime}^' '^$out
}
}

userpass=`{auth/userpasswd 'server='^$server^' service=irc user='^$nick &gt;[2]/dev/null}
if(~ $#userpass 2) {
nick=$userpass(1)
pass=$userpass(2)
}

p='/n/ircrc'
bind '#|' $p
echo connecting to tcp!$server!$port...
aux/trampoline tcp!$server!$port <>$p/data1 &gt;[1=0] &amp;
netdir=$p
work
</pre>
</p>

<h2 id="development">Development</h2>

<p>
As this article is about using Plan 9 as a desktop,
we will only mention development in passing.
The programming language used throughout the system is C,
or more specifically a Plan 9 dialect of C.
The system also has its own set of compilers and linkers,
one set for each supported architecture.
Here is a quick example of how to write and compile a C program:
</p>

<p>
<pre>
; <i>ed take.c</i>
?take.c
<i>a</i> # ed: append text
<i>#include &lt;u.h&gt;</i>
<i>#include &lt;libc.h&gt;</i>
&nbsp;
<i>void</i>
<i>main(int, char*[])</i>
<i>{</i>
<i>print("take me to your leader!\n");</i>
<i>exits(nil);</i>
<i>}</i>
<i>.</i> # ed: end text input
<i>w</i> # ed: write file
112
<i>q</i> # ed: quit
; <i>8c take.c</i>
; <i>8l take.8</i>
; <i>8.out</i>
take me to your leader!
</pre>
</p>

<p>
Don't worry about the <b>ed</b> stuff if you're not used to this editor,
there are alternative text editors in Plan 9 that will also be unfamiliar to you.
Plan 9 users will often open a file with the <b>B</b> command,
which will open the file in whatever text editor happens to be open,
or it will launch the default editor if none is running
(usually <b>sam</b>, add <b>editor=acme</b> to <b>$home/lib/plumbing</b>,
if you prefer <b>acme</b> instead).
At a casual glance the C program looks much like a UNIX equivalent,
but a keen observer will notice many startling differences.
Most Plan 9 programs only have two included headers,
the architecture dependent code <b>u.h</b> and the standard library <b>libc.h</b>.
Notice also that it's perfectly legal for a <b>main</b> to return <b>void</b>,
and that <b>exits</b>,
not <b>exit</b>,
returns a string.
</p>

<p>
There are other differences too.
For one we see that a program in the current directory can be executed just by giving its name,
in UNIX this isn't usually tolerated,
or at least frowned upon.
Another difference is that compiler and linker are two separate programs,
and that each architecture has their own set.
This makes it very easy to cross-compile programs.
For instance,
in the above example a 32-bit PC architecture is assumed,
but on a 32-bit PC you can easily compile 64-bit programs using <b>6c</b> and <b>6l</b>,
or you can compile ARM programs using <b>5c</b> and <b>5l</b> (see <b>2c</b>(1)).
Of course you cannot <i>run</i> ARM programs on PC hardware,
but a Raspberry Pi running Plan 9 can easily compile its software on a PC running Plan 9
(or vice versa for that matter).
In fact it's easy-peasy to cross compile a 32-bit Plan 9 system into 64-bits
(see
<a href="https://fqa.9front.org/fqa5.html#5.2.2.1">section 5.2.2.1</a>
in the 9front fqa).
</p>

<p>
Probably,
the best place to start if you want to develop in Plan 9,
is to read the article
<a href="http://doc.cat-v.org/plan_9/programming/c_programming_in_plan_9">C Programming in Plan 9</a>.
The specific details of the Plan 9 C dialect is discussed in <b>/sys/doc/comp.ps</b>.
Other important papers in this directory are
the <b>acid</b> debugger paper <b>acidpaper.ps</b>
and the <b>mk</b> paper <b>mkfiles.ps</b> (equivalent to <b>make</b> in UNIX).
The other papers here will also give you some useful hints,
but be aware that some of them are quite dated.
Lastly,
manpages and source code is very readable in Plan 9,
so use that for what it's worth!
The <b>src</b> command will let you quickly look up source code for any given command,
eg. <b>src echo</b>.
Another resource that I highly recommend is
<i>Introduction to Operating Systems Abstractions Using Plan 9 from Bell Labs</i>,
by Francisco J. Ballesteros.
You can download the PDF for free from the internet,
and despite its tedious name,
it is a marvelous programming introduction to Plan 9.
Naturally many classic UNIX resources are also useful for Plan 9,
even though many details aren't directly applicable,
such as
<i>The C Programming Language</i> by Kernighan and Pike,
<i>The UNIX Programming Environment</i> by the same,
and <i>The AWK Programming Language</i> by Aho, Kernighan and Weinberger.
</p>

<p>
Beyond shell and <b>awk</b> programming,
there are also some support for external programming languages,
such as POSIX C and <b>sh</b> (see <b>/sys/doc/ape.ps</b>),
Perl, Python, Go, Scheme, Lua,
and Limbo if you install <a href="#inferno">Inferno</a>
(see appendix L in 9fronts <a href="https://fqa.9front.org">fqa</a>).
Generally though C and shell programming are by far the best supported languages.
The Perl port is very old for instance,
and Python was recently dropped from 9front
(you can still get it if you want - see instructions below).
</p>

<h3 id="version_control">Version Control</h3>

<p>
All the Plan 9 file systems
(cwfs or hjfs in 9front,
or fossil or kfs in classic Plan 9)
have built-in support for snapshots.
And like all file systems with snapshot capabilities,
only the difference between the versions are saved,
so a snapshot of a 1 Gb file with 10 Kb of difference,
will only consume 10 Kb of disk space.
Snapshots are usually taken at regular intervals automatically,
but you can take one manually if you want,
e.g. <b>echo dump &gt;&gt;/srv/hjfs.cmd</b>
(use <b>/srv/cwfs.cmd</b> if you are using the cwfs file system).
To read snapshots run <b>9fs dump</b>,
the files will be located in <b>/n/dump</b>.
For example if you are looking for the snapshot of <b>/usr/glenda/prj/code.c</b> taken 23 February 2020,
it will be located in <b>/n/dump/2020/0223/usr/glenda/prj/code.c</b>.
The <b>yesterday</b> command is a quick way to print the path of the most recent snapshot.
<b>history</b> will print all available snapshots where the file content differs.
</p>

<p>
The Bell Labs developers used the built in snapshot feature of Plan 9 as their version control system,
and for personal use this works great.
But if you are collaborating online with other programmers,
or if you are working on non-Plan 9 software,
you probably want to use a more traditional version control software.
Previously the 9front developers used mercurial to maintain their project,
but recently a switch was made to <b>git</b>,
using the Plan 9 port called <b>git9</b>
(it is still possible to get mercurial if you really need it):
</p>

<h4>Installing python and hg (mercurial):</h4>

<p>
<pre>
; <i>cd /tmp</i>
; <i>git/clone gits://git.9front.org/plan9front/pyhg</i>
; <i>cd pyhg</i>
; <i>install.rc</i>
</pre>
</p>

<h3 id="namespaces">Files and Namespaces</h3>

<p>
The big difference between UNIX and Plan 9,
which is especially important for developers to note,
is that while "everything" is a file in UNIX,
<i>everything</i> is a file in Plan 9!
There are no sockets or ioctl for instance,
all networks and devices are controlled through plain files.
It is hard to emphasize just how much simpler this makes programming,
but some illustrative examples can be found in the
<a href="#automation">Automation</a>
sections above.
</p>

<p>
In order to make everything in this dynamic and complex world of ours work as files,
Plan 9 uses some conventions and mechanics that are unfamiliar to UNIX users.
For example,
devices often need control interfaces as well as input/output interfaces,
so Plan 9 often implements a device as a directory with multiple files.
For example audio input/output is handled through <b>/dev/audio</b>,
but the control interface is <b>/dev/audioctl</b>,
and hardware statistics are available through <b>/dev/audiostat</b>
(in 9front that is).
Another example is the <b>/bin</b> directory,
which unlike UNIX contains <i>all</i> available programs on the system.
However,
what that means may differ from program to program,
and notably,
from process to process.
Files in <b>/bin</b> can also be directories that group related programs together.
All games are in <b>/bin/games</b> for instance,
so launching <b>/bin/games/sudoku</b> requires you to type <b>games/sudoku</b>.
</p>

<p>
There are no links in Plan 9,
instead files and directories can be bound to different locations using the <b>bind</b> command.
Lets again consider <b>/bin</b>:
Programs are actually sprinkled in different places throughout the system,
32-bit PC binaries are in <b>/386/bin</b>,
32-bit PC binaries for <b>acme</b> are in <b>/acme/bin/386</b>.
Shell scripts are in <b>/rc/bin</b> and personal shell scripts are in <b>$home/bin/rc</b>,
and so on.
When the system boots,
the relevant program directories are all bound to <b>/bin</b>.
If the system is 32-bit PC <b>/386/bin</b> is bound to <b>/bin</b>,
if it's a 64-bit Sparc system <b>/sparc64/bin</b> is bound instead,
and so on.
The important lesson here is that the correct filestructure is <i>assembled</i> during boot.
To see the full details of how your filestructure is assembled,
use the <b>ns</b> command.
This <i>namespace</i> can be manipulated freely,
and it only effects the current process
(and any child processes executed afterwards).
</p>

<p>
For example,
in the previous version control section,
we mentioned how you can look up old versions of some file.
But this might be tedious if we need to check out many files of a given date,
in such a case it would be simpler to manipulate our namespace instead.
Let say we want to check out all of our project files,
as they were on 23 February 2020,
the following commands should do the trick:
</p>

<p>
<pre>
; <i>9fs dump</i>
; <i>bind /n/dump/2020/0223/usr/glenda/prj $home/prj</i>
</pre>
</p>

<p>
Now all of the files in <b>$home/prj</b> in this window refer to our old copies of 2020.
To go back to the current version of this directory,
just run <b>unmount $home/prj</b>
(that's <b>unmount</b>, not <b>umount</b>).
The <b>bind</b> command can be used much like <b>ln</b> in UNIX,
to create shortcuts from one point to another in the file system.
But it is much more powerful.
For one it doesn't care if the files are in the same file partition,
or even on the same physical machine.
For another you don't need to <i>replace</i> directories,
you can merge them.
That is what Plan 9 does with <b>/bin</b>,
many directories are bound together in this location.
Various flags to <b>bind</b> let you specify if the directory should be prepended or appended,
and whether or not to allow file creation in such a union.
</p>

<p>
Another example of such namespace manipulation is the <b>rcpu</b> (<b>cpu</b> in classic Plan 9) command,
which binds a remote CPU servers processor to the current process,
while the local files and devices,
such as the keyboard,
are kept unmodified.
The window still looks and behaves like a normal Plan 9 terminal,
but it's now using the remote machines processor.
This is handy if the remote machine is fast,
while the local machine is slow or over taxed.
It is also useful if you are testing software for a different architecture,
such as running ARM programs from a PC or vice versa.
Other remote resources can be imported as well,
such as an external audio or ethernet device
(and thus create a very simple MPD/VPN service).
Again,
only the process in question is manipulated,
other running processes are unaffected.
Namespaces lies at the very heart of Plan 9's capabilities,
but it's hard for UNIX users to grasp the concept.
If it helps,
think of each Plan 9 process running inside its own mini-jail.
The difference though is that namespaces in Plan 9 were not primarily devised as a way to <i>isolate</i> resources,
but a way to <i>distribute</i> them.
</p>

<p>
If you really want a jail though,
it's simple enough to implement one:
</p>

<p>
<pre>
; <i>rcpu -u loser</i>    # unprivileged user
; <i>plumber</i>          # isolated inter-proc messaging
; <i>auth/factotum -n</i> # isolated auth services
; <i>rio</i>              # isolated desktop
; <i>mkdir fake</i>
; <i>bind fake $home</i>  # sandbox home directory
; <i>rfork n</i>          # isolated namespace
; <i>rfork e</i>          # isolated environment variables
; <i>rfork s</i>          # isolated signals
; <i>rfork f</i>          # isolated file descriptor table
; <i>rfork m</i>          # disallow mounts
</pre>
</p>

<p>
You can mix and match these commands to give your jail more or less powers,
and you can manipulate files in <b>/dev</b> and <b>/net</b> to grant or deny various devices and networks.
Read <b>fork</b>(2) for more details.
</p>

<p>
Sidenote:
You'll notice that many of the scripts in this article start with <b>rfork e</b>.
This ensures that the variables used in the script are not inadvertently blended with the parent.
On rare occasions you may want that.
The <a href="#script_max"><b>max</b></a> script above for example shares its environment variables with the parent.
It sets the <b>windowsize</b> variable to the current window size,
if it isn't already set.
This value is then later used by another invocation of <b>max</b> to restore a window to its former dimensions.
(thus "unmaximizing" it)
Generally though,
it is hazardous to use global variables,
so start your scripts with <b>rfork e</b>.
It is prudent to use <b>rfork ne</b> if you are working with directories that may change,
such as <b>/mnt</b> or <b>/tmp</b>,
and in rare occasions you may want to run plain <b>rfork</b>
(it defaults to <b>rfork nes</b>)
to isolate your script even further.
It ensures a stoic indifference if your parent commits seppuku.
</p>

<h2 id="web">The Web</h2>

<p>
If you have a wired connection to the internet
(you do if your using a virtual machine),
you should already be connected.
If not run the command <b>ip/ipconfig</b>.
If your having trouble,
you can run the <b>netaudit</b> script,
and see if the network status looks like it's supposed to
(classic Plan 9 does not have this script).
</p>

<h3 id="wifi">Wireless Network</h3>

<p>
Wireless networking is only supported in 9front.
During startup you may see a line similar to
<b>#l1: '/lib/firmware/iwn-6005' does not exist</b>
(you can check startup messages later with <b>cat /dev/kmesg</b>).
This tells you that firmware for the wireless device <b>#l1</b> is missing.
</p>

<p>
9front uses firmware from OpenBSD,
so download the correct package for your device from <a href="https://firmware.openbsd.org">firmware.openbsd.org</a>,
and unpack it in <b>/lib</b>
(if you don't have wired internet access,
put this file on a FAT32 formatted <a href="#memstick">USB stick</a>,
using Ubuntu or Windows or whatever and transfer it that way), 
reboot,
and then you can connect to a wireless network.
To illustrate:
</p>

<p>
<pre>
; <i>cd /lib</i>
; <i>hget https://firmware.openbsd.org/firmware/6.6/iwn-firmware-5.11p1.tgz | tar xz</i>
; <i>reboot</i>
; <i>bind -a '#l1' /net</i>    # you might want to add this to $home/lib/profile
; <i>aux/wpa -s mynetwork -p /net/ether1</i>
!Adding key: proto=wpapsk essid=mynetwork
password: ******
!
; <i>ip/ipconfig ether /net/ether1</i>
</pre>
</p>

<p>
You can easily automate the last two steps with a short script,
and thus connect to a wireless network with <b>wifi mynetwork</b>:
</p>

<p>
<pre>
fn wifi{
aux/wpa -s $1 -p /net/ether1
ip/ipconfig ether /net/ether1
}
fn wifiscan{
# scan for available networks
cat /net/ether1/ifstats
}
</pre>
</p>

<p>
Normally,
9front uses the ethernet as its default network card,
but you can overwrite this.
Adding <b>ether0=type=iwl</b> in <b>plan9.ini</b>,
will tell the system to use the wireless network card as <b>/net/ether0</b>.
(see <b>plan9.ini</b>(8) for more details)
</p>

<p>
PS: Make sure the firmware package has the file you need
(<b>iwn-6005</b> in the above example),
it might not be the latest one.
You also need to
<a href="#sysadmin_software">recompile the kernel</a>
if you want the network card enabled at boot time.
</p>

<h3 id="browsing">Browsing The Web</h3>

<a href="images/browsers.png"><img src="thumbs/browsers.png" title="There are Plan 9 Web Browsers, but they are very basic..."></a>

<p>
The preferred web browser for 9front is <b>mothra</b>,
but the classic Plan 9 browser,
<b>abaco</b>,
is also available.
Both of these web browsers have only basic support for HTML,
they do not support any CSS,
let alone JavaScript.
Also,
you must supply a full URL with a protocol prefix,
eg. <b>https://www.wikipedia.org</b>,
not just <b>wikipedia.org</b>.
</p>

<p>
To open a local HTML file in <b>mothra</b>,
write <b>file:///path/to/file</b>.
To download content from a website,
right-click and choose <b>moth mode</b>.
The mouse cursor will change to a moth,
you can now click on any link or image to download it.
Choose <b>moth mode</b> again,
to return to the default mode,
where clicking on a link will follow it instead of downloading it
(<b>abaco</b> cannot open local files or download content).
</p>

<p>
As for <b>acme</b>,
you cannot use it to browse the web interactively,
but you can do a basic text dump of a webpage,
by middle clicking something like
<b>wurl2txt http://9front.org</b>.
</p>

<p>
Recently,
NetSurf has been ported to Plan 9 by the 9front developers.
The browser is slow and glitchy with a ton of bugs,
and thus provides a fairly convincing web 2.0 experience.
It is still a <i>very</i> simple browser though,
so don't expect to do your online shopping in Plan 9 anytime soon
(you can do
<a href="#youtube">youtube</a> - but not in a browser).
</p>

<h3>Install the Plan 9 port of NetSurf from github:</h3>

<p>
<pre>
; <i>cd $home/src</i>
; <i>git/clone https://github.com/netsurf-plan9/nsport</i>
; <i>cd nsport</i>
; <i>fetch clone http</i>
; <i>mk</i>
; <i>mk install</i>
; <i>netsurf</i>   # browse! (make sure <b>webfs</b> is running)
</pre>
</p>

<p>
<img src="thumbs/netsurf.png" title="NetSurf kinda looks like an actual browser">
</p>

<h3 id="download">Downloading</h3>

<p>
In addition to downloading files interactively with a browser,
you can get files with <b>hget</b>
(it works much like <b>wget</b> in UNIX).
There is also <b>ip/torrent</b> for downloading torrents.
</p>

<h3 id="email">Email</h3>

<p>
Hints on setting up an email server can be found in
<a href="https://fqa.9front.org/fqa7.html#7.7">section 7.7</a>
of the 9front fqa.
And see
<a href="https://fqa.9front.org/fqa8.html#8.4.1.1">section 8.4.1.1</a>
for tips on interfacing the Plan 9 mail server with GMail.
Once you have configured this to your suit your needs,
Plan 9 provides a few alternative email clients.
Probably the most useful one will be <b>acme's</b> mail client <b>Mail</b>,
but you can check out <b>faces</b> and <b>nedmail</b> too if you want.
</p>

<h3 id="chat">Chatting</h3>

<p>
To join the fairly active <b>#cat-v</b> channel,
on <b>irc.oftc.net</b>,
where 9front developers hang out,
run this command <b>ircrc -n mynick -t '#cat-v'</b>.
A few alternative 3rd party IRC clients are also available,
and an xmpp (jabber, eg. Google Talk) client.
You can check out the latter projects
<a href="https://sr.ht/~ft/xmpp/">web site</a>
if you are interested.
</p>

<h3>Install and use the ircs persistent IRC client:</h3>

<p>
<pre>
; <i>cd /tmp</i>
; <i>9fs 9front</i>
; <i>tar xzf /n/extra/src/ircs.tgz</i>
; <i>cd ircs</i>
; <i>mk install</i>
; <i>ircs -p dansimon irc.oftc.net</i>    # start IRC client
; <i>ircx -t '#cat-v'</i>    # start IRC user interface
</pre>
</p>

<p>
An IRC log is available in <b>/tmp/ircs/log</b>.
The first time you run this,
<b>factotum</b> will ask you for your password
(the <b>-p</b> in <b>ircs</b> means 'use a password').
You can add this password permanently the the
<a href="#pim_pass"><b>secstore</b> file factotum</a>,
for a fully automatic authentication.
</p>        

<h3 id="werc">Running a Web Server</h3>

<p>
Plan 9 is quite capable of serving web pages,
just as long as you keep things simple.
To quickly set up a local web page,
do the following:
</p>

<p>
Write some static html file(s),
<b>$home/www/mysite/index.html</b>
for instance.
And check that it works
(eg. <b>mothra file://$home/www/mysite/index.html</b>).
Now we can configure the <b>rc-httpd</b> web server,
by adding the following to <b>/bin/rc-httpd/select-handler</b>:
</p>

<p>
<pre>
if(~ $SERVER_NAME mysite){
PATH_INFO=$location
FS_ROOT=/usr/myname/www/$SERVER_NAME
exec static-or-index
}
</pre>
</p>

<p>
Start the web server by running
<b>aux/listen1 tcp!mysite!80 rc-httpd/rc-httpd</b>
(make sure you middle click the terminal window and select <b>"scroll"</b>).
You should now be able to connect to your web server with <b>mothra http://mysite</b>!
</p>

<p>
<i>PS:</i> Substitute <i>mysite</i> for whatever hostname or ip address is appropriate for your machine.
</p>

<h2 id="media">Multimedia</h2>

<p>
9front has quite decent multimedia support as we shall see,
but classic Plan 9 systems are very limited in this respect.
They only support SoundBlaster cards for instance,
and MP3 file formats
(oddly enough you will find the MP3 decoder and encoder under <b>/bin/games</b> in classic Plan 9).
To enable audio in a <b>qemu</b> virtual machine (for both 9legacy and 9front),
run <b>qemu</b> with the <b>-device sb16</b> flag,
and add this line to <b>plan9.ini</b>:
<b>audio0=type=sb16 port=0x220 irq=5 dma=5</b>.
After this,
you still need to bind the audio device in 9legacy,
like so: <b>bind -a '#A' /dev</b>.
You may want to add this command to
<b>$home/lib/profile</b>.
You don't need to mess with your profile in 9front however,
and audio should just work out of the box on real hardware.
</p>

<h3 id="audio">Audio</h3>

<p>
Adjusting the volume,
to say 80%,
can be done like this:
<b>echo 80 &gt; /dev/volume</b>,
or more precisely:
<b>echo master 80 80 &gt; /dev/volume</b>
(80% for left and right speakers).
But switching between headphones and speakers can be a bit tricky.
If you're are lucky the hardware will just take care of it,
but if you aren't you have to manually redirect audio pins.
On one of my machines the command
<b>echo pin 21 &gt; /dev/audioctl</b>
switches audio output to the jack port,
on another the command
<b>echo pin 33,12,2 &gt; /dev/audioctl</b>
does the trick.
It varies from machine to machine,
you can figure out the correct command by analyzing the output of
<b>cat /dev/audiostat</b>.
This can be a bit daunting,
but don't panic,
just look for words such as <b>jack</b>, <b>speaker</b>, <b>out</b>, <b>pin</b>,
and experiment.
Don't worry,
the machine will not blow up if you get it wrong ;)
</p>

<p>
You might find some of the following functions helpful,
keep in mind though that some specifics here are hardware dependent:
</p>

<p>
<pre>
fn volume{ echo master $1 $1 &gt; /dev/volume }
fn headphones{ echo pin 21 &gt; /dev/audioctl ; volume 40 }
fn speaker{ echo pin 20 &gt; /dev/audioctl ; volume 80 }
fn mute{ volume 0 }
fn pause{ stop pcmconv  | rc }
fn resume{ start pcmconv | rc }
fn skip{ kill pcmconv  | rc }
</pre>
</p>

<p>
Of course,
there are more user friendly 3rd party utilities that can help you out.
<a href="https://git.sr.ht/~ft/jacksense">jacksense</a>
tries to automatically switch between output pins,
whenever you plug in a headset.
And <b>volume.c</b> from the 9front extra repository,
gives you a simple button for adjusting output volume.
</p>

<p>
Technically you can play a raw audio file just by running <b>cat file &gt; /dev/audio</b>,
or you can decode it first:
<b>audio/mp3dec &lt; file.mp3 &gt; /dev/audio</b>.
But 9front includes a userfriendly shell script that makes life much easier:
<b>play file.mp3</b>.
</p>

<p>
Doing audio recording is theoretically possible in 9front,
you first need to redirect the correct pins as in the above example to set audio for input instead of output,
and then a read from <b>/dev/audio</b> will record sound.
Eg. <b>cat &lt; /dev/audio &gt; file</b>,
or <b>oggenc &lt; /dev/audio &gt; file.ogg</b>.
However,
I have not been able to make this work in practice on my test machines,
but maybe you will have better luck then I did.
</p>

<p>
The classic way to play music in Plan 9,
was using the archaic <b>juke</b> player.
You first needed to write a fairly verbose database of your audio files though (see <b>juke</b>(7) if you really want to do this).
Thankfully,
9front recently replaced this rusty jukebox with a cool new audio player called <b>zuke</b>.
It easy and pleasant to use:
</p>

<p>
<pre>
; <i>audio/mkplist $home/music/myalbum &gt; $home/lib/plist/myalbum</i>
; <i>audio/zuke &lt; $home/lib/plist/myalbum</i>
</pre>
</p>

<p>
<a href="images/juke.png"><img src="thumbs/juke.png" title="Juke and Zuke, I recommend Zuke"></a>
</p>

<h3 id="video">Video</h3>

<p>
For the longest time there were no video playing options at all in Plan 9,
but recently a video player called <b>treason</b> has been written by the ever progressive 9front developers.
The video player has limited capabilities,
it cannot skip back and forth,
and worse,
cannot scale the video in any way.
It would have been nice if we could manually set a lower screen resolution;
by running <b>aux/vga -m vesa -l 1024x768x16</b> for instance,
before playing a 480P (DVD quality<a title="Getting an exact vesa mode for your movie resolution is unlikely, but 1024x768, 800x600 or 640x480 might be good enough." href="">*</a>) movie,
to watch it fullscreen.
But that will not work on all video cards,
such as,
notably,
my video card,
but maybe you'll have better luck.
(see the
<a href="#included_emulators">Game Emulators</a>
section for some tips)
In any case,
fullscreen or not,
you can <i>watch</i> Plan 9 in Plan 9,
how cool is that!
</p>

<h4>Install the video player Treason from the developers website:</h4>

<p>
<pre>
; <i>mkdir /tmp/treason</i>
; <i>cd /tmp/treason</i>
; <i>hget https://ftrv.se/_/treason.gz | gunzip | disk/mkext -d .</i>
; <i>./treason/install.rc</i>
; <i>treason der_film.mkv</i> # watch some vids :)
</pre>
</p>

<p>
<a href="images/treason.png"><img src="thumbs/treason.png" title="You can watch Plan 9 in Plan 9"></a>
</p>

<h4 id="youtube">Youtube</h4>

<p>
Oh yes.
With <b>treason</b> installed,
you absolutely can watch Youtube videos in Plan 9.
But first you need to install a Youtube downloader:
</p>

<p>
<pre>
# install nvi, similar to youtube-dl in unix:
; <i>cd /tmp</i>
; <i>git/clone https://git.sr.ht/~ft/nvi</i>
; <i>cd nvi</i>
; <i>mk install</i>

# make some convenient wrapper scripts:

#!/bin/rc
# ytaudio - play audio only from youtube video
nvi -A 251 -a /fd/1 $1 |
mcfs -t audio |
audio/opusdec > /dev/audio

#!/bin/rc
# ytlow - play low quality youtube video
nvi -V 18 -v /tmp/vid.mp4 $1 &&
treason /tmp/vid.mp4 &&
rm /tmp/vid.mp4

#!/bin/rc
# ythigh - play high quality youtube video
nvi -a /tmp/audio -v /tmp/video $1 &&
treason -a /tmp/audio /tmp/video &&
rm /tmp/audio /tmp/video
</pre>
</p>

<p>
You can now play the audio track of <i>Plan 9 from Outer Space</i> with
<b>ytaudio qsb74pW7goU</b>,
watch it in low quality with <b>ytlow</b> or high quality with <b>ythigh</b>.
You can also add a
<a href="#plumbing">plumbing</a>
rule to automatically play a youtube URL,
just add this to <b>$home/lib/plumbing</b> before the <b>include basic</b> line:
</p>

<p>
<pre>
type is text
data matches 'https://(www.)?youtube[^ ]+'
plumb start window ytlow ''''$0''''
</pre>
</p>

<p>
If you now right click on <b>https://www.youtube.com/watch?v=qsb74pW7goU</b> in <b>acme</b>,
or run <b>plumb</b> with that argument
(you need to put it in quotes since '=' has special meaning for the shell),
you'll see a nice low res video of a low quality movie.
</p>

<h2 id="graphic">Graphics</h2>

<h3 id="page">Viewing Images/Documents</h3>

<p>
There is only one program to display images and documents alike,
and that is <b>page</b>.
It is a fantastic application,
despite lacking support for some documents types such as DOCX or ODT,
and poor support for others such as Epubs.
PDF's,
old Microsoft Office documents,
images and other simple formats <i>usually</i> work.
</p>

<h4 id="comic">Reading Comics</h4>

<p>
Comic books are often distributed as CBR or CBZ files,
these are just rared or zipped images,
so to read them <b>unrar</b>
(you can get <b>unrar</b> in 9fronts extra repository - <b>go</b> is a dependency)
or <b>unzip</b> the file,
and then view the extracted images in <b>page</b>:
</p>

<p>
<pre>
; <i>unzip -af voyage_to_venus_1.cbz</i>
; <i>lc voyage_to_venus_1</i>
001.jpg 018.jpg 035.jpg 052.jpg 070.jpg 087.jpg
002.jpg 019.jpg 036.jpg 053.jpg 071.jpg 088.jpg
003.jpg 020.jpg 037.jpg 054.jpg 072.jpg 089.jpg
...
# to view all of these, starting with 001.jpg:
; <i>page voyage_to_venus_1</i>
</pre>
</p>

<h3 id="paint">Creating Images</h3>

<p>
<b>paint</b> is available,
though you would be hard pressed to use it for anything but kindergarten art.
<b>resample</b>(1), <b>crop</b>(1) and <b>rotate</b>(1) on the other hand,
are useful little tools for image manipulation,
see their manpages for more information.
Another good alternative for image manipulation
(as in ImageMagick not as in PhotoShop),
is <b>pico9</b>,
available in the 9front extra repository.
It's still in the early stages of development,
but it's looking good!
</p>

<h3 id="screenshot">Taking a Screenshot</h3>

<p>
Not only is there a file in <b>/dev</b> for your window text,
but there is also a file for your window screen, <b>/dev/window</b>.
To take a screenshot of your current window,
you can run this command:
<b>cat /dev/window &gt; windowdump ; page windowdump</b>.
To take a screenshot of your entire screen,
do this:
<b>cat /dev/screen &gt; screendump</b>.
</p>

<p>
These images are saved in the native Plan 9 image format,
which of course the document/image viewer <b>page</b> can read.
But if you want to use these images on other operating systems,
you should convert them to the more popular PNG or JPEG formats:
<b>cat /dev/screen | topng &gt; sshot.png</b>
or simply
<b>tojpg &lt; /dev/window &gt; window.jpg</b>
</p>

<p>
As for taking a screenshot of a different window then the current one,
take a look at the <a href="#rio_scripting">rio scripting</a> section above.
You can also do this in a more GUI-like fashion,
if you install the 3rd party program
<a href="http://shithub.us/phil9/vshot/HEAD/info.html">vshot</a>.
</p>

<p>
<a href="images/screencast.giff"><img src="thumbs/screencast.giff" title="You can screenshot and screencast in various ways in 9front"></a>
</p>

<h4 id="screencasting">Screencasting</h4>

<p>
There is work in progress on a screen recording program for 9front,
called <b>wrec</b>,
available in the extra repository.
It can do simple screen capturing,
but doesn't record sound yet.
<i>PS:</i> I recommend recording with very few frames per second,
eg <b>wrec -f 3</b>,
for best results.
If you want to scale down a GIF and make it continually loop,
as in the above example,
you can export the file to a UNIX machine with ImageMagick installed and run:
<b>convert -delay 20 -loop 0 -resize 600 screencast.gif small.gif</b>.
</p>

<h2 id="peripherals">Peripherals</h2>

<h3 id="memstick">USB sticks</h3>

<p>
In 9front USB sticks are automatically mounted in <b>/shr</b>,
but if you need to manually mount it,
run <b>ls /dev | grep sd</b> before and after plugging in your memory stick,
to find its device name.
Supposing it's <b>sdUc59fd</b> run the following command to mount the memory stick in <b>/n/dos</b>:
<b>mount &lt;{dossrv -s} /n/dos /dev/sdUc59fd/dos</b>,
and unmount it with <b>unmount /n/dos</b>,
see <b>dossrv</b>(4) for more information.
If the device doesn't show up in <b>/dev</b> after plugging it in,
there is either some hardware/driver issue,
or the device uses a file system that isn't supported.
Traditionally only DOS and Plan 9 file systems have been supported,
but with the addition of <b>ext4srv</b> in the 9front extra repository,
it is also possible to work with Linux file systems.
</p>

<p>
NTFS (Windows file system) is not supported,
so you might need to reformat your memstick to FAT32 (DOS file system) before you can use it.
Assuming it is still called <b>sdUc59fd</b>,
you can do so like this:
</p>

<p>
<pre>
; <i>disk/fdisk /dev/sdUc59fd/data</i>
<i>p</i>      # print a table of partitions
<i>?</i>      # get help instructions
<i>d p1</i>   # delete a couple of partitions
<i>d p2</i>
<i>a p1</i>   # add a new partition
<i>1</i>      # just follow suggested size
<i>971</i>
<i>t p1</i>   # set partition type
<i>?</i>      # list available types
<i>FAT32</i>
<i>w</i>      # write and quit
<i>q</i>
; <i>disk/format -d /dev/sdUc59fd/dos</i>
</pre>
</p>

<h3 id="cds">CD/DVD/BD's</h3>

<p>
To mount an iso image in <b>/n/iso</b> run the command
<b>mount &lt;{9660srv -s &gt;[0=1]} /n/iso /path/to/your/cdrom.iso</b>.
This may look cryptic,
but it's actually very easy to work with CD/DVD/BD's in Plan 9
(see <b>cdfs</b>(4)),
the following demonstration shows how to mount an audio CD
(you only need to specify the device if it isn't <b>/dev/sdD0</b>),
play it,
and rip it:
</p>

<p>
<pre>
; <i>cdfs -d /dev/sdE1</i>
; <i>cat /mnt/cd/a* &gt; /dev/audio</i>
; <i>cp /mnt/cd/a* /tmp/songs</i>
</pre>
</p>

<p>
You might find these custom functions helpful too:
</p>

<p>
<pre>
fn mem{   mount &lt;{dossrv  -s &gt;[0=1]} /n/dos $1 }
fn iso{   mount &lt;{9660srv -s &gt;[0=1]} /n/iso $1 }
fn eject{ echo eject &gt; /mnt/cd/ctl }
fn cdfs{  /bin/cdfs -d /dev/sdE1   }
fn cddb{ # query the internet CD database
cdfs
grep aux/cddb /mnt/cd/ctl | rc
}
fn rip{  # rip a CD and convert it to ogg
cdfs
for(t in /mnt/cd/a*) audio/oggenc &lt; $t &gt; `{basename $t}^.ogg
}
</pre>
</p>

<p>
The next example shows how to burn an audio CD.
Simply change 'a' for 'd' to burn a data disk
(DVD's and Bluerays are always data disks).
The last command fixates the disk,
which is not necessary on rewritable CD's or data disks:
</p>

<p>
<pre>
; <i>cdfs -d /dev/sdE1</i>
; <i>cp /tmp/files/* /mnt/cd/wa</i>
; <i>rm /mnt/cd/wa</i>
</pre>
</p>

<h3 id="printers">Printers</h3>

<p>
Let me save you a <i>lot</i> of trouble:
put <b>LPDEST=stdout</b> in your <b>$home/lib/profile</b>,
now <b>lp</b> will print its postscript to standard output.
You can convert these PS files to PDF if you want,
then copy or email them to a Windows/UNIX machine,
and print out a hardcopy from there:
</p>

<p>
<pre>
; <i>lp doc.html | ps2pdf &gt; doc.pdf</i>
; <i>doctype doc.ms | rc | lp | ssh unixmonster 'cat | lpr'</i>
# or from drawterm (os lets you run a host command)
; <i>doctype doc.ms | rc | lp | os lpr</i>
</pre>
</p>

<h2 id="games">Games and other Fun Stuff</h2>

<a href="images/games.png"><img src="thumbs/games.png" title="It is quite possible to waste time on Plan 9"></a>

<p>
Gaming is a potentially contentious topic when it comes to computers.
Although massively popular of course,
nothing is more detrimental to productivity
(except a modern web browser perhaps).
So the trick to creating a <i>good</i> computer game,
is making it fun enough to distract you for a few minutes of recuperation,
but boring enough that it doesn't keep you from doing important work.
By this definition Plan 9 has a few "good" games.
</p>

<h3 id="included_games">Included Games</h3>

<p>
Plan 9 comes with a collection of games in <b>/bin/games</b>.
My favorites include:
</p>

<p>
<ul>
<li><b>games/sudoku</b></li>
<li><b>games/mahjongg</b></li>
<li><b>games/sokoban</b></li>
<li><b>games/mines</b> (9front only)</li>
</ul>
</p>

<h3 id="included_emulators">Included Game Emulators</h3>

<p>
In 9front you will also find a number of emulators in the <b>game</b> directory,
assuming you can get hold of a legal copy of the Mario World ROM for instance,
you can play it like so:
<b>games/snes -ax 4 mario.sfc</b>
(beware though some of these oldschool games can be dangerously fun!)
</p>

<p>
<ul>
<li><b>games/nes</b> Nintendo</li>
<li><b>games/snes</b> Super Nintendo</li>
<li><b>games/gb</b> GameBoy</li>
<li><b>games/gba</b> GameBoy Advanced</li>
<li><b>games/md</b> Sega Mega Drive</li>
<li><b>games/c64</b> Commodore 64</li>
</ul>
</p>

<p>
In my first attempts at playing these games,
the experience was not perfect.
My rusty old ThinkPad struggled to get good audio out of these games,
and as mentioned in the
<a href="#video">Video</a>
section,
my video card flat out refused to set the screen to a lower resolution.
I managed to circumvent both issues by
<a href="#sysadmin_cpu_terminals">PXE booting</a>
9front in my more powerful desktop machine.
It works great as a Plan 9 movie and gaming console,
and I have a nice little launch script for Zelda:
</p>

<p>
<pre>
#!/bin/rc
# zelda - launch zelda fullscreen
# usage: zelda
# depends: gaming "console" and zelda rom

aux/realemu
aux/vga -m vesa -l 1024x768x16
window -r 0 0 1024 768 games/snes -ax 4 $home/games/snes/zelda.sfc
while(sleep 5){
if(! ps | grep -s snes) aux/vga -m vesa -l 1920x1080x16
}
</pre>
</p>

<h3 id="extra_games">3rd Party Games</h3>

<p>
3rd party games,
or indeed software,
for Plan 9 is rare.
But there are exceptions,
some good ones are <b>2048</b>,
<b>hack9</b> and <b>snake</b>.
</p>


<h4>Install 2048 and hack9 game from the 9front extras:</h4>

<p>
<pre>
; <i>cd /tmp</i>
; <i>9fs 9front</i>
# download packages
; <i>cp /n/extra/src/2048.c .</i>
; <i>tar xzf /n/extra/src/hack9.tgz</i>

# compile and install 2048
; <i>8c 2048.c</i>
; <i>8l 2048.8</i>
; <i>mv 8.out $home/bin/386/2048</i>

# compile and install hack9
; <i>cd hack9</i>
; <i>mk install</i>
</pre>
</p>

<h4>Install snake game from Bell Labs contrib repository:</h4>

<p>
<pre>
; <i>cd /tmp</i>
; <i>9fs sources</i>   # download file
; <i>cp /n/sources/contrib/john/snake.c /tmp</i>
; <i>8c snake.c</i>    # compile and install it
; <i>8l snake.8</i>
; <i>mv 8.out $home/bin/386/snake</i>
</pre>
</p>

<p>
<a href="images/hack9.png"><img src="thumbs/hack9.png" title="It's not exactly Steam, but you do have some 3rd party games for Plan 9"></a>
</p>

<h3 id="edutainment">Edutainment</h3>

<p>
There are a few educational applications in Plan 9,
such as <b>scat</b>,
<b>map</b> and <b>graph</b> for drawing star charts,
maps and graphs
(all of which use <b>plot</b> to actually draw the graphics).
Many of the programs in <b>/bin/games</b> are also more edutainment then actual games.
This includes simulators such as <b>life</b>,
<b>galaxy</b> and <b>timmy</b>.
You also have some very computer science nerdy "games" such as <b>blit</b> (more on this later) and <b>mix</b>
(last four examples are 9front specific).
</p>

<h4 id="edutainment_arithmetic">Arithmetic</h4>

<p>
The classic bsdgames collection provides UNIX with many simple edutainment programs,
most of which are not available in Plan 9.
No matter,
we can just write them from scratch.
Or at least,
I can demonstrate how to implement some of the basic ones.
Who knows,
maybe these simple scripts will inspire you to write an actually entertaining game yourself :^)
</p>

<p>
<pre>
#!/bin/rc
# arithmetic - basic arithmetic quiz
# usage: arithmetic [-q n][-r n][-o '+-*/%^']

# set some default values
rfork e
right=0
wrong=0
questions=20
range=1            # in digits
operands='+-'
start=`{date -n}

# parse optional flags
for(i in $*){
switch($i){
case -q
questions=$2 &amp;&amp; shift 2
case -r
range=$2 &amp;&amp; shift 2
case -o
operands=''$2'' &amp;&amp; shift 2
}
}

# rnum: generate a random digit based on cpu clock
fn rnum{
awk '{ print $2 }' /dev/time | sed 's/.*(.)$/\1/'
}

# ask math questions
opleft=$operands
for(i in `{seq $questions}){
# generate random math puzzle
a=`{rnum}
b=`{rnum}
if(test $range -gt 1){
for(i in `{seq `{echo $range - 1 | bc}}){
a=`{echo $a^`{rnum}}
b=`{echo $b^`{rnum}}
}
}
if(~ $#opleft 0) opleft=$operands
if not{
opused=`{echo $opleft | sed 's/^(.).*/\1/'}
opleft=`{echo $opleft | sed 's/^.(.*)/\1/'}
}
echo $a $"opused $b
correct=`{ echo $a $"opused $b | bc }
answer=`{ read }

# evaluate given answer
while(! ~ $answer $correct){
if(echo $answer | grep -s '^[-+]?[0-9]+$'){
echo What?
wrong=`{ echo $wrong + 1 | bc }
}
if not echo Please type a number '(no decimals!)'
answer=`{ read }
}
echo Right!
right=`{ echo $right + 1 | bc }
}

# print result of math quiz
finish=`{date -n}
time=`{echo $finish - $start | bc}
total=`{echo $right + $wrong | bc}
timepq=`{echo $time / $total | bc}
if(~ $right 0){ prct=0% }
if not prct=`{ echo 'scale=2 ; '$right' / '$total'' |
bc | sed 's/\.//' | sed 's/$/%/' }
echo -n $right right, $wrong wrong '('$prct' correct)'
in $time seconds '('$timepq's per answer)'
</pre>
</p>

<h4 id="edutainment_quiz">Quiz</h4>

<p>
<b>quiz</b> is another simple classic from bsdgames,
it just asks you a bunch of questions and keeps track of your progress.
Originally the UNIX <b>quiz</b> programs could ask you some fairly dated questions about geography,
Star Trek or the <b>ed</b> editor,
but the real beauty of this program is that you can write your own quiz files.
In theory you could even use this program for serious purposes,
such as training vocabulary or prepping for an exam.
</p>

<p>
<pre>
#!/bin/rc
# quiz - ask questions and look for correct answers
# usage: quiz [-as][-q questions][file]
#
# bug: case is normally ignored, but not for exotic unicode characters, this is
#      a grep bug.
# bug: special characters in the correct answers must be escaped (eg. \?\!)

# variables
rfork e
ifs='
'
dir=$home/lib/quiz
is=(Correct answer is)
right=0
wrong=0
printanswer=no
if(~ $1 -a) printanswer=yes &amp;&amp; shift 1
silenterror=no
if(~ $1 -s) silenterror=yes &amp;&amp; shift 1
questions=20
if(~ $1 -q) questions=$2 &amp;&amp; shift 2

# parse args
if(~ $#* 0) ls -p $dir &amp;&amp; exit
if(~ $#* 1) file=$dir/$1
if not echo usage: quiz [-as][-q questions][file] &amp;&amp; exit
if(test `{cat $file | wc -l} -le $questions) questions=`{cat $file | wc -l}

# ask questions, and check answers
for(i in `{sed -e '/^$/d' -e '/^#/d' $file | shuf | sed ''$questions'q'}){
question=`{ echo $i | awk -F@@@ '{ print $1 }' }
if(echo $question | grep -s '^cmd ')
eval `{ echo $question | sed 's/^cmd //'}
if not echo $question
correct=`{ echo $i | awk -F@@@ '{ print $2 }' }
correct_answer=`{ echo $i | awk -F@@@ '{ print $3 }' }
if(~ $#correct_answer 0) correct_answer=$correct
answer=`{ read }
if(echo $answer | grep -si '^'$correct'$'){
if(~ $printanswer yes) echo -n Right! $"is $"correct_answer
if not echo -n Right!
right=`{ echo $right + 1 | bc }
}
if not{
if(~ $silenterror yes) echo -n Wrong!
if not echo -n Wrong! $"is $"correct_answer
wrong=`{ echo $wrong + 1 | bc }
}
read
}

# calculate results
if(~ $right 0){ prct=0% }
if not prct=`{ echo 'scale=2 ; '$right' / '$questions'' |
bc | sed 's/\.//' | sed 's/$/%/' }
if not echo $right right, $wrong wrong '('$prct' correct)'
</pre>
</p>

<p>
This program expects a plain text database in <b>$home/lib/quiz</b> with two,
optionally three,
fields separated by <b>@@@</b>.
The fields are:
question,
answer.
The correct answer can be written as a regex,
to allow for variations,
if so then a third field must also be written,
the default answer.
Here is what the end of my <b>$home/lib/quiz/capitols</b> file looks like:
</p>

<p>
<pre>
Ukraine@@@Kyiv|Kiev@@@Kyiv
United Kingdom@@@London
Uruguay@@@Montevideo
Uzbekistan@@@T[oa]shkent@@@Toshkent
Vanuatu@@@Port Vila
Venezuela@@@Caracas
Vietnam@@@Ha ?Noi@@@Ha Noi
Yemen@@@[ŞS]an'?a'?@@@Şan'a'
Zambia@@@Lusaka
Zimbabwe@@@Harare
</pre>
</p>

<h4 id="edutainment_touchtype">Touchtype</h4>

<p>
Learning to touchtype is a <i>must</i> for any serious computer user,
and even for the seasoned sysadmin it is a skill that one might want to brush up on from time to time.
There are elaborate touchtyping tutors in UNIX,
such as <b>ktouch</b>,
but the basic method of learning this skill is fairly simple:
Print out a picture of your keyboard layout and stick it to the wall,
as you type away,
look up at the picture not down at your keyboard
(ideally you should also place your fingers on the middle row,
with your index fingers on the two keys which have little bumps on them).
This is hard to do in the beginning,
but if you keep at it,
you will gradually learn to touchtype.
The following script will not take away the pain and discipline required to learn this skill,
but it can help you track your progress.
Just retype each line that you are given,
but <i>do not</i> hit backspace and correct your mistakes,
just keep on typing.
When you are done the script will tell you how well/bad your typing skills are.
</p>

<p>
<pre>
#!/bin/rc
# touchtype - check your typing speed and accuracy
# usage: touchtype [ file ]

# choose input sample
rfork ne
tmp=/tmp/touchtype-$pid
out=/tmp/touchtype-out-$pid
fortune &gt; $tmp
if(~ $#* 1) cat $1 &gt; $tmp
ifs='
'

# do some touchtyping
start=`{date -n}
for(line in `{cat $tmp}){
echo $line
read &gt;&gt; $out
}
stop=`{date -n}

# calculate results
time=`{echo $stop - $start | bc}
char=`{cat $tmp | wc -c}
speed=`{echo '('$char' / '$time') * 60' | bc}
err=`{cmp -l $out $tmp | wc -l}
if(~ $#err 0) prc=0
if not prc=`{echo 'scale=2 ; '$err' / '$char'' | bc | sed 's/\.//'}

# print results
rm $tmp $out
echo
echo 'RESULT (<2% errors and >200 c/m is good):'
echo your write speed is $speed c/m with $prc^% errors
</pre>
</p>

<h3 id="telnet_fun">Playing With Telnet</h3>

<p>
Believe it or not,
but there are actually a lot of fun stuff to be done with <b>telnet</b>,
even in 2021!
Not least of which is playing MUD's,
multi-user-dungeons are still alive and kicking.
You can find a list of popular ones on <b>http://mudconnect.com</b>.
Here are some fun <b>telnet</b> examples
(<i>PS:</i> run <b>vt</b> first for a better user experience):
</p>

<p>
<pre>
; <i>telnet discworld.starturtle.net</i> # Play the Discworld MUD
; <i>telnet towel.blinkenlights.nl</i> # Watch Star Wars IV in ASCII
; <i>telnet twenex.org</i> # Login to a shell server with a handful of TTY games
</pre>
</p>

<h3 id="misc_fun">Miscellaneous Fun</h3>

<p>
You can do a lot of fun stuff on Plan 9 that do not strictly fall into the category of "gaming".
A classic example is <b>fortune</b>,
which will display a random quote.
9front also ships with <b>troll</b> and <b>theo</b>,
which does much the same thing,
but are more specific.
Plan 9's <b>fortune</b> is also handy for printing a random line form an arbitrary file
(eg. <b>play `{fortune playlist}</b>).
Another fun program is <b>games/festoon</b>,
which generates a gibberish <b>troff</b> document,
you can for instance use it like so:
<b>games/festoon -pet | pic | eqn | tbl | troff -mm | page</b>
</p>

<p>
Some of the programs in <b>/bin/games</b> are more or less screensavers,
such as <b>juggle</b> and <b>catclock</b>.
9front also throws in <b>mole</b> and <b>packet</b>,
which fit this category.
Lastly,
there is a port of classic UNIX screensavers in the 9front extra repository,
called <b>xsr</b>.
</p>

<h3 id="os_fun">Obscure Operating Systems</h3>

<p>
We have already touch on <b>vmx</b> in
<a href="#virtualizing">the virtualizing section</a> above,
which let you run things like Linux,
and plausibly Windows,
in 9front.
But you can also run a few obscure operating systems more natively,
and these systems may be of special interest,
and provide a lot of fun,
for a Plan 9 fan:
</p>

<h4 id="inferno">Inferno</h4>

<p>
The Inferno project was started a few years after Plan 9 was initially released,
and it was more or less developed in tandem at Bell Labs,
with the same group of developers.
The operating system share much in common with Plan 9,
you will find <b>acme</b> and other similar programs,
and it shares the exact same file system protocol
(although it is referred to as <b>styx</b>,
not <b>9p</b>,
in the docs for historic reasons).
Since everything in Inferno is a file,
you can seamlessly share devices and other resources between it and Plan 9.
</p>

<p>
This last point is especially valuable,
because Inferno was designed to run on top of other operating systems.
It can run on virtually any (old) UNIX system,
Plan 9 of course,
and even in old Internet Explorers!
Inferno presents the network,
audio, memory etc. of these systems as regular files,
and thus provide an elegant bridge between a Plan 9 system and,
say,
a Linux or FreeBSD box.
It can also run in as little as 256 Kb of memory,
<i>a quarter</i> of a Megabyte!
So it is well suited for embedded applications.
</p>

<p>
Sadly though,
Inferno suffers badly from neglect and code rot.
Audio will not work today,
and with it,
any of the multimedia applications that Inferno provides.
Worse,
Inferno is only supported on 32-bit systems,
and it's getting increasingly difficult to even build it on modern systems.
Yet,
there is a ray of hope:
The 9front developers have recently started hacking away at the Inferno code,
and created the forks <b>purgatorio</b> and <b>9ferno</b>.
<b>9ferno</b> does actually build on amd64 Linux and 9front
(use "9front", not "Plan9", as the SYSHOST).
It is still early days,
but it's fun to see that this old project finally gets some love!
You can install and run the original Inferno in Plan 9 (32-bit) like so:
</p>

<p>
<pre>
; <i>cd $home/src</i>
; <i>git/clone https://bitbucket.org/inferno-os/inferno-os</i>
; <i>mkdir /usr/inferno</i>
; <i>dircp inferno-os /usr/inferno</i>
; <i>cd /usr/inferno</i>
; <i>path=(/usr/inferno/Plan9/386/bin $path)</i>
; <i>mk install</i>

# install a new user
; <i>mkdir tmp</i>
; <i>mkdir usr/$myuser</i>
; <i>dircp usr/inferno usr/$myuser</i>

# run inferno and start a desktop
; <i>emu</i>
; <i>wm/wm</i>

# to get purgatorio or 9ferno forks
; <i>cd $home/src</i>
; <i>git/clone gits://git.9front.org/plan9front/purgatorio</i>
; <i>git/clone gits://git.9front.org/plan9front/9ferno</i>
</pre>
</p>

<p>
<a href="images/inferno.png"><img src="thumbs/inferno.png" title="Inferno - if Plan 9 were Windows"></a>
</p>

<p>
Be sure to read the papers in the <b>doc</b> directory here,
especially <b>bltj.pdf</b>,
<b>sh.pdf</b>,
<b>descent/descent.pdf</b>,
and
<b>limbotk/tk.pdf</b>,
which introduces the Inferno operating system,
its <b>rc</b>-inspired shell,
its unique programming language,
Limbo,
and the Tk GUI toolkit for it.
Inferno was written in an entirely new programming language,
Limbo,
a precursor to Go.
Unlike Plan 9,
its approach to GUI's is also much closer to traditional systems.
So if you have experience with Tk,
or really any other toolkit in UNIX or Windows,
you will find it quite easy to develop graphical programs in Inferno.
Btw,
the default startup menu is quite scarce,
but you will find many additional GUI programs under <b>/dis/wm</b>,
and you can modify the startup menu configuration file in <b>/lib/wmsetup</b>.
</p>

<p>
Inferno was intended as a commercial product,
and it has a sort of Windows'y feel to it.
And yet,
despite deep differences,
it is very reminiscent of Plan 9.
It is an interesting blend,
and a fun programming environment.
But be prepared for bugs and limitations,
the project has been quite dead for a long time
(in contrast to Plan 9, which is quite undead).
</p>

<p>
<i>PS:</i> You can get around many of the limitations in Inferno with the <b>os</b> command,
it lets you execute a host program from within Inferno.
For example,
Inferno does not include <b>awk</b>,
<b>tar</b> or <b>lp</b> (<b>lpr</b> in UNIX),
but you can easily write wrapper functions that use these host commands.
You might also want to add some startup shortcuts to your local <b>$home/lib/wmsetup</b>.
(shells will only use <b>$home/lib/profile</b> with the <b>-l</b> (<i>ell</i>) flag)
</p>

<p>
<pre>
# inferno startup shortcuts for plan 9, adjust to suit your needs:
EMU=(-g1600x900 -C x8r8g8b8 -f /fonts/vera/veramono/veramono.12.font -c1)
fn inferno{ /usr/inferno/Plan9/386/bin/emu wm/wm wm/logon -u myuser }

# to halt inferno, run this in an inferno shell
; <i>shutdown -h</i>

# adding awk to inferno (do this within inferno):
; <i>mkdir $home/dis</i>
; <i>echo bind -b $home/dis /dis &gt;&gt; $home/namespace</i>
; <i>touch $home/dis/awk</i>
; <i>chmod +x $home/dis/awk</i>

# then, you can add this to $home/dis/awk:

#!/dis/sh
# awk - a wrapper for awk on host
# usage: awk '{ cmd... }'
# bugs:  does not support awk flags

if {~ $#* 0} { file = /fd/0 } { file = $2 }
os -d $emuroot^`{pwd} awk $1 $file
</pre>
</p>

<h4 id="unix_v8">UNIX V8</h4>

<p>
In the late 80's,
the designers of UNIX continued to work on their operating system,
and developed Research UNIX Version 8 through 10,
before they went on to develop Plan 9.
You can see the prototypes of many Plan 9 ideas in these early UNIX systems.
For example,
<b>mux</b>, <b>jim</b> and <b>face</b>,
are essentially the prototypes for <b>rio</b><a title="Mux for multiplexer, is similar to 8½, by the 3rd release of Plan 9, nicknamed Brazil, from the dystopian movie by that name, the window manager was given colors and renamed rio." href="">*</a>,
<b>sam</b><a title="Sam was Jim Kirks older brother, featured in the episode Operation Annihilate in Star Trek the original series." href="">*</a>
and <b>faces</b>
in Plan 9,
you will find early versions of <b>plot</b> and <b>proof</b> too.
You can run early editions of UNIX with the SIMH emulators,
using the <b>vax780</b> emulator for V8 and the BSD's,
and the <b>pdp11</b> emulator for the earliest editions of UNIX.
To install and run V8:
</p>

<p>
<pre>
# first, install the simh emulators
; <i>cd /tmp</i>
; <i>9fs 9front</i>
; <i>tar xzf /n/extra/src/simh.tgz</i>
; <i>cd simh</i>
; <i>plan9/build_all</i>
; <i>mkdir $home/bin/$objtype/simh</i>
; <i>dircp BIN $home/bin/$objtype/simh</i>

# download v8 and run it
; <i>hget http://9legacy.org/download/unix/v8-simh.tar.bz2 | bunzip2 | tar x</i>
; <i>cd v8-simh</i>
; <i>vt</i>
; <i>simh/vax780 v8.ini</i>
login: <i>root</i>
</pre>
</p>

<p>
<a href="images/blit.png"><img src="thumbs/blit.png" title="V8 - if Plan 9 were UNIX"></a>
</p>

<p>
Using the ANSI terminal <b>vt</b> is not a hard requirement,
but it provides a better experience.
Once the server is running,
right click and choose <b>"raw"</b>.
This will prevent text from echoing twice,
and it will allow you to use key combinations,
like <b>Del</b> and <b>Ctrl-D</b>
(otherwise Plan 9 will interpret these signals).
And yes,
early UNIX's used <b>Delete</b> to kill a process,
just like Plan 9 does.
After halting the system
(see notes below),
right click and choose <b>"cooked"</b>.
When you now hit the <b>Delete</b> key,
Plan 9 will stop the VAX emulator.
Lets add a new user to our V8 system:
</p>

<p>
<pre>
# install a new user

# <i>echo myuser::8:4:mh1092,m069:/usr/myuser: &gt;&gt; /etc/passwd</i>
# <i>mkdir /usr/myuser</i>
# <i>/etc/chown myuser /usr/myuser</i>
# <i>exit</i>
login: <i>myuser</i>

# set up your environment

$ <i>cat &lt;&lt; eof &gt; .profile
TERM=blit # or vt100
export TERM
PATH=$PATH:/etc:/usr/games:/usr/blit/bin:$HOME/bin
export PATH
eof</i>

# halt the system - preferred way (old v7 style also works)

$ <i>su</i>
# kill 1
# /etc/umount -a # v7 style: sync; sync; sync
# /etc/halt

# you can now safely kill the vax780 emulator
</pre>
</p>

<p>
As you can see,
the system is quite similar to Plan 9 in its simplistic approach to user management and shutdown procedures.
We will get back to the <b>TERM</b> value later,
but basically,
if you plan on connecting to V8 with <b>vt</b>,
or a UNIX terminal,
use <b>vt100</b>.
And this is the value you want to set,
if you are running V7 or one of the BSD's in SIMH.
Setting the <b>TERM</b> value will allow you to use programs like <b>vi</b> and <b>rogue</b>.
Setting the <b>PATH</b> variable will make it easier to launch programs,
you can run <b>chown</b> for example,
rather then the more accurate <b>/etc/chown</b>.
If you don't already have it,
I highly recommend that you get
<i>The UNIX Programming Environment</i>
by Kernighan and Pike.
This is <i>the</i> book on UNIX,
whether you use V8 or a Mac or anything in between!
Also,
if you have the interest,
you can look up the abstract papers provided with UNIX Version 7 and 10,
referred to as <i>"Volume 2"</i> of the manual.
They provide some historic context and useful hints for V8.
</p>

<p>
In the olden days,
UNIX ran on a big server somewhere in the basement,
with multiple users connected to it via diskless terminals.
You can simulate this by opening up several windows and connect to the server via telnet: <b>telnet tcp!&lt;mymachine&gt;!8888</b>,
just make sure to change <b>&lt;mymachine&gt;</b> to your actual computer name
(eg. "cirno", <i>not</i> "localhost").
Since these terminals are stateless,
you don't need any shutdown procedure,
just delete the window.
The server however runs a file system,
so it should be halted with the above instructions.
</p>

<p>
As mentioned though,
V8 was meant to be a graphical system,
and it included a window manager,
graphical text editor and other pointy-clicky things.
Bell Labs created their own graphical terminal for V8,
called the Blit
(originally the Jerq,
but for some reason management had problems with that name).
To use graphical programs in V8,
you need to connect to a V8 server with a Blit terminal.
9front includes a <b>blit</b> emulator,
and you can connect it to a V8 server like so:
<b>games/blit -b 19200 -C 000000,00ff00 -t tcp!&lt;mymachine&gt;!8888</b>
(the first two flags here are optional).
You can start the window manager with <b>/usr/blit/bin/mux</b>,
or if you have set your <b>PATH</b> correctly,
just <b>mux</b>.
</p>

<p>
If you access V8 with this <b>blit</b> emulator,
you want to set the <b>TERM</b> variable to <b>blit</b>.
However,
programs such as <b>vi</b> and <b>rogue</b> will not work in <b>mux</b>.
To run such programs you first need to quit the window manager with
<b>mux exit</b>,
and then run these programs in the text terminal.
You will find some fun graphical programs under <b>/usr/blit/bin</b>,
including <b>demo pacman</b> and <b>crabs</b>.
The later spawns a bunch of tiny crabs that wonder about the screen and randomly eats chunks of your windows.
According to Rob Pike it was a favorite pun among the developers to schedule such a program to run 30 minutes into the future,
whenever some boss at Bell Labs needed to use the computer for an important meeting.
Enjoy :)
</p>

<h2 id="office">Office</h2>

<p>
<a href="images/office.png"><img src="thumbs/office.png" title="Reading and writing professional documents is quite possible in Plan 9"></a>
</p>

<p>
There are a great many office suits on most operating systems,
and other utilities besides too numerous to count.
So many are the choices in fact that it's easy to forget that "office" is just a fancy word for working with text.
Plan 9 does not delude it's users:
You need to be a proficient reader and writer to use the system,
and you need to organize and manage your files.
In other words,
you need to have essential office skills to use the system well.
</p>

<h3 id="office_reading">Reading Office Documents</h3>

<p>
As far as it's up to you,
I'm sure all of your documents are plain text as a matter of course.
Plain text is editable,
searchable,
pipeable,
programmable.
You can mangle it freely with standard tools such as <b>grep</b>,
<b>sed</b> and <b>awk</b>,
and it doesn't require a flippin Terrabyte of diskspace.
In Plan 9 text is even more powerful,
it's <i>always</i> unicode,
it's plumbable,
acmeable,
zeroxable,
yesterdayable,
snarfable and devable
(yes,
these are "real" words in Plan 9).
It's the magic goo that holds everything together,
much like in the real world.
You would be insane not to write documents as plain text!
But sadly it's not always up to you.
Your pesky boss may send you important Word documents,
with little to no regard for your peculiar taste in operating systems.
Don't panic!
Many office documents are readable with <b>page</b>
(naturally HTML files can be read with <b>mothra</b>).
Documents that aren't handled by <b>page</b>,
such as DOCX or ODT,
can easily enough be converted to PDF before importing them to your Plan 9 box
(assuming you don't run Plan 9 on <i>all</i> your machines that is).<a title="Well... you *could* spin up Ubuntu virtually with vmx, convert the office files, and move them back to 9front. But that's just ridiculous.">*</a>
</p>

<h4 id="epub">Reading Epubs</h4>

<p>
In theory,
<b>page</b> can handle Epubs,
but in my experience it can't.
Epubs are basically just zipped HTML files,
so it is possible to <b>unzip</b> them,
search around for a "toc" (table of contents) file to find what files constitute what chapters,
and then read them one by one in a web browser.
The following script automates that process:
</p>

<p>
<pre>
#!/bin/rc
# epub2html - convert epub to html
# usage: epub2html file.epub
# bugs:  only one epub at a time

# set some defaults
rfork e
cwd=`{pwd}
fn usage{
echo Usage: epub2html file.epub &gt;[1=2]
exit usage
}    
if(! ~ $#* 1) usage
file=$1
if(! ~ $file /*) file=`{cleanname $cwd/$1}
if(! test -f $file &amp;&amp; ! ~ $1 *.[Ee][Pp][Uu][Bb]) usage
name=`{basename $1 | sed 's/\.[Ee][Pp][Uu][Bb]//'}
dir=$name^_files

# determine directory name of toc file
fn ops{
ops=`{ls -p $1 | grep -i '^o.*ps'}
if(~ $#ops 0) echo $1
if not{
toc=`{ls -p $1/$ops | grep -i 'toc.ncx'}
if(~ $#toc 0) echo $1
if not echo $1/$ops
}
}

# extract epub and chapter information
mkdir -p $dir &amp;&amp; cd $dir
unzip -af $file &gt;/dev/null &gt;[2=1]
ops=`{ops $cwd/$dir} &amp;&amp; cd $ops
cat [Tt][Oo][Cc].[Nn][Cc][Xx] | sed -n '/&lt;navPoint/,/&lt;\/navPoint/p' |
sed -n 's/.*&lt;text&gt;(.*)&lt;\/text&gt;.*/\1/p' &gt; chaps
cat [Tt][Oo][Cc].[Nn][Cc][Xx] | sed -n '/&lt;navPoint/,/&lt;\/navPoint/p' |
sed -n 's/.*src="(.*)".*/\1/p' | sed 's/%20/ /g' &gt; links

# generate html index
cat &lt;&lt;eof &gt; $cwd/$name.html
&lt;!DOCTYPE html&gt;
&lt;html&gt;
&lt;head&gt;
&lt;meta charset="utf-8"&gt;
&lt;title&gt;Contents&lt;/title&gt;
&lt;/head&gt;
&lt;body&gt;
&lt;h1&gt;Contents:&lt;/h1&gt;
eof
for(i in `{seq `{cat links | wc -l}}){
link=`{sed -n $i^p links}
chap=`{sed -n $i^p chaps}
echo '        &lt;a href="'$ops^/^$"link'"&gt;'$"chap'&lt;/a&gt;&lt;br&gt;' \
&gt;&gt; $cwd/$name.html
}
cat &lt;&lt;eof &gt;&gt; $cwd/$name.html
&lt;/body&gt;
&lt;/html&gt;
eof
</pre>
</p>

<p>
This script works surprisingly well for my needs,
but I cannot guarantee that it will handle absolutely all Epubs gracefully.
Feel free to expand or adjust the script to suit your needs.
To wet your appetite,
I will add three additional scripts based on <b>epub2html</b>.
They are fairly self explanatory.
The last one,
<b>eread</b>,
is probably the most interesting;
It extracts the epub directly into private memory and reads the resulting html in <b>mothra</b>.
Once you exit the browser,
the files are discarded.
Thus it provides a fast way to read epubs without messing with temporary files on disk.
Of course,
if you aren't a lazy bum like me,
you might want to patch up <b>page</b> so that it handles Epubs correctly,
instead of monkeying about with shell scripts ;)
</p>

<p>
<pre>
#!/bin/rc
# epub2txt - convert epub to text
# usage:  epub2txt file.epub
# depend: epub2html

# set some defaults
rfork e
if(! ~ $#* 1) exit usage
keep=yes
name=`{basename $1 | sed 's/\.[Ee][Pp][Uu][Bb]//'}
if(! test -f $name.html){
keep=no
epub2html $* || exit $status
}
ifs='
'

# convert extracted epub to text
&gt; $name.txt
for(file in `{awk -F" '/&lt;a/ { print $2 }' $name.html})
html2ms &lt; $"file | deroff | fmt &gt;&gt; $name.txt
if(~ $keep no)
rm -rf $name.html $name^_files
</pre>
</p>

<p>
<pre>
#!/bin/rc
# epub2pdf - convert epub to pdf
# usage:  epub2pdf [-k] file.epub
# depend: epub2html
# bugs:   troff(1) cannot handle any and all fonts,
#         so expect to see Weinberger pinups pop up.

# set some defaults
rfork e
if(! ~ $#* 1) exit usage
keep=yes
name=`{basename $1 | sed 's/\.[Ee][Pp][Uu][Bb]//'}
if(! test -f $name.html){
keep=no
epub2html $* || exit $status
}
temp=/tmp/epub2pdf-$pid
mkdir $temp
ifs='
'

# convert extracted epub to pdf
for(file in `{awk -F" '/&lt;a/ { print $2 }' $name.html})
html2ms &lt; $"file &gt;&gt; $temp/out.ms
doctype $temp/out.ms | rc | dpost -f &gt;[2]/dev/null |\
ps2pdf '-dCompatibilityLevel=1.4' &gt; $name.pdf
rm -rf $temp
if(~ $keep no) rm -rf $name.html $name^_files
</pre>
</p>

<p>
<pre>
#!/bin/rc
# eread - read an epub directly in mothra
# usage:  eread file.epub
# depend: epub2html, mothra

rfork ne
cwd=`{pwd}
name=`{basename $1 | sed 's/\.[Ee][Pp][Uu][Bb]//'}

# extract epub to memory, then read it
ramfs -p; cd /tmp
epub2html `{cleanname $cwd/$1} || exit $status
mothra -a file://tmp/$name.html
</pre>
</p>

<h3 id="office_writing">Writing Office Documents</h3>

<p>
For all it's wondrous benefits,
plain text documents has an obvious flaw:
They don't look good.
If you need to write an article or even just a professional looking letter,
you need something a little more sophisticated then monospace fonts.
Troff is your friend
(an ancient Plan 9 port of Tex is also available,
but i recommend <b>troff</b>).
Don't be too quick to dismiss this venerable old tool!
While <b>man man</b> will print a rather unimpressive monospaced manual,
the command <b>man -t man | page</b>,
produces a much more professional looking document!
Besides the <b>man</b>(6) macros for writing manual pages,
Plan 9 also includes the <b>ms</b>(6) macros for writing generic articles and letters,
naturally with full unicode support
(a feature either missing or clunky in UNIX Troff,
not to mention Tex or DocBook).
You can also use the <b>mpictures</b>(6) macros for including images
(these must be converted to Postscript first,
eg:
<b>jpg -9t &lt; image.jpg | lp -dstdout &gt; image.ps</b>)
and the <b>html2ms</b>/<b>ms2html</b> commands for converting <b>troff</b> articles to/from HTML.
</p>

<p>
Here is a letter in <b>troff</b>
(using <b>ms</b> macros),
and a screenshot of the result:
</p>

<p>
<pre>
.DS L
To: Archduke Poggle of Geonosis
23 Insectoid Str.
Hive
103133
GEONOSIS

From: Emperor Palpatine
Imperial Palace
P0 000001
Senate District
CORUSCANT
.DE
.SH
Dear Archduke
.PP
The so called
.I
undefeatable
.R
Death Star was blown to bits by a bunch of teenagers yesterday.
I must say I am disappointed!
We need to construct a new planet killer ASAP,
and this time lets try to avoid an
.B
Achilles heel
.R
in our design shall we?
.PP
I have some other ideas for further improvements.
First of all we need a
.I
menacing
.R
throne room with a view...
.DS
Yours truly,
Palpatine
.DE
</pre>
</p>

<p>
<img src="thumbs/letter.png" title="important matters of state should always be written in troff!">
</p>

<p>
<b>troff</b> syntax is very simple,
add a <b>troff</b> or macro command,
such as <b>.SH</b> for a section header or <b>.PP</b> for a paragraph on a line by itself,
then the text content after it.
Technically these are <b>ms</b> macro commands,
which differ slightly from <b>man</b> macros,
low level <b>troff</b> commands are written in lower case
(eg. <b>.br</b> or <b>.bp</b> to force a line break or begin a new page).
You can also write certain inline <b>troff</b> commands if you need to
(eg. <b>\fI</b>italics<b>\fR</b>roman<b>\f(CW</b>constant-width fonts).
But you don't need to know all that if you just want to write a simple letter,
in fact <b>.SH</b> and <b>.PP</b> will suffice,
but see <b>ms</b>(6) if you are thirsty for more.
</p>

<p>
Make no mistake,
<b>troff</b> can be used to write highly professional documents.
The <a href="#development">development</a> section mentioned Francisco J. Ballesteros excellent book on Plan 9,
it is worth mentioning that this book was written in Plan 9 using <b>troff</b>.
A less professional,
but perhaps still useful example,
is my article on
<a href="../2020/complexity.pdf">Operating System Complexity</a>.
You can compare this PDF with the
<a href="../2020/complexity.ms">ms source code</a>
for a taste of what writing a <b>troff</b> document looks like.
</p>

<p>
One issue when working with <b>troff</b> is that you need to use a plethora of different <b>troff</b> preprocessors,
macro packages and what not,
in order to compile the source into a useful document.
Run <b>doctype myfile.ms</b> to see what commands are needed to convert the file into pure <b>troff</b>,
this can then be read in <b>page</b> or converted to useful formats.
To illustrate:
</p>

<p>
<pre>
; <i>doctype myfile.ms</i>
tbl myfile.ms | troff -ms -mpictures
; <i>doctype myfile.ms | rc | page</i>
; <i>doctype myfile.ms | rc | dpost &gt; myfile.ps</i>    # ei. postscript
; <i>doctype myfile.ms | rc | dpost | ssh unixmachine 'lpr'</i>
; <i>doctype myfile.ms | rc | dpost | ps2pdf &gt; myfile.pdf</i>
; <i>tbl myfile.ms | nroff -ms &gt; myfile.txt</i>
; <i>ms2html &lt; myfile.ms &gt; myfile.html</i>
</pre>
</p>

<p>
Depending on your document,
some of these conversions may not work very well.
Plain text and HTML conversions are often quite bad,
but Postscript and PDF should work well.
If you work a lot on <b>troff</b> documents you may find it useful to create some shortcuts,
for example:
</p>

<p>
<pre>
fn readms{ doctype $* | rc | page }
fn ms2pdf{ doctype $* | rc | dpost | ps2pdf '-dCompatibilityLevel=1.4' &gt; out.pdf }
</pre>
</p>

<h4 id="office_troff">Tweaking Troff Macros</h4>

<p>
Now,
to be clear,
you <i>don't</i> want to write your documents in pure <b>troff</b>.
Friendly macro packages like <b>ms</b> are there for a reason!
Nevertheless,
there are times when you actually need a more hands on approach.
I regularly print and send personal letters for instance.
The default <b>ms</b> documents have a very small font with wide margins.
To make matters worse,
the macro is not calibrated for an A4 paper size,
which is ubiquitous here in Europe.
The net result is that my <b>ms</b> letters have barely readable fonts with magnanimous margins all around.
Besides,
I prefer Helvetica over the Times font,
so lets change things up a bit.
Here is a very basic macro package that does the job:
</p>

<p>
<pre>
.\" Basic A4 troff macros for personal letters
.\" automatically add margins at top and bottom of page
.de hd
'sp 0.4i
..
.de fo
'bp
..
.wh 0 hd      \" run hd at start of page
.wh -0.6i fo  \" run fo at bottom -0.6 inches

.\" set A4 paper size (8.3i x 11.7i), margins and text size,
.\" and redefine standard fonts to use Helvetica
.pl 11.7i     \" page length
.ll 7.4i      \" line length (8.3 - 0.4 - 0.5 (right margin))
.po 0.4i      \" page offset (left margin)
.ps +2        \" point size  (text size)
.vs +1        \" vertical (line) space
.fp 1 R H     \" redefine font 1 (regular) roman
.fp 2 I HI
.fp 3 B HB

.\" add user friendly macros
.de SH        \" section header
.ft 3
.ps +4
..
.de PP        \" paragraph
.ps 12        \" reset text size and type
.ft 1
.sp 1         \" (vertical) space
.ti 2         \" temporary indent
..
.de R         \" basic font macros
.ft R
..
.de I
.ft I
..
.de B
.ft B
..
</pre>
</p>

<p>
As you can see,
a <b>troff</b> macro command is defined within a <b>.de CMD</b>, <b>..</b> block,
and comments begin with <b>\"</b>.
You can read the <b>troff</b> paper in <b>/sys/doc/troff.ps</b>,
to learn more of whats going on.
And you might want to look at the <b>ms</b> macro package in <b>/sys/lib/tmac/tmac.s</b>,
and perhaps add a few more commands to this bare bones example.
Plan 9 <b>troff</b> comes with only a limited number of installed fonts,
peek at <b>/sys/lib/troff/font/devutf/shell.lib</b> to look at your options.
If you need custom fonts,
colors and URL links in your documents,
you really need to bite the bullet and use GNU troff.
<a href="https://www.youtube.com/@GavinFreeborn/videos">Gavin Freeborn</a> has some nice youtube videos on the subject,
if you're interested.
</p>

<p>
This macro package uses slightly larger fonts then <b>ms</b>,
which has a default point size of 10,
and it hugs the corners of the paper much closer
(about a centimeter of margins on all four sides).
To use the macros in my documents,
I can add <b>.so mymacro.a4</b> to the top of my letter,
and then read it with <b>troff myletter.a4 | page</b>.
But since I'll use it frequently,
it is more convenient to install it system wide:
<b>cp mymacro.a4 /sys/lib/tmac/tmac.a4</b>.
I can now drop the <b>.so mymacro.a4</b> line,
and compile my letter like so:
<b>troff -ma4 myletter.a4 | dpost | ps2pdf '-sPAPERSIZE=a4' > myletter.pdf</b>.
(note that Ghostscript, used by <b>ps2pdf</b>, needs to know the paper size as well as <b>troff</b>)
</p>

<h3 id="office_spell">Spellchecking</h3>

<p>
The spellchecker <b>spell</b>(1),
and the <b>acme</b> equivalent <b>aspell</b>,
is a simple but useful tool for spellchecking English text
(sadly it does not support user supplied dictionaries).
Speaking of which,
<b>dict</b>(7) is an excellent English dictionary,
somewhat equivalent to WordNet in UNIX.
To use this tool you need to install some files first,
see the <b>README</b>'s in <b>/lib/dict</b> for instructions.
</p>

<p>
There is precious little support for non-English languages in any operating system,
but you can use various strategies for spell checking at least,
as an example consider these functions for spell checking Norwegian:
</p>

<p>
<pre>
fn lower{
tr A-ZÆØÅ a-zæøå
}
fn words{
tr -c 'a-zæøåA-ZÆØÅ''' '
' | sed 's/''//g' | sort | uniq
}

temp=/tmp/dict-$pid
dict=/lib/words.no        # Norwegian dictionary
lodict=`{basename $dict}  # Local Norwegian dictionary
fn nolook{
look $* $dict
}
fn nospell{
if(test -f $lodict) dict=$lodict
for(word in `{deroff $* | lower | words | comm -13 $dict -})
if(! grep -s '^'$word'$' $dict) echo $word
}
fn noaddword{
if(test -f $lodict) dict=$lodict
for(word in $*) echo $word &gt;&gt; $dict
words &lt; $dict &gt; $temp &amp;&amp; mv $temp $dict
}
fn nomkdict{
comm -12 &lt;{deroff $* | lower | words} $dict &gt;&gt; $lodict
for(word in `{deroff $* | lower | words | comm -13 $lodict -})
if(! grep -s '^'$word'$' $dict) echo $word &gt;&gt; $lodict
words &lt; $lodict &gt; $temp &amp;&amp; mv $temp $lodict
}

</pre>
</p>

<p>
These functions require you to have a dictionary of correctly spelled Norwegian words in <b>/lib/words.no</b>.
Assuming you have a UNIX machine nearby with the Norwegian wordlist for <b>aspell</b> installed,
you can import the dictionary like so:
<b>ssh myunixpc 'aspell -d nb dump master | aspell -l nb expand' | tcs -f 8859-4 | sort &gt; /lib/words.no</b>
(change <b>"nb"</b> here if you need another language,
eg. <b>"fr"</b> for French).
The <b>lower</b> and <b>words</b> shorthands take the special Norwegian letters æøå into account.
<b>nolook</b> is just a shorthand for Norwegian <b>look</b>(1).
</p>

<p>
Much like <b>spell</b>,
<b>nospell</b> breaks up your document into individual,
unique words stripped of any <b>troff</b> syntax,
and prints any word not found in the dictionary.
(unfortunately <b>comm</b> doesn't handle non-English letters well,
which is why we need an extra <b>grep</b> line to catch words that contain the Norwegian letters æøå)
To add custom words to the dictionary,
use <b>noaddword</b>.
You'll note though that <b>nospell</b> will use a local dictionary file,
if it exists.
Run <b>nomkdict *.ms</b> to populate such a local dictionary,
<b>~/dict.no</b>,
with words matched in the global dictionary,
<b>/lib/dict.no</b>.
You can now freely <b>noaddword</b>'s to the custom list,
without effecting the system dictionary,
and spellchecks will be hundreds of times faster,
since the local dictionary is honed to the vocabulary of your project files.
</p>

<p>
These custom tools are crude,
in particular they do not handle suffixes/prefixes,
so you need a large global dictionary before they become useful.
For instance,
the document you are reading now contains some 4113 unique English words.
<b>spell</b> will flag 1053 of them as spelling errors.<a title="Technical papers will have a *lot* of false positives, underscoring the wisdom of using per-file dictionaries." href="">*</a>
If you use the above strategy coupled with the default dictionary in <b>/lib/words</b>,
containing some 30,000 words,
you will get a whooping 2174 errors.
Using the English <b>aspell</b> dictionary however,
containing some 120,000 words,     
you will only get 853 errors
(the default Plan 9 dictionary intentionally omits suffixes/prefixes).
Of course all of these errors are false positives.
(I hope!)
</p>

<p>
By comparison LibreOffice will give you 828 unique false positives,
which is about as lousy.
The spellchecking mechanics of this massive office suit is certainly more attractive then our crude shell script,
but is it necessarily "better"?
Does it improve your spelling skills to right click in a GUI a thousand times,
rather then manually retyping the correct words one by one?
How easy is it to customize the tool and adapt it to your peculiar idiosyncrasies?
Even with today's <i>impossibly</i> fast computers,
LibreOffice can lag for a minute or two as you correct a false positive by clicking "Ignore All" in a large document.
This office suit is a million times more complex then our tiny shell script
(literally),
but is it a million times better?
</p>

<p>
All of these solutions are unsatisfactory,
but that's life in a nutshell.
The English language being what it is,
an intelligent spellchecker is science fiction tantamount to strong AI.
Our exercise might teach us some additional life lessons too:
1) Simple solutions are good enough,
2) Computers cannot fix the human condition,
3) The life of a writer is tedium.
If you want to take a stab at writing a better spellchecker though,
I recommend ch. 13 in Programming Pearls (Bentley) and ch. 12 in Classic Shell Scripting (Robbins &amp; Beebe).
</p>

<h3 id="office_pim">PIM</h3>

<p>
"PIM" is just a fancy acronym for getting organized.
My former work place was a disorganized disaster zone with half a dozen "professional" project management solutions in place.
Every so often my colleagues would be frustrated enough with the mess that they introduced a new project management tool,
which naturally aggravated the situation further.
The moral?
Software cannot magically clean up your mess,
only you can organize yourself.
</p>

<p>
Plan 9 does not pretend to be your nanny,
but it does give you basic tools that you can use to get yourself organized.
Such tools include
<b>date</b> and <b>cal</b> to keep track of time,
<b>calendar</b> and <b>tel</b> to keep track of appointments and contacts,
and <b>cron</b> to schedule execution of programs
(it requires a CPU+AUTH server).
And with just a little bit of <b>awk</b> it's easy to create your own PIM tools.
We will take a look at a few examples here.
The following scripts are intentionally basic,
likely they will not suit your needs exactly,
but hopefully they can inspire you to write tools that will!
</p>

<h4 id="pim_2do">2do Lists</h4>

<p>
First off,
lets create a simple 2do list manager:
</p>

<p>
<pre>
#!/bin/rc
# 2do - simple 2do list manager
# usage: 2do [list [ id... | task... ]]
# bugs:  a task cannot begin with a number

# set some defaults
rfork ne
dir=$home/lib/2do
mkdir -p $dir
tmp=/tmp/2do-$pid
date=`{date -i}

# parse arguments
if (~ $#* 0) ls -p $dir &amp;&amp; exit
if (~ $#* 1){ grep -v '^#' $dir/$1 | sort -k 2; exit }
list=$1 ; shift
id=1    # id is either 1 or one more then the highest id
if (test -f $dir/$list)
id=`{awk '{ if($1 &gt; id) id=$1 } END { print id+1 }' $dir/$list}

# id: remove tasks; task: add it
if (echo $* | grep -s '^[0-9\ ]+$'){
for (id in $*)
sed '/^'$id' /s/^/#/' $dir/$list &gt; $tmp &amp;&amp; mv $tmp $dir/$list
if not echo $id $date $* &gt;&gt; $dir/$list
</pre>
</p>

<p>
And here is a short demonstration of its usage:
</p>

<p>
<pre>
; <i>for (thing in eggs milk cheese) 2do buy $thing</i>
; <i>2do work start some project</i>
; <i>2do</i>
buy
work
; <i>2do buy</i>
3 2021-03-23 cheese
1 2021-03-23 eggs
2 2021-03-23 milk
; <i>2do buy 1 3</i>
; <i>2do buy</i>
2 2021-03-23 milk
</pre>
</p>

<p>
As you can see,
this <b>2do</b> script is very basic.
It lets you define an arbitrary number of lists that you can add tasks to,
one at a time,
and remove tasks by listing their ID numbers.
Each new task is given a unique ID and today's date,
and the tasks will be listed from oldest to newest.
To remove a list completely just run <b>rm $home/lib/2do/mylist</b>,
and you can of course edit the <b>2do</b> list manually in a text editor if you wish,
eg <b>B $home/lib/2do/mylist</b>.
</p>

<p>
The script can easily be expanded in many interesting ways,
for example you might want to add priorities and sort by priority first,
then by date.
The tasks are not actually removed,
but commented out,
so it is possible to check how many tasks have been completed since the project began and give an ETA of when the list will be completed.
Finally,
you may want to add flags that let you adjust some of the defaults here,
such as setting a date other then today.
Feel free to experiment and play with the code,
and if you have added all of these features and more,
take a step back and consider the difference between your version and the original.
Was it worth the extra complexity?
</p>

<h4 id="pim_que">Queues</h4>

<p>
Our next script is embarrassingly simple,
it's just a crude mechanism for managing a queue,
by printing the next line in a file whenever we run <b>que</b> on it.
But as we shall see,
it turns out to be surprisingly useful.
</p>

<p>
<pre>
#!/bin/rc
# que - a simple queue tracker
# usage: que [-p] file

# set some defaults
rfork ne
tmp=/tmp/que-$pid
pronly=no

# check arguments and errors
if (~ $#* 0 || test $#* -gt 2) {
echo Usage: que [-p] file &gt;[1=2]
exit usage
}
if (~ $1 -p) {
pronly=yes
file=$2
}
if not file=$1
if (! test -f $file){
echo Error: File $file does not exist! &gt;[1=2]
exit nofile
}

# print task and update queue
if (! task=`{grep -n '&lt;--' $file | sed 's/:.*//'}) task=1
next=`{ echo $task + 1 | hoc }
prev=`{ echo $task - 1 | hoc }
if (~ $pronly yes) { sed -n '$prev'p $file; exit }
sed 's/&lt;--//' $file &gt; $tmp
sed -n ''$task'p' $tmp
sed ''$next's/$/&lt;--/' $tmp &gt; $file
</pre>
</p>

<p>
Suppose we are listening through a Red Dwarf audio book,
and we have written a list of these chapters in <b>$home/lib/que/reddwarf</b>,
that look like this:
</p>

<p>
<pre>
/usr/glenda/music/reddwarf/ch1.mp3
/usr/glenda/music/reddwarf/ch2.mp3
/usr/glenda/music/reddwarf/ch3.mp3
...
</pre>
</p>

<p>
If we run <b>que $home/lib/que/reddwarf</b>,
it will print <b>/usr/glenda/music/reddwarf/ch1.mp3</b>,
and our list will now look like this:
</p>

<p>
<pre>
/usr/glenda/music/reddwarf/ch1.mp3
/usr/glenda/music/reddwarf/ch2.mp3&lt;--
/usr/glenda/music/reddwarf/ch3.mp3
...
</pre>
</p>

<p>
The next time we run our command,
<b>que</b> will print <b>ch2.mp3</b> and move the arrow marker to <b>ch3.mp3</b>.
It's easy to automate things further.
For example:
</p>

<p>
<pre>
; <i>fn reddwarf{ play `{que $home/lib/que/reddwarf} }</i>
; <i>reddwarf</i>    # listen to next chapter in our audiobook
; <i>du -a My_Little_Pony | awk '/mp4/ { print $2 }' | sort &gt; $home/lib/que/mlp</i>
; <i>fn mlp{ treason `{que $home/lib/que/mlp} }</i>
; <i>mlp</i>         # watch next episode of My Little Pony
</pre>
</p>

<p>
At times we may want to print our current task in the queue without advancing the marker.
For example,
I regularly attend weekly meetings and keep a list of meeting notes which look like this:
</p>

<p>
<pre>
/usr/dan/jw/litt/work/2022/mwb_E_202209_files/OEBPS/202022327.xhtml
/usr/dan/jw/litt/work/2022/mwb_E_202209_files/OEBPS/202022330.xhtml&lt;--
/usr/dan/jw/litt/work/2022/mwb_E_202209_files/OEBPS/202022332.xhtml
/usr/dan/jw/litt/work/2022/mwb_E_202209_files/OEBPS/202022334.xhtml
...
</pre>
</p>

<p>
The notes are provided by an
<a href="#epub">Epub</a>
that spans several weeks (one for each line).
I have a simple script that extracts the Epub and update my list,
which I need to run maybe three or four times a year.
At the start of each week <b>que</b> is run automatically on this list to advance the marker.
Finally I have a <b>meeting</b> script that runs <b>que -p $home/lib/meeting</b> and open the corresponding HTML notes in <b>mothra</b>.
I may need to run <b>meeting</b> several times a week,
but with this setup it will always refer to the notes for the current week.
</p>

<p>
Of course the details of this example will likely not be relevant for you,
but hopefully it can give you some ideas on how to automate your own workflow.
The weekly notes can easily be daily or monthly notes,
and they do not need to be a file.
It could be a directory of files or a script to run or what have you
(check out the <a href="#plumbing">plumbing</a> section for further tips).
</p>

<h4 id="pim_pass">Password Management</h4>

<p>
All authentication services in Plan 9 are handled by a process called <b>factotum</b>
(a "factotum" is a servant entrusted will the authority to run the masters estate on his behalf).
The idea is somewhat analogous to PAM in UNIX,
but much simpler,
yet more powerful.
No program in Plan 9,
including the kernel,
contain any authentication code whatsoever,
it's all centralized in <b>factotum</b>.
This process should already be running,
but if not you can start it with <b>auth/factotum -n</b>.
And you should add this to your <b>$home/lib/profile</b>,
so that it automatically runs at every boot.
The <b>-n</b> flag here means,
"don't look for a secstore",
more on that later.
You can have more then one instance of <b>factotum</b> running,
just as you can have multiple instance of <b>plumber</b>,
in case you need to isolate some authentication service from the rest of the system.
</p>

<p>
Management of the authentication service is quite easy.
To illustrate:
when logging into a UNIX machine with <b>ssh</b> for the first time,
<b>factotum</b> will notice that it doesn't have the needed key,
and it will duly prompt you for it,
and save the key safely.
Subsequent <b>ssh</b> commands will not ask for a password,
since the authentication service already knows what it is
(the keys will be lost after a reboot though,
but keep reading).
You can see what keys the <b>factotum</b> has stored by running
<b>cat /mnt/factotum/ctl</b>,
it may return something like this:
</p>

<p>
<pre>
key proto=pass server=unixpc service=ssh thumb=5+dUiv4yKNhWR3e+DmVu9wvgX
tu5gN3xPgApEWJGMR user=glenda !password?
</pre>
</p>

<p>
You will notice that secret information,
such as your password,
will never be printed out in plain text.
Now we could have added this key manually to factotum like so:
</p>

<p>
<pre>
; echo 'key proto=pass server=unixpc service=ssh thumb=5+dUiv4yKNhWR3e+D
mVu9wvgXtu5gN3xPgApEWJGMR user=glenda !password='my secret password''
&gt; /mnt/factotum/ctl
# to delete it, do it manually or with delkey(1)
; echo 'delkey service=ssh' &gt; /mnt/factotum/ctl
; delkey ssh | rc
</pre>
</p>

<p>
The real beauty of this service comes into play however,
once you couple it with another service,
ei. <b>secstore</b>.
Plan 9's secure store saves files in non volatile RAM using strong encryption,
and thus persist safely across reboots.
You need to set up a CPU+AUTH server to use this service,
the details on how to do this can be found in
<a href="https://fqa.9front.org/fqa7.html">section 7</a>
(<a href="https://fqa.9front.org/fqa7.html#7.4.3">7.4.3</a>
for <b>secstored</b> specifically)
of the 9front fqa.
Once a <b>secstore</b> is running,
we can write our <b>factotum</b> key database and add it to the vault:
</p>

<p>
<pre>
; ramfs -p; cd /tmp   # write our file to RAM, not to disk
; cat /mnt/factotum/ctl &gt; factotum
; sam factotum        # fill in the passwords
; cat factotum
key proto=pass server=unixpc service=ssh thumb=5+dUiv4yKNhWR3e+DmVu9wvgX
tu5gN3xPgApEWJGMR user=glenda !password='my secret password'
key proto=dp9ik dom=mydomain user=glenda !password='don''t forget me!'
; auth/secstore -p factotum
</pre>
</p>

<p>
You'll notice that we added two keys here,
one for <b>ssh</b> and a Plan 9 user account
(the <b>dom</b> value here is equivalent to <b>authdom</b> in <b>/lib/ndb/local</b>).
We can now change <b>auth/factotum -n</b> in our <b>$home/lib/profile</b> to <b>auth/factotum</b>.
During boot,
<b>factotum</b> will now open up the secure store and read any keys it finds in the encrypted <b>factotum</b> file.
To later edit this file,
just type <b>ipso factotum</b>.
</p>

<p>
You can read more about how Plan 9 security works with <b>page /sys/doc/auth.ps</b>,
but let's talk a little bit more about <b>secstore</b> before we call it quits.
The secure store can be used to encrypt any files we want,
not just the <b>factotum</b> database.
Suppose we use <b>gpg</b> to manage a list of encrypted passwords in UNIX,
and for convenience keep it around on our Plan 9 box as well.
It might look something like this:
</p>

<p>
<pre>
CATEG    NAME   USER   PASSWORD EMAIL            WEBSITE
Bank     PayPal -      123456   myuser@gmail.com paypal.com
Bank     MyBank 123456 password myuser@gmail.com mybank.no
Email    GMail  myuser MySecret myuser@gmail.com gmail.com
...
</pre>
</p>

<p>
We can then do the following:
</p>

<p>
<pre>
# put our custom database in the secret store
; <i>auth/secstore -p passwords</i>
# search the database for passwords
; <i>auth/secstore -G passwords | grep -i bank</i>
; <i>auth/secstore -G passwords | awk '/Bank/ { print $2, $4 }' </i>
# securely edit our database
; <i>ipso -e passwords</i>
</pre>
</p>

<p>
We can also safely export/import our secret database to a UNIX machine:
</p>

<p>
<pre>
# export to unix
; <i>ssh unixpc 'gpg2 --gen-key'</i>
; <i>auth/secstore -G passwords | ssh unixpc 'cat | gpg2 -ser myuser &gt; pass.gpg'</i>

# import from unix
; <i>ramfs -p; cd /tmp</i>
; <i>ssh unixpc 'gpg2 -d pass.gpg' &gt; passwords</i>
; <i>auth/secstore -p passwords</i>
</pre>
</p>

<p>
If you need to constantly import and export such files,
you can easily wrap some of these commands into more user friendly shortcuts.
But suppose we don't have a CPU+AUTH server with a <b>secstore</b> service,
can we still manage our passwords safely?
Sure:
</p>

<p>
<pre>
; <i>ramfs -p; cd /tmp</i>
; <i>B passwords</i>
; <i>auth/aescbc -e &lt; passwords &gt; $home/lib/pass.aes</i>
# and to double check that the password we typed was correct:
; <i>auth/aescbc -d &lt; $home/lib/pass.aes  &gt; /dev/null</i>
# search the encrypted file for a password
; <i>auth/aescbc -d &lt; $home/lib/pass.aes | grep -i bank</i>
</pre>
</p>

<p>
What if we have written something super secret to disk,
is there any way to safely delete the contents?
That depends.
If a copy of the file exists in the read only
<a href="#version_control">dump file system</a>,
then no.
A reinstallation of the operating system is the only way to remove the file.
But if that isn't the case,
it's simple enough to overwrite the contents with blank data:
</p>

<p>
<pre>
# ps: the whitespace in the sed command here is a tab
; <i>dd -if /dev/zero -of myfile -bs 1024 -count `{du myfile | sed 's/    .*//'}</i>
</pre>
</p>

<p>
PS: This is a joke of course,
there is no way to guarantee that data written to a modern harddisk is ever removed,
no matter what the disk may claim to your operating system.
</p>

<h4 id="pim_account">Personal Accounting</h4>

<p>
For many people the word "accounting" sends cold shivers down their spine,
and to be sure,
official business accounting tends to be horrifically complex.
But this is largely due to convoluted legislature,
and unnecessarily paranoid triple checking of the math.
For personal accounting we don't need to worry about all that.
We just need a way to quickly record our expenses,
and a way to check those expenses against a budget.
Here is a simple script that takes care of our first task:
</p>

<p>
<pre>
#!/bin/rc
# account - add records to our personal account
# usage: account [-d date] [-c catg] $$.CC [ comments... ]

# set some defaults
rfork e
account=$home/lib/account
date=`{date -i}
catg=food

# check arguments and errors
if (~ $#* 0) {
echo 'Usage: account [-c catg] $$.CC [ comments... ]'
exit usage
}
for(arg in $*){
switch($arg){
case -c
catg=$2
shift 2
case -d
date=$2
shift 2
}
}
if (echo $date | grep -vs '^[12][09][0-9][0-9]-[01][0-9]-[0-3][0-9]$') {
echo Error: invalid date, use YYYY-MM-DD &gt;[1=2]
exit wrongdate
}
if (echo $1 | grep -vs '^[0-9.]+$') {
echo Error: invalid expense, use $$.CC without prefixes &gt;[1=2]
exit wrongnumber
}

# add record to account
if (~ $catg income) amount=$1
if not amount=-$1
shift
echo $date $amount $catg $* &gt;&gt; $account
</pre>
</p>

<p>
And here is a demonstration of its use:
</p>

<p>
<pre>
; <i>account -d 2021-03-01 -c rent 1000 it sucks to pay rent</i>
; <i>account -d 2021-03-02 -c income 3500 payday!</i>
# this is too much typing, lets reduce it a bit
; <i>fn prompt{ while (echo -n '> ') eval $* `{read} }</i>
; <i>prompt account</i>
&gt; <i>-d 2021-03-04 21.25</i>
&gt; <i>-d 2021-03-06 14.50 groceries</i>
&gt; <i>-c transport 2.50 buss</i>
&gt; <i>-c other 9.50 cinema</i>
&gt; <i>11.35</i> # hit Del key to quit input loop
; <i>date -i</i>
2021-03-09
; <i>cat $home/lib/account</i>
2021-03-01 -1000 rent it sucks to pay rent
2021-03-02 3500 income payday
2021-03-04 -21.25 food
2021-03-06 -14.50 food groceries
2021-03-09 -2.50 transport buss
2021-03-09 -9.50 other cinema
2021-03-09 -11.35 food
</pre>
</p>

<p>
This demonstration illustrates that personal accounting is often quite tedious.
At least our script tries to reduce some of the work.
If we make the habit of typing in our daily expenses,
we do not have to specify a date.
Assuming that most of our expenses are <b>"food"</b> related,
we usually don't need to specify a category either.
The script allows us to give a comment to each input record,
but that is optional.
Note that we don't use + or - in our records,
the script will interpret anything with a category of <b>"income"</b> as +,
anything else as -.
Lastly,
our script requires us to type in one record at a time,
but it feels redundant to type <b>account</b> every time.
So we created a small function called <b>prompt</b> that lets us define a command,
<b>account</b> in this case.
It reads our input a line at a time,
re-evaluates our line as arguments for our command,
and executes it
(somewhat reminiscent of <b>xargs</b> in UNIX,
but with an added loop).
We quit the loop by typing the <b>Delete</b> key.
I find this trick handy in many different situations,
for example,
I might want to look up a bunch of words while writing an article,
<b>prompt look</b> or <b>prompt dict</b> does the trick nicely.
</p>

<p>
If we plan on using this database for computations,
such as summarizing our monthly expenses and checking it against a budget,
it is vital that our database contain valid data.
So we make a couple of extra sanity checks to see if the provided date and expense are correct.
Our checks are not 100% bullet proof,
but it should be good enough for personal use.
So for the next step,
the following script checks our current monthly expenses against a predefined budget:
</p>

<p>
<pre>
#!/bin/rc
# budget - measure monthly expenses against a budget
# usage: budget [YYYY-MM]

# set some defaults
rfork e
account=$home/lib/account_simple
if (~ $#* 0) date=`{date -i | sed 's/...$//'}
if not date=$1
echo $date

awk '
BEGIN {
printf("%-s\n", "-----------------------------")
}
/'$date'.* income/ { income+=$2 }
/'$date'.* rent/   { rent+=$2   }
/'$date'.* save/   { save+=$2   }
/'$date'.* food/   { food+=$2   }
/'$date'.*/        { sum+=$2    }
END {
printf("%-10s%10.2f %-10s\n", "income:", income, "of 3500")
printf("%-10s%10.2f %-10s\n", "rent:", rent, "of -1000")
printf("%-10s%10.2f %-10s\n", "save:", save, "of -200")
printf("%-10s%10.2f %-10s\n", "food:", food, "of -1000")
printf("%-10s%10.2f %-10s\n", "other:",
sum - (income + (rent + save + food)), "of -1000")
printf("%-s\n", "-----------------------------")
printf("%-10s%10.2f\n", "Balance:", sum)
}' $account
</pre>
</p>

<p>
Running the <b>budget</b> command will result in this output:
</p>

<p>
<pre>
2021-03
-----------------------------
income:      3500.00 of 3500
rent:       -1000.00 of -1000
save:           0.00 of -200
food:         -47.10 of -1000
other:        -12.00 of -1000
-----------------------------
Balance:     2440.90
</pre>
</p>

<p>
Naturally our budget here is unrealistically simple,
but it does perhaps illustrate that accounting,
at least for personal expenses,
does not have to be very difficult.
If you are more into spreadsheets and the like,
take a look at the
<a href="#office_sheet">spreadsheets</a> section below,
for an alternative approach to managing your finances.
</p>

<h4 id="pim_time">Time Management</h4>

<p>
There are many elaborate schemes and theories for time management of projects.
I will not really cover that here,
instead I will just look at the very basic tools you'll need for personal time management.
First of all the classic <b>calendar</b> program is well suited to manage your appointments.
If the date happens to be the 24 of March,
and you have a <b>$home/lib/calendar</b> file that looks like this:
</p>

<p>
<pre>
Mar 23    Finish the Plan 9 Desktop Guide already!
Mar 24    Flee the country
Mar 25    Dentist appointment
Mar 26    Go home
</pre>
</p>

<p>
Running the command <b>calendar</b> will print the following lines:
</p>

<p>
<pre>
Mar 24    Flee the country
Mar 25    Dentist appointment
</pre>
</p>

<p>
Calendar will print any appointments matching today's and tomorrow's date,
or on a Friday,
all dates up until the following Monday.
The date and the appointment have to be separated by a tab.
The trick to making this program useful,
besides actually writing down your appointments,
is to configure your system to automatically run the program every day.
Exactly how you want to do this depends greatly on your own setup and tastes,
but one simple solution is to add the following to <b>$home/bin/rc/riostart</b>:
</p>

<p>
<pre>
window rc -c 'calendar; rc'
</pre>
</p>

<p>
If you need a stopwatch,
timer or alarm clock,
the following examples may provide you with some hints:
</p>

<p>
<pre>
# hit enter to stop the clock, "r" is time in seconds
; <i>time read</i>
# set timer for 2 minutes
; <i>sleep 120; play $home/music/sample/beep.mp3</i>
# set of alarm at 17:10 o'clock
; <i>while(! ~ `{date | awk '{ print $4 }'} 17:10*)
sleep 60; play $home/music/alarm.mp3</i>
</pre>
</p>

<h3 id="office_math">Math, Graphs and Units</h3>

<p>
There are three calculators available in Plan 9: <b>bc</b>,
<b>hoc</b> and <b>dc</b>.
All of these have more or less the same capabilities,
and the old UNIX warhorse <b>bc</b> is probably the one you will be most familiar with
(run <b>bc -l</b> to use floating point math).
</p>

<p>
The <b>units</b> command is helpful for converting different units,
such as meter to feet or kilogram to pound
(it has some limitations though).
As for graphs,
one option is to use <b>graph</b>.
Suppose you have the following stock exchange printout:
</p>

<p>
<pre>
98
99
102    "102"
100
97     "97"
</pre>
</p>

<p>
The command <b>graph -y 80 120 -a &lt; stocks | plot</b> will draw a graph,
with the y axis set to vary between 80 and 120,
and the x axis set to increment automatically.
The lowest and highest points in the graph are also labeled with "97" and "102".
Of course you can make much more complicated graphs,
suppose you had three columns of numbers in the database,
one for each company you have invested in
(each optionally tagged with a label).
You can then run the command <b>graph -y 80 120 -a -o 3 -p rgb &lt; stocks | plot</b>,
to produce a graph of the three companies each with its own color
(red, green and blue).
</p>

<p>
<img src="thumbs/graph.png" title="colorful graphs with graph">
</p>

<p>
With all its capabilities,
the <b>graph</b> program has a fatal flaw:
It's clumsy to incorporate its graphs into documents.
A more elegant graph tool for <b>troff</b> documents is <b>grap</b>.
It has much the same capabilities as <b>graph</b> but uses a slightly different syntax.
To add the stock exchange graph from above in a <b>troff</b> document you could write it as follows:
</p>

<p>
<pre>
.G1
98
99
102
100
97
.G2
</pre>
</p>

<p>
And you could view the graph by running the command <b>grap stock.ms | pic | troff | page</b>.
Of course if you have a graph of three companies,
each with its own style and label,
things would become more complicated.
Supposing the plot data is in a file called <b>stocks</b>,
and looks like this:
</p>

<p>
<pre>
1 98  67 88
2 99  69 84
3 102 76 81
4 100 82 77
5 97  84 78
</pre>
</p>

<p>
You could write the grap graph like so:
</p>

<p>
<pre>
.G1
frame invis ht 2 wid 4 left solid bot solid
label left "CompA" left .5 up .7
label left "CompC" left .5
label left "CompB" left .5 down .7
draw compa solid
draw compb dotted
draw compc dashed
copy "stocks" thru X
next compa at $1,$2
next compb at $1,$3
next compc at $1,$4
X
.G2
</pre>
</p>

<p>
<img src="thumbs/grap.png" title="for integration with documents use grap">
</p>

<p>
Like the other <b>troff</b> preprocessors,
such as <b>tbl</b>,
<b>pic</b> and <b>eqn</b>,
it takes a bit of effort to learn the mini-language.
But once you get used to the semantics,
it's easy enough to add fairly advanced tables,
pictographics, math expressions and graphs to your <b>troff</b> documents.
</p>

<h3 id="office_sheet">Spreadsheets</h3>

<p>
You do not have a nice pointy-clicky GUI spreadsheet in Plan 9,
but it's not too hard to replicate the basic functionality.
Let's assume you have a habit of doing your personal accounting in LibreOffice,
and a typical fiscal year looks something like the following screenshot:
</p>

<p>
<img src="thumbs/localc.png" title="accounting in LibreOffice">
</p>

<p>
The crucial step in replicating such a report in Plan 9 is to separate data from presentation.
For instance,
lets write the variable account data in a fixed field database like so:
</p>

<p>
<pre>
Groceries   345     353     321     398     373     362
Health      134     0       0       123     0       142
Transport   262     268     273     352     263     272
Cloths      0       150     0       0       175     225
Other       363     473     481     403     428     393
</pre>
</p>

<p>
With this database in place it's fairly easy to generate the above spreadsheet.
For example, the following <b>awk</b> script will print an ASCII report similar to our LibreOffice screenshot:
</p>

<p>
<pre>
; <i>cat account</i>
#!/bin/rc
# account - print an account report
# usage: account database
# bugs:  requires a very specific input file

date=`{date}
awk &lt;$1 '
BEGIN {
# set some fixed income/expense values
income=3000; rent=1000; lone=250; savings=500; fixed=1750

# print header and fixed monthly values
printf("%10s%6s%6s%6s%6s%6s%6s\n",
"", "Jan", "Feb", "Mar", "Apr", "May", "Jun")
prfixed("Income", income)
print ""
prfixed("Rent", rent)
prfixed("Lone", lone)
prfixed("Savings", savings)
prfixed("FIXED", fixed)
print ""
}
{
# print each line in the db and save their values
prline($1, $2, $3, $4, $5, $6, $7)
jan+=$2; feb+=$3; mar+=$4; apr+=$5; may+=$6; jun+=$7
}
END {
# print summary of expenses
prline("VARIABLE", jan, feb, mar, apr, may, jun)
print ""
prline("Expenses", jan+fixed, feb+fixed, mar+fixed,
apr+fixed, may+fixed, jun+fixed)
prline("BALANCE",
income-fixed-jan, income-fixed-feb, income-fixed-mar,
income-fixed-apr, income-fixed-may, income-fixed-jun)
print ""

# print current year and annual balance
split("'$"date'", date)
printf("%10s %d\n", "Year", date[6])
printf("%10s %d\n", "SUM",
(income*6)-((fixed*6)+jan+feb+mar+apr+may+jun))
}

# a couple of wrapper functions for printf
function prline(tag, jan, feb, mar, apr, may, jun){
printf("%-10s%6d%6d%6d%6d%6d%6d\n",
tag, jan, feb, mar, apr, may, jun)
}
function prfixed(tag, n){
printf("%-10s", tag)
for (i=1; i&lt;=6; i=i+1)
printf("%6d", n)
printf("\n")
}
'
; <i>account database</i>
Jan   Feb   Mar   Apr   May   Jun
Income      3000  3000  3000  3000  3000  3000

Rent        1000  1000  1000  1000  1000  1000
Lone         250   250   250   250   250   250
Savings      500   500   500   500   500   500
FIXED       1750  1750  1750  1750  1750  1750

Groceries    345   353   321   398   373   362
Health       134     0     0   123     0   142
Transport    262   268   273   352   263   272
Cloths         0   150     0     0   175   225
Other        363   473   481   403   428   393
VARIABLE    1104  1244  1075  1276  1239  1394

Expenses    2854  2994  2825  3026  2989  3144
BALANCE      146     6   175   -26    11  -144

Year 2021
SUM 168
</pre>
</p>

<p>
If you are unfamiliar with <b>awk</b>,
I am sure the above example looks quite terrifying.
Settle down,
brew a cup of coffee,
and read the script slowly,
line by line.
The logic is fairly straight forward,
and most of the tedium here has to do with formatting.
For example <b>printf("%-10s%6d%6d%6d%6d%6d%6d\n"...)</b> doesn't look pretty,
but it makes sure that the fields are printed out nicely
(print a line consisting of a 10 character wide string,
followed by six 6 character wide digits followed by a newline).
</p>

<p>
Now it's all well and good to print ASCII tables for our own personal accounting,
but lets assume we need to incorporate such a report in a business document.
ASCII tables went out of fashion in the early 90's,
so we definitely need something more professional to show to our boss.
Don't panic,
<b>tbl</b>(1) has your back!
Consider the following example:
</p>

<p>
<pre>
; <i>cat &lt;&lt; eof &gt; table</i>
<i>.TS</i>
<i>expand center allbox;</i>
<i>l l l l l l l</i>
<i>l n n n n n n.</i>
<i>eof</i>
; <i>account database | sed 's/[ ]+/    /g' &gt;&gt; table</i>
; <i>echo .TE &gt;&gt; table</i>
; <i>tbl table | troff | page</i>
</pre>
</p>

<p>
<a href="images/table.png"><img src="thumbs/table.png" title="accounting in troff"></a>
</p>

<p>
Let's take a step back and explain what is going on here.
The <b>tbl</b>(1) program expects tab separated input fields,
so we use <b>sed</b> to convert our spaces to tabs.
Beyond that our <b>tbl</b> table must start with <b>".TS"</b> and end with <b>".TE"</b>,
and we need a short header that describes what our table should look like.
<b>expand</b>, <b>center</b> and <b>allbox</b> control various visual aspects of the table,
the next two lines state that the first row consists of seven left justified text fields,
and that all following rows after that consist of a left justified text field and six numerical fields.
Look up the <b>tbl</b>(1) manpage for more information,
you can do a lot of cool stuff with it.
To incorporate our table in an <b>ms</b> (ei. <b>troff</b>) document,
just run <b>cat table &gt;&gt; document.ms</b>.
</p>

<p>
At first glance,
our examples look very tedious,
but they are actually not much harder to work with then our LibreOffice example.
The above spreadsheet in LibreOffice consists of 550 characters.
Some of these fields contain code,
for example a field that reads "1104",
may actually be typed <b>"=SUM(B12:B17)"</b>.
Compare that to <b>awk's</b> <b>"jan+=$2"</b>.
In addition to typing in these characters,
we also need to use at least 101 mouse or keyboard actions to manipulate the table,
making a total of 651+ actions.
<p>

</p>
Our <b>awk</b> program is 982 characters,
excluding comments and whitespace,
and our database 118,
making a total of 1100 input actions.
So our <b>awk</b> table requires initially 50% more work to write then our LibreOffice table.
However,
once our <b>awk</b> program is written,
we only need to update the database when we do our accounting,
and that is five times <i>less</i> work then our LibreOffice spreadsheet.
In addition we can freely change our <b>awk</b> code without effecting the data,
we can also use our data with other programs,
we can feed it to <b>graph</b> or a database for instance.
The flexibility of our <b>awk</b> approach,
not to mention computational efficiency,
is far superior!
Proactive laziness is understandably scary for the novice,
but with experience one tends to embrace its wisdom.
</p>

<p>
But lets consider one more problem:
Writing a custom <b>tbl</b> file just to quickly view our data as a <b>troff</b> table is tedious,
can we automate this?
Sure.
Lets write a script called <b>table</b> that automatically writes a <b>tbl</b> table for the file it is given and open it in <b>page</b>:
</p>

<p>
<pre>
#!/bin/rc
# table - convert database to a tbl(1) spreadsheet
# usage: table file
# bugs:  only supports a simple generic spreadsheet

# set some defaults
rfork ne
tmp=/tmp/ttbl-$pid
mkdir -p $tmp
fn sigexit{ rm -rf $tmp }

# workaround: tbl can only handle one page (56 lines) at a time
pages=`{echo `{cat $1 | wc -l} / 56 | hoc}
if(~ $pages [0-9]*.[0-9]*){
pages=`{echo $pages | sed 's/\.*//'}
pages=`{echo $pages + 1 | hoc}
}
s=1
e=56
for(p in `{seq $pages}){
p=`{echo 00$p | sed 's/.*(...$)/\1/'}
sed -n $s,$e^p $1 &gt; $tmp/p$p
s=`{echo $s + 56 | hoc}
e=`{echo $e + 56 | hoc}
}

# generate tbl for each 56 line segment
for(file in $tmp/p*){
tbl=$file.tbl
echo .TS &gt; $tbl
echo 'expand center allbox;' &gt;&gt; $tbl

# create tbl header (header and content lines)
for(word in `{sed 1q $1 | sed 's/[ ]+/_/g'}){
if (echo $word | grep -s '^[0-9.-]+$') echo -n 'n '
if not echo -n 'l '
} &gt;&gt; $tbl
echo &gt;&gt; $tbl
for (word in `{sed -n 2p $1 | sed 's/[ ]+/_/g'}){
if (echo $word | grep -s '^[0-9.-]+$') echo -n 'n '
if not echo -n 'l '
} &gt;&gt; $tbl
echo . &gt;&gt; $tbl

cat $file &gt;&gt; $tbl
echo .TE &gt;&gt; $tbl
}

# compile all segments and print out
for (file in $tmp/p*.tbl){ tbl $file | troff &gt;&gt; $tmp/all }
page $tmp/all
exit    # force clean up
</pre>
</p>

<p>
One complication here is that <b>tbl</b> does not handle tables that overflow the page,
so we need to split very large tables into smaller chunks.
And of course our script cannot magically produce a perfect table for any and all input.
First of all it just scans the first two lines to find out what type of fields it should print,
left justified text or numbers,
it assumes that all following lines have the same fields as the 2nd line.
Lastly our input file must be a tab separated database,
if it isn't we need to transform it first
(eg. <b>sed 's/,/    /g' db.csv &gt; db.tab; table db.tab</b>).
</p>

<h3 id="office_db">Databases</h3>

<p>
"Database" is another one of those IT buzzwords,
that make really simple things sound amazingly complex.
Consider this text file:
</p>

<p>
<pre>
Asia             Japan       120     144
Asia             India       746    1267
Asia             China      1032    3705
Asia             USSR        275    8649
Europe           Germany      61      96
Europe           England      56      94
Europe           France       55     211
North America    Mexico       78     762
North America    USA         237    3615
North America    Canada       25    3852
South America    Brazil      134    3286
</pre>
</p>

<p>
Lo, and behold,
it's a database!
A database is a list of values,
nothing more.
The above table is a database of countries,
listing continent, name, population and area.
We can easily retrieve values from our database with <b>awk</b>,
for instance:
</p>

<p>
<pre>
; <i>echo Asias population is `{awk -F'    ' '
/Asia/ { sum += $3 } END { print sum }' countries}</i>
Asias population is 2173
; <i>echo Germanys population density is `{awk -F'    ' '
/Germany/ { print ($3*1000)/$4 }' countries}</i>
Germanys population density is 635.417
</pre>
</p>

<p>
Naturally these numbers are quite bogus,
since my database is incomplete,
and a bit outdated,
but I trust you get the point.
Of course when people speak of databases,
they often think of <i>relational</i> databases.
That is tables of values that are related with each other through common key values.
For example,
suppose we augment our countries database with a capital database:
</p>

<p>
<pre>
Brazil     Brasilia
Canada     Ottawa
China      Beijing
England    London
France     Paris
Germany    Bonn
India      New Delhi
Japan      Tokyo
Mexico     Mexico City
USA        Washington
USSR       Moscow
</pre>
</p>

<p>
These two databases are related with each other through the common country names,
the second field in our countries database,
and the first in our capitals database.
Lets merge them:
</p>

<p>
<pre>
; <i>sort -t'    ' -k 2 countries &gt; tmp_countries</i>
; <i>sort -t'    ' capitals &gt; tmp_capitals</i>
; <i>join -t'    ' -1 2 tmp_countries tmp_capitals</i>
Brazil     South America     134    3286    Brasilia
Canada     North America      25    3852    Ottawa
China      Asia             1032    3705    Beijing
England    Europe             56      94    London
France     Europe             55     211    Paris
Germany    Europe             61      96    Bonn
India      Asia              746    1267    New Delhi
Japan      Asia              120     144    Tokyo
Mexico     North America      78     762    Mexico City
USA        North America     237    3615    Washington
USSR       Asia              275    8649    Moscow
</pre>
</p>

<p>
You will note one complication here.
Our <b>sort</b> and <b>join</b> commands have the flag <b>-t'    '</b>
(<b>-F'    '</b> for <b>awk</b>),
that is <b>-t</b> followed by a <b>Tab</b> character surrounded by single quotes.
This is because our databases are tab separated values,
this allows us to have fields containing spaces,
such as <b>"North America"</b>.
Without the <b>-t'    '</b> flag,
this would be interpreted as two fields rather then one.
Of course we can use the same approach to work with comma separated values,
just change the flag to <b>-t,</b>.
</p>

<p>
If you try this out yourself,
you will see that we have actually cheated a bit in our examples.
Tab separated databases do not align perfectly,
they actually look more like this:
</p>

<p>
<pre>
Asia    Japan    120    144
Asia    India    746    1267
Asia    China    1032    3705
Asia    USSR    275    8649
Europe    Germany    61    96
Europe    England    56    94
Europe    France    55    211
North America    Mexico    78    762
North America    USA    237    3615
North America    Canada    25    3852
South America    Brazil    134    3286
</pre>
</p>

<p>
In UNIX it is easy to pretty print such text,
just run <b>join -t'    ' -1 2 tmp_countries tmp_capitals | column -t</b>.
Plan 9 however does not have the <b>column</b> command.
The closest equivalent,
<b>mc</b>,
does not have this auto align feature.
But it's not too hard to write an <b>awk</b> script that does the same,
here is one example:
</p>

<p>
<i>PS:</i> This script will <i>replace</i> tabs,
so don't overwrite your tab separated databases with it!
Use it for pretty printing only.
</p>

<p>
<pre>
#!/bin/rc
# column - auto align column output
# usage: column &lt; input &gt; output

cat /fd/0 | awk '
BEGIN {
FS = "\t"; blanks = sprintf("%100s", " ")
number = "^[+-]?([0-9\ ]+[.]?[0-9\ ]*|[.][0-9\ ]+)$"
}

{    row[NR] = $0
for (i = 1; i &lt;= NF; i++){
if ($i ~ number)
nwid[i] = max(nwid[i], length($i))
wid[i] = max(wid[i], length($i))
}
}

END {
for (r = 1; r &lt;= NR; r++){
n = split(row[r], d)
for (i = 1; i &lt;= n; i++){
sep = (i &lt; n) ? "    " : "\n"
if (d[i] ~ number)
printf("%" wid[i] "s%s", numjust(i, d[i]), sep)
else
printf("%-" wid[i] "s%s", d[i], sep)
}
}
}

function max(x, y){ return (x &gt; y) ? x : y }

function numjust(n, s) { # position s in field n
return s substr(blanks, 1, int((wid[n]-nwid[n])/2))
}'
</pre>
</p>

<h4 id="office_db_awk">Awk as a Database</h4>

<p>
OK, so we can merge our relational databases,
but this is still a lot of tedious work.
Can we automate this process?
And besides,
it's not so intuitive to write <b>awk '/Germany/ { print ($3*1000)/$4 }'</b>,
could we possible write <b>awk '$country == "Germany" { print ($population*1000)/$area }'</b>?
Yes.
The following script allows <b>awk</b> to query a relational database.
It only requires you to write a <b>relfile</b> first,
that describe what attributes are where.
The <b>relfile</b> must also contain a table with all available attributes.
If such a file does not exist,
it must be created,
and the instructions for doing so must be provided in the <b>relfile</b>.
</p>

<p>
<pre>
; <i>cat relfile</i>
countries:
continent
country
population
area
capitals:
country
capital
cc:
country
population
area
capital
continent
!sort -t'    ' -k 2 countries &gt; tmp_countries
!sort -t'    ' capitals &gt; tmp_capitals
!join -t'    ' -1 2 tmp_countries tmp_capitals &gt; cc
!rm tmp_* cc
; <i>cat q</i>
#!/bin/rc
# q - awk relational database query
# usage: q query
# depend: relfile

echo $* | awk '
BEGIN { readrel("relfile") }
/./   { doquery($0) }

# parse relfile
function readrel(f) {
while (getline &lt;f &gt; 0 )
if ($0 ~ /^[A-Za-z]+ *:/) {
gsub(/[^A-Za-z]+/, "", $0)
relname[++nrel] = $0
} else if ($0 ~ /^[ \t]*!/)
cmd[nrel, ++ncmd[nrel]] = substr($0,index($0,"!")+1)
else if ($0 ~ /^[ \t]*[A-Za-z]+[ \t]*$/)   # attribute
attr[nrel, $1] = ++nattr[nrel]
else if ($0 !~ /^[ \t]*$/)
print "bad line in relfile:", $0
}

# translate qawk query into corresponding awk query
function doquery(s,    i,j) {
for (i in qattr)
delete qattr[i]
query = s
while (match(s, /\$[A-Za-z]+/)) {
qattr[substr(s, RSTART+1, RLENGTH-1)] = 1
s = substr(s, RSTART+RLENGTH+1)
}
for (i = 1; i &lt;= nrel &amp;&amp; !subset(qattr, attr, i); )
i++
if (i &gt; nrel)
missing(qattr)
else {
for (j in qattr)
gsub("\\$" j, "$" attr[i,j], query)
for (j = 1; j &lt;= ncmd[i]; j++)  # create table i
if (system(cmd[i, j]) != 0) {
print "command failed, query skipped\n", cmd[i,j]
return
}
awkcmd = sprintf("awk -F''\t'' ''%s'' %s", query, relname[i])
#printf("query: %s\n", awkcmd)  # for debugging
system(awkcmd)
}
}

function subset(q, a, r,    i) { # is q a subset of a[r]?
for (i in q)
if (!((r,i) in a))
return 0
return 1
}
function missing(x,    i) {
print "no table contains all of the following attributes:"
for (i in x)
print i
}'
; <i>q '$country == "Germany" { print ($population*1000)/$area }'</i>
635.417
; <i>q '{ printf("%-10s %4.2f\n", $country, ($population*1000)/$area) }'</i>
Japan      833.33
India      588.79
China      278.54
USSR       31.80
Germany    635.42
England    595.74
France     260.66
Mexico     102.36
USA        65.56
Canada     6.49
Brazil     40.78
</pre>
</p>

<p>
Our little <b>q</b> command will likely not topple Oracle anytime soon,
but for personal use <b>awk</b> is both flexible and efficient.
By the way both this section,
and the above spreadsheet section is greatly <s>plagiarized from</s> influenced by chapter 4 in
<i>The Awk Programming Language</i>,
by Aho, Kernighan and Weinberger.
I highly recommend this book for your own personal library,
whether you use Plan 9,
UNIX or flippin DOS.
</p>

<h4 id="office_db_ndb">Ndb as a Database</h4>

<p>
Using <b>awk</b> for databases is fine,
but it's a very UNIX'y way of doing things.
Plan 9 actually has a really good,
and super fast,
database called <b>ndb</b> (network database).
You have probably already used <b>ndb</b> to set up your network configuration file <b>/lib/ndb/local</b>,
and it's for this purpose <b>ndb</b> was created.
But it is actually a fully functional generic database,
with all the trimmings.
Here is how we can implement the above country database in <b>ndb</b>:
</p>

<p>
<pre>
; <i>cat countries.db</i>
country=Japan
continent=Asia
population=120
area=144
capitol=Tokyo

country=France
continent=Europe
population=55
area=211
capitol=Paris

country=Mexico
continent="North America"
population=78
area=762
capitol="Mexico City"
...
; <i>population=`{ndb/query -f countries.db country Germany population}</i>
; <i>area=`{ndb/query -f countries.db country Germany area}</i>
; <i>echo Germanys population density is `{echo '('$population'*1000)/'$area'' | hoc}</i>
Germanys population density is 635.416666667
; <i>ndb/query -f countries.db continent Asia</i>
country=Japan continent=Asia population=120 area=144
country=India continent=Asia population=746 area=1267
country=China continent=Asia population=1032 area=3705
country=USSR continent=Asia population=275 area=8649
; <i>ndb/query -a -f countries.db continent Asia population |
awk '{ sum+=$1 } END { print "Asias population is", sum }'</i>
Asias population is 2173
</pre>
</p>

<p>
<b>ndb</b> entries are very free form,
we can write them in one line,
as in <b>country=Japan continent=Asia population=120...</b>,
in multiple lines indented by spaces or tabs,
or a combination of both.
The empty newline that separate the entries above is not required,
it's just added for the sake of readability.
Note the <b>-a</b> flag in our last <b>ndb</b> example,
without this <b>ndb</b> would return <b>120</b>,
the population of the <i>first</i> entry in continent Asia.
<b>ndb</b> can handle multiple database files,
and make attribute hashes to speed things up.
See <b>ndb</b>(8) for the full details.
</p>

<h2 id="conclusion">Conclusion</h2>

<p>
The primary value of Plan 9 lies in its simplicity.
Making a hard copy of its documentation will not break your bookshelf,
and the source code is actually readable.
This cannot be said for main stream operating systems.
Of course being such a simple system,
there are many many features that popular operating systems provide,
that Plan 9 don't.
If you plan on using this alternative OS as a daily driver,
you really have to pull up your sleeves,
learn some shell,
maybe even some C,
and write a bunch of utilities to do your work.
But it's a fascinating learning experience.
</p>

<p>
Hopefully,
this article has also demonstrated that Plan 9 does not suck quite so badly as you may have thought.
I know I was positively surprised a few times as I was writing it!
Thanks to the good work of the 9front developers,
you can run Plan 9 on modern hardware.
Most of the laptops I have tried it on have just worked,
including essential things like audio and wifi.
As we have seen,
you can do office work,
play games and audio,
work with images and the web.
And with the recent additions of a decent music player,
video player and web browser,
9front is actually starting to look pretty good even as a casual desktop.
Of course the real charm of Plan 9 has always been in its simple and consistent design,
which gives the user tremendous power with modest efforts.
Using it will likely open your mind to the power of UNIX,
much more so in fact,
then UNIX itself will.
Happy hacking :)
</p>
</body>
<!-- 
|                                          |      | 
|                                          |     _|_
_|_                                         |  __/   \__
__/   \__   Program design in the UNIX environment  '---'  \ ))
/  '---'  \                                     | '--_____--'
'--_____--'                 Rob Pike           _|_
Brian W. Kernighan   __/   \__
(( /  '---'  \
'--_____--'
ABSTRACT

Much of the power of the UNIX operating system comes from a style
of program design that makes programs easy to use and, more
important, easy to combine with other programs. This style has
been called the use of software tools, and depends more on how the
programs fit into the programming environment how they can be used
with other programs than on how they are designed internally. But
as the system has become commercially successful and has spread
widely, this style has often been compromised, to the detriment of
all users. Old programs have become encrusted with dubious
features. Newer programs are not always written with attention to
proper separation of function and design for interconnection. This
paper discusses the elements of program design, showing by example
good and bad design, and indicates some possible trends for the
future.

--
Introduction to the paper "Program design in the UNIX environment",
aka. "cat -v considered harmful",
by Rob Pike and Brian W. Kernighan.
___________________________________________________________
/                                                           \
| Greetings, my friends. We are all interested in the future,|
| for that is where you and I are going to spend the rest of |
| our lives. And remember, my friends, future events such as |
| these will affect you in the future   -  CRISWELL PREDICTS |
\___________________________________________________________/
          ,   '   ,  |
           .  '  .  /
             ,~   _/
      - - - (..) - - -
            /><\ 
_______________________________________________________'=oo='________

-->
</html>
