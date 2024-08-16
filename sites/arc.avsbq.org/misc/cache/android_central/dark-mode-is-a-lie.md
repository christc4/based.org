# Dark mode is a lie

> [source](https://www.androidcentral.com/phones/dark-mode-is-a-lie)

>  By Nicholas Sutrich, 2023/08/30

> It's supposedly better for your eyes, but that's not always true

The war between light and dark mode themes on phones has long relied on personal preference to determine a winner. But while dark mode has been called the more eye-friendly mode, I'm here to call its bluff, but that might be for the reason you think.

While it's true that light mode can feel like it's nuking your retinas with the unbelievably bright displays from the best phones, dark mode does the same thing in a completely different way. The crux of the problem is that too many apps and phones use what's known as "dim" dark mode instead of true "lights out" dark mode.

Dim mode might feel more eye-friendly at first but, in actuality, phones that use PWM dimming mean that this mode is worse for your eyes than either fully light or "lights out" dark mode. Some eye conditions mean dark mode is bad in any form, as Wired and Android Authority have written previously about, but I'm here to cover a different reason to toss dim mode out the window.

It's not actually darker

When dark mode first started appearing in apps and on some phones, it was often paired with the promise of battery life savings on phones with OLED screens. After all, a black pixel is completely off on an OLED. Therefore, the more black pixels you have on a display, the less often it's actually using the battery to light a pixel up.

But the problem is that many phones and apps don't have proper "lights off" dark modes. Instead, they replace white backgrounds with a dark grey — or another darker version of your favorite color via the Material You theme engine — which still keeps the display on, completely ruining the core purpose of dark mode in the first place. The X (Twitter) app above showcases all three types.

Worse yet, because most modern OLEDs use PWM flickering to fake dimming, this "dark mode" means your display is flickering more than when the phone is using light mode and could give you terrible headaches or other adverse effects, especially if you use dark mode in a dark room.

Measuring the flicker rate of each color on the display of the Samsung Galaxy Z Fold 5

![image](pwm1.webp)

(Image credit: Nicholas Sutrich / Android Central)

If you don't know what PWM dimming is, here's the briefest explainer I can come up with:

Instead of dropping the voltage to each pixel, lots of modern phones turn the display on and off several hundred times per second. By doing this, displays trick your brain into seeing a display that's brighter or dimmer, depending on how long the display stays off during this cycle.

That means your display is actually flashing at you with full-intensity brightness (2,000 nits on some phones) immediately after the display is completely off. This extreme difference in light levels is incredibly uncomfortable — and even painful — to folks sensitive to flicker. Worse yet, if you're like me, these kinds of displays and lights could eventually cause you to become sensitive to this horrible effect.

Medically, this is known as temporal light modulation, and the negative effects have been documented by studies over the years, yet many smartphone manufacturers seem oblivious to the truth.

You can capture the PWM rate of a phone by using a camera to take a picture or video at a high shutter speed of 1/6000 or higher. PWM shows up as dark lines running across the screen — as you can see in the image above — and thicker lines mean that the display is flashing more intensely per cycle.

Notice which lines are thickest above? That's right: the darkest ones.

Measuring the flicker rate of each color on the display of the Motorola Edge+ (2023) while using DC dimming/anti-flicker mode

![image](pwm2.webp)

(Image credit: Nicholas Sutrich / Android Central)

OLEDs didn't always do this and don't have to do this, but some companies have decided that displays "pop" more or appear to have better colors because of how an OLED works. In reality, this is lazy engineering, and it often makes a display cheaper to manufacture, so these companies are quite literally putting your health at risk just to save a few pennies.

On a phone that uses proper DC-dimming — like the Motorola Edge (2023) pictured above — dark colors actually emit less light instead of flashing at more extreme rates to fake a darker color. To actually fix this problem, we need more phone displays to work like this.

Until then, lights out

Until most phones can be considered PWM-friendly phones, the best stop-gap is to petition for more lights-out modes.

Some manufacturers already do this. Samsung's lights-out dark mode could help reduce the effects of PWM sickness if Samsung's displays weren't so darn intense. Unfortunately, though, even lights-out dark mode doesn't make enough of a difference on Samsung's flickering displays to help people like me who are sensitive to flickering lights.

Conversely, companies like Google are the worst offenders because they use dark backgrounds instead of proper black backgrounds. Ironically, "dark theme" in stock Android and on Pixel phones specifically states that "dark theme uses a black background to help keep your battery alive longer." It's just a shame that reality doesn't match up with what's printed on the tin.

The obvious problem is that light theme is too bright to use in a dark room, and given the nature of PWM dimming, reducing that brightness slider doesn't actually reduce the brightness. It just makes the phone flicker differently.

So, is lights-out mode the only saving grace on these devices? For now, yes. Ultimately, though, we need displays that aren't flashing at our retinas at all levels of brightness.


