<title>Tanenbaum 2008-08-09</title>
<div class="main">
<h1>Interview with Andrew Tanenbaum</h1>
<h2>by Matthew Mikolay
in 2008 August 9<sup>th</sup></h2>

<p><b>MM</b>: When and how did you first get into computer programming and operating system design?</p>
<p><b>AT</b>: I have been programming since I was at MIT as an undergraduate. Operating system design happened much later, in the 1980s, first with Amoeba, an experimental distributed operating system, then in 1984 with MINIX.</p>
<p><b>MM</b>: What influenced you to start developing MINIX?</p>
<p><b>AT</b>: I was teaching a course using UNIX V6 and then AT&T; changed the license forbidding people from teaching it courses, the stupidest thing they could have done. They should have paid bounties to people teaching it in courses. I guess their attitude was "The fewer people who know about UNIX, the better." At that point I decided if I wanted a UNIX-like system to teach, I'd have to write one myself. So I did.</p>
<p><b>MM</b>: Have your students ever helped you in the development of MINIX?</p>
<p><b>AT</b>: In the beginning, no. I wrote V1 entirely alone. Later on, many students had ideas and wrote code. I also got funding to hire some students to write code.</p>
<p><b>MM</b>: What made you decide to make MINIX based on a microkernel rather than a monolithic kernel?</p>
<p><b>AT</b>: Good software engineering principles dictate that your programs are modular. You don't want a bug in one piece to bring down the whole thing if that can be avoided. A microkernel is much better engineered and is more modular and easier to understand. Monolithic kernels are still too big and unreliable. My metric is the TV set. The system should run for 10 years with a total of zero failures for 99.9% of the users.</p>
<p><b>MM</b>: Do you believe that there are certain drawbacks to making MINIX POSIX-compliant?</p>
<p><b>AT</b>: Not really.></p>
<p><b>MM</b>: Will MINIX ever have a windowing system besides X11, or is X11 stable and functional enough for MINIX?</p>
<p><b>AT</b>: Never say never, but X11 seems pretty good to me. I believe it is the only windowing system on Linux.</p>
<p><b>MM</b>: How well does MINIX run on dual-processor machines? Will MINIX ever be optimized for these types of computers?</p>
<p><b>AT</b>: We are just starting to work on multicore. It is MUCH harder than single core. I expect all multicore software to be riddled with errors.</p>
<p><b>MM</b>: Do you expect a lot of Linux users to switch over to MINIX?</p>
<p><b>AT</b>:  Probably not.</p>
<p><b>MM</b>: What other projects have you been working on besides MINIX?</p>
<p><b>AT</b>:  I have been involved with work on RFID security and privacy. See www.rfidvirus.org and www.rfidguardian.org.</p>
<p><b>MM</b>: What can we expect to see developed for MINIX in the future?</p>
<p><b>AT</b>:  We are adding some missing features now like virtual memory and USB support, but the focus of the research is very high reliability and self healing.</p>
<p><b>MM</b>: If Linux's Tux penguin and MINIX's raccoon faced off in a fight to the death, who would win?</p>
<p><b>AT</b>:  Raccoons are quite aggressive. Penguins are not. There would be chicken for dinner.</p>
