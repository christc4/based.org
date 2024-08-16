# History of UNIX dot files (.)

### Original source

<p>Rob Pike discusses the history of UNIX's <code>.</code> and <code>..</code>, here:
[<s data-accessed="2012-08-03" data-defunct-date="2019-04-02">https://plus.google.com/u/0/101960720994009339267/posts/R58WgWwN9jp</s>]. May be accessible with the WayBack machine.

<p>A lesson in shortcuts.</p>

<p>Long ago, as the design of the Unix file system was being worked out, the entries <code>.</code> and <code>..</code> appeared, to make navigation easier. I'm not sure but I believe <code>..</code> went in during the Version 2 rewrite, when the file system became hierarchical (it had a very different structure early on).  When one typed ls, however, these files appeared, so either Ken or Dennis added a simple test to the program. It was in assembler then, but the code in question was equivalent to something like this:</p>

<pre class="c">
   <span class="keyword">if</span> (name[0] == <span class="string">'.'</span>) <span class="keyword">continue</span>;</pre>

<p>This statement was a little shorter than what it should have been, which is</p>

<pre class="c">
   <span class="keyword">if</span> (strcmp(name, <span class="string">"."</span>) == 0 || strcmp(name, <span class="string">".."</span>) == 0) <span class="keyword">continue</span>;</pre>

<p>but hey, it was easy.</p>

<p>Two things resulted.</p>

<p>First, a bad precedent was set. A lot of other lazy programmers introduced bugs by making the same simplification. Actual files beginning with periods are often skipped when they should be counted.</p>

<p>Second, and much worse, the idea of a “hidden” or “dot” file was created. As a consequence, more lazy programmers started dropping files into everyone's home directory. I don't have all that much stuff installed on the machine I'm using to type this, but my home directory has about a hundred dot files and I don't even know what most of them are or whether they're still needed. Every file name evaluation that goes through my home directory is slowed down by this accumulated sludge.</p>

<p>I'm pretty sure the concept of a hidden file was an unintended consequence. It was certainly a mistake.</p>

<p>How many bugs and wasted CPU cycles and instances of human frustration (not to mention bad design) have resulted from that one small shortcut about 40 years ago?</p>

<p>Keep that in mind next time you want to cut a corner in your code.</p>

<p>(For those who object that dot files serve a purpose, I don't dispute that but counter that it's the files that serve the purpose, not the convention for their names. They could just as easily be in $HOME/cfg or $HOME/lib, which is what we did in Plan 9, which had no dot files. Lessons can be learned.)</p>

