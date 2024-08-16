# The Normal(Gaussian) Distribution

Up: [Statistics](statistics)
Brother(s):
TARGET DECK

Normal distributions are represented by:
$$ X \sim \mathcal{N}(\mu, \sigma^2)  $$
And graphically by:
![Pasted image 20240104181516.png](pasted_image_20240104181516.png)

This means that $X$ is a random variable that is normally distributed with mean $\mu$ and variance $\sigma^2$.

The sample space for a normal distribution ($\mathcal{X}$) is $\mathbb{R}$, the whole real line.

It is calculated like so:
$$ \mathbb{P}(X = x; µ, σ^2 ) = p(x|µ, σ^2 ) = \frac{1}{\sqrt{2\pi\sigma^2}}exp(\frac{-(x-\mu)^2}{2\sigma^2}) $$
For an interval, the value is the area under the curve between the two intervals (requires integration)
Like so:
$$ \mathbb{P}(a ≤ X ≤ b) = \int_a^b p(x)dx $$
$$= \int_a^b \frac{1}{\sqrt{2\pi\sigma^2}}exp(\frac{-(x-\mu)^2}{2\sigma^2})dx $$


[Normal Distribution Properties](normal_distribution_properties)

[The Standard Normal Distribution](the_standard_normal_distribution)





























#### Why:
#### How:









