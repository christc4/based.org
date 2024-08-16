# Scaling and Shifting of Normal Variables

Up: [Normal Distribution Properties](normal_distribution_properties)
Brother(s):
TARGET DECK

### Scaling

If $X \sim \mathcal{N}(\mu, \sigma^2)$ then for $\alpha > 0$ and $Y = \alpha X$, $Y \sim \mathcal{N}(\alpha\mu, (\alpha\sigma)^2)$

Proof: [Quadratic Formula of Variance](quadratic_formula_of_variance)
	If $\sigma^2 = \mathbb{V}[X]$ and $\alpha > 0$ then $\mathbb{V}[\alpha X]= \alpha^2 \mathbb{V}[X]$

Example: 
	If:
	$$Y = \frac{X}{2}$$
	Then:
	$$ Y \sim \mathcal{N}(\frac{1}{2}\mu, \frac{1}{4}\sigma^2) $$
****

### Shifting

If $X \sim \mathcal{N}(\mu, \sigma^2)$ then for any $c$ and $Y = X - c$, $Y \sim \mathcal{N}(\mu - c, \sigma^2)$

Proof: [Linearity of Expected Values](linearity_of_expected_values)
	$\mathbb{E}[X-c] = \mathbb{E}[X] - c$

Example A:
	If:
	$$ Y = X - \mu $$
	Then:
	$$ Y \sim \mathcal{N}(0, \sigma^2) $$

Example B: 
	If: 
	$$ Y = \frac{X-\mu}{\sigma} $$ 
	Then:
	$$ Y \sim \mathcal{N}(0, 1) $$
	This process is known as standardisation, and ^ above is [The Standard Normal Distribution](the_standard_normal_distribution).
























#### Why:
#### How:









